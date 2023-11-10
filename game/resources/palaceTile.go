package resources

import (
	"aoi/models"
	"log"
)

type PalaceTile struct {
	Items []TileItem `json:"items"`
}

func NewPalaceTile(id int64) *PalaceTile {
	var item PalaceTile

	item.Items = make([]TileItem, 0)

	tiles := []TileItem{
		{Category: TilePalace, Type: TilePalaceWorker, Name: "2W", Receive: Price{Power: 5}, Action: Price{Worker: 2}, Use: false},
		{Category: TilePalace, Type: TilePalaceSpade, Name: "prist vp", Action: Price{Spade: 2}, Use: false},
		{Category: TilePalace, Type: TilePalaceDowngrade, Name: "tp vp", Receive: Price{Power: 2}, Action: Price{Downgrade: 1, Worker: 1, VP: 3}, Use: false},
		{Category: TilePalace, Type: TilePalaceTpUpgrade, Name: "sh/sa vp", Receive: Price{Power: 2}, Action: Price{TpUpgrade: 1}, Use: false},
		{Category: TilePalace, Type: TilePalaceSchoolTile, Name: "1 spade", Receive: Price{Power: 4}, Once: Price{Tile: 1}, Use: false},
		{Category: TilePalace, Type: TilePalaceScience, Name: "1 bridge", Receive: Price{Book: Book{Any: 1}, Power: 2}, Action: Price{Science: Science{Single: 2}}, Use: false},
		{Category: TilePalace, Type: TilePalaceSchoolVp, Name: "1 science", Receive: Price{Power: 4}, Pass: Price{TeVP: 3}, Use: false},
		{Category: TilePalace, Type: TilePalace6PowerCity, Name: "te science", Receive: Price{Coin: 2, Worker: 1, Power: 2}, Use: false},
		{Category: TilePalace, Type: TilePalaceJump, Name: "4 power", Receive: Price{Prist: 1}, Use: false},
		{Category: TilePalace, Type: TilePalacePower, Name: "4 power", Receive: Price{Coin: 6}, Once: Price{Book: Book{Any: 2}, Power: 12}, Use: false},
		{Category: TilePalace, Type: TilePalaceCity, Name: "6 coin", Receive: Price{Worker: 1}, Once: Price{City: 1}, Use: false},
		{Category: TilePalace, Type: TilePalaceDVp, Name: "6 coin", Receive: Price{Power: 8}, Build: BuildVP{D: 2}, Use: false},
		{Category: TilePalace, Type: TilePalaceTpVp, Name: "6 coin", Action: Price{Coin: 3, Book: Book{Any: 1}}, Build: BuildVP{TP: 3}, Use: false},
		{Category: TilePalace, Type: TilePalaceRiverCity, Name: "6 coin", Receive: Price{Power: 6, ShipUpgrade: 2}, Use: false},
		{Category: TilePalace, Type: TilePalaceBridge, Name: "6 coin", Receive: Price{Power: 6}, Once: Price{Book: Book{Any: 2}, Spade: 2, Bridge: 2}, Use: false},
		{Category: TilePalace, Type: TilePalaceTpBuild, Name: "tp build", Receive: Price{Power: 2, Book: Book{Any: 1}}, Once: Price{Building: TP}, Use: false},
		{Category: TilePalace, Type: TilePalaceVp, Name: "10 vp", Receive: Price{Power: 2}, Once: Price{VP: 10}, Use: false},
	}

	conn := models.NewConnection()
	defer conn.Close()

	gametileManager := models.NewGametileManager(conn)
	items := gametileManager.Find([]interface{}{
		models.Where{Column: "game", Value: id, Compare: "="},
		models.Where{Column: "type", Value: int(TilePalace), Compare: "="},
		models.Ordering("gt_order"),
	})

	for _, v := range items {
		for _, tile := range tiles {
			if v.Number == int(tile.Type) {
				item.Items = append(item.Items, tile)
			}
		}
	}

	return &item
}

func (p *PalaceTile) GetTile(pos int) TileItem {
	log.Println(p.Items[pos].Name)
	return p.Items[pos]
}

func (p *PalaceTile) Setup(pos int) {
	p.Items = append(p.Items[:pos], p.Items[pos+1:]...)
}
