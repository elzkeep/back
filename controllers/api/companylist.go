package api

import (
	"strings"
	"zkeep/controllers"
	"zkeep/models"
)

type CompanylistController struct {
	controllers.Controller
}

func (c *CompanylistController) Search(page int, pagesize int) {    
	
	conn := c.NewConnection()

	manager := models.NewCompanylistManager(conn)

    var args []interface{}
    
    _name := c.Get("name")
    if _name != "" {
		args = append(args, models.Where{Column: "name", Value: _name, Compare: "like"})
    }
    _companyno := c.Get("companyno")
    if _companyno != "" {
        args = append(args, models.Where{Column:"companyno", Value:_companyno, Compare:"like"})
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
    _tel := c.Get("tel")
    if _tel != "" {
        args = append(args, models.Where{Column:"tel", Value:_tel, Compare:"like"})
    }
    _email := c.Get("email")
    if _email != "" {
        args = append(args, models.Where{Column:"email", Value:_email, Compare:"like"})
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
    _company := c.Geti64("company")
    if _company != 0 {
        args = append(args, models.Where{Column:"company", Value:_company, Compare:"="})    
    }
    _buildingcount := c.Geti64("buildingcount")
    if _buildingcount != 0 {
        args = append(args, models.Where{Column:"buildingcount", Value:_buildingcount, Compare:"="})    
    }
    _contractprice := c.Geti("contractprice")
    if _contractprice != 0 {
        args = append(args, models.Where{Column:"contractprice", Value:_contractprice, Compare:"="})    
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