package rest


import (
	"aoi/controllers"
	"aoi/models"

    "strings"
)

type TokenController struct {
	controllers.Controller
}

func (c *TokenController) Read(id int64) {
    
    if c.Session == nil {
        c.Result["code"] = "auth error"
        return
    }
    if c.Session.Level < 3 { 
    
        if c.Session.Id == 0 {
            c.Result["code"] = "auth error"
            return
        }    
    
    }
    
    
	conn := c.NewConnection()

	manager := models.NewTokenManager(conn)
	item := manager.Get(id)

    
    if c.Session.Level < 3 {
    
        if item.User != c.Session.Id {
            c.Result["code"] = "auth error"
            return
        }
    
    }
    
    
    c.Set("item", item)
}

func (c *TokenController) Index(page int, pagesize int) {
    
    if c.Session == nil {
        c.Result["code"] = "auth error"
        return
    }

    if c.Session.Level < 3 {
    
        if c.Session.Id == 0 {
            c.Result["code"] = "auth error"
            return
        }
    
    }
    
    
	conn := c.NewConnection()

	manager := models.NewTokenManager(conn)

    var args []interface{}
    
    _user := c.Geti64("user")
    if _user != 0 {
        args = append(args, models.Where{Column:"user", Value:_user, Compare:"="})    
    }
    _token := c.Get("token")
    if _token != "" {
        args = append(args, models.Where{Column:"token", Value:_token, Compare:"like"})
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
    

    
    if c.Session.Level < 3 {
    
        args = append(args, models.Where{Column:"user", Value:c.Session.Id, Compare:"="})    
    
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
                    str += ", a_" + strings.Trim(v, " ")                
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

func (c *TokenController) Insert(item *models.Token) {
    
    if c.Session == nil {
        item = nil
        return
    }

    if c.Session.Level < 3 {
    
        if c.Session.Id == 0 {
            item = nil
            return
        } else {
            item.User = c.Session.Id
        }    
    
    }
    
    
	conn := c.NewConnection()
    
	manager := models.NewTokenManager(conn)
	manager.Insert(item)

    id := manager.GetIdentity()
    c.Result["id"] = id
    item.Id = id
}

func (c *TokenController) Insertbatch(item *[]models.Token) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    if c.Session == nil {
        item = nil
        return
    }

    if c.Session.Level < 3 {
    
        if c.Session.Id == 0 {
            item = nil
            return
        } else {
            for i := 0; i < rows; i++ {
                (*item)[i].User = c.Session.Id
            }
        }    
    
    }
    
    
	conn := c.NewConnection()
    
	manager := models.NewTokenManager(conn)

    for i := 0; i < rows; i++ {
	    manager.Insert(&((*item)[i]))
    }
}

func (c *TokenController) Update(item *models.Token) {
    
    if c.Session == nil {
        item = nil
        return
    }
    
    if c.Session.Level < 3 {
    
        if c.Session.Id == 0 {
            item = nil
            return
        } else {
            item.User = c.Session.Id
        }    
    
    }
    
    
	conn := c.NewConnection()

	manager := models.NewTokenManager(conn)
	manager.Update(item)
}

func (c *TokenController) Delete(item *models.Token) {
    
    if c.Session == nil {
        item = nil
        return
    }
    
    
    conn := c.NewConnection()

	manager := models.NewTokenManager(conn)

    
    n := manager.Get(item.Id)

    if c.Session.Level < 3 {
    
        if c.Session.Id == 0 {
            item = nil
            return
        } else {
            if n.User != c.Session.Id {
                item = nil
                return
            }
        }    
    
    }
    
	manager.Delete(item.Id)
}

func (c *TokenController) Deletebatch(item *[]models.Token) {
    
    if c.Session == nil {
        item = nil
        return
    }
    
    
    conn := c.NewConnection()

	manager := models.NewTokenManager(conn)

    for _, v := range *item {
        
        n := manager.Get(v.Id)

        if c.Session.Level < 3 {
        
            if c.Session.Id == 0 {
                item = nil
                return
            } else {
                if n.User != c.Session.Id {
                    item = nil
                    return
                }
            }    
        
        }
        
    
	    manager.Delete(v.Id)
    }
}



func (c *TokenController) GetByUser(user int64) *models.Token {
    
    if c.Session == nil {
        c.Result["code"] = "auth error"
        
        return nil
        
    }

    if c.Session.Level < 3 {
    
    if c.Session.Id == 0 {
        c.Result["code"] = "auth error"
        
        return nil
        
    } else {
        
        
        user = c.Session.Id
        
        
    }
    
    }
    
    conn := c.NewConnection()

	_manager := models.NewTokenManager(conn)
    
    item := _manager.GetByUser(user)
    
    c.Set("item", item)
    
    
    
    return item
    
}


// @Put()
func (c *TokenController) UpdateUser(user int64, id int64) {
    
    if c.Session == nil {
        return
    }    
    
    
	conn := c.NewConnection()

	_manager := models.NewTokenManager(conn)
	_manager.UpdateUser(user, id)
}

// @Put()
func (c *TokenController) UpdateToken(token string, id int64) {
    
    if c.Session == nil {
        return
    }    
    
    
	conn := c.NewConnection()

	_manager := models.NewTokenManager(conn)
	_manager.UpdateToken(token, id)
}

// @Put()
func (c *TokenController) UpdateStatus(status int, id int64) {
    
    if c.Session == nil {
        return
    }    
    
    
	conn := c.NewConnection()

	_manager := models.NewTokenManager(conn)
	_manager.UpdateStatus(status, id)
}






