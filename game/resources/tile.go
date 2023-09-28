package resources

import "math/rand"

type TileCategory int

const (
	_ TileCategory = iota
	TileSH
	TileTe
	TileInnovation
)

type TileType int

const (
	TileWorker TileType = iota
	TileSpade
	TileDowngrade
	TileTpUpgrade
	TileTeTile
	TileScience
	TileTeVp
	Tile6PowerCity
	TileJump
	TileCity
	TileDVp
	TileTpVp
	TileRiverCity
	TileBridge
	TileTpBuild
	TileVp
)

type TileItem struct {
	Type     TileType     `json:"type"`
	Category TileCategory `json:"category"`
	Name     string       `json:"name"`
	Receive  Price        `json:"receive"`
	Action   Price        `json:"action"`
	Once     Price        `json:"once"`
	Pass     Price        `json:"pass"`
	Use      bool         `json:"use"`
}

type Tile struct {
	Items []TileItem `json:"items"`
}

func NewTile() *Tile {
	var item Tile

	item.Items = make([]TileItem, 0)

	return &item
}

func (p *Tile) Init(count int) {
	items := []TileItem{
		TileItem{Category: TileSH, Type: TileWorker, Name: "side vp", Receive: Price{Power: 5}, Action: Price{Worker: 2}, Use: false},
		TileItem{Category: TileSH, Type: TileSpade, Name: "prist vp", Action: Price{Spade: 2}, Use: false},
		TileItem{Category: TileSH, Type: TileDowngrade, Name: "tp vp", Receive: Price{Power: 2}, Action: Price{Downgrade: 1}, Use: false},
		TileItem{Category: TileSH, Type: TileTpUpgrade, Name: "sh/sa vp", Receive: Price{Power: 2}, Action: Price{TpUpgrade: 1}, Use: false},
		TileItem{Category: TileSH, Type: TileTeTile, Name: "1 spade", Receive: Price{Power: 4}, Once: Price{Tile: 1}, Use: false},
		TileItem{Category: TileSH, Type: TileScience, Name: "1 bridge", Receive: Price{Book: 1, Power: 2}, Action: Price{Science: Science{Single: 2}}, Use: false},
		TileItem{Category: TileSH, Type: TileTeVp, Name: "1 science", Receive: Price{Power: 4}, Pass: Price{TeVP: 3}, Use: false},
		TileItem{Category: TileSH, Type: Tile6PowerCity, Name: "te science", Receive: Price{Coin: 2, Worker: 1, Power: 2}, Use: false},
		TileItem{Category: TileSH, Type: TileJump, Name: "4 power", Receive: Price{Prist: 1}, Use: false},
		TileItem{Category: TileSH, Type: TileCity, Name: "6 coin", Receive: Price{Worker: 1}, Once: Price{City: 1}, Use: false},
		TileItem{Category: TileSH, Type: TileDVp, Name: "6 coin", Receive: Price{Power: 8}, Use: false},
		TileItem{Category: TileSH, Type: TileTpVp, Name: "6 coin", Action: Price{Coin: 3, Book: 1}, Use: false},
		TileItem{Category: TileSH, Type: TileRiverCity, Name: "6 coin", Receive: Price{Power: 6}, Use: false},
		TileItem{Category: TileSH, Type: TileBridge, Name: "6 coin", Receive: Price{Power: 6}, Once: Price{Book: 2, Spade: 2, Bridge: 2}, Use: false},
		TileItem{Category: TileSH, Type: TileTpBuild, Name: "6 coin", Receive: Price{Power: 2, Book: 1}, Use: false},
	}

	p.Items = append(p.Items, TileItem{Category: TileSH, Type: TileVp, Name: "6 coin", Receive: Price{Power: 2}, Once: Price{VP: 10}, Use: false})

	rand.Shuffle(len(items), func(i, j int) { items[i], items[j] = items[j], items[i] })

	for i := 0; i < count+1; i++ {
		p.Items = append(p.Items, items[i])
	}
}

func (p *Tile) GetTile(pos TileType) *TileItem {
	return &p.Items[pos]
}
