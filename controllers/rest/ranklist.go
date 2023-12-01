package rest


import (
	"aoi/controllers"
	"aoi/models"

    "strings"
)

type RanklistController struct {
	controllers.Controller
}

func (c *RanklistController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewRanklistManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *RanklistController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewRanklistManager(conn)

    var args []interface{}
    
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"="})
        
    }
    _elo := c.Geti("elo")
    if _elo != 0 {
        args = append(args, models.Where{Column:"elo", Value:_elo, Compare:"="})    
    }
    _score := c.Geti("score")
    if _score != 0 {
        args = append(args, models.Where{Column:"score", Value:_score, Compare:"="})    
    }
    _count := c.Geti64("count")
    if _count != 0 {
        args = append(args, models.Where{Column:"count", Value:_count, Compare:"="})    
    }
    _rank1 := c.Geti("rank1")
    if _rank1 != 0 {
        args = append(args, models.Where{Column:"rank1", Value:_rank1, Compare:"="})    
    }
    _rank2 := c.Geti("rank2")
    if _rank2 != 0 {
        args = append(args, models.Where{Column:"rank2", Value:_rank2, Compare:"="})    
    }
    _rank3 := c.Geti("rank3")
    if _rank3 != 0 {
        args = append(args, models.Where{Column:"rank3", Value:_rank3, Compare:"="})    
    }
    _rank4 := c.Geti("rank4")
    if _rank4 != 0 {
        args = append(args, models.Where{Column:"rank4", Value:_rank4, Compare:"="})    
    }
    _rank5 := c.Geti("rank5")
    if _rank5 != 0 {
        args = append(args, models.Where{Column:"rank5", Value:_rank5, Compare:"="})    
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










func (c *RanklistController) Sum() {
    
    
	conn := c.NewConnection()

	manager := models.NewRanklistManager(conn)

    var args []interface{}
    
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"="})
        
    }
    _elo := c.Geti("elo")
    if _elo != 0 {
        args = append(args, models.Where{Column:"elo", Value:_elo, Compare:"="})    
    }
    _score := c.Geti("score")
    if _score != 0 {
        args = append(args, models.Where{Column:"score", Value:_score, Compare:"="})    
    }
    _count := c.Geti64("count")
    if _count != 0 {
        args = append(args, models.Where{Column:"count", Value:_count, Compare:"="})    
    }
    _rank1 := c.Geti("rank1")
    if _rank1 != 0 {
        args = append(args, models.Where{Column:"rank1", Value:_rank1, Compare:"="})    
    }
    _rank2 := c.Geti("rank2")
    if _rank2 != 0 {
        args = append(args, models.Where{Column:"rank2", Value:_rank2, Compare:"="})    
    }
    _rank3 := c.Geti("rank3")
    if _rank3 != 0 {
        args = append(args, models.Where{Column:"rank3", Value:_rank3, Compare:"="})    
    }
    _rank4 := c.Geti("rank4")
    if _rank4 != 0 {
        args = append(args, models.Where{Column:"rank4", Value:_rank4, Compare:"="})    
    }
    _rank5 := c.Geti("rank5")
    if _rank5 != 0 {
        args = append(args, models.Where{Column:"rank5", Value:_rank5, Compare:"="})    
    }
    

    
    
    item := manager.Sum(args)
	c.Set("item", item)
}

