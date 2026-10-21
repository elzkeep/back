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
	"strings"
	"sync"
	"zkeep/global"
	"zkeep/models"

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

// HashPasswordOnWrite 는 사용자 생성·수정 요청의 passwd 를 bcrypt 해시로 바꾼다 (빈 값은 그대로 — 수정 시 기존 유지)
func HashPasswordOnWrite(c *fiber.Ctx) error {
	path := c.Path()
	method := c.Method()
	if !((method == fiber.MethodPost && (path == "/api/user" || path == "/api/user/batch")) ||
		(method == fiber.MethodPut && (path == "/api/user" || path == "/api/user/passwd"))) {
		return c.Next()
	}

	var body interface{}
	if json.Unmarshal(c.Body(), &body) != nil {
		return c.Next()
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

	switch t := body.(type) {
	case map[string]interface{}:
		if !hashOne(t) {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"code": "error", "message": "password error"})
		}
	case []interface{}:
		for _, v := range t {
			if obj, ok := v.(map[string]interface{}); ok && !hashOne(obj) {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"code": "error", "message": "password error"})
			}
		}
	default:
		return c.Next()
	}

	if out, err := json.Marshal(body); err == nil {
		c.Request().SetBody(out)
	}

	return c.Next()
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
