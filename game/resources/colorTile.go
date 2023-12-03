package resources

import (
	"aoi/game/color"
	"aoi/models"
)

type ColorTile struct {
	Items    []TileItem `json:"items"`
	Original []TileItem `json:"-"`
}

func NewColorTile(id int64) *ColorTile {
	var item ColorTile

	item.Items = make([]TileItem, 0)

	tiles := []TileItem{
		{Category: TileColor, Type: TileColorRed, Color: color.Red, Name: "Red", Once: Price{Book: Book{Any: 1}, Worker: 1}, Use: false},
		{Category: TileColor, Type: TileColorYellow, Color: color.Yellow, Name: "Yellow", Once: Price{Spade: 1}, Use: false},
		{Category: TileColor, Type: TileColorBrown, Color: color.Brown, Name: "Brown", Use: false},
		{Category: TileColor, Type: TileColorBlack, Color: color.Black, Name: "Black", Once: Price{Power: 2, Prist: 1}, Use: false},
		{Category: TileColor, Type: TileColorBlue, Color: color.Blue, Name: "Blue", Once: Price{ShipUpgrade: 1}, Use: false},
		{Category: TileColor, Type: TileColorGreen, Color: color.Green, Name: "Green", Once: Price{Power: 1, Science: Science{Banking: 1, Law: 1, Engineering: 1, Medicine: 1}}, Use: false},
		{Category: TileColor, Type: TileColorGray, Color: color.Gray, Name: "Gray", Use: false},
	}

	conn := models.NewConnection()
	defer conn.Close()

	gametileManager := models.NewGametileManager(conn)
	items := gametileManager.Find([]interface{}{
		models.Where{Column: "game", Value: id, Compare: "="},
		models.Where{Column: "type", Value: int(TileColor), Compare: "="},
		models.Ordering("gt_order"),
	})

	for _, v := range items {
		for _, tile := range tiles {
			if v.Number == int(tile.Type) {
				item.Items = append(item.Items, tile)
			}
		}
	}

	item.Original = make([]TileItem, len(item.Items))
	copy(item.Original, item.Items)

	return &item
}

func (p *ColorTile) Copy() *ColorTile {
	var item ColorTile

	item.Items = make([]TileItem, len(p.Original))
	copy(item.Items, p.Original)
	for i := range item.Items {
		item.Items[i].Use = false
	}

	return &item
}

func (p *ColorTile) GetColorTile(value color.Color) TileItem {
	pos := int(value) - 2

	return p.Items[pos]
}
