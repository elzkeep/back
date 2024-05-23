package rest


import (
	"zkeep/controllers"
	"zkeep/models"

    "strings"
)

type WebquestionController struct {
	controllers.Controller
}



func (c *WebquestionController) Insert(item *models.Webquestion) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewWebquestionManager(conn)
	manager.Insert(item)

    id := manager.GetIdentity()
    c.Result["id"] = id
    item.Id = id
}

func (c *WebquestionController) Insertbatch(item *[]models.Webquestion) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewWebquestionManager(conn)

    for i := 0; i < rows; i++ {
	    manager.Insert(&((*item)[i]))
    }
}

func (c *WebquestionController) Update(item *models.Webquestion) {
    
    
	conn := c.NewConnection()

	manager := models.NewWebquestionManager(conn)
	manager.Update(item)
}

func (c *WebquestionController) Delete(item *models.Webquestion) {
    
    
    conn := c.NewConnection()

	manager := models.NewWebquestionManager(conn)

    
	manager.Delete(item.Id)
}

func (c *WebquestionController) Deletebatch(item *[]models.Webquestion) {
    
    
    conn := c.NewConnection()

	manager := models.NewWebquestionManager(conn)

    for _, v := range *item {
        
    
	    manager.Delete(v.Id)
    }
}

func (c *WebquestionController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewWebquestionManager(conn)

    var args []interface{}
    
    _email := c.Get("email")
    if _email != "" {
        args = append(args, models.Where{Column:"email", Value:_email, Compare:"="})
    }
    _tel := c.Get("tel")
    if _tel != "" {
        args = append(args, models.Where{Column:"tel", Value:_tel, Compare:"="})
    }
    _content := c.Get("content")
    if _content != "" {
        args = append(args, models.Where{Column:"content", Value:_content, Compare:"="})
        
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


func (c *WebquestionController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewWebquestionManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *WebquestionController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewWebquestionManager(conn)

    var args []interface{}
    
    _email := c.Get("email")
    if _email != "" {
        args = append(args, models.Where{Column:"email", Value:_email, Compare:"="})
    }
    _tel := c.Get("tel")
    if _tel != "" {
        args = append(args, models.Where{Column:"tel", Value:_tel, Compare:"="})
    }
    _content := c.Get("content")
    if _content != "" {
        args = append(args, models.Where{Column:"content", Value:_content, Compare:"="})
        
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
                    str += ", wq_" + strings.Trim(v, " ")                
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
func (c *WebquestionController) UpdateEmail(email string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewWebquestionManager(conn)
	_manager.UpdateEmail(email, id)
}

// @Put()
func (c *WebquestionController) UpdateTel(tel string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewWebquestionManager(conn)
	_manager.UpdateTel(tel, id)
}

// @Put()
func (c *WebquestionController) UpdateContent(content string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewWebquestionManager(conn)
	_manager.UpdateContent(content, id)
}






