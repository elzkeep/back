package resources

import (
	"aoi/models"
	"log"
)

type InnovationTile struct {
	Items [][]TileItem `json:"items"`
	Price []Price      `json:"price"`
}

func NewInnovationTile(id int64, count int) *InnovationTile {
	var item InnovationTile

	item.Items = make([][]TileItem, 6)

	for i := 0; i < 6; i++ {
		item.Items[i] = make([]TileItem, 0)
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

	for _, v := range targets {
		log.Println(v)
	}

	pos := 0
	for i, v := range counts[count-1] {
		tiles := make([]TileItem, 0)

		for j := 0; j < v; j++ {
			tiles = append(tiles, targets[pos])
			pos++
		}

		item.Items[i] = tiles
	}

	return &item
}

/*
func (p *InnovationTile) GetTile(pos int) TileItem {
	return p.Items[pos]
}

func (p *InnovationTile) Setup(pos int) {
	p.Items = append(p.Items[:pos], p.Items[pos+1:]...)
}
*/
