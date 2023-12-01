package api

import (
	"aoi/controllers"
	"aoi/models"
	"strings"
)

type GamelistController struct {
	controllers.Controller
}

func (c *GamelistController) Search(status int, page int, pagesize int) {
	user := c.Session.Id

	conn := c.NewConnection()

	manager := models.NewGamelistManager(conn)
	gameuserManager := models.NewGameuserManager(conn)

	var args []interface{}

	args = append(args, models.Where{Column: "gameuser", Value: user, Compare: "="})

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
					str += ", g_" + strings.Trim(v, " ")
				}
			}
		}

		args = append(args, models.Ordering(str))
	}

	items := manager.Find(args)

	for i, v := range items {
		gameusers := gameuserManager.Find([]interface{}{
			models.Where{Column: "game", Value: v.Id, Compare: "="},
		})

		users := make([]models.Gameuser, 0)

		for _, v := range gameusers {
			user := v.Extra["user"].(models.User)
			user.Passwd = ""

			v.Extra["user"] = user
			users = append(users, v)
		}

		items[i].Extra["users"] = users
	}

	c.Set("items", items)

	total := manager.Count(args)
	c.Set("total", total)
}
