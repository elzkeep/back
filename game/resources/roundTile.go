package resources

import (
	"log"
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
		{Category: TileRound, Type: TileRoundEdgeVP, Name: "side VP", Build: BuildVP{River: 2}, Ship: 1, Use: false},
		{Category: TileRound, Type: TileRoundPristVP, Name: "P VP", Receive: Price{Prist: 1}, Build: BuildVP{Prist: 2}, Use: false},
		{Category: TileRound, Type: TileRoundTpVP, Name: "TP VP", Receive: Price{Power: 3}, Build: BuildVP{TP: 3}, Use: false},
		{Category: TileRound, Type: TileRoundShVP, Name: "SH/SA VP", Receive: Price{Worker: 1}, Pass: Price{ShVP: 4}, Use: false},
		{Category: TileRound, Type: TileRoundSpade, Name: "spd", Receive: Price{Book: Book{Any: 1}}, Action: Price{Spade: 1}, Use: false},
		{Category: TileRound, Type: TileRoundBridge, Name: "bridge", Receive: Price{Book: Book{Any: 1}}, Action: Price{Bridge: 1}, Use: false},
		{Category: TileRound, Type: TileRoundScienceCube, Name: "1 science", Receive: Price{Worker: 2}, Action: Price{Science: Science{Single: 1}}, Use: false},
		{Category: TileRound, Type: TileRoundSchoolScienceCoin, Name: "te science", Receive: Price{Coin: 4}, Pass: Price{Science: Science{Any: 1}}, Use: false},
		{Category: TileRound, Type: TileRoundPower, Name: "4PW", Receive: Price{Coin: 2, Power: 4}, Use: false},
		{Category: TileRound, Type: TileRoundCoin, Name: "6C", Receive: Price{Coin: 6}, Use: false},
	}

	rand.Shuffle(len(items), func(i, j int) { items[i], items[j] = items[j], items[i] })

	/*
		for i := 0; i < count+3; i++ {
			p.Items = append(p.Items, items[i])
		}
	*/

	p.Items = items
}

func (p *RoundTile) Start() {
	for i, _ := range p.Items {
		p.Items[i].Use = false
		p.Items[i].Coin++
	}
}

func (p *RoundTile) Pass(pos int) TileItem {
	log.Println("pos", pos)
	ret := p.Items[pos]
	log.Println(ret.Name)

	p.Items = append(p.Items[:pos], p.Items[pos+1:]...)

	return ret
}

func (p *RoundTile) Add(tile TileItem) {
	p.Items = append(p.Items, tile)
}
