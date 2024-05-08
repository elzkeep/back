package main

import (
	"runtime"
	"zkeep/config"
	"zkeep/global"
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

	conn := models.NewConnection()

	buildingManager := models.NewBuildingManager(conn)
	facilityManager := models.NewFacilityManager(conn)

	buildings := buildingManager.Find(nil)

	for _, b := range buildings {
		id := b.Id
		items := facilityManager.Find([]interface{}{
			models.Where{Column: "building", Value: id, Compare: "="},
		})

		var total float64 = 0
		for _, v := range items {
			if v.Category == 10 {
				total += float64(global.Atoi(v.Value2))
			}

			if v.Category == 20 {
				total += float64(global.Atoi(v.Value12))
			}

			if v.Category == 30 {
				total += float64(global.Atoi(v.Value6))
			}

			if v.Category == 40 {
				total += float64(global.Atoi(v.Value5))
			}

			if v.Category == 50 {
				total += float64(global.Atoi(v.Value4))
			}

			if v.Category == 60 {
				total += float64(global.Atoi(v.Value12))
			}

			if v.Category == 70 {
				total += float64(global.Atoi(v.Value4))
			}

			if v.Category == 90 {
				total += float64(global.Atoi(v.Value6))
			}
		}

		buildingManager.UpdateTotalweight(models.Double(total), id)
	}

	conn.Close()

	services.Cron()
	services.Chat()
	services.Notify()
	//services.Fcm()

	services.Http()

}
