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
				c.Locals("user", &(claims.User))
				return c.Next()
			}
		}
	}

	path := c.Path()
	u, _ := url.Parse(path)

	// API 외 경로(정적 파일, SPA 화면)는 인증 대상이 아님
	if !strings.HasPrefix(u.Path, "/api/") && u.Path != "/api" {
		return c.Next()
	}

	if c.Method() == "GET" && u.Path == "/api/user" {
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

	if c.Method() == "POST" && u.Path == "/api/user/findid" {
		return c.Next()
	}

	if c.Method() == "POST" && u.Path == "/api/user/password" {
		return c.Next()
	}

	if c.Method() == "POST" && u.Path == "/api/user" {
		return c.Next()
	}

	if c.Method() == "POST" && u.Path == "/api/company" {
		return c.Next()
	}

	if c.Method() == "GET" && u.Path == "/api/company" {
		return c.Next()
	}

	if c.Method() == "GET" && u.Path == "/api/department" {
		return c.Next()
	}

	if u.Path == "/api/jwt" {
		return c.Next()
	}

	log.Println("Jwt header is broken")

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"code":    "error",
		"message": "unauthorized",
	})
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
