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

		apiGroup.Get("/billing/search/:page", func(c *fiber.Ctx) error {
			page_, _ := strconv.Atoi(c.Params("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller api.BillingController
			controller.Init(c)
			controller.Search(page_, pagesize_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/billing/make", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var durationtype_ int
			if v, flag := results["durationtype"]; flag {
				durationtype_ = int(v.(float64))
			}
			var base_ int
			if v, flag := results["base"]; flag {
				base_ = int(v.(float64))
			}
			var year_ int
			if v, flag := results["year"]; flag {
				year_ = int(v.(float64))
			}
			var month_ int
			if v, flag := results["month"]; flag {
				month_ = int(v.(float64))
			}
			var durationmonth_ []int
			if v, flag := results["durationmonth"]; flag {
				durationmonth_= getArrayCommai(v.(string))
			}
			var ids_ []int64
			if v, flag := results["ids"]; flag {
				ids_= getArrayCommal(v.(string))
			}
			var price_ []int
			if v, flag := results["price"]; flag {
				price_= getArrayCommai(v.(string))
			}
			var vat_ []int
			if v, flag := results["vat"]; flag {
				vat_= getArrayCommai(v.(string))
			}
			var controller api.BillingController
			controller.Init(c)
			controller.Make(durationtype_, base_, year_, month_, durationmonth_, ids_, price_, vat_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/billing/process", func(c *fiber.Ctx) error {
			item_ := &models.Billing{}
			c.BodyParser(item_)
			var controller api.BillingController
			controller.Init(c)
			controller.Process(item_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/billinglist/initdata", func(c *fiber.Ctx) error {
			var controller api.BillinglistController
			controller.Init(c)
			controller.InitData()
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/billinguserlist/excel/:company", func(c *fiber.Ctx) error {
			company_, _ := strconv.ParseInt(c.Params("company"), 10, 64)
			startdate_ := c.Query("startdate")
			enddate_ := c.Query("enddate")
			users_ := getArrayCommal(c.Query("users"))
			var controller api.BillinguserlistController
			controller.Init(c)
			controller.Excel(company_, startdate_, enddate_, users_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/building/score", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller api.BuildingController
			controller.Init(c)
			controller.Score(id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/company/search/:page", func(c *fiber.Ctx) error {
			page_, _ := strconv.Atoi(c.Params("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller api.CompanyController
			controller.Init(c)
			controller.Search(page_, pagesize_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/company/upload/:filename", func(c *fiber.Ctx) error {
			filename_ := c.Params("filename")
			var controller api.CompanyController
			controller.Init(c)
			controller.Upload(filename_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/company/totalscore", func(c *fiber.Ctx) error {
			var controller api.CompanyController
			controller.Init(c)
			controller.Totalscore()
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/companylist/search/:page", func(c *fiber.Ctx) error {
			page_, _ := strconv.Atoi(c.Params("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller api.CompanylistController
			controller.Init(c)
			controller.Search(page_, pagesize_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/customer", func(c *fiber.Ctx) error {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller api.CustomerController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/customer/status/:id", func(c *fiber.Ctx) error {
			id_, _ := strconv.ParseInt(c.Params("id"), 10, 64)
			var controller api.CustomerController
			controller.Init(c)
			controller.Status(id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/customer/upload/:filename", func(c *fiber.Ctx) error {
			filename_ := c.Params("filename")
			var controller api.CustomerController
			controller.Init(c)
			controller.Upload(filename_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/customer/maxnumber/:id", func(c *fiber.Ctx) error {
			id_, _ := strconv.ParseInt(c.Params("id"), 10, 64)
			var controller api.CustomerController
			controller.Init(c)
			controller.MaxNumber(id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/customer/initdata", func(c *fiber.Ctx) error {
			var controller api.CustomerController
			controller.Init(c)
			controller.InitData()
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/dashboard/initdata", func(c *fiber.Ctx) error {
			var controller api.DashboardController
			controller.Init(c)
			controller.InitData()
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

		apiGroup.Get("/download/giro/:ids", func(c *fiber.Ctx) error {
			ids_ := getArrayCommal(c.Params("ids"))
			var controller api.DownloadController
			controller.Init(c)
			controller.Giro(ids_)
			controller.Close()
            return nil
		})

		apiGroup.Get("/download/girowork/:ids", func(c *fiber.Ctx) error {
			ids_ := getArrayCommal(c.Params("ids"))
			var controller api.DownloadController
			controller.Init(c)
			controller.GiroWork(ids_)
			controller.Close()
            return nil
		})

		apiGroup.Get("/download/company", func(c *fiber.Ctx) error {
			var controller api.DownloadController
			controller.Init(c)
			controller.Company()
			controller.Close()
            return nil
		})

		apiGroup.Get("/download/user", func(c *fiber.Ctx) error {
			var controller api.DownloadController
			controller.Init(c)
			controller.User()
			controller.Close()
            return nil
		})

		apiGroup.Get("/download/companyexample", func(c *fiber.Ctx) error {
			var controller api.DownloadController
			controller.Init(c)
			controller.CompanyExample()
			controller.Close()
            return nil
		})

		apiGroup.Get("/download/customerexample", func(c *fiber.Ctx) error {
			var controller api.DownloadController
			controller.Init(c)
			controller.CustomerExample()
			controller.Close()
            return nil
		})

		apiGroup.Get("/download/userexample", func(c *fiber.Ctx) error {
			var controller api.DownloadController
			controller.Init(c)
			controller.UserExample()
			controller.Close()
            return nil
		})

		apiGroup.Get("/download/all/:category", func(c *fiber.Ctx) error {
			category_, _ := strconv.Atoi(c.Params("category"))
			var controller api.DownloadController
			controller.Init(c)
			controller.All(category_)
			controller.Close()
            return nil
		})

		apiGroup.Post("/external/giro", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var filename_ []string
			if v, flag := results["filename"]; flag {
			    strs := make([]string, 0)
			    for _, str := range v.([]interface{}) {
			        strs = append(strs, str.(string))
			    }
				filename_ = strs
			}
			var controller api.ExternalController
			controller.Init(c)
			controller.Giro(filename_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/external", func(c *fiber.Ctx) error {
			filenames_ := c.Query("filenames")
			type_, _ := strconv.Atoi(c.Query("type"))
			var controller api.ExternalController
			controller.Init(c)
			controller.Index(filenames_, type_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/external/user/:filename", func(c *fiber.Ctx) error {
			filename_ := c.Params("filename")
			var controller api.ExternalController
			controller.Init(c)
			controller.User(filename_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/external/all/:category", func(c *fiber.Ctx) error {
			category_, _ := strconv.Atoi(c.Params("category"))
			filename_ := c.Query("filename")
			var controller api.ExternalController
			controller.Init(c)
			controller.All(category_, filename_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/mail/index", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var to_ string
			if v, flag := results["to"]; flag {
				to_ = v.(string)
			}
			var subject_ string
			if v, flag := results["subject"]; flag {
				subject_ = v.(string)
			}
			var body_ string
			if v, flag := results["body"]; flag {
				body_ = v.(string)
			}
			var controller api.MailController
			controller.Init(c)
			controller.Index(to_, subject_, body_)
			controller.Close()
			return c.JSON(controller.Result)
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

		apiGroup.Get("/report/download/:id", func(c *fiber.Ctx) error {
			id_, _ := strconv.ParseInt(c.Params("id"), 10, 64)
			var controller api.ReportController
			controller.Init(c)
			controller.Download(id_)
			controller.Close()
            return nil
		})

		apiGroup.Post("/sms/index", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var to_ string
			if v, flag := results["to"]; flag {
				to_ = v.(string)
			}
			var message_ string
			if v, flag := results["message"]; flag {
				message_ = v.(string)
			}
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

		apiGroup.Get("/user/search", func(c *fiber.Ctx) error {
			var controller api.UserController
			controller.Init(c)
			controller.Search()
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/user/upload/:filename", func(c *fiber.Ctx) error {
			filename_ := c.Params("filename")
			var controller api.UserController
			controller.Init(c)
			controller.Upload(filename_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/user", func(c *fiber.Ctx) error {
			item_ := &models.User{}
			c.BodyParser(item_)
			var controller api.UserController
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
			var controller api.UserController
			controller.Init(c)
			if item_ != nil {
				controller.Deletebatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/userlist/initdata", func(c *fiber.Ctx) error {
			var controller api.UserlistController
			controller.Init(c)
			controller.InitData()
			controller.Close()
			return c.JSON(controller.Result)
		})

	}

	{

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

		apiGroup.Get("/billing/count", func(c *fiber.Ctx) error {
			var controller rest.BillingController
			controller.Init(c)
			controller.Count()
			controller.Close()
			return c.JSON(controller.Result)
		})

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

		apiGroup.Put("/billing/title", func(c *fiber.Ctx) error {
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
			var controller rest.BillingController
			controller.Init(c)
			controller.UpdateTitle(title_, id_)
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

		apiGroup.Put("/billing/vat", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var vat_ int
			if v, flag := results["vat"]; flag {
				vat_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.BillingController
			controller.Init(c)
			controller.UpdateVat(vat_, id_)
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

		apiGroup.Put("/billing/month", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var month_ string
			if v, flag := results["month"]; flag {
				month_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.BillingController
			controller.Init(c)
			controller.UpdateMonth(month_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/billing/endmonth", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var endmonth_ string
			if v, flag := results["endmonth"]; flag {
				endmonth_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.BillingController
			controller.Init(c)
			controller.UpdateEndmonth(endmonth_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/billing/period", func(c *fiber.Ctx) error {
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
			var controller rest.BillingController
			controller.Init(c)
			controller.UpdatePeriod(period_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/billing/remark", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var remark_ string
			if v, flag := results["remark"]; flag {
				remark_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.BillingController
			controller.Init(c)
			controller.UpdateRemark(remark_, id_)
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

		apiGroup.Delete("/billinghistory/bybilling", func(c *fiber.Ctx) error {
			item_ := &models.Billinghistory{}
			c.BodyParser(item_)
			var controller rest.BillinghistoryController
			controller.Init(c)
			controller.DeleteByBilling(item_.Billing)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/billinghistory", func(c *fiber.Ctx) error {
			item_ := &models.Billinghistory{}
			c.BodyParser(item_)
			var controller rest.BillinghistoryController
			controller.Init(c)
			if item_ != nil {
				controller.Insert(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/billinghistory/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Billinghistory{}
			c.BodyParser(item_)
			var controller rest.BillinghistoryController
			controller.Init(c)
			if item_ != nil {
				controller.Insertbatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/billinghistory", func(c *fiber.Ctx) error {
			item_ := &models.Billinghistory{}
			c.BodyParser(item_)
			var controller rest.BillinghistoryController
			controller.Init(c)
			if item_ != nil {
				controller.Update(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/billinghistory", func(c *fiber.Ctx) error {
			item_ := &models.Billinghistory{}
			c.BodyParser(item_)
			var controller rest.BillinghistoryController
			controller.Init(c)
			if item_ != nil {
				controller.Delete(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/billinghistory/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Billinghistory{}
			c.BodyParser(item_)
			var controller rest.BillinghistoryController
			controller.Init(c)
			if item_ != nil {
				controller.Deletebatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/billinghistory/count", func(c *fiber.Ctx) error {
			var controller rest.BillinghistoryController
			controller.Init(c)
			controller.Count()
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/billinghistory/:id", func(c *fiber.Ctx) error {
			id_, _ := strconv.ParseInt(c.Params("id"), 10, 64)
			var controller rest.BillinghistoryController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/billinghistory", func(c *fiber.Ctx) error {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.BillinghistoryController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/billinghistory/price", func(c *fiber.Ctx) error {
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
			var controller rest.BillinghistoryController
			controller.Init(c)
			controller.UpdatePrice(price_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/billinghistory/company", func(c *fiber.Ctx) error {
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
			var controller rest.BillinghistoryController
			controller.Init(c)
			controller.UpdateCompany(company_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/billinghistory/building", func(c *fiber.Ctx) error {
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
			var controller rest.BillinghistoryController
			controller.Init(c)
			controller.UpdateBuilding(building_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/billinghistory/billing", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var billing_ int64
			if v, flag := results["billing"]; flag {
				billing_ = int64(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.BillinghistoryController
			controller.Init(c)
			controller.UpdateBilling(billing_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/billinghistory/sum", func(c *fiber.Ctx) error {
			var controller rest.BillinghistoryController
			controller.Init(c)
			controller.Sum()
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

		apiGroup.Get("/billinglist/sum", func(c *fiber.Ctx) error {
			var controller rest.BillinglistController
			controller.Init(c)
			controller.Sum()
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/billinguserlist/count", func(c *fiber.Ctx) error {
			var controller rest.BillinguserlistController
			controller.Init(c)
			controller.Count()
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/billinguserlist/:id", func(c *fiber.Ctx) error {
			id_, _ := strconv.ParseInt(c.Params("id"), 10, 64)
			var controller rest.BillinguserlistController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/billinguserlist", func(c *fiber.Ctx) error {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.BillinguserlistController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/billinguserlist/sum", func(c *fiber.Ctx) error {
			var controller rest.BillinguserlistController
			controller.Init(c)
			controller.Sum()
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/building/get/companyname/:company", func(c *fiber.Ctx) error {
			company_, _ := strconv.ParseInt(c.Params("company"), 10, 64)
			name_ := c.Query("name")
			var controller rest.BuildingController
			controller.Init(c)
			controller.GetByCompanyName(company_, name_)
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

		apiGroup.Get("/building/count", func(c *fiber.Ctx) error {
			var controller rest.BuildingController
			controller.Init(c)
			controller.Count()
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

		apiGroup.Put("/building/companyno", func(c *fiber.Ctx) error {
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
			var controller rest.BuildingController
			controller.Init(c)
			controller.UpdateCompanyno(companyno_, id_)
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

		apiGroup.Put("/building/zip", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var zip_ string
			if v, flag := results["zip"]; flag {
				zip_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.BuildingController
			controller.Init(c)
			controller.UpdateZip(zip_, id_)
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

		apiGroup.Put("/building/postzip", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var postzip_ string
			if v, flag := results["postzip"]; flag {
				postzip_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.BuildingController
			controller.Init(c)
			controller.UpdatePostzip(postzip_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/building/postaddress", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var postaddress_ string
			if v, flag := results["postaddress"]; flag {
				postaddress_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.BuildingController
			controller.Init(c)
			controller.UpdatePostaddress(postaddress_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/building/postaddressetc", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var postaddressetc_ string
			if v, flag := results["postaddressetc"]; flag {
				postaddressetc_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.BuildingController
			controller.Init(c)
			controller.UpdatePostaddressetc(postaddressetc_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/building/postname", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var postname_ string
			if v, flag := results["postname"]; flag {
				postname_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.BuildingController
			controller.Init(c)
			controller.UpdatePostname(postname_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/building/posttel", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var posttel_ string
			if v, flag := results["posttel"]; flag {
				posttel_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.BuildingController
			controller.Init(c)
			controller.UpdatePosttel(posttel_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/building/cmsnumber", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var cmsnumber_ string
			if v, flag := results["cmsnumber"]; flag {
				cmsnumber_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.BuildingController
			controller.Init(c)
			controller.UpdateCmsnumber(cmsnumber_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/building/cmsbank", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var cmsbank_ string
			if v, flag := results["cmsbank"]; flag {
				cmsbank_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.BuildingController
			controller.Init(c)
			controller.UpdateCmsbank(cmsbank_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/building/cmsaccountno", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var cmsaccountno_ string
			if v, flag := results["cmsaccountno"]; flag {
				cmsaccountno_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.BuildingController
			controller.Init(c)
			controller.UpdateCmsaccountno(cmsaccountno_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/building/cmsconfirm", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var cmsconfirm_ string
			if v, flag := results["cmsconfirm"]; flag {
				cmsconfirm_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.BuildingController
			controller.Init(c)
			controller.UpdateCmsconfirm(cmsconfirm_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/building/cmsstartdate", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var cmsstartdate_ string
			if v, flag := results["cmsstartdate"]; flag {
				cmsstartdate_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.BuildingController
			controller.Init(c)
			controller.UpdateCmsstartdate(cmsstartdate_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/building/cmsenddate", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var cmsenddate_ string
			if v, flag := results["cmsenddate"]; flag {
				cmsenddate_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.BuildingController
			controller.Init(c)
			controller.UpdateCmsenddate(cmsenddate_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/building/contractvolumn", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var contractvolumn_ models.Double
			contractvolumn__ref := &contractvolumn_
			c.BodyParser(contractvolumn__ref)
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.BuildingController
			controller.Init(c)
			controller.UpdateContractvolumn(contractvolumn_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/building/receivevolumn", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var receivevolumn_ models.Double
			receivevolumn__ref := &receivevolumn_
			c.BodyParser(receivevolumn__ref)
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.BuildingController
			controller.Init(c)
			controller.UpdateReceivevolumn(receivevolumn_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/building/generatevolumn", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var generatevolumn_ models.Double
			generatevolumn__ref := &generatevolumn_
			c.BodyParser(generatevolumn__ref)
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.BuildingController
			controller.Init(c)
			controller.UpdateGeneratevolumn(generatevolumn_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/building/sunlightvolumn", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var sunlightvolumn_ models.Double
			sunlightvolumn__ref := &sunlightvolumn_
			c.BodyParser(sunlightvolumn__ref)
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.BuildingController
			controller.Init(c)
			controller.UpdateSunlightvolumn(sunlightvolumn_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/building/volttype", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var volttype_ int
			if v, flag := results["volttype"]; flag {
				volttype_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.BuildingController
			controller.Init(c)
			controller.UpdateVolttype(volttype_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/building/weight", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var weight_ models.Double
			weight__ref := &weight_
			c.BodyParser(weight__ref)
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.BuildingController
			controller.Init(c)
			controller.UpdateWeight(weight_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/building/totalweight", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var totalweight_ models.Double
			totalweight__ref := &totalweight_
			c.BodyParser(totalweight__ref)
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.BuildingController
			controller.Init(c)
			controller.UpdateTotalweight(totalweight_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/building/checkcount", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var checkcount_ int
			if v, flag := results["checkcount"]; flag {
				checkcount_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.BuildingController
			controller.Init(c)
			controller.UpdateCheckcount(checkcount_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/building/receivevolt", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var receivevolt_ int
			if v, flag := results["receivevolt"]; flag {
				receivevolt_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.BuildingController
			controller.Init(c)
			controller.UpdateReceivevolt(receivevolt_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/building/generatevolt", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var generatevolt_ int
			if v, flag := results["generatevolt"]; flag {
				generatevolt_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.BuildingController
			controller.Init(c)
			controller.UpdateGeneratevolt(generatevolt_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/building/periodic", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var periodic_ int
			if v, flag := results["periodic"]; flag {
				periodic_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.BuildingController
			controller.Init(c)
			controller.UpdatePeriodic(periodic_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/building/businesscondition", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var businesscondition_ string
			if v, flag := results["businesscondition"]; flag {
				businesscondition_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.BuildingController
			controller.Init(c)
			controller.UpdateBusinesscondition(businesscondition_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/building/businessitem", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var businessitem_ string
			if v, flag := results["businessitem"]; flag {
				businessitem_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.BuildingController
			controller.Init(c)
			controller.UpdateBusinessitem(businessitem_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/building/usage", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var usage_ string
			if v, flag := results["usage"]; flag {
				usage_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.BuildingController
			controller.Init(c)
			controller.UpdateUsage(usage_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/building/district", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var district_ string
			if v, flag := results["district"]; flag {
				district_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.BuildingController
			controller.Init(c)
			controller.UpdateDistrict(district_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/building/check", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var check_ int
			if v, flag := results["check"]; flag {
				check_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.BuildingController
			controller.Init(c)
			controller.UpdateCheck(check_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/building/checkpost", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var checkpost_ int
			if v, flag := results["checkpost"]; flag {
				checkpost_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.BuildingController
			controller.Init(c)
			controller.UpdateCheckpost(checkpost_, id_)
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

		apiGroup.Get("/calendarcompanylist/count", func(c *fiber.Ctx) error {
			var controller rest.CalendarcompanylistController
			controller.Init(c)
			controller.Count()
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

		apiGroup.Get("/calendarcompanylist/sum", func(c *fiber.Ctx) error {
			var controller rest.CalendarcompanylistController
			controller.Init(c)
			controller.Sum()
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/company/get/companyno/:companyno", func(c *fiber.Ctx) error {
			companyno_ := c.Params("companyno")
			var controller rest.CompanyController
			controller.Init(c)
			controller.GetByCompanyno(companyno_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/company/get/name/:name", func(c *fiber.Ctx) error {
			name_ := c.Params("name")
			var controller rest.CompanyController
			controller.Init(c)
			controller.GetByName(name_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/company", func(c *fiber.Ctx) error {
			item_ := &models.Company{}
			c.BodyParser(item_)
			var apicontroller api.CompanyController
            apicontroller.Init(c)
			var controller rest.CompanyController
			controller.Init(c)
			if item_ != nil {
				controller.Insert(item_)
				apicontroller.Post_Insert(item_)
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

		apiGroup.Get("/company/count", func(c *fiber.Ctx) error {
			var controller rest.CompanyController
			controller.Init(c)
			controller.Count()
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

		apiGroup.Put("/company/tel", func(c *fiber.Ctx) error {
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
			var controller rest.CompanyController
			controller.Init(c)
			controller.UpdateTel(tel_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/company/email", func(c *fiber.Ctx) error {
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
			var controller rest.CompanyController
			controller.Init(c)
			controller.UpdateEmail(email_, id_)
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

		apiGroup.Put("/company/bankname", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var bankname_ string
			if v, flag := results["bankname"]; flag {
				bankname_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CompanyController
			controller.Init(c)
			controller.UpdateBankname(bankname_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/company/bankno", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var bankno_ string
			if v, flag := results["bankno"]; flag {
				bankno_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CompanyController
			controller.Init(c)
			controller.UpdateBankno(bankno_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/company/businesscondition", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var businesscondition_ string
			if v, flag := results["businesscondition"]; flag {
				businesscondition_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CompanyController
			controller.Init(c)
			controller.UpdateBusinesscondition(businesscondition_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/company/businessitem", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var businessitem_ string
			if v, flag := results["businessitem"]; flag {
				businessitem_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CompanyController
			controller.Init(c)
			controller.UpdateBusinessitem(businessitem_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/company/giro", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var giro_ string
			if v, flag := results["giro"]; flag {
				giro_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CompanyController
			controller.Init(c)
			controller.UpdateGiro(giro_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/company/egirologinid", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var egirologinid_ string
			if v, flag := results["egirologinid"]; flag {
				egirologinid_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CompanyController
			controller.Init(c)
			controller.UpdateEgirologinid(egirologinid_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/company/egiropasswd", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var egiropasswd_ string
			if v, flag := results["egiropasswd"]; flag {
				egiropasswd_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CompanyController
			controller.Init(c)
			controller.UpdateEgiropasswd(egiropasswd_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/company/content", func(c *fiber.Ctx) error {
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
			var controller rest.CompanyController
			controller.Init(c)
			controller.UpdateContent(content_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/company/x1", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var x1_ models.Double
			x1__ref := &x1_
			c.BodyParser(x1__ref)
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CompanyController
			controller.Init(c)
			controller.UpdateX1(x1_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/company/y1", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var y1_ models.Double
			y1__ref := &y1_
			c.BodyParser(y1__ref)
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CompanyController
			controller.Init(c)
			controller.UpdateY1(y1_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/company/x2", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var x2_ models.Double
			x2__ref := &x2_
			c.BodyParser(x2__ref)
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CompanyController
			controller.Init(c)
			controller.UpdateX2(x2_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/company/y2", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var y2_ models.Double
			y2__ref := &y2_
			c.BodyParser(y2__ref)
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CompanyController
			controller.Init(c)
			controller.UpdateY2(y2_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/company/x3", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var x3_ models.Double
			x3__ref := &x3_
			c.BodyParser(x3__ref)
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CompanyController
			controller.Init(c)
			controller.UpdateX3(x3_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/company/y3", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var y3_ models.Double
			y3__ref := &y3_
			c.BodyParser(y3__ref)
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CompanyController
			controller.Init(c)
			controller.UpdateY3(y3_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/company/x4", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var x4_ models.Double
			x4__ref := &x4_
			c.BodyParser(x4__ref)
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CompanyController
			controller.Init(c)
			controller.UpdateX4(x4_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/company/y4", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var y4_ models.Double
			y4__ref := &y4_
			c.BodyParser(y4__ref)
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CompanyController
			controller.Init(c)
			controller.UpdateY4(y4_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/company/x5", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var x5_ models.Double
			x5__ref := &x5_
			c.BodyParser(x5__ref)
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CompanyController
			controller.Init(c)
			controller.UpdateX5(x5_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/company/y5", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var y5_ models.Double
			y5__ref := &y5_
			c.BodyParser(y5__ref)
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CompanyController
			controller.Init(c)
			controller.UpdateY5(y5_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/company/x6", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var x6_ models.Double
			x6__ref := &x6_
			c.BodyParser(x6__ref)
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CompanyController
			controller.Init(c)
			controller.UpdateX6(x6_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/company/y6", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var y6_ models.Double
			y6__ref := &y6_
			c.BodyParser(y6__ref)
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CompanyController
			controller.Init(c)
			controller.UpdateY6(y6_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/company/x7", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var x7_ models.Double
			x7__ref := &x7_
			c.BodyParser(x7__ref)
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CompanyController
			controller.Init(c)
			controller.UpdateX7(x7_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/company/y7", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var y7_ models.Double
			y7__ref := &y7_
			c.BodyParser(y7__ref)
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CompanyController
			controller.Init(c)
			controller.UpdateY7(y7_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/company/x8", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var x8_ models.Double
			x8__ref := &x8_
			c.BodyParser(x8__ref)
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CompanyController
			controller.Init(c)
			controller.UpdateX8(x8_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/company/y8", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var y8_ models.Double
			y8__ref := &y8_
			c.BodyParser(y8__ref)
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CompanyController
			controller.Init(c)
			controller.UpdateY8(y8_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/company/x9", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var x9_ models.Double
			x9__ref := &x9_
			c.BodyParser(x9__ref)
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CompanyController
			controller.Init(c)
			controller.UpdateX9(x9_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/company/y9", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var y9_ models.Double
			y9__ref := &y9_
			c.BodyParser(y9__ref)
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CompanyController
			controller.Init(c)
			controller.UpdateY9(y9_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/company/x10", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var x10_ models.Double
			x10__ref := &x10_
			c.BodyParser(x10__ref)
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CompanyController
			controller.Init(c)
			controller.UpdateX10(x10_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/company/y10", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var y10_ models.Double
			y10__ref := &y10_
			c.BodyParser(y10__ref)
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CompanyController
			controller.Init(c)
			controller.UpdateY10(y10_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/company/x11", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var x11_ models.Double
			x11__ref := &x11_
			c.BodyParser(x11__ref)
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CompanyController
			controller.Init(c)
			controller.UpdateX11(x11_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/company/y11", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var y11_ models.Double
			y11__ref := &y11_
			c.BodyParser(y11__ref)
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CompanyController
			controller.Init(c)
			controller.UpdateY11(y11_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/company/x12", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var x12_ models.Double
			x12__ref := &x12_
			c.BodyParser(x12__ref)
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CompanyController
			controller.Init(c)
			controller.UpdateX12(x12_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/company/y12", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var y12_ models.Double
			y12__ref := &y12_
			c.BodyParser(y12__ref)
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CompanyController
			controller.Init(c)
			controller.UpdateY12(y12_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/company/x13", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var x13_ models.Double
			x13__ref := &x13_
			c.BodyParser(x13__ref)
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CompanyController
			controller.Init(c)
			controller.UpdateX13(x13_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/company/y13", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var y13_ models.Double
			y13__ref := &y13_
			c.BodyParser(y13__ref)
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CompanyController
			controller.Init(c)
			controller.UpdateY13(y13_, id_)
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

		apiGroup.Get("/companylicense/count", func(c *fiber.Ctx) error {
			var controller rest.CompanylicenseController
			controller.Init(c)
			controller.Count()
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

		apiGroup.Put("/companylicense/number", func(c *fiber.Ctx) error {
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
			var controller rest.CompanylicenseController
			controller.Init(c)
			controller.UpdateNumber(number_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/companylicense/takingdate", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var takingdate_ string
			if v, flag := results["takingdate"]; flag {
				takingdate_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CompanylicenseController
			controller.Init(c)
			controller.UpdateTakingdate(takingdate_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/companylicense/educationdate", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var educationdate_ string
			if v, flag := results["educationdate"]; flag {
				educationdate_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CompanylicenseController
			controller.Init(c)
			controller.UpdateEducationdate(educationdate_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/companylicense/educationinstitution", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var educationinstitution_ string
			if v, flag := results["educationinstitution"]; flag {
				educationinstitution_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CompanylicenseController
			controller.Init(c)
			controller.UpdateEducationinstitution(educationinstitution_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/companylicense/specialeducationdate", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var specialeducationdate_ string
			if v, flag := results["specialeducationdate"]; flag {
				specialeducationdate_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CompanylicenseController
			controller.Init(c)
			controller.UpdateSpecialeducationdate(specialeducationdate_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/companylicense/specialeducationinstitution", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var specialeducationinstitution_ string
			if v, flag := results["specialeducationinstitution"]; flag {
				specialeducationinstitution_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CompanylicenseController
			controller.Init(c)
			controller.UpdateSpecialeducationinstitution(specialeducationinstitution_, id_)
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

		apiGroup.Post("/companylicensecategory", func(c *fiber.Ctx) error {
			item_ := &models.Companylicensecategory{}
			c.BodyParser(item_)
			var controller rest.CompanylicensecategoryController
			controller.Init(c)
			if item_ != nil {
				controller.Insert(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/companylicensecategory/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Companylicensecategory{}
			c.BodyParser(item_)
			var controller rest.CompanylicensecategoryController
			controller.Init(c)
			if item_ != nil {
				controller.Insertbatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/companylicensecategory", func(c *fiber.Ctx) error {
			item_ := &models.Companylicensecategory{}
			c.BodyParser(item_)
			var controller rest.CompanylicensecategoryController
			controller.Init(c)
			if item_ != nil {
				controller.Update(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/companylicensecategory", func(c *fiber.Ctx) error {
			item_ := &models.Companylicensecategory{}
			c.BodyParser(item_)
			var controller rest.CompanylicensecategoryController
			controller.Init(c)
			if item_ != nil {
				controller.Delete(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/companylicensecategory/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Companylicensecategory{}
			c.BodyParser(item_)
			var controller rest.CompanylicensecategoryController
			controller.Init(c)
			if item_ != nil {
				controller.Deletebatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/companylicensecategory/count", func(c *fiber.Ctx) error {
			var controller rest.CompanylicensecategoryController
			controller.Init(c)
			controller.Count()
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/companylicensecategory/:id", func(c *fiber.Ctx) error {
			id_, _ := strconv.ParseInt(c.Params("id"), 10, 64)
			var controller rest.CompanylicensecategoryController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/companylicensecategory", func(c *fiber.Ctx) error {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.CompanylicensecategoryController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/companylicensecategory/name", func(c *fiber.Ctx) error {
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
			var controller rest.CompanylicensecategoryController
			controller.Init(c)
			controller.UpdateName(name_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/companylicensecategory/order", func(c *fiber.Ctx) error {
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
			var controller rest.CompanylicensecategoryController
			controller.Init(c)
			controller.UpdateOrder(order_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/companylist/count", func(c *fiber.Ctx) error {
			var controller rest.CompanylistController
			controller.Init(c)
			controller.Count()
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/companylist/:id", func(c *fiber.Ctx) error {
			id_, _ := strconv.ParseInt(c.Params("id"), 10, 64)
			var controller rest.CompanylistController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/companylist", func(c *fiber.Ctx) error {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.CompanylistController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/customer/count/companybuilding/:company", func(c *fiber.Ctx) error {
			company_, _ := strconv.ParseInt(c.Params("company"), 10, 64)
			building_, _ := strconv.ParseInt(c.Query("building"), 10, 64)
			var controller rest.CustomerController
			controller.Init(c)
			controller.CountByCompanyBuilding(company_, building_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/customer/get/companybuilding/:company", func(c *fiber.Ctx) error {
			company_, _ := strconv.ParseInt(c.Params("company"), 10, 64)
			building_, _ := strconv.ParseInt(c.Query("building"), 10, 64)
			var controller rest.CustomerController
			controller.Init(c)
			controller.GetByCompanyBuilding(company_, building_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/customer/bycompanybuilding", func(c *fiber.Ctx) error {
			item_ := &models.Customer{}
			c.BodyParser(item_)
			var controller rest.CustomerController
			controller.Init(c)
			controller.DeleteByCompanyBuilding(item_.Company, item_.Building)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/customer/bycompany", func(c *fiber.Ctx) error {
			item_ := &models.Customer{}
			c.BodyParser(item_)
			var controller rest.CustomerController
			controller.Init(c)
			controller.DeleteByCompany(item_.Company)
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

		apiGroup.Get("/customer/count", func(c *fiber.Ctx) error {
			var controller rest.CustomerController
			controller.Init(c)
			controller.Count()
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

		apiGroup.Put("/customer/number", func(c *fiber.Ctx) error {
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
			var controller rest.CustomerController
			controller.Init(c)
			controller.UpdateNumber(number_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/customer/kepconumber", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var kepconumber_ string
			if v, flag := results["kepconumber"]; flag {
				kepconumber_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CustomerController
			controller.Init(c)
			controller.UpdateKepconumber(kepconumber_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/customer/kesconumber", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var kesconumber_ string
			if v, flag := results["kesconumber"]; flag {
				kesconumber_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CustomerController
			controller.Init(c)
			controller.UpdateKesconumber(kesconumber_, id_)
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

		apiGroup.Put("/customer/usevat", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var usevat_ int
			if v, flag := results["usevat"]; flag {
				usevat_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CustomerController
			controller.Init(c)
			controller.UpdateUsevat(usevat_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/customer/contracttotalprice", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var contracttotalprice_ int
			if v, flag := results["contracttotalprice"]; flag {
				contracttotalprice_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CustomerController
			controller.Init(c)
			controller.UpdateContracttotalprice(contracttotalprice_, id_)
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

		apiGroup.Put("/customer/contractvat", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var contractvat_ int
			if v, flag := results["contractvat"]; flag {
				contractvat_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CustomerController
			controller.Init(c)
			controller.UpdateContractvat(contractvat_, id_)
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

		apiGroup.Put("/customer/contracttype", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var contracttype_ int
			if v, flag := results["contracttype"]; flag {
				contracttype_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CustomerController
			controller.Init(c)
			controller.UpdateContracttype(contracttype_, id_)
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

		apiGroup.Put("/customer/billingtype", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var billingtype_ int
			if v, flag := results["billingtype"]; flag {
				billingtype_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CustomerController
			controller.Init(c)
			controller.UpdateBillingtype(billingtype_, id_)
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

		apiGroup.Put("/customer/address", func(c *fiber.Ctx) error {
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
			var controller rest.CustomerController
			controller.Init(c)
			controller.UpdateAddress(address_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/customer/addressetc", func(c *fiber.Ctx) error {
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
			var controller rest.CustomerController
			controller.Init(c)
			controller.UpdateAddressetc(addressetc_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/customer/collectmonth", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var collectmonth_ int
			if v, flag := results["collectmonth"]; flag {
				collectmonth_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CustomerController
			controller.Init(c)
			controller.UpdateCollectmonth(collectmonth_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/customer/collectday", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var collectday_ int
			if v, flag := results["collectday"]; flag {
				collectday_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CustomerController
			controller.Init(c)
			controller.UpdateCollectday(collectday_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/customer/manager", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var manager_ string
			if v, flag := results["manager"]; flag {
				manager_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CustomerController
			controller.Init(c)
			controller.UpdateManager(manager_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/customer/tel", func(c *fiber.Ctx) error {
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
			var controller rest.CustomerController
			controller.Init(c)
			controller.UpdateTel(tel_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/customer/fax", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var fax_ string
			if v, flag := results["fax"]; flag {
				fax_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CustomerController
			controller.Init(c)
			controller.UpdateFax(fax_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/customer/periodic", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var periodic_ string
			if v, flag := results["periodic"]; flag {
				periodic_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CustomerController
			controller.Init(c)
			controller.UpdatePeriodic(periodic_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/customer/lastdate", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var lastdate_ string
			if v, flag := results["lastdate"]; flag {
				lastdate_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CustomerController
			controller.Init(c)
			controller.UpdateLastdate(lastdate_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/customer/remark", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var remark_ string
			if v, flag := results["remark"]; flag {
				remark_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CustomerController
			controller.Init(c)
			controller.UpdateRemark(remark_, id_)
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

		apiGroup.Put("/customer/salesuser", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var salesuser_ int64
			if v, flag := results["salesuser"]; flag {
				salesuser_ = int64(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CustomerController
			controller.Init(c)
			controller.UpdateSalesuser(salesuser_, id_)
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

		apiGroup.Get("/customercompany/get/companycustomer/:company", func(c *fiber.Ctx) error {
			company_, _ := strconv.ParseInt(c.Params("company"), 10, 64)
			customer_, _ := strconv.ParseInt(c.Query("customer"), 10, 64)
			var controller rest.CustomercompanyController
			controller.Init(c)
			controller.GetByCompanyCustomer(company_, customer_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/customercompany/bycompany", func(c *fiber.Ctx) error {
			item_ := &models.Customercompany{}
			c.BodyParser(item_)
			var controller rest.CustomercompanyController
			controller.Init(c)
			controller.DeleteByCompany(item_.Company)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/customercompany", func(c *fiber.Ctx) error {
			item_ := &models.Customercompany{}
			c.BodyParser(item_)
			var controller rest.CustomercompanyController
			controller.Init(c)
			if item_ != nil {
				controller.Insert(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/customercompany/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Customercompany{}
			c.BodyParser(item_)
			var controller rest.CustomercompanyController
			controller.Init(c)
			if item_ != nil {
				controller.Insertbatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/customercompany", func(c *fiber.Ctx) error {
			item_ := &models.Customercompany{}
			c.BodyParser(item_)
			var controller rest.CustomercompanyController
			controller.Init(c)
			if item_ != nil {
				controller.Update(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/customercompany", func(c *fiber.Ctx) error {
			item_ := &models.Customercompany{}
			c.BodyParser(item_)
			var controller rest.CustomercompanyController
			controller.Init(c)
			if item_ != nil {
				controller.Delete(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/customercompany/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Customercompany{}
			c.BodyParser(item_)
			var controller rest.CustomercompanyController
			controller.Init(c)
			if item_ != nil {
				controller.Deletebatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
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

		apiGroup.Put("/customercompany/company", func(c *fiber.Ctx) error {
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
			var controller rest.CustomercompanyController
			controller.Init(c)
			controller.UpdateCompany(company_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/customercompany/customer", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var customer_ int64
			if v, flag := results["customer"]; flag {
				customer_ = int64(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.CustomercompanyController
			controller.Init(c)
			controller.UpdateCustomer(customer_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/customercompanylist/count", func(c *fiber.Ctx) error {
			var controller rest.CustomercompanylistController
			controller.Init(c)
			controller.Count()
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/customercompanylist/:id", func(c *fiber.Ctx) error {
			id_, _ := strconv.ParseInt(c.Params("id"), 10, 64)
			var controller rest.CustomercompanylistController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/customercompanylist", func(c *fiber.Ctx) error {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.CustomercompanylistController
			controller.Init(c)
			controller.Index(page_, pagesize_)
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

		apiGroup.Get("/data/count", func(c *fiber.Ctx) error {
			var controller rest.DataController
			controller.Init(c)
			controller.Count()
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

		apiGroup.Get("/department/get/companyname/:company", func(c *fiber.Ctx) error {
			company_, _ := strconv.ParseInt(c.Params("company"), 10, 64)
			name_ := c.Query("name")
			var controller rest.DepartmentController
			controller.Init(c)
			controller.GetByCompanyName(company_, name_)
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

		apiGroup.Get("/department/count", func(c *fiber.Ctx) error {
			var controller rest.DepartmentController
			controller.Init(c)
			controller.Count()
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

		apiGroup.Put("/department/status", func(c *fiber.Ctx) error {
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
			var controller rest.DepartmentController
			controller.Init(c)
			controller.UpdateStatus(status_, id_)
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

		apiGroup.Put("/department/parent", func(c *fiber.Ctx) error {
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
			var controller rest.DepartmentController
			controller.Init(c)
			controller.UpdateParent(parent_, id_)
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

		apiGroup.Put("/department/master", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var master_ int64
			if v, flag := results["master"]; flag {
				master_ = int64(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.DepartmentController
			controller.Init(c)
			controller.UpdateMaster(master_, id_)
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

		apiGroup.Delete("/facility/bybuilding", func(c *fiber.Ctx) error {
			item_ := &models.Facility{}
			c.BodyParser(item_)
			var controller rest.FacilityController
			controller.Init(c)
			controller.DeleteByBuilding(item_.Building)
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

		apiGroup.Get("/facility/count", func(c *fiber.Ctx) error {
			var controller rest.FacilityController
			controller.Init(c)
			controller.Count()
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

		apiGroup.Put("/facility/type", func(c *fiber.Ctx) error {
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
			var controller rest.FacilityController
			controller.Init(c)
			controller.UpdateType(type_, id_)
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

		apiGroup.Post("/giro", func(c *fiber.Ctx) error {
			item_ := &models.Giro{}
			c.BodyParser(item_)
			var controller rest.GiroController
			controller.Init(c)
			if item_ != nil {
				controller.Insert(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/giro/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Giro{}
			c.BodyParser(item_)
			var controller rest.GiroController
			controller.Init(c)
			if item_ != nil {
				controller.Insertbatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/giro", func(c *fiber.Ctx) error {
			item_ := &models.Giro{}
			c.BodyParser(item_)
			var controller rest.GiroController
			controller.Init(c)
			if item_ != nil {
				controller.Update(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/giro", func(c *fiber.Ctx) error {
			item_ := &models.Giro{}
			c.BodyParser(item_)
			var controller rest.GiroController
			controller.Init(c)
			if item_ != nil {
				controller.Delete(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/giro/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Giro{}
			c.BodyParser(item_)
			var controller rest.GiroController
			controller.Init(c)
			if item_ != nil {
				controller.Deletebatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/giro/count", func(c *fiber.Ctx) error {
			var controller rest.GiroController
			controller.Init(c)
			controller.Count()
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/giro/:id", func(c *fiber.Ctx) error {
			id_, _ := strconv.ParseInt(c.Params("id"), 10, 64)
			var controller rest.GiroController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/giro", func(c *fiber.Ctx) error {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.GiroController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/giro/insertdate", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var insertdate_ string
			if v, flag := results["insertdate"]; flag {
				insertdate_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.GiroController
			controller.Init(c)
			controller.UpdateInsertdate(insertdate_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/giro/number", func(c *fiber.Ctx) error {
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
			var controller rest.GiroController
			controller.Init(c)
			controller.UpdateNumber(number_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/giro/price", func(c *fiber.Ctx) error {
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
			var controller rest.GiroController
			controller.Init(c)
			controller.UpdatePrice(price_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/giro/acceptdate", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var acceptdate_ string
			if v, flag := results["acceptdate"]; flag {
				acceptdate_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.GiroController
			controller.Init(c)
			controller.UpdateAcceptdate(acceptdate_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/giro/charge", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var charge_ int
			if v, flag := results["charge"]; flag {
				charge_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.GiroController
			controller.Init(c)
			controller.UpdateCharge(charge_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/giro/type", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var type_ string
			if v, flag := results["type"]; flag {
				type_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.GiroController
			controller.Init(c)
			controller.UpdateType(type_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/giro/content", func(c *fiber.Ctx) error {
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
			var controller rest.GiroController
			controller.Init(c)
			controller.UpdateContent(content_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/giro/sum", func(c *fiber.Ctx) error {
			var controller rest.GiroController
			controller.Init(c)
			controller.Sum()
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

		apiGroup.Get("/item/count", func(c *fiber.Ctx) error {
			var controller rest.ItemController
			controller.Init(c)
			controller.Count()
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

		apiGroup.Get("/license/get/userlicensecategory/:user", func(c *fiber.Ctx) error {
			user_, _ := strconv.ParseInt(c.Params("user"), 10, 64)
			licensecategory_, _ := strconv.ParseInt(c.Query("licensecategory"), 10, 64)
			var controller rest.LicenseController
			controller.Init(c)
			controller.GetByUserLicensecategory(user_, licensecategory_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/license/byuser", func(c *fiber.Ctx) error {
			item_ := &models.License{}
			c.BodyParser(item_)
			var controller rest.LicenseController
			controller.Init(c)
			controller.DeleteByUser(item_.User)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/license/find/user/:user", func(c *fiber.Ctx) error {
			user_, _ := strconv.ParseInt(c.Params("user"), 10, 64)
			var controller rest.LicenseController
			controller.Init(c)
			controller.FindByUser(user_)
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

		apiGroup.Get("/license/count", func(c *fiber.Ctx) error {
			var controller rest.LicenseController
			controller.Init(c)
			controller.Count()
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

		apiGroup.Put("/license/number", func(c *fiber.Ctx) error {
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
			var controller rest.LicenseController
			controller.Init(c)
			controller.UpdateNumber(number_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/license/takingdate", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var takingdate_ string
			if v, flag := results["takingdate"]; flag {
				takingdate_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.LicenseController
			controller.Init(c)
			controller.UpdateTakingdate(takingdate_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/license/educationdate", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var educationdate_ string
			if v, flag := results["educationdate"]; flag {
				educationdate_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.LicenseController
			controller.Init(c)
			controller.UpdateEducationdate(educationdate_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/license/educationinstitution", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var educationinstitution_ string
			if v, flag := results["educationinstitution"]; flag {
				educationinstitution_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.LicenseController
			controller.Init(c)
			controller.UpdateEducationinstitution(educationinstitution_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/license/specialeducationdate", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var specialeducationdate_ string
			if v, flag := results["specialeducationdate"]; flag {
				specialeducationdate_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.LicenseController
			controller.Init(c)
			controller.UpdateSpecialeducationdate(specialeducationdate_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/license/specialeducationinstitution", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var specialeducationinstitution_ string
			if v, flag := results["specialeducationinstitution"]; flag {
				specialeducationinstitution_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.LicenseController
			controller.Init(c)
			controller.UpdateSpecialeducationinstitution(specialeducationinstitution_, id_)
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

		apiGroup.Get("/licensecategory/get/name/:name", func(c *fiber.Ctx) error {
			name_ := c.Params("name")
			var controller rest.LicensecategoryController
			controller.Init(c)
			controller.GetByName(name_)
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

		apiGroup.Get("/licensecategory/count", func(c *fiber.Ctx) error {
			var controller rest.LicensecategoryController
			controller.Init(c)
			controller.Count()
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

		apiGroup.Get("/licenselevel/get/name/:name", func(c *fiber.Ctx) error {
			name_ := c.Params("name")
			var controller rest.LicenselevelController
			controller.Init(c)
			controller.GetByName(name_)
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

		apiGroup.Get("/licenselevel/count", func(c *fiber.Ctx) error {
			var controller rest.LicenselevelController
			controller.Init(c)
			controller.Count()
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

		apiGroup.Get("/notice/count", func(c *fiber.Ctx) error {
			var controller rest.NoticeController
			controller.Init(c)
			controller.Count()
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

		apiGroup.Get("/report/count", func(c *fiber.Ctx) error {
			var controller rest.ReportController
			controller.Init(c)
			controller.Count()
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

		apiGroup.Get("/reportlist/count", func(c *fiber.Ctx) error {
			var controller rest.ReportlistController
			controller.Init(c)
			controller.Count()
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/reportlist/:id", func(c *fiber.Ctx) error {
			id_, _ := strconv.ParseInt(c.Params("id"), 10, 64)
			var controller rest.ReportlistController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/reportlist", func(c *fiber.Ctx) error {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.ReportlistController
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

		apiGroup.Get("/statisticsday/sum", func(c *fiber.Ctx) error {
			var controller rest.StatisticsdayController
			controller.Init(c)
			controller.Sum()
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

		apiGroup.Get("/statisticsmonth/sum", func(c *fiber.Ctx) error {
			var controller rest.StatisticsmonthController
			controller.Init(c)
			controller.Sum()
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

		apiGroup.Get("/statisticsyear/sum", func(c *fiber.Ctx) error {
			var controller rest.StatisticsyearController
			controller.Init(c)
			controller.Sum()
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

		apiGroup.Get("/user/get/companyname/:company", func(c *fiber.Ctx) error {
			company_, _ := strconv.ParseInt(c.Params("company"), 10, 64)
			name_ := c.Query("name")
			var controller rest.UserController
			controller.Init(c)
			controller.GetByCompanyName(company_, name_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/user/get/companytelname/:company", func(c *fiber.Ctx) error {
			company_, _ := strconv.ParseInt(c.Params("company"), 10, 64)
			tel_ := c.Query("tel")
			name_ := c.Query("name")
			var controller rest.UserController
			controller.Init(c)
			controller.GetByCompanyTelName(company_, tel_, name_)
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

		apiGroup.Get("/user/count", func(c *fiber.Ctx) error {
			var controller rest.UserController
			controller.Init(c)
			controller.Count()
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

		apiGroup.Put("/user/zip", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var zip_ string
			if v, flag := results["zip"]; flag {
				zip_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.UserController
			controller.Init(c)
			controller.UpdateZip(zip_, id_)
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

		apiGroup.Put("/user/approval", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var approval_ int
			if v, flag := results["approval"]; flag {
				approval_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.UserController
			controller.Init(c)
			controller.UpdateApproval(approval_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/user/educationdate", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var educationdate_ string
			if v, flag := results["educationdate"]; flag {
				educationdate_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.UserController
			controller.Init(c)
			controller.UpdateEducationdate(educationdate_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/user/educationinstitution", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var educationinstitution_ string
			if v, flag := results["educationinstitution"]; flag {
				educationinstitution_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.UserController
			controller.Init(c)
			controller.UpdateEducationinstitution(educationinstitution_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/user/specialeducationdate", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var specialeducationdate_ string
			if v, flag := results["specialeducationdate"]; flag {
				specialeducationdate_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.UserController
			controller.Init(c)
			controller.UpdateSpecialeducationdate(specialeducationdate_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/user/specialeducationinstitution", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var specialeducationinstitution_ string
			if v, flag := results["specialeducationinstitution"]; flag {
				specialeducationinstitution_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.UserController
			controller.Init(c)
			controller.UpdateSpecialeducationinstitution(specialeducationinstitution_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/user/rejectreason", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var rejectreason_ string
			if v, flag := results["rejectreason"]; flag {
				rejectreason_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.UserController
			controller.Init(c)
			controller.UpdateRejectreason(rejectreason_, id_)
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

		apiGroup.Get("/userlist/count", func(c *fiber.Ctx) error {
			var controller rest.UserlistController
			controller.Init(c)
			controller.Count()
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

		apiGroup.Get("/userlist/sum", func(c *fiber.Ctx) error {
			var controller rest.UserlistController
			controller.Init(c)
			controller.Sum()
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/webfaq", func(c *fiber.Ctx) error {
			item_ := &models.Webfaq{}
			c.BodyParser(item_)
			var controller rest.WebfaqController
			controller.Init(c)
			if item_ != nil {
				controller.Insert(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/webfaq/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Webfaq{}
			c.BodyParser(item_)
			var controller rest.WebfaqController
			controller.Init(c)
			if item_ != nil {
				controller.Insertbatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/webfaq", func(c *fiber.Ctx) error {
			item_ := &models.Webfaq{}
			c.BodyParser(item_)
			var controller rest.WebfaqController
			controller.Init(c)
			if item_ != nil {
				controller.Update(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/webfaq", func(c *fiber.Ctx) error {
			item_ := &models.Webfaq{}
			c.BodyParser(item_)
			var controller rest.WebfaqController
			controller.Init(c)
			if item_ != nil {
				controller.Delete(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/webfaq/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Webfaq{}
			c.BodyParser(item_)
			var controller rest.WebfaqController
			controller.Init(c)
			if item_ != nil {
				controller.Deletebatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/webfaq/count", func(c *fiber.Ctx) error {
			var controller rest.WebfaqController
			controller.Init(c)
			controller.Count()
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/webfaq/:id", func(c *fiber.Ctx) error {
			id_, _ := strconv.ParseInt(c.Params("id"), 10, 64)
			var controller rest.WebfaqController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/webfaq", func(c *fiber.Ctx) error {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.WebfaqController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/webfaq/category", func(c *fiber.Ctx) error {
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
			var controller rest.WebfaqController
			controller.Init(c)
			controller.UpdateCategory(category_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/webfaq/title", func(c *fiber.Ctx) error {
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
			var controller rest.WebfaqController
			controller.Init(c)
			controller.UpdateTitle(title_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/webfaq/content", func(c *fiber.Ctx) error {
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
			var controller rest.WebfaqController
			controller.Init(c)
			controller.UpdateContent(content_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/webfaq/order", func(c *fiber.Ctx) error {
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
			var controller rest.WebfaqController
			controller.Init(c)
			controller.UpdateOrder(order_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/webjoin", func(c *fiber.Ctx) error {
			item_ := &models.Webjoin{}
			c.BodyParser(item_)
			var controller rest.WebjoinController
			controller.Init(c)
			if item_ != nil {
				controller.Insert(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/webjoin/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Webjoin{}
			c.BodyParser(item_)
			var controller rest.WebjoinController
			controller.Init(c)
			if item_ != nil {
				controller.Insertbatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/webjoin", func(c *fiber.Ctx) error {
			item_ := &models.Webjoin{}
			c.BodyParser(item_)
			var controller rest.WebjoinController
			controller.Init(c)
			if item_ != nil {
				controller.Update(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/webjoin", func(c *fiber.Ctx) error {
			item_ := &models.Webjoin{}
			c.BodyParser(item_)
			var controller rest.WebjoinController
			controller.Init(c)
			if item_ != nil {
				controller.Delete(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/webjoin/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Webjoin{}
			c.BodyParser(item_)
			var controller rest.WebjoinController
			controller.Init(c)
			if item_ != nil {
				controller.Deletebatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/webjoin/count", func(c *fiber.Ctx) error {
			var controller rest.WebjoinController
			controller.Init(c)
			controller.Count()
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/webjoin/:id", func(c *fiber.Ctx) error {
			id_, _ := strconv.ParseInt(c.Params("id"), 10, 64)
			var controller rest.WebjoinController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/webjoin", func(c *fiber.Ctx) error {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.WebjoinController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/webjoin/category", func(c *fiber.Ctx) error {
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
			var controller rest.WebjoinController
			controller.Init(c)
			controller.UpdateCategory(category_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/webjoin/name", func(c *fiber.Ctx) error {
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
			var controller rest.WebjoinController
			controller.Init(c)
			controller.UpdateName(name_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/webjoin/manager", func(c *fiber.Ctx) error {
			var results map[string]interface{}
			jsonData := c.Body()
			json.Unmarshal(jsonData, &results)
			var manager_ string
			if v, flag := results["manager"]; flag {
				manager_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.WebjoinController
			controller.Init(c)
			controller.UpdateManager(manager_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/webjoin/tel", func(c *fiber.Ctx) error {
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
			var controller rest.WebjoinController
			controller.Init(c)
			controller.UpdateTel(tel_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/webjoin/email", func(c *fiber.Ctx) error {
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
			var controller rest.WebjoinController
			controller.Init(c)
			controller.UpdateEmail(email_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/webnotice", func(c *fiber.Ctx) error {
			item_ := &models.Webnotice{}
			c.BodyParser(item_)
			var controller rest.WebnoticeController
			controller.Init(c)
			if item_ != nil {
				controller.Insert(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/webnotice/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Webnotice{}
			c.BodyParser(item_)
			var controller rest.WebnoticeController
			controller.Init(c)
			if item_ != nil {
				controller.Insertbatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/webnotice", func(c *fiber.Ctx) error {
			item_ := &models.Webnotice{}
			c.BodyParser(item_)
			var controller rest.WebnoticeController
			controller.Init(c)
			if item_ != nil {
				controller.Update(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/webnotice", func(c *fiber.Ctx) error {
			item_ := &models.Webnotice{}
			c.BodyParser(item_)
			var controller rest.WebnoticeController
			controller.Init(c)
			if item_ != nil {
				controller.Delete(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/webnotice/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Webnotice{}
			c.BodyParser(item_)
			var controller rest.WebnoticeController
			controller.Init(c)
			if item_ != nil {
				controller.Deletebatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/webnotice/count", func(c *fiber.Ctx) error {
			var controller rest.WebnoticeController
			controller.Init(c)
			controller.Count()
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/webnotice/:id", func(c *fiber.Ctx) error {
			id_, _ := strconv.ParseInt(c.Params("id"), 10, 64)
			var controller rest.WebnoticeController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/webnotice", func(c *fiber.Ctx) error {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.WebnoticeController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/webnotice/title", func(c *fiber.Ctx) error {
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
			var controller rest.WebnoticeController
			controller.Init(c)
			controller.UpdateTitle(title_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/webnotice/content", func(c *fiber.Ctx) error {
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
			var controller rest.WebnoticeController
			controller.Init(c)
			controller.UpdateContent(content_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/webnotice/image", func(c *fiber.Ctx) error {
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
			var controller rest.WebnoticeController
			controller.Init(c)
			controller.UpdateImage(image_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/webnotice/category", func(c *fiber.Ctx) error {
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
			var controller rest.WebnoticeController
			controller.Init(c)
			controller.UpdateCategory(category_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/webquestion", func(c *fiber.Ctx) error {
			item_ := &models.Webquestion{}
			c.BodyParser(item_)
			var controller rest.WebquestionController
			controller.Init(c)
			if item_ != nil {
				controller.Insert(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Post("/webquestion/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Webquestion{}
			c.BodyParser(item_)
			var controller rest.WebquestionController
			controller.Init(c)
			if item_ != nil {
				controller.Insertbatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/webquestion", func(c *fiber.Ctx) error {
			item_ := &models.Webquestion{}
			c.BodyParser(item_)
			var controller rest.WebquestionController
			controller.Init(c)
			if item_ != nil {
				controller.Update(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/webquestion", func(c *fiber.Ctx) error {
			item_ := &models.Webquestion{}
			c.BodyParser(item_)
			var controller rest.WebquestionController
			controller.Init(c)
			if item_ != nil {
				controller.Delete(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Delete("/webquestion/batch", func(c *fiber.Ctx) error {
			item_ := &[]models.Webquestion{}
			c.BodyParser(item_)
			var controller rest.WebquestionController
			controller.Init(c)
			if item_ != nil {
				controller.Deletebatch(item_)
			} else {
			    controller.Result["code"] = "error"
			}
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/webquestion/count", func(c *fiber.Ctx) error {
			var controller rest.WebquestionController
			controller.Init(c)
			controller.Count()
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/webquestion/:id", func(c *fiber.Ctx) error {
			id_, _ := strconv.ParseInt(c.Params("id"), 10, 64)
			var controller rest.WebquestionController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Get("/webquestion", func(c *fiber.Ctx) error {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.WebquestionController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/webquestion/email", func(c *fiber.Ctx) error {
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
			var controller rest.WebquestionController
			controller.Init(c)
			controller.UpdateEmail(email_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/webquestion/tel", func(c *fiber.Ctx) error {
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
			var controller rest.WebquestionController
			controller.Init(c)
			controller.UpdateTel(tel_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

		apiGroup.Put("/webquestion/content", func(c *fiber.Ctx) error {
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
			var controller rest.WebquestionController
			controller.Init(c)
			controller.UpdateContent(content_, id_)
			controller.Close()
			return c.JSON(controller.Result)
		})

	}

}