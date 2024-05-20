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

	list := f.File.GetSheetList()

	sheet := list[0]
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
			generatorFacility.Value12 = fmt.Sprintf("%v", generator)
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

		customerItem.Number = global.Atoi(f.GetCell("B", pos))
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
	departmentManager := models.NewDepartmentManager(conn)
	licenseManager := models.NewLicenseManager(conn)
	licensecategoryManager := models.NewLicensecategoryManager(conn)

	fullFilename := path.Join(config.UploadPath, filename)
	f := global.NewExcelReader(fullFilename)
	if f == nil {
		log.Println("not found file")
		return
	}

	list := f.File.GetSheetList()

	sheet := list[0]
	f.SetSheet(sheet)

	start := true
	pos := 1
	for {
		item := models.User{}

		no := f.GetCell("A", pos)

		if start == true {
			if no == "No." {
				start = false
			}

			pos++
			continue
		}

		if no == "" {
			break
		}

		name := f.GetCell("C", pos)
		licenseno := f.GetCell("L", pos)
		zip := f.GetCell("G", pos)
		address := f.GetCell("H", pos)
		tel := f.GetCell("I", pos)
		email := f.GetCell("J", pos)
		licensename := f.GetCell("K", pos)

		educationdate := f.GetCell("N", pos)
		educationinstitution := f.GetCell("P", pos)
		specialeducationdate := ""
		specialeducationinstitution := ""
		joindate := f.GetCell("S", pos)
		status := f.GetCell("T", pos)

		userItem := userManager.GetByCompanyName(company, name)

		if userItem != nil {
			item = *userItem
		}

		departmentName := f.GetCell("E", pos)
		department := departmentManager.GetByCompanyName(session.Company, departmentName)
		if department == nil {
			department = &models.Department{
				Name:    departmentName,
				Status:  1,
				Company: session.Company,
			}

			departmentManager.Insert(department)

			department.Id = departmentManager.GetIdentity()
		}

		item.Department = department.Id
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
			licensecategoryItem.Id = licensecategoryManager.GetIdentity()
		}

		licenseItem := licenseManager.GetByUserLicensecategory(item.Id, licensecategoryItem.Id)

		if licenseItem == nil {
			//licenseManager.DeleteByUser(item.Id)
			licenseManager.Insert(&models.License{User: item.Id, Number: licenseno, Licensecategory: licensecategoryItem.Id})
		}

		pos++
	}

	conn.Commit()
}

func (c *ExternalController) All(filename string) {
	session := c.Session

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

	sheet := "수용가 현황"
	f.SetSheet(sheet)

	pos := 2
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

		var userId int64 = 0

		if userName != "" {
			userFind := userManager.GetByCompanyName(session.Company, userName)

			if userFind == nil {
				userItem := models.User{}
				userItem.Level = user.LevelNormal
				userItem.Company = session.Company
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
			userFind := userManager.GetByCompanyName(session.Company, salesuserName)

			if userFind == nil {
				userItem := models.User{}
				userItem.Level = user.LevelNormal
				userItem.Company = session.Company
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

		item.Name = f.GetCell("T", pos)
		item.Companyno = f.GetCell("U", pos)
		item.Ceo = f.GetCell("V", pos)
		item.Businesscondition = f.GetCell("W", pos)
		item.Businessitem = f.GetCell("X", pos)
		item.Address = f.GetCell("Y", pos)
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

		customercompany := customercompanyManager.GetByCompanyCustomer(session.Company, companyId)

		if customercompany == nil {
			customercompanyManager.Insert(&models.Customercompany{Company: session.Company, Customer: companyId})
		}

		building.Name = f.GetCell("B", pos)
		building.Address = f.GetCell("C", pos)
		building.Contractvolumn = models.Double(global.Atol(f.GetCell("F", pos)))
		building.Receivevolumn = models.Double(global.Atol(f.GetCell("G", pos)))
		building.Generatevolumn = models.Double(global.Atol(f.GetCell("H", pos)))
		building.Sunlightvolumn = models.Double(global.Atol(f.GetCell("I", pos)))
		building.Ceo = f.GetCell("D", pos)

		weight := global.Atof(f.GetCell("K", pos))
		building.Weight = models.Double(weight)
		/*
			volttype := f.GetCell("I", pos)

			if volttype == "고압" {
				building.Volttype = 2
			} else {
				building.Volttype = 1
			}
		*/

		building.Checkcount = global.Atoi(f.GetCell("L", pos))

		building.Receivevolt = global.Atoi(strings.ReplaceAll(f.GetCell("O", pos), "V", ""))
		building.Usage = f.GetCell("R", pos)
		building.District = f.GetCell("S", pos)
		building.Company = companyId

		basic := global.Atoi(f.GetCell("G", pos))
		generator := global.Atoi(f.GetCell("H", pos))
		sunlight := global.Atoi(f.GetCell("I", pos))

		basicFacility := models.Facility{}
		generatorFacility := models.Facility{}
		sunlightFacility := models.Facility{}

		if basic > 0 {
			basicFacility.Category = 10
			basicFacility.Value2 = fmt.Sprintf("%v", basic)
		}

		if generator > 0 {
			generatorFacility.Category = 20
			generatorFacility.Value12 = fmt.Sprintf("%v", generator)
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

		customerItem.Number = global.Atoi(f.GetCell("A", pos))
		customerItem.Managername = f.GetCell("AD", pos)
		customerItem.Managertel = f.GetCell("AE", pos)
		customerItem.Manageremail = f.GetCell("AF", pos)
		customerItem.Address = f.GetCell("Y", pos)
		customerItem.Manager = f.GetCell("AD", pos)
		customerItem.Contractprice = global.Atoi(f.GetCell("AJ", pos))
		customerItem.Contractvat = global.Atoi(f.GetCell("AK", pos))

		if f.GetCell("AB", pos) != "" || f.GetCell("AC", pos) != "" {
			customerItem.Status = 2
		} else {
			customerItem.Status = 1

		}

		customerItem.Contractstartdate = strings.ReplaceAll(f.GetCell("AA", pos), ".", "-")
		customerItem.Contractenddate = strings.ReplaceAll(f.GetCell("AB", pos), ".", "-")
		customerItem.Type = customer.TypeOutsourcing

		str := f.GetCell("AN", pos)

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

		if f.GetCell("AM", pos) == "지로" {
			customerItem.Billingtype = 1
		} else {
			customerItem.Billingtype = 2
		}

		customerItem.Building = buildingId
		customerItem.User = userId
		customerItem.Salesuser = salesuserId
		customerItem.Company = session.Company

		customerManager.DeleteByCompanyBuilding(session.Company, buildingId)
		customerManager.Insert(&customerItem)

		pos++
	}

	departmentManager := models.NewDepartmentManager(conn)
	licenseManager := models.NewLicenseManager(conn)
	licensecategoryManager := models.NewLicensecategoryManager(conn)

	sheet = "소속회원"
	f.SetSheet(sheet)

	pos = 1
	for {
		item := models.User{}

		name := f.GetCell("B", pos)
		licenseno := ""
		zip := ""
		address := f.GetCell("F", pos)
		tel := f.GetCell("E", pos)
		email := f.GetCell("D", pos)
		licensename := ""

		educationdate := ""
		educationinstitution := ""
		specialeducationdate := ""
		specialeducationinstitution := ""
		joindate := f.GetCell("J", pos)
		status := f.GetCell("H", pos)

		userItem := userManager.GetByCompanyName(session.Company, name)

		if userItem != nil {
			item = *userItem
		}

		departmentName := f.GetCell("A", pos)
		department := departmentManager.GetByCompanyName(session.Company, departmentName)
		if department == nil {
			department = &models.Department{
				Name:    departmentName,
				Status:  1,
				Company: session.Company,
			}

			departmentManager.Insert(department)

			department.Id = departmentManager.GetIdentity()
		}

		item.Department = department.Id
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
			item.Company = session.Company
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
			licensecategoryItem.Id = licensecategoryManager.GetIdentity()
		}

		licenseItem := licenseManager.GetByUserLicensecategory(item.Id, licensecategoryItem.Id)

		if licenseItem == nil {
			//licenseManager.DeleteByUser(item.Id)
			licenseManager.Insert(&models.License{User: item.Id, Number: licenseno, Licensecategory: licensecategoryItem.Id})
		}

		pos++
	}

	conn.Commit()
}
