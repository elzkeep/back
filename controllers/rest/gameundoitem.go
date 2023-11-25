package rest


import (
	"aoi/controllers"
	"aoi/models"

    "strings"
)

type GameundoitemController struct {
	controllers.Controller
}

func (c *GameundoitemController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewGameundoitemManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *GameundoitemController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewGameundoitemManager(conn)

    var args []interface{}
    
    _status := c.Geti("status")
    if _status != 0 {
        args = append(args, models.Where{Column:"status", Value:_status, Compare:"="})    
    }
    _gameundo := c.Geti64("gameundo")
    if _gameundo != 0 {
        args = append(args, models.Where{Column:"gameundo", Value:_gameundo, Compare:"="})    
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
                    str += ", gi_" + strings.Trim(v, " ")                
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

func (c *GameundoitemController) Insert(item *models.Gameundoitem) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewGameundoitemManager(conn)
	manager.Insert(item)

    id := manager.GetIdentity()
    c.Result["id"] = id
    item.Id = id
}

func (c *GameundoitemController) Insertbatch(item *[]models.Gameundoitem) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewGameundoitemManager(conn)

    for i := 0; i < rows; i++ {
	    manager.Insert(&((*item)[i]))
    }
}

func (c *GameundoitemController) Update(item *models.Gameundoitem) {
    
    
	conn := c.NewConnection()

	manager := models.NewGameundoitemManager(conn)
	manager.Update(item)
}

func (c *GameundoitemController) Delete(item *models.Gameundoitem) {
    
    
    conn := c.NewConnection()

	manager := models.NewGameundoitemManager(conn)

    
	manager.Delete(item.Id)
}

func (c *GameundoitemController) Deletebatch(item *[]models.Gameundoitem) {
    
    
    conn := c.NewConnection()

	manager := models.NewGameundoitemManager(conn)

    for _, v := range *item {
        
    
	    manager.Delete(v.Id)
    }
}



// @Put()
func (c *GameundoitemController) UpdateStatus(status int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewGameundoitemManager(conn)
	_manager.UpdateStatus(status, id)
}

// @Put()
func (c *GameundoitemController) UpdateGameundo(gameundo int64, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewGameundoitemManager(conn)
	_manager.UpdateGameundo(gameundo, id)
}

// @Put()
func (c *GameundoitemController) UpdateGame(game int64, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewGameundoitemManager(conn)
	_manager.UpdateGame(game, id)
}

// @Put()
func (c *GameundoitemController) UpdateUser(user int64, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewGameundoitemManager(conn)
	_manager.UpdateUser(user, id)
}






