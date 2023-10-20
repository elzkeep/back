package rest


import (
	"aoi/controllers"
	"aoi/models"

    "strings"
)

type GameController struct {
	controllers.Controller
}

func (c *GameController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewGameManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *GameController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewGameManager(conn)

    var args []interface{}
    
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"="})
        
    }
    _count := c.Geti("count")
    if _count != 0 {
        args = append(args, models.Where{Column:"count", Value:_count, Compare:"="})    
    }
    _map := c.Geti64("map")
    if _map != 0 {
        args = append(args, models.Where{Column:"map", Value:_map, Compare:"="})    
    }
    _type := c.Geti("type")
    if _type != 0 {
        args = append(args, models.Where{Column:"type", Value:_type, Compare:"="})    
    }
    _status := c.Geti("status")
    if _status != 0 {
        args = append(args, models.Where{Column:"status", Value:_status, Compare:"="})    
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

func (c *GameController) Insert(item *models.Game) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewGameManager(conn)
	manager.Insert(item)

    id := manager.GetIdentity()
    c.Result["id"] = id
    item.Id = id
}

func (c *GameController) Insertbatch(item *[]models.Game) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewGameManager(conn)

    for i := 0; i < rows; i++ {
	    manager.Insert(&((*item)[i]))
    }
}

func (c *GameController) Update(item *models.Game) {
    
    
	conn := c.NewConnection()

	manager := models.NewGameManager(conn)
	manager.Update(item)
}

func (c *GameController) Delete(item *models.Game) {
    
    
    conn := c.NewConnection()

	manager := models.NewGameManager(conn)

    
	manager.Delete(item.Id)
}

func (c *GameController) Deletebatch(item *[]models.Game) {
    
    
    conn := c.NewConnection()

	manager := models.NewGameManager(conn)

    for _, v := range *item {
        
    
	    manager.Delete(v.Id)
    }
}



// @Put()
func (c *GameController) UpdateName(name string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewGameManager(conn)
	_manager.UpdateName(name, id)
}

// @Put()
func (c *GameController) UpdateCount(count int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewGameManager(conn)
	_manager.UpdateCount(count, id)
}

// @Put()
func (c *GameController) UpdateMap(mapid int64, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewGameManager(conn)
	_manager.UpdateMap(mapid, id)
}

// @Put()
func (c *GameController) UpdateType(typeid int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewGameManager(conn)
	_manager.UpdateType(typeid, id)
}

// @Put()
func (c *GameController) UpdateStatus(status int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewGameManager(conn)
	_manager.UpdateStatus(status, id)
}






func (c *GameController) Sum() {
    
    
	conn := c.NewConnection()

	manager := models.NewGameManager(conn)

    var args []interface{}
    
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"="})
        
    }
    _count := c.Geti("count")
    if _count != 0 {
        args = append(args, models.Where{Column:"count", Value:_count, Compare:"="})    
    }
    _map := c.Geti64("map")
    if _map != 0 {
        args = append(args, models.Where{Column:"map", Value:_map, Compare:"="})    
    }
    _type := c.Geti("type")
    if _type != 0 {
        args = append(args, models.Where{Column:"type", Value:_type, Compare:"="})    
    }
    _status := c.Geti("status")
    if _status != 0 {
        args = append(args, models.Where{Column:"status", Value:_status, Compare:"="})    
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

