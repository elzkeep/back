package resources

import (
	"aoi/models"
)

type SchoolTileType int

const (
	Banking SchoolTileType = iota
	Law
	Engineering
	Medicine
)

type SchoolTile struct {
	Items    [][]SchoolTileItem `json:"items"`
	Original [][]SchoolTileItem `json:"-"`
	Count    int                `json:"-"`
}

type SchoolTileItem struct {
	Count int      `json:"count"`
	Tile  TileItem `json:"tile"`
}

func NewSchoolTile(id int64, count int) *SchoolTile {
	var item SchoolTile

	if count == 5 {
		count = 4
	}

	item.Count = count
	item.Items = make([][]SchoolTileItem, 0)

	for i := 0; i < 4; i++ {
		items := make([]SchoolTileItem, 0)
		for j := 0; j < 3; j++ {
			items = append(items, SchoolTileItem{})
		}
		item.Items = append(item.Items, items)
	}

	item.Original = make([][]SchoolTileItem, 0)

	for i := 0; i < 4; i++ {
		items := make([]SchoolTileItem, 0)
		for j := 0; j < 3; j++ {
			items = append(items, SchoolTileItem{})
		}
		item.Original = append(item.Original, items)
	}

	tiles := []TileItem{
		{Category: TileSchool, Type: TileSchoolWorker, Name: "Worker", Receive: Price{Worker: 1, Science: Science{Any: 1}}, Use: false},
		{Category: TileSchool, Type: TileSchoolSpade, Name: "Spade", Once: Price{Spade: 2}, Use: false},
		{Category: TileSchool, Type: TileSchoolPrist, Name: "Prist", Build: BuildVP{Prist: 2}, Use: false},
		{Category: TileSchool, Type: TileSchoolEdgeVP, Name: "EdgeVP", Build: BuildVP{Edge: 3}, Use: false},
		{Category: TileSchool, Type: TileSchoolCoin, Name: "Coin", Receive: Price{Coin: 2, VP: 3}, Use: false},
		{Category: TileSchool, Type: TileSchoolAnnex, Name: "Annex", Once: Price{Annex: 2}, Use: false},
		{Category: TileSchool, Type: TileSchoolNeutral, Name: "1 science", Receive: Price{Power: 2, Coin: 2}, Once: Price{Building: WHITE_TOWER}, Use: false},
		{Category: TileSchool, Type: TileSchoolBook, Name: "1 science", Receive: Price{Power: 1, Book: Book{Any: 1}}, Use: false},
		{Category: TileSchool, Type: TileSchoolVP, Name: "te science", Once: Price{Coin: 2, Worker: 1, VP: 5}, Use: false},
		{Category: TileSchool, Type: TileSchoolPower, Name: "6 coin", Action: Price{Power: 4}, Use: false},
		{Category: TileSchool, Type: TileSchoolPassCity, Name: "6 coin", Pass: Price{CityVP: 2}, Use: false},
		{Category: TileSchool, Type: TileSchoolPassPrist, Name: "6 coin", Pass: Price{ScienceVP: 1}, Use: false},
	}

	conn := models.NewConnection()
	defer conn.Close()

	gametileManager := models.NewGametileManager(conn)
	items := gametileManager.Find([]interface{}{
		models.Where{Column: "game", Value: id, Compare: "="},
		models.Where{Column: "type", Value: int(TileSchool), Compare: "="},
		models.Ordering("gt_order"),
	})

	i := 0
	j := 0
	for _, v := range items {
		for _, tile := range tiles {
			if v.Number == int(tile.Type) {
				item.Items[i][j].Tile = tile
				item.Items[i][j].Count = count

				item.Original[i][j].Tile = tile
				item.Original[i][j].Count = count

				j++
				if j == 3 {
					i++
					j = 0
				}
			}
		}
	}

	return &item
}

func (p *SchoolTile) Copy() *SchoolTile {
	var item SchoolTile

	item.Items = make([][]SchoolTileItem, 0)

	for i := 0; i < 4; i++ {
		items := make([]SchoolTileItem, 0)
		for j := 0; j < 3; j++ {
			tile := p.Original[i][j]
			tile.Tile.Use = false
			items = append(items, tile)
		}
		item.Items = append(item.Items, items)
	}

	return &item
}

func (p *SchoolTile) GetTile(pos SchoolTileType, level int) SchoolTileItem {
	return p.Items[pos][level]
}
