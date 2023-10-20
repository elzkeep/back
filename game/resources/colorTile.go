package resources

import (
	"aoi/game/color"
	"math/rand"
)

var _colorTiles []TileItem

func init() {
	items := []TileItem{
		{Category: TileColor, Type: TileColorRed, Color: color.Red, Name: "Red", Once: Price{Book: Book{Any: 1}, Worker: 1}, Use: false},
		{Category: TileColor, Type: TileColorYellow, Color: color.Yellow, Name: "Yellow", Once: Price{Spade: 1}, Use: false},
		{Category: TileColor, Type: TileColorBrown, Color: color.Brown, Name: "Brown", Use: false},
		{Category: TileColor, Type: TileColorBlack, Color: color.Black, Name: "Black", Once: Price{Power: 2, Prist: 1}, Use: false},
		{Category: TileColor, Type: TileColorBlue, Color: color.Blue, Name: "Blue", Once: Price{ShipUpgrade: 1}, Use: false},
		{Category: TileColor, Type: TileColorGreen, Color: color.Green, Name: "Green", Once: Price{Power: 1, Science: Science{Banking: 1, Law: 1, Engineering: 1, Medicine: 1}}, Use: false},
		{Category: TileColor, Type: TileColorGray, Color: color.Gray, Name: "Gray", Use: false},
	}

	rand.Shuffle(len(items), func(i, j int) { items[i], items[j] = items[j], items[i] })

	_colorTiles = items
}

func GetColorTile(value color.Color) TileItem {
	pos := int(value) - 2

	return _colorTiles[pos]
}
