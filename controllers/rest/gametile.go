package rest


import (
	"aoi/controllers"
	"aoi/models"

    "strings"
)

type GametileController struct {
	controllers.Controller
}

func (c *GametileController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewGametileManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *GametileController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewGametileManager(conn)

    var args []interface{}
    
    _type := c.Geti("type")
    if _type != 0 {
        args = append(args, models.Where{Column:"type", Value:_type, Compare:"="})    
    }
    _number := c.Geti("number")
    if _number != 0 {
        args = append(args, models.Where{Column:"number", Value:_number, Compare:"="})    
    }
    _order := c.Geti("order")
    if _order != 0 {
        args = append(args, models.Where{Column:"order", Value:_order, Compare:"="})    
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
                    str += ", gt_" + strings.Trim(v, " ")                
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

func (c *GametileController) Insert(item *models.Gametile) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewGametileManager(conn)
	manager.Insert(item)

    id := manager.GetIdentity()
    c.Result["id"] = id
    item.Id = id
}

func (c *GametileController) Insertbatch(item *[]models.Gametile) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewGametileManager(conn)

    for i := 0; i < rows; i++ {
	    manager.Insert(&((*item)[i]))
    }
}

func (c *GametileController) Update(item *models.Gametile) {
    
    
	conn := c.NewConnection()

	manager := models.NewGametileManager(conn)
	manager.Update(item)
}

func (c *GametileController) Delete(item *models.Gametile) {
    
    
    conn := c.NewConnection()

	manager := models.NewGametileManager(conn)

    
	manager.Delete(item.Id)
}

func (c *GametileController) Deletebatch(item *[]models.Gametile) {
    
    
    conn := c.NewConnection()

	manager := models.NewGametileManager(conn)

    for _, v := range *item {
        
    
	    manager.Delete(v.Id)
    }
}



func (c *GametileController) FindByGame(game int64) []models.Gametile {
    
    conn := c.NewConnection()

	_manager := models.NewGametileManager(conn)
    
    item := _manager.FindByGame(game)
    
    
    c.Set("items", item)
    
    
    return item
    
}


func (c *GametileController) CountByGame(game int64) int {
    
    conn := c.NewConnection()

	_manager := models.NewGametileManager(conn)
    
    item := _manager.CountByGame(game)
    
    
    
    c.Set("count", item)
    
    return item
    
}

// @Delete()
func (c *GametileController) DeleteByGame(game int64) {
    
    conn := c.NewConnection()

	_manager := models.NewGametileManager(conn)
    
    _manager.DeleteByGame(game)
    
}


// @Put()
func (c *GametileController) UpdateType(typeid int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewGametileManager(conn)
	_manager.UpdateType(typeid, id)
}

// @Put()
func (c *GametileController) UpdateNumber(number int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewGametileManager(conn)
	_manager.UpdateNumber(number, id)
}

// @Put()
func (c *GametileController) UpdateOrder(order int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewGametileManager(conn)
	_manager.UpdateOrder(order, id)
}

// @Put()
func (c *GametileController) UpdateGame(game int64, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewGametileManager(conn)
	_manager.UpdateGame(game, id)
}






