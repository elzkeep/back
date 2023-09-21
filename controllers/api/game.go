package api

import (
	"aoi/controllers"
	"aoi/game"
	"log"
)

type GameController struct {
	controllers.Controller
}

func (c *GameController) Make(id int64) {
	game.Make(id)
}

func (c *GameController) Map(id int64) {
	g := game.Get(id)

	if g == nil {
		game.Make(id)
		g = game.Get(id)
	}
	c.Set("item", g.Map)

	g.Map.Print()
}

// @Post()
func (c *GameController) Command(id int64, cmd string) {
	g := game.Get(id)

	if g == nil {
		game.Make(id)
		g = game.Get(id)
	}

	ret := game.Command(g, cmd)

	log.Println("command ret")
	log.Println("-------------------------")
	log.Println(ret)
	log.Println("-------------------------")
	//c.Set("item", ret)
}
