package api

import (
	"aoi/controllers"
	"aoi/global"
	"aoi/models"
	"fmt"
)

type RanklistController struct {
	controllers.Controller
}

func (c *RanklistController) Elo(status int, page int, pagesize int) {
	conn := c.NewConnection()
	manager := models.NewRanklistManager(conn)

	var args []interface{}

	if status == 2 {
		date := global.GetMonthStartDatetime()
		query := fmt.Sprintf("select gu_user as r_id, u_name as r_name, sum(gu_elo) + 1000.0 as r_elo, convert(avg(gu_score), int) as r_score, count(*) as  r_count, sum(if(gu_rank = 1, 1, 0)) as r_rank1, sum(if(gu_rank = 2, 1, 0)) as r_rank2, sum(if(gu_rank = 3, 1, 0)) as r_rank3, sum(if(gu_rank = 4, 1, 0)) as r_rank4, sum(if(gu_rank = 5, 1, 0)) as r_rank5 from game_tb, gameuser_tb, user_tb where g_status = 4 and g_id = gu_game and gu_user = u_id and gu_date >= '%v' group by gu_user, u_name", date)
		args = append(args, models.Base{Query: query})
	} else if status == 3 {
		date := global.GetTodayDatetime()
		query := fmt.Sprintf("select gu_user as r_id, u_name as r_name, sum(gu_elo) + 1000.0 as r_elo, convert(avg(gu_score), int) as r_score, count(*) as  r_count, sum(if(gu_rank = 1, 1, 0)) as r_rank1, sum(if(gu_rank = 2, 1, 0)) as r_rank2, sum(if(gu_rank = 3, 1, 0)) as r_rank3, sum(if(gu_rank = 4, 1, 0)) as r_rank4, sum(if(gu_rank = 5, 1, 0)) as r_rank5 from game_tb, gameuser_tb, user_tb where g_status = 4 and g_id = gu_game and gu_user = u_id and gu_date >= '%v' group by gu_user, u_name", date)
		args = append(args, models.Base{Query: query})
	} else {
		query := "select gu_user as r_id, u_name as r_name, sum(gu_elo) + 1000.0 as r_elo, convert(avg(gu_score), int) as r_score, count(*) as  r_count, sum(if(gu_rank = 1, 1, 0)) as r_rank1, sum(if(gu_rank = 2, 1, 0)) as r_rank2, sum(if(gu_rank = 3, 1, 0)) as r_rank3, sum(if(gu_rank = 4, 1, 0)) as r_rank4, sum(if(gu_rank = 5, 1, 0)) as r_rank5 from game_tb, gameuser_tb, user_tb where g_status = 4 and g_id = gu_game and gu_user = u_id group by gu_user, u_name"
		args = append(args, models.Base{Query: query})
	}

	if page != 0 && pagesize != 0 {
		args = append(args, models.Paging(page, pagesize))
	}

	args = append(args, models.Ordering("elo desc"))

	items := manager.Find(args)
	c.Set("items", items)

	total := manager.Count(args)
	c.Set("total", total)

}
