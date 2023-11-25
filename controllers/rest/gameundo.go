package rest


import (
	"aoi/controllers"
	"aoi/models"

    "strings"
)

type GameundoController struct {
	controllers.Controller
}

func (c *GameundoController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewGameundoManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *GameundoController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewGameundoManager(conn)

    var args []interface{}
    
    _status := c.Geti("status")
    if _status != 0 {
        args = append(args, models.Where{Column:"status", Value:_status, Compare:"="})    
    }
    _gamehistory := c.Geti64("gamehistory")
    if _gamehistory != 0 {
        args = append(args, models.Where{Column:"gamehistory", Value:_gamehistory, Compare:"="})    
    }
    _game := c.Geti64("game")
    if _game != 0 {
        args = append(args, models.Where{Column:"game", Value:_game, Compare:"="})    
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
                    str += ", gn_" + strings.Trim(v, " ")                
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

func (c *GameundoController) Insert(item *models.Gameundo) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewGameundoManager(conn)
	manager.Insert(item)

    id := manager.GetIdentity()
    c.Result["id"] = id
    item.Id = id
}

func (c *GameundoController) Insertbatch(item *[]models.Gameundo) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewGameundoManager(conn)

    for i := 0; i < rows; i++ {
	    manager.Insert(&((*item)[i]))
    }
}

func (c *GameundoController) Update(item *models.Gameundo) {
    
    
	conn := c.NewConnection()

	manager := models.NewGameundoManager(conn)
	manager.Update(item)
}

func (c *GameundoController) Delete(item *models.Gameundo) {
    
    
    conn := c.NewConnection()

	manager := models.NewGameundoManager(conn)

    
	manager.Delete(item.Id)
}

func (c *GameundoController) Deletebatch(item *[]models.Gameundo) {
    
    
    conn := c.NewConnection()

	manager := models.NewGameundoManager(conn)

    for _, v := range *item {
        
    
	    manager.Delete(v.Id)
    }
}



// @Put()
func (c *GameundoController) UpdateStatus(status int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewGameundoManager(conn)
	_manager.UpdateStatus(status, id)
}

// @Put()
func (c *GameundoController) UpdateGamehistory(gamehistory int64, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewGameundoManager(conn)
	_manager.UpdateGamehistory(gamehistory, id)
}

// @Put()
func (c *GameundoController) UpdateGame(game int64, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewGameundoManager(conn)
	_manager.UpdateGame(game, id)
}

// @Put()
func (c *GameundoController) UpdateUser(user int64, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewGameundoManager(conn)
	_manager.UpdateUser(user, id)
}






