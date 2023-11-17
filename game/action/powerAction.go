package action

import (
	"aoi/game/resources"
)

type PowerActionType int

const (
	Bridge PowerActionType = iota
	Prist
	Worker
	Coin
	Spade
	Spade2
)

type PowerAction struct {
	Items []PowerActionItem `json:"items"`
}

func NewPowerAction() *PowerAction {
	var item PowerAction

	item.Items = make([]PowerActionItem, 0)

	item.Items = []PowerActionItem{
		PowerActionItem{Type: Bridge, Name: "bridge", Power: 3, Receive: resources.Price{Bridge: 1}, Use: false},
		PowerActionItem{Type: Prist, Name: "P", Power: 3, Receive: resources.Price{Prist: 1}, Use: false},
		PowerActionItem{Type: Worker, Name: "2W", Power: 4, Receive: resources.Price{Worker: 2}, Use: false},
		PowerActionItem{Type: Coin, Name: "7C", Power: 4, Receive: resources.Price{Coin: 7}, Use: false},
		PowerActionItem{Type: Spade, Name: "1 spd", Power: 4, Receive: resources.Price{Spade: 1}, Use: false},
		PowerActionItem{Type: Spade2, Name: "2 spd", Power: 6, Receive: resources.Price{Spade: 2}, Use: false},
	}

	return &item
}

func (p *PowerAction) GetNeedPower(pos int) int {
	return p.Items[pos].Power
}

func (p *PowerAction) IsUse(pos int) bool {
	return p.Items[pos].Use
}

func (p *PowerAction) Action(pos int) PowerActionItem {
	p.Items[pos].Use = true
	return p.Items[pos]
}

func (p *PowerAction) Start() {
	count := len(p.Items)

	for i := 0; i < count; i++ {
		p.Items[i].Use = false
	}
}
