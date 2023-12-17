package services

import (
	"context"
	"zkeep/global"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"

	log "github.com/sirupsen/logrus"
)

func Fcm() {
	log.Println("Fcm Service Start")

	go func() {
		ch := global.GetChannel()

		ctx := context.Background()
		opt := option.WithCredentialsFile("./fcm.json")
		app, err := firebase.NewApp(ctx, nil, opt)
		if err != nil {
			log.Printf("error initializing app: %v\n", err)
			return
		}

		client, err := app.Messaging(ctx)
		if err != nil {
			log.Fatalf("error getting Messaging client: %v\n", err)
		}

		for {
			select {
			case item := <-ch:
				log.Println("send message")
				log.Println(item)
				noti := messaging.Notification{Title: item.Title}
				message := &messaging.MulticastMessage{
					Data:         item.Message,
					Notification: &noti,
					Tokens:       item.Target,
				}

				br, err := client.SendMulticast(context.Background(), message)
				if err != nil {
					log.Println(err)
				}

				log.Printf("%d messages were sent successfully\n", br.SuccessCount)
			}
		}
	}()
}
