package resources

import (
	"aoi/game/color"
	"log"
)

var _colorTiles []TileItem

func init() {
	_colorTiles = []TileItem{
		TileItem{Category: TileColor, Type: TileColorRed, Color: color.Red, Name: "1 science", Once: Price{Book: 1, Worker: 1}, Use: false},
		TileItem{Category: TileColor, Type: TileColorYellow, Color: color.Yellow, Name: "side VP", Once: Price{Spade: 1}, Use: false},
		TileItem{Category: TileColor, Type: TileColorBrown, Color: color.Brown, Name: "spd", Use: false},
		TileItem{Category: TileColor, Type: TileColorBlack, Color: color.Black, Name: "bridge", Once: Price{Power: 2, Prist: 1}, Use: false},
		TileItem{Category: TileColor, Type: TileColorBlue, Color: color.Blue, Name: "TP VP", Once: Price{ShipUpgrade: 1}, Use: false},
		TileItem{Category: TileColor, Type: TileColorGreen, Color: color.Green, Name: "P VP", Once: Price{Power: 1, Science: Science{Banking: 1, Law: 1, Engineering: 1, Medicine: 1}}, Use: false},
		TileItem{Category: TileColor, Type: TileColorGray, Color: color.Gray, Name: "SH/SA VP", Use: false},
	}
}

func GetColorTile(value color.Color) TileItem {
	pos := int(value) - 2

	log.Println("pos : ", pos)
	return _colorTiles[pos]
}
