package api

import (
	"aoi/controllers"
	"aoi/game"
	"aoi/models"
	gamemodel "aoi/models/game"
	"log"
)

type GameController struct {
	controllers.Controller
}

// @Post()
func (c *GameController) Make(item *models.Game) {
	user := c.Session.Id

	conn := c.NewConnection()

	gameManager := models.NewGameManager(conn)
	gameuserManager := models.NewGameuserManager(conn)

	item.Status = gamemodel.StatusReady
	gameManager.Insert(item)

	id := gameManager.GetIdentity()

	var gameuser models.Gameuser
	gameuser.User = user
	gameuser.Game = id

	gameuserManager.Insert(&gameuser)

	game.Make(id)
}

// @Post()
func (c *GameController) Join(id int64) {
	user := c.Session.Id

	g := game.Get(id)

	if g == nil {
		c.Set("code", "not found game")
		return
	}

	conn := c.NewConnection()

	gameManager := models.NewGameManager(conn)
	gameuserManager := models.NewGameuserManager(conn)

	c.Lock()

	game := gameManager.Get(id)

	if game.Status != gamemodel.StatusReady {
		c.Set("code", "status error")

		c.Unlock()
		return
	}

	if gameuserManager.CountByGameUser(id, user) > 0 {
		c.Set("code", "already")

		c.Unlock()
		return
	}

	count := gameuserManager.CountByGame(id)
	if count >= game.Count {
		c.Set("code", "full")

		c.Unlock()
		return
	}

	var item models.Gameuser
	item.User = user
	item.Game = id

	gameuserManager.Insert(&item)
	count++

	g.AddUser(user)

	if count == game.Count {
		game.Status = gamemodel.StatusFaction
		gameManager.Update(game)
	}

	c.Unlock()
}

func (c *GameController) Game(id int64) {
	g := game.Get(id)

	if g == nil {
		game.Make(id)
		g = game.Get(id)
	}

	c.Set("item", g)
}

func (c *GameController) Map(id int64) {
	g := game.Get(id)

	if g == nil {
		game.Make(id)
		g = game.Get(id)
	}

	c.Set("item", g.Map)
}

// @Post()
func (c *GameController) Command(id int64, cmd string) {
	user := c.Session.Id

	g := game.Get(id)

	ret := game.Command(g, id, user, cmd)

	if ret != nil {
		log.Println("-------------------------")
		log.Println(ret)
		log.Println("-------------------------")
	}
	//c.Set("item", ret)
}
