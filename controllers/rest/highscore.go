package rest


import (
	"aoi/controllers"
	"aoi/models"

    "strings"
)

type HighscoreController struct {
	controllers.Controller
}

func (c *HighscoreController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewHighscoreManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *HighscoreController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewHighscoreManager(conn)

    var args []interface{}
    
    _count := c.Geti("count")
    if _count != 0 {
        args = append(args, models.Where{Column:"count", Value:_count, Compare:"="})    
    }
    _faction := c.Geti("faction")
    if _faction != 0 {
        args = append(args, models.Where{Column:"faction", Value:_faction, Compare:"="})    
    }
    _score := c.Geti("score")
    if _score != 0 {
        args = append(args, models.Where{Column:"score", Value:_score, Compare:"="})    
    }
    _user := c.Geti64("user")
    if _user != 0 {
        args = append(args, models.Where{Column:"user", Value:_user, Compare:"="})    
    }
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"="})
        
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
                    str += ", u_" + strings.Trim(v, " ")                
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










func (c *HighscoreController) Sum() {
    
    
	conn := c.NewConnection()

	manager := models.NewHighscoreManager(conn)

    var args []interface{}
    
    _count := c.Geti("count")
    if _count != 0 {
        args = append(args, models.Where{Column:"count", Value:_count, Compare:"="})    
    }
    _faction := c.Geti("faction")
    if _faction != 0 {
        args = append(args, models.Where{Column:"faction", Value:_faction, Compare:"="})    
    }
    _score := c.Geti("score")
    if _score != 0 {
        args = append(args, models.Where{Column:"score", Value:_score, Compare:"="})    
    }
    _user := c.Geti64("user")
    if _user != 0 {
        args = append(args, models.Where{Column:"user", Value:_user, Compare:"="})    
    }
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"="})
        
    }
    

    
    
    item := manager.Sum(args)
	c.Set("item", item)
}

