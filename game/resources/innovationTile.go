package resources

import "math/rand"

type InnovationTile struct {
	Items [][]TileItem `json:"items"`
}

func NewInnovationTile() *InnovationTile {
	var item InnovationTile

	item.Items = make([][]TileItem, 0)

	for i := 0; i < 6; i++ {
		item.Items[i] = make([]TileItem, 0)
	}

	return &item
}

func (p *InnovationTile) Init(count int) {
	items := []TileItem{
		{Category: TileInnovation, Type: TileInnovationKind, Name: "2W", Once: Price{VP: 10}, Use: false},
		{Category: TileInnovation, Type: TileInnovationCount, Name: "prist vp", Use: false},
		{Category: TileInnovation, Type: TileInnovationSchool, Name: "tp vp", Once: Price{TeVP: 5}, Use: false},
		{Category: TileInnovation, Type: TileInnovationCity, Name: "sh/sa vp", Once: Price{CityVP: 5}, Use: false},
		{Category: TileInnovation, Type: TileInnovationScience, Name: "1 spade", Use: false},
		{Category: TileInnovation, Type: TileInnovationCluster, Name: "1 bridge", Use: false},
		{Category: TileInnovation, Type: TileInnovationD, Name: "1 science", Once: Price{DVP: 2}, Use: false},
		{Category: TileInnovation, Type: TileInnovationUpgrade, Name: "te science", Once: Price{Prist: 1, SpadeUpgrade: 1, ShipUpgrade: 1}, Use: false},
		{Category: TileInnovation, Type: TileInnovationBridge, Name: "4 power", Use: false},
		{Category: TileInnovation, Type: TileInnovationFreeD, Name: "4 power", Receive: Price{Worker: 3}, Once: Price{Building: WHITE_D}, Use: false},
		{Category: TileInnovation, Type: TileInnovationFreeTP, Name: "6 coin", Receive: Price{Coin: 5}, Once: Price{Building: WHITE_TP}, Use: false},
		{Category: TileInnovation, Type: TileInnovationFreeSchool, Name: "6 coin", Once: Price{Tile: 1, Building: WHITE_TE}, Use: false},
		{Category: TileInnovation, Type: TileInnovationFreeSA, Name: "6 coin", Receive: Price{VP: 2}, Once: Price{Building: WHITE_SA}, Use: false},
		{Category: TileInnovation, Type: TileInnovationFreeSH, Name: "6 coin", Receive: Price{Power: 4}, Once: Price{Building: WHITE_SH, Token: 2}, Use: false},
		{Category: TileInnovation, Type: TileInnovationFreeMT, Name: "6 coin", Once: Price{VP: 7, Building: WHITE_MT}, Use: false},
	}

	rand.Shuffle(len(items), func(i, j int) { items[i], items[j] = items[j], items[i] })

	counts := [][]int{
		{1, 1, 1, 1, 1, 1},
		{2, 2, 2, 2},
		{2, 2, 2, 2, 1, 1},
		{3, 3, 3, 3, 3, 3},
	}

	pos := 0
	for i, v := range counts[count-2] {
		tiles := make([]TileItem, 0)

		for j := 0; j < v; j++ {
			tiles = append(tiles, items[pos])
			pos++
		}

		p.Items[i] = tiles
	}
}

/*
func (p *InnovationTile) GetTile(pos int) TileItem {
	return p.Items[pos]
}

func (p *InnovationTile) Setup(pos int) {
	p.Items = append(p.Items[:pos], p.Items[pos+1:]...)
}
*/
