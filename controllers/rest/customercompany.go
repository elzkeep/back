package rest


import (
	"zkeep/controllers"
	"zkeep/models"

    "strings"
)

type CustomercompanyController struct {
	controllers.Controller
}

func (c *CustomercompanyController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewCustomercompanyManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *CustomercompanyController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewCustomercompanyManager(conn)

    var args []interface{}
    
    _company := c.Geti64("company")
    if _company != 0 {
        args = append(args, models.Where{Column:"company", Value:_company, Compare:"="})    
    }
    _customer := c.Geti64("customer")
    if _customer != 0 {
        args = append(args, models.Where{Column:"customer", Value:_customer, Compare:"="})    
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
                    str += ", cc_" + strings.Trim(v, " ")                
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

func (c *CustomercompanyController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewCustomercompanyManager(conn)

    var args []interface{}
    
    _company := c.Geti64("company")
    if _company != 0 {
        args = append(args, models.Where{Column:"company", Value:_company, Compare:"="})    
    }
    _customer := c.Geti64("customer")
    if _customer != 0 {
        args = append(args, models.Where{Column:"customer", Value:_customer, Compare:"="})    
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

func (c *CustomercompanyController) Insert(item *models.Customercompany) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewCustomercompanyManager(conn)
	manager.Insert(item)

    id := manager.GetIdentity()
    c.Result["id"] = id
    item.Id = id
}

func (c *CustomercompanyController) Insertbatch(item *[]models.Customercompany) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewCustomercompanyManager(conn)

    for i := 0; i < rows; i++ {
	    manager.Insert(&((*item)[i]))
    }
}

func (c *CustomercompanyController) Update(item *models.Customercompany) {
    
    
	conn := c.NewConnection()

	manager := models.NewCustomercompanyManager(conn)
	manager.Update(item)
}

func (c *CustomercompanyController) Delete(item *models.Customercompany) {
    
    
    conn := c.NewConnection()

	manager := models.NewCustomercompanyManager(conn)

    
	manager.Delete(item.Id)
}

func (c *CustomercompanyController) Deletebatch(item *[]models.Customercompany) {
    
    
    conn := c.NewConnection()

	manager := models.NewCustomercompanyManager(conn)

    for _, v := range *item {
        
    
	    manager.Delete(v.Id)
    }
}



func (c *CustomercompanyController) GetByCompanyCustomer(company int64 ,customer int64) *models.Customercompany {
    
    conn := c.NewConnection()

	_manager := models.NewCustomercompanyManager(conn)
    
    item := _manager.GetByCompanyCustomer(company, customer)
    
    c.Set("item", item)
    
    
    
    return item
    
}


// @Put()
func (c *CustomercompanyController) UpdateCompany(company int64, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCustomercompanyManager(conn)
	_manager.UpdateCompany(company, id)
}

// @Put()
func (c *CustomercompanyController) UpdateCustomer(customer int64, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCustomercompanyManager(conn)
	_manager.UpdateCustomer(customer, id)
}






