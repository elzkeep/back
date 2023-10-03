package resources

import (
	"math/rand"
)

type RoundTile struct {
	Items []TileItem `json:"items"`
}

func NewRoundTile() *RoundTile {
	var item RoundTile

	item.Items = make([]TileItem, 0)

	return &item
}

func (p *RoundTile) Init(count int) {
	items := []TileItem{
		TileItem{Category: TileRound, Type: TileRoundSideVP, Name: "side VP", Ship: 1, Use: false},
		TileItem{Category: TileRound, Type: TileRoundPristVP, Name: "P VP", Receive: Price{Prist: 1}, Use: false},
		TileItem{Category: TileRound, Type: TileRoundTpVP, Name: "TP VP", Receive: Price{Power: 3}, Use: false},
		TileItem{Category: TileRound, Type: TileRoundShVP, Name: "SH/SA VP", Receive: Price{Worker: 1}, Pass: Price{ShVP: 4}, Use: false},
		TileItem{Category: TileRound, Type: TileRoundSpade, Name: "spd", Receive: Price{Book: 1}, Action: Price{Spade: 1}, Use: false},
		TileItem{Category: TileRound, Type: TileRoundBridge, Name: "bridge", Receive: Price{Book: 1}, Action: Price{Bridge: 1}, Use: false},
		TileItem{Category: TileRound, Type: TileRoundScienceCube, Name: "1 science", Receive: Price{Worker: 2}, Action: Price{Science: Science{Single: 1}}, Use: false},
		TileItem{Category: TileRound, Type: TileRoundSchoolScienceCoin, Name: "te science", Receive: Price{Coin: 4}, Pass: Price{Science: Science{Any: 1}}, Use: false},
		TileItem{Category: TileRound, Type: TileRoundPower, Name: "4PW", Receive: Price{Coin: 2}, Use: false},
		TileItem{Category: TileRound, Type: TileRoundCoin, Name: "6C", Receive: Price{Coin: 6}, Use: false},
	}

	rand.Shuffle(len(items), func(i, j int) { items[i], items[j] = items[j], items[i] })

	for i := 0; i < count+3; i++ {
		p.Items = append(p.Items, items[i])
	}
}

func (p *RoundTile) Start() {
	for i, _ := range p.Items {
		if p.Items[i].Use == true {
			return
		}

		p.Items[i].Receive.Coin++
	}
}

func (p *RoundTile) GetTile(pos int) *TileItem {
	return &p.Items[pos]
}

func (p *RoundTile) Pass(pos int) *TileItem {
	ret := &p.Items[pos]

	/*
		p.Items = append(p.Items[:pos], p.Items[pos+1:]...)
		p.Items = append(p.Items, *tile)
	*/

	return ret
}
