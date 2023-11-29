package services

import (
	"aoi/global"

	log "github.com/sirupsen/logrus"
)

func Notify() {
	log.Println("Start Notify Service")
	go func() {
		ch := global.GetChannel()

		for {
			select {
			case item := <-ch:
				chat.Broadcast(item.Id, []byte(item.Title))
			}
		}
	}()
}
