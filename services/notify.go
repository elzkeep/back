package services

import (
	"aoi/global"

	"github.com/antoniodipinto/ikisocket"
	log "github.com/sirupsen/logrus"
)

func Notify() {
	log.Println("Start Notify Service")
	go func() {
		ch := global.GetChannel()

		for {
			select {
			case item := <-ch:
				ikisocket.Broadcast([]byte(item.Title))
			}
		}
	}()
}
