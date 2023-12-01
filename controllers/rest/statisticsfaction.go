package rest


import (
	"aoi/controllers"
	"aoi/models"

    "strings"
)

type StatisticsfactionController struct {
	controllers.Controller
}

func (c *StatisticsfactionController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewStatisticsfactionManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *StatisticsfactionController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewStatisticsfactionManager(conn)

    var args []interface{}
    
    _faction := c.Geti("faction")
    if _faction != 0 {
        args = append(args, models.Where{Column:"faction", Value:_faction, Compare:"="})    
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
    _count := c.Geti64("count")
    if _count != 0 {
        args = append(args, models.Where{Column:"count", Value:_count, Compare:"="})    
    }
    _avg := c.Geti64("avg")
    if _avg != 0 {
        args = append(args, models.Where{Column:"avg", Value:_avg, Compare:"="})    
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
                    str += ", gu_" + strings.Trim(v, " ")                
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










func (c *StatisticsfactionController) Sum() {
    
    
	conn := c.NewConnection()

	manager := models.NewStatisticsfactionManager(conn)

    var args []interface{}
    
    _faction := c.Geti("faction")
    if _faction != 0 {
        args = append(args, models.Where{Column:"faction", Value:_faction, Compare:"="})    
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
    _count := c.Geti64("count")
    if _count != 0 {
        args = append(args, models.Where{Column:"count", Value:_count, Compare:"="})    
    }
    _avg := c.Geti64("avg")
    if _avg != 0 {
        args = append(args, models.Where{Column:"avg", Value:_avg, Compare:"="})    
    }
    

    
    
    item := manager.Sum(args)
	c.Set("item", item)
}

