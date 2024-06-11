package api

import (
	"zkeep/controllers"
	"zkeep/models"
)

type UserlistController struct {
	controllers.Controller
}

func (c *UserlistController) InitData() {
	session := c.Session

	conn := c.NewConnection()

	userlistManager := models.NewUserlistManager(conn)

	items := userlistManager.Find([]interface{}{
		models.Where{Column: "company", Value: session.Company, Compare: "="},
	})

	var users int
	var totalusers int
	var score float64
	var totalscore float64
	for _, v := range items {
		if v.Status == 1 {
			users++
		}

		totalusers++

		if v.Status != 1 {
			continue
		}

		score += float64(v.Score)
		totalscore += float64(v.Totalscore)
	}

	c.Set("users", users)
	c.Set("totalusers", totalusers)
	c.Set("score", score)
	c.Set("totalscore", totalscore)
}
