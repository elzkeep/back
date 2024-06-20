package api

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"regexp"
	"strings"
	"sync"
	"zkeep/config"
	"zkeep/controllers"
	"zkeep/global"
	"zkeep/models"
	"zkeep/models/billing"
	"zkeep/models/company"
	"zkeep/models/customer"
	"zkeep/models/user"
)

var cols map[string]int

func init() {
	cols = map[string]int{
		"A": 0,
		"B": 1,
		"C": 2,
		"D": 3,
		"E": 4,
		"F": 5,
		"G": 6,
		"H": 7,
		"I": 8,
		"J": 9,
		"K": 10,
		"L": 11,
		"M": 12,
		"N": 13,
		"O": 14,
		"P": 15,
		"Q": 16,
		"R": 17,
		"S": 18,
		"T": 19,
		"U": 20,
		"V": 21,
		"W": 22,
		"X": 23,
		"Y": 24,
		"Z": 25,

		"AA": 26,
		"AB": 27,
		"AC": 28,
		"AD": 39,
		"AE": 30,
		"AF": 31,
		"AG": 32,
		"AH": 33,
		"AI": 34,
		"AJ": 35,
		"AK": 36,
		"AL": 37,
		"AM": 38,
		"AN": 39,
		"AO": 40,
		"AP": 41,
		"AQ": 42,
		"AR": 43,
		"AS": 44,
		"AT": 45,
		"AU": 46,
		"AV": 47,
		"AW": 48,
		"AX": 49,
		"AY": 50,
		"AZ": 51,

		"BA": 52,
		"BB": 53,
		"BC": 54,
		"BD": 55,
		"BE": 56,
		"BF": 57,
		"BG": 58,
		"BH": 59,
		"BI": 60,
		"BJ": 61,
		"BK": 62,
		"BL": 63,
		"BM": 64,
		"BN": 65,
		"BO": 66,
		"BP": 67,
		"BQ": 68,
		"BR": 69,
		"BS": 70,
		"BT": 71,
		"BU": 72,
		"BV": 73,
		"BW": 74,
		"BX": 75,
		"BY": 76,
		"BZ": 77,
	}
}

type ExternalController struct {
	controllers.Controller
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
			numberStr := line[51:70]
			log.Println(numberStr)
			number := global.Atol(numberStr) - 1000000000
			log.Println(number)
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

func (c *ExternalController) Index(filenames string, typeid int) {
	user := c.Session

	conn := c.NewConnection()
	customerManager := models.NewCustomerManager(conn)
	customercompanyManager := models.NewCustomercompanyManager(conn)

	customerManager.DeleteByCompany(user.Company)
	customercompanyManager.DeleteByCompany(user.Company)

	files := strings.Split(filenames, ",")

	for _, v := range files {
		Thread(v, typeid, user.Company)
	}
}

func Thread(filename string, typeid int, myCompanyId int64) {
	fullFilename := path.Join(config.UploadPath, filename)
	f := global.NewExcelReader(fullFilename)
	if f == nil {
		log.Println("not found file")
		return
	}

	list := f.File.GetSheetList()

	sheet := list[0]
	f.SetSheet(sheet)

	cells := f.GetRows(sheet)

	log.Println("len", len(cells))
	f.Close()

	max := 10

	wg := new(sync.WaitGroup)

	for i := 0; i < max; i++ {
		wg.Add(1)
		go func(start int) {
			ExcelProcess(start, max, typeid, myCompanyId, cells)
			wg.Done()
		}(i)
	}

	wg.Wait()
}

func GetCell(str string, cells []string) string {
	col := cols[str]

	if col >= len(cells) {
		return ""
	}

	return cells[col]
}

func ExcelProcessCheck(start int, max int, typeid int, myCompanyId int64, cells [][]string, ch chan int) {
	rows := len(cells)

	log.Println("start", start)
	pos := start
	for {
		if pos < 4 {
			pos += max
			continue
		}

		if pos >= rows {
			break
		}

		cell := cells[pos]

		no := GetCell("A", cell)

		if no == "" {
			break
		}

		//log.Println("pos", pos)
		ch <- pos

		pos += max
	}
}

func ExcelProcess(start int, max int, typeid int, myCompanyId int64, cells [][]string) {
	rows := len(cells)

	re, _ := regexp.Compile("[0-9][0-9]-[0-9][0-9]-2[0-9]")

	conn := models.NewConnection()
	defer conn.Close()

	companyManager := models.NewCompanyManager(conn)
	customercompanyManager := models.NewCustomercompanyManager(conn)
	buildingManager := models.NewBuildingManager(conn)
	customerManager := models.NewCustomerManager(conn)
	userManager := models.NewUserManager(conn)
	facilityManager := models.NewFacilityManager(conn)

	pos := start
	for {
		if pos < 4 {
			pos += max
			continue
		}

		if pos >= rows {
			break
		}

		item := models.Company{}
		building := models.Building{}
		customerItem := models.Customer{}

		cell := cells[pos]

		no := GetCell("A", cell)

		if no == "" {
			break
		}

		userName := GetCell("M", cell)
		if userName == "" {
			userName = GetCell("L", cell)
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
				userItem.Status = user.StatusNotuse
				userItem.Approval = user.ApprovalComplete
				userItem.Score = 60

				userManager.Insert(&userItem)
				userId = userManager.GetIdentity()
			} else {
				if userFind.Level != user.LevelAdmin {
					userFind.Level = user.LevelNormal
					userFind.Company = myCompanyId
					userFind.Name = userName
					userFind.Status = user.StatusNotuse
					userFind.Approval = user.ApprovalComplete

					userManager.Update(userFind)
				}

				userId = userFind.Id
			}
		}

		salesuserName := GetCell("N", cell)
		var salesuserId int64 = 0

		if salesuserName != "" {
			userFind := userManager.GetByCompanyName(myCompanyId, salesuserName)

			if userFind == nil {
				userItem := models.User{}
				userItem.Level = user.LevelNormal
				userItem.Company = myCompanyId
				userItem.Name = salesuserName
				userItem.Loginid = userItem.Name
				userItem.Status = user.StatusNotuse
				userItem.Approval = user.ApprovalComplete
				userItem.Score = 60

				userManager.Insert(&userItem)
				salesuserId = userManager.GetIdentity()
			} else {
				if userFind.Level != user.LevelAdmin {
					userFind.Level = user.LevelNormal
					userFind.Company = myCompanyId
					userFind.Name = salesuserName
					userFind.Status = user.StatusNotuse
					userFind.Approval = user.ApprovalComplete

					userManager.Update(userFind)
				}

				salesuserId = userFind.Id
			}
		}

		item.Name = GetCell("Y", cell)
		item.Companyno = GetCell("Z", cell)
		item.Ceo = GetCell("AA", cell)
		item.Businesscondition = GetCell("AB", cell)
		item.Businessitem = GetCell("AC", cell)
		item.Address = GetCell("AD", cell)
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
		} else {
			//log.Println("already :", companyId, item.Name, item.Companyno)
		}

		building.Name = GetCell("C", cell)
		building.Address = GetCell("D", cell)
		building.Contractvolumn = models.Double(global.Atol(GetCell("E", cell)))
		building.Receivevolumn = models.Double(global.Atol(GetCell("F", cell)))
		building.Generatevolumn = models.Double(global.Atol(GetCell("G", cell)))
		building.Sunlightvolumn = models.Double(global.Atol(GetCell("H", cell)))
		building.Ceo = GetCell("AN", cell)

		weight := global.Atof(GetCell("E", cell))
		building.Weight = models.Double(weight)
		volttype := GetCell("I", cell)

		if volttype == "고압" || volttype == "특고압" || volttype == "특 고압" {
			building.Volttype = 2
		} else {
			building.Volttype = 1
		}

		building.Checkcount = global.Atoi(GetCell("K", cell))

		building.Receivevolt = global.Atoi(strings.ReplaceAll(GetCell("O", cell), "V", ""))
		building.Usage = GetCell("T", cell)
		building.District = GetCell("U", cell)
		building.Company = companyId

		basic := global.Atoi(GetCell("F", cell))
		generator := global.Atoi(GetCell("G", cell))

		sunlight := global.Atoi(GetCell("H", cell))

		basicFacility := models.Facility{}
		generatorFacility := models.Facility{}
		sunlightFacility := models.Facility{}

		if basic > 0 {
			basicFacility.Category = 10
			basicFacility.Value2 = fmt.Sprintf("%v", basic)
			basicFacility.Type = building.Volttype
		}

		if generator > 0 {
			generatorFacility.Category = 20
			generatorFacility.Value12 = fmt.Sprintf("%v", generator)
			generatorFacility.Type = building.Volttype
		}

		if sunlight > 0 {
			sunlightFacility.Category = 30
			sunlightFacility.Value6 = fmt.Sprintf("%v", sunlight)
			sunlightFacility.Type = building.Volttype
		}

		building.Totalweight = models.Double(basic + generator + sunlight)

		building.Postaddress = GetCell("AM", cell)
		building.Postname = GetCell("AN", cell)
		building.Posttel = GetCell("AO", cell)

		building.Cmsnumber = GetCell("AW", cell)
		building.Cmsbank = GetCell("AX", cell)
		building.Cmsaccountno = GetCell("AY", cell)
		building.Cmsconfirm = GetCell("AZ", cell)
		building.Cmsstartdate = GetCell("BA", cell)
		building.Cmsenddate = GetCell("BB", cell)

		var buildingId int64 = 0

		buildingFind := buildingManager.GetByCompanyName(companyId, building.Name)
		if buildingFind == nil {
			buildingManager.Insert(&building)
			buildingId = buildingManager.GetIdentity()
		} else {
			buildingId = buildingFind.Id
		}

		facilityManager.DeleteByBuilding(buildingId)
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

		CalculateScore2(conn, buildingId)

		customerItem.Number = global.Atoi(GetCell("B", cell))
		customerItem.Managername = GetCell("V", cell)
		customerItem.Managertel = GetCell("W", cell)
		customerItem.Manageremail = GetCell("X", cell)
		customerItem.Address = GetCell("AM", cell)
		customerItem.Manager = GetCell("AN", cell)
		customerItem.Contractprice = global.Atoi(GetCell("AQ", cell))
		customerItem.Contractvat = global.Atoi(GetCell("AR", cell))
		customerItem.Contracttype = 1
		customerItem.Status = typeid

		contractstartdate := GetCell("AE", cell)
		if re.MatchString(contractstartdate) {
			customerItem.Contractstartdate = fmt.Sprintf("20%v-%v", contractstartdate[6:], contractstartdate[:5])
		} else {
			customerItem.Contractstartdate = strings.ReplaceAll(contractstartdate, ".", "-")
		}

		contractenddate := GetCell("AF", cell)
		if re.MatchString(contractenddate) {
			customerItem.Contractenddate = fmt.Sprintf("20%v-%v", contractenddate[6:], contractenddate[:5])
		} else {
			customerItem.Contractenddate = strings.ReplaceAll(contractenddate, ".", "-")
		}

		customerItem.Remark = GetCell("AJ", cell)
		customerItem.Type = customer.TypeOutsourcing

		if typeid == 1 {
			if GetCell("AH", cell) != "" {
				customerItem.Status = 2
			}
		}

		r, _ := regexp.Compile("[0-9]+")

		billStr := GetCell("AS", cell)
		billdate := r.FindString(billStr)

		customerItem.Billingdate = global.Atoi(billdate)

		str := GetCell("AU", cell)

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

		if GetCell("AT", cell) == "지로" {
			customerItem.Billingtype = 1
		} else {
			customerItem.Billingtype = 2
		}

		customerItem.Building = buildingId
		customerItem.User = userId
		customerItem.Salesuser = salesuserId
		customerItem.Company = myCompanyId

		customerManager.DeleteByCompanyBuilding(myCompanyId, buildingId)
		err := customerManager.Insert(&customerItem)

		if err != nil {
			log.Println(err)
		}

		pos += max
	}
}

func ExcelProcessOld(filename string, typeid int, myCompanyId int64) {
	db := models.NewConnection()
	defer db.Close()

	re, _ := regexp.Compile("[0-9][0-9]-[0-9][0-9]-2[0-9]")

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

	values := f.GetRows(sheet)
	log.Println(values)

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
				if typeid == 1 {
					userItem.Status = user.StatusUse
				} else {
					userItem.Status = user.StatusNotuse
				}
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
				if typeid == 1 {
					userItem.Status = user.StatusUse
				} else {
					userItem.Status = user.StatusNotuse
				}
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

		if volttype == "고압" || volttype == "특고압" || volttype == "특 고압" {
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

		CalculateScore(conn, buildingId)

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

		contractstartdate := f.GetCell("AE", pos)
		if re.MatchString(contractstartdate) {
			customerItem.Contractstartdate = fmt.Sprintf("20%v-%v", contractstartdate[6:], contractstartdate[:5])
		} else {
			customerItem.Contractstartdate = strings.ReplaceAll(contractstartdate, ".", "-")
		}

		contractenddate := f.GetCell("AF", pos)
		if re.MatchString(contractenddate) {
			customerItem.Contractenddate = fmt.Sprintf("20%v-%v", contractenddate[6:], contractenddate[:5])
		} else {
			customerItem.Contractenddate = strings.ReplaceAll(contractenddate, ".", "-")
		}

		customerItem.Remark = f.GetCell("AJ", pos)
		customerItem.Type = customer.TypeOutsourcing

		if typeid == 1 {
			if f.GetCell("AH", pos) != "" {
				customerItem.Status = 2
			}
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

	re, _ := regexp.Compile("[0-9][0-9]-[0-9][0-9]-2[0-9]")

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
		if re.MatchString(joindate) {
			joindate = fmt.Sprintf("20%v-%v", joindate[6:], joindate[:5])
		}

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
		item.Status = user.StatusUse

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
			if userItem.Level != user.LevelAdmin {
				userManager.Update(&item)
			}
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

	fullFilename := path.Join(config.UploadPath, filename)
	f := global.NewExcelReader(fullFilename)
	if f == nil {
		log.Println("not found file")
		return
	}

	userCells := make([][]string, 0)
	if category != 1 {
		sheet := "소속회원"
		f.SetSheet(sheet)
		userCells = f.GetRows(sheet)
	}

	companyCells := make([][]string, 0)
	if category != 2 {
		sheet := "고객 현황"
		f.SetSheet(sheet)
		companyCells = f.GetRows(sheet)
	}

	f.Close()

	max := 10

	wg := new(sync.WaitGroup)

	for i := 0; i < max; i++ {
		wg.Add(1)
		AllProcess(i, max, category, session.Company, userCells, companyCells)
		wg.Done()
	}

	wg.Wait()
}

func AllProcess(start int, max int, category int, myCompanyId int64, userCells [][]string, companyCells [][]string) {
	conn := models.NewConnection()
	defer conn.Close()

	re, _ := regexp.Compile("[0-9][0-9]-[0-9][0-9]-2[0-9]")

	companyManager := models.NewCompanyManager(conn)
	customercompanyManager := models.NewCustomercompanyManager(conn)
	buildingManager := models.NewBuildingManager(conn)
	customerManager := models.NewCustomerManager(conn)
	userManager := models.NewUserManager(conn)

	if category != 1 {
		departmentManager := models.NewDepartmentManager(conn)
		licenseManager := models.NewLicenseManager(conn)
		licensecategoryManager := models.NewLicensecategoryManager(conn)
		licenselevelManager := models.NewLicenselevelManager(conn)

		rows := len(userCells)

		pos := start
		for {
			if pos < 1 {
				pos += max
				continue
			}

			if pos >= rows {
				break
			}

			cell := userCells[pos]

			item := models.User{}

			loginid := GetCell("B", cell)

			if loginid == "" {
				break
			}

			if loginid == "로그인아이디" {
				pos++
				continue
			}

			zip := ""
			address := GetCell("F", cell)
			tel := GetCell("E", cell)
			email := GetCell("D", cell)
			name := GetCell("C", cell)

			educationdate := ""
			educationinstitution := ""
			specialeducationdate := ""
			specialeducationinstitution := ""

			joindate := GetCell("J", cell)
			if re.MatchString(joindate) {
				joindate = fmt.Sprintf("20%v-%v", joindate[6:], joindate[:5])
			}

			status := GetCell("H", cell)

			userItem := userManager.GetByCompanyName(myCompanyId, name)

			if userItem != nil {
				item = *userItem
			}

			departmentName := GetCell("A", cell)
			department := departmentManager.GetByCompanyName(myCompanyId, departmentName)
			if department == nil {
				department = &models.Department{
					Name:    departmentName,
					Status:  1,
					Company: myCompanyId,
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
			item.Status = user.StatusUse

			if status == "퇴직" || status == "사용 안함" || status == "사용안함" || status == "미사용" {
				item.Status = user.StatusNotuse
			} else {
				item.Status = user.StatusUse
			}

			if userItem == nil {
				item.Company = myCompanyId
				item.Loginid = loginid
				item.Passwd = "0000"
				item.Score = 60

				userManager.Insert(&item)
				item.Id = userManager.GetIdentity()
			} else {
				if userItem.Level != user.LevelAdmin {
					userManager.Update(&item)
				}
			}

			licensename := GetCell("K", cell)
			licenseno := GetCell("L", cell)
			licenselevel := GetCell("M", cell)
			licensedate := GetCell("N", cell)

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

			pos += max
		}
	}

	if category != 2 {
		rows := len(companyCells)

		pos := start
		for {
			if pos < 2 {
				pos += max
				continue
			}

			if pos >= rows {
				break
			}

			cell := companyCells[pos]

			item := models.Company{}
			customerItem := models.Customer{}

			no := GetCell("A", cell)

			if no == "" {
				break
			}

			item.Name = GetCell("B", cell)
			item.Companyno = GetCell("C", cell)
			item.Ceo = GetCell("D", cell)
			item.Address = GetCell("E", cell)
			item.Addressetc = GetCell("F", cell)
			item.Tel = GetCell("G", cell)
			item.Email = GetCell("H", cell)
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

			buildingName := GetCell("I", cell)
			building := buildingManager.GetByCompanyName(companyId, buildingName)
			if building == nil {
				building = &models.Building{}
			}

			building.Name = GetCell("I", cell)
			building.Companyno = GetCell("J", cell)
			building.Ceo = GetCell("K", cell)
			building.Zip = GetCell("L", cell)

			building.Address = GetCell("M", cell)
			building.Addressetc = GetCell("N", cell)

			building.Businesscondition = GetCell("O", cell)
			building.Businessitem = GetCell("P", cell)
			building.Usage = GetCell("Q", cell)

			contracttype := GetCell("R", cell)
			if contracttype == "안전관리" {
				customerItem.Contracttype = 1
			} else if contracttype == "유지보수" {
				customerItem.Contracttype = 2
			} else if contracttype == "안전관리+유지보수" {
				customerItem.Contracttype = 3
			} else {
				customerItem.Contracttype = 1
			}

			weight := global.Atof(GetCell("S", cell))
			building.Weight = models.Double(weight)

			userName := GetCell("T", cell)

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
					if userFind.Level != user.LevelAdmin {
						userId = userFind.Id
					}
				}
			}

			salesuserName := GetCell("U", cell)
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

			contractstartdate := GetCell("V", cell)

			if re.MatchString(contractstartdate) {
				customerItem.Contractstartdate = fmt.Sprintf("20%v-%v", contractstartdate[6:], contractstartdate[:5])
			} else {
				customerItem.Contractstartdate = strings.ReplaceAll(contractstartdate, ".", "-")
			}

			contractenddate := GetCell("W", cell)
			if re.MatchString(contractenddate) {
				customerItem.Contractenddate = fmt.Sprintf("20%v-%v", contractenddate[6:], contractenddate[:5])
			} else {
				customerItem.Contractenddate = strings.ReplaceAll(contractenddate, ".", "-")
			}

			building.District = GetCell("X", cell)

			customerItem.Kepconumber = GetCell("Y", cell)
			customerItem.Kesconumber = GetCell("Z", cell)

			customerItem.Periodic = GetCell("AA", cell)

			lastdate := GetCell("AB", cell)

			if re.MatchString(lastdate) {
				customerItem.Lastdate = fmt.Sprintf("20%v-%v", lastdate[6:], lastdate[:5])
			} else {
				customerItem.Lastdate = strings.ReplaceAll(lastdate, ".", "-")
			}

			building.Company = companyId

			building.Postzip = GetCell("AI", cell)
			building.Postaddress = GetCell("AJ", cell)
			building.Postaddress = GetCell("AK", cell)
			building.Postname = GetCell("AL", cell)
			building.Posttel = GetCell("AM", cell)

			var buildingId int64 = 0

			if building.Id == 0 {
				buildingManager.Insert(building)
				buildingId = buildingManager.GetIdentity()
			} else {
				buildingId = building.Id
			}

			CalculateScore2(conn, buildingId)

			customerItem.Number = global.Atoi(GetCell("A", cell))

			customerItem.Managername = GetCell("AC", cell)
			customerItem.Managertel = GetCell("AD", cell)
			customerItem.Manageremail = GetCell("AE", cell)

			customerItem.Billingname = GetCell("AF", cell)
			customerItem.Billingtel = GetCell("AG", cell)
			customerItem.Billingemail = GetCell("AH", cell)

			customerItem.Fax = GetCell("AN", cell)
			customerItem.Status = 1

			customerItem.Contractprice = global.Atoi(GetCell("AO", cell))
			customerItem.Contractvat = global.Atoi(GetCell("AP", cell))

			customerItem.Type = customer.TypeOutsourcing

			customerItem.Billingdate = global.Atoi(strings.TrimSpace(strings.ReplaceAll(GetCell("AQ", cell), "일", "")))

			billingtype := GetCell("AR", cell)
			if billingtype == "지로" {
				customerItem.Billingtype = 1
			} else if billingtype == "계산서" || billingtype == "세금계산서" || billingtype == "세금 계산서" || billingtype == "이체" || billingtype == "계좌이체" || billingtype == "계좌 이체" {
				customerItem.Billingtype = 2
			} else if billingtype == "카드" {
				customerItem.Billingtype = 3
			} else if billingtype == "CMS" || billingtype == "자동이체" || billingtype == "자동 이체" {
				customerItem.Billingtype = 4
			} else if billingtype == "소매 매출" || billingtype == "소매매출" {
				customerItem.Billingtype = 5
			} else {
				customerItem.Billingtype = 2
			}

			str := GetCell("AS", cell)

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

			customerItem.Remark = GetCell("AT", cell)

			customerItem.Building = buildingId
			customerItem.User = userId
			customerItem.Salesuser = salesuserId
			customerItem.Company = myCompanyId

			customerManager.DeleteByCompanyBuilding(myCompanyId, buildingId)
			customerManager.Insert(&customerItem)

			pos += max
		}
	}
}
