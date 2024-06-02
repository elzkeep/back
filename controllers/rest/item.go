package rest


import (
	"zkeep/controllers"
	"zkeep/models"

    "strings"
)

type ItemController struct {
	controllers.Controller
}


// @Delete()
func (c *ItemController) DeleteByReportTopcategory(report int64 ,topcategory int) {
    
    conn := c.NewConnection()

	_manager := models.NewItemManager(conn)
    
    _manager.DeleteByReportTopcategory(report, topcategory)
    
}


func (c *ItemController) Insert(item *models.Item) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewItemManager(conn)
	manager.Insert(item)

    id := manager.GetIdentity()
    c.Result["id"] = id
    item.Id = id
}

func (c *ItemController) Insertbatch(item *[]models.Item) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewItemManager(conn)

    for i := 0; i < rows; i++ {
	    manager.Insert(&((*item)[i]))
    }
}

func (c *ItemController) Update(item *models.Item) {
    
    
	conn := c.NewConnection()

	manager := models.NewItemManager(conn)
	manager.Update(item)
}

func (c *ItemController) Delete(item *models.Item) {
    
    
    conn := c.NewConnection()

	manager := models.NewItemManager(conn)

    
	manager.Delete(item.Id)
}

func (c *ItemController) Deletebatch(item *[]models.Item) {
    
    
    conn := c.NewConnection()

	manager := models.NewItemManager(conn)

    for _, v := range *item {
        
    
	    manager.Delete(v.Id)
    }
}

func (c *ItemController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewItemManager(conn)

    var args []interface{}
    
    _title := c.Get("title")
    if _title != "" {
        args = append(args, models.Where{Column:"title", Value:_title, Compare:"="})
        
    }
    _type := c.Geti("type")
    if _type != 0 {
        args = append(args, models.Where{Column:"type", Value:_type, Compare:"="})    
    }
    _value1 := c.Geti("value1")
    if _value1 != 0 {
        args = append(args, models.Where{Column:"value1", Value:_value1, Compare:"="})    
    }
    _value2 := c.Geti("value2")
    if _value2 != 0 {
        args = append(args, models.Where{Column:"value2", Value:_value2, Compare:"="})    
    }
    _value3 := c.Geti("value3")
    if _value3 != 0 {
        args = append(args, models.Where{Column:"value3", Value:_value3, Compare:"="})    
    }
    _value4 := c.Geti("value4")
    if _value4 != 0 {
        args = append(args, models.Where{Column:"value4", Value:_value4, Compare:"="})    
    }
    _value5 := c.Geti("value5")
    if _value5 != 0 {
        args = append(args, models.Where{Column:"value5", Value:_value5, Compare:"="})    
    }
    _value6 := c.Geti("value6")
    if _value6 != 0 {
        args = append(args, models.Where{Column:"value6", Value:_value6, Compare:"="})    
    }
    _value7 := c.Geti("value7")
    if _value7 != 0 {
        args = append(args, models.Where{Column:"value7", Value:_value7, Compare:"="})    
    }
    _value8 := c.Geti("value8")
    if _value8 != 0 {
        args = append(args, models.Where{Column:"value8", Value:_value8, Compare:"="})    
    }
    _value := c.Geti("value")
    if _value != 0 {
        args = append(args, models.Where{Column:"value", Value:_value, Compare:"="})    
    }
    _content := c.Get("content")
    if _content != "" {
        args = append(args, models.Where{Column:"content", Value:_content, Compare:"="})
        
    }
    _unit := c.Get("unit")
    if _unit != "" {
        args = append(args, models.Where{Column:"unit", Value:_unit, Compare:"="})
    }
    _status := c.Geti("status")
    if _status != 0 {
        args = append(args, models.Where{Column:"status", Value:_status, Compare:"="})    
    }
    _reason := c.Geti("reason")
    if _reason != 0 {
        args = append(args, models.Where{Column:"reason", Value:_reason, Compare:"="})    
    }
    _reasontext := c.Get("reasontext")
    if _reasontext != "" {
        args = append(args, models.Where{Column:"reasontext", Value:_reasontext, Compare:"="})
    }
    _action := c.Geti("action")
    if _action != 0 {
        args = append(args, models.Where{Column:"action", Value:_action, Compare:"="})    
    }
    _actiontext := c.Get("actiontext")
    if _actiontext != "" {
        args = append(args, models.Where{Column:"actiontext", Value:_actiontext, Compare:"="})
    }
    _image := c.Get("image")
    if _image != "" {
        args = append(args, models.Where{Column:"image", Value:_image, Compare:"="})
    }
    _order := c.Geti("order")
    if _order != 0 {
        args = append(args, models.Where{Column:"order", Value:_order, Compare:"="})    
    }
    _topcategory := c.Geti("topcategory")
    if _topcategory != 0 {
        args = append(args, models.Where{Column:"topcategory", Value:_topcategory, Compare:"="})    
    }
    _data := c.Geti64("data")
    if _data != 0 {
        args = append(args, models.Where{Column:"data", Value:_data, Compare:"="})    
    }
    _report := c.Geti64("report")
    if _report != 0 {
        args = append(args, models.Where{Column:"report", Value:_report, Compare:"="})    
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


func (c *ItemController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewItemManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *ItemController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewItemManager(conn)

    var args []interface{}
    
    _title := c.Get("title")
    if _title != "" {
        args = append(args, models.Where{Column:"title", Value:_title, Compare:"="})
        
    }
    _type := c.Geti("type")
    if _type != 0 {
        args = append(args, models.Where{Column:"type", Value:_type, Compare:"="})    
    }
    _value1 := c.Geti("value1")
    if _value1 != 0 {
        args = append(args, models.Where{Column:"value1", Value:_value1, Compare:"="})    
    }
    _value2 := c.Geti("value2")
    if _value2 != 0 {
        args = append(args, models.Where{Column:"value2", Value:_value2, Compare:"="})    
    }
    _value3 := c.Geti("value3")
    if _value3 != 0 {
        args = append(args, models.Where{Column:"value3", Value:_value3, Compare:"="})    
    }
    _value4 := c.Geti("value4")
    if _value4 != 0 {
        args = append(args, models.Where{Column:"value4", Value:_value4, Compare:"="})    
    }
    _value5 := c.Geti("value5")
    if _value5 != 0 {
        args = append(args, models.Where{Column:"value5", Value:_value5, Compare:"="})    
    }
    _value6 := c.Geti("value6")
    if _value6 != 0 {
        args = append(args, models.Where{Column:"value6", Value:_value6, Compare:"="})    
    }
    _value7 := c.Geti("value7")
    if _value7 != 0 {
        args = append(args, models.Where{Column:"value7", Value:_value7, Compare:"="})    
    }
    _value8 := c.Geti("value8")
    if _value8 != 0 {
        args = append(args, models.Where{Column:"value8", Value:_value8, Compare:"="})    
    }
    _value := c.Geti("value")
    if _value != 0 {
        args = append(args, models.Where{Column:"value", Value:_value, Compare:"="})    
    }
    _content := c.Get("content")
    if _content != "" {
        args = append(args, models.Where{Column:"content", Value:_content, Compare:"="})
        
    }
    _unit := c.Get("unit")
    if _unit != "" {
        args = append(args, models.Where{Column:"unit", Value:_unit, Compare:"="})
    }
    _status := c.Geti("status")
    if _status != 0 {
        args = append(args, models.Where{Column:"status", Value:_status, Compare:"="})    
    }
    _reason := c.Geti("reason")
    if _reason != 0 {
        args = append(args, models.Where{Column:"reason", Value:_reason, Compare:"="})    
    }
    _reasontext := c.Get("reasontext")
    if _reasontext != "" {
        args = append(args, models.Where{Column:"reasontext", Value:_reasontext, Compare:"="})
    }
    _action := c.Geti("action")
    if _action != 0 {
        args = append(args, models.Where{Column:"action", Value:_action, Compare:"="})    
    }
    _actiontext := c.Get("actiontext")
    if _actiontext != "" {
        args = append(args, models.Where{Column:"actiontext", Value:_actiontext, Compare:"="})
    }
    _image := c.Get("image")
    if _image != "" {
        args = append(args, models.Where{Column:"image", Value:_image, Compare:"="})
    }
    _order := c.Geti("order")
    if _order != 0 {
        args = append(args, models.Where{Column:"order", Value:_order, Compare:"="})    
    }
    _topcategory := c.Geti("topcategory")
    if _topcategory != 0 {
        args = append(args, models.Where{Column:"topcategory", Value:_topcategory, Compare:"="})    
    }
    _data := c.Geti64("data")
    if _data != 0 {
        args = append(args, models.Where{Column:"data", Value:_data, Compare:"="})    
    }
    _report := c.Geti64("report")
    if _report != 0 {
        args = append(args, models.Where{Column:"report", Value:_report, Compare:"="})    
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
                    str += ", i_" + strings.Trim(v, " ")                
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
func (c *ItemController) UpdateTitle(title string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewItemManager(conn)
	_manager.UpdateTitle(title, id)
}
// @Put()
func (c *ItemController) UpdateType(typeid int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewItemManager(conn)
	_manager.UpdateType(typeid, id)
}
// @Put()
func (c *ItemController) UpdateValue1(value1 int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewItemManager(conn)
	_manager.UpdateValue1(value1, id)
}
// @Put()
func (c *ItemController) UpdateValue2(value2 int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewItemManager(conn)
	_manager.UpdateValue2(value2, id)
}
// @Put()
func (c *ItemController) UpdateValue3(value3 int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewItemManager(conn)
	_manager.UpdateValue3(value3, id)
}
// @Put()
func (c *ItemController) UpdateValue4(value4 int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewItemManager(conn)
	_manager.UpdateValue4(value4, id)
}
// @Put()
func (c *ItemController) UpdateValue5(value5 int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewItemManager(conn)
	_manager.UpdateValue5(value5, id)
}
// @Put()
func (c *ItemController) UpdateValue6(value6 int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewItemManager(conn)
	_manager.UpdateValue6(value6, id)
}
// @Put()
func (c *ItemController) UpdateValue7(value7 int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewItemManager(conn)
	_manager.UpdateValue7(value7, id)
}
// @Put()
func (c *ItemController) UpdateValue8(value8 int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewItemManager(conn)
	_manager.UpdateValue8(value8, id)
}
// @Put()
func (c *ItemController) UpdateValue(value int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewItemManager(conn)
	_manager.UpdateValue(value, id)
}
// @Put()
func (c *ItemController) UpdateContent(content string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewItemManager(conn)
	_manager.UpdateContent(content, id)
}
// @Put()
func (c *ItemController) UpdateUnit(unit string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewItemManager(conn)
	_manager.UpdateUnit(unit, id)
}
// @Put()
func (c *ItemController) UpdateStatus(status int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewItemManager(conn)
	_manager.UpdateStatus(status, id)
}
// @Put()
func (c *ItemController) UpdateReason(reason int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewItemManager(conn)
	_manager.UpdateReason(reason, id)
}
// @Put()
func (c *ItemController) UpdateReasontext(reasontext string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewItemManager(conn)
	_manager.UpdateReasontext(reasontext, id)
}
// @Put()
func (c *ItemController) UpdateAction(action int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewItemManager(conn)
	_manager.UpdateAction(action, id)
}
// @Put()
func (c *ItemController) UpdateActiontext(actiontext string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewItemManager(conn)
	_manager.UpdateActiontext(actiontext, id)
}
// @Put()
func (c *ItemController) UpdateImage(image string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewItemManager(conn)
	_manager.UpdateImage(image, id)
}
// @Put()
func (c *ItemController) UpdateOrder(order int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewItemManager(conn)
	_manager.UpdateOrder(order, id)
}
// @Put()
func (c *ItemController) UpdateTopcategory(topcategory int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewItemManager(conn)
	_manager.UpdateTopcategory(topcategory, id)
}
// @Put()
func (c *ItemController) UpdateData(data int64, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewItemManager(conn)
	_manager.UpdateData(data, id)
}
// @Put()
func (c *ItemController) UpdateReport(report int64, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewItemManager(conn)
	_manager.UpdateReport(report, id)
}





