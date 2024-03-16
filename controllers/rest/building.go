package rest


import (
	"zkeep/controllers"
	"zkeep/models"

    "strings"
)

type BuildingController struct {
	controllers.Controller
}

func (c *BuildingController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewBuildingManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *BuildingController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewBuildingManager(conn)

    var args []interface{}
    
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"like"})
        
    }
    _conpanyno := c.Get("conpanyno")
    if _conpanyno != "" {
        args = append(args, models.Where{Column:"conpanyno", Value:_conpanyno, Compare:"like"})
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
    _score := c.Geti("score")
    if _score != 0 {
        args = append(args, models.Where{Column:"score", Value:_score, Compare:"="})    
    }
    _status := c.Geti("status")
    if _status != 0 {
        args = append(args, models.Where{Column:"status", Value:_status, Compare:"="})    
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
                    str += ", b_" + strings.Trim(v, " ")                
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

func (c *BuildingController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewBuildingManager(conn)

    var args []interface{}
    
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"like"})
        
    }
    _conpanyno := c.Get("conpanyno")
    if _conpanyno != "" {
        args = append(args, models.Where{Column:"conpanyno", Value:_conpanyno, Compare:"like"})
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
    _score := c.Geti("score")
    if _score != 0 {
        args = append(args, models.Where{Column:"score", Value:_score, Compare:"="})    
    }
    _status := c.Geti("status")
    if _status != 0 {
        args = append(args, models.Where{Column:"status", Value:_status, Compare:"="})    
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

func (c *BuildingController) Insert(item *models.Building) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewBuildingManager(conn)
	manager.Insert(item)

    id := manager.GetIdentity()
    c.Result["id"] = id
    item.Id = id
}

func (c *BuildingController) Insertbatch(item *[]models.Building) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewBuildingManager(conn)

    for i := 0; i < rows; i++ {
	    manager.Insert(&((*item)[i]))
    }
}

func (c *BuildingController) Update(item *models.Building) {
    
    
	conn := c.NewConnection()

	manager := models.NewBuildingManager(conn)
	manager.Update(item)
}

func (c *BuildingController) Delete(item *models.Building) {
    
    
    conn := c.NewConnection()

	manager := models.NewBuildingManager(conn)

    
	manager.Delete(item.Id)
}

func (c *BuildingController) Deletebatch(item *[]models.Building) {
    
    
    conn := c.NewConnection()

	manager := models.NewBuildingManager(conn)

    for _, v := range *item {
        
    
	    manager.Delete(v.Id)
    }
}



// @Put()
func (c *BuildingController) UpdateName(name string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewBuildingManager(conn)
	_manager.UpdateName(name, id)
}

// @Put()
func (c *BuildingController) UpdateConpanyno(conpanyno string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewBuildingManager(conn)
	_manager.UpdateConpanyno(conpanyno, id)
}

// @Put()
func (c *BuildingController) UpdateCeo(ceo string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewBuildingManager(conn)
	_manager.UpdateCeo(ceo, id)
}

// @Put()
func (c *BuildingController) UpdateAddress(address string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewBuildingManager(conn)
	_manager.UpdateAddress(address, id)
}

// @Put()
func (c *BuildingController) UpdateAddressetc(addressetc string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewBuildingManager(conn)
	_manager.UpdateAddressetc(addressetc, id)
}

// @Put()
func (c *BuildingController) UpdateScore(score models.Double, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewBuildingManager(conn)
	_manager.UpdateScore(score, id)
}

// @Put()
func (c *BuildingController) UpdateStatus(status int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewBuildingManager(conn)
	_manager.UpdateStatus(status, id)
}

// @Put()
func (c *BuildingController) UpdateCompany(company int64, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewBuildingManager(conn)
	_manager.UpdateCompany(company, id)
}






func (c *BuildingController) Sum() {
    
    
	conn := c.NewConnection()

	manager := models.NewBuildingManager(conn)

    var args []interface{}
    
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"like"})
        
    }
    _conpanyno := c.Get("conpanyno")
    if _conpanyno != "" {
        args = append(args, models.Where{Column:"conpanyno", Value:_conpanyno, Compare:"like"})
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
    _score := c.Geti("score")
    if _score != 0 {
        args = append(args, models.Where{Column:"score", Value:_score, Compare:"="})    
    }
    _status := c.Geti("status")
    if _status != 0 {
        args = append(args, models.Where{Column:"status", Value:_status, Compare:"="})    
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
    

    
    
    item := manager.Sum(args)
	c.Set("item", item)
}

