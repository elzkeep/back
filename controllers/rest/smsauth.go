package rest


import (
	"aoi/controllers"
	"aoi/models"

    "strings"
)

type SmsauthController struct {
	controllers.Controller
}

func (c *SmsauthController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewSmsauthManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *SmsauthController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewSmsauthManager(conn)

    var args []interface{}
    
    _hp := c.Get("hp")
    if _hp != "" {
        args = append(args, models.Where{Column:"hp", Value:_hp, Compare:"like"})
    }
    _number := c.Get("number")
    if _number != "" {
        args = append(args, models.Where{Column:"number", Value:_number, Compare:"like"})
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
                    str += ", sa_" + strings.Trim(v, " ")                
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

func (c *SmsauthController) Insert(item *models.Smsauth) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewSmsauthManager(conn)
	manager.Insert(item)

    id := manager.GetIdentity()
    c.Result["id"] = id
    item.Id = id
}

func (c *SmsauthController) Insertbatch(item *[]models.Smsauth) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewSmsauthManager(conn)

    for i := 0; i < rows; i++ {
	    manager.Insert(&((*item)[i]))
    }
}

func (c *SmsauthController) Update(item *models.Smsauth) {
    
    
	conn := c.NewConnection()

	manager := models.NewSmsauthManager(conn)
	manager.Update(item)
}

func (c *SmsauthController) Delete(item *models.Smsauth) {
    
    
    conn := c.NewConnection()

	manager := models.NewSmsauthManager(conn)

    
	manager.Delete(item.Id)
}

func (c *SmsauthController) Deletebatch(item *[]models.Smsauth) {
    
    
    conn := c.NewConnection()

	manager := models.NewSmsauthManager(conn)

    for _, v := range *item {
        
    
	    manager.Delete(v.Id)
    }
}



func (c *SmsauthController) GetByHpNumber(hp string ,number string) *models.Smsauth {
    
    conn := c.NewConnection()

	_manager := models.NewSmsauthManager(conn)
    
    item := _manager.GetByHpNumber(hp, number)
    
    c.Set("item", item)
    
    
    
    return item
    
}

// @Delete()
func (c *SmsauthController) DeleteByHp(hp string) {
    
    conn := c.NewConnection()

	_manager := models.NewSmsauthManager(conn)
    
    _manager.DeleteByHp(hp)
    
}


// @Put()
func (c *SmsauthController) UpdateHp(hp string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewSmsauthManager(conn)
	_manager.UpdateHp(hp, id)
}

// @Put()
func (c *SmsauthController) UpdateNumber(number string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewSmsauthManager(conn)
	_manager.UpdateNumber(number, id)
}






