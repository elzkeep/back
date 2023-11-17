package main

import (
	"aoi/config"
	"aoi/game"
	"aoi/models"
	"aoi/services"
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})

	log.Printf("Version=" + config.Version + " Mode=" + config.Mode)
	log.Info("Server Start")

	models.InitCache()

	game.Init()

	services.Cron()
	services.Chat()
	services.Notify()
	//services.Fcm()

	services.Http()
}
