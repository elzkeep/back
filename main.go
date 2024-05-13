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

		typeid := 1
		var total float64 = 0
		for _, v := range items {
			if v.Category == 10 {
				total += float64(global.Atoi(v.Value2))

				typeid = global.Atoi(v.Value3)
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

		var score float64 = 0.0

		if typeid != 2 {
			if total <= 50 {
				score = 0.7
			} else if total <= 100 {
				score = 0.8
			} else if total <= 200 {
				score = 0.9
			} else if total <= 300 {
				score = 1
			} else if total <= 400 {
				score = 1.5
			} else {
				score = 1.8
			}
		} else {
			if total <= 100 {
				score = 1
			} else if total <= 200 {
				score = 1.2
			} else if total <= 300 {
				score = 1.3
			} else if total <= 400 {
				score = 2
			} else if total <= 500 {
				score = 2.4
			} else if total <= 600 {
				score = 3
			} else if total <= 700 {
				score = 4
			} else if total <= 800 {
				score = 5
			} else if total <= 900 {
				score = 6
			} else if total <= 1000 {
				score = 7.5
			} else if total <= 1250 {
				score = 10
			} else if total <= 1500 {
				score = 12
			} else if total <= 2000 {
				score = 15
			} else {
				score = 20
			}
		}

		buildingManager.UpdateTotalweight(models.Double(score), id)

	}

	conn.Close()

	services.Cron()
	services.Chat()
	services.Notify()
	//services.Fcm()

	services.Http()

}
