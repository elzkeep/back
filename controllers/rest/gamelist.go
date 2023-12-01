package rest


import (
	"aoi/controllers"
	"aoi/models"

    "strings"
)

type GamelistController struct {
	controllers.Controller
}

func (c *GamelistController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewGamelistManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *GamelistController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewGamelistManager(conn)

    var args []interface{}
    
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"="})
        
    }
    _count := c.Geti("count")
    if _count != 0 {
        args = append(args, models.Where{Column:"count", Value:_count, Compare:"="})    
    }
    _join := c.Geti("join")
    if _join != 0 {
        args = append(args, models.Where{Column:"join", Value:_join, Compare:"="})    
    }
    _map := c.Geti64("map")
    if _map != 0 {
        args = append(args, models.Where{Column:"map", Value:_map, Compare:"="})    
    }
    _illusionists := c.Geti("illusionists")
    if _illusionists != 0 {
        args = append(args, models.Where{Column:"illusionists", Value:_illusionists, Compare:"="})    
    }
    _type := c.Geti("type")
    if _type != 0 {
        args = append(args, models.Where{Column:"type", Value:_type, Compare:"="})    
    }
    _status := c.Geti("status")
    if _status != 0 {
        args = append(args, models.Where{Column:"status", Value:_status, Compare:"="})    
    }
    _startenddate := c.Get("startenddate")
    _endenddate := c.Get("endenddate")
    if _startenddate != "" && _endenddate != "" {        
        var v [2]string
        v[0] = _startenddate
        v[1] = _endenddate  
        args = append(args, models.Where{Column:"enddate", Value:v, Compare:"between"})    
    } else if  _startenddate != "" {          
        args = append(args, models.Where{Column:"enddate", Value:_startenddate, Compare:">="})
    } else if  _endenddate != "" {          
        args = append(args, models.Where{Column:"enddate", Value:_endenddate, Compare:"<="})            
    }
    _user := c.Geti64("user")
    if _user != 0 {
        args = append(args, models.Where{Column:"user", Value:_user, Compare:"="})    
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
    _gameuser := c.Geti64("gameuser")
    if _gameuser != 0 {
        args = append(args, models.Where{Column:"gameuser", Value:_gameuser, Compare:"="})    
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
                    str += ", g_" + strings.Trim(v, " ")                
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










func (c *GamelistController) Sum() {
    
    
	conn := c.NewConnection()

	manager := models.NewGamelistManager(conn)

    var args []interface{}
    
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"="})
        
    }
    _count := c.Geti("count")
    if _count != 0 {
        args = append(args, models.Where{Column:"count", Value:_count, Compare:"="})    
    }
    _join := c.Geti("join")
    if _join != 0 {
        args = append(args, models.Where{Column:"join", Value:_join, Compare:"="})    
    }
    _map := c.Geti64("map")
    if _map != 0 {
        args = append(args, models.Where{Column:"map", Value:_map, Compare:"="})    
    }
    _illusionists := c.Geti("illusionists")
    if _illusionists != 0 {
        args = append(args, models.Where{Column:"illusionists", Value:_illusionists, Compare:"="})    
    }
    _type := c.Geti("type")
    if _type != 0 {
        args = append(args, models.Where{Column:"type", Value:_type, Compare:"="})    
    }
    _status := c.Geti("status")
    if _status != 0 {
        args = append(args, models.Where{Column:"status", Value:_status, Compare:"="})    
    }
    _startenddate := c.Get("startenddate")
    _endenddate := c.Get("endenddate")
    if _startenddate != "" && _endenddate != "" {        
        var v [2]string
        v[0] = _startenddate
        v[1] = _endenddate  
        args = append(args, models.Where{Column:"enddate", Value:v, Compare:"between"})    
    } else if  _startenddate != "" {          
        args = append(args, models.Where{Column:"enddate", Value:_startenddate, Compare:">="})
    } else if  _endenddate != "" {          
        args = append(args, models.Where{Column:"enddate", Value:_endenddate, Compare:"<="})            
    }
    _user := c.Geti64("user")
    if _user != 0 {
        args = append(args, models.Where{Column:"user", Value:_user, Compare:"="})    
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
    _gameuser := c.Geti64("gameuser")
    if _gameuser != 0 {
        args = append(args, models.Where{Column:"gameuser", Value:_gameuser, Compare:"="})    
    }
    

    
    
    item := manager.Sum(args)
	c.Set("item", item)
}

