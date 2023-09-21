package main

import (
	"aoi/config"
	"aoi/game/resources"
	"aoi/models"
	"aoi/services"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Println(int(resources.D))
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})

	log.Printf("Version=" + config.Version + " Mode=" + config.Mode)
	log.Info("Server Start")

	models.InitCache()

	services.Cron()
	//services.Fcm()

	services.Http()
}
