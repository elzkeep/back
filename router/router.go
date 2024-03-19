package router

import (
    "encoding/json"
    "strconv"
    "strings"
	"zkeep/controllers/api"
	"zkeep/controllers/rest"
    "zkeep/models"
    "zkeep/models/user"
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

		apiGroup.Get("/billing/print/:ids", func(c *fiber.Ctx) error {
			ids_ := getArrayCommal(c.Query("ids"))
			var controller api.BillingController
			controller.Init(c)
			controller.Print(ids_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/dashboard", func(c *fiber.Ctx) error {
			company_, _ := strconv.ParseInt(c.Query("company"), 10, 64)
			var controller api.DashboardController
			controller.Init(c)
			controller.Index(company_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/dataitem/process", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var datas_ []models.Dataitem
			datas__ref := &datas_
			c.BodyParser(datas__ref)
			var controller api.DataitemController
			controller.Init(c)
			controller.Process(datas_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/download/file/:id", func(c *fiber.Ctx) error {
			id_, _ := strconv.ParseInt(c.Params("id"), 10, 64)
			var controller api.DownloadController
			controller.Init(c)
			controller.File(id_)
			controller.Close()
            return nil
		})

		apiGroup.Get("/report/search/:page", func(c *fiber.Ctx) error {
			page_, _ := strconv.Atoi(c.Params("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller api.ReportController
			controller.Init(c)
			controller.Search(page_, pagesize_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/report", func(c *fiber.Ctx) error {
			item_ := &models.Report{}
			c.BodyParser(item_)
			var controller api.ReportController
			controller.Init(c)
			if item_ != nil {
				controller.Insert(item_)
			} else {
			    controller.Result["code"] = "error"
			}
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

		apiGroup.Get("/user/search/:page", func(c *fiber.Ctx) error {
			page_, _ := strconv.Atoi(c.Params("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller api.UserController
			controller.Init(c)
			controller.Search(page_, pagesize_)
			controller.Close()
			return c.JSON(controller.Result)
		})

	}

	{

		apiGroup.Get("/billing/:id", func(c *fiber.Ctx) error {
			id_, _ := strconv.ParseInt(c.Params("id"), 10, 64)
			var controller rest.BillingController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/billing", func(c *fiber.Ctx) error {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.BillingController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/billing/count", func(c *fiber.Ctx) error {
			var controller rest.BillingController
			controller.Init(c)
			controller.Count()
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/billing", func(c *fiber.Ctx) error {
			item_ := &models.Billing{}
			c.BodyParser(item_)
			var controller rest.BillingController
			controller.Init(c)
			if item_ != nil {
				controller.Insert(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/billing/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Billing{}
			c.BodyParser(item_)
			var controller rest.BillingController
			controller.Init(c)
			if item_ != nil {
				controller.Insertbatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/billing", func(c *fiber.Ctx) error {
			item_ := &models.Billing{}
			c.BodyParser(item_)
			var controller rest.BillingController
			controller.Init(c)
			if item_ != nil {
				controller.Update(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/billing", func(c *fiber.Ctx) error {
			item_ := &models.Billing{}
			c.BodyParser(item_)
			var controller rest.BillingController
			controller.Init(c)
			if item_ != nil {
				controller.Delete(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/billing/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Billing{}
			c.BodyParser(item_)
			var controller rest.BillingController
			controller.Init(c)
			if item_ != nil {
				controller.Deletebatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/billing/price", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var price_ int
			if v, flag := results["price"]; flag {
				price_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.BillingController
			controller.Init(c)
			controller.UpdatePrice(price_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/billing/status", func(c *fiber.Ctx) error {
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
			var controller rest.BillingController
			controller.Init(c)
			controller.UpdateStatus(status_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/billing/giro", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var giro_ int
			if v, flag := results["giro"]; flag {
				giro_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.BillingController
			controller.Init(c)
			controller.UpdateGiro(giro_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/billing/billdate", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var billdate_ string
			if v, flag := results["billdate"]; flag {
				billdate_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.BillingController
			controller.Init(c)
			controller.UpdateBilldate(billdate_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/billing/company", func(c *fiber.Ctx) error {
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
			var controller rest.BillingController
			controller.Init(c)
			controller.UpdateCompany(company_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/billing/building", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var building_ int64
			if v, flag := results["building"]; flag {
				building_ = int64(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.BillingController
			controller.Init(c)
			controller.UpdateBuilding(building_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/billing/sum", func(c *fiber.Ctx) error {
			var controller rest.BillingController
			controller.Init(c)
			controller.Sum()
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/billinglist/:id", func(c *fiber.Ctx) error {
			id_, _ := strconv.ParseInt(c.Params("id"), 10, 64)
			var controller rest.BillinglistController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/billinglist", func(c *fiber.Ctx) error {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.BillinglistController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/billinglist/count", func(c *fiber.Ctx) error {
			var controller rest.BillinglistController
			controller.Init(c)
			controller.Count()
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/billinglist/sum", func(c *fiber.Ctx) error {
			var controller rest.BillinglistController
			controller.Init(c)
			controller.Sum()
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/building/:id", func(c *fiber.Ctx) error {
			id_, _ := strconv.ParseInt(c.Params("id"), 10, 64)
			var controller rest.BuildingController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/building", func(c *fiber.Ctx) error {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.BuildingController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/building/count", func(c *fiber.Ctx) error {
			var controller rest.BuildingController
			controller.Init(c)
			controller.Count()
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/building", func(c *fiber.Ctx) error {
			item_ := &models.Building{}
			c.BodyParser(item_)
			var controller rest.BuildingController
			controller.Init(c)
			if item_ != nil {
				controller.Insert(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/building/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Building{}
			c.BodyParser(item_)
			var controller rest.BuildingController
			controller.Init(c)
			if item_ != nil {
				controller.Insertbatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/building", func(c *fiber.Ctx) error {
			item_ := &models.Building{}
			c.BodyParser(item_)
			var controller rest.BuildingController
			controller.Init(c)
			if item_ != nil {
				controller.Update(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/building", func(c *fiber.Ctx) error {
			item_ := &models.Building{}
			c.BodyParser(item_)
			var controller rest.BuildingController
			controller.Init(c)
			if item_ != nil {
				controller.Delete(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/building/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Building{}
			c.BodyParser(item_)
			var controller rest.BuildingController
			controller.Init(c)
			if item_ != nil {
				controller.Deletebatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/building/name", func(c *fiber.Ctx) error {
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
			var controller rest.BuildingController
			controller.Init(c)
			controller.UpdateName(name_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/building/conpanyno", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var conpanyno_ string
			if v, flag := results["conpanyno"]; flag {
				conpanyno_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.BuildingController
			controller.Init(c)
			controller.UpdateConpanyno(conpanyno_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/building/ceo", func(c *fiber.Ctx) error {
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
			var controller rest.BuildingController
			controller.Init(c)
			controller.UpdateCeo(ceo_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/building/address", func(c *fiber.Ctx) error {
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
			var controller rest.BuildingController
			controller.Init(c)
			controller.UpdateAddress(address_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/building/addressetc", func(c *fiber.Ctx) error {
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
			var controller rest.BuildingController
			controller.Init(c)
			controller.UpdateAddressetc(addressetc_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/building/score", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var score_ models.Double
			score__ref := &score_
			c.BodyParser(score__ref)
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.BuildingController
			controller.Init(c)
			controller.UpdateScore(score_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/building/status", func(c *fiber.Ctx) error {
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
			var controller rest.BuildingController
			controller.Init(c)
			controller.UpdateStatus(status_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/building/company", func(c *fiber.Ctx) error {
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
			var controller rest.BuildingController
			controller.Init(c)
			controller.UpdateCompany(company_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/building/sum", func(c *fiber.Ctx) error {
			var controller rest.BuildingController
			controller.Init(c)
			controller.Sum()
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/calendarcompanylist/:id", func(c *fiber.Ctx) error {
			id_, _ := strconv.ParseInt(c.Params("id"), 10, 64)
			var controller rest.CalendarcompanylistController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/calendarcompanylist", func(c *fiber.Ctx) error {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.CalendarcompanylistController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/calendarcompanylist/count", func(c *fiber.Ctx) error {
			var controller rest.CalendarcompanylistController
			controller.Init(c)
			controller.Count()
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/calendarcompanylist/sum", func(c *fiber.Ctx) error {
			var controller rest.CalendarcompanylistController
			controller.Init(c)
			controller.Sum()
			controller.Close()
			return c.JSON(controller.Result)
		})

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

		apiGroup.Get("/company/count", func(c *fiber.Ctx) error {
			var controller rest.CompanyController
			controller.Init(c)
			controller.Count()
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

		apiGroup.Get("/companylicense/:id", func(c *fiber.Ctx) error {
			id_, _ := strconv.ParseInt(c.Params("id"), 10, 64)
			var controller rest.CompanylicenseController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/companylicense", func(c *fiber.Ctx) error {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.CompanylicenseController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/companylicense/count", func(c *fiber.Ctx) error {
			var controller rest.CompanylicenseController
			controller.Init(c)
			controller.Count()
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/companylicense", func(c *fiber.Ctx) error {
			item_ := &models.Companylicense{}
			c.BodyParser(item_)
			var controller rest.CompanylicenseController
			controller.Init(c)
			if item_ != nil {
				controller.Insert(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/companylicense/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Companylicense{}
			c.BodyParser(item_)
			var controller rest.CompanylicenseController
			controller.Init(c)
			if item_ != nil {
				controller.Insertbatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/companylicense", func(c *fiber.Ctx) error {
			item_ := &models.Companylicense{}
			c.BodyParser(item_)
			var controller rest.CompanylicenseController
			controller.Init(c)
			if item_ != nil {
				controller.Update(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/companylicense", func(c *fiber.Ctx) error {
			item_ := &models.Companylicense{}
			c.BodyParser(item_)
			var controller rest.CompanylicenseController
			controller.Init(c)
			if item_ != nil {
				controller.Delete(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/companylicense/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Companylicense{}
			c.BodyParser(item_)
			var controller rest.CompanylicenseController
			controller.Init(c)
			if item_ != nil {
				controller.Deletebatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/companylicense/company", func(c *fiber.Ctx) error {
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
			var controller rest.CompanylicenseController
			controller.Init(c)
			controller.UpdateCompany(company_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/companylicense/licensecategory", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var licensecategory_ int64
			if v, flag := results["licensecategory"]; flag {
				licensecategory_ = int64(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CompanylicenseController
			controller.Init(c)
			controller.UpdateLicensecategory(licensecategory_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/companylicense/licenselevel", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var licenselevel_ int64
			if v, flag := results["licenselevel"]; flag {
				licenselevel_ = int64(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CompanylicenseController
			controller.Init(c)
			controller.UpdateLicenselevel(licenselevel_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/customer/:id", func(c *fiber.Ctx) error {
			id_, _ := strconv.ParseInt(c.Params("id"), 10, 64)
			var controller rest.CustomerController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/customer", func(c *fiber.Ctx) error {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.CustomerController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/customer/count", func(c *fiber.Ctx) error {
			var controller rest.CustomerController
			controller.Init(c)
			controller.Count()
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/customer", func(c *fiber.Ctx) error {
			item_ := &models.Customer{}
			c.BodyParser(item_)
			var controller rest.CustomerController
			controller.Init(c)
			if item_ != nil {
				controller.Insert(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/customer/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Customer{}
			c.BodyParser(item_)
			var controller rest.CustomerController
			controller.Init(c)
			if item_ != nil {
				controller.Insertbatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/customer", func(c *fiber.Ctx) error {
			item_ := &models.Customer{}
			c.BodyParser(item_)
			var controller rest.CustomerController
			controller.Init(c)
			if item_ != nil {
				controller.Update(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/customer", func(c *fiber.Ctx) error {
			item_ := &models.Customer{}
			c.BodyParser(item_)
			var controller rest.CustomerController
			controller.Init(c)
			if item_ != nil {
				controller.Delete(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/customer/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Customer{}
			c.BodyParser(item_)
			var controller rest.CustomerController
			controller.Init(c)
			if item_ != nil {
				controller.Deletebatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/customer/type", func(c *fiber.Ctx) error {
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
			var controller rest.CustomerController
			controller.Init(c)
			controller.UpdateType(type_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/customer/checkdate", func(c *fiber.Ctx) error {
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
			var controller rest.CustomerController
			controller.Init(c)
			controller.UpdateCheckdate(checkdate_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/customer/managername", func(c *fiber.Ctx) error {
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
			var controller rest.CustomerController
			controller.Init(c)
			controller.UpdateManagername(managername_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/customer/managertel", func(c *fiber.Ctx) error {
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
			var controller rest.CustomerController
			controller.Init(c)
			controller.UpdateManagertel(managertel_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/customer/manageremail", func(c *fiber.Ctx) error {
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
			var controller rest.CustomerController
			controller.Init(c)
			controller.UpdateManageremail(manageremail_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/customer/contractstartdate", func(c *fiber.Ctx) error {
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
			var controller rest.CustomerController
			controller.Init(c)
			controller.UpdateContractstartdate(contractstartdate_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/customer/contractenddate", func(c *fiber.Ctx) error {
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
			var controller rest.CustomerController
			controller.Init(c)
			controller.UpdateContractenddate(contractenddate_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/customer/contractprice", func(c *fiber.Ctx) error {
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
			var controller rest.CustomerController
			controller.Init(c)
			controller.UpdateContractprice(contractprice_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/customer/contractday", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var contractday_ int
			if v, flag := results["contractday"]; flag {
				contractday_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CustomerController
			controller.Init(c)
			controller.UpdateContractday(contractday_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/customer/billingdate", func(c *fiber.Ctx) error {
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
			var controller rest.CustomerController
			controller.Init(c)
			controller.UpdateBillingdate(billingdate_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/customer/billingname", func(c *fiber.Ctx) error {
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
			var controller rest.CustomerController
			controller.Init(c)
			controller.UpdateBillingname(billingname_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/customer/billingtel", func(c *fiber.Ctx) error {
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
			var controller rest.CustomerController
			controller.Init(c)
			controller.UpdateBillingtel(billingtel_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/customer/billingemail", func(c *fiber.Ctx) error {
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
			var controller rest.CustomerController
			controller.Init(c)
			controller.UpdateBillingemail(billingemail_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/customer/status", func(c *fiber.Ctx) error {
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
			var controller rest.CustomerController
			controller.Init(c)
			controller.UpdateStatus(status_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/customer/user", func(c *fiber.Ctx) error {
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
			var controller rest.CustomerController
			controller.Init(c)
			controller.UpdateUser(user_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/customer/company", func(c *fiber.Ctx) error {
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
			var controller rest.CustomerController
			controller.Init(c)
			controller.UpdateCompany(company_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/customer/building", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var building_ int64
			if v, flag := results["building"]; flag {
				building_ = int64(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CustomerController
			controller.Init(c)
			controller.UpdateBuilding(building_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/customercompany/:id", func(c *fiber.Ctx) error {
			id_, _ := strconv.ParseInt(c.Params("id"), 10, 64)
			var controller rest.CustomercompanyController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/customercompany", func(c *fiber.Ctx) error {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.CustomercompanyController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/customercompany/count", func(c *fiber.Ctx) error {
			var controller rest.CustomercompanyController
			controller.Init(c)
			controller.Count()
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/customercompany/sum", func(c *fiber.Ctx) error {
			var controller rest.CustomercompanyController
			controller.Init(c)
			controller.Sum()
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

		apiGroup.Get("/data/count", func(c *fiber.Ctx) error {
			var controller rest.DataController
			controller.Init(c)
			controller.Count()
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

		apiGroup.Delete("/data/byreporttopcategory", func(c *fiber.Ctx) error {
			item_ := &models.Data{}
			c.BodyParser(item_)
			var controller rest.DataController
			controller.Init(c)
			controller.DeleteByReportTopcategory(item_.Report, item_.Topcategory)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/data/topcategory", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var topcategory_ int
			if v, flag := results["topcategory"]; flag {
				topcategory_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.DataController
			controller.Init(c)
			controller.UpdateTopcategory(topcategory_, id_)
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

		apiGroup.Get("/department/:id", func(c *fiber.Ctx) error {
			id_, _ := strconv.ParseInt(c.Params("id"), 10, 64)
			var controller rest.DepartmentController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/department", func(c *fiber.Ctx) error {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.DepartmentController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/department/count", func(c *fiber.Ctx) error {
			var controller rest.DepartmentController
			controller.Init(c)
			controller.Count()
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/department", func(c *fiber.Ctx) error {
			item_ := &models.Department{}
			c.BodyParser(item_)
			var controller rest.DepartmentController
			controller.Init(c)
			if item_ != nil {
				controller.Insert(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/department/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Department{}
			c.BodyParser(item_)
			var controller rest.DepartmentController
			controller.Init(c)
			if item_ != nil {
				controller.Insertbatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/department", func(c *fiber.Ctx) error {
			item_ := &models.Department{}
			c.BodyParser(item_)
			var controller rest.DepartmentController
			controller.Init(c)
			if item_ != nil {
				controller.Update(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/department", func(c *fiber.Ctx) error {
			item_ := &models.Department{}
			c.BodyParser(item_)
			var controller rest.DepartmentController
			controller.Init(c)
			if item_ != nil {
				controller.Delete(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/department/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Department{}
			c.BodyParser(item_)
			var controller rest.DepartmentController
			controller.Init(c)
			if item_ != nil {
				controller.Deletebatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/department/name", func(c *fiber.Ctx) error {
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
			var controller rest.DepartmentController
			controller.Init(c)
			controller.UpdateName(name_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/department/order", func(c *fiber.Ctx) error {
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
			var controller rest.DepartmentController
			controller.Init(c)
			controller.UpdateOrder(order_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/department/company", func(c *fiber.Ctx) error {
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
			var controller rest.DepartmentController
			controller.Init(c)
			controller.UpdateCompany(company_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/facility/:id", func(c *fiber.Ctx) error {
			id_, _ := strconv.ParseInt(c.Params("id"), 10, 64)
			var controller rest.FacilityController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/facility", func(c *fiber.Ctx) error {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.FacilityController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/facility/count", func(c *fiber.Ctx) error {
			var controller rest.FacilityController
			controller.Init(c)
			controller.Count()
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/facility", func(c *fiber.Ctx) error {
			item_ := &models.Facility{}
			c.BodyParser(item_)
			var controller rest.FacilityController
			controller.Init(c)
			if item_ != nil {
				controller.Insert(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/facility/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Facility{}
			c.BodyParser(item_)
			var controller rest.FacilityController
			controller.Init(c)
			if item_ != nil {
				controller.Insertbatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/facility", func(c *fiber.Ctx) error {
			item_ := &models.Facility{}
			c.BodyParser(item_)
			var controller rest.FacilityController
			controller.Init(c)
			if item_ != nil {
				controller.Update(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/facility", func(c *fiber.Ctx) error {
			item_ := &models.Facility{}
			c.BodyParser(item_)
			var controller rest.FacilityController
			controller.Init(c)
			if item_ != nil {
				controller.Delete(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/facility/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Facility{}
			c.BodyParser(item_)
			var controller rest.FacilityController
			controller.Init(c)
			if item_ != nil {
				controller.Deletebatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/facility/bybuildingcategory", func(c *fiber.Ctx) error {
			item_ := &models.Facility{}
			c.BodyParser(item_)
			var controller rest.FacilityController
			controller.Init(c)
			controller.DeleteByBuildingCategory(item_.Building, item_.Category)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/facility/category", func(c *fiber.Ctx) error {
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
			var controller rest.FacilityController
			controller.Init(c)
			controller.UpdateCategory(category_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/facility/parent", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var parent_ int64
			if v, flag := results["parent"]; flag {
				parent_ = int64(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.FacilityController
			controller.Init(c)
			controller.UpdateParent(parent_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/facility/name", func(c *fiber.Ctx) error {
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
			var controller rest.FacilityController
			controller.Init(c)
			controller.UpdateName(name_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/facility/value1", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var value1_ string
			if v, flag := results["value1"]; flag {
				value1_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.FacilityController
			controller.Init(c)
			controller.UpdateValue1(value1_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/facility/value2", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var value2_ string
			if v, flag := results["value2"]; flag {
				value2_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.FacilityController
			controller.Init(c)
			controller.UpdateValue2(value2_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/facility/value3", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var value3_ string
			if v, flag := results["value3"]; flag {
				value3_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.FacilityController
			controller.Init(c)
			controller.UpdateValue3(value3_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/facility/value4", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var value4_ string
			if v, flag := results["value4"]; flag {
				value4_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.FacilityController
			controller.Init(c)
			controller.UpdateValue4(value4_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/facility/value5", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var value5_ string
			if v, flag := results["value5"]; flag {
				value5_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.FacilityController
			controller.Init(c)
			controller.UpdateValue5(value5_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/facility/value6", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var value6_ string
			if v, flag := results["value6"]; flag {
				value6_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.FacilityController
			controller.Init(c)
			controller.UpdateValue6(value6_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/facility/value7", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var value7_ string
			if v, flag := results["value7"]; flag {
				value7_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.FacilityController
			controller.Init(c)
			controller.UpdateValue7(value7_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/facility/value8", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var value8_ string
			if v, flag := results["value8"]; flag {
				value8_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.FacilityController
			controller.Init(c)
			controller.UpdateValue8(value8_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/facility/value9", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var value9_ string
			if v, flag := results["value9"]; flag {
				value9_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.FacilityController
			controller.Init(c)
			controller.UpdateValue9(value9_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/facility/value10", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var value10_ string
			if v, flag := results["value10"]; flag {
				value10_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.FacilityController
			controller.Init(c)
			controller.UpdateValue10(value10_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/facility/value11", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var value11_ string
			if v, flag := results["value11"]; flag {
				value11_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.FacilityController
			controller.Init(c)
			controller.UpdateValue11(value11_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/facility/value12", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var value12_ string
			if v, flag := results["value12"]; flag {
				value12_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.FacilityController
			controller.Init(c)
			controller.UpdateValue12(value12_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/facility/value13", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var value13_ string
			if v, flag := results["value13"]; flag {
				value13_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.FacilityController
			controller.Init(c)
			controller.UpdateValue13(value13_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/facility/value14", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var value14_ string
			if v, flag := results["value14"]; flag {
				value14_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.FacilityController
			controller.Init(c)
			controller.UpdateValue14(value14_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/facility/value15", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var value15_ string
			if v, flag := results["value15"]; flag {
				value15_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.FacilityController
			controller.Init(c)
			controller.UpdateValue15(value15_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/facility/value16", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var value16_ string
			if v, flag := results["value16"]; flag {
				value16_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.FacilityController
			controller.Init(c)
			controller.UpdateValue16(value16_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/facility/value17", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var value17_ string
			if v, flag := results["value17"]; flag {
				value17_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.FacilityController
			controller.Init(c)
			controller.UpdateValue17(value17_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/facility/value18", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var value18_ string
			if v, flag := results["value18"]; flag {
				value18_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.FacilityController
			controller.Init(c)
			controller.UpdateValue18(value18_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/facility/value19", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var value19_ string
			if v, flag := results["value19"]; flag {
				value19_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.FacilityController
			controller.Init(c)
			controller.UpdateValue19(value19_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/facility/value20", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var value20_ string
			if v, flag := results["value20"]; flag {
				value20_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.FacilityController
			controller.Init(c)
			controller.UpdateValue20(value20_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/facility/value21", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var value21_ string
			if v, flag := results["value21"]; flag {
				value21_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.FacilityController
			controller.Init(c)
			controller.UpdateValue21(value21_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/facility/value22", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var value22_ string
			if v, flag := results["value22"]; flag {
				value22_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.FacilityController
			controller.Init(c)
			controller.UpdateValue22(value22_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/facility/value23", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var value23_ string
			if v, flag := results["value23"]; flag {
				value23_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.FacilityController
			controller.Init(c)
			controller.UpdateValue23(value23_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/facility/value24", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var value24_ string
			if v, flag := results["value24"]; flag {
				value24_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.FacilityController
			controller.Init(c)
			controller.UpdateValue24(value24_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/facility/value25", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var value25_ string
			if v, flag := results["value25"]; flag {
				value25_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.FacilityController
			controller.Init(c)
			controller.UpdateValue25(value25_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/facility/content", func(c *fiber.Ctx) error {
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
			var controller rest.FacilityController
			controller.Init(c)
			controller.UpdateContent(content_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/facility/building", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var building_ int64
			if v, flag := results["building"]; flag {
				building_ = int64(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.FacilityController
			controller.Init(c)
			controller.UpdateBuilding(building_, id_)
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

		apiGroup.Get("/item/count", func(c *fiber.Ctx) error {
			var controller rest.ItemController
			controller.Init(c)
			controller.Count()
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

		apiGroup.Delete("/item/byreporttopcategory", func(c *fiber.Ctx) error {
			item_ := &models.Item{}
			c.BodyParser(item_)
			var controller rest.ItemController
			controller.Init(c)
			controller.DeleteByReportTopcategory(item_.Report, item_.Topcategory)
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

		apiGroup.Put("/item/content", func(c *fiber.Ctx) error {
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
			var controller rest.ItemController
			controller.Init(c)
			controller.UpdateContent(content_, id_)
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

		apiGroup.Put("/item/topcategory", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var topcategory_ int
			if v, flag := results["topcategory"]; flag {
				topcategory_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.ItemController
			controller.Init(c)
			controller.UpdateTopcategory(topcategory_, id_)
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

		apiGroup.Get("/license/:id", func(c *fiber.Ctx) error {
			id_, _ := strconv.ParseInt(c.Params("id"), 10, 64)
			var controller rest.LicenseController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/license", func(c *fiber.Ctx) error {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.LicenseController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/license/count", func(c *fiber.Ctx) error {
			var controller rest.LicenseController
			controller.Init(c)
			controller.Count()
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/license", func(c *fiber.Ctx) error {
			item_ := &models.License{}
			c.BodyParser(item_)
			var controller rest.LicenseController
			controller.Init(c)
			if item_ != nil {
				controller.Insert(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/license/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.License{}
			c.BodyParser(item_)
			var controller rest.LicenseController
			controller.Init(c)
			if item_ != nil {
				controller.Insertbatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/license", func(c *fiber.Ctx) error {
			item_ := &models.License{}
			c.BodyParser(item_)
			var controller rest.LicenseController
			controller.Init(c)
			if item_ != nil {
				controller.Update(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/license", func(c *fiber.Ctx) error {
			item_ := &models.License{}
			c.BodyParser(item_)
			var controller rest.LicenseController
			controller.Init(c)
			if item_ != nil {
				controller.Delete(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/license/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.License{}
			c.BodyParser(item_)
			var controller rest.LicenseController
			controller.Init(c)
			if item_ != nil {
				controller.Deletebatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/license/user", func(c *fiber.Ctx) error {
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
			var controller rest.LicenseController
			controller.Init(c)
			controller.UpdateUser(user_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/license/licensecategory", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var licensecategory_ int64
			if v, flag := results["licensecategory"]; flag {
				licensecategory_ = int64(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.LicenseController
			controller.Init(c)
			controller.UpdateLicensecategory(licensecategory_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/license/licenselevel", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var licenselevel_ int64
			if v, flag := results["licenselevel"]; flag {
				licenselevel_ = int64(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.LicenseController
			controller.Init(c)
			controller.UpdateLicenselevel(licenselevel_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/licensecategory/:id", func(c *fiber.Ctx) error {
			id_, _ := strconv.ParseInt(c.Params("id"), 10, 64)
			var controller rest.LicensecategoryController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/licensecategory", func(c *fiber.Ctx) error {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.LicensecategoryController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/licensecategory/count", func(c *fiber.Ctx) error {
			var controller rest.LicensecategoryController
			controller.Init(c)
			controller.Count()
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/licensecategory", func(c *fiber.Ctx) error {
			item_ := &models.Licensecategory{}
			c.BodyParser(item_)
			var controller rest.LicensecategoryController
			controller.Init(c)
			if item_ != nil {
				controller.Insert(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/licensecategory/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Licensecategory{}
			c.BodyParser(item_)
			var controller rest.LicensecategoryController
			controller.Init(c)
			if item_ != nil {
				controller.Insertbatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/licensecategory", func(c *fiber.Ctx) error {
			item_ := &models.Licensecategory{}
			c.BodyParser(item_)
			var controller rest.LicensecategoryController
			controller.Init(c)
			if item_ != nil {
				controller.Update(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/licensecategory", func(c *fiber.Ctx) error {
			item_ := &models.Licensecategory{}
			c.BodyParser(item_)
			var controller rest.LicensecategoryController
			controller.Init(c)
			if item_ != nil {
				controller.Delete(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/licensecategory/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Licensecategory{}
			c.BodyParser(item_)
			var controller rest.LicensecategoryController
			controller.Init(c)
			if item_ != nil {
				controller.Deletebatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/licensecategory/name", func(c *fiber.Ctx) error {
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
			var controller rest.LicensecategoryController
			controller.Init(c)
			controller.UpdateName(name_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/licensecategory/order", func(c *fiber.Ctx) error {
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
			var controller rest.LicensecategoryController
			controller.Init(c)
			controller.UpdateOrder(order_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/licenselevel/:id", func(c *fiber.Ctx) error {
			id_, _ := strconv.ParseInt(c.Params("id"), 10, 64)
			var controller rest.LicenselevelController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/licenselevel", func(c *fiber.Ctx) error {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.LicenselevelController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/licenselevel/count", func(c *fiber.Ctx) error {
			var controller rest.LicenselevelController
			controller.Init(c)
			controller.Count()
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/licenselevel", func(c *fiber.Ctx) error {
			item_ := &models.Licenselevel{}
			c.BodyParser(item_)
			var controller rest.LicenselevelController
			controller.Init(c)
			if item_ != nil {
				controller.Insert(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/licenselevel/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Licenselevel{}
			c.BodyParser(item_)
			var controller rest.LicenselevelController
			controller.Init(c)
			if item_ != nil {
				controller.Insertbatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/licenselevel", func(c *fiber.Ctx) error {
			item_ := &models.Licenselevel{}
			c.BodyParser(item_)
			var controller rest.LicenselevelController
			controller.Init(c)
			if item_ != nil {
				controller.Update(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/licenselevel", func(c *fiber.Ctx) error {
			item_ := &models.Licenselevel{}
			c.BodyParser(item_)
			var controller rest.LicenselevelController
			controller.Init(c)
			if item_ != nil {
				controller.Delete(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/licenselevel/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Licenselevel{}
			c.BodyParser(item_)
			var controller rest.LicenselevelController
			controller.Init(c)
			if item_ != nil {
				controller.Deletebatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/licenselevel/name", func(c *fiber.Ctx) error {
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
			var controller rest.LicenselevelController
			controller.Init(c)
			controller.UpdateName(name_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/licenselevel/order", func(c *fiber.Ctx) error {
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
			var controller rest.LicenselevelController
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

		apiGroup.Get("/notice/count", func(c *fiber.Ctx) error {
			var controller rest.NoticeController
			controller.Init(c)
			controller.Count()
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

		apiGroup.Put("/notice/title", func(c *fiber.Ctx) error {
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
			var controller rest.NoticeController
			controller.Init(c)
			controller.UpdateTitle(title_, id_)
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

		apiGroup.Get("/report/count", func(c *fiber.Ctx) error {
			var controller rest.ReportController
			controller.Init(c)
			controller.Count()
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

		apiGroup.Put("/report/content", func(c *fiber.Ctx) error {
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
			var controller rest.ReportController
			controller.Init(c)
			controller.UpdateContent(content_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/report/image", func(c *fiber.Ctx) error {
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
			var controller rest.ReportController
			controller.Init(c)
			controller.UpdateImage(image_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/report/sign1", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var sign1_ string
			if v, flag := results["sign1"]; flag {
				sign1_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.ReportController
			controller.Init(c)
			controller.UpdateSign1(sign1_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/report/sign2", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var sign2_ string
			if v, flag := results["sign2"]; flag {
				sign2_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.ReportController
			controller.Init(c)
			controller.UpdateSign2(sign2_, id_)
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

		apiGroup.Put("/report/user", func(c *fiber.Ctx) error {
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
			var controller rest.ReportController
			controller.Init(c)
			controller.UpdateUser(user_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/report/building", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var building_ int64
			if v, flag := results["building"]; flag {
				building_ = int64(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.ReportController
			controller.Init(c)
			controller.UpdateBuilding(building_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/statisticsday/:id", func(c *fiber.Ctx) error {
			id_, _ := strconv.ParseInt(c.Params("id"), 10, 64)
			var controller rest.StatisticsdayController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/statisticsday", func(c *fiber.Ctx) error {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.StatisticsdayController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/statisticsday/count", func(c *fiber.Ctx) error {
			var controller rest.StatisticsdayController
			controller.Init(c)
			controller.Count()
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/statisticsday/sum", func(c *fiber.Ctx) error {
			var controller rest.StatisticsdayController
			controller.Init(c)
			controller.Sum()
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/statisticsmonth/:id", func(c *fiber.Ctx) error {
			id_, _ := strconv.ParseInt(c.Params("id"), 10, 64)
			var controller rest.StatisticsmonthController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/statisticsmonth", func(c *fiber.Ctx) error {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.StatisticsmonthController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/statisticsmonth/count", func(c *fiber.Ctx) error {
			var controller rest.StatisticsmonthController
			controller.Init(c)
			controller.Count()
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/statisticsmonth/sum", func(c *fiber.Ctx) error {
			var controller rest.StatisticsmonthController
			controller.Init(c)
			controller.Sum()
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/statisticsyear/:id", func(c *fiber.Ctx) error {
			id_, _ := strconv.ParseInt(c.Params("id"), 10, 64)
			var controller rest.StatisticsyearController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/statisticsyear", func(c *fiber.Ctx) error {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.StatisticsyearController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/statisticsyear/count", func(c *fiber.Ctx) error {
			var controller rest.StatisticsyearController
			controller.Init(c)
			controller.Count()
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/statisticsyear/sum", func(c *fiber.Ctx) error {
			var controller rest.StatisticsyearController
			controller.Init(c)
			controller.Sum()
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

		apiGroup.Get("/user/count", func(c *fiber.Ctx) error {
			var controller rest.UserController
			controller.Init(c)
			controller.Count()
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

		apiGroup.Get("/user/get/loginid/:loginid", func(c *fiber.Ctx) error {
			loginid_ := c.Params("loginid")
			var controller rest.UserController
			controller.Init(c)
			controller.GetByLoginid(loginid_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/user/count/loginid/:loginid", func(c *fiber.Ctx) error {
			loginid_ := c.Params("loginid")
			var controller rest.UserController
			controller.Init(c)
			controller.CountByLoginid(loginid_)
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

		apiGroup.Get("/user/count/company/:company", func(c *fiber.Ctx) error {
			company_, _ := strconv.ParseInt(c.Params("company"), 10, 64)
			var controller rest.UserController
			controller.Init(c)
			controller.CountByCompany(company_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/user/loginid", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var loginid_ string
			if v, flag := results["loginid"]; flag {
				loginid_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.UserController
			controller.Init(c)
			controller.UpdateLoginid(loginid_, id_)
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

		apiGroup.Put("/user/tel", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var tel_ string
			if v, flag := results["tel"]; flag {
				tel_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.UserController
			controller.Init(c)
			controller.UpdateTel(tel_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/user/address", func(c *fiber.Ctx) error {
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
			var controller rest.UserController
			controller.Init(c)
			controller.UpdateAddress(address_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/user/addressetc", func(c *fiber.Ctx) error {
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
			var controller rest.UserController
			controller.Init(c)
			controller.UpdateAddressetc(addressetc_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/user/joindate", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var joindate_ string
			if v, flag := results["joindate"]; flag {
				joindate_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.UserController
			controller.Init(c)
			controller.UpdateJoindate(joindate_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/user/careeryear", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var careeryear_ int
			if v, flag := results["careeryear"]; flag {
				careeryear_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.UserController
			controller.Init(c)
			controller.UpdateCareeryear(careeryear_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/user/careermonth", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var careermonth_ int
			if v, flag := results["careermonth"]; flag {
				careermonth_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.UserController
			controller.Init(c)
			controller.UpdateCareermonth(careermonth_, id_)
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

		apiGroup.Put("/user/score", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var score_ models.Double
			score__ref := &score_
			c.BodyParser(score__ref)
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.UserController
			controller.Init(c)
			controller.UpdateScore(score_, id_)
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

		apiGroup.Put("/user/company", func(c *fiber.Ctx) error {
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
			var controller rest.UserController
			controller.Init(c)
			controller.UpdateCompany(company_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/user/department", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var department_ int64
			if v, flag := results["department"]; flag {
				department_ = int64(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.UserController
			controller.Init(c)
			controller.UpdateDepartment(department_, id_)
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

		apiGroup.Get("/userlist/:id", func(c *fiber.Ctx) error {
			id_, _ := strconv.ParseInt(c.Params("id"), 10, 64)
			var controller rest.UserlistController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/userlist", func(c *fiber.Ctx) error {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.UserlistController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/userlist/count", func(c *fiber.Ctx) error {
			var controller rest.UserlistController
			controller.Init(c)
			controller.Count()
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/userlist/sum", func(c *fiber.Ctx) error {
			var controller rest.UserlistController
			controller.Init(c)
			controller.Sum()
			controller.Close()
			return c.JSON(controller.Result)
		})

	}

}