package api

import (
	"fmt"
	"log"
	"path"
	"zkeep/config"
	"zkeep/controllers"
	"zkeep/global"
	"zkeep/models"

	"strings"
)

type CustomerController struct {
	controllers.Controller
}

func (c *CustomerController) Index(page int, pagesize int) {

	conn := c.NewConnection()

	manager := models.NewCustomerManager(conn)

	var args []interface{}

	_name := c.Get("name")
	if _name != "" {
		args = append(args, models.Custom{Query: fmt.Sprintf("(b_name like '%%%v%%' or c_name like '%%%v%%')", _name, _name)})
	}

	_type := c.Geti("type")
	if _type != 0 {
		args = append(args, models.Where{Column: "type", Value: _type, Compare: "="})
	}
	_checkdate := c.Geti("checkdate")
	if _checkdate != 0 {
		args = append(args, models.Where{Column: "checkdate", Value: _checkdate, Compare: "="})
	}
	_managername := c.Get("managername")
	if _managername != "" {
		args = append(args, models.Where{Column: "managername", Value: _managername, Compare: "like"})
	}
	_managertel := c.Get("managertel")
	if _managertel != "" {
		args = append(args, models.Where{Column: "managertel", Value: _managertel, Compare: "like"})
	}
	_manageremail := c.Get("manageremail")
	if _manageremail != "" {
		args = append(args, models.Where{Column: "manageremail", Value: _manageremail, Compare: "like"})
	}
	_startcontractstartdate := c.Get("startcontractstartdate")
	_endcontractstartdate := c.Get("endcontractstartdate")
	if _startcontractstartdate != "" && _endcontractstartdate != "" {
		var v [2]string
		v[0] = _startcontractstartdate
		v[1] = _endcontractstartdate
		args = append(args, models.Where{Column: "contractstartdate", Value: v, Compare: "between"})
	} else if _startcontractstartdate != "" {
		args = append(args, models.Where{Column: "contractstartdate", Value: _startcontractstartdate, Compare: ">="})
	} else if _endcontractstartdate != "" {
		args = append(args, models.Where{Column: "contractstartdate", Value: _endcontractstartdate, Compare: "<="})
	}
	_startcontractenddate := c.Get("startcontractenddate")
	_endcontractenddate := c.Get("endcontractenddate")
	if _startcontractenddate != "" && _endcontractenddate != "" {
		var v [2]string
		v[0] = _startcontractenddate
		v[1] = _endcontractenddate
		args = append(args, models.Where{Column: "contractenddate", Value: v, Compare: "between"})
	} else if _startcontractenddate != "" {
		args = append(args, models.Where{Column: "contractenddate", Value: _startcontractenddate, Compare: ">="})
	} else if _endcontractenddate != "" {
		args = append(args, models.Where{Column: "contractenddate", Value: _endcontractenddate, Compare: "<="})
	}
	_contractprice := c.Geti("contractprice")
	if _contractprice != 0 {
		args = append(args, models.Where{Column: "contractprice", Value: _contractprice, Compare: "="})
	}
	_contractvat := c.Geti("contractvat")
	if _contractvat != 0 {
		args = append(args, models.Where{Column: "contractvat", Value: _contractvat, Compare: "="})
	}
	_contractday := c.Geti("contractday")
	if _contractday != 0 {
		args = append(args, models.Where{Column: "contractday", Value: _contractday, Compare: "="})
	}
	_billingdate := c.Geti("billingdate")
	if _billingdate != 0 {
		args = append(args, models.Where{Column: "billingdate", Value: _billingdate, Compare: "="})
	}
	_billingtype := c.Geti("billingtype")
	if _billingtype != 0 {
		args = append(args, models.Where{Column: "billingtype", Value: _billingtype, Compare: "="})
	}
	_billingname := c.Get("billingname")
	if _billingname != "" {
		args = append(args, models.Where{Column: "billingname", Value: _billingname, Compare: "like"})
	}
	_billingtel := c.Get("billingtel")
	if _billingtel != "" {
		args = append(args, models.Where{Column: "billingtel", Value: _billingtel, Compare: "like"})
	}
	_billingemail := c.Get("billingemail")
	if _billingemail != "" {
		args = append(args, models.Where{Column: "billingemail", Value: _billingemail, Compare: "like"})
	}
	_address := c.Get("address")
	if _address != "" {
		args = append(args, models.Where{Column: "address", Value: _address, Compare: "like"})
	}
	_addressetc := c.Get("addressetc")
	if _addressetc != "" {
		args = append(args, models.Where{Column: "addressetc", Value: _addressetc, Compare: "like"})
	}
	_collectmonth := c.Geti("collectmonth")
	if _collectmonth != 0 {
		args = append(args, models.Where{Column: "collectmonth", Value: _collectmonth, Compare: "="})
	}
	_collectday := c.Geti("collectday")
	if _collectday != 0 {
		args = append(args, models.Where{Column: "collectday", Value: _collectday, Compare: "="})
	}
	_manager := c.Get("manager")
	if _manager != "" {
		args = append(args, models.Where{Column: "manager", Value: _manager, Compare: "like"})
	}
	_tel := c.Get("tel")
	if _tel != "" {
		args = append(args, models.Where{Column: "tel", Value: _tel, Compare: "like"})
	}
	_fax := c.Get("fax")
	if _fax != "" {
		args = append(args, models.Where{Column: "fax", Value: _fax, Compare: "like"})
	}
	_status := c.Geti("status")
	if _status != 0 {
		args = append(args, models.Where{Column: "status", Value: _status, Compare: "="})
	}
	_user := c.Geti64("user")
	if _user != 0 {
		args = append(args, models.Where{Column: "user", Value: _user, Compare: "="})
	}
	_company := c.Geti64("company")
	if _company != 0 {
		args = append(args, models.Where{Column: "company", Value: _company, Compare: "="})
	}
	_building := c.Geti64("building")
	if _building != 0 {
		args = append(args, models.Where{Column: "building", Value: _building, Compare: "="})
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
					str += ", cu_" + strings.Trim(v, " ")
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

func (c *CustomerController) Status(id int64) {
	conn := c.NewConnection()

	session := c.Session

	customerManager := models.NewCustomerManager(conn)
	items := customerManager.Find([]interface{}{
		models.Where{Column: "company", Value: session.Company, Compare: "="},
		models.Where{Column: "status", Value: 1, Compare: "="},
	})

	user := customerManager.Count([]interface{}{
		models.Where{Column: "company", Value: session.Company, Compare: "="},
		models.Where{Column: "status", Value: 2, Compare: "="},
	})

	var money int64 = 0
	var score float32 = 0.0

	for _, v := range items {
		money += int64(v.Contractprice)
		money += int64(v.Contractvat)

		building := v.Extra["building"].(models.Building)

		score += float32(building.Score)
	}

	log.Println("customer status", id)

	c.Set("currentuser", len(items))
	c.Set("user", user)
	c.Set("money", money)
	c.Set("score", score)

}

func (c *CustomerController) Upload(filename string) {
	conn := c.NewConnection()

	companyManager := models.NewCompanyManager(conn)
	buildingManager := models.NewBuildingManager(conn)
	customerManager := models.NewCustomerManager(conn)
	customercompanyManager := models.NewCustomercompanyManager(conn)
	userManager := models.NewUserManager(conn)

	user := c.Session

	fullFilename := path.Join(config.UploadPath, filename)
	f := global.NewExcelReader(fullFilename)
	if f == nil {
		log.Println("not found file")
		return
	}

	sheet := "Sheet1"
	f.SetSheet(sheet)

	pos := 2
	for {
		name := f.GetCell("A", pos)

		if name == "" {
			log.Println("brake")
			break
		}

		companyno := f.GetCell("B", pos)

		item := companyManager.GetByCompanyno(companyno)
		if item == nil {
			item = &models.Company{}
			item.Companyno = companyno
			item.Name = name
			item.Ceo = f.GetCell("C", pos)
			item.Address = f.GetCell("G", pos)
			item.Tel = f.GetCell("J", pos)
			item.Email = f.GetCell("K", pos)

			companyManager.Insert(item)

			item.Id = companyManager.GetIdentity()
		}

		customercompany := &models.Customercompany{}
		customercompany.Company = user.Company
		customercompany.Customer = item.Id
		customercompanyManager.Insert(customercompany)

		buildingName := f.GetCell("D", pos)
		building := buildingManager.GetByCompanyName(item.Id, buildingName)

		if building == nil {
			building = &models.Building{}
			building.Name = buildingName
			building.Companyno = f.GetCell("E", pos)
			building.Ceo = f.GetCell("F", pos)
			if building.Companyno == "" {
				building.Companyno = item.Companyno
			}
			if building.Ceo == "" {
				building.Ceo = item.Ceo
			}
			building.Address = f.GetCell("G", pos)

			buildingManager.Insert(building)

			building.Id = buildingManager.GetIdentity()
		}

		customer := customerManager.GetByCompanyBuilding(item.Id, building.Id)
		if customer == nil {
			customer = &models.Customer{}
			customer.Company = item.Id
			customer.Building = building.Id
		}

		typeid := f.GetCell("H", pos)
		if typeid == "1" || typeid == "직영" {
			customer.Type = 1
		} else {
			customer.Type = 2
		}

		customer.Managername = f.GetCell("I", pos)
		customer.Managertel = f.GetCell("J", pos)
		customer.Manageremail = f.GetCell("K", pos)
		customer.Billingname = f.GetCell("L", pos)
		customer.Billingtel = f.GetCell("M", pos)
		customer.Billingemail = f.GetCell("N", pos)
		customer.Contractprice = global.Atoi(f.GetCell("O", pos))
		customer.Contractvat = global.Atoi(f.GetCell("P", pos))
		customer.Collectday = global.Atoi(f.GetCell("Q", pos))
		if customer.Collectmonth == 0 {
			customer.Collectmonth = 1
		}

		billingtype := f.GetCell("H", pos)
		if billingtype == "1" || billingtype == "계산서" {
			customer.Billingtype = 1
		} else {
			customer.Billingtype = 2
		}

		userName := f.GetCell("S", pos)
		if userName != "" {
			user := userManager.GetByCompanyName(user.Company, userName)
			if user != nil {
				customer.User = user.Id
			}
		}

		if customer.Id == 0 {
			customerManager.Insert(customer)
		} else {
			customerManager.Update(customer)
		}

		pos++
	}
}
