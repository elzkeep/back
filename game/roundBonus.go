package game

import (
	"aoi/game/resources"
	"aoi/models"
)

type RoundBonusType int

const (
	DVP RoundBonusType = iota
	TpVP
	TeVP
	ShSaVP
	SpadeVP
	ScienceVP
	CityVP
	AdvanceVP
	InnovationVP
	EdgeVP
)

type RoundBonusItem struct {
	Type    RoundBonusType    `json:"type"`
	Name    string            `json:"name"`
	Science resources.Science `json:"science"`
	Receive resources.Price   `json:"receive"`
	Build   resources.BuildVP `json:"build"`
}

type RoundBonus struct {
	Items      []RoundBonusItem `json:"items"`
	FinalRound RoundBonusItem   `json:"final"`
	Tiles      []RoundBonusItem `json:"-"`
}

func NewRoundBonus(id int64) *RoundBonus {
	var item RoundBonus

	item.Items = make([]RoundBonusItem, 0)

	tiles := []RoundBonusItem{
		{Type: DVP, Name: "D >> 2", Build: resources.BuildVP{D: 2}, Science: resources.Science{Law: 3}, Receive: resources.Price{Prist: 1}},
		{Type: DVP, Name: "D >> 2", Build: resources.BuildVP{D: 2}, Science: resources.Science{Banking: 3}, Receive: resources.Price{Power: 4}},
		{Type: TpVP, Name: "TP >> 3", Build: resources.BuildVP{TP: 3}, Science: resources.Science{Law: 3}, Receive: resources.Price{Book: resources.Book{Any: 1}}},
		{Type: TpVP, Name: "TP >> 3", Build: resources.BuildVP{TP: 3}, Science: resources.Science{Medicine: 4}, Receive: resources.Price{Spade: 1}},
		{Type: TeVP, Name: "TE >> 4", Build: resources.BuildVP{TE: 4}, Science: resources.Science{Banking: 1}, Receive: resources.Price{Coin: 1}},
		{Type: ShSaVP, Name: "SA/SH >> 5", Build: resources.BuildVP{SHSA: 5}, Science: resources.Science{Medicine: 2}, Receive: resources.Price{Worker: 1}},
		{Type: ShSaVP, Name: "SA/SH >> 5", Build: resources.BuildVP{SHSA: 5}, Science: resources.Science{Banking: 2}, Receive: resources.Price{Worker: 1}},
		{Type: SpadeVP, Name: "SPADE >> 2", Build: resources.BuildVP{Spade: 2}, Science: resources.Science{Engineering: 1}, Receive: resources.Price{Coin: 1}},
		{Type: ScienceVP, Name: "SCIENCE >> 1", Build: resources.BuildVP{Science: 1}, Science: resources.Science{Medicine: 3}, Receive: resources.Price{Book: resources.Book{Any: 1}}},
		{Type: CityVP, Name: "CITY >> 5", Build: resources.BuildVP{City: 5}, Science: resources.Science{Engineering: 4}, Receive: resources.Price{Spade: 1}},
		{Type: AdvanceVP, Name: "ADVANCE >> 3", Build: resources.BuildVP{Advance: 3}, Science: resources.Science{Engineering: 3}, Receive: resources.Price{Prist: 1}},
		{Type: InnovationVP, Name: "INNOVATION >> 5", Build: resources.BuildVP{Innovation: 5}, Science: resources.Science{Law: 2}, Receive: resources.Price{Power: 3}},
	}

	finalRound := []RoundBonusItem{
		{Type: DVP, Name: "D >> 2", Build: resources.BuildVP{D: 2}},
		{Type: TpVP, Name: "TP >> 3", Build: resources.BuildVP{TP: 3}},
		{Type: TeVP, Name: "TE >> 4", Build: resources.BuildVP{TE: 4}},
		{Type: EdgeVP, Name: "EDGE >> 3", Build: resources.BuildVP{Edge: 3}},
	}

	conn := models.NewConnection()
	defer conn.Close()

	gametileManager := models.NewGametileManager(conn)
	items := gametileManager.Find([]interface{}{
		models.Where{Column: "game", Value: id, Compare: "="},
		models.Where{Column: "type", Value: int(resources.TileRoundBonus), Compare: "="},
		models.Ordering("gt_order"),
	})

	for _, v := range items[:6] {
		item.Items = append(item.Items, tiles[v.Number])
	}

	item.FinalRound = finalRound[items[6].Number]

	item.Tiles = tiles

	return &item
}

func (p *RoundBonus) Get(pos int) RoundBonusItem {
	return p.Items[pos-1]
}

func (p *RoundBonus) GetBuildVP(pos int) resources.BuildVP {
	if pos < 1 {
		return resources.BuildVP{}
	}

	return p.Items[pos-1].Build
}
