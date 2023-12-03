package router

import (
    "encoding/json"
    "strconv"
    "strings"
	"aoi/controllers/api"
	"aoi/controllers/rest"
    "aoi/models"
    "aoi/models/user"
	"github.com/gofiber/fiber/v2"
)

func getArrayCommal(name string) []int64 {
	values := strings.Split(name, ",")

	var items []int64
	for _, item := range values {
        n, _ := strconv.ParseInt(item, 10, 64)
		items = append(items, n)
	}

	return items
}

func getArrayCommai(name string) []int {
	values := strings.Split(name, ",")

	var items []int
	for _, item := range values {
        n, _ := strconv.Atoi(item)
		items = append(items, n)
	}

	return items
}

func SetRouter(r *fiber.App) {

    r.Get("/api/jwt", func(c *fiber.Ctx) error {
		loginid := c.Query("loginid")
        passwd := c.Query("passwd")
        return c.JSON(JwtAuth(c, loginid, passwd))
	})
	apiGroup := r.Group("/api")
	r.Use(JwtAuthRequired)
	{

		apiGroup.Get("/download/file/:id", func(c *fiber.Ctx) error {
			id_, _ := strconv.ParseInt(c.Params("id"), 10, 64)
			var controller api.DownloadController
			controller.Init(c)
			controller.File(id_)
			controller.Close()
            return nil
		})

		apiGroup.Post("/game/make", func(c *fiber.Ctx) error {
			item_ := &models.Game{}
			c.BodyParser(item_)
			var controller api.GameController
			controller.Init(c)
			controller.Make(item_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/game/join", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller api.GameController
			controller.Init(c)
			controller.Join(id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/game/game/:id", func(c *fiber.Ctx) error {
			id_, _ := strconv.ParseInt(c.Params("id"), 10, 64)
			var controller api.GameController
			controller.Init(c)
			controller.Game(id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/game/replay/:id", func(c *fiber.Ctx) error {
			id_, _ := strconv.ParseInt(c.Params("id"), 10, 64)
			pos_, _ := strconv.Atoi(c.Query("pos"))
			var controller api.GameController
			controller.Init(c)
			controller.Replay(id_, pos_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/game/map/:id", func(c *fiber.Ctx) error {
			id_, _ := strconv.ParseInt(c.Params("id"), 10, 64)
			var controller api.GameController
			controller.Init(c)
			controller.Map(id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/game/command", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var cmd_ string
			if v, flag := results["cmd"]; flag {
				cmd_ = v.(string)
			}
			var controller api.GameController
			controller.Init(c)
			controller.Command(id_, cmd_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/game/undo", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var history_ int64
			if v, flag := results["history"]; flag {
				history_ = int64(v.(float64))
			}
			var controller api.GameController
			controller.Init(c)
			controller.Undo(id_, history_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/game/undoconfirm", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var undo_ int64
			if v, flag := results["undo"]; flag {
				undo_ = int64(v.(float64))
			}
			var status_ int
			if v, flag := results["status"]; flag {
				status_ = int(v.(float64))
			}
			var controller api.GameController
			controller.Init(c)
			controller.Undoconfirm(id_, undo_, status_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/game/search/:status", func(c *fiber.Ctx) error {
			status_, _ := strconv.Atoi(c.Params("status"))
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller api.GameController
			controller.Init(c)
			controller.Search(status_, page_, pagesize_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/game", func(c *fiber.Ctx) error {
			item_ := &models.Game{}
			c.BodyParser(item_)
			var controller api.GameController
			controller.Init(c)
			if item_ != nil {
				controller.Delete(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/gamelist/search/:status", func(c *fiber.Ctx) error {
			status_, _ := strconv.Atoi(c.Params("status"))
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller api.GamelistController
			controller.Init(c)
			controller.Search(status_, page_, pagesize_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/ranklist/elo/:status", func(c *fiber.Ctx) error {
			status_, _ := strconv.Atoi(c.Params("status"))
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller api.RanklistController
			controller.Init(c)
			controller.Elo(status_, page_, pagesize_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/sms", func(c *fiber.Ctx) error {
			to_ := c.Query("to")
			message_ := c.Query("message")
			var controller api.SmsController
			controller.Init(c)
			controller.Index(to_, message_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/statisticscolor/search/:status", func(c *fiber.Ctx) error {
			status_, _ := strconv.Atoi(c.Params("status"))
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller api.StatisticscolorController
			controller.Init(c)
			controller.Search(status_, page_, pagesize_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/upload/index", func(c *fiber.Ctx) error {
			var controller api.UploadController
			controller.Init(c)
			controller.Index()
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/user/elo/:name", func(c *fiber.Ctx) error {
			name_ := c.Params("name")
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller api.UserController
			controller.Init(c)
			controller.Elo(name_, page_, pagesize_)
			controller.Close()
			return c.JSON(controller.Result)
		})

	}

	{

		apiGroup.Get("/game/:id", func(c *fiber.Ctx) error {
			id_, _ := strconv.ParseInt(c.Params("id"), 10, 64)
			var controller rest.GameController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/game", func(c *fiber.Ctx) error {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.GameController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/game", func(c *fiber.Ctx) error {
			item_ := &models.Game{}
			c.BodyParser(item_)
			var controller rest.GameController
			controller.Init(c)
			if item_ != nil {
				controller.Insert(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/game/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Game{}
			c.BodyParser(item_)
			var controller rest.GameController
			controller.Init(c)
			if item_ != nil {
				controller.Insertbatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/game", func(c *fiber.Ctx) error {
			item_ := &models.Game{}
			c.BodyParser(item_)
			var controller rest.GameController
			controller.Init(c)
			if item_ != nil {
				controller.Update(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/gamehistory/:id", func(c *fiber.Ctx) error {
			id_, _ := strconv.ParseInt(c.Params("id"), 10, 64)
			var controller rest.GamehistoryController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/gamehistory", func(c *fiber.Ctx) error {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.GamehistoryController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/gamehistory", func(c *fiber.Ctx) error {
			item_ := &models.Gamehistory{}
			c.BodyParser(item_)
			var controller rest.GamehistoryController
			controller.Init(c)
			if item_ != nil {
				controller.Insert(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/gamehistory/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Gamehistory{}
			c.BodyParser(item_)
			var controller rest.GamehistoryController
			controller.Init(c)
			if item_ != nil {
				controller.Insertbatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/gamehistory", func(c *fiber.Ctx) error {
			item_ := &models.Gamehistory{}
			c.BodyParser(item_)
			var controller rest.GamehistoryController
			controller.Init(c)
			if item_ != nil {
				controller.Update(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/gamehistory", func(c *fiber.Ctx) error {
			item_ := &models.Gamehistory{}
			c.BodyParser(item_)
			var controller rest.GamehistoryController
			controller.Init(c)
			if item_ != nil {
				controller.Delete(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/gamehistory/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Gamehistory{}
			c.BodyParser(item_)
			var controller rest.GamehistoryController
			controller.Init(c)
			if item_ != nil {
				controller.Deletebatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/gamehistory/find/game/:game", func(c *fiber.Ctx) error {
			game_, _ := strconv.ParseInt(c.Params("game"), 10, 64)
			var controller rest.GamehistoryController
			controller.Init(c)
			controller.FindByGame(game_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/gamehistory/count/game/:game", func(c *fiber.Ctx) error {
			game_, _ := strconv.ParseInt(c.Params("game"), 10, 64)
			var controller rest.GamehistoryController
			controller.Init(c)
			controller.CountByGame(game_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/gamehistory/bygame", func(c *fiber.Ctx) error {
			item_ := &models.Gamehistory{}
			c.BodyParser(item_)
			var controller rest.GamehistoryController
			controller.Init(c)
			controller.DeleteByGame(item_.Game)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/gamehistory/round", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var round_ int
			if v, flag := results["round"]; flag {
				round_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.GamehistoryController
			controller.Init(c)
			controller.UpdateRound(round_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/gamehistory/command", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var command_ string
			if v, flag := results["command"]; flag {
				command_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.GamehistoryController
			controller.Init(c)
			controller.UpdateCommand(command_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/gamehistory/vp", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var vp_ int
			if v, flag := results["vp"]; flag {
				vp_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.GamehistoryController
			controller.Init(c)
			controller.UpdateVp(vp_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/gamehistory/user", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var user_ int64
			if v, flag := results["user"]; flag {
				user_ = int64(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.GamehistoryController
			controller.Init(c)
			controller.UpdateUser(user_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/gamehistory/game", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var game_ int64
			if v, flag := results["game"]; flag {
				game_ = int64(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.GamehistoryController
			controller.Init(c)
			controller.UpdateGame(game_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/gamelist/:id", func(c *fiber.Ctx) error {
			id_, _ := strconv.ParseInt(c.Params("id"), 10, 64)
			var controller rest.GamelistController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/gamelist", func(c *fiber.Ctx) error {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.GamelistController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/gamelist/sum", func(c *fiber.Ctx) error {
			var controller rest.GamelistController
			controller.Init(c)
			controller.Sum()
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/gametile/:id", func(c *fiber.Ctx) error {
			id_, _ := strconv.ParseInt(c.Params("id"), 10, 64)
			var controller rest.GametileController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/gametile", func(c *fiber.Ctx) error {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.GametileController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/gametile", func(c *fiber.Ctx) error {
			item_ := &models.Gametile{}
			c.BodyParser(item_)
			var controller rest.GametileController
			controller.Init(c)
			if item_ != nil {
				controller.Insert(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/gametile/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Gametile{}
			c.BodyParser(item_)
			var controller rest.GametileController
			controller.Init(c)
			if item_ != nil {
				controller.Insertbatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/gametile", func(c *fiber.Ctx) error {
			item_ := &models.Gametile{}
			c.BodyParser(item_)
			var controller rest.GametileController
			controller.Init(c)
			if item_ != nil {
				controller.Update(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/gametile", func(c *fiber.Ctx) error {
			item_ := &models.Gametile{}
			c.BodyParser(item_)
			var controller rest.GametileController
			controller.Init(c)
			if item_ != nil {
				controller.Delete(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/gametile/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Gametile{}
			c.BodyParser(item_)
			var controller rest.GametileController
			controller.Init(c)
			if item_ != nil {
				controller.Deletebatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/gametile/find/game/:game", func(c *fiber.Ctx) error {
			game_, _ := strconv.ParseInt(c.Params("game"), 10, 64)
			var controller rest.GametileController
			controller.Init(c)
			controller.FindByGame(game_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/gametile/count/game/:game", func(c *fiber.Ctx) error {
			game_, _ := strconv.ParseInt(c.Params("game"), 10, 64)
			var controller rest.GametileController
			controller.Init(c)
			controller.CountByGame(game_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/gametile/bygame", func(c *fiber.Ctx) error {
			item_ := &models.Gametile{}
			c.BodyParser(item_)
			var controller rest.GametileController
			controller.Init(c)
			controller.DeleteByGame(item_.Game)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/gametile/type", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var type_ int
			if v, flag := results["type"]; flag {
				type_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.GametileController
			controller.Init(c)
			controller.UpdateType(type_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/gametile/number", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var number_ int
			if v, flag := results["number"]; flag {
				number_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.GametileController
			controller.Init(c)
			controller.UpdateNumber(number_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/gametile/order", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var order_ int
			if v, flag := results["order"]; flag {
				order_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.GametileController
			controller.Init(c)
			controller.UpdateOrder(order_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/gametile/game", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var game_ int64
			if v, flag := results["game"]; flag {
				game_ = int64(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.GametileController
			controller.Init(c)
			controller.UpdateGame(game_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/gameundo/:id", func(c *fiber.Ctx) error {
			id_, _ := strconv.ParseInt(c.Params("id"), 10, 64)
			var controller rest.GameundoController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/gameundo", func(c *fiber.Ctx) error {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.GameundoController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/gameundo", func(c *fiber.Ctx) error {
			item_ := &models.Gameundo{}
			c.BodyParser(item_)
			var controller rest.GameundoController
			controller.Init(c)
			if item_ != nil {
				controller.Insert(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/gameundo/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Gameundo{}
			c.BodyParser(item_)
			var controller rest.GameundoController
			controller.Init(c)
			if item_ != nil {
				controller.Insertbatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/gameundo", func(c *fiber.Ctx) error {
			item_ := &models.Gameundo{}
			c.BodyParser(item_)
			var controller rest.GameundoController
			controller.Init(c)
			if item_ != nil {
				controller.Update(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/gameundo", func(c *fiber.Ctx) error {
			item_ := &models.Gameundo{}
			c.BodyParser(item_)
			var controller rest.GameundoController
			controller.Init(c)
			if item_ != nil {
				controller.Delete(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/gameundo/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Gameundo{}
			c.BodyParser(item_)
			var controller rest.GameundoController
			controller.Init(c)
			if item_ != nil {
				controller.Deletebatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/gameundo/find/game/:game", func(c *fiber.Ctx) error {
			game_, _ := strconv.ParseInt(c.Params("game"), 10, 64)
			var controller rest.GameundoController
			controller.Init(c)
			controller.FindByGame(game_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/gameundo/count/game/:game", func(c *fiber.Ctx) error {
			game_, _ := strconv.ParseInt(c.Params("game"), 10, 64)
			var controller rest.GameundoController
			controller.Init(c)
			controller.CountByGame(game_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/gameundo/bygame", func(c *fiber.Ctx) error {
			item_ := &models.Gameundo{}
			c.BodyParser(item_)
			var controller rest.GameundoController
			controller.Init(c)
			controller.DeleteByGame(item_.Game)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/gameundo/status", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var status_ int
			if v, flag := results["status"]; flag {
				status_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.GameundoController
			controller.Init(c)
			controller.UpdateStatus(status_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/gameundo/gamehistory", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var gamehistory_ int64
			if v, flag := results["gamehistory"]; flag {
				gamehistory_ = int64(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.GameundoController
			controller.Init(c)
			controller.UpdateGamehistory(gamehistory_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/gameundo/game", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var game_ int64
			if v, flag := results["game"]; flag {
				game_ = int64(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.GameundoController
			controller.Init(c)
			controller.UpdateGame(game_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/gameundo/user", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var user_ int64
			if v, flag := results["user"]; flag {
				user_ = int64(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.GameundoController
			controller.Init(c)
			controller.UpdateUser(user_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/gameundoitem/:id", func(c *fiber.Ctx) error {
			id_, _ := strconv.ParseInt(c.Params("id"), 10, 64)
			var controller rest.GameundoitemController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/gameundoitem", func(c *fiber.Ctx) error {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.GameundoitemController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/gameundoitem", func(c *fiber.Ctx) error {
			item_ := &models.Gameundoitem{}
			c.BodyParser(item_)
			var controller rest.GameundoitemController
			controller.Init(c)
			if item_ != nil {
				controller.Insert(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/gameundoitem/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Gameundoitem{}
			c.BodyParser(item_)
			var controller rest.GameundoitemController
			controller.Init(c)
			if item_ != nil {
				controller.Insertbatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/gameundoitem", func(c *fiber.Ctx) error {
			item_ := &models.Gameundoitem{}
			c.BodyParser(item_)
			var controller rest.GameundoitemController
			controller.Init(c)
			if item_ != nil {
				controller.Update(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/gameundoitem", func(c *fiber.Ctx) error {
			item_ := &models.Gameundoitem{}
			c.BodyParser(item_)
			var controller rest.GameundoitemController
			controller.Init(c)
			if item_ != nil {
				controller.Delete(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/gameundoitem/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Gameundoitem{}
			c.BodyParser(item_)
			var controller rest.GameundoitemController
			controller.Init(c)
			if item_ != nil {
				controller.Deletebatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/gameundoitem/find/game/:game", func(c *fiber.Ctx) error {
			game_, _ := strconv.ParseInt(c.Params("game"), 10, 64)
			var controller rest.GameundoitemController
			controller.Init(c)
			controller.FindByGame(game_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/gameundoitem/count/game/:game", func(c *fiber.Ctx) error {
			game_, _ := strconv.ParseInt(c.Params("game"), 10, 64)
			var controller rest.GameundoitemController
			controller.Init(c)
			controller.CountByGame(game_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/gameundoitem/bygame", func(c *fiber.Ctx) error {
			item_ := &models.Gameundoitem{}
			c.BodyParser(item_)
			var controller rest.GameundoitemController
			controller.Init(c)
			controller.DeleteByGame(item_.Game)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/gameundoitem/status", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var status_ int
			if v, flag := results["status"]; flag {
				status_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.GameundoitemController
			controller.Init(c)
			controller.UpdateStatus(status_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/gameundoitem/gameundo", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var gameundo_ int64
			if v, flag := results["gameundo"]; flag {
				gameundo_ = int64(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.GameundoitemController
			controller.Init(c)
			controller.UpdateGameundo(gameundo_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/gameundoitem/game", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var game_ int64
			if v, flag := results["game"]; flag {
				game_ = int64(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.GameundoitemController
			controller.Init(c)
			controller.UpdateGame(game_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/gameundoitem/user", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var user_ int64
			if v, flag := results["user"]; flag {
				user_ = int64(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.GameundoitemController
			controller.Init(c)
			controller.UpdateUser(user_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/gameuser/:id", func(c *fiber.Ctx) error {
			id_, _ := strconv.ParseInt(c.Params("id"), 10, 64)
			var controller rest.GameuserController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/gameuser", func(c *fiber.Ctx) error {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.GameuserController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/gameuser", func(c *fiber.Ctx) error {
			item_ := &models.Gameuser{}
			c.BodyParser(item_)
			var controller rest.GameuserController
			controller.Init(c)
			if item_ != nil {
				controller.Insert(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/gameuser/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Gameuser{}
			c.BodyParser(item_)
			var controller rest.GameuserController
			controller.Init(c)
			if item_ != nil {
				controller.Insertbatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/gameuser", func(c *fiber.Ctx) error {
			item_ := &models.Gameuser{}
			c.BodyParser(item_)
			var controller rest.GameuserController
			controller.Init(c)
			if item_ != nil {
				controller.Update(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/gameuser", func(c *fiber.Ctx) error {
			item_ := &models.Gameuser{}
			c.BodyParser(item_)
			var controller rest.GameuserController
			controller.Init(c)
			if item_ != nil {
				controller.Delete(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/gameuser/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Gameuser{}
			c.BodyParser(item_)
			var controller rest.GameuserController
			controller.Init(c)
			if item_ != nil {
				controller.Deletebatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/gameuser/find/game/:game", func(c *fiber.Ctx) error {
			game_, _ := strconv.ParseInt(c.Params("game"), 10, 64)
			var controller rest.GameuserController
			controller.Init(c)
			controller.FindByGame(game_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/gameuser/count/game/:game", func(c *fiber.Ctx) error {
			game_, _ := strconv.ParseInt(c.Params("game"), 10, 64)
			var controller rest.GameuserController
			controller.Init(c)
			controller.CountByGame(game_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/gameuser/get/gameuser/:game", func(c *fiber.Ctx) error {
			game_, _ := strconv.ParseInt(c.Params("game"), 10, 64)
			user_, _ := strconv.ParseInt(c.Query("user"), 10, 64)
			var controller rest.GameuserController
			controller.Init(c)
			controller.GetByGameUser(game_, user_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/gameuser/count/gameuser/:game", func(c *fiber.Ctx) error {
			game_, _ := strconv.ParseInt(c.Params("game"), 10, 64)
			user_, _ := strconv.ParseInt(c.Query("user"), 10, 64)
			var controller rest.GameuserController
			controller.Init(c)
			controller.CountByGameUser(game_, user_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/gameuser/bygame", func(c *fiber.Ctx) error {
			item_ := &models.Gameuser{}
			c.BodyParser(item_)
			var controller rest.GameuserController
			controller.Init(c)
			controller.DeleteByGame(item_.Game)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/gameuser/order", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var order_ int
			if v, flag := results["order"]; flag {
				order_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.GameuserController
			controller.Init(c)
			controller.UpdateOrder(order_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/gameuser/faction", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var faction_ int
			if v, flag := results["faction"]; flag {
				faction_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.GameuserController
			controller.Init(c)
			controller.UpdateFaction(faction_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/gameuser/color", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var color_ int
			if v, flag := results["color"]; flag {
				color_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.GameuserController
			controller.Init(c)
			controller.UpdateColor(color_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/gameuser/score", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var score_ int
			if v, flag := results["score"]; flag {
				score_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.GameuserController
			controller.Init(c)
			controller.UpdateScore(score_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/gameuser/rank", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var rank_ int
			if v, flag := results["rank"]; flag {
				rank_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.GameuserController
			controller.Init(c)
			controller.UpdateRank(rank_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/gameuser/elo", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var elo_ models.Double
			elo__ref := &elo_
			c.BodyParser(elo__ref)
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.GameuserController
			controller.Init(c)
			controller.UpdateElo(elo_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/gameuser/user", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var user_ int64
			if v, flag := results["user"]; flag {
				user_ = int64(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.GameuserController
			controller.Init(c)
			controller.UpdateUser(user_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/gameuser/game", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var game_ int64
			if v, flag := results["game"]; flag {
				game_ = int64(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.GameuserController
			controller.Init(c)
			controller.UpdateGame(game_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/map/:id", func(c *fiber.Ctx) error {
			id_, _ := strconv.ParseInt(c.Params("id"), 10, 64)
			var controller rest.MapController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/map", func(c *fiber.Ctx) error {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.MapController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/map", func(c *fiber.Ctx) error {
			item_ := &models.Map{}
			c.BodyParser(item_)
			var controller rest.MapController
			controller.Init(c)
			if item_ != nil {
				controller.Insert(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/map/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Map{}
			c.BodyParser(item_)
			var controller rest.MapController
			controller.Init(c)
			if item_ != nil {
				controller.Insertbatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/map", func(c *fiber.Ctx) error {
			item_ := &models.Map{}
			c.BodyParser(item_)
			var controller rest.MapController
			controller.Init(c)
			if item_ != nil {
				controller.Update(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/map", func(c *fiber.Ctx) error {
			item_ := &models.Map{}
			c.BodyParser(item_)
			var controller rest.MapController
			controller.Init(c)
			if item_ != nil {
				controller.Delete(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/map/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Map{}
			c.BodyParser(item_)
			var controller rest.MapController
			controller.Init(c)
			if item_ != nil {
				controller.Deletebatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/map/name", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var name_ string
			if v, flag := results["name"]; flag {
				name_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.MapController
			controller.Init(c)
			controller.UpdateName(name_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/map/content", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var content_ string
			if v, flag := results["content"]; flag {
				content_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.MapController
			controller.Init(c)
			controller.UpdateContent(content_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/map/order", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var order_ int
			if v, flag := results["order"]; flag {
				order_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.MapController
			controller.Init(c)
			controller.UpdateOrder(order_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/notice/:id", func(c *fiber.Ctx) error {
			id_, _ := strconv.ParseInt(c.Params("id"), 10, 64)
			var controller rest.NoticeController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/notice", func(c *fiber.Ctx) error {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.NoticeController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/notice", func(c *fiber.Ctx) error {
			item_ := &models.Notice{}
			c.BodyParser(item_)
			var controller rest.NoticeController
			controller.Init(c)
			if item_ != nil {
				controller.Insert(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/notice/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Notice{}
			c.BodyParser(item_)
			var controller rest.NoticeController
			controller.Init(c)
			if item_ != nil {
				controller.Insertbatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/notice", func(c *fiber.Ctx) error {
			item_ := &models.Notice{}
			c.BodyParser(item_)
			var controller rest.NoticeController
			controller.Init(c)
			if item_ != nil {
				controller.Update(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/notice", func(c *fiber.Ctx) error {
			item_ := &models.Notice{}
			c.BodyParser(item_)
			var controller rest.NoticeController
			controller.Init(c)
			if item_ != nil {
				controller.Delete(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/notice/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Notice{}
			c.BodyParser(item_)
			var controller rest.NoticeController
			controller.Init(c)
			if item_ != nil {
				controller.Deletebatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/notice/content", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var content_ string
			if v, flag := results["content"]; flag {
				content_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.NoticeController
			controller.Init(c)
			controller.UpdateContent(content_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/notice/status", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var status_ int
			if v, flag := results["status"]; flag {
				status_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.NoticeController
			controller.Init(c)
			controller.UpdateStatus(status_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/ranklist/:id", func(c *fiber.Ctx) error {
			id_, _ := strconv.ParseInt(c.Params("id"), 10, 64)
			var controller rest.RanklistController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/ranklist", func(c *fiber.Ctx) error {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.RanklistController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/ranklist/sum", func(c *fiber.Ctx) error {
			var controller rest.RanklistController
			controller.Init(c)
			controller.Sum()
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/smsauth/:id", func(c *fiber.Ctx) error {
			id_, _ := strconv.ParseInt(c.Params("id"), 10, 64)
			var controller rest.SmsauthController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/smsauth", func(c *fiber.Ctx) error {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.SmsauthController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/smsauth", func(c *fiber.Ctx) error {
			item_ := &models.Smsauth{}
			c.BodyParser(item_)
			var controller rest.SmsauthController
			controller.Init(c)
			if item_ != nil {
				controller.Insert(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/smsauth/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Smsauth{}
			c.BodyParser(item_)
			var controller rest.SmsauthController
			controller.Init(c)
			if item_ != nil {
				controller.Insertbatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/smsauth", func(c *fiber.Ctx) error {
			item_ := &models.Smsauth{}
			c.BodyParser(item_)
			var controller rest.SmsauthController
			controller.Init(c)
			if item_ != nil {
				controller.Update(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/smsauth", func(c *fiber.Ctx) error {
			item_ := &models.Smsauth{}
			c.BodyParser(item_)
			var controller rest.SmsauthController
			controller.Init(c)
			if item_ != nil {
				controller.Delete(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/smsauth/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Smsauth{}
			c.BodyParser(item_)
			var controller rest.SmsauthController
			controller.Init(c)
			if item_ != nil {
				controller.Deletebatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/smsauth/get/hpnumber/:hp", func(c *fiber.Ctx) error {
			hp_ := c.Params("hp")
			number_ := c.Query("number")
			var controller rest.SmsauthController
			controller.Init(c)
			controller.GetByHpNumber(hp_, number_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/smsauth/byhp", func(c *fiber.Ctx) error {
			item_ := &models.Smsauth{}
			c.BodyParser(item_)
			var controller rest.SmsauthController
			controller.Init(c)
			controller.DeleteByHp(item_.Hp)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/smsauth/hp", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var hp_ string
			if v, flag := results["hp"]; flag {
				hp_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.SmsauthController
			controller.Init(c)
			controller.UpdateHp(hp_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/smsauth/number", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var number_ string
			if v, flag := results["number"]; flag {
				number_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.SmsauthController
			controller.Init(c)
			controller.UpdateNumber(number_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/statistics/:id", func(c *fiber.Ctx) error {
			id_, _ := strconv.ParseInt(c.Params("id"), 10, 64)
			var controller rest.StatisticsController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/statistics", func(c *fiber.Ctx) error {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.StatisticsController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/statistics/sum", func(c *fiber.Ctx) error {
			var controller rest.StatisticsController
			controller.Init(c)
			controller.Sum()
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/statisticscolor/:id", func(c *fiber.Ctx) error {
			id_, _ := strconv.ParseInt(c.Params("id"), 10, 64)
			var controller rest.StatisticscolorController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/statisticscolor", func(c *fiber.Ctx) error {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.StatisticscolorController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/statisticscolor/sum", func(c *fiber.Ctx) error {
			var controller rest.StatisticscolorController
			controller.Init(c)
			controller.Sum()
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/statisticsfaction/:id", func(c *fiber.Ctx) error {
			id_, _ := strconv.ParseInt(c.Params("id"), 10, 64)
			var controller rest.StatisticsfactionController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/statisticsfaction", func(c *fiber.Ctx) error {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.StatisticsfactionController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/statisticsfaction/sum", func(c *fiber.Ctx) error {
			var controller rest.StatisticsfactionController
			controller.Init(c)
			controller.Sum()
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/token/:id", func(c *fiber.Ctx) error {
			id_, _ := strconv.ParseInt(c.Params("id"), 10, 64)
			var controller rest.TokenController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/token", func(c *fiber.Ctx) error {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.TokenController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/token", func(c *fiber.Ctx) error {
			item_ := &models.Token{}
			c.BodyParser(item_)
			var controller rest.TokenController
			controller.Init(c)
			if item_ != nil {
				controller.Insert(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/token/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Token{}
			c.BodyParser(item_)
			var controller rest.TokenController
			controller.Init(c)
			if item_ != nil {
				controller.Insertbatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/token", func(c *fiber.Ctx) error {
			item_ := &models.Token{}
			c.BodyParser(item_)
			var controller rest.TokenController
			controller.Init(c)
			if item_ != nil {
				controller.Update(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/token", func(c *fiber.Ctx) error {
			item_ := &models.Token{}
			c.BodyParser(item_)
			var controller rest.TokenController
			controller.Init(c)
			if item_ != nil {
				controller.Delete(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/token/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Token{}
			c.BodyParser(item_)
			var controller rest.TokenController
			controller.Init(c)
			if item_ != nil {
				controller.Deletebatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/token/get/user/:user", func(c *fiber.Ctx) error {
			user_, _ := strconv.ParseInt(c.Params("user"), 10, 64)
			var controller rest.TokenController
			controller.Init(c)
			controller.GetByUser(user_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/token/user", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var user_ int64
			if v, flag := results["user"]; flag {
				user_ = int64(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.TokenController
			controller.Init(c)
			controller.UpdateUser(user_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/token/token", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var token_ string
			if v, flag := results["token"]; flag {
				token_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.TokenController
			controller.Init(c)
			controller.UpdateToken(token_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/token/status", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var status_ int
			if v, flag := results["status"]; flag {
				status_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.TokenController
			controller.Init(c)
			controller.UpdateStatus(status_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/user/:id", func(c *fiber.Ctx) error {
			id_, _ := strconv.ParseInt(c.Params("id"), 10, 64)
			var controller rest.UserController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/user", func(c *fiber.Ctx) error {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.UserController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/user", func(c *fiber.Ctx) error {
			item_ := &models.User{}
			c.BodyParser(item_)
			var controller rest.UserController
			controller.Init(c)
			if item_ != nil {
				controller.Insert(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/user/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.User{}
			c.BodyParser(item_)
			var controller rest.UserController
			controller.Init(c)
			if item_ != nil {
				controller.Insertbatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/user", func(c *fiber.Ctx) error {
			item_ := &models.User{}
			c.BodyParser(item_)
			var controller rest.UserController
			controller.Init(c)
			if item_ != nil {
				controller.Update(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/user", func(c *fiber.Ctx) error {
			item_ := &models.User{}
			c.BodyParser(item_)
			var controller rest.UserController
			controller.Init(c)
			if item_ != nil {
				controller.Delete(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/user/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.User{}
			c.BodyParser(item_)
			var controller rest.UserController
			controller.Init(c)
			if item_ != nil {
				controller.Deletebatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/user/get/email/:email", func(c *fiber.Ctx) error {
			email_ := c.Params("email")
			var controller rest.UserController
			controller.Init(c)
			controller.GetByEmail(email_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/user/count/email/:email", func(c *fiber.Ctx) error {
			email_ := c.Params("email")
			var controller rest.UserController
			controller.Init(c)
			controller.CountByEmail(email_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/user/find/level/:level", func(c *fiber.Ctx) error {
			var level_ user.Level
			level__, _ := strconv.Atoi(c.Params("level"))
			level_ = user.Level(level__)
			var controller rest.UserController
			controller.Init(c)
			controller.FindByLevel(level_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/user/imagebyid", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var image_ string
			if v, flag := results["image"]; flag {
				image_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.UserController
			controller.Init(c)
			controller.UpdateImageById(image_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/user/email", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var email_ string
			if v, flag := results["email"]; flag {
				email_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.UserController
			controller.Init(c)
			controller.UpdateEmail(email_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/user/passwd", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var passwd_ string
			if v, flag := results["passwd"]; flag {
				passwd_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.UserController
			controller.Init(c)
			controller.UpdatePasswd(passwd_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/user/name", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var name_ string
			if v, flag := results["name"]; flag {
				name_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.UserController
			controller.Init(c)
			controller.UpdateName(name_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/user/level", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var level_ int
			if v, flag := results["level"]; flag {
				level_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.UserController
			controller.Init(c)
			controller.UpdateLevel(level_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/user/status", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var status_ int
			if v, flag := results["status"]; flag {
				status_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.UserController
			controller.Init(c)
			controller.UpdateStatus(status_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/user/elo", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var elo_ models.Double
			elo__ref := &elo_
			c.BodyParser(elo__ref)
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.UserController
			controller.Init(c)
			controller.UpdateElo(elo_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/user/count", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var count_ int
			if v, flag := results["count"]; flag {
				count_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.UserController
			controller.Init(c)
			controller.UpdateCount(count_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/user/image", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var image_ string
			if v, flag := results["image"]; flag {
				image_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.UserController
			controller.Init(c)
			controller.UpdateImage(image_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/user/profile", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var profile_ string
			if v, flag := results["profile"]; flag {
				profile_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.UserController
			controller.Init(c)
			controller.UpdateProfile(profile_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/user/sum", func(c *fiber.Ctx) error {
			var controller rest.UserController
			controller.Init(c)
			controller.Sum()
			controller.Close()
			return c.JSON(controller.Result)
		})

	}

}