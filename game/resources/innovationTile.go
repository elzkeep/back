package resources

import (
	"aoi/models"
)

type InnovationTile struct {
	Items    [][]TileItem `json:"items"`
	Price    []Price      `json:"price"`
	Original [][]TileItem `json:"-"`
}

func NewInnovationTile(id int64, count int) *InnovationTile {
	var item InnovationTile

	item.Items = make([][]TileItem, 6)

	for i := 0; i < 6; i++ {
		item.Items[i] = make([]TileItem, 0)
	}

	item.Original = make([][]TileItem, 6)

	for i := 0; i < 6; i++ {
		item.Original[i] = make([]TileItem, 0)
	}

	item.Price = []Price{
		{Book: Book{Any: 3, Banking: 2}},
		{Book: Book{Any: 3, Law: 2}},
		{Book: Book{Any: 3, Engineering: 2}},
		{Book: Book{Any: 3, Medicine: 2}},
		{Book: Book{Any: 1, Banking: 2, Law: 2}},
		{Book: Book{Any: 1, Engineering: 2, Medicine: 2}},
	}

	tiles := []TileItem{
		{Category: TileInnovation, Type: TileInnovationSpade, Name: "spade", Action: Price{Spade: 1}, Once: Price{Book: Book{Any: 1}, Science: Science{Banking: 1, Law: 1, Engineering: 1, Medicine: 1}}, Use: false},
		{Category: TileInnovation, Type: TileInnovationTP, Name: "tp", Pass: Price{TpVP: 2}, Use: false},
		{Category: TileInnovation, Type: TileInnovationPrist, Name: "prist", Action: Price{Prist: 1, VP: 3}, Use: false},

		{Category: TileInnovation, Type: TileInnovationKind, Name: "kind", Once: Price{VP: 10}, Use: false},
		{Category: TileInnovation, Type: TileInnovationCount, Name: "count", Use: false},
		{Category: TileInnovation, Type: TileInnovationSchool, Name: "school", Once: Price{TeVP: 5}, Use: false},
		{Category: TileInnovation, Type: TileInnovationCity, Name: "city", Once: Price{CityVP: 5}, Use: false},
		{Category: TileInnovation, Type: TileInnovationScience, Name: "science", Use: false},
		{Category: TileInnovation, Type: TileInnovationCluster, Name: "cluster", Use: false},
		{Category: TileInnovation, Type: TileInnovationD, Name: "d", Once: Price{DVP: 2}, Use: false},
		{Category: TileInnovation, Type: TileInnovationUpgrade, Name: "upgrade", Once: Price{Prist: 1, SpadeUpgrade: 1, ShipUpgrade: 1}, Use: false},
		{Category: TileInnovation, Type: TileInnovationBridge, Name: "bridge", Use: false},

		{Category: TileInnovation, Type: TileInnovationFreeD, Name: "freeD", Receive: Price{Worker: 3}, Once: Price{Building: WHITE_D}, Use: false},
		{Category: TileInnovation, Type: TileInnovationFreeTP, Name: "freeTP", Receive: Price{Coin: 5}, Once: Price{Building: WHITE_TP}, Use: false},
		{Category: TileInnovation, Type: TileInnovationFreeSchool, Name: "freeSchool", Once: Price{Tile: 1, Building: WHITE_TE}, Use: false},
		{Category: TileInnovation, Type: TileInnovationFreeSA, Name: "freeSA", Receive: Price{VP: 2}, Once: Price{Building: WHITE_SA}, Use: false},
		{Category: TileInnovation, Type: TileInnovationFreeSH, Name: "freeSH", Receive: Price{Power: 4}, Once: Price{Building: WHITE_SH, Token: 2}, Use: false},
		{Category: TileInnovation, Type: TileInnovationFreeMT, Name: "freeMT", Once: Price{VP: 7, Building: WHITE_MT}, Use: false},
	}

	conn := models.NewConnection()
	defer conn.Close()

	gametileManager := models.NewGametileManager(conn)
	items := gametileManager.Find([]interface{}{
		models.Where{Column: "game", Value: id, Compare: "="},
		models.Where{Column: "type", Value: int(TileInnovation), Compare: "="},
		models.Ordering("gt_order"),
	})

	targets := make([]TileItem, 0)

	for _, v := range items {
		for _, tile := range tiles {
			if v.Number == int(tile.Type) {
				targets = append(targets, tile)
			}
		}
	}

	counts := [][]int{
		{1, 1, 1, 1, 0, 0},
		{1, 1, 1, 1, 1, 1},
		{2, 2, 2, 2, 0, 0},
		{2, 2, 2, 2, 1, 1},
		{3, 3, 3, 3, 0, 0},
	}

	pos := 0
	for i, v := range counts[count-1] {
		tiles := make([]TileItem, 0)

		for j := 0; j < v; j++ {
			tiles = append(tiles, targets[pos])
			pos++
		}

		item.Items[i] = tiles
		item.Original[i] = tiles
	}

	return &item
}

func (p *InnovationTile) Copy() *InnovationTile {
	var item InnovationTile

	item.Items = make([][]TileItem, 6)

	for i := 0; i < 6; i++ {
		item.Items[i] = make([]TileItem, len(p.Original[i]))
		copy(item.Items[i], p.Original[i])
		for j := range item.Items[i] {
			item.Items[i][j].Use = false
		}
	}

	item.Price = make([]Price, len(p.Price))
	copy(item.Price, p.Price)

	return &item
}

func (p *InnovationTile) GetTile(pos int, index int) TileItem {
	return p.Items[pos][index]
}

func (p *InnovationTile) Setup(pos int, index int) {
	p.Items[pos][index].Use = true
}
