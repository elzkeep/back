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

	db := c.NewConnection()

	conn, _ := db.Begin()
	defer conn.Rollback()

	dataManager := models.NewDataManager(conn)
	itemManager := models.NewItemManager(conn)

	if len(datas) == 0 {
		return
	}

	reportId := datas[0].Data.Report
	topcategory := datas[0].Data.Topcategory

	dataManager.DeleteByReportTopcategory(reportId, topcategory)
	itemManager.DeleteByReportTopcategory(reportId, topcategory)

	for _, v := range datas {
		log.Println(v.Data)
		log.Println("id", v.Data.Id)
		log.Println("title", v.Data.Title)
		err := dataManager.Insert(&v.Data)
		log.Println(err)
		id := dataManager.GetIdentity()
		for _, item := range v.Items {
			item.Data = id
			item.Report = v.Data.Report
			item.Topcategory = v.Data.Topcategory
			itemManager.Insert(&item)
		}
	}

	conn.Commit()
}
