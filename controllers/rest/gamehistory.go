package rest


import (
	"aoi/controllers"
	"aoi/models"

    "strings"
)

type GamehistoryController struct {
	controllers.Controller
}

func (c *GamehistoryController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewGamehistoryManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *GamehistoryController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewGamehistoryManager(conn)

    var args []interface{}
    
    _round := c.Geti("round")
    if _round != 0 {
        args = append(args, models.Where{Column:"round", Value:_round, Compare:"="})    
    }
    _command := c.Get("command")
    if _command != "" {
        args = append(args, models.Where{Column:"command", Value:_command, Compare:"like"})
    }
    _vp := c.Geti("vp")
    if _vp != 0 {
        args = append(args, models.Where{Column:"vp", Value:_vp, Compare:"="})    
    }
    _user := c.Geti64("user")
    if _user != 0 {
        args = append(args, models.Where{Column:"user", Value:_user, Compare:"="})    
    }
    _game := c.Geti64("game")
    if _game != 0 {
        args = append(args, models.Where{Column:"game", Value:_game, Compare:"="})    
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
                    str += ", gh_" + strings.Trim(v, " ")                
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

func (c *GamehistoryController) Insert(item *models.Gamehistory) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewGamehistoryManager(conn)
	manager.Insert(item)

    id := manager.GetIdentity()
    c.Result["id"] = id
    item.Id = id
}

func (c *GamehistoryController) Insertbatch(item *[]models.Gamehistory) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewGamehistoryManager(conn)

    for i := 0; i < rows; i++ {
	    manager.Insert(&((*item)[i]))
    }
}

func (c *GamehistoryController) Update(item *models.Gamehistory) {
    
    
	conn := c.NewConnection()

	manager := models.NewGamehistoryManager(conn)
	manager.Update(item)
}

func (c *GamehistoryController) Delete(item *models.Gamehistory) {
    
    
    conn := c.NewConnection()

	manager := models.NewGamehistoryManager(conn)

    
	manager.Delete(item.Id)
}

func (c *GamehistoryController) Deletebatch(item *[]models.Gamehistory) {
    
    
    conn := c.NewConnection()

	manager := models.NewGamehistoryManager(conn)

    for _, v := range *item {
        
    
	    manager.Delete(v.Id)
    }
}



// @Put()
func (c *GamehistoryController) UpdateRound(round int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewGamehistoryManager(conn)
	_manager.UpdateRound(round, id)
}

// @Put()
func (c *GamehistoryController) UpdateCommand(command string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewGamehistoryManager(conn)
	_manager.UpdateCommand(command, id)
}

// @Put()
func (c *GamehistoryController) UpdateVp(vp int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewGamehistoryManager(conn)
	_manager.UpdateVp(vp, id)
}

// @Put()
func (c *GamehistoryController) UpdateUser(user int64, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewGamehistoryManager(conn)
	_manager.UpdateUser(user, id)
}

// @Put()
func (c *GamehistoryController) UpdateGame(game int64, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewGamehistoryManager(conn)
	_manager.UpdateGame(game, id)
}






