package api

import (
	"fmt"
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
	user := c.Session

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
	facilityManager := models.NewFacilityManager(conn)

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
		log.Println("POS:", pos)
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
				userItem.Loginid = userItem.Name
				userItem.Passwd = "0000"
				userItem.Status = user.StatusUse
				userItem.Approval = user.ApprovalComplete
				userItem.Score = 60

				userManager.Insert(&userItem)
				userId = userManager.GetIdentity()
			} else {
				userId = userFind.Id
			}
		}

		salesuserName := f.GetCell("N", pos)
		var salesuserId int64 = 0

		if salesuserName != "" {
			userFind := userManager.GetByCompanyName(myCompanyId, salesuserName)

			if userFind == nil {
				userItem := models.User{}
				userItem.Level = user.LevelNormal
				userItem.Company = myCompanyId
				userItem.Name = salesuserName
				userItem.Loginid = userItem.Name
				userItem.Passwd = "0000"
				userItem.Status = user.StatusUse
				userItem.Approval = user.ApprovalComplete
				userItem.Score = 60

				userManager.Insert(&userItem)
				salesuserId = userManager.GetIdentity()
			} else {
				salesuserId = userFind.Id
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

		if item.Companyno == "" {
			companyFind := companyManager.GetByName(item.Name)

			if companyFind == nil {
				companyManager.Insert(&item)
				companyId = companyManager.GetIdentity()
			} else {
				companyId = companyFind.Id
			}

		} else {
			companyFind := companyManager.GetByCompanyno(item.Companyno)

			if companyFind == nil {
				companyManager.Insert(&item)
				companyId = companyManager.GetIdentity()
			} else {
				companyId = companyFind.Id
			}
		}

		customercompany := customercompanyManager.GetByCompanyCustomer(myCompanyId, companyId)

		if customercompany == nil {
			customercompanyManager.Insert(&models.Customercompany{Company: myCompanyId, Customer: companyId})
		}

		building.Name = f.GetCell("C", pos)
		building.Address = f.GetCell("D", pos)
		building.Contractvolumn = models.Double(global.Atol(f.GetCell("E", pos)))
		building.Receivevolumn = models.Double(global.Atol(f.GetCell("F", pos)))
		building.Generatevolumn = models.Double(global.Atol(f.GetCell("G", pos)))
		building.Sunlightvolumn = models.Double(global.Atol(f.GetCell("H", pos)))
		building.Ceo = f.GetCell("AN", pos)

		weight := global.Atof(f.GetCell("E", pos))
		building.Weight = models.Double(weight)
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

		basic := global.Atoi(f.GetCell("F", pos))
		generator := global.Atoi(f.GetCell("G", pos))
		sunlight := global.Atoi(f.GetCell("H", pos))

		basicFacility := models.Facility{}
		generatorFacility := models.Facility{}
		sunlightFacility := models.Facility{}

		if basic > 0 {
			basicFacility.Category = 10
			basicFacility.Value2 = fmt.Sprintf("%v", basic)
		}

		if generator > 0 {
			generatorFacility.Category = 20
			generatorFacility.Value3 = fmt.Sprintf("%v", generator)
		}

		if sunlight > 0 {
			sunlightFacility.Category = 30
			sunlightFacility.Value6 = fmt.Sprintf("%v", sunlight)
		}

		building.Totalweight = models.Double(basic + generator + sunlight)

		var buildingId int64 = 0

		buildingFind := buildingManager.GetByCompanyName(companyId, building.Name)
		if buildingFind == nil {
			buildingManager.Insert(&building)
			buildingId = buildingManager.GetIdentity()
		} else {
			buildingId = buildingFind.Id
		}

		if basic > 0 {
			basicFacility.Building = buildingId

			facilityManager.Insert(&basicFacility)
		}

		if generator > 0 {
			generatorFacility.Building = buildingId

			facilityManager.Insert(&generatorFacility)
		}

		if sunlight > 0 {
			sunlightFacility.Building = buildingId

			facilityManager.Insert(&sunlightFacility)
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
		customerItem.Salesuser = salesuserId
		customerItem.Company = myCompanyId

		customerManager.DeleteByCompanyBuilding(myCompanyId, buildingId)
		customerManager.Insert(&customerItem)

		pos++
	}

	conn.Commit()
}

func (c *ExternalController) User(filename string) {
	session := c.Session
	company := session.Company

	db := models.NewConnection()
	defer db.Close()

	conn, _ := db.Begin()
	defer conn.Rollback()

	userManager := models.NewUserManager(conn)
	licenseManager := models.NewLicenseManager(conn)
	licensecategoryManager := models.NewLicensecategoryManager(conn)

	fullFilename := path.Join(config.UploadPath, filename)
	f := global.NewExcelReader(fullFilename)
	if f == nil {
		log.Println("not found file")
		return
	}

	sheet := "안전관리자목록"
	f.SetSheet(sheet)

	pos := 4
	for {
		item := models.User{}

		no := f.GetCell("A", pos)

		if no == "" {
			break
		}

		name := f.GetCell("C", pos)
		licenseno := f.GetCell("D", pos)
		zip := f.GetCell("F", pos)
		address := f.GetCell("G", pos)
		tel := f.GetCell("H", pos)
		email := f.GetCell("I", pos)
		license := f.GetCell("j", pos)

		educationdate := f.GetCell("M", pos)
		educationinstitution := f.GetCell("N", pos)
		specialeducationdate := f.GetCell("O", pos)
		specialeducationinstitution := f.GetCell("P", pos)
		joindate := f.GetCell("Q", pos)
		status := f.GetCell("R", pos)

		licensename := ""
		temp := strings.Fields(license)
		if len(temp) >= 2 {
			licensename = temp[0]
			licenseno = temp[1]
		} else {
			licensename = license
		}

		userItem := userManager.GetByCompanyName(company, name)

		if userItem != nil {
			item = *userItem
		}

		item.Name = name
		item.Zip = zip
		item.Address = address
		item.Tel = tel
		item.Email = email
		item.Educationdate = educationdate
		item.Educationinstitution = educationinstitution
		item.Specialeducationdate = specialeducationdate
		item.Specialeducationinstitution = specialeducationinstitution
		item.Joindate = joindate
		item.Approval = user.ApprovalComplete
		item.Level = user.LevelNormal

		if status == "재직" {
			item.Status = user.StatusUse
		} else {
			item.Status = user.StatusNotuse
		}

		if userItem == nil {
			item.Company = company
			item.Loginid = item.Name
			item.Passwd = "0000"
			item.Score = 60

			userManager.Insert(&item)
			item.Id = userManager.GetIdentity()
		} else {
			userManager.Update(&item)
		}

		licensecategoryItem := licensecategoryManager.GetByName(licensename)
		if licensecategoryItem == nil {
			licensecategoryManager.Insert(&models.Licensecategory{Name: licensename, Order: 0})
			licensecategory := licensecategoryManager.GetIdentity()

			licenseManager.Insert(&models.License{User: session.Id, Number: licenseno, Licensecategory: licensecategory})
		} else {
			licenseItem := licenseManager.GetByUserLicensecategory(session.Id, licensecategoryItem.Id)

			if licenseItem == nil {
				licenseManager.Insert(&models.License{User: session.Id, Number: licenseno, Licensecategory: licensecategoryItem.Id})
			}
		}

		pos++
	}

	conn.Commit()
}
