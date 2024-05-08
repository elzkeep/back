package rest


import (
	"zkeep/controllers"
	"zkeep/models"

    "strings"
)

type CustomerController struct {
	controllers.Controller
}



func (c *CustomerController) CountByCompanyBuilding(company int64 ,building int64) int {
    
    conn := c.NewConnection()

	_manager := models.NewCustomerManager(conn)
    
    item := _manager.CountByCompanyBuilding(company, building)
    
    
    
    c.Set("count", item)
    
    return item
    
}


func (c *CustomerController) GetByCompanyBuilding(company int64 ,building int64) *models.Customer {
    
    conn := c.NewConnection()

	_manager := models.NewCustomerManager(conn)
    
    item := _manager.GetByCompanyBuilding(company, building)
    
    c.Set("item", item)
    
    
    
    return item
    
}

// @Delete()
func (c *CustomerController) DeleteByCompanyBuilding(company int64 ,building int64) {
    
    conn := c.NewConnection()

	_manager := models.NewCustomerManager(conn)
    
    _manager.DeleteByCompanyBuilding(company, building)
    
}


func (c *CustomerController) Insert(item *models.Customer) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewCustomerManager(conn)
	manager.Insert(item)

    id := manager.GetIdentity()
    c.Result["id"] = id
    item.Id = id
}

func (c *CustomerController) Insertbatch(item *[]models.Customer) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewCustomerManager(conn)

    for i := 0; i < rows; i++ {
	    manager.Insert(&((*item)[i]))
    }
}

func (c *CustomerController) Update(item *models.Customer) {
    
    
	conn := c.NewConnection()

	manager := models.NewCustomerManager(conn)
	manager.Update(item)
}

func (c *CustomerController) Delete(item *models.Customer) {
    
    
    conn := c.NewConnection()

	manager := models.NewCustomerManager(conn)

    
	manager.Delete(item.Id)
}

func (c *CustomerController) Deletebatch(item *[]models.Customer) {
    
    
    conn := c.NewConnection()

	manager := models.NewCustomerManager(conn)

    for _, v := range *item {
        
    
	    manager.Delete(v.Id)
    }
}

func (c *CustomerController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewCustomerManager(conn)

    var args []interface{}
    
    _number := c.Geti("number")
    if _number != 0 {
        args = append(args, models.Where{Column:"number", Value:_number, Compare:"="})    
    }
    _type := c.Geti("type")
    if _type != 0 {
        args = append(args, models.Where{Column:"type", Value:_type, Compare:"="})    
    }
    _checkdate := c.Geti("checkdate")
    if _checkdate != 0 {
        args = append(args, models.Where{Column:"checkdate", Value:_checkdate, Compare:"="})    
    }
    _managername := c.Get("managername")
    if _managername != "" {
        args = append(args, models.Where{Column:"managername", Value:_managername, Compare:"like"})
    }
    _managertel := c.Get("managertel")
    if _managertel != "" {
        args = append(args, models.Where{Column:"managertel", Value:_managertel, Compare:"like"})
    }
    _manageremail := c.Get("manageremail")
    if _manageremail != "" {
        args = append(args, models.Where{Column:"manageremail", Value:_manageremail, Compare:"like"})
    }
    _startcontractstartdate := c.Get("startcontractstartdate")
    _endcontractstartdate := c.Get("endcontractstartdate")
    if _startcontractstartdate != "" && _endcontractstartdate != "" {        
        var v [2]string
        v[0] = _startcontractstartdate
        v[1] = _endcontractstartdate  
        args = append(args, models.Where{Column:"contractstartdate", Value:v, Compare:"between"})    
    } else if  _startcontractstartdate != "" {          
        args = append(args, models.Where{Column:"contractstartdate", Value:_startcontractstartdate, Compare:">="})
    } else if  _endcontractstartdate != "" {          
        args = append(args, models.Where{Column:"contractstartdate", Value:_endcontractstartdate, Compare:"<="})            
    }
    _startcontractenddate := c.Get("startcontractenddate")
    _endcontractenddate := c.Get("endcontractenddate")
    if _startcontractenddate != "" && _endcontractenddate != "" {        
        var v [2]string
        v[0] = _startcontractenddate
        v[1] = _endcontractenddate  
        args = append(args, models.Where{Column:"contractenddate", Value:v, Compare:"between"})    
    } else if  _startcontractenddate != "" {          
        args = append(args, models.Where{Column:"contractenddate", Value:_startcontractenddate, Compare:">="})
    } else if  _endcontractenddate != "" {          
        args = append(args, models.Where{Column:"contractenddate", Value:_endcontractenddate, Compare:"<="})            
    }
    _contractprice := c.Geti("contractprice")
    if _contractprice != 0 {
        args = append(args, models.Where{Column:"contractprice", Value:_contractprice, Compare:"="})    
    }
    _contractvat := c.Geti("contractvat")
    if _contractvat != 0 {
        args = append(args, models.Where{Column:"contractvat", Value:_contractvat, Compare:"="})    
    }
    _contractday := c.Geti("contractday")
    if _contractday != 0 {
        args = append(args, models.Where{Column:"contractday", Value:_contractday, Compare:"="})    
    }
    _billingdate := c.Geti("billingdate")
    if _billingdate != 0 {
        args = append(args, models.Where{Column:"billingdate", Value:_billingdate, Compare:"="})    
    }
    _billingtype := c.Geti("billingtype")
    if _billingtype != 0 {
        args = append(args, models.Where{Column:"billingtype", Value:_billingtype, Compare:"="})    
    }
    _billingname := c.Get("billingname")
    if _billingname != "" {
        args = append(args, models.Where{Column:"billingname", Value:_billingname, Compare:"like"})
    }
    _billingtel := c.Get("billingtel")
    if _billingtel != "" {
        args = append(args, models.Where{Column:"billingtel", Value:_billingtel, Compare:"like"})
    }
    _billingemail := c.Get("billingemail")
    if _billingemail != "" {
        args = append(args, models.Where{Column:"billingemail", Value:_billingemail, Compare:"like"})
    }
    _address := c.Get("address")
    if _address != "" {
        args = append(args, models.Where{Column:"address", Value:_address, Compare:"like"})
    }
    _addressetc := c.Get("addressetc")
    if _addressetc != "" {
        args = append(args, models.Where{Column:"addressetc", Value:_addressetc, Compare:"like"})
    }
    _collectmonth := c.Geti("collectmonth")
    if _collectmonth != 0 {
        args = append(args, models.Where{Column:"collectmonth", Value:_collectmonth, Compare:"="})    
    }
    _collectday := c.Geti("collectday")
    if _collectday != 0 {
        args = append(args, models.Where{Column:"collectday", Value:_collectday, Compare:"="})    
    }
    _manager := c.Get("manager")
    if _manager != "" {
        args = append(args, models.Where{Column:"manager", Value:_manager, Compare:"like"})
    }
    _tel := c.Get("tel")
    if _tel != "" {
        args = append(args, models.Where{Column:"tel", Value:_tel, Compare:"like"})
    }
    _fax := c.Get("fax")
    if _fax != "" {
        args = append(args, models.Where{Column:"fax", Value:_fax, Compare:"like"})
    }
    _status := c.Geti("status")
    if _status != 0 {
        args = append(args, models.Where{Column:"status", Value:_status, Compare:"="})    
    }
    _salesuser := c.Geti64("salesuser")
    if _salesuser != 0 {
        args = append(args, models.Where{Column:"salesuser", Value:_salesuser, Compare:"="})    
    }
    _user := c.Geti64("user")
    if _user != 0 {
        args = append(args, models.Where{Column:"user", Value:_user, Compare:"="})    
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


func (c *CustomerController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewCustomerManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *CustomerController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewCustomerManager(conn)

    var args []interface{}
    
    _number := c.Geti("number")
    if _number != 0 {
        args = append(args, models.Where{Column:"number", Value:_number, Compare:"="})    
    }
    _type := c.Geti("type")
    if _type != 0 {
        args = append(args, models.Where{Column:"type", Value:_type, Compare:"="})    
    }
    _checkdate := c.Geti("checkdate")
    if _checkdate != 0 {
        args = append(args, models.Where{Column:"checkdate", Value:_checkdate, Compare:"="})    
    }
    _managername := c.Get("managername")
    if _managername != "" {
        args = append(args, models.Where{Column:"managername", Value:_managername, Compare:"like"})
    }
    _managertel := c.Get("managertel")
    if _managertel != "" {
        args = append(args, models.Where{Column:"managertel", Value:_managertel, Compare:"like"})
    }
    _manageremail := c.Get("manageremail")
    if _manageremail != "" {
        args = append(args, models.Where{Column:"manageremail", Value:_manageremail, Compare:"like"})
    }
    _startcontractstartdate := c.Get("startcontractstartdate")
    _endcontractstartdate := c.Get("endcontractstartdate")
    if _startcontractstartdate != "" && _endcontractstartdate != "" {        
        var v [2]string
        v[0] = _startcontractstartdate
        v[1] = _endcontractstartdate  
        args = append(args, models.Where{Column:"contractstartdate", Value:v, Compare:"between"})    
    } else if  _startcontractstartdate != "" {          
        args = append(args, models.Where{Column:"contractstartdate", Value:_startcontractstartdate, Compare:">="})
    } else if  _endcontractstartdate != "" {          
        args = append(args, models.Where{Column:"contractstartdate", Value:_endcontractstartdate, Compare:"<="})            
    }
    _startcontractenddate := c.Get("startcontractenddate")
    _endcontractenddate := c.Get("endcontractenddate")
    if _startcontractenddate != "" && _endcontractenddate != "" {        
        var v [2]string
        v[0] = _startcontractenddate
        v[1] = _endcontractenddate  
        args = append(args, models.Where{Column:"contractenddate", Value:v, Compare:"between"})    
    } else if  _startcontractenddate != "" {          
        args = append(args, models.Where{Column:"contractenddate", Value:_startcontractenddate, Compare:">="})
    } else if  _endcontractenddate != "" {          
        args = append(args, models.Where{Column:"contractenddate", Value:_endcontractenddate, Compare:"<="})            
    }
    _contractprice := c.Geti("contractprice")
    if _contractprice != 0 {
        args = append(args, models.Where{Column:"contractprice", Value:_contractprice, Compare:"="})    
    }
    _contractvat := c.Geti("contractvat")
    if _contractvat != 0 {
        args = append(args, models.Where{Column:"contractvat", Value:_contractvat, Compare:"="})    
    }
    _contractday := c.Geti("contractday")
    if _contractday != 0 {
        args = append(args, models.Where{Column:"contractday", Value:_contractday, Compare:"="})    
    }
    _billingdate := c.Geti("billingdate")
    if _billingdate != 0 {
        args = append(args, models.Where{Column:"billingdate", Value:_billingdate, Compare:"="})    
    }
    _billingtype := c.Geti("billingtype")
    if _billingtype != 0 {
        args = append(args, models.Where{Column:"billingtype", Value:_billingtype, Compare:"="})    
    }
    _billingname := c.Get("billingname")
    if _billingname != "" {
        args = append(args, models.Where{Column:"billingname", Value:_billingname, Compare:"like"})
    }
    _billingtel := c.Get("billingtel")
    if _billingtel != "" {
        args = append(args, models.Where{Column:"billingtel", Value:_billingtel, Compare:"like"})
    }
    _billingemail := c.Get("billingemail")
    if _billingemail != "" {
        args = append(args, models.Where{Column:"billingemail", Value:_billingemail, Compare:"like"})
    }
    _address := c.Get("address")
    if _address != "" {
        args = append(args, models.Where{Column:"address", Value:_address, Compare:"like"})
    }
    _addressetc := c.Get("addressetc")
    if _addressetc != "" {
        args = append(args, models.Where{Column:"addressetc", Value:_addressetc, Compare:"like"})
    }
    _collectmonth := c.Geti("collectmonth")
    if _collectmonth != 0 {
        args = append(args, models.Where{Column:"collectmonth", Value:_collectmonth, Compare:"="})    
    }
    _collectday := c.Geti("collectday")
    if _collectday != 0 {
        args = append(args, models.Where{Column:"collectday", Value:_collectday, Compare:"="})    
    }
    _manager := c.Get("manager")
    if _manager != "" {
        args = append(args, models.Where{Column:"manager", Value:_manager, Compare:"like"})
    }
    _tel := c.Get("tel")
    if _tel != "" {
        args = append(args, models.Where{Column:"tel", Value:_tel, Compare:"like"})
    }
    _fax := c.Get("fax")
    if _fax != "" {
        args = append(args, models.Where{Column:"fax", Value:_fax, Compare:"like"})
    }
    _status := c.Geti("status")
    if _status != 0 {
        args = append(args, models.Where{Column:"status", Value:_status, Compare:"="})    
    }
    _salesuser := c.Geti64("salesuser")
    if _salesuser != 0 {
        args = append(args, models.Where{Column:"salesuser", Value:_salesuser, Compare:"="})    
    }
    _user := c.Geti64("user")
    if _user != 0 {
        args = append(args, models.Where{Column:"user", Value:_user, Compare:"="})    
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
                    str += ", cu_" + strings.Trim(v, " ")                
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
func (c *CustomerController) UpdateNumber(number int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCustomerManager(conn)
	_manager.UpdateNumber(number, id)
}

// @Put()
func (c *CustomerController) UpdateType(typeid int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCustomerManager(conn)
	_manager.UpdateType(typeid, id)
}

// @Put()
func (c *CustomerController) UpdateCheckdate(checkdate int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCustomerManager(conn)
	_manager.UpdateCheckdate(checkdate, id)
}

// @Put()
func (c *CustomerController) UpdateManagername(managername string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCustomerManager(conn)
	_manager.UpdateManagername(managername, id)
}

// @Put()
func (c *CustomerController) UpdateManagertel(managertel string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCustomerManager(conn)
	_manager.UpdateManagertel(managertel, id)
}

// @Put()
func (c *CustomerController) UpdateManageremail(manageremail string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCustomerManager(conn)
	_manager.UpdateManageremail(manageremail, id)
}

// @Put()
func (c *CustomerController) UpdateContractstartdate(contractstartdate string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCustomerManager(conn)
	_manager.UpdateContractstartdate(contractstartdate, id)
}

// @Put()
func (c *CustomerController) UpdateContractenddate(contractenddate string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCustomerManager(conn)
	_manager.UpdateContractenddate(contractenddate, id)
}

// @Put()
func (c *CustomerController) UpdateContractprice(contractprice int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCustomerManager(conn)
	_manager.UpdateContractprice(contractprice, id)
}

// @Put()
func (c *CustomerController) UpdateContractvat(contractvat int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCustomerManager(conn)
	_manager.UpdateContractvat(contractvat, id)
}

// @Put()
func (c *CustomerController) UpdateContractday(contractday int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCustomerManager(conn)
	_manager.UpdateContractday(contractday, id)
}

// @Put()
func (c *CustomerController) UpdateBillingdate(billingdate int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCustomerManager(conn)
	_manager.UpdateBillingdate(billingdate, id)
}

// @Put()
func (c *CustomerController) UpdateBillingtype(billingtype int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCustomerManager(conn)
	_manager.UpdateBillingtype(billingtype, id)
}

// @Put()
func (c *CustomerController) UpdateBillingname(billingname string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCustomerManager(conn)
	_manager.UpdateBillingname(billingname, id)
}

// @Put()
func (c *CustomerController) UpdateBillingtel(billingtel string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCustomerManager(conn)
	_manager.UpdateBillingtel(billingtel, id)
}

// @Put()
func (c *CustomerController) UpdateBillingemail(billingemail string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCustomerManager(conn)
	_manager.UpdateBillingemail(billingemail, id)
}

// @Put()
func (c *CustomerController) UpdateAddress(address string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCustomerManager(conn)
	_manager.UpdateAddress(address, id)
}

// @Put()
func (c *CustomerController) UpdateAddressetc(addressetc string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCustomerManager(conn)
	_manager.UpdateAddressetc(addressetc, id)
}

// @Put()
func (c *CustomerController) UpdateCollectmonth(collectmonth int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCustomerManager(conn)
	_manager.UpdateCollectmonth(collectmonth, id)
}

// @Put()
func (c *CustomerController) UpdateCollectday(collectday int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCustomerManager(conn)
	_manager.UpdateCollectday(collectday, id)
}

// @Put()
func (c *CustomerController) UpdateManager(manager string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCustomerManager(conn)
	_manager.UpdateManager(manager, id)
}

// @Put()
func (c *CustomerController) UpdateTel(tel string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCustomerManager(conn)
	_manager.UpdateTel(tel, id)
}

// @Put()
func (c *CustomerController) UpdateFax(fax string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCustomerManager(conn)
	_manager.UpdateFax(fax, id)
}

// @Put()
func (c *CustomerController) UpdateStatus(status int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCustomerManager(conn)
	_manager.UpdateStatus(status, id)
}

// @Put()
func (c *CustomerController) UpdateSalesuser(salesuser int64, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCustomerManager(conn)
	_manager.UpdateSalesuser(salesuser, id)
}

// @Put()
func (c *CustomerController) UpdateUser(user int64, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCustomerManager(conn)
	_manager.UpdateUser(user, id)
}

// @Put()
func (c *CustomerController) UpdateCompany(company int64, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCustomerManager(conn)
	_manager.UpdateCompany(company, id)
}

// @Put()
func (c *CustomerController) UpdateBuilding(building int64, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCustomerManager(conn)
	_manager.UpdateBuilding(building, id)
}






