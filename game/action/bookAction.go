package action

import (
	"aoi/game/resources"
	"math/rand"
)

type BookActionType int

type BookAction struct {
	Items []BookActionItem
}

func NewBookAction() *BookAction {
	var item BookAction

	item.Items = make([]BookActionItem, 0)

	items := []BookActionItem{
		BookActionItem{Name: "5 power", Book: 1, Receive: resources.Price{Power: 5}, Use: false},
		BookActionItem{Name: "2 science", Book: 1, Receive: resources.Price{Science: resources.Science{Single: 2}}, Use: false},
		BookActionItem{Name: "6 coin", Book: 2, Receive: resources.Price{Coin: 6}, Use: false},
		BookActionItem{Name: "tp upgrade", Book: 2, Receive: resources.Price{TpUpgrade: 1}, Use: false},
		BookActionItem{Name: "tp vp", Book: 2, Receive: resources.Price{TpVP: 2}, Use: false},
		BookActionItem{Name: "3 space", Book: 3, Receive: resources.Price{Spade: 3}, Use: false},
	}

	rand.Shuffle(len(items), func(i, j int) { items[i], items[j] = items[j], items[i] })

	for i := 0; i < 3; i++ {
		item.Items = append(item.Items, items[i])
	}

	return &item
}

func (p *BookAction) GetNeedBook(pos int) int {
	return p.Items[pos].Book
}

func (p *BookAction) IsUse(pos int) bool {
	return p.Items[pos].Use
}

func (p *BookAction) Action(pos int) BookActionItem {
	p.Items[pos].Use = true
	return p.Items[int(pos)]
}

func (p *BookAction) Pass() {
	count := len(p.Items)

	for i := 0; i < count; i++ {
		p.Items[i].Use = false
	}
}
