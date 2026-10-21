package services

import (
	"fmt"
	"sync"
	"zkeep/global"

	"github.com/antoniodipinto/ikisocket"
)

type ChatService struct {
	Use     bool
	Clients map[string]string
	Rooms   map[int64]map[string]string

	// Clients, Rooms는 여러 웹소켓 핸들러에서 동시에 접근하므로 보호 필요
	mutex sync.RWMutex
}

type MessageObject struct {
	Data string `json:"data"`
	From string `json:"from"`
	To   string `json:"to"`
}

var chat ChatService

func (p *ChatService) Join(id string, room int64, uuid string) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.Clients[id] = uuid
	if _, exists := p.Rooms[room]; !exists {
		p.Rooms[room] = make(map[string]string)
	}
	p.Rooms[room][uuid] = uuid
}

func (p *ChatService) Leave(room int64, uuid string) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	delete(p.Rooms[room], uuid)
	if len(p.Rooms[room]) == 0 {
		delete(p.Rooms, room)
	}
}

// 전송 중에는 잠금을 잡지 않도록 방 참가자 목록을 복사해서 반환
func (p *ChatService) members(room int64) []string {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	items := make([]string, 0, len(p.Rooms[room]))
	for _, v := range p.Rooms[room] {
		items = append(items, v)
	}

	return items
}

func (p *ChatService) Broadcast(room int64, message []byte) {
	for _, v := range p.members(room) {
		err := ikisocket.EmitTo(v, message)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (p *ChatService) BroadcastWithout(room int64, user string, message []byte) {
	for _, v := range p.members(room) {
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

		chat.Leave(id, ep.Kws.UUID)
	})

	ikisocket.On(ikisocket.EventClose, func(ep *ikisocket.EventPayload) {
		id := global.Atol(ep.Kws.GetStringAttribute("id"))

		chat.Leave(id, ep.Kws.UUID)
	})

	ikisocket.On(ikisocket.EventError, func(ep *ikisocket.EventPayload) {
		id := ep.Kws.GetStringAttribute("id")

		fmt.Println(fmt.Sprintf("Error event - User: %s", id))
	})
}
