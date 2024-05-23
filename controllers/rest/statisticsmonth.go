package rest


import (
	"zkeep/controllers"
	"zkeep/models"

    "strings"
)

type StatisticsmonthController struct {
	controllers.Controller
}





func (c *StatisticsmonthController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewStatisticsmonthManager(conn)

    var args []interface{}
    
    _year := c.Get("year")
    if _year != "" {
        args = append(args, models.Where{Column:"year", Value:_year, Compare:"="})
    }
    _duration := c.Get("duration")
    if _duration != "" {
        args = append(args, models.Where{Column:"duration", Value:_duration, Compare:"="})
    }
    _total := c.Geti64("total")
    if _total != 0 {
        args = append(args, models.Where{Column:"total", Value:_total, Compare:"="})    
    }
    _totalprice := c.Geti64("totalprice")
    if _totalprice != 0 {
        args = append(args, models.Where{Column:"totalprice", Value:_totalprice, Compare:"="})    
    }
    _startdate := c.Get("startdate")
    _enddate := c.Get("enddate")
    if _startdate != "" && _enddate != "" {        
        var v [2]string
        v[0] = _startdate
        v[1] = _enddate  
        args = append(args, models.Where{Column:"date", Value:v, Compare:"between"})    
    } else if  _startdate != "" {          
        args = append(args, models.Where{Column:"date", Value:_startdate, Compare:">="})
    } else if  _enddate != "" {          
        args = append(args, models.Where{Column:"date", Value:_enddate, Compare:"<="})            
    }
    

    
    
    total := manager.Count(args)
	c.Set("total", total)
}


func (c *StatisticsmonthController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewStatisticsmonthManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *StatisticsmonthController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewStatisticsmonthManager(conn)

    var args []interface{}
    
    _year := c.Get("year")
    if _year != "" {
        args = append(args, models.Where{Column:"year", Value:_year, Compare:"="})
    }
    _duration := c.Get("duration")
    if _duration != "" {
        args = append(args, models.Where{Column:"duration", Value:_duration, Compare:"="})
    }
    _total := c.Geti64("total")
    if _total != 0 {
        args = append(args, models.Where{Column:"total", Value:_total, Compare:"="})    
    }
    _totalprice := c.Geti64("totalprice")
    if _totalprice != 0 {
        args = append(args, models.Where{Column:"totalprice", Value:_totalprice, Compare:"="})    
    }
    _startdate := c.Get("startdate")
    _enddate := c.Get("enddate")
    if _startdate != "" && _enddate != "" {        
        var v [2]string
        v[0] = _startdate
        v[1] = _enddate  
        args = append(args, models.Where{Column:"date", Value:v, Compare:"between"})    
    } else if  _startdate != "" {          
        args = append(args, models.Where{Column:"date", Value:_startdate, Compare:">="})
    } else if  _enddate != "" {          
        args = append(args, models.Where{Column:"date", Value:_enddate, Compare:"<="})            
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






func (c *StatisticsmonthController) Sum() {
    
    
	conn := c.NewConnection()

	manager := models.NewStatisticsmonthManager(conn)

    var args []interface{}
    
    _year := c.Get("year")
    if _year != "" {
        args = append(args, models.Where{Column:"year", Value:_year, Compare:"like"})
    }
    _duration := c.Get("duration")
    if _duration != "" {
        args = append(args, models.Where{Column:"duration", Value:_duration, Compare:"like"})
    }
    _total := c.Geti64("total")
    if _total != 0 {
        args = append(args, models.Where{Column:"total", Value:_total, Compare:"="})    
    }
    _totalprice := c.Geti64("totalprice")
    if _totalprice != 0 {
        args = append(args, models.Where{Column:"totalprice", Value:_totalprice, Compare:"="})    
    }
    _startdate := c.Get("startdate")
    _enddate := c.Get("enddate")
    if _startdate != "" && _enddate != "" {        
        var v [2]string
        v[0] = _startdate
        v[1] = _enddate  
        args = append(args, models.Where{Column:"date", Value:v, Compare:"between"})    
    } else if  _startdate != "" {          
        args = append(args, models.Where{Column:"date", Value:_startdate, Compare:">="})
    } else if  _enddate != "" {          
        args = append(args, models.Where{Column:"date", Value:_enddate, Compare:"<="})            
    }
    

    
    
    item := manager.Sum(args)
	c.Set("item", item)
}

