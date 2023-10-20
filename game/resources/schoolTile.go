package resources

import "math/rand"

type SchoolTileType int

const (
	Banking SchoolTileType = iota
	Law
	Engineering
	Medicine
)

type SchoolTile struct {
	Items [][]SchoolTileItem `json:"items"`
}

type SchoolTileItem struct {
	Count int      `json:"count"`
	Tile  TileItem `json:"tile"`
}

func NewSchoolTile() *SchoolTile {
	var item SchoolTile

	item.Items = make([][]SchoolTileItem, 0)

	for i := 0; i < 4; i++ {
		items := make([]SchoolTileItem, 0)
		for j := 0; j < 3; j++ {
			items = append(items, SchoolTileItem{})
		}
		item.Items = append(item.Items, items)
	}

	return &item
}

func (p *SchoolTile) Init(count int) {
	items := []TileItem{
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

	rand.Shuffle(len(items), func(i, j int) { items[i], items[j] = items[j], items[i] })

	pos := 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			p.Items[i][j].Tile = items[pos]
			p.Items[i][j].Count = count
			pos++
		}
	}
}

func (p *SchoolTile) GetTile(pos SchoolTileType, level int) SchoolTileItem {
	return p.Items[pos][level]
}
