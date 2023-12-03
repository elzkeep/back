package api

import (
	"aoi/controllers"
	"aoi/models"
	"strings"
)

type StatisticscolorController struct {
	controllers.Controller
}

func (c *StatisticscolorController) Search(status int, page int, pagesize int) {
	conn := c.NewConnection()
	manager := models.NewStatisticscolorManager(conn)

	var args []interface{}

	if status == 2 {
		query := "select gu_color, sum(if(gu_rank = 1, 1, 0))as gu_rank1, sum(if(gu_rank = 2, 1, 0)) as gu_rank2, sum(if(gu_rank = 3, 1, 0)) as gu_rank3, sum(if(gu_rank = 4, 1, 0)) as gu_rank4, sum(if(gu_rank = 5, 1, 0)) as gu_rank5, count(*) as gu_count, convert(avg(gu_score), int) as gu_avg from game_tb, gameuser_tb where g_status = 4 and g_map = 1 and g_id = gu_game group by gu_color"
		args = append(args, models.Base{Query: query})
	} else if status == 3 {
		query := "select gu_color, sum(if(gu_rank = 1, 1, 0))as gu_rank1, sum(if(gu_rank = 2, 1, 0)) as gu_rank2, sum(if(gu_rank = 3, 1, 0)) as gu_rank3, sum(if(gu_rank = 4, 1, 0)) as gu_rank4, sum(if(gu_rank = 5, 1, 0)) as gu_rank5, count(*) as gu_count, convert(avg(gu_score), int) as gu_avg from game_tb, gameuser_tb where g_status = 4 and g_map = 2 and g_id = gu_game group by gu_color"
		args = append(args, models.Base{Query: query})
	} else {
		query := "select gu_color, sum(if(gu_rank = 1, 1, 0))as gu_rank1, sum(if(gu_rank = 2, 1, 0)) as gu_rank2, sum(if(gu_rank = 3, 1, 0)) as gu_rank3, sum(if(gu_rank = 4, 1, 0)) as gu_rank4, sum(if(gu_rank = 5, 1, 0)) as gu_rank5, count(*) as gu_count, convert(avg(gu_score), int) as gu_avg from game_tb, gameuser_tb where g_status = 4 and g_id = gu_game group by gu_color"
		args = append(args, models.Base{Query: query})
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
