package api

import (
	"fmt"
	"log"
	"net/url"
	"zkeep/controllers"
	"zkeep/models"

	"github.com/gofiber/fiber/v2"

	"strings"
)

type ReportController struct {
	controllers.Controller
}

func (c *ReportController) Search(page int, pagesize int) {
	log.Println("Search")

	if c.Session == nil {
		c.Error("invalid request")
		return
	}

	conn := c.NewConnection()

	manager := models.NewReportManager(conn)

	var args []interface{}

	_search := c.Get("search")
	if _search != "" {
		search := models.EscapeString(_search)
		args = append(args, models.Custom{Query: fmt.Sprintf("(r_title like '%%%v%%' or b_name like '%%%v%%' or c_address like '%%%v%%')", search, search, search)})
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
	_user := c.Geti64("user")
	if _user != 0 {
		args = append(args, models.Where{Column: "user", Value: _user, Compare: "="})
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

	log.Println(args)
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

	log.Println(args)
	items := manager.Find(args)
	c.Set("items", items)

	total := manager.Count(args)
	c.Set("total", total)
}

func (c *ReportController) Insert(item *models.Report) {
	if c.Session == nil || item == nil {
		c.Error("invalid request")
		return
	}

	// 전체관리자 외에는 본인 회사로만 등록
	if !c.IsRootadmin() {
		item.Company = c.Session.Company
	}

	conn := c.NewConnection()

	manager := models.NewReportManager(conn)
	manager.Insert(item)

	id := manager.GetIdentity()
	c.Result["id"] = id
	item.Id = id
}

// GET /api/report/download/:id — 점검 보고서 PDF (reportpdf.go). 다른 업체 보고서는 거부
func (c *ReportController) Download(id int64) {
	ctx := c.Context

	if c.Session == nil {
		ctx.Status(fiber.StatusUnauthorized)
		return
	}

	conn := c.NewConnection()
	report := models.NewReportManager(conn).Get(id)
	if report == nil || (!c.IsRootadmin() && report.Company != c.Session.Company) {
		ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"code": "error", "message": "report not found"})
		return
	}

	buf, err := BuildReportPdf(loadReportData(conn, report))
	if err != nil {
		log.Println("report pdf :", err)
		ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"code": "error", "message": "pdf error"})
		return
	}

	date := report.Checkdate
	if len(date) >= 10 {
		date = date[:10]
	}
	filename := fmt.Sprintf("전기설비점검기록표_%v_%v.pdf", date, report.Id)

	ctx.Set("Content-Type", "application/pdf")
	ctx.Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"report_%v.pdf\"; filename*=UTF-8''%v", report.Id, url.PathEscape(filename)))
	ctx.Set("Cache-Control", "no-store")
	ctx.Send(buf)
}
