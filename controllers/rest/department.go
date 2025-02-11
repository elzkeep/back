package rest


import (
	"zkeep/controllers"
	"zkeep/models"

    "strings"
)

type DepartmentController struct {
	controllers.Controller
}



func (c *DepartmentController) GetByCompanyName(company int64 ,name string) *models.Department {
    
    conn := c.NewConnection()

	_manager := models.NewDepartmentManager(conn)
    
    item := _manager.GetByCompanyName(company, name)
    
    c.Set("item", item)
    
    
    
    return item
    
}


func (c *DepartmentController) Insert(item *models.Department) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewDepartmentManager(conn)
	manager.Insert(item)

    id := manager.GetIdentity()
    c.Result["id"] = id
    item.Id = id
}

func (c *DepartmentController) Insertbatch(item *[]models.Department) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewDepartmentManager(conn)

    for i := 0; i < rows; i++ {
	    manager.Insert(&((*item)[i]))
    }
}

func (c *DepartmentController) Update(item *models.Department) {
    
    
	conn := c.NewConnection()

	manager := models.NewDepartmentManager(conn)
	manager.Update(item)
}

func (c *DepartmentController) Delete(item *models.Department) {
    
    
    conn := c.NewConnection()

	manager := models.NewDepartmentManager(conn)

    
	manager.Delete(item.Id)
}

func (c *DepartmentController) Deletebatch(item *[]models.Department) {
    
    
    conn := c.NewConnection()

	manager := models.NewDepartmentManager(conn)

    for _, v := range *item {
        
    
	    manager.Delete(v.Id)
    }
}

func (c *DepartmentController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewDepartmentManager(conn)

    var args []interface{}
    
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"="})
        
    }
    _status := c.Geti("status")
    if _status != 0 {
        args = append(args, models.Where{Column:"status", Value:_status, Compare:"="})    
    }
    _order := c.Geti("order")
    if _order != 0 {
        args = append(args, models.Where{Column:"order", Value:_order, Compare:"="})    
    }
    _parent := c.Geti64("parent")
    if _parent != 0 {
        args = append(args, models.Where{Column:"parent", Value:_parent, Compare:"="})    
    }
    _company := c.Geti64("company")
    if _company != 0 {
        args = append(args, models.Where{Column:"company", Value:_company, Compare:"="})    
    }
    _master := c.Geti64("master")
    if _master != 0 {
        args = append(args, models.Where{Column:"master", Value:_master, Compare:"="})    
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


func (c *DepartmentController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewDepartmentManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *DepartmentController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewDepartmentManager(conn)

    var args []interface{}
    
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"="})
        
    }
    _status := c.Geti("status")
    if _status != 0 {
        args = append(args, models.Where{Column:"status", Value:_status, Compare:"="})    
    }
    _order := c.Geti("order")
    if _order != 0 {
        args = append(args, models.Where{Column:"order", Value:_order, Compare:"="})    
    }
    _parent := c.Geti64("parent")
    if _parent != 0 {
        args = append(args, models.Where{Column:"parent", Value:_parent, Compare:"="})    
    }
    _company := c.Geti64("company")
    if _company != 0 {
        args = append(args, models.Where{Column:"company", Value:_company, Compare:"="})    
    }
    _master := c.Geti64("master")
    if _master != 0 {
        args = append(args, models.Where{Column:"master", Value:_master, Compare:"="})    
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
                    str += ", de_" + strings.Trim(v, " ")                
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
func (c *DepartmentController) UpdateName(name string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewDepartmentManager(conn)
	_manager.UpdateName(name, id)
}
// @Put()
func (c *DepartmentController) UpdateStatus(status int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewDepartmentManager(conn)
	_manager.UpdateStatus(status, id)
}
// @Put()
func (c *DepartmentController) UpdateOrder(order int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewDepartmentManager(conn)
	_manager.UpdateOrder(order, id)
}
// @Put()
func (c *DepartmentController) UpdateParent(parent int64, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewDepartmentManager(conn)
	_manager.UpdateParent(parent, id)
}
// @Put()
func (c *DepartmentController) UpdateCompany(company int64, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewDepartmentManager(conn)
	_manager.UpdateCompany(company, id)
}
// @Put()
func (c *DepartmentController) UpdateMaster(master int64, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewDepartmentManager(conn)
	_manager.UpdateMaster(master, id)
}





