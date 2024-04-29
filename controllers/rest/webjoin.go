package rest


import (
	"zkeep/controllers"
	"zkeep/models"

    "strings"
)

type WebjoinController struct {
	controllers.Controller
}



func (c *WebjoinController) Insert(item *models.Webjoin) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewWebjoinManager(conn)
	manager.Insert(item)

    id := manager.GetIdentity()
    c.Result["id"] = id
    item.Id = id
}

func (c *WebjoinController) Insertbatch(item *[]models.Webjoin) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewWebjoinManager(conn)

    for i := 0; i < rows; i++ {
	    manager.Insert(&((*item)[i]))
    }
}

func (c *WebjoinController) Update(item *models.Webjoin) {
    
    
	conn := c.NewConnection()

	manager := models.NewWebjoinManager(conn)
	manager.Update(item)
}

func (c *WebjoinController) Delete(item *models.Webjoin) {
    
    
    conn := c.NewConnection()

	manager := models.NewWebjoinManager(conn)

    
	manager.Delete(item.Id)
}

func (c *WebjoinController) Deletebatch(item *[]models.Webjoin) {
    
    
    conn := c.NewConnection()

	manager := models.NewWebjoinManager(conn)

    for _, v := range *item {
        
    
	    manager.Delete(v.Id)
    }
}

func (c *WebjoinController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewWebjoinManager(conn)

    var args []interface{}
    
    _category := c.Geti("category")
    if _category != 0 {
        args = append(args, models.Where{Column:"category", Value:_category, Compare:"="})    
    }
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"like"})
        
    }
    _manager := c.Get("manager")
    if _manager != "" {
        args = append(args, models.Where{Column:"manager", Value:_manager, Compare:"like"})
    }
    _tel := c.Get("tel")
    if _tel != "" {
        args = append(args, models.Where{Column:"tel", Value:_tel, Compare:"like"})
    }
    _email := c.Get("email")
    if _email != "" {
        args = append(args, models.Where{Column:"email", Value:_email, Compare:"like"})
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
    

    
    
    total := manager.Count(args)
	c.Set("total", total)
}


func (c *WebjoinController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewWebjoinManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *WebjoinController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewWebjoinManager(conn)

    var args []interface{}
    
    _category := c.Geti("category")
    if _category != 0 {
        args = append(args, models.Where{Column:"category", Value:_category, Compare:"="})    
    }
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"like"})
        
    }
    _manager := c.Get("manager")
    if _manager != "" {
        args = append(args, models.Where{Column:"manager", Value:_manager, Compare:"like"})
    }
    _tel := c.Get("tel")
    if _tel != "" {
        args = append(args, models.Where{Column:"tel", Value:_tel, Compare:"like"})
    }
    _email := c.Get("email")
    if _email != "" {
        args = append(args, models.Where{Column:"email", Value:_email, Compare:"like"})
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
                    str += ", wj_" + strings.Trim(v, " ")                
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

// @Put()
func (c *WebjoinController) UpdateCategory(category int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewWebjoinManager(conn)
	_manager.UpdateCategory(category, id)
}

// @Put()
func (c *WebjoinController) UpdateName(name string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewWebjoinManager(conn)
	_manager.UpdateName(name, id)
}

// @Put()
func (c *WebjoinController) UpdateManager(manager string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewWebjoinManager(conn)
	_manager.UpdateManager(manager, id)
}

// @Put()
func (c *WebjoinController) UpdateTel(tel string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewWebjoinManager(conn)
	_manager.UpdateTel(tel, id)
}

// @Put()
func (c *WebjoinController) UpdateEmail(email string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewWebjoinManager(conn)
	_manager.UpdateEmail(email, id)
}






