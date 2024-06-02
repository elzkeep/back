package rest


import (
	"zkeep/controllers"
	"zkeep/models"

    "strings"
)

type LicenselevelController struct {
	controllers.Controller
}



func (c *LicenselevelController) GetByName(name string) *models.Licenselevel {
    
    conn := c.NewConnection()

	_manager := models.NewLicenselevelManager(conn)
    
    item := _manager.GetByName(name)
    
    c.Set("item", item)
    
    
    
    return item
    
}


func (c *LicenselevelController) Insert(item *models.Licenselevel) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewLicenselevelManager(conn)
	manager.Insert(item)

    id := manager.GetIdentity()
    c.Result["id"] = id
    item.Id = id
}

func (c *LicenselevelController) Insertbatch(item *[]models.Licenselevel) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewLicenselevelManager(conn)

    for i := 0; i < rows; i++ {
	    manager.Insert(&((*item)[i]))
    }
}

func (c *LicenselevelController) Update(item *models.Licenselevel) {
    
    
	conn := c.NewConnection()

	manager := models.NewLicenselevelManager(conn)
	manager.Update(item)
}

func (c *LicenselevelController) Delete(item *models.Licenselevel) {
    
    
    conn := c.NewConnection()

	manager := models.NewLicenselevelManager(conn)

    
	manager.Delete(item.Id)
}

func (c *LicenselevelController) Deletebatch(item *[]models.Licenselevel) {
    
    
    conn := c.NewConnection()

	manager := models.NewLicenselevelManager(conn)

    for _, v := range *item {
        
    
	    manager.Delete(v.Id)
    }
}

func (c *LicenselevelController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewLicenselevelManager(conn)

    var args []interface{}
    
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"="})
        
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


func (c *LicenselevelController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewLicenselevelManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *LicenselevelController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewLicenselevelManager(conn)

    var args []interface{}
    
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"="})
        
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
                    str += ", ll_" + strings.Trim(v, " ")                
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
func (c *LicenselevelController) UpdateName(name string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewLicenselevelManager(conn)
	_manager.UpdateName(name, id)
}
// @Put()
func (c *LicenselevelController) UpdateOrder(order int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewLicenselevelManager(conn)
	_manager.UpdateOrder(order, id)
}





