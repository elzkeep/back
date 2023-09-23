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
	apiGroup := r.Group("/api")
	{

		apiGroup.Get("/download/file/:id", func(c *fiber.Ctx) error {
			id_, _ := strconv.ParseInt(c.Params("id"), 10, 64)
			var controller api.DownloadController
			controller.Init(c)
			controller.File(id_)
			controller.Close()
            return nil
		})

		apiGroup.Get("/game/make/:id", func(c *fiber.Ctx) error {
			id_, _ := strconv.ParseInt(c.Params("id"), 10, 64)
			var controller api.GameController
			controller.Init(c)
			controller.Make(id_)
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

		apiGroup.Get("/sms", func(c *fiber.Ctx) error {
			to_ := c.Query("to")
			message_ := c.Query("message")
			var controller api.SmsController
			controller.Init(c)
			controller.Index(to_, message_)
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

	}

	{

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

	}

}