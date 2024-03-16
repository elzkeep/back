package rest


import (
	"zkeep/controllers"
	"zkeep/models"

    "strings"
)

type DataController struct {
	controllers.Controller
}

func (c *DataController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewDataManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *DataController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewDataManager(conn)

    var args []interface{}
    
    _topcategory := c.Geti("topcategory")
    if _topcategory != 0 {
        args = append(args, models.Where{Column:"topcategory", Value:_topcategory, Compare:"="})    
    }
    _title := c.Get("title")
    if _title != "" {
        args = append(args, models.Where{Column:"title", Value:_title, Compare:"="})
        
    }
    _type := c.Geti("type")
    if _type != 0 {
        args = append(args, models.Where{Column:"type", Value:_type, Compare:"="})    
    }
    _category := c.Geti("category")
    if _category != 0 {
        args = append(args, models.Where{Column:"category", Value:_category, Compare:"="})    
    }
    _order := c.Geti("order")
    if _order != 0 {
        args = append(args, models.Where{Column:"order", Value:_order, Compare:"="})    
    }
    _report := c.Geti64("report")
    if _report != 0 {
        args = append(args, models.Where{Column:"report", Value:_report, Compare:"="})    
    }
    _company := c.Geti64("company")
    if _company != 0 {
        args = append(args, models.Where{Column:"company", Value:_company, Compare:"="})    
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
                    str += ", d_" + strings.Trim(v, " ")                
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

func (c *DataController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewDataManager(conn)

    var args []interface{}
    
    _topcategory := c.Geti("topcategory")
    if _topcategory != 0 {
        args = append(args, models.Where{Column:"topcategory", Value:_topcategory, Compare:"="})    
    }
    _title := c.Get("title")
    if _title != "" {
        args = append(args, models.Where{Column:"title", Value:_title, Compare:"="})
        
    }
    _type := c.Geti("type")
    if _type != 0 {
        args = append(args, models.Where{Column:"type", Value:_type, Compare:"="})    
    }
    _category := c.Geti("category")
    if _category != 0 {
        args = append(args, models.Where{Column:"category", Value:_category, Compare:"="})    
    }
    _order := c.Geti("order")
    if _order != 0 {
        args = append(args, models.Where{Column:"order", Value:_order, Compare:"="})    
    }
    _report := c.Geti64("report")
    if _report != 0 {
        args = append(args, models.Where{Column:"report", Value:_report, Compare:"="})    
    }
    _company := c.Geti64("company")
    if _company != 0 {
        args = append(args, models.Where{Column:"company", Value:_company, Compare:"="})    
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

func (c *DataController) Insert(item *models.Data) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewDataManager(conn)
	manager.Insert(item)

    id := manager.GetIdentity()
    c.Result["id"] = id
    item.Id = id
}

func (c *DataController) Insertbatch(item *[]models.Data) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewDataManager(conn)

    for i := 0; i < rows; i++ {
	    manager.Insert(&((*item)[i]))
    }
}

func (c *DataController) Update(item *models.Data) {
    
    
	conn := c.NewConnection()

	manager := models.NewDataManager(conn)
	manager.Update(item)
}

func (c *DataController) Delete(item *models.Data) {
    
    
    conn := c.NewConnection()

	manager := models.NewDataManager(conn)

    
	manager.Delete(item.Id)
}

func (c *DataController) Deletebatch(item *[]models.Data) {
    
    
    conn := c.NewConnection()

	manager := models.NewDataManager(conn)

    for _, v := range *item {
        
    
	    manager.Delete(v.Id)
    }
}


// @Delete()
func (c *DataController) DeleteByReportTopcategory(report int64 ,topcategory int) {
    
    conn := c.NewConnection()

	_manager := models.NewDataManager(conn)
    
    _manager.DeleteByReportTopcategory(report, topcategory)
    
}


// @Put()
func (c *DataController) UpdateTopcategory(topcategory int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewDataManager(conn)
	_manager.UpdateTopcategory(topcategory, id)
}

// @Put()
func (c *DataController) UpdateTitle(title string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewDataManager(conn)
	_manager.UpdateTitle(title, id)
}

// @Put()
func (c *DataController) UpdateType(typeid int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewDataManager(conn)
	_manager.UpdateType(typeid, id)
}

// @Put()
func (c *DataController) UpdateCategory(category int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewDataManager(conn)
	_manager.UpdateCategory(category, id)
}

// @Put()
func (c *DataController) UpdateOrder(order int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewDataManager(conn)
	_manager.UpdateOrder(order, id)
}

// @Put()
func (c *DataController) UpdateReport(report int64, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewDataManager(conn)
	_manager.UpdateReport(report, id)
}

// @Put()
func (c *DataController) UpdateCompany(company int64, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewDataManager(conn)
	_manager.UpdateCompany(company, id)
}






