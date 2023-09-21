package resources

import "math/rand"

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
	Type    RoundTileType
	Name    string
	Receive Price
	Use     bool
	Ship    int
}

type RoundTile struct {
	Items []RoundTileItem
}

func NewRoundTile() *RoundTile {
	var item RoundTile

	item.Items = make([]RoundTileItem, 0)

	return &item
}

func (p *RoundTile) Init(count int) {
	items := []RoundTileItem{
		RoundTileItem{Type: SideVP, Name: "side vp", Ship: 1, Use: false},
		RoundTileItem{Type: PristVP, Name: "prist vp", Receive: Price{Prist: 1}, Use: false},
		RoundTileItem{Type: TpVP, Name: "tp vp", Receive: Price{Power: 3}, Use: false},
		RoundTileItem{Type: ShVP, Name: "sh/sa vp", Receive: Price{Worker: 1}, Use: false},
		RoundTileItem{Type: Spade, Name: "1 spade", Receive: Price{Book: 1}, Use: false},
		RoundTileItem{Type: Bridge, Name: "1 bridge", Receive: Price{Book: 1}, Use: false},
		RoundTileItem{Type: ScienceCube, Name: "1 science", Receive: Price{Worker: 2}, Use: false},
		RoundTileItem{Type: TeScienceCoin, Name: "te science", Receive: Price{Coin: 4}, Use: false},
		RoundTileItem{Type: Power, Name: "4 power", Receive: Price{Coin: 2}, Use: false},
		RoundTileItem{Type: Coin, Name: "6 coin", Receive: Price{Coin: 6}, Use: false},
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

func (p *RoundTile) GetTile(pos RoundTileType) *RoundTileItem {
	return &p.Items[pos]
}
