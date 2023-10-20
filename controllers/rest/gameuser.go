package rest


import (
	"aoi/controllers"
	"aoi/models"

    "strings"
)

type GameuserController struct {
	controllers.Controller
}

func (c *GameuserController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewGameuserManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *GameuserController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewGameuserManager(conn)

    var args []interface{}
    
    _order := c.Geti("order")
    if _order != 0 {
        args = append(args, models.Where{Column:"order", Value:_order, Compare:"="})    
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
                    str += ", gu_" + strings.Trim(v, " ")                
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

func (c *GameuserController) Insert(item *models.Gameuser) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewGameuserManager(conn)
	manager.Insert(item)

    id := manager.GetIdentity()
    c.Result["id"] = id
    item.Id = id
}

func (c *GameuserController) Insertbatch(item *[]models.Gameuser) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewGameuserManager(conn)

    for i := 0; i < rows; i++ {
	    manager.Insert(&((*item)[i]))
    }
}

func (c *GameuserController) Update(item *models.Gameuser) {
    
    
	conn := c.NewConnection()

	manager := models.NewGameuserManager(conn)
	manager.Update(item)
}

func (c *GameuserController) Delete(item *models.Gameuser) {
    
    
    conn := c.NewConnection()

	manager := models.NewGameuserManager(conn)

    
	manager.Delete(item.Id)
}

func (c *GameuserController) Deletebatch(item *[]models.Gameuser) {
    
    
    conn := c.NewConnection()

	manager := models.NewGameuserManager(conn)

    for _, v := range *item {
        
    
	    manager.Delete(v.Id)
    }
}



func (c *GameuserController) CountByGame(game int64) int {
    
    conn := c.NewConnection()

	_manager := models.NewGameuserManager(conn)
    
    item := _manager.CountByGame(game)
    
    
    
    c.Set("count", item)
    
    return item
    
}


func (c *GameuserController) CountByGameUser(game int64 ,user int64) int {
    
    conn := c.NewConnection()

	_manager := models.NewGameuserManager(conn)
    
    item := _manager.CountByGameUser(game, user)
    
    
    
    c.Set("count", item)
    
    return item
    
}


// @Put()
func (c *GameuserController) UpdateOrder(order int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewGameuserManager(conn)
	_manager.UpdateOrder(order, id)
}

// @Put()
func (c *GameuserController) UpdateUser(user int64, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewGameuserManager(conn)
	_manager.UpdateUser(user, id)
}

// @Put()
func (c *GameuserController) UpdateGame(game int64, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewGameuserManager(conn)
	_manager.UpdateGame(game, id)
}






