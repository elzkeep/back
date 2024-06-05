package main_test

import (
	"fmt"
	"log"
	"path"
	"regexp"
	"strings"
	"sync"
	"testing"
	"zkeep/config"
	"zkeep/controllers/api"
	"zkeep/global"
	"zkeep/models"
	"zkeep/models/company"
	"zkeep/models/customer"
	"zkeep/models/user"
)

var cols map[string]int

func TestExcel(t *testing.T) {
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

	filename := "external.xlsx"
	typeid := 1
	myCompanyId := int64(7213)

	db := models.NewConnection()
	defer db.Close()

	conn, _ := db.Begin()
	defer conn.Rollback()

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

	f.Close()

	rows := len(cells)
	log.Println("ROWS", rows)
	max := 10

	wg := new(sync.WaitGroup)

	for i := 0; i < max; i++ {
		wg.Add(1)
		go func(start int) {
			conn := models.NewConnection()
			defer conn.Close()

			companyManager := models.NewCompanyManager(conn)
			customercompanyManager := models.NewCustomercompanyManager(conn)
			buildingManager := models.NewBuildingManager(conn)
			customerManager := models.NewCustomerManager(conn)
			userManager := models.NewUserManager(conn)

			log.Println("START POS", start)
			pos := start
			for {
				log.Println("POS", pos)

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

				no := cell[cols["A"]]

				if no == "" {
					break
				}

				userName := cell[cols["M"]]
				if userName == "" {
					userName = cell[cols["L"]]
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
						userId = userFind.Id
					}
				}

				salesuserName := cell[cols["N"]]
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

				item.Name = cell[cols["Y"]]
				item.Companyno = cell[cols["Z"]]
				item.Ceo = cell[cols["AA"]]
				item.Businesscondition = cell[cols["AB"]]
				item.Businessitem = cell[cols["AC"]]
				item.Address = cell[cols["AD"]]
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

				building.Name = cell[cols["C"]]
				building.Address = cell[cols["D"]]
				building.Contractvolumn = models.Double(global.Atol(cell[cols["E"]]))
				building.Receivevolumn = models.Double(global.Atol(cell[cols["F"]]))
				building.Generatevolumn = models.Double(global.Atol(cell[cols["G"]]))
				building.Sunlightvolumn = models.Double(global.Atol(cell[cols["H"]]))
				building.Ceo = cell[cols["AN"]]

				weight := global.Atof(cell[cols["E"]])
				building.Weight = models.Double(weight)
				volttype := cell[cols["I"]]

				if volttype == "고압" || volttype == "특고압" || volttype == "특 고압" {
					building.Volttype = 2
				} else {
					building.Volttype = 1
				}

				building.Checkcount = global.Atoi(cell[cols["K"]])

				building.Receivevolt = global.Atoi(strings.ReplaceAll(cell[cols["O"]], "V", ""))
				building.Usage = cell[cols["T"]]
				building.District = cell[cols["U"]]
				building.Company = companyId

				basic := global.Atoi(cell[cols["F"]])
				generator := global.Atoi(cell[cols["G"]])
				sunlight := global.Atoi(cell[cols["H"]])

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

				building.Postaddress = cell[cols["AM"]]
				building.Postname = cell[cols["AN"]]
				building.Posttel = cell[cols["AO"]]

				var buildingId int64 = 0

				buildingFind := buildingManager.GetByCompanyName(companyId, building.Name)
				if buildingFind == nil {
					buildingManager.Insert(&building)
					buildingId = buildingManager.GetIdentity()
				} else {
					buildingId = buildingFind.Id
				}

				api.CalculateScore2(conn, buildingId)

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

				customerItem.Number = global.Atoi(cell[cols["B"]])
				customerItem.Managername = cell[cols["V"]]
				customerItem.Managertel = cell[cols["W"]]
				customerItem.Manageremail = cell[cols["X"]]
				customerItem.Address = cell[cols["AM"]]
				customerItem.Manager = cell[cols["AN"]]
				customerItem.Contractprice = global.Atoi(cell[cols["AQ"]])
				customerItem.Contractvat = global.Atoi(cell[cols["AR"]])
				customerItem.Status = typeid
				customerItem.Contractstartdate = strings.ReplaceAll(cell[cols["AE"]], ".", "-")
				customerItem.Contractenddate = strings.ReplaceAll(cell[cols["AF"]], ".", "-")
				customerItem.Remark = cell[cols["AJ"]]
				customerItem.Type = customer.TypeOutsourcing

				if typeid == 1 {
					if cell[cols["AH"]] != "" {
						customerItem.Status = 2
					}
				}

				r, _ := regexp.Compile("[0-9]+")

				billStr := cell[cols["AS"]]
				billdate := r.FindString(billStr)

				customerItem.Billingdate = global.Atoi(billdate)

				str := cell[cols["AU"]]

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

				if cell[cols["AT"]] == "지로" {
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

				pos += max
			}

			wg.Done()
		}(i)
	}

	wg.Wait()
}
