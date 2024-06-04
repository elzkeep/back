package main

import (
	"runtime"
	"zkeep/config"
	"zkeep/controllers/api"
	"zkeep/models"
	"zkeep/services"

	log "github.com/sirupsen/logrus"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})

	log.Printf("Version=" + config.Version + " Mode=" + config.Mode)
	log.Info("Server Start")

	models.InitCache()

	//DeleteFacility()
	UpdateScore()

	services.Cron()
	services.Chat()
	services.Notify()
	//services.Fcm()

	services.Http()

}

func DeleteFacility() {
	conn := models.NewConnection()

	buildingManager := models.NewBuildingManager(conn)
	facilityManager := models.NewFacilityManager(conn)

	buildings := buildingManager.Find(nil)

	for _, v := range buildings {
		items := facilityManager.Find([]interface{}{
			models.Where{Column: "building", Value: v.Id, Compare: "="},
			models.Where{Column: "category", Value: 10, Compare: "="},
		})

		if len(items) <= 1 {
			continue
		}

		for i, v2 := range items {
			if i == len(items)-1 {
				continue
			}

			facilityManager.Delete(v2.Id)
		}
	}
}

func UpdateScore() {
	db := models.NewConnection()
	defer db.Close()

	conn, _ := db.Begin()
	defer conn.Rollback()

	buildingManager := models.NewBuildingManager(conn)

	buildings := buildingManager.Find(nil)

	for _, b := range buildings {
		id := b.Id

		api.CalculateScore(conn, id)
	}

	conn.Commit()
}
