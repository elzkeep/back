package rest


import (
	"zkeep/controllers"
	"zkeep/models"

    "strings"
)

type BillinghistoryController struct {
	controllers.Controller
}


// @Delete()
func (c *BillinghistoryController) DeleteByBilling(billing int64) {
    
    conn := c.NewConnection()

	_manager := models.NewBillinghistoryManager(conn)
    
    _manager.DeleteByBilling(billing)
    
}


func (c *BillinghistoryController) Insert(item *models.Billinghistory) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewBillinghistoryManager(conn)
	manager.Insert(item)

    id := manager.GetIdentity()
    c.Result["id"] = id
    item.Id = id
}

func (c *BillinghistoryController) Insertbatch(item *[]models.Billinghistory) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewBillinghistoryManager(conn)

    for i := 0; i < rows; i++ {
	    manager.Insert(&((*item)[i]))
    }
}

func (c *BillinghistoryController) Update(item *models.Billinghistory) {
    
    
	conn := c.NewConnection()

	manager := models.NewBillinghistoryManager(conn)
	manager.Update(item)
}

func (c *BillinghistoryController) Delete(item *models.Billinghistory) {
    
    
    conn := c.NewConnection()

	manager := models.NewBillinghistoryManager(conn)

    
	manager.Delete(item.Id)
}

func (c *BillinghistoryController) Deletebatch(item *[]models.Billinghistory) {
    
    
    conn := c.NewConnection()

	manager := models.NewBillinghistoryManager(conn)

    for _, v := range *item {
        
    
	    manager.Delete(v.Id)
    }
}

func (c *BillinghistoryController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewBillinghistoryManager(conn)

    var args []interface{}
    
    _price := c.Geti("price")
    if _price != 0 {
        args = append(args, models.Where{Column:"price", Value:_price, Compare:"="})    
    }
    _company := c.Geti64("company")
    if _company != 0 {
        args = append(args, models.Where{Column:"company", Value:_company, Compare:"="})    
    }
    _building := c.Geti64("building")
    if _building != 0 {
        args = append(args, models.Where{Column:"building", Value:_building, Compare:"="})    
    }
    _billing := c.Geti64("billing")
    if _billing != 0 {
        args = append(args, models.Where{Column:"billing", Value:_billing, Compare:"="})    
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


func (c *BillinghistoryController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewBillinghistoryManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *BillinghistoryController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewBillinghistoryManager(conn)

    var args []interface{}
    
    _price := c.Geti("price")
    if _price != 0 {
        args = append(args, models.Where{Column:"price", Value:_price, Compare:"="})    
    }
    _company := c.Geti64("company")
    if _company != 0 {
        args = append(args, models.Where{Column:"company", Value:_company, Compare:"="})    
    }
    _building := c.Geti64("building")
    if _building != 0 {
        args = append(args, models.Where{Column:"building", Value:_building, Compare:"="})    
    }
    _billing := c.Geti64("billing")
    if _billing != 0 {
        args = append(args, models.Where{Column:"billing", Value:_billing, Compare:"="})    
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
                    str += ", bh_" + strings.Trim(v, " ")                
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
func (c *BillinghistoryController) UpdatePrice(price int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewBillinghistoryManager(conn)
	_manager.UpdatePrice(price, id)
}
// @Put()
func (c *BillinghistoryController) UpdateCompany(company int64, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewBillinghistoryManager(conn)
	_manager.UpdateCompany(company, id)
}
// @Put()
func (c *BillinghistoryController) UpdateBuilding(building int64, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewBillinghistoryManager(conn)
	_manager.UpdateBuilding(building, id)
}
// @Put()
func (c *BillinghistoryController) UpdateBilling(billing int64, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewBillinghistoryManager(conn)
	_manager.UpdateBilling(billing, id)
}





func (c *BillinghistoryController) Sum() {
    
    
	conn := c.NewConnection()

	manager := models.NewBillinghistoryManager(conn)

    var args []interface{}
    
    _price := c.Geti("price")
    if _price != 0 {
        args = append(args, models.Where{Column:"price", Value:_price, Compare:"="})    
    }
    _company := c.Geti64("company")
    if _company != 0 {
        args = append(args, models.Where{Column:"company", Value:_company, Compare:"="})    
    }
    _building := c.Geti64("building")
    if _building != 0 {
        args = append(args, models.Where{Column:"building", Value:_building, Compare:"="})    
    }
    _billing := c.Geti64("billing")
    if _billing != 0 {
        args = append(args, models.Where{Column:"billing", Value:_billing, Compare:"="})    
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

