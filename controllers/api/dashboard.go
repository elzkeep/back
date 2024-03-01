package api

import (
	"zkeep/controllers"
	"zkeep/models"
)

type DashboardController struct {
	controllers.Controller
}

func (c *DashboardController) Index(company int64) {
	conn := c.NewConnection()

	userManager := models.NewUserManager(conn)
	buildingManager := models.NewBuildingManager(conn)

	userscore := userManager.Sum([]interface{}{models.Where{Column: "company", Value: company, Compare: "="}})
	buildingscore := buildingManager.Sum([]interface{}{models.Where{Column: "company", Value: company, Compare: "="}})

	c.Set("userscore", userscore)
	c.Set("buildingscore", buildingscore)
}
