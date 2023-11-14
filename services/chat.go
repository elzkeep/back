package services

import (
	"fmt"

	"github.com/antoniodipinto/ikisocket"
)

type ChatService struct {
	Use     bool
	Clients map[string]string
}

type MessageObject struct {
	Data string `json:"data"`
	From string `json:"from"`
	To   string `json:"to"`
}

var chat ChatService

func Chat() {
	chat.Use = true
	chat.Clients = make(map[string]string)

	ikisocket.On(ikisocket.EventConnect, func(ep *ikisocket.EventPayload) {
		//fmt.Println(fmt.Sprintf("Connection event 1 - User: %s", ep.Kws.GetStringAttribute("id")))
	})

	ikisocket.On(ikisocket.EventMessage, func(ep *ikisocket.EventPayload) {
		//id := ep.Kws.GetStringAttribute("id")
		//fmt.Println(fmt.Sprintf("Message event - User: %s - Message: %s", id, string(ep.Data)))

		/*
			message := MessageObject{}

			// Unmarshal the json message
			// {
			//  "from": "<user-id>",
			//  "to": "<recipient-user-id>",
			//  "data": "hello"
			//}
			err := json.Unmarshal(ep.Data, &message)
			if err != nil {
				fmt.Println(err)
				return
			}

			// Emit the message directly to specified user
			err = ep.Kws.EmitTo(chat.Clients[message.To], ep.Data)
			if err != nil {
				fmt.Println(err)
			}
		*/
	})

	ikisocket.On(ikisocket.EventDisconnect, func(ep *ikisocket.EventPayload) {
		id := ep.Kws.GetStringAttribute("id")

		delete(chat.Clients, id)
		fmt.Println(fmt.Sprintf("Disconnection event - User: %s", id))
	})

	ikisocket.On(ikisocket.EventClose, func(ep *ikisocket.EventPayload) {
		id := ep.Kws.GetStringAttribute("id")

		delete(chat.Clients, id)
		fmt.Println(fmt.Sprintf("Close event - User: %s", id))
	})

	ikisocket.On(ikisocket.EventError, func(ep *ikisocket.EventPayload) {
		id := ep.Kws.GetStringAttribute("id")

		fmt.Println(fmt.Sprintf("Error event - User: %s", id))
	})
}
