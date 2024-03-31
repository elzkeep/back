package rest


import (
	"zkeep/controllers"
	"zkeep/models"

    "strings"
)

type CompanylicenseController struct {
	controllers.Controller
}

func (c *CompanylicenseController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewCompanylicenseManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *CompanylicenseController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewCompanylicenseManager(conn)

    var args []interface{}
    
    _number := c.Get("number")
    if _number != "" {
        args = append(args, models.Where{Column:"number", Value:_number, Compare:"like"})
    }
    _starttakingdate := c.Get("starttakingdate")
    _endtakingdate := c.Get("endtakingdate")
    if _starttakingdate != "" && _endtakingdate != "" {        
        var v [2]string
        v[0] = _starttakingdate
        v[1] = _endtakingdate  
        args = append(args, models.Where{Column:"takingdate", Value:v, Compare:"between"})    
    } else if  _starttakingdate != "" {          
        args = append(args, models.Where{Column:"takingdate", Value:_starttakingdate, Compare:">="})
    } else if  _endtakingdate != "" {          
        args = append(args, models.Where{Column:"takingdate", Value:_endtakingdate, Compare:"<="})            
    }
    _starteducationdate := c.Get("starteducationdate")
    _endeducationdate := c.Get("endeducationdate")
    if _starteducationdate != "" && _endeducationdate != "" {        
        var v [2]string
        v[0] = _starteducationdate
        v[1] = _endeducationdate  
        args = append(args, models.Where{Column:"educationdate", Value:v, Compare:"between"})    
    } else if  _starteducationdate != "" {          
        args = append(args, models.Where{Column:"educationdate", Value:_starteducationdate, Compare:">="})
    } else if  _endeducationdate != "" {          
        args = append(args, models.Where{Column:"educationdate", Value:_endeducationdate, Compare:"<="})            
    }
    _educationinstitution := c.Get("educationinstitution")
    if _educationinstitution != "" {
        args = append(args, models.Where{Column:"educationinstitution", Value:_educationinstitution, Compare:"like"})
    }
    _startspecialeducationdate := c.Get("startspecialeducationdate")
    _endspecialeducationdate := c.Get("endspecialeducationdate")
    if _startspecialeducationdate != "" && _endspecialeducationdate != "" {        
        var v [2]string
        v[0] = _startspecialeducationdate
        v[1] = _endspecialeducationdate  
        args = append(args, models.Where{Column:"specialeducationdate", Value:v, Compare:"between"})    
    } else if  _startspecialeducationdate != "" {          
        args = append(args, models.Where{Column:"specialeducationdate", Value:_startspecialeducationdate, Compare:">="})
    } else if  _endspecialeducationdate != "" {          
        args = append(args, models.Where{Column:"specialeducationdate", Value:_endspecialeducationdate, Compare:"<="})            
    }
    _specialeducationinstitution := c.Get("specialeducationinstitution")
    if _specialeducationinstitution != "" {
        args = append(args, models.Where{Column:"specialeducationinstitution", Value:_specialeducationinstitution, Compare:"like"})
    }
    _company := c.Geti64("company")
    if _company != 0 {
        args = append(args, models.Where{Column:"company", Value:_company, Compare:"="})    
    }
    _licensecategory := c.Geti64("licensecategory")
    if _licensecategory != 0 {
        args = append(args, models.Where{Column:"licensecategory", Value:_licensecategory, Compare:"="})    
    }
    _licenselevel := c.Geti64("licenselevel")
    if _licenselevel != 0 {
        args = append(args, models.Where{Column:"licenselevel", Value:_licenselevel, Compare:"="})    
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
                    str += ", l_" + strings.Trim(v, " ")                
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

func (c *CompanylicenseController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewCompanylicenseManager(conn)

    var args []interface{}
    
    _number := c.Get("number")
    if _number != "" {
        args = append(args, models.Where{Column:"number", Value:_number, Compare:"like"})
    }
    _starttakingdate := c.Get("starttakingdate")
    _endtakingdate := c.Get("endtakingdate")
    if _starttakingdate != "" && _endtakingdate != "" {        
        var v [2]string
        v[0] = _starttakingdate
        v[1] = _endtakingdate  
        args = append(args, models.Where{Column:"takingdate", Value:v, Compare:"between"})    
    } else if  _starttakingdate != "" {          
        args = append(args, models.Where{Column:"takingdate", Value:_starttakingdate, Compare:">="})
    } else if  _endtakingdate != "" {          
        args = append(args, models.Where{Column:"takingdate", Value:_endtakingdate, Compare:"<="})            
    }
    _starteducationdate := c.Get("starteducationdate")
    _endeducationdate := c.Get("endeducationdate")
    if _starteducationdate != "" && _endeducationdate != "" {        
        var v [2]string
        v[0] = _starteducationdate
        v[1] = _endeducationdate  
        args = append(args, models.Where{Column:"educationdate", Value:v, Compare:"between"})    
    } else if  _starteducationdate != "" {          
        args = append(args, models.Where{Column:"educationdate", Value:_starteducationdate, Compare:">="})
    } else if  _endeducationdate != "" {          
        args = append(args, models.Where{Column:"educationdate", Value:_endeducationdate, Compare:"<="})            
    }
    _educationinstitution := c.Get("educationinstitution")
    if _educationinstitution != "" {
        args = append(args, models.Where{Column:"educationinstitution", Value:_educationinstitution, Compare:"like"})
    }
    _startspecialeducationdate := c.Get("startspecialeducationdate")
    _endspecialeducationdate := c.Get("endspecialeducationdate")
    if _startspecialeducationdate != "" && _endspecialeducationdate != "" {        
        var v [2]string
        v[0] = _startspecialeducationdate
        v[1] = _endspecialeducationdate  
        args = append(args, models.Where{Column:"specialeducationdate", Value:v, Compare:"between"})    
    } else if  _startspecialeducationdate != "" {          
        args = append(args, models.Where{Column:"specialeducationdate", Value:_startspecialeducationdate, Compare:">="})
    } else if  _endspecialeducationdate != "" {          
        args = append(args, models.Where{Column:"specialeducationdate", Value:_endspecialeducationdate, Compare:"<="})            
    }
    _specialeducationinstitution := c.Get("specialeducationinstitution")
    if _specialeducationinstitution != "" {
        args = append(args, models.Where{Column:"specialeducationinstitution", Value:_specialeducationinstitution, Compare:"like"})
    }
    _company := c.Geti64("company")
    if _company != 0 {
        args = append(args, models.Where{Column:"company", Value:_company, Compare:"="})    
    }
    _licensecategory := c.Geti64("licensecategory")
    if _licensecategory != 0 {
        args = append(args, models.Where{Column:"licensecategory", Value:_licensecategory, Compare:"="})    
    }
    _licenselevel := c.Geti64("licenselevel")
    if _licenselevel != 0 {
        args = append(args, models.Where{Column:"licenselevel", Value:_licenselevel, Compare:"="})    
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

func (c *CompanylicenseController) Insert(item *models.Companylicense) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewCompanylicenseManager(conn)
	manager.Insert(item)

    id := manager.GetIdentity()
    c.Result["id"] = id
    item.Id = id
}

func (c *CompanylicenseController) Insertbatch(item *[]models.Companylicense) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewCompanylicenseManager(conn)

    for i := 0; i < rows; i++ {
	    manager.Insert(&((*item)[i]))
    }
}

func (c *CompanylicenseController) Update(item *models.Companylicense) {
    
    
	conn := c.NewConnection()

	manager := models.NewCompanylicenseManager(conn)
	manager.Update(item)
}

func (c *CompanylicenseController) Delete(item *models.Companylicense) {
    
    
    conn := c.NewConnection()

	manager := models.NewCompanylicenseManager(conn)

    
	manager.Delete(item.Id)
}

func (c *CompanylicenseController) Deletebatch(item *[]models.Companylicense) {
    
    
    conn := c.NewConnection()

	manager := models.NewCompanylicenseManager(conn)

    for _, v := range *item {
        
    
	    manager.Delete(v.Id)
    }
}



// @Put()
func (c *CompanylicenseController) UpdateNumber(number string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanylicenseManager(conn)
	_manager.UpdateNumber(number, id)
}

// @Put()
func (c *CompanylicenseController) UpdateTakingdate(takingdate string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanylicenseManager(conn)
	_manager.UpdateTakingdate(takingdate, id)
}

// @Put()
func (c *CompanylicenseController) UpdateEducationdate(educationdate string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanylicenseManager(conn)
	_manager.UpdateEducationdate(educationdate, id)
}

// @Put()
func (c *CompanylicenseController) UpdateEducationinstitution(educationinstitution string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanylicenseManager(conn)
	_manager.UpdateEducationinstitution(educationinstitution, id)
}

// @Put()
func (c *CompanylicenseController) UpdateSpecialeducationdate(specialeducationdate string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanylicenseManager(conn)
	_manager.UpdateSpecialeducationdate(specialeducationdate, id)
}

// @Put()
func (c *CompanylicenseController) UpdateSpecialeducationinstitution(specialeducationinstitution string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanylicenseManager(conn)
	_manager.UpdateSpecialeducationinstitution(specialeducationinstitution, id)
}

// @Put()
func (c *CompanylicenseController) UpdateCompany(company int64, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanylicenseManager(conn)
	_manager.UpdateCompany(company, id)
}

// @Put()
func (c *CompanylicenseController) UpdateLicensecategory(licensecategory int64, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanylicenseManager(conn)
	_manager.UpdateLicensecategory(licensecategory, id)
}

// @Put()
func (c *CompanylicenseController) UpdateLicenselevel(licenselevel int64, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewCompanylicenseManager(conn)
	_manager.UpdateLicenselevel(licenselevel, id)
}






