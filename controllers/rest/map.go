package rest


import (
	"aoi/controllers"
	"aoi/models"

    "strings"
)

type MapController struct {
	controllers.Controller
}

func (c *MapController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewMapManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *MapController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewMapManager(conn)

    var args []interface{}
    
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"="})
        
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
                    str += ", m_" + strings.Trim(v, " ")                
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

func (c *MapController) Insert(item *models.Map) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewMapManager(conn)
	manager.Insert(item)

    id := manager.GetIdentity()
    c.Result["id"] = id
    item.Id = id
}

func (c *MapController) Insertbatch(item *[]models.Map) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewMapManager(conn)

    for i := 0; i < rows; i++ {
	    manager.Insert(&((*item)[i]))
    }
}

func (c *MapController) Update(item *models.Map) {
    
    
	conn := c.NewConnection()

	manager := models.NewMapManager(conn)
	manager.Update(item)
}

func (c *MapController) Delete(item *models.Map) {
    
    
    conn := c.NewConnection()

	manager := models.NewMapManager(conn)

    
	manager.Delete(item.Id)
}

func (c *MapController) Deletebatch(item *[]models.Map) {
    
    
    conn := c.NewConnection()

	manager := models.NewMapManager(conn)

    for _, v := range *item {
        
    
	    manager.Delete(v.Id)
    }
}



// @Put()
func (c *MapController) UpdateName(name string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewMapManager(conn)
	_manager.UpdateName(name, id)
}

// @Put()
func (c *MapController) UpdateContent(content string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewMapManager(conn)
	_manager.UpdateContent(content, id)
}

// @Put()
func (c *MapController) UpdateOrder(order int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewMapManager(conn)
	_manager.UpdateOrder(order, id)
}






