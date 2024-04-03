package api

import (
	"log"
	"path"
	"strings"
	"zkeep/config"
	"zkeep/controllers"
	"zkeep/global"
	"zkeep/models"
)

type CompanyController struct {
	controllers.Controller
}

func (c *CompanyController) Search(page int, pagesize int) {
	conn := c.NewConnection()

	user := c.Session

	manager := models.NewCompanyManager(conn)

	var args []interface{}

	args = append(args, models.Where{Column: "company", Value: user.Company, Compare: "="})

	_name := c.Get("name")
	if _name != "" {
		args = append(args, models.Where{Column: "name", Value: _name, Compare: "like"})

	}
	_companyno := c.Get("companyno")
	if _companyno != "" {
		args = append(args, models.Where{Column: "companyno", Value: _companyno, Compare: "like"})
	}
	_ceo := c.Get("ceo")
	if _ceo != "" {
		args = append(args, models.Where{Column: "ceo", Value: _ceo, Compare: "like"})
	}
	_tel := c.Get("tel")
	if _tel != "" {
		args = append(args, models.Where{Column: "tel", Value: _tel, Compare: "like"})
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
					str += ", c_" + strings.Trim(v, " ")
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

func (c *CompanyController) Post_Insert(item *models.Company) {
	conn := c.NewConnection()

	customercompanyManager := models.NewCustomercompanyManager(conn)

	log.Println("item", item)
	user := c.Session

	log.Println("user", user)
	customercompany := models.Customercompany{Company: user.Company, Customer: item.Id}
	customercompanyManager.Insert(&customercompany)
}

func (c *CompanyController) Upload(filename string) {
	conn := c.NewConnection()

	companyManager := models.NewCompanyManager(conn)
	customercompanyManager := models.NewCustomercompanyManager(conn)

	user := c.Session

	fullFilename := path.Join(config.UploadPath, filename)
	f := global.NewExcelReader(fullFilename)
	if f == nil {
		log.Println("not found file")
		return
	}

	sheet := "Sheet1"
	f.SetSheet(sheet)

	pos := 1
	for {
		item := &models.Company{}
		item.Name = f.GetCell("A", pos)

		if item.Name == "" {
			log.Println("brake")
			break
		}

		item.Companyno = f.GetCell("B", pos)
		item.Ceo = f.GetCell("C", pos)
		item.Address = f.GetCell("D", pos)
		item.Addressetc = f.GetCell("E", pos)
		item.Tel = f.GetCell("F", pos)
		item.Email = f.GetCell("G", pos)

		log.Println(item)
		companyManager.Insert(item)

		id := companyManager.GetIdentity()

		customercompany := &models.Customercompany{}
		customercompany.Company = user.Company
		customercompany.Customer = id
		customercompanyManager.Insert(customercompany)

		pos++
	}
}
