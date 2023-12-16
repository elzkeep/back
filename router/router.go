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

		apiGroup.Get("/company/:id", func(c *fiber.Ctx) error {
			id_, _ := strconv.ParseInt(c.Params("id"), 10, 64)
			var controller rest.CompanyController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/company", func(c *fiber.Ctx) error {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.CompanyController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/company", func(c *fiber.Ctx) error {
			item_ := &models.Company{}
			c.BodyParser(item_)
			var controller rest.CompanyController
			controller.Init(c)
			if item_ != nil {
				controller.Insert(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/company/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Company{}
			c.BodyParser(item_)
			var controller rest.CompanyController
			controller.Init(c)
			if item_ != nil {
				controller.Insertbatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/company", func(c *fiber.Ctx) error {
			item_ := &models.Company{}
			c.BodyParser(item_)
			var controller rest.CompanyController
			controller.Init(c)
			if item_ != nil {
				controller.Update(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/company", func(c *fiber.Ctx) error {
			item_ := &models.Company{}
			c.BodyParser(item_)
			var controller rest.CompanyController
			controller.Init(c)
			if item_ != nil {
				controller.Delete(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/company/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Company{}
			c.BodyParser(item_)
			var controller rest.CompanyController
			controller.Init(c)
			if item_ != nil {
				controller.Deletebatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/company/name", func(c *fiber.Ctx) error {
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
			var controller rest.CompanyController
			controller.Init(c)
			controller.UpdateName(name_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/company/companyno", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var companyno_ string
			if v, flag := results["companyno"]; flag {
				companyno_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CompanyController
			controller.Init(c)
			controller.UpdateCompanyno(companyno_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/company/ceo", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var ceo_ string
			if v, flag := results["ceo"]; flag {
				ceo_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CompanyController
			controller.Init(c)
			controller.UpdateCeo(ceo_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/company/address", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var address_ string
			if v, flag := results["address"]; flag {
				address_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CompanyController
			controller.Init(c)
			controller.UpdateAddress(address_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/company/addressetc", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var addressetc_ string
			if v, flag := results["addressetc"]; flag {
				addressetc_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CompanyController
			controller.Init(c)
			controller.UpdateAddressetc(addressetc_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/company/type", func(c *fiber.Ctx) error {
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
			var controller rest.CompanyController
			controller.Init(c)
			controller.UpdateType(type_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/company/checkdate", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var checkdate_ int
			if v, flag := results["checkdate"]; flag {
				checkdate_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CompanyController
			controller.Init(c)
			controller.UpdateCheckdate(checkdate_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/company/managername", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var managername_ string
			if v, flag := results["managername"]; flag {
				managername_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CompanyController
			controller.Init(c)
			controller.UpdateManagername(managername_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/company/managertel", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var managertel_ string
			if v, flag := results["managertel"]; flag {
				managertel_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CompanyController
			controller.Init(c)
			controller.UpdateManagertel(managertel_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/company/manageremail", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var manageremail_ string
			if v, flag := results["manageremail"]; flag {
				manageremail_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CompanyController
			controller.Init(c)
			controller.UpdateManageremail(manageremail_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/company/contractstartdate", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var contractstartdate_ string
			if v, flag := results["contractstartdate"]; flag {
				contractstartdate_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CompanyController
			controller.Init(c)
			controller.UpdateContractstartdate(contractstartdate_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/company/contractenddate", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var contractenddate_ string
			if v, flag := results["contractenddate"]; flag {
				contractenddate_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CompanyController
			controller.Init(c)
			controller.UpdateContractenddate(contractenddate_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/company/contractprice", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var contractprice_ int
			if v, flag := results["contractprice"]; flag {
				contractprice_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CompanyController
			controller.Init(c)
			controller.UpdateContractprice(contractprice_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/company/billingdate", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var billingdate_ int
			if v, flag := results["billingdate"]; flag {
				billingdate_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CompanyController
			controller.Init(c)
			controller.UpdateBillingdate(billingdate_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/company/billingname", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var billingname_ string
			if v, flag := results["billingname"]; flag {
				billingname_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CompanyController
			controller.Init(c)
			controller.UpdateBillingname(billingname_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/company/billingtel", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var billingtel_ string
			if v, flag := results["billingtel"]; flag {
				billingtel_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CompanyController
			controller.Init(c)
			controller.UpdateBillingtel(billingtel_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/company/billingemail", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var billingemail_ string
			if v, flag := results["billingemail"]; flag {
				billingemail_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CompanyController
			controller.Init(c)
			controller.UpdateBillingemail(billingemail_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/company/status", func(c *fiber.Ctx) error {
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
			var controller rest.CompanyController
			controller.Init(c)
			controller.UpdateStatus(status_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/data/:id", func(c *fiber.Ctx) error {
			id_, _ := strconv.ParseInt(c.Params("id"), 10, 64)
			var controller rest.DataController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/data", func(c *fiber.Ctx) error {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.DataController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/data", func(c *fiber.Ctx) error {
			item_ := &models.Data{}
			c.BodyParser(item_)
			var controller rest.DataController
			controller.Init(c)
			if item_ != nil {
				controller.Insert(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/data/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Data{}
			c.BodyParser(item_)
			var controller rest.DataController
			controller.Init(c)
			if item_ != nil {
				controller.Insertbatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/data", func(c *fiber.Ctx) error {
			item_ := &models.Data{}
			c.BodyParser(item_)
			var controller rest.DataController
			controller.Init(c)
			if item_ != nil {
				controller.Update(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/data", func(c *fiber.Ctx) error {
			item_ := &models.Data{}
			c.BodyParser(item_)
			var controller rest.DataController
			controller.Init(c)
			if item_ != nil {
				controller.Delete(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/data/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Data{}
			c.BodyParser(item_)
			var controller rest.DataController
			controller.Init(c)
			if item_ != nil {
				controller.Deletebatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/data/title", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var title_ string
			if v, flag := results["title"]; flag {
				title_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.DataController
			controller.Init(c)
			controller.UpdateTitle(title_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/data/type", func(c *fiber.Ctx) error {
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
			var controller rest.DataController
			controller.Init(c)
			controller.UpdateType(type_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/data/category", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var category_ int
			if v, flag := results["category"]; flag {
				category_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.DataController
			controller.Init(c)
			controller.UpdateCategory(category_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/data/order", func(c *fiber.Ctx) error {
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
			var controller rest.DataController
			controller.Init(c)
			controller.UpdateOrder(order_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/data/report", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var report_ int64
			if v, flag := results["report"]; flag {
				report_ = int64(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.DataController
			controller.Init(c)
			controller.UpdateReport(report_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/data/company", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var company_ int64
			if v, flag := results["company"]; flag {
				company_ = int64(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.DataController
			controller.Init(c)
			controller.UpdateCompany(company_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

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

		apiGroup.Delete("/game", func(c *fiber.Ctx) error {
			item_ := &models.Game{}
			c.BodyParser(item_)
			var controller rest.GameController
			controller.Init(c)
			if item_ != nil {
				controller.Delete(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/game/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Game{}
			c.BodyParser(item_)
			var controller rest.GameController
			controller.Init(c)
			if item_ != nil {
				controller.Deletebatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/game/name", func(c *fiber.Ctx) error {
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
			var controller rest.GameController
			controller.Init(c)
			controller.UpdateName(name_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/game/count", func(c *fiber.Ctx) error {
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
			var controller rest.GameController
			controller.Init(c)
			controller.UpdateCount(count_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/game/join", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var join_ int
			if v, flag := results["join"]; flag {
				join_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.GameController
			controller.Init(c)
			controller.UpdateJoin(join_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/game/map", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var mapid_ int64
			if v, flag := results["mapid"]; flag {
				mapid_ = int64(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.GameController
			controller.Init(c)
			controller.UpdateMap(mapid_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/game/illusionists", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var illusionists_ int
			if v, flag := results["illusionists"]; flag {
				illusionists_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.GameController
			controller.Init(c)
			controller.UpdateIllusionists(illusionists_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/game/type", func(c *fiber.Ctx) error {
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
			var controller rest.GameController
			controller.Init(c)
			controller.UpdateType(type_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/game/status", func(c *fiber.Ctx) error {
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
			var controller rest.GameController
			controller.Init(c)
			controller.UpdateStatus(status_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/game/enddate", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var enddate_ string
			if v, flag := results["enddate"]; flag {
				enddate_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.GameController
			controller.Init(c)
			controller.UpdateEnddate(enddate_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/game/user", func(c *fiber.Ctx) error {
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
			var controller rest.GameController
			controller.Init(c)
			controller.UpdateUser(user_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/game/sum", func(c *fiber.Ctx) error {
			var controller rest.GameController
			controller.Init(c)
			controller.Sum()
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

		apiGroup.Get("/highscore/:id", func(c *fiber.Ctx) error {
			id_, _ := strconv.ParseInt(c.Params("id"), 10, 64)
			var controller rest.HighscoreController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/highscore", func(c *fiber.Ctx) error {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.HighscoreController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/highscore/sum", func(c *fiber.Ctx) error {
			var controller rest.HighscoreController
			controller.Init(c)
			controller.Sum()
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/item/:id", func(c *fiber.Ctx) error {
			id_, _ := strconv.ParseInt(c.Params("id"), 10, 64)
			var controller rest.ItemController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/item", func(c *fiber.Ctx) error {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.ItemController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/item", func(c *fiber.Ctx) error {
			item_ := &models.Item{}
			c.BodyParser(item_)
			var controller rest.ItemController
			controller.Init(c)
			if item_ != nil {
				controller.Insert(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/item/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Item{}
			c.BodyParser(item_)
			var controller rest.ItemController
			controller.Init(c)
			if item_ != nil {
				controller.Insertbatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/item", func(c *fiber.Ctx) error {
			item_ := &models.Item{}
			c.BodyParser(item_)
			var controller rest.ItemController
			controller.Init(c)
			if item_ != nil {
				controller.Update(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/item", func(c *fiber.Ctx) error {
			item_ := &models.Item{}
			c.BodyParser(item_)
			var controller rest.ItemController
			controller.Init(c)
			if item_ != nil {
				controller.Delete(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/item/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Item{}
			c.BodyParser(item_)
			var controller rest.ItemController
			controller.Init(c)
			if item_ != nil {
				controller.Deletebatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/item/title", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var title_ string
			if v, flag := results["title"]; flag {
				title_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.ItemController
			controller.Init(c)
			controller.UpdateTitle(title_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/item/type", func(c *fiber.Ctx) error {
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
			var controller rest.ItemController
			controller.Init(c)
			controller.UpdateType(type_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/item/value1", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var value1_ int
			if v, flag := results["value1"]; flag {
				value1_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.ItemController
			controller.Init(c)
			controller.UpdateValue1(value1_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/item/value2", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var value2_ int
			if v, flag := results["value2"]; flag {
				value2_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.ItemController
			controller.Init(c)
			controller.UpdateValue2(value2_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/item/value3", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var value3_ int
			if v, flag := results["value3"]; flag {
				value3_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.ItemController
			controller.Init(c)
			controller.UpdateValue3(value3_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/item/value4", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var value4_ int
			if v, flag := results["value4"]; flag {
				value4_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.ItemController
			controller.Init(c)
			controller.UpdateValue4(value4_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/item/value5", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var value5_ int
			if v, flag := results["value5"]; flag {
				value5_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.ItemController
			controller.Init(c)
			controller.UpdateValue5(value5_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/item/value6", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var value6_ int
			if v, flag := results["value6"]; flag {
				value6_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.ItemController
			controller.Init(c)
			controller.UpdateValue6(value6_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/item/value7", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var value7_ int
			if v, flag := results["value7"]; flag {
				value7_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.ItemController
			controller.Init(c)
			controller.UpdateValue7(value7_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/item/value8", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var value8_ int
			if v, flag := results["value8"]; flag {
				value8_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.ItemController
			controller.Init(c)
			controller.UpdateValue8(value8_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/item/value", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var value_ int
			if v, flag := results["value"]; flag {
				value_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.ItemController
			controller.Init(c)
			controller.UpdateValue(value_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/item/unit", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var unit_ string
			if v, flag := results["unit"]; flag {
				unit_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.ItemController
			controller.Init(c)
			controller.UpdateUnit(unit_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/item/status", func(c *fiber.Ctx) error {
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
			var controller rest.ItemController
			controller.Init(c)
			controller.UpdateStatus(status_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/item/reason", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var reason_ int
			if v, flag := results["reason"]; flag {
				reason_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.ItemController
			controller.Init(c)
			controller.UpdateReason(reason_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/item/reasontext", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var reasontext_ string
			if v, flag := results["reasontext"]; flag {
				reasontext_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.ItemController
			controller.Init(c)
			controller.UpdateReasontext(reasontext_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/item/action", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var action_ int
			if v, flag := results["action"]; flag {
				action_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.ItemController
			controller.Init(c)
			controller.UpdateAction(action_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/item/actiontext", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var actiontext_ string
			if v, flag := results["actiontext"]; flag {
				actiontext_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.ItemController
			controller.Init(c)
			controller.UpdateActiontext(actiontext_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/item/image", func(c *fiber.Ctx) error {
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
			var controller rest.ItemController
			controller.Init(c)
			controller.UpdateImage(image_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/item/order", func(c *fiber.Ctx) error {
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
			var controller rest.ItemController
			controller.Init(c)
			controller.UpdateOrder(order_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/item/data", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var data_ int64
			if v, flag := results["data"]; flag {
				data_ = int64(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.ItemController
			controller.Init(c)
			controller.UpdateData(data_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/item/report", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var report_ int64
			if v, flag := results["report"]; flag {
				report_ = int64(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.ItemController
			controller.Init(c)
			controller.UpdateReport(report_, id_)
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

		apiGroup.Get("/report/:id", func(c *fiber.Ctx) error {
			id_, _ := strconv.ParseInt(c.Params("id"), 10, 64)
			var controller rest.ReportController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/report", func(c *fiber.Ctx) error {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.ReportController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/report", func(c *fiber.Ctx) error {
			item_ := &models.Report{}
			c.BodyParser(item_)
			var controller rest.ReportController
			controller.Init(c)
			if item_ != nil {
				controller.Insert(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/report/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Report{}
			c.BodyParser(item_)
			var controller rest.ReportController
			controller.Init(c)
			if item_ != nil {
				controller.Insertbatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/report", func(c *fiber.Ctx) error {
			item_ := &models.Report{}
			c.BodyParser(item_)
			var controller rest.ReportController
			controller.Init(c)
			if item_ != nil {
				controller.Update(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/report", func(c *fiber.Ctx) error {
			item_ := &models.Report{}
			c.BodyParser(item_)
			var controller rest.ReportController
			controller.Init(c)
			if item_ != nil {
				controller.Delete(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/report/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Report{}
			c.BodyParser(item_)
			var controller rest.ReportController
			controller.Init(c)
			if item_ != nil {
				controller.Deletebatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/report/title", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var title_ string
			if v, flag := results["title"]; flag {
				title_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.ReportController
			controller.Init(c)
			controller.UpdateTitle(title_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/report/period", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var period_ int
			if v, flag := results["period"]; flag {
				period_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.ReportController
			controller.Init(c)
			controller.UpdatePeriod(period_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/report/number", func(c *fiber.Ctx) error {
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
			var controller rest.ReportController
			controller.Init(c)
			controller.UpdateNumber(number_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/report/checkdate", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var checkdate_ string
			if v, flag := results["checkdate"]; flag {
				checkdate_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.ReportController
			controller.Init(c)
			controller.UpdateCheckdate(checkdate_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/report/checktime", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var checktime_ string
			if v, flag := results["checktime"]; flag {
				checktime_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.ReportController
			controller.Init(c)
			controller.UpdateChecktime(checktime_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/report/status", func(c *fiber.Ctx) error {
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
			var controller rest.ReportController
			controller.Init(c)
			controller.UpdateStatus(status_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/report/company", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var company_ int64
			if v, flag := results["company"]; flag {
				company_ = int64(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.ReportController
			controller.Init(c)
			controller.UpdateCompany(company_, id_)
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