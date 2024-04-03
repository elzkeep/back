package api

import (
	"log"
	"path"
	"regexp"
	"strings"
	"zkeep/config"
	"zkeep/controllers"
	"zkeep/global"
	"zkeep/models"
	"zkeep/models/company"
	"zkeep/models/customer"
	"zkeep/models/user"
)

type ExternalController struct {
	controllers.Controller
}

func (c *ExternalController) Index(filenames string, typeid int) {
	log.Println("external index")
	log.Println(filenames)

	user := c.Session

	log.Println(user)

	files := strings.Split(filenames, ",")

	for _, v := range files {
		ExcelProcess(v, typeid, user.Company)
	}
}

func ExcelProcess(filename string, typeid int, myCompanyId int64) {
	db := models.NewConnection()
	defer db.Close()

	conn, _ := db.Begin()
	defer conn.Rollback()

	companyManager := models.NewCompanyManager(conn)
	customercompanyManager := models.NewCustomercompanyManager(conn)
	buildingManager := models.NewBuildingManager(conn)
	customerManager := models.NewCustomerManager(conn)
	userManager := models.NewUserManager(conn)

	fullFilename := path.Join(config.UploadPath, filename)
	f := global.NewExcelReader(fullFilename)
	if f == nil {
		log.Println("not found file")
		return
	}

	sheet := "수용가현황"
	f.SetSheet(sheet)

	pos := 5
	for {
		item := models.Company{}
		building := models.Building{}
		customerItem := models.Customer{}

		no := f.GetCell("A", pos)

		if no == "" {
			break
		}

		userName := f.GetCell("M", pos)
		if userName == "" {
			userName = f.GetCell("L", pos)
		}

		var userId int64 = 0

		if userName != "" {
			userFind := userManager.GetByCompanyName(myCompanyId, userName)

			if userFind == nil {
				userItem := models.User{}
				userItem.Level = user.LevelNormal
				userItem.Company = myCompanyId
				userItem.Name = userName
				userItem.Loginid = item.Name
				userItem.Passwd = "0000"
				userItem.Status = user.StatusUse

				userManager.Insert(&userItem)
				userId = userManager.GetIdentity()
			} else {
				userId = userFind.Id
			}
		}

		item.Name = f.GetCell("Y", pos)
		item.Companyno = f.GetCell("Z", pos)
		item.Ceo = f.GetCell("AA", pos)
		item.Businesscondition = f.GetCell("AB", pos)
		item.Businessitem = f.GetCell("AC", pos)
		item.Address = f.GetCell("AD", pos)
		item.Type = company.TypeBuilding

		var companyId int64 = 0

		companyFind := companyManager.GetByCompanyno(item.Companyno)

		if companyFind == nil {
			companyManager.Insert(&item)
			companyId = companyManager.GetIdentity()
		} else {
			companyId = companyFind.Id
		}

		customercompany := customercompanyManager.GetByCompanyCustomer(myCompanyId, companyId)

		if customercompany == nil {
			log.Println("my :", myCompanyId, "customer :", companyId)
			customercompanyManager.Insert(&models.Customercompany{Company: myCompanyId, Customer: companyId})
		}

		building.Name = f.GetCell("C", pos)
		building.Address = f.GetCell("D", pos)
		building.Contractvolumn = models.Double(global.Atol(f.GetCell("E", pos)))
		building.Receivevolumn = models.Double(global.Atol(f.GetCell("F", pos)))
		building.Generatevolumn = models.Double(global.Atol(f.GetCell("G", pos)))
		building.Sunlightvolumn = models.Double(global.Atol(f.GetCell("H", pos)))
		building.Ceo = f.GetCell("AN", pos)

		volttype := f.GetCell("I", pos)

		if volttype == "고압" {
			building.Volttype = 2
		} else {
			building.Volttype = 1
		}

		building.Checkcount = global.Atoi(f.GetCell("K", pos))

		building.Receivevolt = global.Atoi(strings.ReplaceAll(f.GetCell("O", pos), "V", ""))
		building.Usage = f.GetCell("T", pos)
		building.District = f.GetCell("U", pos)
		building.Company = companyId

		var buildingId int64 = 0

		buildingFind := buildingManager.GetByCompanyName(companyId, building.Name)
		if buildingFind == nil {
			buildingManager.Insert(&building)
			buildingId = buildingManager.GetIdentity()
		} else {
			buildingId = buildingFind.Id
		}

		customerItem.Managername = f.GetCell("V", pos)
		customerItem.Managertel = f.GetCell("W", pos)
		customerItem.Manageremail = f.GetCell("X", pos)
		customerItem.Address = f.GetCell("AM", pos)
		customerItem.Manager = f.GetCell("AN", pos)
		customerItem.Contractprice = global.Atoi(f.GetCell("AQ", pos))
		customerItem.Contractvat = global.Atoi(f.GetCell("AR", pos))
		customerItem.Status = typeid
		customerItem.Contractstartdate = strings.ReplaceAll(f.GetCell("AE", pos), ".", "-")
		customerItem.Contractenddate = strings.ReplaceAll(f.GetCell("AF", pos), ".", "-")
		customerItem.Type = customer.TypeOutsourcing

		str := f.GetCell("AU", pos)

		r, _ := regexp.Compile("[0-9]+")
		collectday := r.FindString(str)

		if strings.Contains(str, "당월") {
			customerItem.Collectmonth = 1
		} else {
			customerItem.Collectmonth = 2
		}

		if collectday == "" {
			customerItem.Collectday = 0
		} else {
			customerItem.Collectday = global.Atoi(collectday)
		}

		if f.GetCell("AT", pos) == "지로" {
			customerItem.Billingtype = 1
		} else {
			customerItem.Billingtype = 2
		}

		customerItem.Building = buildingId
		customerItem.User = userId
		customerItem.Company = myCompanyId

		customerManager.DeleteByCompanyBuilding(myCompanyId, buildingId)
		customerManager.Insert(&customerItem)

		pos++
	}

	conn.Commit()
}
