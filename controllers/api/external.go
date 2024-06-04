package api

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"regexp"
	"strings"
	"zkeep/config"
	"zkeep/controllers"
	"zkeep/global"
	"zkeep/models"
	"zkeep/models/billing"
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
	//facilityManager := models.NewFacilityManager(conn)

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

		if volttype == "고압" || volttype == "특고압" {
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

		building.Postaddress = f.GetCell("AM", pos)
		building.Postname = f.GetCell("AN", pos)
		building.Posttel = f.GetCell("AO", pos)

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

			//facilityManager.Insert(&basicFacility)
		}

		if generator > 0 {
			generatorFacility.Building = buildingId

			//facilityManager.Insert(&generatorFacility)
		}

		if sunlight > 0 {
			sunlightFacility.Building = buildingId

			//facilityManager.Insert(&sunlightFacility)
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
		customerItem.Remark = f.GetCell("AJ", pos)
		customerItem.Type = customer.TypeOutsourcing

		if f.GetCell("AH", pos) != "" {
			customerItem.Status = 2
		}

		r, _ := regexp.Compile("[0-9]+")

		billStr := f.GetCell("AS", pos)
		billdate := r.FindString(billStr)

		customerItem.Billingdate = global.Atoi(billdate)

		str := f.GetCell("AU", pos)

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

		licensenames := strings.Split(licensename, "\n")
		licensenos := strings.Split(licenseno, "\n")

		for i, v := range licensenames {
			licensename := v
			licenseno := ""

			if len(licensenos)-1 >= i {
				licenseno = licensenos[i]
			}

			licensecategoryItem := licensecategoryManager.GetByName(licensename)
			if licensecategoryItem == nil {
				licensecategoryItem = &models.Licensecategory{Name: licensename, Order: 0}
				licensecategoryManager.Insert(licensecategoryItem)
				licensecategoryItem.Id = licensecategoryManager.GetIdentity()
			}

			licenseItem := licenseManager.GetByUserLicensecategory(item.Id, licensecategoryItem.Id)

			if licenseItem == nil {
				//licenseManager.DeleteByUser(item.Id)
				licenseManager.Insert(&models.License{User: item.Id, Number: licenseno, Licensecategory: licensecategoryItem.Id})
			}
		}

		pos++
	}

	conn.Commit()
}

func (c *ExternalController) All(category int, filename string) {
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

	fullFilename := path.Join(config.UploadPath, filename)
	f := global.NewExcelReader(fullFilename)
	if f == nil {
		log.Println("not found file")
		return
	}

	if category != 1 {
		departmentManager := models.NewDepartmentManager(conn)
		licenseManager := models.NewLicenseManager(conn)
		licensecategoryManager := models.NewLicensecategoryManager(conn)
		licenselevelManager := models.NewLicenselevelManager(conn)

		sheet := "소속회원"
		f.SetSheet(sheet)

		pos := 1
		for {
			item := models.User{}

			name := f.GetCell("B", pos)

			if name == "" {
				break
			}

			if name == "로그인아이디" {
				pos++
				continue
			}

			zip := ""
			address := f.GetCell("F", pos)
			tel := f.GetCell("E", pos)
			email := f.GetCell("D", pos)

			educationdate := ""
			educationinstitution := ""
			specialeducationdate := ""
			specialeducationinstitution := ""
			joindate := f.GetCell("J", pos)
			status := f.GetCell("H", pos)

			userItem := userManager.GetByCompanyTelName(session.Company, tel, name)

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

			licensename := f.GetCell("K", pos)
			licenseno := f.GetCell("L", pos)
			licenselevel := f.GetCell("M", pos)
			licensedate := f.GetCell("N", pos)

			if licensename != "" {
				licensenames := strings.Split(licensename, "\n")
				licensenos := strings.Split(licenseno, "\n")
				licenselevels := strings.Split(licenselevel, "\n")
				licensedates := strings.Split(licensedate, "\n")

				for i, v := range licensenames {
					licensename := v
					licenseno := ""
					licenselevel := ""
					licensedate := ""

					if len(licensenos)-1 >= i {
						licenseno = licensenos[i]
					}

					if len(licenselevels)-1 >= i {
						licenselevel = licenselevels[i]
					}

					if len(licensedates)-1 >= i {
						licensedate = licensedates[i]
					}

					licensecategoryItem := licensecategoryManager.GetByName(licensename)
					if licensecategoryItem == nil {
						licensecategoryItem = &models.Licensecategory{Name: licensename, Order: 0}
						licensecategoryManager.Insert(licensecategoryItem)
						licensecategoryItem.Id = licensecategoryManager.GetIdentity()
					}

					licenselevelItem := licenselevelManager.GetByName(licenselevel)
					if licenselevelItem == nil {
						licenselevelItem = &models.Licenselevel{Name: licenselevel, Order: 0}
						licenselevelManager.Insert(licenselevelItem)
						licenselevelItem.Id = licenselevelManager.GetIdentity()
					}

					licenseItem := licenseManager.GetByUserLicensecategory(item.Id, licensecategoryItem.Id)

					if licenseItem == nil {
						licenseManager.Insert(&models.License{User: item.Id, Number: licenseno, Licensecategory: licensecategoryItem.Id, Licenselevel: licenselevelItem.Id, Takingdate: licensedate})
					}
				}
			}

			pos++
		}
	}

	if category != 2 {
		sheet := "고객 현황"
		f.SetSheet(sheet)

		pos := 2
		for {
			log.Println("POS:", pos)
			item := models.Company{}
			customerItem := models.Customer{}

			no := f.GetCell("A", pos)

			if no == "" {
				break
			}

			item.Name = f.GetCell("B", pos)
			item.Companyno = f.GetCell("C", pos)
			item.Ceo = f.GetCell("D", pos)
			item.Address = f.GetCell("E", pos)
			item.Addressetc = f.GetCell("F", pos)
			item.Tel = f.GetCell("G", pos)
			item.Email = f.GetCell("H", pos)
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

			buildingName := f.GetCell("I", pos)
			building := buildingManager.GetByCompanyName(companyId, buildingName)
			if building == nil {
				building = &models.Building{}
			}

			building.Name = f.GetCell("I", pos)
			building.Companyno = f.GetCell("J", pos)
			building.Ceo = f.GetCell("K", pos)
			building.Zip = f.GetCell("L", pos)

			building.Address = f.GetCell("M", pos)
			building.Addressetc = f.GetCell("N", pos)

			building.Businesscondition = f.GetCell("O", pos)
			building.Businessitem = f.GetCell("P", pos)
			building.Usage = f.GetCell("Q", pos)

			contracttype := f.GetCell("R", pos)
			if contracttype == "안전관리" {
				customerItem.Contracttype = 1
			} else if contracttype == "유지보수" {
				customerItem.Contracttype = 2
			} else if contracttype == "안전관리+유지보수" {
				customerItem.Contracttype = 3
			} else {
				customerItem.Contracttype = 1
			}

			weight := global.Atof(f.GetCell("S", pos))
			building.Weight = models.Double(weight)

			userName := f.GetCell("T", pos)

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

			salesuserName := f.GetCell("U", pos)
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

			customerItem.Contractstartdate = strings.ReplaceAll(f.GetCell("V", pos), ".", "-")
			customerItem.Contractenddate = strings.ReplaceAll(f.GetCell("W", pos), ".", "-")

			building.District = f.GetCell("X", pos)

			customerItem.Kepconumber = f.GetCell("Y", pos)
			customerItem.Kesconumber = f.GetCell("Z", pos)

			customerItem.Periodic = f.GetCell("AA", pos)

			customerItem.Lastdate = strings.ReplaceAll(f.GetCell("AB", pos), ".", "-")

			building.Company = companyId

			building.Postzip = f.GetCell("AI", pos)
			building.Postaddress = f.GetCell("AJ", pos)
			building.Postname = f.GetCell("AK", pos)
			building.Posttel = f.GetCell("AL", pos)

			var buildingId int64 = 0

			if building.Id == 0 {
				buildingManager.Insert(building)
				buildingId = buildingManager.GetIdentity()
			} else {
				buildingId = building.Id
			}

			customerItem.Number = global.Atoi(f.GetCell("A", pos))

			customerItem.Managername = f.GetCell("AC", pos)
			customerItem.Managertel = f.GetCell("AD", pos)
			customerItem.Manageremail = f.GetCell("AE", pos)

			customerItem.Billingname = f.GetCell("AF", pos)
			customerItem.Billingtel = f.GetCell("AG", pos)
			customerItem.Billingemail = f.GetCell("AH", pos)

			customerItem.Fax = f.GetCell("AM", pos)
			customerItem.Status = 1

			customerItem.Contractprice = global.Atoi(f.GetCell("AK", pos))
			customerItem.Contractvat = global.Atoi(f.GetCell("AL", pos))

			customerItem.Type = customer.TypeOutsourcing

			customerItem.Billingdate = global.Atoi(strings.TrimSpace(strings.ReplaceAll(f.GetCell("AM", pos), "일", "")))

			if f.GetCell("AN", pos) == "지로" {
				customerItem.Billingtype = 1
			} else {
				customerItem.Billingtype = 2
			}

			str := f.GetCell("AO", pos)

			r, _ := regexp.Compile("[0-9]+")
			collectday := r.FindString(str)

			if strings.Contains(str, "당월") || strings.Contains(str, "매월") {
				customerItem.Collectmonth = 1
			} else {
				customerItem.Collectmonth = 2
			}

			if collectday == "" {
				customerItem.Collectday = 0
			} else {
				customerItem.Collectday = global.Atoi(collectday)
			}

			customerItem.Remark = f.GetCell("AP", pos)

			customerItem.Building = buildingId
			customerItem.User = userId
			customerItem.Salesuser = salesuserId
			customerItem.Company = session.Company

			customerManager.DeleteByCompanyBuilding(session.Company, buildingId)
			customerManager.Insert(&customerItem)

			pos++
		}
	}

	conn.Commit()
}

// @POST()
func (c *ExternalController) Giro(filename []string) {
	session := c.Session

	conn := c.NewConnection()

	billingManager := models.NewBillingManager(conn)
	giroManager := models.NewGiroManager(conn)

	for _, v := range filename {
		fullFilename := path.Join(config.UploadPath, v)
		data, err := os.Open(fullFilename)
		if err != nil {
			continue
		}
		scanner := bufio.NewScanner(data)
		scanner.Split(bufio.ScanLines)
		var txtlines []string

		for scanner.Scan() {
			txtlines = append(txtlines, scanner.Text())
		}

		data.Close()

		for i := 1; i < len(txtlines)-1; i++ {
			line := txtlines[i]
			log.Println(i, line)
			//220000001202401252024012900429030042903025002000343000001242024010151620000000069300 0260
			acceptdate := line[9:17]
			insertdate := line[17:25]
			numberStr := line[51:71]
			number := global.Atol(numberStr) - 1000000000
			price := global.Atoi(line[71:84])
			typeid := line[84:85]
			charge := global.Atoi(line[85:])
			log.Println(acceptdate, insertdate, number, price, typeid, charge)

			item := models.Giro{
				Acceptdate: acceptdate,
				Insertdate: insertdate,
				Number:     numberStr,
				Price:      price,
				Type:       typeid,
				Charge:     charge,
				Content:    line,
			}
			giroManager.Insert(&item)

			billingItem := billingManager.Get(number)

			if billingItem != nil && billingItem.Id == number {
				if billingItem.Company != session.Company {
					continue
				}

				billingItem.Status = billing.StatusComplete
				billingManager.Update(billingItem)
			}
		}
	}
}
