package rest


import (
	"zkeep/controllers"
	"zkeep/models"

    "strings"
)

type GiroController struct {
	controllers.Controller
}



func (c *GiroController) Insert(item *models.Giro) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewGiroManager(conn)
	manager.Insert(item)

    id := manager.GetIdentity()
    c.Result["id"] = id
    item.Id = id
}

func (c *GiroController) Insertbatch(item *[]models.Giro) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewGiroManager(conn)

    for i := 0; i < rows; i++ {
	    manager.Insert(&((*item)[i]))
    }
}

func (c *GiroController) Update(item *models.Giro) {
    
    
	conn := c.NewConnection()

	manager := models.NewGiroManager(conn)
	manager.Update(item)
}

func (c *GiroController) Delete(item *models.Giro) {
    
    
    conn := c.NewConnection()

	manager := models.NewGiroManager(conn)

    
	manager.Delete(item.Id)
}

func (c *GiroController) Deletebatch(item *[]models.Giro) {
    
    
    conn := c.NewConnection()

	manager := models.NewGiroManager(conn)

    for _, v := range *item {
        
    
	    manager.Delete(v.Id)
    }
}

func (c *GiroController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewGiroManager(conn)

    var args []interface{}
    
    _startinsertdate := c.Get("startinsertdate")
    _endinsertdate := c.Get("endinsertdate")
    if _startinsertdate != "" && _endinsertdate != "" {        
        var v [2]string
        v[0] = _startinsertdate
        v[1] = _endinsertdate  
        args = append(args, models.Where{Column:"insertdate", Value:v, Compare:"between"})    
    } else if  _startinsertdate != "" {          
        args = append(args, models.Where{Column:"insertdate", Value:_startinsertdate, Compare:">="})
    } else if  _endinsertdate != "" {          
        args = append(args, models.Where{Column:"insertdate", Value:_endinsertdate, Compare:"<="})            
    }
    _number := c.Get("number")
    if _number != "" {
        args = append(args, models.Where{Column:"number", Value:_number, Compare:"="})
    }
    _price := c.Geti("price")
    if _price != 0 {
        args = append(args, models.Where{Column:"price", Value:_price, Compare:"="})    
    }
    _startacceptdate := c.Get("startacceptdate")
    _endacceptdate := c.Get("endacceptdate")
    if _startacceptdate != "" && _endacceptdate != "" {        
        var v [2]string
        v[0] = _startacceptdate
        v[1] = _endacceptdate  
        args = append(args, models.Where{Column:"acceptdate", Value:v, Compare:"between"})    
    } else if  _startacceptdate != "" {          
        args = append(args, models.Where{Column:"acceptdate", Value:_startacceptdate, Compare:">="})
    } else if  _endacceptdate != "" {          
        args = append(args, models.Where{Column:"acceptdate", Value:_endacceptdate, Compare:"<="})            
    }
    _charge := c.Geti("charge")
    if _charge != 0 {
        args = append(args, models.Where{Column:"charge", Value:_charge, Compare:"="})    
    }
    _type := c.Get("type")
    if _type != "" {
        args = append(args, models.Where{Column:"type", Value:_type, Compare:"="})
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


func (c *GiroController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewGiroManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *GiroController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewGiroManager(conn)

    var args []interface{}
    
    _startinsertdate := c.Get("startinsertdate")
    _endinsertdate := c.Get("endinsertdate")
    if _startinsertdate != "" && _endinsertdate != "" {        
        var v [2]string
        v[0] = _startinsertdate
        v[1] = _endinsertdate  
        args = append(args, models.Where{Column:"insertdate", Value:v, Compare:"between"})    
    } else if  _startinsertdate != "" {          
        args = append(args, models.Where{Column:"insertdate", Value:_startinsertdate, Compare:">="})
    } else if  _endinsertdate != "" {          
        args = append(args, models.Where{Column:"insertdate", Value:_endinsertdate, Compare:"<="})            
    }
    _number := c.Get("number")
    if _number != "" {
        args = append(args, models.Where{Column:"number", Value:_number, Compare:"="})
    }
    _price := c.Geti("price")
    if _price != 0 {
        args = append(args, models.Where{Column:"price", Value:_price, Compare:"="})    
    }
    _startacceptdate := c.Get("startacceptdate")
    _endacceptdate := c.Get("endacceptdate")
    if _startacceptdate != "" && _endacceptdate != "" {        
        var v [2]string
        v[0] = _startacceptdate
        v[1] = _endacceptdate  
        args = append(args, models.Where{Column:"acceptdate", Value:v, Compare:"between"})    
    } else if  _startacceptdate != "" {          
        args = append(args, models.Where{Column:"acceptdate", Value:_startacceptdate, Compare:">="})
    } else if  _endacceptdate != "" {          
        args = append(args, models.Where{Column:"acceptdate", Value:_endacceptdate, Compare:"<="})            
    }
    _charge := c.Geti("charge")
    if _charge != 0 {
        args = append(args, models.Where{Column:"charge", Value:_charge, Compare:"="})    
    }
    _type := c.Get("type")
    if _type != "" {
        args = append(args, models.Where{Column:"type", Value:_type, Compare:"="})
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

// @Put()
func (c *GiroController) UpdateInsertdate(insertdate string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewGiroManager(conn)
	_manager.UpdateInsertdate(insertdate, id)
}
// @Put()
func (c *GiroController) UpdateNumber(number string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewGiroManager(conn)
	_manager.UpdateNumber(number, id)
}
// @Put()
func (c *GiroController) UpdatePrice(price int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewGiroManager(conn)
	_manager.UpdatePrice(price, id)
}
// @Put()
func (c *GiroController) UpdateAcceptdate(acceptdate string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewGiroManager(conn)
	_manager.UpdateAcceptdate(acceptdate, id)
}
// @Put()
func (c *GiroController) UpdateCharge(charge int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewGiroManager(conn)
	_manager.UpdateCharge(charge, id)
}
// @Put()
func (c *GiroController) UpdateType(typeid string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewGiroManager(conn)
	_manager.UpdateType(typeid, id)
}
// @Put()
func (c *GiroController) UpdateContent(content string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewGiroManager(conn)
	_manager.UpdateContent(content, id)
}





func (c *GiroController) Sum() {
    
    
	conn := c.NewConnection()

	manager := models.NewGiroManager(conn)

    var args []interface{}
    
    _startinsertdate := c.Get("startinsertdate")
    _endinsertdate := c.Get("endinsertdate")
    if _startinsertdate != "" && _endinsertdate != "" {        
        var v [2]string
        v[0] = _startinsertdate
        v[1] = _endinsertdate  
        args = append(args, models.Where{Column:"insertdate", Value:v, Compare:"between"})    
    } else if  _startinsertdate != "" {          
        args = append(args, models.Where{Column:"insertdate", Value:_startinsertdate, Compare:">="})
    } else if  _endinsertdate != "" {          
        args = append(args, models.Where{Column:"insertdate", Value:_endinsertdate, Compare:"<="})            
    }
    _number := c.Get("number")
    if _number != "" {
        args = append(args, models.Where{Column:"number", Value:_number, Compare:"like"})
    }
    _price := c.Geti("price")
    if _price != 0 {
        args = append(args, models.Where{Column:"price", Value:_price, Compare:"="})    
    }
    _startacceptdate := c.Get("startacceptdate")
    _endacceptdate := c.Get("endacceptdate")
    if _startacceptdate != "" && _endacceptdate != "" {        
        var v [2]string
        v[0] = _startacceptdate
        v[1] = _endacceptdate  
        args = append(args, models.Where{Column:"acceptdate", Value:v, Compare:"between"})    
    } else if  _startacceptdate != "" {          
        args = append(args, models.Where{Column:"acceptdate", Value:_startacceptdate, Compare:">="})
    } else if  _endacceptdate != "" {          
        args = append(args, models.Where{Column:"acceptdate", Value:_endacceptdate, Compare:"<="})            
    }
    _charge := c.Geti("charge")
    if _charge != 0 {
        args = append(args, models.Where{Column:"charge", Value:_charge, Compare:"="})    
    }
    _type := c.Get("type")
    if _type != "" {
        args = append(args, models.Where{Column:"type", Value:_type, Compare:"like"})
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
    

    
    
    item := manager.Sum(args)
	c.Set("item", item)
}

