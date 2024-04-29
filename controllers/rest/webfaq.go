package rest


import (
	"zkeep/controllers"
	"zkeep/models"

    "strings"
)

type WebfaqController struct {
	controllers.Controller
}



func (c *WebfaqController) Insert(item *models.Webfaq) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewWebfaqManager(conn)
	manager.Insert(item)

    id := manager.GetIdentity()
    c.Result["id"] = id
    item.Id = id
}

func (c *WebfaqController) Insertbatch(item *[]models.Webfaq) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewWebfaqManager(conn)

    for i := 0; i < rows; i++ {
	    manager.Insert(&((*item)[i]))
    }
}

func (c *WebfaqController) Update(item *models.Webfaq) {
    
    
	conn := c.NewConnection()

	manager := models.NewWebfaqManager(conn)
	manager.Update(item)
}

func (c *WebfaqController) Delete(item *models.Webfaq) {
    
    
    conn := c.NewConnection()

	manager := models.NewWebfaqManager(conn)

    
	manager.Delete(item.Id)
}

func (c *WebfaqController) Deletebatch(item *[]models.Webfaq) {
    
    
    conn := c.NewConnection()

	manager := models.NewWebfaqManager(conn)

    for _, v := range *item {
        
    
	    manager.Delete(v.Id)
    }
}

func (c *WebfaqController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewWebfaqManager(conn)

    var args []interface{}
    
    _category := c.Geti("category")
    if _category != 0 {
        args = append(args, models.Where{Column:"category", Value:_category, Compare:"="})    
    }
    _title := c.Get("title")
    if _title != "" {
        args = append(args, models.Where{Column:"title", Value:_title, Compare:"like"})
        
    }
    _content := c.Get("content")
    if _content != "" {
        args = append(args, models.Where{Column:"content", Value:_content, Compare:"="})
        
    }
    _order := c.Geti("order")
    if _order != 0 {
        args = append(args, models.Where{Column:"order", Value:_order, Compare:"="})    
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


func (c *WebfaqController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewWebfaqManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *WebfaqController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewWebfaqManager(conn)

    var args []interface{}
    
    _category := c.Geti("category")
    if _category != 0 {
        args = append(args, models.Where{Column:"category", Value:_category, Compare:"="})    
    }
    _title := c.Get("title")
    if _title != "" {
        args = append(args, models.Where{Column:"title", Value:_title, Compare:"like"})
        
    }
    _content := c.Get("content")
    if _content != "" {
        args = append(args, models.Where{Column:"content", Value:_content, Compare:"="})
        
    }
    _order := c.Geti("order")
    if _order != 0 {
        args = append(args, models.Where{Column:"order", Value:_order, Compare:"="})    
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
                    str += ", wf_" + strings.Trim(v, " ")                
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
func (c *WebfaqController) UpdateCategory(category int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewWebfaqManager(conn)
	_manager.UpdateCategory(category, id)
}

// @Put()
func (c *WebfaqController) UpdateTitle(title string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewWebfaqManager(conn)
	_manager.UpdateTitle(title, id)
}

// @Put()
func (c *WebfaqController) UpdateContent(content string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewWebfaqManager(conn)
	_manager.UpdateContent(content, id)
}

// @Put()
func (c *WebfaqController) UpdateOrder(order int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewWebfaqManager(conn)
	_manager.UpdateOrder(order, id)
}





