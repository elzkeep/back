package router

// 응답·요청 공통 보호 — services/http.go 에서 라우터보다 먼저 등록한다.
//   - SanitizeResponse: /api 응답 JSON 의 비밀번호 값(passwd, egiropasswd)을 비운다.
//     모델 Extra 로 붙는 사용자·업체 객체(점검 목록의 담당자 등)에도 평문 비밀번호가 실려 나가던 것을 막는다.
//   - KeepSecretsOnUpdate: 응답에서 비운 값이 수정 화면을 거쳐 빈 값으로 저장되지 않게, 업체 전체 수정(PUT /api/company)에서
//     egiropasswd 가 비어 있으면 기존 값을 채운다 (사용자 비밀번호는 rest/user.go Update 가 같은 처리를 한다).

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"strings"
	"sync"
	"zkeep/global"
	"zkeep/models"
	"zkeep/models/user"

	"github.com/gofiber/fiber/v2"
)

// 미들웨어(인증·업체 분리)는 요청마다 실행되므로 커넥션 풀 하나를 같이 쓴다 (컨트롤러는 요청마다 새로 연다)
var guardDB *sql.DB
var guardDBOnce sync.Once

func sharedDB() *sql.DB {
	guardDBOnce.Do(func() {
		guardDB = models.NewConnection()
	})

	return guardDB
}

// errStop: 미들웨어가 응답을 이미 썼으니 다음 단계로 가지 말라는 표시 (호출한 쪽이 nil 로 바꿔 돌려준다)
var errStop = errors.New("stop")

var secretKeys = map[string]bool{"passwd": true, "egiropasswd": true}

func stripSecrets(v interface{}) {
	switch t := v.(type) {
	case map[string]interface{}:
		for k, child := range t {
			if secretKeys[k] {
				t[k] = ""
				continue
			}
			stripSecrets(child)
		}
	case []interface{}:
		for _, child := range t {
			stripSecrets(child)
		}
	}
}

func SanitizeResponse(c *fiber.Ctx) error {
	err := c.Next()

	if !strings.HasPrefix(c.Path(), "/api/") {
		return err
	}

	if !strings.HasPrefix(string(c.Response().Header.ContentType()), fiber.MIMEApplicationJSON) {
		return err
	}

	body := c.Response().Body()
	if !bytes.Contains(body, []byte(`"passwd"`)) && !bytes.Contains(body, []byte(`"egiropasswd"`)) {
		return err
	}

	var v interface{}
	if json.Unmarshal(body, &v) != nil {
		return err
	}
	stripSecrets(v)

	if out, e := json.Marshal(v); e == nil {
		c.Response().SetBody(out)
	}

	return err
}

func KeepSecretsOnUpdate(c *fiber.Ctx) error {
	// 인증 전 단계라 토큰이 없는 요청은 DB 를 보지 않는다 (토큰 검증은 JwtAuthRequired)
	if c.Method() != fiber.MethodPut || c.Path() != "/api/company" || !strings.HasPrefix(c.Get("Authorization"), "Bearer ") {
		return c.Next()
	}

	var obj map[string]interface{}
	if json.Unmarshal(c.Body(), &obj) != nil || obj == nil {
		return c.Next()
	}

	if v, _ := obj["egiropasswd"].(string); v != "" {
		return c.Next()
	}

	conn := sharedDB()
	if conn == nil {
		return c.Next()
	}

	var old string
	if conn.QueryRow("select c_egiropasswd from company_tb where c_id = ?", toID(obj["id"])).Scan(&old) == nil {
		obj["egiropasswd"] = old
		if body, err := json.Marshal(obj); err == nil {
			c.Request().SetBody(body)
		}
	}

	return c.Next()
}

// 사용자 생성·수정 요청의 passwd 를 bcrypt 해시로 바꾼다 (빈 값은 그대로 — 수정 시 기존 유지).
// bcrypt 는 비싸서(건당 수십 ms) 인증이 끝난 뒤(JwtAuthRequired), 그 쓰기를 할 수 있는 사용자에게만 한다:
// 전체관리자·업체관리자, 또는 자기 정보 수정(PUT /api/user, /api/user/passwd). 일괄 등록은 maxPasswordBatch 건까지
const maxPasswordBatch = 200

func isPasswordWrite(c *fiber.Ctx) bool {
	path := c.Path()
	method := c.Method()
	return (method == fiber.MethodPost && (path == "/api/user" || path == "/api/user/batch")) ||
		(method == fiber.MethodPut && (path == "/api/user" || path == "/api/user/passwd"))
}

// hashUserPasswords 는 처리했으면 nil, 요청을 멈춰야 하면 응답을 쓰고 errStop 을 돌려준다
func hashUserPasswords(c *fiber.Ctx, session *models.User) error {
	if !isPasswordWrite(c) {
		return nil
	}

	// 업체관리자 미만은 자기 정보 수정만 가능하다 (나머지는 tenantGuard 가 거부) — 거부될 요청은 해시하지 않는다
	if session.Level < user.LevelAdmin && c.Method() != fiber.MethodPut {
		return nil
	}

	var body interface{}
	if json.Unmarshal(c.Body(), &body) != nil {
		return nil
	}

	hashOne := func(obj map[string]interface{}) bool {
		for k, v := range obj {
			if strings.ToLower(k) != "passwd" {
				continue
			}
			if p, _ := v.(string); p != "" {
				hash := global.HashPassword(p)
				if hash == "" {
					return false
				}
				obj[k] = hash
			}
		}
		return true
	}

	failed := func() error {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"code": "error", "message": "password error"})
		return errStop
	}

	switch t := body.(type) {
	case map[string]interface{}:
		if session.Level < user.LevelAdmin && toID(t["id"]) != session.Id {
			return nil // 남의 정보 — tenantGuard 가 거부
		}
		if !hashOne(t) {
			return failed()
		}
	case []interface{}:
		if len(t) > maxPasswordBatch {
			c.Status(fiber.StatusRequestEntityTooLarge).JSON(fiber.Map{"code": "error", "message": "한 번에 등록할 수 있는 사용자 수를 넘었습니다"})
			return errStop
		}
		for _, v := range t {
			if obj, ok := v.(map[string]interface{}); ok && !hashOne(obj) {
				return failed()
			}
		}
	default:
		return nil
	}

	if out, err := json.Marshal(body); err == nil {
		c.Request().SetBody(out)
	}

	return nil
}

// 업체 관리자(admin) 이상만 쓰는 커스텀 API — 엑셀 일괄 등록(기존 데이터 삭제·덮어쓰기 포함), 메일·SMS 발송
var adminOnlyPrefixes = []string{
	"/api/external",
	"/api/company/upload/",
	"/api/customer/upload/",
	"/api/user/upload/",
	"/api/mail/",
	"/api/sms/",
}

func requiresTenantAdmin(path string) bool {
	for _, p := range adminOnlyPrefixes {
		if strings.HasSuffix(p, "/") {
			if strings.HasPrefix(path, p) {
				return true
			}
		} else if path == p || strings.HasPrefix(path, p+"/") {
			return true
		}
	}

	return false
}
