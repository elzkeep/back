package api

import (
	"aoi/controllers"
	"aoi/game"
	"aoi/global"
	"aoi/models"
	gm "aoi/models/game"
	"log"
	"strings"
	"time"
)

type GameController struct {
	controllers.Controller
}

// @Post()
func (c *GameController) Make(item *models.Game) {
	user := c.Session.Id
	item.User = user

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
	game.Lock(id)
	defer game.Unlock(id)

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

func (c *GameController) Replay(id int64, pos int) {
	game.Lock(id)
	defer game.Unlock(id)

	g := game.Get(id)

	log.Println("Replace", id, pos)

	if g == nil {
		game.MakeGame(id)

		g = game.Get(id)
		if g == nil {
			c.Set("code", "not found game")
			return
		}
	}

	replay := game.Replay(id, pos)

	c.Set("item", replay)
}

func (c *GameController) Map(id int64) {
	game.Lock(id)
	defer game.Unlock(id)

	g := game.Get(id)

	if g == nil {
		c.Set("code", "not found game")
		return
	}

	c.Set("item", g.Map)
}

// @Post()
func (c *GameController) Command(id int64, cmd string) {
	game.Lock(id)
	defer game.Unlock(id)

	user := c.Session.Id

	g := game.Get(id)

	if g == nil {
		c.Set("code", "not found game")
		return
	}

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
}

// @Post()
func (c *GameController) Undo(id int64, history int64) {
	game.Lock(id)
	defer game.Unlock(id)

	user := c.Session.Id

	g := game.Get(id)

	if g == nil {
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
	game.Lock(id)
	defer game.Unlock(id)

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

func (c *GameController) Search(status int, page int, pagesize int) {
	user := c.Session.Id

	conn := c.NewConnection()

	manager := models.NewGameManager(conn)
	gameuserManager := models.NewGameuserManager(conn)

	var args []interface{}

	if status == 2 {
		args = append(args, models.Where{Column: "status", Value: 1, Compare: ">"})

		now := time.Now().AddDate(0, 0, -1)
		date := global.GetDatetime(now)
		args = append(args, models.Where{Column: "date", Value: date, Compare: ">="})
	} else if status == 3 {
		args = append(args, models.Where{Column: "status", Value: 1, Compare: "="})
	} else if status == 4 {
		args = append(args, models.Where{Column: "user", Value: user, Compare: "="})
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

func (c *GameController) Delete(item *models.Game) {
	user := c.Session.Id

	db := c.NewConnection()

	conn, _ := db.Begin()
	defer conn.Rollback()

	gameManager := models.NewGameManager(conn)
	gameuserManager := models.NewGameuserManager(conn)
	gametileManager := models.NewGametileManager(conn)
	gamehistoryManager := models.NewGamehistoryManager(conn)
	gameundoManager := models.NewGameundoManager(conn)
	gameundoitemManager := models.NewGameundoitemManager(conn)

	id := item.Id

	old := gameManager.Get(id)

	if old.User != user {
		return
	}

	if old.Status == gm.StatusEnd {
		return
	}

	gameManager.Delete(id)

	gameuserManager.DeleteByGame(id)
	gametileManager.DeleteByGame(id)
	gamehistoryManager.DeleteByGame(id)
	gameundoManager.DeleteByGame(id)
	gameundoitemManager.DeleteByGame(id)

	conn.Commit()
}
