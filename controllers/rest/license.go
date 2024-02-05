package rest


import (
	"zkeep/controllers"
	"zkeep/models"

    "strings"
)

type LicenseController struct {
	controllers.Controller
}

func (c *LicenseController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewLicenseManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *LicenseController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewLicenseManager(conn)

    var args []interface{}
    
    _user := c.Geti64("user")
    if _user != 0 {
        args = append(args, models.Where{Column:"user", Value:_user, Compare:"="})    
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

func (c *LicenseController) Insert(item *models.License) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewLicenseManager(conn)
	manager.Insert(item)

    id := manager.GetIdentity()
    c.Result["id"] = id
    item.Id = id
}

func (c *LicenseController) Insertbatch(item *[]models.License) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewLicenseManager(conn)

    for i := 0; i < rows; i++ {
	    manager.Insert(&((*item)[i]))
    }
}

func (c *LicenseController) Update(item *models.License) {
    
    
	conn := c.NewConnection()

	manager := models.NewLicenseManager(conn)
	manager.Update(item)
}

func (c *LicenseController) Delete(item *models.License) {
    
    
    conn := c.NewConnection()

	manager := models.NewLicenseManager(conn)

    
	manager.Delete(item.Id)
}

func (c *LicenseController) Deletebatch(item *[]models.License) {
    
    
    conn := c.NewConnection()

	manager := models.NewLicenseManager(conn)

    for _, v := range *item {
        
    
	    manager.Delete(v.Id)
    }
}



// @Put()
func (c *LicenseController) UpdateUser(user int64, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewLicenseManager(conn)
	_manager.UpdateUser(user, id)
}

// @Put()
func (c *LicenseController) UpdateLicensecategory(licensecategory int64, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewLicenseManager(conn)
	_manager.UpdateLicensecategory(licensecategory, id)
}

// @Put()
func (c *LicenseController) UpdateLicenselevel(licenselevel int64, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewLicenseManager(conn)
	_manager.UpdateLicenselevel(licenselevel, id)
}






