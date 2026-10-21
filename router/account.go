package router

// 회원 탈퇴 — 앱에서 가입할 수 있으면 앱 안에서 탈퇴도 되어야 한다 (App Store 5.1.1(v), Google Play 계정 삭제 정책).
// 점검·청구 기록이 사용자 id 를 참조하므로 행은 지우지 않고 개인정보를 지운 뒤 사용 중지한다.
// 인증이 필요한 라우트라 JwtAuthRequired 뒤(router.SetRouter 다음)에 등록한다 — services/http.go

import (
	"fmt"
	"time"
	"zkeep/global"
	"zkeep/models"
	"zkeep/models/user"

	"github.com/gofiber/fiber/v2"
)

func SetAccountRouter(r *fiber.App) {
	r.Post("/api/account/withdraw", signupBody, accountWithdraw)
}

// POST /api/account/withdraw {passwd} — 본인 확인 후 탈퇴
func accountWithdraw(c *fiber.Ctx) error {
	session, _ := c.Locals("user").(*models.User)
	if session == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"code": "error", "message": "unauthorized"})
	}

	var req struct {
		Passwd string `json:"passwd"`
	}
	if err := c.BodyParser(&req); err != nil || req.Passwd == "" || len(req.Passwd) > global.MaxPasswordInput {
		return signupError(c, "비밀번호를 입력하세요")
	}

	if session.Level == user.LevelRootadmin {
		return signupError(c, "전체관리자 계정은 탈퇴할 수 없습니다")
	}

	conn := sharedDB()
	if conn == nil {
		return serverError(c)
	}

	var stored string
	if err := conn.QueryRow("select u_passwd from user_tb where u_id = ?", session.Id).Scan(&stored); err != nil {
		return serverError(c)
	}
	if ok, _ := global.CheckPassword(stored, req.Passwd); !ok {
		return signupError(c, "비밀번호가 맞지 않습니다")
	}

	// 「마지막 관리자」 확인과 탈퇴 처리를 한 트랜잭션에서 — 같은 업체 관리자 행을 잠가
	// 관리자 둘이 동시에 탈퇴해 직원만 남는(관리자 0명) 경우를 막는다
	tx, err := conn.Begin()
	if err != nil {
		return serverError(c)
	}
	defer tx.Rollback()

	if session.Level == user.LevelAdmin {
		rows, err := tx.Query("select u_id from user_tb where u_company = ? and u_level = ? and u_status = ? for update", session.Company, int(user.LevelAdmin), int(user.StatusUse))
		if err != nil {
			return serverError(c)
		}
		admins := 0
		for rows.Next() {
			var id int64
			if rows.Scan(&id) == nil && id != session.Id {
				admins++
			}
		}
		rows.Close()

		var others int
		if err := tx.QueryRow("select count(*) from user_tb where u_company = ? and u_status = ? and u_id <> ?", session.Company, int(user.StatusUse), session.Id).Scan(&others); err != nil {
			return serverError(c)
		}
		if admins == 0 && others > 0 {
			return signupError(c, "다른 직원에게 관리자 권한을 넘긴 뒤 탈퇴하세요")
		}
	}

	loginid := fmt.Sprintf("withdrawn_%v_%v", session.Id, time.Now().Unix())
	if _, err := tx.Exec(`update user_tb set u_loginid = ?, u_passwd = '', u_name = '탈퇴회원', u_email = '', u_tel = '',
		u_zip = '', u_address = '', u_addressetc = '', u_educationdate = '1000-01-01', u_educationinstitution = '',
		u_specialeducationdate = '1000-01-01', u_specialeducationinstitution = '', u_rejectreason = '', u_status = ?
		where u_id = ?`, loginid, int(user.StatusNotuse), session.Id); err != nil {
		return serverError(c)
	}
	if _, err := tx.Exec("delete from license_tb where l_user = ?", session.Id); err != nil {
		return serverError(c)
	}
	if err := tx.Commit(); err != nil {
		return serverError(c)
	}

	return c.JSON(fiber.Map{"code": "ok"})
}
