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
	Items []PowerActionItem
}

func NewPowerAction() *PowerAction {
	var item PowerAction

	item.Items = make([]PowerActionItem, 0)

	item.Items = []PowerActionItem{
		PowerActionItem{Name: "1 bridge", Power: 3, Receive: resources.Price{Bridge: 1}, Use: false},
		PowerActionItem{Name: "2 prist", Power: 3, Receive: resources.Price{Prist: 1}, Use: false},
		PowerActionItem{Name: "2 worker", Power: 4, Receive: resources.Price{Worker: 2}, Use: false},
		PowerActionItem{Name: "7 coin", Power: 4, Receive: resources.Price{Coin: 7}, Use: false},
		PowerActionItem{Name: "1 spade", Power: 4, Receive: resources.Price{Spade: 1}, Use: false},
		PowerActionItem{Name: "2 space", Power: 6, Receive: resources.Price{Spade: 2}, Use: false},
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

func (p *PowerAction) Pass() {
	count := len(p.Items)

	for i := 0; i < count; i++ {
		p.Items[i].Use = false
	}
}
