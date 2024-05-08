package api

import (
	"fmt"
	"log"
	"strings"
	"time"
	"zkeep/controllers"
	"zkeep/global"
	"zkeep/models"
	"zkeep/models/billing"
)

type BillingController struct {
	controllers.Controller
}

func (c *BillingController) Search(page int, pagesize int) {

	conn := c.NewConnection()

	manager := models.NewBillingManager(conn)

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

type Pair struct {
	Year   int
	Month  int
	Period int
}

// @POST()
func (c *BillingController) Make(durationtype int, base int, year int, month int, durationmonth []int, ids []int64) {
	session := c.Session

	conn := c.NewConnection()

	customerManager := models.NewCustomerManager(conn)
	billingManager := models.NewBillingManager(conn)

	customers := customerManager.Find([]interface{}{
		models.Where{Column: "company", Value: session.Company, Compare: "="},
		models.Where{Column: "building", Value: ids, Compare: "in"},
	})

	log.Println("durationtype:", durationtype, durationmonth)
	now := time.Now()

	months := make([]Pair, 0)
	if durationtype == 2 {
		if len(durationmonth) == 0 {
			return
		}

		months = append(months, Pair{Year: year, Month: durationmonth[0], Period: 1})

		inc := 1
		for _, v := range durationmonth[1:] {
			current := months[len(months)-1]
			if v == current.Month+inc {
				current.Month = v
				months[len(months)-1].Period++
				inc++
			} else {
				months = append(months, Pair{Year: year, Month: v, Period: 1})
				inc = 1
			}
		}
	} else {
		year = now.Year()
		currentMonth := int(now.Month())
		targetMonth := 1

		if base == 1 {
			targetMonth = currentMonth
		} else {
			if currentMonth == 12 {
				year++
				targetMonth = 1
			} else {
				targetMonth = currentMonth + 1
			}
		}

		months = append(months, Pair{Year: year, Month: targetMonth, Period: month})
	}

	today := global.GetCurrentDatetime()

	for _, d := range months {
		yearmonth := fmt.Sprintf("%04d-%02d", d.Year, d.Month)

		for _, v := range customers {

			cnt := billingManager.Count([]interface{}{
				models.Where{Column: "company", Value: session.Company, Compare: "="},
				models.Where{Column: "building", Value: v.Building, Compare: "="},
				models.Where{Column: "month", Value: yearmonth, Compare: "="},
				models.Where{Column: "period", Value: d.Period, Compare: "="},
			})

			if cnt > 0 {
				continue
			}

			item := models.Billing{}
			item.Price = (v.Contractprice + v.Contractvat) * d.Period
			item.Status = billing.StatusWait
			item.Giro = billing.GiroWait
			item.Billdate = today
			item.Month = yearmonth
			item.Company = session.Company
			item.Building = v.Building
			item.Period = d.Period

			billingManager.Insert(&item)
		}
	}
}

// @Put()
func (c *BillingController) Process(item *models.Billing) {
	user := c.Session

	conn := c.NewConnection()

	manager := models.NewBillingManager(conn)
	billinghistoryManager := models.NewBillinghistoryManager(conn)

	manager.Update(item)

	billinghistoryManager.DeleteByBilling(item.Id)

	var billinghistory models.Billinghistory
	billinghistory.Price = item.Price
	billinghistory.Company = user.Company
	billinghistory.Building = item.Building
	billinghistory.Billing = item.Id

	if item.Status == billing.StatusComplete {
		billinghistoryManager.Insert(&billinghistory)
	}
}
