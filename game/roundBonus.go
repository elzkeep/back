package game

import (
	"aoi/game/resources"
	"math/rand"
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
}

func NewRoundBonus() *RoundBonus {
	var item RoundBonus

	item.Items = make([]RoundBonusItem, 0)

	items := []RoundBonusItem{
		RoundBonusItem{Type: DVP, Name: "D >> 2", Build: resources.BuildVP{D: 2}, Science: resources.Science{Law: 3}, Receive: resources.Price{Prist: 1}},
		RoundBonusItem{Type: DVP, Name: "D >> 2", Build: resources.BuildVP{D: 2}, Science: resources.Science{Banking: 3}, Receive: resources.Price{Power: 4}},
		RoundBonusItem{Type: TpVP, Name: "TP >> 3", Build: resources.BuildVP{TP: 3}, Science: resources.Science{Law: 3}, Receive: resources.Price{Book: resources.Book{Any: 1}}},
		RoundBonusItem{Type: TpVP, Name: "TP >> 3", Build: resources.BuildVP{TP: 3}, Science: resources.Science{Medicine: 4}, Receive: resources.Price{Spade: 1}},
		RoundBonusItem{Type: TeVP, Name: "TE >> 4", Build: resources.BuildVP{TE: 4}, Science: resources.Science{Banking: 1}, Receive: resources.Price{Coin: 1}},
		RoundBonusItem{Type: ShSaVP, Name: "SA/SH >> 5", Build: resources.BuildVP{SHSA: 5}, Science: resources.Science{Medicine: 2}, Receive: resources.Price{Worker: 1}},
		RoundBonusItem{Type: ShSaVP, Name: "SA/SH >> 5", Build: resources.BuildVP{SHSA: 5}, Science: resources.Science{Banking: 2}, Receive: resources.Price{Worker: 1}},
		RoundBonusItem{Type: SpadeVP, Name: "SPADE >> 2", Build: resources.BuildVP{Spade: 2}, Science: resources.Science{Engineering: 1}, Receive: resources.Price{Coin: 1}},
		RoundBonusItem{Type: ScienceVP, Name: "SCIENCE >> 1", Build: resources.BuildVP{Science: 1}, Science: resources.Science{Medicine: 3}, Receive: resources.Price{Book: resources.Book{Any: 1}}},
		RoundBonusItem{Type: CityVP, Name: "CITY >> 5", Build: resources.BuildVP{City: 5}, Science: resources.Science{Engineering: 4}, Receive: resources.Price{Spade: 1}},
		RoundBonusItem{Type: AdvanceVP, Name: "ADVANCE >> 3", Build: resources.BuildVP{Advance: 3}, Science: resources.Science{Engineering: 3}, Receive: resources.Price{Prist: 1}},
		RoundBonusItem{Type: InnovationVP, Name: "INNOVATION >> 5", Build: resources.BuildVP{Innovation: 5}, Science: resources.Science{Law: 2}, Receive: resources.Price{Power: 3}},
	}

	rand.Shuffle(len(items), func(i, j int) { items[i], items[j] = items[j], items[i] })

	for i := 0; i < 6; i++ {
		item.Items = append(item.Items, items[i])
	}

	finalRound := []RoundBonusItem{
		RoundBonusItem{Type: DVP, Name: "D >> 2", Build: resources.BuildVP{D: 2}},
		RoundBonusItem{Type: TpVP, Name: "TP >> 3", Build: resources.BuildVP{TP: 3}},
		RoundBonusItem{Type: TeVP, Name: "TE >> 4", Build: resources.BuildVP{TE: 4}},
		RoundBonusItem{Type: EdgeVP, Name: "EDGE >> 3", Build: resources.BuildVP{Edge: 3}},
	}

	item.FinalRound = finalRound[rand.Intn(len(finalRound))]

	return &item
}

func (p *RoundBonus) Get(pos int) RoundBonusItem {
	return p.Items[pos-1]
}

func (p *RoundBonus) GetBuildVP(pos int) resources.BuildVP {
	return p.Items[pos-1].Build
}
