package main

import (
	"aoi/config"
	"aoi/game"
	"aoi/game/action"
	"aoi/game/resources"
	"aoi/models"
	"aoi/services"
	"encoding/gob"

	log "github.com/sirupsen/logrus"
)

func main() {
	gob.Register(&game.Game{})
	gob.Register(&game.Map{})
	gob.Register(&game.Science{})
	//gob.Register(&factions.FactionInterface{})
	gob.Register(&action.PowerAction{})
	gob.Register(&action.BookAction{})
	gob.Register(&resources.RoundTile{})
	gob.Register(&game.RoundBonus{})
	gob.Register(&resources.PalaceTile{})
	gob.Register(&resources.SchoolTile{})
	gob.Register(&resources.InnovationTile{})
	gob.Register(&resources.FactionTile{})
	gob.Register(&resources.ColorTile{})
	gob.Register(&game.City{})
	gob.Register(&game.Turn{})
	//gob.Register(&color.Color{})
	gob.Register(&resources.Resource{})
	gob.Register(&resources.TileItem{})
	gob.Register(&resources.Position{})
	gob.Register(&resources.BridgePosition{})
	gob.Register(&resources.Price{})
	gob.Register(&resources.CityItem{})
	//gob.Register(&resources.Building{})
	gob.Register(&game.Mapitem{})

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
