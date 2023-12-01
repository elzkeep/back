package main

import (
	"aoi/config"
	"aoi/game"
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

	/*
		conn := models.NewConnection()
		defer conn.Close()

		gameManager := models.NewGameManager(conn)


			gameuserManager := models.NewGameuserManager(conn)
			gametileManager := models.NewGametileManager(conn)
			gamehistoryManager := models.NewGamehistoryManager(conn)
			gameundoManager := models.NewGameundoManager(conn)
			gameundoitemManager := models.NewGameundoitemManager(conn)

			for i := 1; i <= 140; i++ {
				var id int64
				id = int64(i)
				game := gameManager.Get(id)

				if game != nil {
					continue
				}

				gameuserManager.DeleteByGame(id)
				gametileManager.DeleteByGame(id)
				gamehistoryManager.DeleteByGame(id)
				gameundoManager.DeleteByGame(id)
				gameundoitemManager.DeleteByGame(id)
			}
	*/

	game.Init()

	/*
		userManager := models.NewUserManager(conn)
		users := userManager.Find(nil)
		for _, v := range users {
			v.Count = 0
			v.Elo = models.Double(1000.0)
			userManager.Update(&v)
		}

		items := gameManager.Find(nil)
		for _, v := range items {
			game.MakeGame(v.Id)
		}
	*/

	services.Cron()
	services.Chat()
	services.Notify()
	//services.Fcm()

	services.Http()

}
