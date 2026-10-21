package api

import (
	"fmt"
	"log"
	"path"
	"strings"
	"zkeep/config"
	"zkeep/controllers"
	"zkeep/global"
	"zkeep/models"
	"zkeep/models/user"
)

type UserController struct {
	controllers.Controller
}

func (c *UserController) Search() {

	log.Println("search")

	if c.Session == nil {
		c.Error("invalid request")
		return
	}

	conn := c.NewConnection()

	manager := models.NewUserManager(conn)

	var args []interface{}

	_loginid := c.Get("loginid")
	if _loginid != "" {
		args = append(args, models.Where{Column: "loginid", Value: _loginid, Compare: "like"})
	}
	_name := c.Get("name")
	if _name != "" {
		name := models.EscapeString(_name)
		query := fmt.Sprintf("(u_loginid like '%%%v%%' or u_name like '%%%v%%' or u_email like '%%%v%%')", name, name, name)
		args = append(args, models.Custom{Query: query})
	}
	_email := c.Get("email")
	if _email != "" {
		args = append(args, models.Where{Column: "email", Value: _email, Compare: "like"})
	}
	_address := c.Get("address")
	if _address != "" {
		args = append(args, models.Where{Column: "address", Value: _address, Compare: "like"})
	}
	_addressetc := c.Get("addressetc")
	if _addressetc != "" {
		args = append(args, models.Where{Column: "addressetc", Value: _addressetc, Compare: "like"})
	}
	_startjoindate := c.Get("startjoindate")
	_endjoindate := c.Get("endjoindate")
	if _startjoindate != "" && _endjoindate != "" {
		var v [2]string
		v[0] = _startjoindate
		v[1] = _endjoindate
		args = append(args, models.Where{Column: "joindate", Value: v, Compare: "between"})
	} else if _startjoindate != "" {
		args = append(args, models.Where{Column: "joindate", Value: _startjoindate, Compare: ">="})
	} else if _endjoindate != "" {
		args = append(args, models.Where{Column: "joindate", Value: _endjoindate, Compare: "<="})
	}
	_careeryear := c.Geti("careeryear")
	if _careeryear != 0 {
		args = append(args, models.Where{Column: "careeryear", Value: _careeryear, Compare: "="})
	}
	_careermonth := c.Geti("careermonth")
	if _careermonth != 0 {
		args = append(args, models.Where{Column: "careermonth", Value: _careermonth, Compare: "="})
	}
	_level := c.Geti("level")
	if _level != 0 {
		args = append(args, models.Where{Column: "level", Value: _level, Compare: "="})
	}
	_score := c.Geti("score")
	if _score != 0 {
		args = append(args, models.Where{Column: "score", Value: _score, Compare: "="})
	}
	_approval := c.Geti("approval")
	if _approval != 0 {
		args = append(args, models.Where{Column: "approval", Value: _approval, Compare: "="})
	}
	_status := c.Geti("status")
	if _status != 0 {
		args = append(args, models.Where{Column: "status", Value: _status, Compare: "="})
	}
	_company := c.TenantCompany(c.Geti64("company"))
	if _company != 0 {
		args = append(args, models.Where{Column: "company", Value: _company, Compare: "="})
	}
	_startdate := c.Get("startdate")
	_enddate := c.Get("enddate")
	if _startdate != "" && _enddate != "" {
		var v [2]string
		v[0] = _startdate
		v[1] = _enddate
		args = append(args, models.Where{Column: "date", Value: v, Compare: "between"})
	} else if _startdate != "" {
		args = append(args, models.Where{Column: "date", Value: _startdate, Compare: ">="})
	} else if _enddate != "" {
		args = append(args, models.Where{Column: "date", Value: _enddate, Compare: "<="})
	}

	page := 0
	pagesize := 0
	if page != 0 && pagesize != 0 {
		args = append(args, models.Paging(page, pagesize))
	}

	orderby := c.Get("orderby")
	if orderby == "" {
		if page != 0 && pagesize != 0 {
			orderby = "id desc"
			args = append(args, models.Ordering(orderby))
		}
	} else {
		orderbys := strings.Split(orderby, ",")

		str := ""
		for i, v := range orderbys {
			if i == 0 {
				str += v
			} else {
				if strings.Contains(v, "_") {
					str += ", " + strings.Trim(v, " ")
				} else {
					str += ", u_" + strings.Trim(v, " ")
				}
			}
		}

		args = append(args, models.Ordering(str))
	}

	items := manager.Find(args)
	c.Set("items", items)

	total := manager.Count(args)
	c.Set("total", total)
}

func (c *UserController) Upload(filename string) {
	conn := c.NewConnection()

	departmentManager := models.NewDepartmentManager(conn)
	userManager := models.NewUserManager(conn)

	session := c.Session

	if session == nil || !controllers.IsSafeFilename(filename) {
		c.Error("invalid request")
		return
	}

	fullFilename := path.Join(config.UploadPath, filename)
	f := global.NewExcelReader(fullFilename)
	if f == nil {
		log.Println("not found file")
		return
	}
	defer f.Close()

	sheet := "Sheet1"
	f.SetSheet(sheet)

	skipped := make([]string, 0)

	pos := 1
	for {
		name := f.GetCell("C", pos)

		log.Println(name)
		if name == "" {
			log.Println("brake")
			break
		}

		departmentName := f.GetCell("A", pos)

		if departmentName == "팀" && name == "이름" {
			pos++
			continue
		}

		userItem := userManager.GetByCompanyName(session.Company, name)

		// 새 계정은 K열 초기비밀번호가 있어야 등록 (기존 계정의 비밀번호는 바꾸지 않음)
		passwd := strings.TrimSpace(f.GetCell("K", pos))
		if userItem == nil && passwd == "" {
			skipped = append(skipped, name)
			pos++
			continue
		}
		department := departmentManager.GetByCompanyName(session.Company, departmentName)
		if department == nil {
			department = &models.Department{
				Name:    departmentName,
				Status:  1,
				Company: session.Company,
			}

			departmentManager.Insert(department)

			department.Id = departmentManager.GetIdentity()
		}

		loginid := f.GetCell("B", pos)
		email := f.GetCell("D", pos)
		tel := f.GetCell("E", pos)
		address := f.GetCell("F", pos)
		levelStr := f.GetCell("G", pos)
		statusStr := f.GetCell("H", pos)
		scoreStr := f.GetCell("I", pos)
		date := f.GetCell("J", pos)

		if userItem != nil {
			if userItem.Loginid == loginid {
				pos++
				continue
			}
		}

		level := user.LevelNormal
		if levelStr == "팀장" {
			level = user.LevelManager
		} else if levelStr == "관리자" {
			level = user.LevelAdmin
		}

		status := user.StatusUse
		if statusStr == "사용안함" {
			status = user.StatusNotuse
		}

		score := 0
		if strings.Contains(scoreStr, "/") {
			temp := strings.Split(scoreStr, "/")
			if len(temp) == 2 {
				score = global.Atoi(temp[1])
			}
		} else {
			score = global.Atoi(scoreStr)
		}

		if score == 0 {
			score = 60
		}

		if userItem != nil {
			userItem.Email = email
			userItem.Tel = tel
			userItem.Address = address
			userItem.Level = level
			userItem.Status = status

			userItem.Score = models.Double(score)
			userItem.Company = session.Company
			userItem.Department = department.Id
			userItem.Approval = user.ApprovalComplete

			userManager.Update(userItem)
		} else {
			user := models.User{
				Name:       name,
				Loginid:    loginid,
				Passwd:     global.HashPassword(passwd),
				Email:      email,
				Tel:        tel,
				Address:    address,
				Level:      level,
				Status:     status,
				Score:      models.Double(score),
				Company:    session.Company,
				Department: department.Id,
				Approval:   user.ApprovalComplete,
				Date:       date,
			}

			err := userManager.Insert(&user)
			if err != nil {
				log.Println(err)
			}
		}

		pos++
	}

	// 초기비밀번호가 없어 등록하지 않은 이름
	c.Set("skipped", skipped)
}

// 삭제 가능한 사용자인지 DB에 저장된 값으로 확인 (클라이언트가 보낸 level은 신뢰하지 않음)
func (c *UserController) deletable(old *models.User) bool {
	if old == nil {
		return false
	}

	if !c.IsRootadmin() && old.Company != c.Session.Company {
		return false
	}

	if old.Level == user.LevelAdmin || old.Level == user.LevelRootadmin {
		return false
	}

	return true
}

func (c *UserController) Delete(item *models.User) {
	if c.Session == nil || item == nil {
		c.Error("invalid request")
		return
	}

	conn := c.NewConnection()

	manager := models.NewUserManager(conn)
	old := manager.Get(item.Id)
	if old == nil || (!c.IsRootadmin() && old.Company != c.Session.Company) {
		c.Error("user not found")
		return
	}

	if !c.deletable(old) {
		return
	}

	manager.Delete(item.Id)
}

func (c *UserController) Deletebatch(item *[]models.User) {
	if c.Session == nil || item == nil {
		c.Error("invalid request")
		return
	}

	conn := c.NewConnection()

	manager := models.NewUserManager(conn)

	for _, v := range *item {
		if !c.deletable(manager.Get(v.Id)) {
			continue
		}

		manager.Delete(v.Id)
	}
}
