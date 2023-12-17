package services

import (
	"fmt"
	"zkeep/global"

	"github.com/antoniodipinto/ikisocket"
)

type ChatService struct {
	Use     bool
	Clients map[string]string
	Rooms   map[int64]map[string]string
}

type MessageObject struct {
	Data string `json:"data"`
	From string `json:"from"`
	To   string `json:"to"`
}

var chat ChatService

func (p *ChatService) Broadcast(room int64, message []byte) {
	for _, v := range p.Rooms[room] {
		err := ikisocket.EmitTo(v, message)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (p *ChatService) BroadcastWithout(room int64, user string, message []byte) {
	for _, v := range p.Rooms[room] {
		if user == v {
			continue
		}
		err := ikisocket.EmitTo(v, message)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func Chat() {
	chat.Use = true
	chat.Clients = make(map[string]string)
	chat.Rooms = make(map[int64]map[string]string)

	ikisocket.On(ikisocket.EventConnect, func(ep *ikisocket.EventPayload) {
		//fmt.Println(fmt.Sprintf("Connection event 1 - User: %s", ep.Kws.GetStringAttribute("id")))
		//
		//id := ep.Kws.GetStringAttribute("id")
		err := ep.Kws.EmitTo(ep.Kws.UUID, []byte("pong"))
		if err != nil {
			fmt.Println(err)
		}
	})

	ikisocket.On(ikisocket.EventMessage, func(ep *ikisocket.EventPayload) {
		//id := ep.Kws.GetStringAttribute("id")
		//fmt.Println(fmt.Sprintf("Message event - User: %s - Message: %s", id, string(ep.Data)))
		if string(ep.Data) == "ping" {
			err := ep.Kws.EmitTo(ep.Kws.UUID, []byte("pong"))
			if err != nil {
				fmt.Println(err)
			}
		}

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
		id := global.Atol(ep.Kws.GetStringAttribute("id"))

		delete(chat.Rooms[id], ep.Kws.UUID)
	})

	ikisocket.On(ikisocket.EventClose, func(ep *ikisocket.EventPayload) {
		id := global.Atol(ep.Kws.GetStringAttribute("id"))

		delete(chat.Rooms[id], ep.Kws.UUID)
	})

	ikisocket.On(ikisocket.EventError, func(ep *ikisocket.EventPayload) {
		id := ep.Kws.GetStringAttribute("id")

		fmt.Println(fmt.Sprintf("Error event - User: %s", id))
	})
}
