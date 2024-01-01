package api

import (
	"fmt"
	"zkeep/controllers"
	"zkeep/models"

	"strings"
)

type ReportController struct {
	controllers.Controller
}

func (c *ReportController) Search(page int, pagesize int) {

	conn := c.NewConnection()

	manager := models.NewReportManager(conn)

	var args []interface{}

	_search := c.Get("search")
	if _search != "" {
		args = append(args, models.Custom{Query: fmt.Sprintf("(r_title like '%%%v%%' or c_buildingname like '%%%v%%' or c_address like '%%%v%%')", _search)})
	}

	_title := c.Get("title")
	if _title != "" {
		args = append(args, models.Where{Column: "title", Value: _title, Compare: "="})

	}
	_period := c.Geti("period")
	if _period != 0 {
		args = append(args, models.Where{Column: "period", Value: _period, Compare: "="})
	}
	_number := c.Geti("number")
	if _number != 0 {
		args = append(args, models.Where{Column: "number", Value: _number, Compare: "="})
	}
	_checkdate := c.Get("checkdate")
	if _checkdate != "" {
		args = append(args, models.Where{Column: "checkdate", Value: _checkdate, Compare: "like"})
	}
	_checktime := c.Get("checktime")
	if _checktime != "" {
		args = append(args, models.Where{Column: "checktime", Value: _checktime, Compare: "like"})
	}
	_content := c.Get("content")
	if _content != "" {
		args = append(args, models.Where{Column: "content", Value: _content, Compare: "="})

	}
	_status := c.Geti("status")
	if _status != 0 {
		args = append(args, models.Where{Column: "status", Value: _status, Compare: "="})
	}
	_company := c.Geti64("company")
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
					str += ", r_" + strings.Trim(v, " ")
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

func (c *ReportController) Insert(item *models.Report) {

	conn := c.NewConnection()

	manager := models.NewReportManager(conn)
	manager.Insert(item)

	id := manager.GetIdentity()
	c.Result["id"] = id
	item.Id = id
}
