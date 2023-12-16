package main

import (
	"aoi/config"
	"aoi/models"
	"aoi/services"
	"math/rand"
	"runtime"
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().UnixNano())

	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})

	log.Printf("Version=" + config.Version + " Mode=" + config.Mode)
	log.Info("Server Start")

	models.InitCache()

	services.Cron()
	services.Chat()
	services.Notify()
	//services.Fcm()

	services.Http()

}
