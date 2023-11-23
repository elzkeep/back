package api

import (
	"aoi/controllers"
	"aoi/models"
)

type UserController struct {
	controllers.Controller
}

func (c *UserController) Elo(name string, page int, pagesize int) {
	conn := c.NewConnection()
	manager := models.NewUserManager(conn)

	var args []interface{}

	args = append(args, models.Where{Column: "count", Value: 1, Compare: ">="})

	if page != 0 && pagesize != 0 {
		args = append(args, models.Paging(page, pagesize))
	}

	args = append(args, models.Ordering("elo desc"))

	items := manager.Find(args)
	c.Set("items", items)

	total := manager.Count(args)
	c.Set("total", total)

}
