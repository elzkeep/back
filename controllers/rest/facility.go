package rest


import (
	"zkeep/controllers"
	"zkeep/models"

    "strings"
)

type FacilityController struct {
	controllers.Controller
}


// @Delete()
func (c *FacilityController) DeleteByBuildingCategory(building int64 ,category int) {
    
    conn := c.NewConnection()

	_manager := models.NewFacilityManager(conn)
    
    _manager.DeleteByBuildingCategory(building, category)
    
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

func (c *FacilityController) Count() {
    
    
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
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"="})
        
    }
    _type := c.Geti("type")
    if _type != 0 {
        args = append(args, models.Where{Column:"type", Value:_type, Compare:"="})    
    }
    _value1 := c.Get("value1")
    if _value1 != "" {
        args = append(args, models.Where{Column:"value1", Value:_value1, Compare:"="})
    }
    _value2 := c.Get("value2")
    if _value2 != "" {
        args = append(args, models.Where{Column:"value2", Value:_value2, Compare:"="})
    }
    _value3 := c.Get("value3")
    if _value3 != "" {
        args = append(args, models.Where{Column:"value3", Value:_value3, Compare:"="})
    }
    _value4 := c.Get("value4")
    if _value4 != "" {
        args = append(args, models.Where{Column:"value4", Value:_value4, Compare:"="})
    }
    _value5 := c.Get("value5")
    if _value5 != "" {
        args = append(args, models.Where{Column:"value5", Value:_value5, Compare:"="})
    }
    _value6 := c.Get("value6")
    if _value6 != "" {
        args = append(args, models.Where{Column:"value6", Value:_value6, Compare:"="})
    }
    _value7 := c.Get("value7")
    if _value7 != "" {
        args = append(args, models.Where{Column:"value7", Value:_value7, Compare:"="})
    }
    _value8 := c.Get("value8")
    if _value8 != "" {
        args = append(args, models.Where{Column:"value8", Value:_value8, Compare:"="})
    }
    _value9 := c.Get("value9")
    if _value9 != "" {
        args = append(args, models.Where{Column:"value9", Value:_value9, Compare:"="})
    }
    _value10 := c.Get("value10")
    if _value10 != "" {
        args = append(args, models.Where{Column:"value10", Value:_value10, Compare:"="})
    }
    _value11 := c.Get("value11")
    if _value11 != "" {
        args = append(args, models.Where{Column:"value11", Value:_value11, Compare:"="})
    }
    _value12 := c.Get("value12")
    if _value12 != "" {
        args = append(args, models.Where{Column:"value12", Value:_value12, Compare:"="})
    }
    _value13 := c.Get("value13")
    if _value13 != "" {
        args = append(args, models.Where{Column:"value13", Value:_value13, Compare:"="})
    }
    _value14 := c.Get("value14")
    if _value14 != "" {
        args = append(args, models.Where{Column:"value14", Value:_value14, Compare:"="})
    }
    _value15 := c.Get("value15")
    if _value15 != "" {
        args = append(args, models.Where{Column:"value15", Value:_value15, Compare:"="})
    }
    _value16 := c.Get("value16")
    if _value16 != "" {
        args = append(args, models.Where{Column:"value16", Value:_value16, Compare:"="})
    }
    _value17 := c.Get("value17")
    if _value17 != "" {
        args = append(args, models.Where{Column:"value17", Value:_value17, Compare:"="})
    }
    _value18 := c.Get("value18")
    if _value18 != "" {
        args = append(args, models.Where{Column:"value18", Value:_value18, Compare:"="})
    }
    _value19 := c.Get("value19")
    if _value19 != "" {
        args = append(args, models.Where{Column:"value19", Value:_value19, Compare:"="})
    }
    _value20 := c.Get("value20")
    if _value20 != "" {
        args = append(args, models.Where{Column:"value20", Value:_value20, Compare:"="})
    }
    _value21 := c.Get("value21")
    if _value21 != "" {
        args = append(args, models.Where{Column:"value21", Value:_value21, Compare:"="})
    }
    _value22 := c.Get("value22")
    if _value22 != "" {
        args = append(args, models.Where{Column:"value22", Value:_value22, Compare:"="})
    }
    _value23 := c.Get("value23")
    if _value23 != "" {
        args = append(args, models.Where{Column:"value23", Value:_value23, Compare:"="})
    }
    _value24 := c.Get("value24")
    if _value24 != "" {
        args = append(args, models.Where{Column:"value24", Value:_value24, Compare:"="})
    }
    _value25 := c.Get("value25")
    if _value25 != "" {
        args = append(args, models.Where{Column:"value25", Value:_value25, Compare:"="})
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
    

    
    
    total := manager.Count(args)
	c.Set("total", total)
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
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"="})
        
    }
    _type := c.Geti("type")
    if _type != 0 {
        args = append(args, models.Where{Column:"type", Value:_type, Compare:"="})    
    }
    _value1 := c.Get("value1")
    if _value1 != "" {
        args = append(args, models.Where{Column:"value1", Value:_value1, Compare:"="})
    }
    _value2 := c.Get("value2")
    if _value2 != "" {
        args = append(args, models.Where{Column:"value2", Value:_value2, Compare:"="})
    }
    _value3 := c.Get("value3")
    if _value3 != "" {
        args = append(args, models.Where{Column:"value3", Value:_value3, Compare:"="})
    }
    _value4 := c.Get("value4")
    if _value4 != "" {
        args = append(args, models.Where{Column:"value4", Value:_value4, Compare:"="})
    }
    _value5 := c.Get("value5")
    if _value5 != "" {
        args = append(args, models.Where{Column:"value5", Value:_value5, Compare:"="})
    }
    _value6 := c.Get("value6")
    if _value6 != "" {
        args = append(args, models.Where{Column:"value6", Value:_value6, Compare:"="})
    }
    _value7 := c.Get("value7")
    if _value7 != "" {
        args = append(args, models.Where{Column:"value7", Value:_value7, Compare:"="})
    }
    _value8 := c.Get("value8")
    if _value8 != "" {
        args = append(args, models.Where{Column:"value8", Value:_value8, Compare:"="})
    }
    _value9 := c.Get("value9")
    if _value9 != "" {
        args = append(args, models.Where{Column:"value9", Value:_value9, Compare:"="})
    }
    _value10 := c.Get("value10")
    if _value10 != "" {
        args = append(args, models.Where{Column:"value10", Value:_value10, Compare:"="})
    }
    _value11 := c.Get("value11")
    if _value11 != "" {
        args = append(args, models.Where{Column:"value11", Value:_value11, Compare:"="})
    }
    _value12 := c.Get("value12")
    if _value12 != "" {
        args = append(args, models.Where{Column:"value12", Value:_value12, Compare:"="})
    }
    _value13 := c.Get("value13")
    if _value13 != "" {
        args = append(args, models.Where{Column:"value13", Value:_value13, Compare:"="})
    }
    _value14 := c.Get("value14")
    if _value14 != "" {
        args = append(args, models.Where{Column:"value14", Value:_value14, Compare:"="})
    }
    _value15 := c.Get("value15")
    if _value15 != "" {
        args = append(args, models.Where{Column:"value15", Value:_value15, Compare:"="})
    }
    _value16 := c.Get("value16")
    if _value16 != "" {
        args = append(args, models.Where{Column:"value16", Value:_value16, Compare:"="})
    }
    _value17 := c.Get("value17")
    if _value17 != "" {
        args = append(args, models.Where{Column:"value17", Value:_value17, Compare:"="})
    }
    _value18 := c.Get("value18")
    if _value18 != "" {
        args = append(args, models.Where{Column:"value18", Value:_value18, Compare:"="})
    }
    _value19 := c.Get("value19")
    if _value19 != "" {
        args = append(args, models.Where{Column:"value19", Value:_value19, Compare:"="})
    }
    _value20 := c.Get("value20")
    if _value20 != "" {
        args = append(args, models.Where{Column:"value20", Value:_value20, Compare:"="})
    }
    _value21 := c.Get("value21")
    if _value21 != "" {
        args = append(args, models.Where{Column:"value21", Value:_value21, Compare:"="})
    }
    _value22 := c.Get("value22")
    if _value22 != "" {
        args = append(args, models.Where{Column:"value22", Value:_value22, Compare:"="})
    }
    _value23 := c.Get("value23")
    if _value23 != "" {
        args = append(args, models.Where{Column:"value23", Value:_value23, Compare:"="})
    }
    _value24 := c.Get("value24")
    if _value24 != "" {
        args = append(args, models.Where{Column:"value24", Value:_value24, Compare:"="})
    }
    _value25 := c.Get("value25")
    if _value25 != "" {
        args = append(args, models.Where{Column:"value25", Value:_value25, Compare:"="})
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
func (c *FacilityController) UpdateName(name string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewFacilityManager(conn)
	_manager.UpdateName(name, id)
}
// @Put()
func (c *FacilityController) UpdateType(typeid int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewFacilityManager(conn)
	_manager.UpdateType(typeid, id)
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
func (c *FacilityController) UpdateValue11(value11 string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewFacilityManager(conn)
	_manager.UpdateValue11(value11, id)
}
// @Put()
func (c *FacilityController) UpdateValue12(value12 string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewFacilityManager(conn)
	_manager.UpdateValue12(value12, id)
}
// @Put()
func (c *FacilityController) UpdateValue13(value13 string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewFacilityManager(conn)
	_manager.UpdateValue13(value13, id)
}
// @Put()
func (c *FacilityController) UpdateValue14(value14 string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewFacilityManager(conn)
	_manager.UpdateValue14(value14, id)
}
// @Put()
func (c *FacilityController) UpdateValue15(value15 string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewFacilityManager(conn)
	_manager.UpdateValue15(value15, id)
}
// @Put()
func (c *FacilityController) UpdateValue16(value16 string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewFacilityManager(conn)
	_manager.UpdateValue16(value16, id)
}
// @Put()
func (c *FacilityController) UpdateValue17(value17 string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewFacilityManager(conn)
	_manager.UpdateValue17(value17, id)
}
// @Put()
func (c *FacilityController) UpdateValue18(value18 string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewFacilityManager(conn)
	_manager.UpdateValue18(value18, id)
}
// @Put()
func (c *FacilityController) UpdateValue19(value19 string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewFacilityManager(conn)
	_manager.UpdateValue19(value19, id)
}
// @Put()
func (c *FacilityController) UpdateValue20(value20 string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewFacilityManager(conn)
	_manager.UpdateValue20(value20, id)
}
// @Put()
func (c *FacilityController) UpdateValue21(value21 string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewFacilityManager(conn)
	_manager.UpdateValue21(value21, id)
}
// @Put()
func (c *FacilityController) UpdateValue22(value22 string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewFacilityManager(conn)
	_manager.UpdateValue22(value22, id)
}
// @Put()
func (c *FacilityController) UpdateValue23(value23 string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewFacilityManager(conn)
	_manager.UpdateValue23(value23, id)
}
// @Put()
func (c *FacilityController) UpdateValue24(value24 string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewFacilityManager(conn)
	_manager.UpdateValue24(value24, id)
}
// @Put()
func (c *FacilityController) UpdateValue25(value25 string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewFacilityManager(conn)
	_manager.UpdateValue25(value25, id)
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





