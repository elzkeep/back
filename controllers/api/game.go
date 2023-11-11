package api

import (
	"aoi/controllers"
	"aoi/game"
	"aoi/models"
	"log"
)

type GameController struct {
	controllers.Controller
}

// @Post()
func (c *GameController) Make(item *models.Game) {
	user := c.Session.Id

	game.Make(user, item)
}

// @Post()
func (c *GameController) Join(id int64) {
	user := c.Session.Id

	err := game.Join(user, id)

	if err != nil {
		c.Set("code", err)
	}
}

func (c *GameController) Game(id int64) {
	g := game.Get(id)

	if g == nil {
		c.Set("code", "not found game")
		return
	}

	c.Set("item", g)
}

func (c *GameController) Map(id int64) {
	g := game.Get(id)

	if g == nil {
		c.Set("code", "not found game")
		return
	}

	c.Set("item", g.Map)
}

// @Post()
func (c *GameController) Command(id int64, cmd string) {
	user := c.Session.Id

	g := game.Get(id)

	if g == nil {
		c.Set("code", "not found game")
		return
	}

	ret := game.Command(g, id, user, cmd, true)

	if ret != nil {
		log.Println("-------------------------")
		log.Println(ret)
		log.Println("-------------------------")
	}
	//c.Set("item", ret)
}
