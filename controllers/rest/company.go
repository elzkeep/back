package rest


import (
	"aoi/controllers"
	"aoi/models"

    "strings"
)

type CompanyController struct {
	controllers.Controller
}

func (c *CompanyController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewCompanyManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *CompanyController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewCompanyManager(conn)

    var args []interface{}
    
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"="})
        
    }
    _companyno := c.Get("companyno")
    if _companyno != "" {
        args = append(args, models.Where{Column:"companyno", Value:_companyno, Compare:"like"})
    }
    _ceo := c.Get("ceo")
    if _ceo != "" {
        args = append(args, models.Where{Column:"ceo", Value:_ceo, Compare:"like"})
    }
    _address := c.Get("address")
    if _address != "" {
        args = append(args, models.Where{Column:"address", Value:_address, Compare:"like"})
    }
    _addressetc := c.Get("addressetc")
    if _addressetc != "" {
        args = append(args, models.Where{Column:"addressetc", Value:_addressetc, Compare:"like"})
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
    _billingdate := c.Geti("billingdate")
    if _billingdate != 0 {
        args = append(args, models.Where{Column:"billingdate", Value:_billingdate, Compare:"="})    
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
    _status := c.Geti("status")
    if _status != 0 {
        args = append(args, models.Where{Column:"status", Value:_status, Compare:"="})    
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
                    str += ", c_" + strings.Trim(v, " ")                
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

func (c *CompanyController) Insert(item *models.Company) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewCompanyManager(conn)
	manager.Insert(item)

    id := manager.GetIdentity()
    c.Result["id"] = id
    item.Id = id
}

func (c *CompanyController) Insertbatch(item *[]models.Company) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewCompanyManager(conn)

    for i := 0; i < rows; i++ {
	    manager.Insert(&((*item)[i]))
    }
}

func (c *CompanyController) Update(item *models.Company) {
    
    
	conn := c.NewConnection()

	manager := models.NewCompanyManager(conn)
	manager.Update(item)
}

func (c *CompanyController) Delete(item *models.Company) {
    
    
    conn := c.NewConnection()

	manager := models.NewCompanyManager(conn)

    
	manager.Delete(item.Id)
}

func (c *CompanyController) Deletebatch(item *[]models.Company) {
    
    
    conn := c.NewConnection()

	manager := models.NewCompanyManager(conn)

    for _, v := range *item {
        
    
	    manager.Delete(v.Id)
    }
}



// @Put()
func (c *CompanyController) UpdateName(name string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanyManager(conn)
	_manager.UpdateName(name, id)
}

// @Put()
func (c *CompanyController) UpdateCompanyno(companyno string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanyManager(conn)
	_manager.UpdateCompanyno(companyno, id)
}

// @Put()
func (c *CompanyController) UpdateCeo(ceo string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanyManager(conn)
	_manager.UpdateCeo(ceo, id)
}

// @Put()
func (c *CompanyController) UpdateAddress(address string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanyManager(conn)
	_manager.UpdateAddress(address, id)
}

// @Put()
func (c *CompanyController) UpdateAddressetc(addressetc string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanyManager(conn)
	_manager.UpdateAddressetc(addressetc, id)
}

// @Put()
func (c *CompanyController) UpdateType(typeid int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanyManager(conn)
	_manager.UpdateType(typeid, id)
}

// @Put()
func (c *CompanyController) UpdateCheckdate(checkdate int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanyManager(conn)
	_manager.UpdateCheckdate(checkdate, id)
}

// @Put()
func (c *CompanyController) UpdateManagername(managername string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanyManager(conn)
	_manager.UpdateManagername(managername, id)
}

// @Put()
func (c *CompanyController) UpdateManagertel(managertel string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanyManager(conn)
	_manager.UpdateManagertel(managertel, id)
}

// @Put()
func (c *CompanyController) UpdateManageremail(manageremail string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanyManager(conn)
	_manager.UpdateManageremail(manageremail, id)
}

// @Put()
func (c *CompanyController) UpdateContractstartdate(contractstartdate string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanyManager(conn)
	_manager.UpdateContractstartdate(contractstartdate, id)
}

// @Put()
func (c *CompanyController) UpdateContractenddate(contractenddate string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanyManager(conn)
	_manager.UpdateContractenddate(contractenddate, id)
}

// @Put()
func (c *CompanyController) UpdateContractprice(contractprice int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanyManager(conn)
	_manager.UpdateContractprice(contractprice, id)
}

// @Put()
func (c *CompanyController) UpdateBillingdate(billingdate int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanyManager(conn)
	_manager.UpdateBillingdate(billingdate, id)
}

// @Put()
func (c *CompanyController) UpdateBillingname(billingname string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanyManager(conn)
	_manager.UpdateBillingname(billingname, id)
}

// @Put()
func (c *CompanyController) UpdateBillingtel(billingtel string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanyManager(conn)
	_manager.UpdateBillingtel(billingtel, id)
}

// @Put()
func (c *CompanyController) UpdateBillingemail(billingemail string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanyManager(conn)
	_manager.UpdateBillingemail(billingemail, id)
}

// @Put()
func (c *CompanyController) UpdateStatus(status int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanyManager(conn)
	_manager.UpdateStatus(status, id)
}






