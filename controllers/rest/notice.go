package rest


import (
	"zkeep/controllers"
	"zkeep/models"

    "strings"
)

type NoticeController struct {
	controllers.Controller
}

func (c *NoticeController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewNoticeManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *NoticeController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewNoticeManager(conn)

    var args []interface{}
    
    _title := c.Get("title")
    if _title != "" {
        args = append(args, models.Where{Column:"title", Value:_title, Compare:"="})
        
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
                    str += ", n_" + strings.Trim(v, " ")                
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

func (c *NoticeController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewNoticeManager(conn)

    var args []interface{}
    
    _title := c.Get("title")
    if _title != "" {
        args = append(args, models.Where{Column:"title", Value:_title, Compare:"="})
        
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

func (c *NoticeController) Insert(item *models.Notice) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewNoticeManager(conn)
	manager.Insert(item)

    id := manager.GetIdentity()
    c.Result["id"] = id
    item.Id = id
}

func (c *NoticeController) Insertbatch(item *[]models.Notice) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewNoticeManager(conn)

    for i := 0; i < rows; i++ {
	    manager.Insert(&((*item)[i]))
    }
}

func (c *NoticeController) Update(item *models.Notice) {
    
    
	conn := c.NewConnection()

	manager := models.NewNoticeManager(conn)
	manager.Update(item)
}

func (c *NoticeController) Delete(item *models.Notice) {
    
    
    conn := c.NewConnection()

	manager := models.NewNoticeManager(conn)

    
	manager.Delete(item.Id)
}

func (c *NoticeController) Deletebatch(item *[]models.Notice) {
    
    
    conn := c.NewConnection()

	manager := models.NewNoticeManager(conn)

    for _, v := range *item {
        
    
	    manager.Delete(v.Id)
    }
}



// @Put()
func (c *NoticeController) UpdateTitle(title string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewNoticeManager(conn)
	_manager.UpdateTitle(title, id)
}

// @Put()
func (c *NoticeController) UpdateContent(content string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewNoticeManager(conn)
	_manager.UpdateContent(content, id)
}






