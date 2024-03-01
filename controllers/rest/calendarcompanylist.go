package rest


import (
	"zkeep/controllers"
	"zkeep/models"

    "strings"
)

type CalendarcompanylistController struct {
	controllers.Controller
}

func (c *CalendarcompanylistController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewCalendarcompanylistManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *CalendarcompanylistController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewCalendarcompanylistManager(conn)

    var args []interface{}
    
    _company := c.Geti64("company")
    if _company != 0 {
        args = append(args, models.Where{Column:"company", Value:_company, Compare:"="})    
    }
    _month := c.Get("month")
    if _month != "" {
        args = append(args, models.Where{Column:"month", Value:_month, Compare:"like"})
    }
    _day := c.Get("day")
    if _day != "" {
        args = append(args, models.Where{Column:"day", Value:_day, Compare:"like"})
    }
    _status := c.Geti("status")
    if _status != 0 {
        args = append(args, models.Where{Column:"status", Value:_status, Compare:"="})    
    }
    _count := c.Geti64("count")
    if _count != 0 {
        args = append(args, models.Where{Column:"count", Value:_count, Compare:"="})    
    }
    _startcheckdate := c.Get("startcheckdate")
    _endcheckdate := c.Get("endcheckdate")
    if _startcheckdate != "" && _endcheckdate != "" {        
        var v [2]string
        v[0] = _startcheckdate
        v[1] = _endcheckdate  
        args = append(args, models.Where{Column:"checkdate", Value:v, Compare:"between"})    
    } else if  _startcheckdate != "" {          
        args = append(args, models.Where{Column:"checkdate", Value:_startcheckdate, Compare:">="})
    } else if  _endcheckdate != "" {          
        args = append(args, models.Where{Column:"checkdate", Value:_endcheckdate, Compare:"<="})            
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










func (c *CalendarcompanylistController) Sum() {
    
    
	conn := c.NewConnection()

	manager := models.NewCalendarcompanylistManager(conn)

    var args []interface{}
    
    _company := c.Geti64("company")
    if _company != 0 {
        args = append(args, models.Where{Column:"company", Value:_company, Compare:"="})    
    }
    _month := c.Get("month")
    if _month != "" {
        args = append(args, models.Where{Column:"month", Value:_month, Compare:"like"})
    }
    _day := c.Get("day")
    if _day != "" {
        args = append(args, models.Where{Column:"day", Value:_day, Compare:"like"})
    }
    _status := c.Geti("status")
    if _status != 0 {
        args = append(args, models.Where{Column:"status", Value:_status, Compare:"="})    
    }
    _count := c.Geti64("count")
    if _count != 0 {
        args = append(args, models.Where{Column:"count", Value:_count, Compare:"="})    
    }
    _startcheckdate := c.Get("startcheckdate")
    _endcheckdate := c.Get("endcheckdate")
    if _startcheckdate != "" && _endcheckdate != "" {        
        var v [2]string
        v[0] = _startcheckdate
        v[1] = _endcheckdate  
        args = append(args, models.Where{Column:"checkdate", Value:v, Compare:"between"})    
    } else if  _startcheckdate != "" {          
        args = append(args, models.Where{Column:"checkdate", Value:_startcheckdate, Compare:">="})
    } else if  _endcheckdate != "" {          
        args = append(args, models.Where{Column:"checkdate", Value:_endcheckdate, Compare:"<="})            
    }
    

    
    
    item := manager.Sum(args)
	c.Set("item", item)
}

