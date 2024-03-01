package rest

import (
	"log"
	"zkeep/controllers"
	"zkeep/models"

	"zkeep/models/user"

	"strings"
)

type UserController struct {
	controllers.Controller
}

func (c *UserController) Read(id int64) {

	conn := c.NewConnection()

	manager := models.NewUserManager(conn)
	item := manager.Get(id)

	c.Set("item", item)
}

func (c *UserController) Index(page int, pagesize int) {

	conn := c.NewConnection()

	manager := models.NewUserManager(conn)

	var args []interface{}

	_loginid := c.Get("loginid")
	if _loginid != "" {
		args = append(args, models.Where{Column: "loginid", Value: _loginid, Compare: "like"})
	}
	_passwd := c.Get("passwd")
	if _passwd != "" {
		args = append(args, models.Where{Column: "passwd", Value: _passwd, Compare: "like"})
	}
	_name := c.Get("name")
	if _name != "" {
		args = append(args, models.Where{Column: "name", Value: _name, Compare: "="})

	}
	_email := c.Get("email")
	if _email != "" {
		args = append(args, models.Where{Column: "email", Value: _email, Compare: "like"})
	}
	_tel := c.Get("tel")
	if _tel != "" {
		args = append(args, models.Where{Column: "tel", Value: _tel, Compare: "like"})
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
	_status := c.Geti("status")
	if _status != 0 {
		args = append(args, models.Where{Column: "status", Value: _status, Compare: "="})
	}
	_company := c.Geti64("company")
	if _company != 0 {
		args = append(args, models.Where{Column: "company", Value: _company, Compare: "="})
	}
	_department := c.Geti64("department")
	if _department != 0 {
		args = append(args, models.Where{Column: "department", Value: _department, Compare: "="})
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

func (c *UserController) Insert(item *models.User) {

	conn := c.NewConnection()

	manager := models.NewUserManager(conn)
	manager.Insert(item)

	id := manager.GetIdentity()
	c.Result["id"] = id
	item.Id = id
}

func (c *UserController) Insertbatch(item *[]models.User) {
	if item == nil || len(*item) == 0 {
		return
	}

	rows := len(*item)

	conn := c.NewConnection()

	manager := models.NewUserManager(conn)

	for i := 0; i < rows; i++ {
		manager.Insert(&((*item)[i]))
	}
}

func (c *UserController) Update(item *models.User) {

	conn := c.NewConnection()

	manager := models.NewUserManager(conn)
	manager.Update(item)
}

func (c *UserController) Delete(item *models.User) {

	conn := c.NewConnection()

	manager := models.NewUserManager(conn)

	manager.Delete(item.Id)
}

func (c *UserController) Deletebatch(item *[]models.User) {

	conn := c.NewConnection()

	manager := models.NewUserManager(conn)

	for _, v := range *item {

		manager.Delete(v.Id)
	}
}

func (c *UserController) GetByLoginid(loginid string) *models.User {

	conn := c.NewConnection()

	_manager := models.NewUserManager(conn)

	item := _manager.GetByLoginid(loginid)

	c.Set("item", item)

	return item

}

func (c *UserController) CountByLoginid(loginid string) int {

	conn := c.NewConnection()

	_manager := models.NewUserManager(conn)

	item := _manager.CountByLoginid(loginid)

	c.Set("count", item)

	return item

}

func (c *UserController) FindByLevel(level user.Level) []models.User {

	conn := c.NewConnection()

	_manager := models.NewUserManager(conn)

	item := _manager.FindByLevel(level)

	c.Set("items", item)

	return item

}

// @Put()
func (c *UserController) UpdateLoginid(loginid string, id int64) {

	conn := c.NewConnection()

	_manager := models.NewUserManager(conn)
	_manager.UpdateLoginid(loginid, id)
}

// @Put()
func (c *UserController) UpdatePasswd(passwd string, id int64) {

	conn := c.NewConnection()

	_manager := models.NewUserManager(conn)
	_manager.UpdatePasswd(passwd, id)
}

// @Put()
func (c *UserController) UpdateName(name string, id int64) {

	conn := c.NewConnection()

	_manager := models.NewUserManager(conn)
	_manager.UpdateName(name, id)
}

// @Put()
func (c *UserController) UpdateEmail(email string, id int64) {

	conn := c.NewConnection()

	_manager := models.NewUserManager(conn)
	_manager.UpdateEmail(email, id)
}

// @Put()
func (c *UserController) UpdateTel(tel string, id int64) {

	conn := c.NewConnection()

	_manager := models.NewUserManager(conn)
	_manager.UpdateTel(tel, id)
}

// @Put()
func (c *UserController) UpdateAddress(address string, id int64) {

	conn := c.NewConnection()

	_manager := models.NewUserManager(conn)
	_manager.UpdateAddress(address, id)
}

// @Put()
func (c *UserController) UpdateAddressetc(addressetc string, id int64) {

	conn := c.NewConnection()

	_manager := models.NewUserManager(conn)
	_manager.UpdateAddressetc(addressetc, id)
}

// @Put()
func (c *UserController) UpdateJoindate(joindate string, id int64) {

	conn := c.NewConnection()

	_manager := models.NewUserManager(conn)
	_manager.UpdateJoindate(joindate, id)
}

// @Put()
func (c *UserController) UpdateCareeryear(careeryear int, id int64) {

	conn := c.NewConnection()

	_manager := models.NewUserManager(conn)
	_manager.UpdateCareeryear(careeryear, id)
}

// @Put()
func (c *UserController) UpdateCareermonth(careermonth int, id int64) {

	conn := c.NewConnection()

	_manager := models.NewUserManager(conn)
	_manager.UpdateCareermonth(careermonth, id)
}

// @Put()
func (c *UserController) UpdateLevel(level int, id int64) {

	conn := c.NewConnection()

	_manager := models.NewUserManager(conn)
	_manager.UpdateLevel(level, id)
}

// @Put()
func (c *UserController) UpdateScore(score models.Double, id int64) {

	conn := c.NewConnection()

	_manager := models.NewUserManager(conn)
	_manager.UpdateScore(score, id)
}

// @Put()
func (c *UserController) UpdateStatus(status int, id int64) {

	conn := c.NewConnection()

	_manager := models.NewUserManager(conn)
	_manager.UpdateStatus(status, id)
}

// @Put()
func (c *UserController) UpdateCompany(company int64, id int64) {

	conn := c.NewConnection()

	_manager := models.NewUserManager(conn)
	_manager.UpdateCompany(company, id)
}

// @Put()
func (c *UserController) UpdateDepartment(department int64, id int64) {

	conn := c.NewConnection()

	_manager := models.NewUserManager(conn)
	_manager.UpdateDepartment(department, id)
}

func (c *UserController) Sum() {
	log.Println("rest sum")

	conn := c.NewConnection()

	manager := models.NewUserManager(conn)

	var args []interface{}

	_loginid := c.Get("loginid")
	if _loginid != "" {
		args = append(args, models.Where{Column: "loginid", Value: _loginid, Compare: "like"})
	}
	_passwd := c.Get("passwd")
	if _passwd != "" {
		args = append(args, models.Where{Column: "passwd", Value: _passwd, Compare: "like"})
	}
	_name := c.Get("name")
	if _name != "" {
		args = append(args, models.Where{Column: "name", Value: _name, Compare: "="})

	}
	_email := c.Get("email")
	if _email != "" {
		args = append(args, models.Where{Column: "email", Value: _email, Compare: "like"})
	}
	_tel := c.Get("tel")
	if _tel != "" {
		args = append(args, models.Where{Column: "tel", Value: _tel, Compare: "like"})
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
	_status := c.Geti("status")
	if _status != 0 {
		args = append(args, models.Where{Column: "status", Value: _status, Compare: "="})
	}
	_company := c.Geti64("company")
	if _company != 0 {
		args = append(args, models.Where{Column: "company", Value: _company, Compare: "="})
	}
	_department := c.Geti64("department")
	if _department != 0 {
		args = append(args, models.Where{Column: "department", Value: _department, Compare: "="})
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

	item := manager.Sum(args)
	c.Set("item", item)
}
