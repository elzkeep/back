package action

import (
	"aoi/game/resources"
	"aoi/models"
	"sort"
)

type BookActionType int

const (
	Power5 BookActionType = iota
	Science
	Coin6
	TpUpgrade
	TpVP
	Spade3
)

type BookAction struct {
	Items    []BookActionItem `json:"items"`
	Original []BookActionItem `json:"-"`
}

func NewBookAction(id int64) *BookAction {
	var item BookAction

	item.Items = make([]BookActionItem, 0)

	tiles := []BookActionItem{
		{Type: Power5, Name: "5PW", Book: 1, Receive: resources.Price{Power: 5}, Use: false},
		{Type: Science, Name: "2 science", Book: 1, Receive: resources.Price{Science: resources.Science{Single: 2}}, Use: false},
		{Type: Coin6, Name: "6C", Book: 2, Receive: resources.Price{Coin: 6}, Use: false},
		{Type: TpUpgrade, Name: "TP up", Book: 2, Receive: resources.Price{TpUpgrade: 1}, Use: false},
		{Type: TpVP, Name: "TP VP", Book: 2, Receive: resources.Price{TpVP: 2}, Use: false},
		{Type: Spade3, Name: "3 spd", Book: 3, Receive: resources.Price{Spade: 3}, Use: false},
	}

	conn := models.NewConnection()
	defer conn.Close()

	gametileManager := models.NewGametileManager(conn)
	items := gametileManager.Find([]interface{}{
		models.Where{Column: "game", Value: id, Compare: "="},
		models.Where{Column: "type", Value: int(resources.TileBookAction), Compare: "="},
		models.Ordering("gt_order"),
	})

	for _, v := range items {
		for _, tile := range tiles {
			if v.Number == int(tile.Type) {
				item.Items = append(item.Items, tile)
			}
		}
	}

	sort.Slice(item.Items, func(i, j int) bool {
		return item.Items[i].Book > item.Items[j].Book
	})

	item.Original = make([]BookActionItem, len(item.Items))
	copy(item.Original, item.Items)

	return &item
}

func (p *BookAction) Copy() *BookAction {
	var item BookAction

	item.Items = make([]BookActionItem, len(p.Original))
	copy(item.Items, p.Original)
	for i := range item.Items {
		item.Items[i].Use = false
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

func (p *BookAction) Start() {
	count := len(p.Items)

	for i := 0; i < count; i++ {
		p.Items[i].Use = false
	}
}
