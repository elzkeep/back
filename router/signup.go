package router

// 회원가입 전용 공개 API — 비로그인 사용자가 쓸 수 있는 범위를 여기로 모은다.
// 예전에는 GET /api/user, POST /api/user, GET/POST /api/company, GET /api/department 를 비로그인에 열어 두고
// 클라이언트가 레벨·승인 상태를 정해 보냈다(전체관리자 계정 생성, 사용자 목록·비밀번호 확인 가능).
// 이제 레벨·승인·업체 종류는 서버가 정하고, 목록은 가입 화면에 필요한 id·이름만 내려준다.
// router.go(생성 파일) 밖에서 JwtAuthRequired 보다 먼저 등록해야 공개 라우트가 된다 — services/http.go 참고.

import (
	"database/sql"
	"strings"
	"time"
	"zkeep/models"
	"zkeep/models/company"
	"zkeep/models/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

const (
	signupBodyLimit   = 64 * 1024
	signupMaxLicenses = 10
)

// 공개 API 요청 횟수 제한 키 — 앞단 프록시(nginx) 뒤에서는 c.IP() 가 모두 같으므로 X-Real-IP / X-Forwarded-For 를 우선
func clientKey(c *fiber.Ctx) string {
	if v := strings.TrimSpace(c.Get("X-Real-IP")); v != "" {
		return v
	}
	if v := c.Get("X-Forwarded-For"); v != "" {
		return strings.TrimSpace(strings.Split(v, ",")[0])
	}
	return c.IP()
}

func newPublicLimiter(max int, window time.Duration) fiber.Handler {
	return limiter.New(limiter.Config{
		Max:          max,
		Expiration:   window,
		KeyGenerator: clientKey,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{"code": "error", "message": "잠시 후 다시 시도하세요"})
		},
	})
}

type signupUser struct {
	Loginid    string `json:"loginid"`
	Passwd     string `json:"passwd"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Tel        string `json:"tel"`
	Zip        string `json:"zip"`
	Address    string `json:"address"`
	Addressetc string `json:"addressetc"`
	Company    int64  `json:"company"`
	Department int64  `json:"department"`
}

type signupCompany struct {
	Name       string `json:"name"`
	Companyno  string `json:"companyno"`
	Ceo        string `json:"ceo"`
	Tel        string `json:"tel"`
	Email      string `json:"email"`
	Address    string `json:"address"`
	Addressetc string `json:"addressetc"`
}

type signupLicense struct {
	Licensecategory int64  `json:"licensecategory"`
	Licenselevel    int64  `json:"licenselevel"`
	Takingdate      string `json:"takingdate"`
}

type signupRequest struct {
	Company  signupCompany   `json:"company"`
	User     signupUser      `json:"user"`
	Licenses []signupLicense `json:"licenses"`
}

func signupError(c *fiber.Ctx, message string) error {
	return c.JSON(fiber.Map{"code": "error", "message": message})
}

func signupBody(c *fiber.Ctx) error {
	if len(c.Body()) > signupBodyLimit {
		return c.Status(fiber.StatusRequestEntityTooLarge).JSON(fiber.Map{"code": "error", "message": "요청이 너무 큽니다"})
	}
	return c.Next()
}

func SetSignupRouter(r *fiber.App) {
	check := newPublicLimiter(60, 10*time.Minute)
	create := newPublicLimiter(10, time.Hour)

	r.Post("/api/signup/check", check, signupBody, signupCheck)
	r.Get("/api/signup/companies", check, signupCompanies)
	r.Get("/api/signup/departments", check, signupDepartments)
	r.Post("/api/signup/company", create, signupBody, signupCreateCompany)
	r.Post("/api/signup/user", create, signupBody, signupCreateUser)
}

// POST /api/signup/check {loginid, email, companyno, except} → 각 값의 사용 여부(true = 이미 있음)만 돌려준다.
// except 는 사용자 수정 화면에서 자기 자신을 빼고 검사할 사용자 id
func signupCheck(c *fiber.Ctx) error {
	var req struct {
		Loginid   string `json:"loginid"`
		Email     string `json:"email"`
		Companyno string `json:"companyno"`
		Except    int64  `json:"except"`
	}
	if err := c.BodyParser(&req); err != nil {
		return signupError(c, "잘못된 요청입니다")
	}

	conn := models.NewConnection()
	if conn == nil {
		return signupError(c, "database error")
	}
	defer conn.Close()

	failed := false
	exists := func(query string, args ...interface{}) bool {
		if s, _ := args[0].(string); strings.TrimSpace(s) == "" {
			return false
		}
		var n int
		if err := conn.QueryRow(query, args...).Scan(&n); err != nil {
			failed = true
			return false
		}
		return n > 0
	}

	result := fiber.Map{
		"code":      "ok",
		"loginid":   exists("select count(*) from user_tb where u_loginid = ? and u_id <> ?", req.Loginid, req.Except),
		"email":     exists("select count(*) from user_tb where u_email = ? and u_id <> ?", req.Email, req.Except),
		"companyno": exists("select count(*) from company_tb where c_companyno = ? and c_type = ?", req.Companyno, int(company.TypeWork)),
	}
	if failed {
		return signupError(c, "database error")
	}

	return c.JSON(result)
}

// GET /api/signup/companies → 가입할 수 있는 점검업체(id, name) 목록
func signupCompanies(c *fiber.Ctx) error {
	conn := models.NewConnection()
	if conn == nil {
		return signupError(c, "database error")
	}
	defer conn.Close()

	rows, err := conn.Query("select c_id, c_name from company_tb where c_type = ? order by c_name", int(company.TypeWork))
	if err != nil {
		return signupError(c, "database error")
	}
	defer rows.Close()

	items := []fiber.Map{}
	for rows.Next() {
		var id int64
		var name string
		if rows.Scan(&id, &name) == nil {
			items = append(items, fiber.Map{"id": id, "name": name})
		}
	}

	return c.JSON(fiber.Map{"code": "ok", "items": items})
}

// GET /api/signup/departments?company= → 그 업체의 팀(id, name) 목록
func signupDepartments(c *fiber.Ctx) error {
	companyId := int64(c.QueryInt("company"))
	if companyId <= 0 {
		return c.JSON(fiber.Map{"code": "ok", "items": []fiber.Map{}})
	}

	conn := models.NewConnection()
	if conn == nil {
		return signupError(c, "database error")
	}
	defer conn.Close()

	rows, err := conn.Query("select de_id, de_name from department_tb where de_company = ? order by de_order, de_id", companyId)
	if err != nil {
		return signupError(c, "database error")
	}
	defer rows.Close()

	items := []fiber.Map{}
	for rows.Next() {
		var id int64
		var name string
		if rows.Scan(&id, &name) == nil {
			items = append(items, fiber.Map{"id": id, "name": name})
		}
	}

	return c.JSON(fiber.Map{"code": "ok", "items": items})
}

func validateSignupUser(conn *sql.DB, u *signupUser) string {
	u.Loginid = strings.TrimSpace(u.Loginid)
	u.Email = strings.TrimSpace(u.Email)
	if u.Loginid == "" {
		return "로그인 아이디를 입력하세요"
	}
	if u.Passwd == "" {
		return "비밀번호를 입력하세요"
	}
	if strings.TrimSpace(u.Name) == "" {
		return "이름을 입력하세요"
	}

	var n int
	if err := conn.QueryRow("select count(*) from user_tb where u_loginid = ?", u.Loginid).Scan(&n); err != nil {
		return "database error"
	}
	if n > 0 {
		return "이미 등록된 로그인아이디 입니다"
	}
	if u.Email != "" {
		if err := conn.QueryRow("select count(*) from user_tb where u_email = ?", u.Email).Scan(&n); err != nil {
			return "database error"
		}
		if n > 0 {
			return "이미 등록된 이메일 입니다"
		}
	}

	return ""
}

func insertSignupUser(tx *sql.Tx, u signupUser, level user.Level, approval user.Approval, licenses []signupLicense) (int64, error) {
	if len(licenses) > signupMaxLicenses {
		licenses = licenses[:signupMaxLicenses]
	}

	userManager := models.NewUserManager(tx)
	item := models.User{
		Loginid:    u.Loginid,
		Passwd:     u.Passwd,
		Name:       u.Name,
		Email:      u.Email,
		Tel:        u.Tel,
		Zip:        u.Zip,
		Address:    u.Address,
		Addressetc: u.Addressetc,
		Level:      level,
		Approval:   approval,
		Status:     user.StatusUse,
		Company:    u.Company,
		Department: u.Department,
	}
	if err := userManager.Insert(&item); err != nil {
		return 0, err
	}
	id := userManager.GetIdentity()

	licenseManager := models.NewLicenseManager(tx)
	for _, l := range licenses {
		if l.Licensecategory <= 0 {
			continue
		}
		license := models.License{User: id, Licensecategory: l.Licensecategory, Licenselevel: l.Licenselevel, Takingdate: l.Takingdate}
		if err := licenseManager.Insert(&license); err != nil {
			return 0, err
		}
	}

	return id, nil
}

// POST /api/signup/company {company, user, licenses} → 점검업체 + 그 업체 관리자 계정
func signupCreateCompany(c *fiber.Ctx) error {
	var req signupRequest
	if err := c.BodyParser(&req); err != nil {
		return signupError(c, "잘못된 요청입니다")
	}

	req.Company.Companyno = strings.TrimSpace(req.Company.Companyno)
	if strings.TrimSpace(req.Company.Name) == "" {
		return signupError(c, "사업장 이름을 입력하세요")
	}
	if req.Company.Companyno == "" {
		return signupError(c, "사업자번호를 입력하세요")
	}

	conn := models.NewConnection()
	if conn == nil {
		return signupError(c, "database error")
	}
	defer conn.Close()

	if msg := validateSignupUser(conn, &req.User); msg != "" {
		return signupError(c, msg)
	}
	var n int
	if err := conn.QueryRow("select count(*) from company_tb where c_companyno = ? and c_type = ?", req.Company.Companyno, int(company.TypeWork)).Scan(&n); err != nil {
		return signupError(c, "database error")
	}
	if n > 0 {
		return signupError(c, "이미 등록된 사업장입니다")
	}

	tx, err := conn.Begin()
	if err != nil {
		return signupError(c, "database error")
	}
	defer tx.Rollback()

	companyManager := models.NewCompanyManager(tx)
	item := models.Company{
		Name:       req.Company.Name,
		Companyno:  req.Company.Companyno,
		Ceo:        req.Company.Ceo,
		Tel:        req.Company.Tel,
		Email:      req.Company.Email,
		Address:    req.Company.Address,
		Addressetc: req.Company.Addressetc,
		Type:       company.TypeWork,
	}
	if err := companyManager.Insert(&item); err != nil {
		return signupError(c, "database error")
	}
	companyId := companyManager.GetIdentity()

	req.User.Company = companyId
	req.User.Department = 0
	userId, err := insertSignupUser(tx, req.User, user.LevelAdmin, user.ApprovalComplete, req.Licenses)
	if err != nil {
		return signupError(c, "database error")
	}

	if err := tx.Commit(); err != nil {
		return signupError(c, "database error")
	}

	return c.JSON(fiber.Map{"code": "ok", "company": companyId, "user": userId})
}

// POST /api/signup/user {user, licenses} → 기존 점검업체에 일반 직원으로 가입 (업체 관리자 승인 대기)
func signupCreateUser(c *fiber.Ctx) error {
	var req signupRequest
	if err := c.BodyParser(&req); err != nil {
		return signupError(c, "잘못된 요청입니다")
	}

	conn := models.NewConnection()
	if conn == nil {
		return signupError(c, "database error")
	}
	defer conn.Close()

	if msg := validateSignupUser(conn, &req.User); msg != "" {
		return signupError(c, msg)
	}

	var n int
	if err := conn.QueryRow("select count(*) from company_tb where c_id = ? and c_type = ?", req.User.Company, int(company.TypeWork)).Scan(&n); err != nil || n == 0 {
		return signupError(c, "소속 업체를 선택하세요")
	}
	if req.User.Department != 0 {
		if err := conn.QueryRow("select count(*) from department_tb where de_id = ? and de_company = ?", req.User.Department, req.User.Company).Scan(&n); err != nil || n == 0 {
			return signupError(c, "팀을 다시 선택하세요")
		}
	}

	tx, err := conn.Begin()
	if err != nil {
		return signupError(c, "database error")
	}
	defer tx.Rollback()

	userId, err := insertSignupUser(tx, req.User, user.LevelNormal, user.ApprovalWait, req.Licenses)
	if err != nil {
		return signupError(c, "database error")
	}

	if err := tx.Commit(); err != nil {
		return signupError(c, "database error")
	}

	return c.JSON(fiber.Map{"code": "ok", "user": userId})
}
