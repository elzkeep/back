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
func (c *BillingController) Make(durationtype int, base int, year int, month int, durationmonth []int, ids []int64, price []int, vat []int, remark []string) {
	session := c.Session

	if session == nil {
		c.Error("invalid request")
		return
	}

	if len(ids) == 0 || len(price) != len(ids) || len(vat) != len(ids) || len(remark) != len(ids) {
		c.Error("invalid request")
		return
	}

	now := time.Now()

	months := make([]Pair, 0)
	if durationtype == 2 {
		if len(durationmonth) == 0 || year <= 0 {
			c.Error("invalid period")
			return
		}

		// 월이 이전 월보다 작아지면 다음 해로 넘어간 것으로 처리 (예: 11,12,1,2 → 11월~다음해 2월)
		// 연속 여부는 년*12+월 기준으로 판단하므로 12월→1월도 연속으로 묶임
		seen := make(map[int]bool)
		prev := 0
		for i, v := range durationmonth {
			if v < 1 || v > 12 || seen[v] {
				c.Error("invalid period")
				return
			}
			seen[v] = true

			current := year*12 + v - 1
			for i > 0 && current < prev {
				current += 12
			}

			if i > 0 && current == prev+1 {
				months[len(months)-1].Period++
			} else {
				months = append(months, Pair{Year: current / 12, Month: current%12 + 1, Period: 1})
			}

			prev = current
		}
	} else {
		if month < 1 {
			c.Error("invalid period")
			return
		}

		year = now.Year()
		currentMonth := int(now.Month())
		targetMonth := 1

		if base == 1 {
			targetMonth = currentMonth
		} else if base == 2 {
			if currentMonth == 12 {
				year++
				targetMonth = 1
			} else {
				targetMonth = currentMonth + 1
			}
		} else {
			if currentMonth == 1 {
				year--
				targetMonth = 12
			} else {
				targetMonth = currentMonth - 1
			}
		}

		months = append(months, Pair{Year: year, Month: targetMonth, Period: month})
	}

	conn := c.NewConnection()

	// 진행(1) 상태인 고객만 청구
	customerManager := models.NewCustomerManager(conn)
	customers := customerManager.Find([]interface{}{
		models.Where{Column: "company", Value: session.Company, Compare: "="},
		models.Where{Column: "building", Value: ids, Compare: "in"},
		models.Where{Column: "status", Value: 1, Compare: "="},
	})

	tx, err := conn.Begin()
	if err != nil {
		log.Println(err)
		c.Error("database error")
		return
	}
	defer tx.Rollback()

	billingManager := models.NewBillingManager(tx)

	today := global.GetCurrentDatetime()

	log.Println("months", months)
	log.Println(ids)
	log.Println(price)
	log.Println(vat)
	log.Println(remark)

	for _, d := range months {
		yearmonth := fmt.Sprintf("%04d-%02d", d.Year, d.Month)
		t := time.Date(d.Year, time.Month(d.Month), 1, 0, 0, 0, 0, time.Local)
		log.Println(t)
		ed := t.AddDate(0, d.Period-1, 0)
		log.Println(ed.Year(), ed.Month())
		endmonth := fmt.Sprintf("%04d-%02d", ed.Year(), ed.Month())
		for _, v := range customers {
			priceValue := 0
			vatValue := 0
			remarkValue := ""

			for i, id := range ids {
				if id != v.Building {
					continue
				}

				priceValue = price[i]
				vatValue = vat[i]
				remarkValue = remark[i]
				break
			}

			title := ""
			if d.Period == 1 {
				title = fmt.Sprintf("%v년 %v월분", d.Year, d.Month)
			} else {
				if d.Year == ed.Year() {
					title = fmt.Sprintf("%v년 %v월~%v월분", d.Year, d.Month, int(ed.Month()))
				} else {
					title = fmt.Sprintf("%v년 %v월 ~ %v년 %v월분", d.Year, d.Month, ed.Year(), int(ed.Month()))
				}
			}

			item := models.Billing{
				Title:        title,
				Price:        priceValue * d.Period,
				Vat:          vatValue * d.Period,
				Remark:       remarkValue,
				Status:       billing.StatusWait,
				Giro:         billing.GiroWait,
				Billdate:     today,
				Month:        yearmonth,
				Endmonth:     endmonth,
				Company:      session.Company,
				Building:     v.Building,
				Period:       d.Period,
				Billingtype:  v.Billingtype,
				Depositprice: 0,
			}

			if err := billingManager.Insert(&item); err != nil {
				log.Println(err)
				c.Error("billing insert error")
				return
			}
		}
	}

	if err := tx.Commit(); err != nil {
		log.Println(err)
		c.Error("database error")
	}
}
