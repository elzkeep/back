package api

import (
	"aoi/controllers"
	"aoi/game"
	"aoi/global"
	"aoi/models"
	"log"
	"strings"
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
	game.Lock(id)
	defer game.Unlock(id)

	user := c.Session.Id

	err := game.Join(user, id)

	if err != nil {
		c.Set("code", err)
	}
}

func (c *GameController) Game(id int64) {
	g := game.Get(id)

	if g == nil {
		game.MakeGame(id)

		g = game.Get(id)
		if g == nil {
			c.Set("code", "not found game")
			return
		}
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

	g.Lock()

	cmds := strings.Split(cmd, ", ")

	for _, v := range cmds {
		ret := game.Command(g, id, user, v, true, 0)

		if ret != nil {
			log.Println("-------------------------")
			log.Println(ret)
			log.Println("-------------------------")
			c.Error(ret.Error())
		}
	}

	g.Unlock()
}

// @Post()
func (c *GameController) Undo(id int64, history int64) {
	log.Println("undo call", id, history)
	user := c.Session.Id

	g := game.Get(id)

	if g == nil {
		log.Println("not found")
		c.Set("code", "not found game")
		return
	}

	ret := game.Undo(g, id, history, user)

	if ret != nil {
		log.Println("-------------------------")
		log.Println(ret)
		log.Println("-------------------------")
		c.Error(ret.Error())
	} else {
		msg := global.Notify{Id: id, Title: "undo"}
		global.SendNotify(msg)
	}
}

// @Post()
func (c *GameController) Undoconfirm(id int64, undo int64, status int) {
	user := c.Session.Id

	g := game.Get(id)

	ret := game.UndoConfirm(g, id, undo, user, status)

	if ret != nil {
		log.Println("-------------------------")
		log.Println(ret)
		log.Println("-------------------------")
		c.Error(ret.Error())
	} else {
		msg := global.Notify{Id: id, Title: "undo"}
		global.SendNotify(msg)
	}
}
