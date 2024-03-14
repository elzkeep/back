package rest


import (
	"zkeep/controllers"
	"zkeep/models"

    "strings"
)

type ReportController struct {
	controllers.Controller
}

func (c *ReportController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewReportManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *ReportController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewReportManager(conn)

    var args []interface{}
    
    _title := c.Get("title")
    if _title != "" {
        args = append(args, models.Where{Column:"title", Value:_title, Compare:"="})
        
    }
    _period := c.Geti("period")
    if _period != 0 {
        args = append(args, models.Where{Column:"period", Value:_period, Compare:"="})    
    }
    _number := c.Geti("number")
    if _number != 0 {
        args = append(args, models.Where{Column:"number", Value:_number, Compare:"="})    
    }
    _startcheckdate := c.Get("startcheckdate")
    _endcheckdate := c.Get("endcheckdate")
    if _startcheckdate != "" && _endcheckdate != "" {        
        var v [2]string
        v[0] = _startcheckdate
        v[1] = _endcheckdate  
        args = append(args, models.Where{Column:"checkdate", Value:v, Compare:"between"})    
    } else if  _startcheckdate != "" {          
        args = append(args, models.Where{Column:"checkdate", Value:_startcheckdate, Compare:">="})
    } else if  _endcheckdate != "" {          
        args = append(args, models.Where{Column:"checkdate", Value:_endcheckdate, Compare:"<="})            
    }
    _checktime := c.Get("checktime")
    if _checktime != "" {
        args = append(args, models.Where{Column:"checktime", Value:_checktime, Compare:"like"})
    }
    _content := c.Get("content")
    if _content != "" {
        args = append(args, models.Where{Column:"content", Value:_content, Compare:"="})
        
    }
    _image := c.Get("image")
    if _image != "" {
        args = append(args, models.Where{Column:"image", Value:_image, Compare:"like"})
    }
    _sign1 := c.Get("sign1")
    if _sign1 != "" {
        args = append(args, models.Where{Column:"sign1", Value:_sign1, Compare:"like"})
    }
    _sign2 := c.Get("sign2")
    if _sign2 != "" {
        args = append(args, models.Where{Column:"sign2", Value:_sign2, Compare:"like"})
    }
    _status := c.Geti("status")
    if _status != 0 {
        args = append(args, models.Where{Column:"status", Value:_status, Compare:"="})    
    }
    _company := c.Geti64("company")
    if _company != 0 {
        args = append(args, models.Where{Column:"company", Value:_company, Compare:"="})    
    }
    _user := c.Geti64("user")
    if _user != 0 {
        args = append(args, models.Where{Column:"user", Value:_user, Compare:"="})    
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
                    str += ", r_" + strings.Trim(v, " ")                
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

func (c *ReportController) Insert(item *models.Report) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewReportManager(conn)
	manager.Insert(item)

    id := manager.GetIdentity()
    c.Result["id"] = id
    item.Id = id
}

func (c *ReportController) Insertbatch(item *[]models.Report) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewReportManager(conn)

    for i := 0; i < rows; i++ {
	    manager.Insert(&((*item)[i]))
    }
}

func (c *ReportController) Update(item *models.Report) {
    
    
	conn := c.NewConnection()

	manager := models.NewReportManager(conn)
	manager.Update(item)
}

func (c *ReportController) Delete(item *models.Report) {
    
    
    conn := c.NewConnection()

	manager := models.NewReportManager(conn)

    
	manager.Delete(item.Id)
}

func (c *ReportController) Deletebatch(item *[]models.Report) {
    
    
    conn := c.NewConnection()

	manager := models.NewReportManager(conn)

    for _, v := range *item {
        
    
	    manager.Delete(v.Id)
    }
}



// @Put()
func (c *ReportController) UpdateTitle(title string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewReportManager(conn)
	_manager.UpdateTitle(title, id)
}

// @Put()
func (c *ReportController) UpdatePeriod(period int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewReportManager(conn)
	_manager.UpdatePeriod(period, id)
}

// @Put()
func (c *ReportController) UpdateNumber(number int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewReportManager(conn)
	_manager.UpdateNumber(number, id)
}

// @Put()
func (c *ReportController) UpdateCheckdate(checkdate string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewReportManager(conn)
	_manager.UpdateCheckdate(checkdate, id)
}

// @Put()
func (c *ReportController) UpdateChecktime(checktime string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewReportManager(conn)
	_manager.UpdateChecktime(checktime, id)
}

// @Put()
func (c *ReportController) UpdateContent(content string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewReportManager(conn)
	_manager.UpdateContent(content, id)
}

// @Put()
func (c *ReportController) UpdateImage(image string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewReportManager(conn)
	_manager.UpdateImage(image, id)
}

// @Put()
func (c *ReportController) UpdateSign1(sign1 string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewReportManager(conn)
	_manager.UpdateSign1(sign1, id)
}

// @Put()
func (c *ReportController) UpdateSign2(sign2 string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewReportManager(conn)
	_manager.UpdateSign2(sign2, id)
}

// @Put()
func (c *ReportController) UpdateStatus(status int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewReportManager(conn)
	_manager.UpdateStatus(status, id)
}

// @Put()
func (c *ReportController) UpdateCompany(company int64, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewReportManager(conn)
	_manager.UpdateCompany(company, id)
}

// @Put()
func (c *ReportController) UpdateUser(user int64, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewReportManager(conn)
	_manager.UpdateUser(user, id)
}

// @Put()
func (c *ReportController) UpdateBuilding(building int64, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewReportManager(conn)
	_manager.UpdateBuilding(building, id)
}






