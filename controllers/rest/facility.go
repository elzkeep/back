package rest


import (
	"zkeep/controllers"
	"zkeep/models"

    "strings"
)

type FacilityController struct {
	controllers.Controller
}

func (c *FacilityController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewFacilityManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *FacilityController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewFacilityManager(conn)

    var args []interface{}
    
    _category := c.Geti("category")
    if _category != 0 {
        args = append(args, models.Where{Column:"category", Value:_category, Compare:"="})    
    }
    _parent := c.Geti64("parent")
    if _parent != 0 {
        args = append(args, models.Where{Column:"parent", Value:_parent, Compare:"="})    
    }
    _value1 := c.Get("value1")
    if _value1 != "" {
        args = append(args, models.Where{Column:"value1", Value:_value1, Compare:"like"})
    }
    _value2 := c.Get("value2")
    if _value2 != "" {
        args = append(args, models.Where{Column:"value2", Value:_value2, Compare:"like"})
    }
    _value3 := c.Get("value3")
    if _value3 != "" {
        args = append(args, models.Where{Column:"value3", Value:_value3, Compare:"like"})
    }
    _value4 := c.Get("value4")
    if _value4 != "" {
        args = append(args, models.Where{Column:"value4", Value:_value4, Compare:"like"})
    }
    _value5 := c.Get("value5")
    if _value5 != "" {
        args = append(args, models.Where{Column:"value5", Value:_value5, Compare:"like"})
    }
    _value6 := c.Get("value6")
    if _value6 != "" {
        args = append(args, models.Where{Column:"value6", Value:_value6, Compare:"like"})
    }
    _value7 := c.Get("value7")
    if _value7 != "" {
        args = append(args, models.Where{Column:"value7", Value:_value7, Compare:"like"})
    }
    _value8 := c.Get("value8")
    if _value8 != "" {
        args = append(args, models.Where{Column:"value8", Value:_value8, Compare:"like"})
    }
    _value9 := c.Get("value9")
    if _value9 != "" {
        args = append(args, models.Where{Column:"value9", Value:_value9, Compare:"like"})
    }
    _value10 := c.Get("value10")
    if _value10 != "" {
        args = append(args, models.Where{Column:"value10", Value:_value10, Compare:"like"})
    }
    _content := c.Get("content")
    if _content != "" {
        args = append(args, models.Where{Column:"content", Value:_content, Compare:"="})
        
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
                    str += ", f_" + strings.Trim(v, " ")                
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

func (c *FacilityController) Insert(item *models.Facility) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewFacilityManager(conn)
	manager.Insert(item)

    id := manager.GetIdentity()
    c.Result["id"] = id
    item.Id = id
}

func (c *FacilityController) Insertbatch(item *[]models.Facility) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewFacilityManager(conn)

    for i := 0; i < rows; i++ {
	    manager.Insert(&((*item)[i]))
    }
}

func (c *FacilityController) Update(item *models.Facility) {
    
    
	conn := c.NewConnection()

	manager := models.NewFacilityManager(conn)
	manager.Update(item)
}

func (c *FacilityController) Delete(item *models.Facility) {
    
    
    conn := c.NewConnection()

	manager := models.NewFacilityManager(conn)

    
	manager.Delete(item.Id)
}

func (c *FacilityController) Deletebatch(item *[]models.Facility) {
    
    
    conn := c.NewConnection()

	manager := models.NewFacilityManager(conn)

    for _, v := range *item {
        
    
	    manager.Delete(v.Id)
    }
}


// @Delete()
func (c *FacilityController) DeleteByBuildingCategory(building int64 ,category int) {
    
    conn := c.NewConnection()

	_manager := models.NewFacilityManager(conn)
    
    _manager.DeleteByBuildingCategory(building, category)
    
}


// @Put()
func (c *FacilityController) UpdateCategory(category int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewFacilityManager(conn)
	_manager.UpdateCategory(category, id)
}

// @Put()
func (c *FacilityController) UpdateParent(parent int64, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewFacilityManager(conn)
	_manager.UpdateParent(parent, id)
}

// @Put()
func (c *FacilityController) UpdateValue1(value1 string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewFacilityManager(conn)
	_manager.UpdateValue1(value1, id)
}

// @Put()
func (c *FacilityController) UpdateValue2(value2 string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewFacilityManager(conn)
	_manager.UpdateValue2(value2, id)
}

// @Put()
func (c *FacilityController) UpdateValue3(value3 string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewFacilityManager(conn)
	_manager.UpdateValue3(value3, id)
}

// @Put()
func (c *FacilityController) UpdateValue4(value4 string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewFacilityManager(conn)
	_manager.UpdateValue4(value4, id)
}

// @Put()
func (c *FacilityController) UpdateValue5(value5 string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewFacilityManager(conn)
	_manager.UpdateValue5(value5, id)
}

// @Put()
func (c *FacilityController) UpdateValue6(value6 string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewFacilityManager(conn)
	_manager.UpdateValue6(value6, id)
}

// @Put()
func (c *FacilityController) UpdateValue7(value7 string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewFacilityManager(conn)
	_manager.UpdateValue7(value7, id)
}

// @Put()
func (c *FacilityController) UpdateValue8(value8 string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewFacilityManager(conn)
	_manager.UpdateValue8(value8, id)
}

// @Put()
func (c *FacilityController) UpdateValue9(value9 string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewFacilityManager(conn)
	_manager.UpdateValue9(value9, id)
}

// @Put()
func (c *FacilityController) UpdateValue10(value10 string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewFacilityManager(conn)
	_manager.UpdateValue10(value10, id)
}

// @Put()
func (c *FacilityController) UpdateContent(content string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewFacilityManager(conn)
	_manager.UpdateContent(content, id)
}

// @Put()
func (c *FacilityController) UpdateBuilding(building int64, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewFacilityManager(conn)
	_manager.UpdateBuilding(building, id)
}






