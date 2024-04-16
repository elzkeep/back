package rest


import (
	"zkeep/controllers"
	"zkeep/models"

    "strings"
)

type BillingController struct {
	controllers.Controller
}



func (c *BillingController) Insert(item *models.Billing) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewBillingManager(conn)
	manager.Insert(item)

    id := manager.GetIdentity()
    c.Result["id"] = id
    item.Id = id
}

func (c *BillingController) Insertbatch(item *[]models.Billing) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewBillingManager(conn)

    for i := 0; i < rows; i++ {
	    manager.Insert(&((*item)[i]))
    }
}

func (c *BillingController) Update(item *models.Billing) {
    
    
	conn := c.NewConnection()

	manager := models.NewBillingManager(conn)
	manager.Update(item)
}

func (c *BillingController) Delete(item *models.Billing) {
    
    
    conn := c.NewConnection()

	manager := models.NewBillingManager(conn)

    
	manager.Delete(item.Id)
}

func (c *BillingController) Deletebatch(item *[]models.Billing) {
    
    
    conn := c.NewConnection()

	manager := models.NewBillingManager(conn)

    for _, v := range *item {
        
    
	    manager.Delete(v.Id)
    }
}

func (c *BillingController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewBillingManager(conn)

    var args []interface{}
    
    _price := c.Geti("price")
    if _price != 0 {
        args = append(args, models.Where{Column:"price", Value:_price, Compare:"="})    
    }
    _status := c.Geti("status")
    if _status != 0 {
        args = append(args, models.Where{Column:"status", Value:_status, Compare:"="})    
    }
    _giro := c.Geti("giro")
    if _giro != 0 {
        args = append(args, models.Where{Column:"giro", Value:_giro, Compare:"="})    
    }
    _startbilldate := c.Get("startbilldate")
    _endbilldate := c.Get("endbilldate")
    if _startbilldate != "" && _endbilldate != "" {        
        var v [2]string
        v[0] = _startbilldate
        v[1] = _endbilldate  
        args = append(args, models.Where{Column:"billdate", Value:v, Compare:"between"})    
    } else if  _startbilldate != "" {          
        args = append(args, models.Where{Column:"billdate", Value:_startbilldate, Compare:">="})
    } else if  _endbilldate != "" {          
        args = append(args, models.Where{Column:"billdate", Value:_endbilldate, Compare:"<="})            
    }
    _month := c.Get("month")
    if _month != "" {
        args = append(args, models.Where{Column:"month", Value:_month, Compare:"like"})
    }
    _period := c.Geti("period")
    if _period != 0 {
        args = append(args, models.Where{Column:"period", Value:_period, Compare:"="})    
    }
    _company := c.Geti64("company")
    if _company != 0 {
        args = append(args, models.Where{Column:"company", Value:_company, Compare:"="})    
    }
    _building := c.Geti64("building")
    if _building != 0 {
        args = append(args, models.Where{Column:"building", Value:_building, Compare:"="})    
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


func (c *BillingController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewBillingManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *BillingController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewBillingManager(conn)

    var args []interface{}
    
    _price := c.Geti("price")
    if _price != 0 {
        args = append(args, models.Where{Column:"price", Value:_price, Compare:"="})    
    }
    _status := c.Geti("status")
    if _status != 0 {
        args = append(args, models.Where{Column:"status", Value:_status, Compare:"="})    
    }
    _giro := c.Geti("giro")
    if _giro != 0 {
        args = append(args, models.Where{Column:"giro", Value:_giro, Compare:"="})    
    }
    _startbilldate := c.Get("startbilldate")
    _endbilldate := c.Get("endbilldate")
    if _startbilldate != "" && _endbilldate != "" {        
        var v [2]string
        v[0] = _startbilldate
        v[1] = _endbilldate  
        args = append(args, models.Where{Column:"billdate", Value:v, Compare:"between"})    
    } else if  _startbilldate != "" {          
        args = append(args, models.Where{Column:"billdate", Value:_startbilldate, Compare:">="})
    } else if  _endbilldate != "" {          
        args = append(args, models.Where{Column:"billdate", Value:_endbilldate, Compare:"<="})            
    }
    _month := c.Get("month")
    if _month != "" {
        args = append(args, models.Where{Column:"month", Value:_month, Compare:"like"})
    }
    _period := c.Geti("period")
    if _period != 0 {
        args = append(args, models.Where{Column:"period", Value:_period, Compare:"="})    
    }
    _company := c.Geti64("company")
    if _company != 0 {
        args = append(args, models.Where{Column:"company", Value:_company, Compare:"="})    
    }
    _building := c.Geti64("building")
    if _building != 0 {
        args = append(args, models.Where{Column:"building", Value:_building, Compare:"="})    
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
                    str += ", bi_" + strings.Trim(v, " ")                
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
func (c *BillingController) UpdatePrice(price int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewBillingManager(conn)
	_manager.UpdatePrice(price, id)
}

// @Put()
func (c *BillingController) UpdateStatus(status int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewBillingManager(conn)
	_manager.UpdateStatus(status, id)
}

// @Put()
func (c *BillingController) UpdateGiro(giro int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewBillingManager(conn)
	_manager.UpdateGiro(giro, id)
}

// @Put()
func (c *BillingController) UpdateBilldate(billdate string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewBillingManager(conn)
	_manager.UpdateBilldate(billdate, id)
}

// @Put()
func (c *BillingController) UpdateMonth(month string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewBillingManager(conn)
	_manager.UpdateMonth(month, id)
}

// @Put()
func (c *BillingController) UpdatePeriod(period int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewBillingManager(conn)
	_manager.UpdatePeriod(period, id)
}

// @Put()
func (c *BillingController) UpdateCompany(company int64, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewBillingManager(conn)
	_manager.UpdateCompany(company, id)
}

// @Put()
func (c *BillingController) UpdateBuilding(building int64, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewBillingManager(conn)
	_manager.UpdateBuilding(building, id)
}






func (c *BillingController) Sum() {
    
    
	conn := c.NewConnection()

	manager := models.NewBillingManager(conn)

    var args []interface{}
    
    _price := c.Geti("price")
    if _price != 0 {
        args = append(args, models.Where{Column:"price", Value:_price, Compare:"="})    
    }
    _status := c.Geti("status")
    if _status != 0 {
        args = append(args, models.Where{Column:"status", Value:_status, Compare:"="})    
    }
    _giro := c.Geti("giro")
    if _giro != 0 {
        args = append(args, models.Where{Column:"giro", Value:_giro, Compare:"="})    
    }
    _startbilldate := c.Get("startbilldate")
    _endbilldate := c.Get("endbilldate")
    if _startbilldate != "" && _endbilldate != "" {        
        var v [2]string
        v[0] = _startbilldate
        v[1] = _endbilldate  
        args = append(args, models.Where{Column:"billdate", Value:v, Compare:"between"})    
    } else if  _startbilldate != "" {          
        args = append(args, models.Where{Column:"billdate", Value:_startbilldate, Compare:">="})
    } else if  _endbilldate != "" {          
        args = append(args, models.Where{Column:"billdate", Value:_endbilldate, Compare:"<="})            
    }
    _month := c.Get("month")
    if _month != "" {
        args = append(args, models.Where{Column:"month", Value:_month, Compare:"like"})
    }
    _period := c.Geti("period")
    if _period != 0 {
        args = append(args, models.Where{Column:"period", Value:_period, Compare:"="})    
    }
    _company := c.Geti64("company")
    if _company != 0 {
        args = append(args, models.Where{Column:"company", Value:_company, Compare:"="})    
    }
    _building := c.Geti64("building")
    if _building != 0 {
        args = append(args, models.Where{Column:"building", Value:_building, Compare:"="})    
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

