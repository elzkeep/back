package router

import (
	"errors"
	"log"
	"net/url"
	"os"
	"strings"
	"time"
	"zkeep/models"
	"zkeep/models/user"

	"github.com/gofiber/fiber/v2"

	"github.com/dgrijalva/jwt-go/v4"
)

type AuthTokenClaims struct {
	User               models.User `json:"user"`
	jwt.StandardClaims             // 표준 토큰 Claims
}

var _secretCode string = jwtSecret()

// JWT_SECRET 환경변수가 있으면 사용, 없으면 기존 키 유지 (발급된 토큰 호환)
func jwtSecret() string {
	if v := os.Getenv("JWT_SECRET"); v != "" {
		return v
	}

	return "WkaQHd100%"
}

var JwtAuthRequired = func(c *fiber.Ctx) error {
	var token string

	if values := c.Get("Authorization"); len(values) > 0 {
		str := values

		if len(str) > 7 && str[:7] == "Bearer " {
			token = str[7:]

			claims := AuthTokenClaims{}
			key := func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("Unexpected Signing Method")
				}
				return []byte(_secretCode), nil
			}

			_, err := jwt.ParseWithClaims(token, &claims, key)
			if err == nil {
				// 토큰의 사용자 정보는 발급 시점 값이라(유효기간이 길다) 레벨·업체·상태는 DB 에서 다시 읽는다
				current, err := currentUser(claims.User.Id)
				if err != nil {
					// DB 장애를 인증 실패(401 → 클라이언트 강제 로그아웃)로 돌려주지 않는다
					return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
						"code":    "error",
						"message": "database error",
					})
				}
				if current == nil || !activeUser(current) {
					return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
						"code":    "error",
						"message": "unauthorized",
					})
				}
				c.Locals("user", current)
				return tenantGuard(c, current)
			}
		}
	}

	path := c.Path()
	u, _ := url.Parse(path)

	// API 외 경로(정적 파일, SPA 화면)는 인증 대상이 아님
	if !strings.HasPrefix(u.Path, "/api/") && u.Path != "/api" {
		return c.Next()
	}

	if c.Method() == "GET" && u.Path == "/api/webnotice" {
		return c.Next()
	}

	// 비로그인 사용자는 등록(POST)만 허용
	if c.Method() == "POST" && u.Path == "/api/webquestion" {
		return c.Next()
	}

	if c.Method() == "GET" && u.Path == "/api/web" {
		return c.Next()
	}

	// 비로그인 사용자는 등록(POST)만 허용
	if c.Method() == "POST" && u.Path == "/api/webjoin" {
		return c.Next()
	}

	// 회원가입 화면의 자격 종류·등급 선택지 (기준 정보)
	if c.Method() == "GET" && (u.Path == "/api/licensecategory" || u.Path == "/api/licenselevel") {
		return c.Next()
	}

	// 회원가입은 /api/signup/* (router/signup.go, 이 미들웨어보다 먼저 등록) — 사용자·업체 API 는 더 이상 공개하지 않는다

	if u.Path == "/api/jwt" {
		return c.Next()
	}

	log.Println("Jwt header is broken")

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"code":    "error",
		"message": "unauthorized",
	})
}

// activeUser: 사용 중이고, 가입 승인이 대기·거절이 아닌 사용자 (전체관리자는 승인 대상 아님).
// 승인 값이 비어 있는(0) 예전 계정은 승인된 것으로 본다
func activeUser(item *models.User) bool {
	if item.Status != user.StatusUse {
		return false
	}
	if item.Level == user.LevelRootadmin {
		return true
	}
	return item.Approval != user.ApprovalWait && item.Approval != user.ApprovalReject
}

// currentUser 는 DB 의 현재 사용자 정보를 읽는다. 없으면 (nil, nil), DB 오류면 err
func currentUser(id int64) (*models.User, error) {
	if id <= 0 {
		return nil, nil
	}

	conn := sharedDB()
	if conn == nil {
		return nil, errors.New("database error")
	}

	var n int
	if err := conn.QueryRow("select count(*) from user_tb where u_id = ?", id).Scan(&n); err != nil {
		return nil, err
	}
	if n == 0 {
		return nil, nil
	}

	item := models.NewUserManager(conn).Get(id)
	if item == nil {
		return nil, errors.New("user load error")
	}
	item.Passwd = ""

	return item, nil
}

func JwtAuth(c *fiber.Ctx, loginid string, passwd string) map[string]interface{} {
	conn := models.NewConnection()
	if conn == nil {
		return map[string]interface{}{
			"code":    "error",
			"message": "database error",
		}
	}
	defer conn.Close()

	manager := models.NewUserManager(conn)
	item := manager.GetByLoginid(loginid)

	if item == nil {
		return map[string]interface{}{
			"code":    "error",
			"message": "user not found",
		}
	}

	if item.Passwd != passwd {
		return map[string]interface{}{
			"code":    "error",
			"message": "wrong password",
		}
	}

	if item.Status != user.StatusUse {
		return map[string]interface{}{
			"code":    "error",
			"message": "status error",
		}
	}

	if !activeUser(item) {
		return map[string]interface{}{
			"code":    "error",
			"message": "approval error",
		}
	}

	// 토큰 payload는 누구나 디코딩할 수 있으므로 비밀번호를 넣지 않음
	item.Passwd = ""

	at := AuthTokenClaims{
		User: *item,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Hour * 24 * 365 * 10)),
		},
	}

	atoken := jwt.NewWithClaims(jwt.SigningMethodHS256, &at)
	signedAuthToken, _ := atoken.SignedString([]byte(_secretCode))

	item.Passwd = ""
	return map[string]interface{}{
		"code":  "ok",
		"token": signedAuthToken,
		"user":  item,
	}
}
