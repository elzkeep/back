package rest


import (
	"zkeep/controllers"
	"zkeep/models"

    "strings"
)

type UserlistController struct {
	controllers.Controller
}

func (c *UserlistController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewUserlistManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *UserlistController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewUserlistManager(conn)

    var args []interface{}
    
    _loginid := c.Get("loginid")
    if _loginid != "" {
        args = append(args, models.Where{Column:"loginid", Value:_loginid, Compare:"like"})
    }
    _passwd := c.Get("passwd")
    if _passwd != "" {
        args = append(args, models.Where{Column:"passwd", Value:_passwd, Compare:"like"})
    }
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"="})
        
    }
    _email := c.Get("email")
    if _email != "" {
        args = append(args, models.Where{Column:"email", Value:_email, Compare:"like"})
    }
    _tel := c.Get("tel")
    if _tel != "" {
        args = append(args, models.Where{Column:"tel", Value:_tel, Compare:"like"})
    }
    _address := c.Get("address")
    if _address != "" {
        args = append(args, models.Where{Column:"address", Value:_address, Compare:"like"})
    }
    _addressetc := c.Get("addressetc")
    if _addressetc != "" {
        args = append(args, models.Where{Column:"addressetc", Value:_addressetc, Compare:"like"})
    }
    _startjoindate := c.Get("startjoindate")
    _endjoindate := c.Get("endjoindate")
    if _startjoindate != "" && _endjoindate != "" {        
        var v [2]string
        v[0] = _startjoindate
        v[1] = _endjoindate  
        args = append(args, models.Where{Column:"joindate", Value:v, Compare:"between"})    
    } else if  _startjoindate != "" {          
        args = append(args, models.Where{Column:"joindate", Value:_startjoindate, Compare:">="})
    } else if  _endjoindate != "" {          
        args = append(args, models.Where{Column:"joindate", Value:_endjoindate, Compare:"<="})            
    }
    _careeryear := c.Geti("careeryear")
    if _careeryear != 0 {
        args = append(args, models.Where{Column:"careeryear", Value:_careeryear, Compare:"="})    
    }
    _careermonth := c.Geti("careermonth")
    if _careermonth != 0 {
        args = append(args, models.Where{Column:"careermonth", Value:_careermonth, Compare:"="})    
    }
    _level := c.Geti("level")
    if _level != 0 {
        args = append(args, models.Where{Column:"level", Value:_level, Compare:"="})    
    }
    _score := c.Geti("score")
    if _score != 0 {
        args = append(args, models.Where{Column:"score", Value:_score, Compare:"="})    
    }
    _status := c.Geti("status")
    if _status != 0 {
        args = append(args, models.Where{Column:"status", Value:_status, Compare:"="})    
    }
    _company := c.Geti64("company")
    if _company != 0 {
        args = append(args, models.Where{Column:"company", Value:_company, Compare:"="})    
    }
    _department := c.Geti64("department")
    if _department != 0 {
        args = append(args, models.Where{Column:"department", Value:_department, Compare:"="})    
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
    _totalscore := c.Geti("totalscore")
    if _totalscore != 0 {
        args = append(args, models.Where{Column:"totalscore", Value:_totalscore, Compare:"="})    
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
                    str += ", u_" + strings.Trim(v, " ")                
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










