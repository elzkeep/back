package api

import (
	"zkeep/controllers"
	"zkeep/models"

	"strings"
)

type BillinguserlistController struct {
	controllers.Controller
}

func (c *BillinguserlistController) Excel(company int64, startdate string, enddate string, users []int64) {

	conn := c.NewConnection()

	manager := models.NewBillinguserlistManager(conn)

	var args []interface{}

	_price := c.Geti("price")
	if _price != 0 {
		args = append(args, models.Where{Column: "price", Value: _price, Compare: "="})
	}
	_status := c.Geti("status")
	if _status != 0 {
		args = append(args, models.Where{Column: "status", Value: _status, Compare: "="})
	}
	_giro := c.Geti("giro")
	if _giro != 0 {
		args = append(args, models.Where{Column: "giro", Value: _giro, Compare: "="})
	}
	_startbilldate := c.Get("startbilldate")
	_endbilldate := c.Get("endbilldate")
	if _startbilldate != "" && _endbilldate != "" {
		var v [2]string
		v[0] = _startbilldate
		v[1] = _endbilldate
		args = append(args, models.Where{Column: "billdate", Value: v, Compare: "between"})
	} else if _startbilldate != "" {
		args = append(args, models.Where{Column: "billdate", Value: _startbilldate, Compare: ">="})
	} else if _endbilldate != "" {
		args = append(args, models.Where{Column: "billdate", Value: _endbilldate, Compare: "<="})
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
	_buildingname := c.Get("buildingname")
	if _buildingname != "" {
		args = append(args, models.Where{Column: "buildingname", Value: _buildingname, Compare: "like"})
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
	_user := c.Geti64("user")
	if _user != 0 {
		args = append(args, models.Where{Column: "user", Value: _user, Compare: "="})
	}
	_username := c.Get("username")
	if _username != "" {
		args = append(args, models.Where{Column: "username", Value: _username, Compare: "like"})
	}

	if len(users) > 0 {
		args = append(args, models.Where{Column: "user", Value: users, Compare: "in"})
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
					str += ", bi_" + strings.Trim(v, " ")
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
