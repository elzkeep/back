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
	Items    []PowerActionItem `json:"items"`
	Original []PowerActionItem `json:"-"`
}

func NewPowerAction() *PowerAction {
	var item PowerAction

	item.Items = make([]PowerActionItem, 0)

	item.Items = []PowerActionItem{
		{Type: Bridge, Name: "bridge", Power: 3, Receive: resources.Price{Bridge: 1}, Use: false},
		{Type: Prist, Name: "P", Power: 3, Receive: resources.Price{Prist: 1}, Use: false},
		{Type: Worker, Name: "2W", Power: 4, Receive: resources.Price{Worker: 2}, Use: false},
		{Type: Coin, Name: "7C", Power: 4, Receive: resources.Price{Coin: 7}, Use: false},
		{Type: Spade, Name: "1 spd", Power: 4, Receive: resources.Price{Spade: 1}, Use: false},
		{Type: Spade2, Name: "2 spd", Power: 6, Receive: resources.Price{Spade: 2}, Use: false},
	}

	item.Original = make([]PowerActionItem, len(item.Items))
	copy(item.Original, item.Items)

	return &item
}

func (p *PowerAction) Copy() *PowerAction {
	var item PowerAction

	item.Items = make([]PowerActionItem, len(p.Original))
	copy(item.Items, p.Original)
	for i := range item.Items {
		item.Items[i].Use = false
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
