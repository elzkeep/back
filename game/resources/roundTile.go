package resources

import (
	"math/rand"
)

type RoundTileType int

const (
	SideVP RoundTileType = iota
	PristVP
	TpVP
	ShVP
	Spade
	Bridge
	ScienceCube
	TeScienceCoin
	Power
	Coin
)

type RoundTileItem struct {
	Type    RoundTileType `json:"type"`
	Name    string        `json:"name"`
	Receive Price         `json:"receive"`
	Pass    Price         `json:"pass"`
	Use     bool          `json:"use"`
	Ship    int           `json:"ship"`
}

type RoundTile struct {
	Items []RoundTileItem `json:"items"`
}

func NewRoundTile() *RoundTile {
	var item RoundTile

	item.Items = make([]RoundTileItem, 0)

	return &item
}

func (p *RoundTile) Init(count int) {
	items := []RoundTileItem{
		RoundTileItem{Type: SideVP, Name: "side VP", Ship: 1, Use: false},
		RoundTileItem{Type: PristVP, Name: "P VP", Receive: Price{Prist: 1}, Use: false},
		RoundTileItem{Type: TpVP, Name: "TP VP", Receive: Price{Power: 3}, Use: false},
		RoundTileItem{Type: ShVP, Name: "SH/SA VP", Receive: Price{Worker: 1}, Pass: Price{ShVP: 4}, Use: false},
		RoundTileItem{Type: Spade, Name: "spd", Receive: Price{Book: 1}, Use: false},
		RoundTileItem{Type: Bridge, Name: "bridge", Receive: Price{Book: 1}, Use: false},
		RoundTileItem{Type: ScienceCube, Name: "1 science", Receive: Price{Worker: 2}, Use: false},
		RoundTileItem{Type: TeScienceCoin, Name: "te science", Receive: Price{Coin: 4}, Pass: Price{Science: Science{Any: 1}}, Use: false},
		RoundTileItem{Type: Power, Name: "4PW", Receive: Price{Coin: 2}, Use: false},
		RoundTileItem{Type: Coin, Name: "6C", Receive: Price{Coin: 6}, Use: false},
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

func (p *RoundTile) GetTile(pos int) *RoundTileItem {
	return &p.Items[pos]
}

func (p *RoundTile) Pass(pos int) *RoundTileItem {
	ret := &p.Items[pos]

	/*
		p.Items = append(p.Items[:pos], p.Items[pos+1:]...)
		p.Items = append(p.Items, *tile)
	*/

	return ret
}
