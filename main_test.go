package main_test

import (
	"fmt"
	"log"
	"path"
	"regexp"
	"strings"
	"testing"
	"zkeep/config"
	"zkeep/global"
	"zkeep/models"
	"zkeep/models/user"
)

func GetCell(col string, row int) string {
	return fmt.Sprintf("%v%v", col, row)
}

func TestExcel(t *testing.T) {
	var companyId int64 = 1

	db := models.NewConnection()
	defer db.Close()

	conn, _ := db.Begin()
	defer conn.Rollback()

	companyManager := models.NewCompanyManager(conn)
	buildingManager := models.NewBuildingManager(conn)
	customerManager := models.NewCustomerManager(conn)
	userManager := models.NewUserManager(conn)

	filename := "1.xlsx"
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
		customer := models.Customer{}

		no := f.GetCell("A", pos)

		if no == "" {
			break
		}

		userName := f.GetCell("M", pos)

		userFind := userManager.GetByCompanyName(companyId, userName)

		var userId int64 = 0
		if userFind == nil {
			userItem := models.User{}
			userItem.Level = user.LevelNormal
			userItem.Company = companyId
			userItem.Loginid = item.Name
			userItem.Passwd = "0000"

			userManager.Insert(&userItem)
			userId = userManager.GetIdentity()
		} else {
			userId = userFind.Id
		}

		item.Name = f.GetCell("Y", pos)
		item.Companyno = f.GetCell("Z", pos)
		item.Ceo = f.GetCell("AA", pos)
		item.Businesscondition = f.GetCell("AB", pos)
		item.Businessitem = f.GetCell("AC", pos)
		item.Address = f.GetCell("AD", pos)

		var companyId int64 = 0

		companyFind := companyManager.GetByCompanyno(item.Companyno)

		if companyFind == nil {
			companyManager.Insert(&item)
			companyId = companyManager.GetIdentity()
		} else {
			companyId = companyFind.Id
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

		customer.Managername = f.GetCell("V", pos)
		customer.Managertel = f.GetCell("W", pos)
		customer.Manageremail = f.GetCell("X", pos)
		customer.Address = f.GetCell("AM", pos)
		customer.Manager = f.GetCell("AN", pos)
		customer.Contractprice = global.Atoi(f.GetCell("AQ", pos))
		customer.Contractvat = global.Atoi(f.GetCell("AR", pos))
		customer.Status = 1

		str := f.GetCell("AU", pos)

		r, _ := regexp.Compile("[0-9]+")
		collectday := r.FindString(str)

		if strings.Contains(str, "당월") {
			customer.Collectmonth = 1
		} else {
			customer.Collectmonth = 2
		}

		customer.Collectday = global.Atoi(collectday)

		if f.GetCell("AT", pos) == "지로" {
			customer.Billingtype = 1
		} else {
			customer.Billingtype = 2
		}

		customer.Building = buildingId
		customer.User = userId
		customer.Company = companyId

		customerManager.DeleteByCompanyBuilding(companyId, buildingId)
		customerManager.Insert(&customer)

		pos++
	}

	conn.Commit()
}
