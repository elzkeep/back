package api

import (
	"log"
	"zkeep/controllers"
	"zkeep/models"
)

type DataitemController struct {
	controllers.Controller
}

// @Post()
func (c *DataitemController) Process(datas []models.Dataitem) {
	log.Println("dataitem/process")

	session := c.Session

	if session == nil {
		c.Error("invalid request")
		return
	}

	if len(datas) == 0 {
		return
	}

	db := c.NewConnection()

	// 다른 회사의 점검 보고서 데이터는 수정할 수 없음
	reportManager := models.NewReportManager(db)
	checked := make(map[int64]bool)
	for _, v := range datas {
		if checked[v.Data.Report] {
			continue
		}

		report := reportManager.Get(v.Data.Report)
		if report == nil || (!c.IsRootadmin() && report.Company != session.Company) {
			c.Error("report not found")
			return
		}

		checked[v.Data.Report] = true
	}

	conn, err := db.Begin()
	if err != nil {
		log.Println(err)
		c.Error("database error")
		return
	}
	defer conn.Rollback()

	dataManager := models.NewDataManager(conn)
	itemManager := models.NewItemManager(conn)

	reportId := datas[0].Data.Report
	topcategory := datas[0].Data.Topcategory

	dataManager.DeleteByReportTopcategory(reportId, topcategory)
	itemManager.DeleteByReportTopcategory(reportId, topcategory)

	for _, v := range datas {
		// id·업체는 클라이언트 값을 쓰지 않는다 (위에서 소유를 확인한 점검의 업체로 고정)
		v.Data.Id = 0
		if report := reportManager.Get(v.Data.Report); report != nil {
			v.Data.Company = report.Company
		}
		err := dataManager.Insert(&v.Data)
		log.Println(err)
		id := dataManager.GetIdentity()
		for _, item := range v.Items {
			item.Id = 0
			item.Data = id
			item.Report = v.Data.Report
			item.Topcategory = v.Data.Topcategory
			itemManager.Insert(&item)
		}
	}

	conn.Commit()
}
