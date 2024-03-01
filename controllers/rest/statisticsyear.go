package rest


import (
	"zkeep/controllers"
	"zkeep/models"

    "strings"
)

type StatisticsyearController struct {
	controllers.Controller
}

func (c *StatisticsyearController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewStatisticsyearManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *StatisticsyearController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewStatisticsyearManager(conn)

    var args []interface{}
    
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
    _startbilldate := c.Get("startbilldate")
    _endbilldate := c.Get("endbilldate")
    if _startbilldate != "" && _endbilldate != "" {        
        var v [2]string
        v[0] = _startbilldate
        v[1] = _endbilldate  
        args = append(args, models.Where{Column:"billdate", Value:v, Compare:"between"})    
    } else if  _startbilldate != "" {          
        args = append(args, models.Where{Column:"billdate", Value:_startbilldate, Compare:">="})
    } else if  _endbilldate != "" {          
        args = append(args, models.Where{Column:"billdate", Value:_endbilldate, Compare:"<="})            
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










func (c *StatisticsyearController) Sum() {
    
    
	conn := c.NewConnection()

	manager := models.NewStatisticsyearManager(conn)

    var args []interface{}
    
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
    _startbilldate := c.Get("startbilldate")
    _endbilldate := c.Get("endbilldate")
    if _startbilldate != "" && _endbilldate != "" {        
        var v [2]string
        v[0] = _startbilldate
        v[1] = _endbilldate  
        args = append(args, models.Where{Column:"billdate", Value:v, Compare:"between"})    
    } else if  _startbilldate != "" {          
        args = append(args, models.Where{Column:"billdate", Value:_startbilldate, Compare:">="})
    } else if  _endbilldate != "" {          
        args = append(args, models.Where{Column:"billdate", Value:_endbilldate, Compare:"<="})            
    }
    

    
    
    item := manager.Sum(args)
	c.Set("item", item)
}

