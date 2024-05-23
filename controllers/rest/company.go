package rest


import (
	"zkeep/controllers"
	"zkeep/models"

    "strings"
)

type CompanyController struct {
	controllers.Controller
}



func (c *CompanyController) GetByCompanyno(companyno string) *models.Company {
    
    conn := c.NewConnection()

	_manager := models.NewCompanyManager(conn)
    
    item := _manager.GetByCompanyno(companyno)
    
    c.Set("item", item)
    
    
    
    return item
    
}


func (c *CompanyController) GetByName(name string) *models.Company {
    
    conn := c.NewConnection()

	_manager := models.NewCompanyManager(conn)
    
    item := _manager.GetByName(name)
    
    c.Set("item", item)
    
    
    
    return item
    
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

func (c *CompanyController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewCompanyManager(conn)

    var args []interface{}
    
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"like"})
        
    }
    _companyno := c.Get("companyno")
    if _companyno != "" {
        args = append(args, models.Where{Column:"companyno", Value:_companyno, Compare:"="})
    }
    _ceo := c.Get("ceo")
    if _ceo != "" {
        args = append(args, models.Where{Column:"ceo", Value:_ceo, Compare:"="})
    }
    _tel := c.Get("tel")
    if _tel != "" {
        args = append(args, models.Where{Column:"tel", Value:_tel, Compare:"="})
    }
    _email := c.Get("email")
    if _email != "" {
        args = append(args, models.Where{Column:"email", Value:_email, Compare:"="})
    }
    _address := c.Get("address")
    if _address != "" {
        args = append(args, models.Where{Column:"address", Value:_address, Compare:"="})
    }
    _addressetc := c.Get("addressetc")
    if _addressetc != "" {
        args = append(args, models.Where{Column:"addressetc", Value:_addressetc, Compare:"="})
    }
    _type := c.Geti("type")
    if _type != 0 {
        args = append(args, models.Where{Column:"type", Value:_type, Compare:"="})    
    }
    _billingname := c.Get("billingname")
    if _billingname != "" {
        args = append(args, models.Where{Column:"billingname", Value:_billingname, Compare:"="})
    }
    _billingtel := c.Get("billingtel")
    if _billingtel != "" {
        args = append(args, models.Where{Column:"billingtel", Value:_billingtel, Compare:"="})
    }
    _billingemail := c.Get("billingemail")
    if _billingemail != "" {
        args = append(args, models.Where{Column:"billingemail", Value:_billingemail, Compare:"="})
    }
    _bankname := c.Get("bankname")
    if _bankname != "" {
        args = append(args, models.Where{Column:"bankname", Value:_bankname, Compare:"="})
    }
    _bankno := c.Get("bankno")
    if _bankno != "" {
        args = append(args, models.Where{Column:"bankno", Value:_bankno, Compare:"="})
    }
    _businesscondition := c.Get("businesscondition")
    if _businesscondition != "" {
        args = append(args, models.Where{Column:"businesscondition", Value:_businesscondition, Compare:"="})
    }
    _businessitem := c.Get("businessitem")
    if _businessitem != "" {
        args = append(args, models.Where{Column:"businessitem", Value:_businessitem, Compare:"="})
    }
    _giro := c.Get("giro")
    if _giro != "" {
        args = append(args, models.Where{Column:"giro", Value:_giro, Compare:"="})
    }
    _content := c.Get("content")
    if _content != "" {
        args = append(args, models.Where{Column:"content", Value:_content, Compare:"="})
        
    }
    _x1 := c.Geti("x1")
    if _x1 != 0 {
        args = append(args, models.Where{Column:"x1", Value:_x1, Compare:"="})    
    }
    _y1 := c.Geti("y1")
    if _y1 != 0 {
        args = append(args, models.Where{Column:"y1", Value:_y1, Compare:"="})    
    }
    _x2 := c.Geti("x2")
    if _x2 != 0 {
        args = append(args, models.Where{Column:"x2", Value:_x2, Compare:"="})    
    }
    _y2 := c.Geti("y2")
    if _y2 != 0 {
        args = append(args, models.Where{Column:"y2", Value:_y2, Compare:"="})    
    }
    _x3 := c.Geti("x3")
    if _x3 != 0 {
        args = append(args, models.Where{Column:"x3", Value:_x3, Compare:"="})    
    }
    _y3 := c.Geti("y3")
    if _y3 != 0 {
        args = append(args, models.Where{Column:"y3", Value:_y3, Compare:"="})    
    }
    _x4 := c.Geti("x4")
    if _x4 != 0 {
        args = append(args, models.Where{Column:"x4", Value:_x4, Compare:"="})    
    }
    _y4 := c.Geti("y4")
    if _y4 != 0 {
        args = append(args, models.Where{Column:"y4", Value:_y4, Compare:"="})    
    }
    _x5 := c.Geti("x5")
    if _x5 != 0 {
        args = append(args, models.Where{Column:"x5", Value:_x5, Compare:"="})    
    }
    _y5 := c.Geti("y5")
    if _y5 != 0 {
        args = append(args, models.Where{Column:"y5", Value:_y5, Compare:"="})    
    }
    _x6 := c.Geti("x6")
    if _x6 != 0 {
        args = append(args, models.Where{Column:"x6", Value:_x6, Compare:"="})    
    }
    _y6 := c.Geti("y6")
    if _y6 != 0 {
        args = append(args, models.Where{Column:"y6", Value:_y6, Compare:"="})    
    }
    _x7 := c.Geti("x7")
    if _x7 != 0 {
        args = append(args, models.Where{Column:"x7", Value:_x7, Compare:"="})    
    }
    _y7 := c.Geti("y7")
    if _y7 != 0 {
        args = append(args, models.Where{Column:"y7", Value:_y7, Compare:"="})    
    }
    _x8 := c.Geti("x8")
    if _x8 != 0 {
        args = append(args, models.Where{Column:"x8", Value:_x8, Compare:"="})    
    }
    _y8 := c.Geti("y8")
    if _y8 != 0 {
        args = append(args, models.Where{Column:"y8", Value:_y8, Compare:"="})    
    }
    _x9 := c.Geti("x9")
    if _x9 != 0 {
        args = append(args, models.Where{Column:"x9", Value:_x9, Compare:"="})    
    }
    _y9 := c.Geti("y9")
    if _y9 != 0 {
        args = append(args, models.Where{Column:"y9", Value:_y9, Compare:"="})    
    }
    _x10 := c.Geti("x10")
    if _x10 != 0 {
        args = append(args, models.Where{Column:"x10", Value:_x10, Compare:"="})    
    }
    _y10 := c.Geti("y10")
    if _y10 != 0 {
        args = append(args, models.Where{Column:"y10", Value:_y10, Compare:"="})    
    }
    _x11 := c.Geti("x11")
    if _x11 != 0 {
        args = append(args, models.Where{Column:"x11", Value:_x11, Compare:"="})    
    }
    _y11 := c.Geti("y11")
    if _y11 != 0 {
        args = append(args, models.Where{Column:"y11", Value:_y11, Compare:"="})    
    }
    _x12 := c.Geti("x12")
    if _x12 != 0 {
        args = append(args, models.Where{Column:"x12", Value:_x12, Compare:"="})    
    }
    _y12 := c.Geti("y12")
    if _y12 != 0 {
        args = append(args, models.Where{Column:"y12", Value:_y12, Compare:"="})    
    }
    _x13 := c.Geti("x13")
    if _x13 != 0 {
        args = append(args, models.Where{Column:"x13", Value:_x13, Compare:"="})    
    }
    _y13 := c.Geti("y13")
    if _y13 != 0 {
        args = append(args, models.Where{Column:"y13", Value:_y13, Compare:"="})    
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
    

    
    
    total := manager.Count(args)
	c.Set("total", total)
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
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"like"})
        
    }
    _companyno := c.Get("companyno")
    if _companyno != "" {
        args = append(args, models.Where{Column:"companyno", Value:_companyno, Compare:"="})
    }
    _ceo := c.Get("ceo")
    if _ceo != "" {
        args = append(args, models.Where{Column:"ceo", Value:_ceo, Compare:"="})
    }
    _tel := c.Get("tel")
    if _tel != "" {
        args = append(args, models.Where{Column:"tel", Value:_tel, Compare:"="})
    }
    _email := c.Get("email")
    if _email != "" {
        args = append(args, models.Where{Column:"email", Value:_email, Compare:"="})
    }
    _address := c.Get("address")
    if _address != "" {
        args = append(args, models.Where{Column:"address", Value:_address, Compare:"="})
    }
    _addressetc := c.Get("addressetc")
    if _addressetc != "" {
        args = append(args, models.Where{Column:"addressetc", Value:_addressetc, Compare:"="})
    }
    _type := c.Geti("type")
    if _type != 0 {
        args = append(args, models.Where{Column:"type", Value:_type, Compare:"="})    
    }
    _billingname := c.Get("billingname")
    if _billingname != "" {
        args = append(args, models.Where{Column:"billingname", Value:_billingname, Compare:"="})
    }
    _billingtel := c.Get("billingtel")
    if _billingtel != "" {
        args = append(args, models.Where{Column:"billingtel", Value:_billingtel, Compare:"="})
    }
    _billingemail := c.Get("billingemail")
    if _billingemail != "" {
        args = append(args, models.Where{Column:"billingemail", Value:_billingemail, Compare:"="})
    }
    _bankname := c.Get("bankname")
    if _bankname != "" {
        args = append(args, models.Where{Column:"bankname", Value:_bankname, Compare:"="})
    }
    _bankno := c.Get("bankno")
    if _bankno != "" {
        args = append(args, models.Where{Column:"bankno", Value:_bankno, Compare:"="})
    }
    _businesscondition := c.Get("businesscondition")
    if _businesscondition != "" {
        args = append(args, models.Where{Column:"businesscondition", Value:_businesscondition, Compare:"="})
    }
    _businessitem := c.Get("businessitem")
    if _businessitem != "" {
        args = append(args, models.Where{Column:"businessitem", Value:_businessitem, Compare:"="})
    }
    _giro := c.Get("giro")
    if _giro != "" {
        args = append(args, models.Where{Column:"giro", Value:_giro, Compare:"="})
    }
    _content := c.Get("content")
    if _content != "" {
        args = append(args, models.Where{Column:"content", Value:_content, Compare:"="})
        
    }
    _x1 := c.Geti("x1")
    if _x1 != 0 {
        args = append(args, models.Where{Column:"x1", Value:_x1, Compare:"="})    
    }
    _y1 := c.Geti("y1")
    if _y1 != 0 {
        args = append(args, models.Where{Column:"y1", Value:_y1, Compare:"="})    
    }
    _x2 := c.Geti("x2")
    if _x2 != 0 {
        args = append(args, models.Where{Column:"x2", Value:_x2, Compare:"="})    
    }
    _y2 := c.Geti("y2")
    if _y2 != 0 {
        args = append(args, models.Where{Column:"y2", Value:_y2, Compare:"="})    
    }
    _x3 := c.Geti("x3")
    if _x3 != 0 {
        args = append(args, models.Where{Column:"x3", Value:_x3, Compare:"="})    
    }
    _y3 := c.Geti("y3")
    if _y3 != 0 {
        args = append(args, models.Where{Column:"y3", Value:_y3, Compare:"="})    
    }
    _x4 := c.Geti("x4")
    if _x4 != 0 {
        args = append(args, models.Where{Column:"x4", Value:_x4, Compare:"="})    
    }
    _y4 := c.Geti("y4")
    if _y4 != 0 {
        args = append(args, models.Where{Column:"y4", Value:_y4, Compare:"="})    
    }
    _x5 := c.Geti("x5")
    if _x5 != 0 {
        args = append(args, models.Where{Column:"x5", Value:_x5, Compare:"="})    
    }
    _y5 := c.Geti("y5")
    if _y5 != 0 {
        args = append(args, models.Where{Column:"y5", Value:_y5, Compare:"="})    
    }
    _x6 := c.Geti("x6")
    if _x6 != 0 {
        args = append(args, models.Where{Column:"x6", Value:_x6, Compare:"="})    
    }
    _y6 := c.Geti("y6")
    if _y6 != 0 {
        args = append(args, models.Where{Column:"y6", Value:_y6, Compare:"="})    
    }
    _x7 := c.Geti("x7")
    if _x7 != 0 {
        args = append(args, models.Where{Column:"x7", Value:_x7, Compare:"="})    
    }
    _y7 := c.Geti("y7")
    if _y7 != 0 {
        args = append(args, models.Where{Column:"y7", Value:_y7, Compare:"="})    
    }
    _x8 := c.Geti("x8")
    if _x8 != 0 {
        args = append(args, models.Where{Column:"x8", Value:_x8, Compare:"="})    
    }
    _y8 := c.Geti("y8")
    if _y8 != 0 {
        args = append(args, models.Where{Column:"y8", Value:_y8, Compare:"="})    
    }
    _x9 := c.Geti("x9")
    if _x9 != 0 {
        args = append(args, models.Where{Column:"x9", Value:_x9, Compare:"="})    
    }
    _y9 := c.Geti("y9")
    if _y9 != 0 {
        args = append(args, models.Where{Column:"y9", Value:_y9, Compare:"="})    
    }
    _x10 := c.Geti("x10")
    if _x10 != 0 {
        args = append(args, models.Where{Column:"x10", Value:_x10, Compare:"="})    
    }
    _y10 := c.Geti("y10")
    if _y10 != 0 {
        args = append(args, models.Where{Column:"y10", Value:_y10, Compare:"="})    
    }
    _x11 := c.Geti("x11")
    if _x11 != 0 {
        args = append(args, models.Where{Column:"x11", Value:_x11, Compare:"="})    
    }
    _y11 := c.Geti("y11")
    if _y11 != 0 {
        args = append(args, models.Where{Column:"y11", Value:_y11, Compare:"="})    
    }
    _x12 := c.Geti("x12")
    if _x12 != 0 {
        args = append(args, models.Where{Column:"x12", Value:_x12, Compare:"="})    
    }
    _y12 := c.Geti("y12")
    if _y12 != 0 {
        args = append(args, models.Where{Column:"y12", Value:_y12, Compare:"="})    
    }
    _x13 := c.Geti("x13")
    if _x13 != 0 {
        args = append(args, models.Where{Column:"x13", Value:_x13, Compare:"="})    
    }
    _y13 := c.Geti("y13")
    if _y13 != 0 {
        args = append(args, models.Where{Column:"y13", Value:_y13, Compare:"="})    
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
func (c *CompanyController) UpdateTel(tel string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanyManager(conn)
	_manager.UpdateTel(tel, id)
}

// @Put()
func (c *CompanyController) UpdateEmail(email string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanyManager(conn)
	_manager.UpdateEmail(email, id)
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
func (c *CompanyController) UpdateBankname(bankname string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanyManager(conn)
	_manager.UpdateBankname(bankname, id)
}

// @Put()
func (c *CompanyController) UpdateBankno(bankno string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanyManager(conn)
	_manager.UpdateBankno(bankno, id)
}

// @Put()
func (c *CompanyController) UpdateBusinesscondition(businesscondition string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanyManager(conn)
	_manager.UpdateBusinesscondition(businesscondition, id)
}

// @Put()
func (c *CompanyController) UpdateBusinessitem(businessitem string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanyManager(conn)
	_manager.UpdateBusinessitem(businessitem, id)
}

// @Put()
func (c *CompanyController) UpdateGiro(giro string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanyManager(conn)
	_manager.UpdateGiro(giro, id)
}

// @Put()
func (c *CompanyController) UpdateContent(content string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanyManager(conn)
	_manager.UpdateContent(content, id)
}

// @Put()
func (c *CompanyController) UpdateX1(x1 models.Double, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanyManager(conn)
	_manager.UpdateX1(x1, id)
}

// @Put()
func (c *CompanyController) UpdateY1(y1 models.Double, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanyManager(conn)
	_manager.UpdateY1(y1, id)
}

// @Put()
func (c *CompanyController) UpdateX2(x2 models.Double, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanyManager(conn)
	_manager.UpdateX2(x2, id)
}

// @Put()
func (c *CompanyController) UpdateY2(y2 models.Double, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanyManager(conn)
	_manager.UpdateY2(y2, id)
}

// @Put()
func (c *CompanyController) UpdateX3(x3 models.Double, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanyManager(conn)
	_manager.UpdateX3(x3, id)
}

// @Put()
func (c *CompanyController) UpdateY3(y3 models.Double, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanyManager(conn)
	_manager.UpdateY3(y3, id)
}

// @Put()
func (c *CompanyController) UpdateX4(x4 models.Double, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanyManager(conn)
	_manager.UpdateX4(x4, id)
}

// @Put()
func (c *CompanyController) UpdateY4(y4 models.Double, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanyManager(conn)
	_manager.UpdateY4(y4, id)
}

// @Put()
func (c *CompanyController) UpdateX5(x5 models.Double, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanyManager(conn)
	_manager.UpdateX5(x5, id)
}

// @Put()
func (c *CompanyController) UpdateY5(y5 models.Double, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanyManager(conn)
	_manager.UpdateY5(y5, id)
}

// @Put()
func (c *CompanyController) UpdateX6(x6 models.Double, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanyManager(conn)
	_manager.UpdateX6(x6, id)
}

// @Put()
func (c *CompanyController) UpdateY6(y6 models.Double, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanyManager(conn)
	_manager.UpdateY6(y6, id)
}

// @Put()
func (c *CompanyController) UpdateX7(x7 models.Double, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanyManager(conn)
	_manager.UpdateX7(x7, id)
}

// @Put()
func (c *CompanyController) UpdateY7(y7 models.Double, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanyManager(conn)
	_manager.UpdateY7(y7, id)
}

// @Put()
func (c *CompanyController) UpdateX8(x8 models.Double, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanyManager(conn)
	_manager.UpdateX8(x8, id)
}

// @Put()
func (c *CompanyController) UpdateY8(y8 models.Double, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanyManager(conn)
	_manager.UpdateY8(y8, id)
}

// @Put()
func (c *CompanyController) UpdateX9(x9 models.Double, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanyManager(conn)
	_manager.UpdateX9(x9, id)
}

// @Put()
func (c *CompanyController) UpdateY9(y9 models.Double, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanyManager(conn)
	_manager.UpdateY9(y9, id)
}

// @Put()
func (c *CompanyController) UpdateX10(x10 models.Double, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanyManager(conn)
	_manager.UpdateX10(x10, id)
}

// @Put()
func (c *CompanyController) UpdateY10(y10 models.Double, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanyManager(conn)
	_manager.UpdateY10(y10, id)
}

// @Put()
func (c *CompanyController) UpdateX11(x11 models.Double, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanyManager(conn)
	_manager.UpdateX11(x11, id)
}

// @Put()
func (c *CompanyController) UpdateY11(y11 models.Double, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanyManager(conn)
	_manager.UpdateY11(y11, id)
}

// @Put()
func (c *CompanyController) UpdateX12(x12 models.Double, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanyManager(conn)
	_manager.UpdateX12(x12, id)
}

// @Put()
func (c *CompanyController) UpdateY12(y12 models.Double, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanyManager(conn)
	_manager.UpdateY12(y12, id)
}

// @Put()
func (c *CompanyController) UpdateX13(x13 models.Double, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanyManager(conn)
	_manager.UpdateX13(x13, id)
}

// @Put()
func (c *CompanyController) UpdateY13(y13 models.Double, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanyManager(conn)
	_manager.UpdateY13(y13, id)
}

// @Put()
func (c *CompanyController) UpdateStatus(status int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanyManager(conn)
	_manager.UpdateStatus(status, id)
}






