package router

// 업체(테넌트) 분리 — 생성 CRUD(controllers/rest, router.go)는 업체 검사를 하지 않는다.
// JwtAuthRequired 가 토큰을 확인한 뒤 tenantGuard 를 거치게 해서, 전체관리자가 아닌 사용자의 요청을 자기 업체 범위로 묶는다.
//   - 목록·개수(GET /api/x, /x/count, /x/sum): 업체 컬럼이 있는 테이블은 company 쿼리를 세션 업체로 고정
//   - id 로 접근(GET /x/:id, PUT, DELETE, batch, 필드별 PUT): 그 행이 내 업체 소유인지 확인
//   - 쓰기(POST/PUT): 업체 컬럼을 세션 업체로 덮어쓰고, 참조(건물·점검·사용자 등)가 내 업체 것인지 확인
//   - 응답 후처리: 업체 컬럼이 없는 테이블(업체·건물·설비·점검항목·자격증)은 내 것이 아닌 행을 걸러낸다
//   - 기준 정보(자격 종류·공지 등)는 읽기만, e지로 수납·가입문의·문의는 전체관리자만
// 생성 라우트 모양이 아닌 경로(/x/search/:page, /x/make 등)는 controllers/api 가 직접 검사한다.
// 소유 판정:
//   - 업체: 내 업체, customercompany 로 연결한 건물주 업체, 내 계약(customer) 건물의 건물주 업체
//   - 건물: 내 계약이 있는 건물, 또는 내가 연결한 건물주 업체의 건물 중 다른 업체 계약이 없는 것(계약 전 등록 단계)

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"strings"
	"zkeep/models"
	"zkeep/models/company"
	"zkeep/models/user"

	"github.com/gofiber/fiber/v2"
)

type ownerKind int

const (
	kindTenant       ownerKind = iota + 1 // 테이블에 업체 컬럼이 있다
	kindCompany                           // company_tb
	kindBuilding                          // building_tb
	kindFacility                          // facility_tb (건물 소속)
	kindItem                              // item_tb (점검 소속)
	kindLicense                           // license_tb (사용자 소속)
	kindOwnerCompany                      // customercompany.customer 로 연결할 건물주 업체
	kindGlobal                            // 기준 정보 — 읽기는 로그인 사용자 모두, 쓰기는 전체관리자
	kindRoot                              // 전체관리자 전용
)

type entityRule struct {
	kind   ownerKind
	table  string
	prefix string
	view   bool              // 뷰 — 쓰기 없음
	refs   map[string]string // 쓰기 때 검증할 참조 필드 → 엔티티 이름
}

func (r entityRule) idColumn() string      { return r.prefix + "_id" }
func (r entityRule) companyColumn() string { return r.prefix + "_company" }

var tenantRules = map[string]entityRule{
	"billing":             {kind: kindTenant, table: "billing_tb", prefix: "bi", refs: map[string]string{"building": "building"}},
	"billinghistory":      {kind: kindTenant, table: "billinghistory_tb", prefix: "bh", refs: map[string]string{"building": "building", "billing": "billing"}},
	"billinglist":         {kind: kindTenant, table: "billinglist_vw", prefix: "bi", view: true},
	"billinguserlist":     {kind: kindTenant, table: "billinguserlist_vw", prefix: "bi", view: true},
	"calendarcompanylist": {kind: kindTenant, table: "calendarcompanylist_vw", prefix: "r", view: true},
	"companylicense":      {kind: kindTenant, table: "companylicense_tb", prefix: "l"},
	"customer":            {kind: kindTenant, table: "customer_tb", prefix: "cu", refs: map[string]string{"building": "building", "user": "user", "salesuser": "user"}},
	"customercompany":     {kind: kindTenant, table: "customercompany_tb", prefix: "cc", refs: map[string]string{"customer": "ownercompany"}},
	"customercompanylist": {kind: kindTenant, table: "customercompanylist_vw", prefix: "c", view: true},
	"companylist":         {kind: kindTenant, table: "companylist_vw", prefix: "c", view: true},
	"data":                {kind: kindTenant, table: "data_tb", prefix: "d", refs: map[string]string{"report": "report"}},
	"department":          {kind: kindTenant, table: "department_tb", prefix: "de", refs: map[string]string{"parent": "department", "master": "user"}},
	"report":              {kind: kindTenant, table: "report_tb", prefix: "r", refs: map[string]string{"building": "building", "user": "user"}},
	"reportlist":          {kind: kindTenant, table: "reportlist_vw", prefix: "r", view: true},
	"statisticsday":       {kind: kindTenant, table: "statisticsday_vw", prefix: "bi", view: true},
	"statisticsmonth":     {kind: kindTenant, table: "statisticsmonth_vw", prefix: "bi", view: true},
	"statisticsyear":      {kind: kindTenant, table: "statisticsyear_vw", prefix: "bi", view: true},
	"user":                {kind: kindTenant, table: "user_tb", prefix: "u", refs: map[string]string{"department": "department"}},
	"userlist":            {kind: kindTenant, table: "userlist_vw", prefix: "u", view: true},

	"company":      {kind: kindCompany, table: "company_tb", prefix: "c"},
	"ownercompany": {kind: kindOwnerCompany, table: "company_tb", prefix: "c"},
	"building":     {kind: kindBuilding, table: "building_tb", prefix: "b", refs: map[string]string{"company": "company"}},
	"facility":     {kind: kindFacility, table: "facility_tb", prefix: "f", refs: map[string]string{"building": "building", "parent": "facility"}},
	"item":         {kind: kindItem, table: "item_tb", prefix: "i", refs: map[string]string{"report": "report", "data": "data"}},
	"license":      {kind: kindLicense, table: "license_tb", prefix: "l", refs: map[string]string{"user": "user"}},

	"licensecategory":        {kind: kindGlobal},
	"licenselevel":           {kind: kindGlobal},
	"companylicensecategory": {kind: kindGlobal},
	"notice":                 {kind: kindGlobal},
	"webfaq":                 {kind: kindGlobal},
	"webnotice":              {kind: kindGlobal},

	"giro":        {kind: kindRoot},
	"webjoin":     {kind: kindRoot},
	"webquestion": {kind: kindRoot},
}

// 생성 GET 메서드 중 경로 인자가 업체·사용자인 것 — "엔티티/동작/컬럼" → 인자 검사 종류
var tenantParamRoutes = map[string]string{
	"building/get/companyname":            "company",
	"customer/count/companybuilding":      "self",
	"customer/get/companybuilding":        "self",
	"customercompany/get/companycustomer": "self",
	"department/get/companyname":          "self",
	"user/count/company":                  "self",
	"user/get/companyname":                "self",
	"user/get/companytelname":             "self",
	"license/get/userlicensecategory":     "user",
	"license/find/user":                   "user",
}

// 사용자 레벨·승인·상태·업체는 업체 관리자만 바꾼다
var userPrivilegedFields = map[string]bool{"level": true, "approval": true, "status": true, "company": true, "score": true}

type tenantScope struct {
	conn    *sql.DB
	session *models.User
}

func (s *tenantScope) company() int64 { return s.session.Company }

func (s *tenantScope) isTenantAdmin() bool { return s.session.Level >= user.LevelAdmin }

// ownerCondition 은 "이 엔티티의 행 중 내 업체 소유" 조건과 인자를 돌려준다
func (s *tenantScope) ownerCondition(rule entityRule) (string, []interface{}) {
	S := s.company()
	buildingCond := "(b_id in (select cu_building from customer_tb where cu_company = ?) or ((b_company = ? or b_company in (select cc_customer from customercompany_tb where cc_company = ?)) and not exists (select 1 from customer_tb where cu_building = b_id and cu_company <> ?)))"
	buildingArgs := []interface{}{S, S, S, S}

	switch rule.kind {
	case kindTenant:
		return rule.companyColumn() + " = ?", []interface{}{S}
	case kindCompany:
		// 연결·계약으로 보이는 업체는 건물주 업체뿐 — 다른 점검업체(테넌트)의 업체 정보(계좌·e지로)는 보이지 않는다
		// (예전 엑셀 일괄등록이 c_type 을 비워 둔 건물주가 있어 「점검업체가 아님」으로 본다)
		return "(c_id = ? or (c_type <> ? and (c_id in (select cc_customer from customercompany_tb where cc_company = ?) or c_id in (select b_company from building_tb join customer_tb on cu_building = b_id where cu_company = ?))))", []interface{}{S, int(company.TypeWork), S, S}
	case kindOwnerCompany:
		// 새로 연결할 수 있는 건물주: 점검업체가 아니고, 다른 업체와 연결·계약이 없는 것
		return "(c_type <> ? and not exists (select 1 from customercompany_tb where cc_customer = c_id and cc_company <> ?) and not exists (select 1 from building_tb join customer_tb on cu_building = b_id where b_company = c_id and cu_company <> ?))", []interface{}{int(company.TypeWork), S, S}
	case kindBuilding:
		return buildingCond, buildingArgs
	case kindFacility:
		return "f_building in (select b_id from building_tb where " + buildingCond + ")", buildingArgs
	case kindItem:
		return "i_report in (select r_id from report_tb where r_company = ?)", []interface{}{S}
	case kindLicense:
		return "l_user in (select u_id from user_tb where u_company = ?)", []interface{}{S}
	}

	return "1 = 0", nil
}

// ownedIDs 는 ids 중 내 업체 소유인 id 집합을 돌려준다
func (s *tenantScope) ownedIDs(rule entityRule, ids []int64) (map[int64]bool, error) {
	owned := map[int64]bool{}
	if len(ids) == 0 {
		return owned, nil
	}

	cond, args := s.ownerCondition(rule)

	// placeholder 상한을 넘지 않게 나눠서 묻는다
	const chunk = 500
	for start := 0; start < len(ids); start += chunk {
		end := start + chunk
		if end > len(ids) {
			end = len(ids)
		}
		part := ids[start:end]

		holders := make([]string, len(part))
		params := make([]interface{}, 0, len(part)+len(args))
		for i, id := range part {
			holders[i] = "?"
			params = append(params, id)
		}
		params = append(params, args...)

		query := fmt.Sprintf("select distinct %v from %v where %v in (%v) and %v", rule.idColumn(), rule.table, rule.idColumn(), strings.Join(holders, ","), cond)
		rows, err := s.conn.Query(query, params...)
		if err != nil {
			log.Println("tenant owner query :", err)
			return nil, err
		}
		for rows.Next() {
			var id int64
			if rows.Scan(&id) == nil {
				owned[id] = true
			}
		}
		rows.Close()
	}

	return owned, nil
}

func (s *tenantScope) ownsAll(entity string, ids []int64) bool {
	rule, ok := tenantRules[entity]
	if !ok || rule.table == "" {
		return false
	}

	unique := []int64{}
	seen := map[int64]bool{}
	for _, id := range ids {
		if id > 0 && !seen[id] {
			seen[id] = true
			unique = append(unique, id)
		}
	}
	if len(unique) == 0 {
		return true
	}

	owned, err := s.ownedIDs(rule, unique)
	if err != nil {
		return false
	}

	return len(owned) == len(unique)
}

// 요청 JSON 값을 id 로 — 숫자, 숫자 문자열, {id: n} 객체(앱 모델의 참조 직렬화)를 받는다
func toID(v interface{}) int64 {
	switch t := v.(type) {
	case float64:
		return int64(t)
	case string:
		n, _ := strconv.ParseInt(t, 10, 64)
		return n
	case map[string]interface{}:
		return toID(t["id"])
	}

	return 0
}

func forbidden(c *fiber.Ctx) error {
	return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"code": "error", "message": "forbidden"})
}

func serverError(c *fiber.Ctx) error {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"code": "error", "message": "database error"})
}

// normalizeKeys 는 요청 JSON 키를 소문자로 맞춘다. 핸들러는 구조체에 대소문자 무시로 바인딩하므로
// "Building"·"Level" 같은 키가 검사를 피해 가지 않게 한다. 대소문자만 다른 키가 둘이거나 ASCII 가 아닌 키는 거부
func normalizeKeys(obj map[string]interface{}) (map[string]interface{}, bool) {
	out := make(map[string]interface{}, len(obj))
	for k, v := range obj {
		for _, r := range k {
			if r > 127 {
				return nil, false
			}
		}
		lower := strings.ToLower(k)
		if _, dup := out[lower]; dup {
			return nil, false
		}
		out[lower] = v
	}
	return out, true
}

type routeShape int

const (
	shapeCustom routeShape = iota // 생성 CRUD 모양이 아님 → controllers/api 가 검사
	shapeList
	shapeRead
	shapeParamGet
	shapeInsert
	shapeInsertBatch
	shapeUpdate
	shapeFieldUpdate
	shapeDelete
	shapeDeleteBatch
	shapeDeleteBy
)

func classifyRoute(method string, rest []string) routeShape {
	switch method {
	case fiber.MethodGet, fiber.MethodHead:
		if len(rest) == 0 || (len(rest) == 1 && (rest[0] == "count" || rest[0] == "sum")) {
			return shapeList
		}
		if len(rest) == 1 {
			if _, err := strconv.ParseInt(rest[0], 10, 64); err == nil {
				return shapeRead
			}
		}
		if len(rest) == 3 && (rest[0] == "get" || rest[0] == "count" || rest[0] == "find") {
			return shapeParamGet
		}
	case fiber.MethodPost:
		if len(rest) == 0 {
			return shapeInsert
		}
		if len(rest) == 1 && rest[0] == "batch" {
			return shapeInsertBatch
		}
	case fiber.MethodPut:
		if len(rest) == 0 {
			return shapeUpdate
		}
		if len(rest) == 1 {
			return shapeFieldUpdate
		}
	case fiber.MethodDelete:
		if len(rest) == 0 {
			return shapeDelete
		}
		if len(rest) == 1 && rest[0] == "batch" {
			return shapeDeleteBatch
		}
		if len(rest) == 1 && strings.HasPrefix(rest[0], "by") {
			return shapeDeleteBy
		}
	}

	return shapeCustom
}

// tenantGuard 는 인증된 요청에서 호출된다(c.Locals("user") 설정 후)
func tenantGuard(c *fiber.Ctx, session *models.User) error {
	if session.Level == user.LevelRootadmin {
		return c.Next()
	}

	u, err := url.Parse(c.Path())
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": "error", "message": "bad request"})
	}

	if requiresTenantAdmin(u.Path) && session.Level < user.LevelAdmin {
		return forbidden(c)
	}
	segments := strings.Split(strings.Trim(strings.TrimPrefix(u.Path, "/api"), "/"), "/")
	if len(segments) == 0 || segments[0] == "" {
		return c.Next()
	}

	entity := strings.ToLower(segments[0])
	rest := segments[1:]
	rule, ok := tenantRules[entity]
	if !ok {
		return c.Next()
	}

	method := c.Method()
	switch method {
	case fiber.MethodGet, fiber.MethodHead, fiber.MethodPost, fiber.MethodPut, fiber.MethodDelete:
	default:
		return forbidden(c)
	}
	if method == fiber.MethodHead {
		method = fiber.MethodGet
	}
	shape := classifyRoute(method, rest)

	switch rule.kind {
	case kindRoot:
		// 가입문의·문의 등록은 누구나
		if method == fiber.MethodPost && shape == shapeInsert && (entity == "webjoin" || entity == "webquestion") {
			return c.Next()
		}
		return forbidden(c)
	case kindGlobal:
		if method == fiber.MethodGet {
			return c.Next()
		}
		return forbidden(c)
	}

	if shape == shapeCustom {
		return c.Next()
	}

	if rule.view && method != fiber.MethodGet {
		return forbidden(c)
	}

	conn := sharedDB()
	if conn == nil {
		return serverError(c)
	}

	scope := &tenantScope{conn: conn, session: session}

	switch shape {
	case shapeList:
		aggregate := len(rest) == 1 // count, sum
		if rule.kind == kindTenant {
			c.Request().URI().QueryArgs().Set("company", strconv.FormatInt(scope.company(), 10))
			if aggregate {
				return c.Next() // 업체가 고정된 집계 — 응답에 행이 없다
			}
		} else if aggregate {
			if rest[0] == "sum" {
				return forbidden(c) // 업체 컬럼이 없는 테이블의 합계는 전 업체 합계
			}
			return c.Next()
		}
		return scope.filterResponse(c, entity, rule)
	case shapeRead:
		id, _ := strconv.ParseInt(rest[0], 10, 64)
		if id <= 0 || !scope.ownsAll(entity, []int64{id}) {
			return forbidden(c)
		}
		return scope.filterResponse(c, entity, rule)
	case shapeParamGet:
		key := entity + "/" + rest[0] + "/" + rest[1]
		value, _ := strconv.ParseInt(rest[2], 10, 64)
		switch tenantParamRoutes[key] {
		case "self":
			if value != scope.company() {
				return forbidden(c)
			}
		case "company":
			if !scope.ownsAll("company", []int64{value}) {
				return forbidden(c)
			}
		case "user":
			if !scope.ownsAll("user", []int64{value}) {
				return forbidden(c)
			}
		}
		return scope.filterResponse(c, entity, rule)
	}

	return scope.guardWrite(c, entity, rule, shape, rest)
}

func (s *tenantScope) guardWrite(c *fiber.Ctx, entity string, rule entityRule, shape routeShape, rest []string) error {
	// 다시 쓴 JSON 본문이 그대로 핸들러에 가도록 JSON·비압축만 받는다
	if !strings.HasPrefix(strings.ToLower(string(c.Request().Header.ContentType())), fiber.MIMEApplicationJSON) || len(c.Request().Header.Peek(fiber.HeaderContentEncoding)) > 0 {
		return forbidden(c)
	}

	var raw []map[string]interface{}
	batch := shape == shapeInsertBatch || shape == shapeDeleteBatch

	if batch {
		if err := json.Unmarshal(c.Body(), &raw); err != nil {
			return forbidden(c)
		}
	} else {
		var obj map[string]interface{}
		if err := json.Unmarshal(c.Body(), &obj); err != nil || obj == nil {
			return forbidden(c)
		}
		raw = []map[string]interface{}{obj}
	}

	objects := make([]map[string]interface{}, 0, len(raw))
	for _, obj := range raw {
		if obj == nil {
			return forbidden(c)
		}
		normalized, ok := normalizeKeys(obj)
		if !ok {
			return forbidden(c)
		}
		objects = append(objects, normalized)
	}
	if len(objects) == 0 {
		return forbidden(c)
	}

	field := ""
	if shape == shapeFieldUpdate {
		field = strings.ToLower(rest[0])
		if field == "company" && rule.kind == kindTenant {
			return forbidden(c)
		}
	}

	if entity == "user" && !s.allowUserWrite(shape, field, objects) {
		return forbidden(c)
	}

	// 기존 행을 가리키는 쓰기는 소유 확인
	if shape == shapeUpdate || shape == shapeFieldUpdate || shape == shapeDelete || shape == shapeDeleteBatch {
		ids := []int64{}
		for _, obj := range objects {
			id := toID(obj["id"])
			if id <= 0 {
				return forbidden(c)
			}
			ids = append(ids, id)
		}
		if !s.ownsAll(entity, ids) {
			return forbidden(c)
		}
	}

	refIDs := map[string][]int64{}
	for _, obj := range objects {
		if rule.kind == kindTenant {
			if _, has := obj["company"]; has || shape == shapeInsert || shape == shapeInsertBatch || shape == shapeUpdate || shape == shapeDeleteBy {
				obj["company"] = s.company()
			}
		}

		for name, value := range obj {
			refEntity, isRef := rule.refs[name]
			if !isRef {
				continue
			}
			if shape == shapeFieldUpdate && name != field {
				continue
			}
			if id := toID(value); id > 0 {
				refIDs[refEntity] = append(refIDs[refEntity], id)
			}
		}
	}

	for refEntity, ids := range refIDs {
		if !s.ownsAll(refEntity, ids) {
			return forbidden(c)
		}
	}

	var body []byte
	var err error
	if batch {
		body, err = json.Marshal(objects)
	} else {
		body, err = json.Marshal(objects[0])
	}
	if err != nil {
		return forbidden(c)
	}
	c.Request().SetBody(body)

	return c.Next()
}

// 사용자 쓰기: 전체관리자 레벨은 못 만들고, 일반·팀장은 자기 정보(권한 필드 제외)만 고친다.
// 전체 수정(PUT /api/user)에서는 권한 필드를 DB 의 현재 값으로 되돌린다
func (s *tenantScope) allowUserWrite(shape routeShape, field string, objects []map[string]interface{}) bool {
	for _, obj := range objects {
		if v, has := obj["level"]; has && toID(v) >= int64(user.LevelRootadmin) {
			return false
		}
	}
	if shape == shapeFieldUpdate && field == "level" {
		if v, has := objects[0]["level"]; has && toID(v) >= int64(user.LevelRootadmin) {
			return false
		}
	}

	if s.isTenantAdmin() {
		// 업체 관리자는 자기보다 높은 레벨(전체관리자 포함)의 사용자를 만들거나 고치거나 지우지 못한다
		for _, obj := range objects {
			if v, has := obj["level"]; has && toID(v) > int64(s.session.Level) {
				return false
			}
		}
		ids := []int64{}
		for _, obj := range objects {
			if id := toID(obj["id"]); id > 0 && id != s.session.Id {
				ids = append(ids, id)
			}
		}
		if len(ids) > 0 {
			holders := strings.TrimSuffix(strings.Repeat("?,", len(ids)), ",")
			params := []interface{}{}
			for _, id := range ids {
				params = append(params, id)
			}
			var maxLevel sql.NullInt64
			if err := s.conn.QueryRow("select max(u_level) from user_tb where u_id in ("+holders+")", params...).Scan(&maxLevel); err != nil {
				return false
			}
			if maxLevel.Valid && maxLevel.Int64 > int64(s.session.Level) {
				return false
			}
		}
		return true
	}

	switch shape {
	case shapeFieldUpdate:
		if userPrivilegedFields[field] {
			return false
		}
	case shapeUpdate:
	default:
		return false
	}

	obj := objects[0]
	if toID(obj["id"]) != s.session.Id {
		return false
	}

	if shape == shapeUpdate {
		var level, approval, status int
		var score float64
		err := s.conn.QueryRow("select u_level, u_approval, u_status, u_score from user_tb where u_id = ?", s.session.Id).Scan(&level, &approval, &status, &score)
		if err != nil {
			return false
		}
		obj["level"] = level
		obj["approval"] = approval
		obj["status"] = status
		obj["score"] = score
	}

	return true
}

// filterResponse 는 핸들러를 실행한 뒤 응답의 item/items 중 내 업체 것이 아닌 행을 걸러낸다
func (s *tenantScope) filterResponse(c *fiber.Ctx, entity string, rule entityRule) error {
	if err := c.Next(); err != nil {
		return err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(c.Response().Body(), &result); err != nil {
		return nil
	}

	changed := false
	S := float64(s.company())

	failed := false
	keep := func(rows []map[string]interface{}) map[int]bool {
		ok := map[int]bool{}
		if rule.kind == kindTenant {
			for i, row := range rows {
				if v, has := row["company"]; has && v == S {
					ok[i] = true
				}
			}
			return ok
		}

		ids := []int64{}
		for _, row := range rows {
			ids = append(ids, toID(row["id"]))
		}
		owned, err := s.ownedIDs(rule, ids)
		if err != nil {
			failed = true
			return ok
		}
		for i, row := range rows {
			if owned[toID(row["id"])] {
				ok[i] = true
			}
		}
		return ok
	}

	if raw, has := result["items"].([]interface{}); has {
		rows := []map[string]interface{}{}
		for _, r := range raw {
			if m, isMap := r.(map[string]interface{}); isMap {
				rows = append(rows, m)
			}
		}
		ok := keep(rows)
		filtered := []interface{}{}
		for i, row := range rows {
			if ok[i] {
				filtered = append(filtered, row)
			}
		}
		if len(filtered) != len(raw) {
			if total, isNum := result["total"].(float64); isNum {
				result["total"] = total - float64(len(raw)-len(filtered))
			}
			result["items"] = filtered
			changed = true
		}
	}

	if row, has := result["item"].(map[string]interface{}); has {
		if !keep([]map[string]interface{}{row})[0] {
			result["item"] = nil
			changed = true
		}
	}

	if failed {
		return serverError(c)
	}

	if changed {
		return c.JSON(result)
	}

	return nil
}
