package resources

import (
	"aoi/models"
)

type FactionTile struct {
	Items    []TileItem `json:"items"`
	Original []TileItem `json:"-"`
}

var _factionTile []TileItem

func NewFactionTile(id int64) *FactionTile {
	var item FactionTile

	item.Items = make([]TileItem, 0)

	tiles := []TileItem{
		{Category: TileFaction, Type: TileFactionBlessed, Name: "Blessed", Once: Price{Science: Science{Banking: 1, Law: 1, Engineering: 1, Medicine: 1}}, Use: false},
		{Category: TileFaction, Type: TileFactionFelines, Name: "Felines", Once: Price{Science: Science{Banking: 1, Medicine: 1}}, Use: false},
		{Category: TileFaction, Type: TileFactionGoblins, Name: "Goblins", Once: Price{Science: Science{Banking: 1, Engineering: 1}, Worker: 1}, Use: false},
		{Category: TileFaction, Type: TileFactionIllusionists, Name: "Illusionists", Once: Price{Science: Science{Medicine: 2}}, Use: false},
		{Category: TileFaction, Type: TileFactionInventors, Name: "Inventors", Once: Price{Tile: 1}, Use: false},
		{Category: TileFaction, Type: TileFactionLizards, Name: "Lizards", Once: Price{Science: Science{Any: 2}}, Use: false},
		{Category: TileFaction, Type: TileFactionMoles, Name: "Moles", Once: Price{Science: Science{Engineering: 2}}, Action: Price{Bridge: 1}, Use: false},
		{Category: TileFaction, Type: TileFactionMonks, Name: "Monks", Once: Price{Science: Science{Law: 1}}, Use: false},
		{Category: TileFaction, Type: TileFactionNavigators, Name: "Navigators", Once: Price{Science: Science{Law: 3}}, Build: BuildVP{River: 2}, Use: false},
		{Category: TileFaction, Type: TileFactionOmar, Name: "Omar", Once: Price{Science: Science{Banking: 1, Engineering: 1}, Building: WHITE_TOWER}, Receive: Price{Coin: 2, Power: 2}, Use: false},
		{Category: TileFaction, Type: TileFactionPhilosophers, Name: "Philosophers", Once: Price{Science: Science{Banking: 2}}, Action: Price{Book: Book{Any: 1}}, Use: false},
		{Category: TileFaction, Type: TileFactionPsychics, Name: "Psychics", Once: Price{Science: Science{Banking: 1, Medicine: 1}, Worker: 1}, Action: Price{Power: 5}, Use: false},
	}

	conn := models.NewConnection()
	defer conn.Close()

	gametileManager := models.NewGametileManager(conn)
	items := gametileManager.Find([]interface{}{
		models.Where{Column: "game", Value: id, Compare: "="},
		models.Where{Column: "type", Value: int(TileFaction), Compare: "="},
		models.Ordering("gt_order"),
	})

	for _, v := range items {
		for _, tile := range tiles {
			if v.Number == int(tile.Type) {
				item.Items = append(item.Items, tile)
			}
		}
	}

	_factionTile = make([]TileItem, len(item.Items))
	copy(_factionTile, item.Items)

	item.Original = make([]TileItem, len(item.Items))
	copy(item.Original, item.Items)

	return &item
}

func (p *FactionTile) Copy() *FactionTile {
	var item FactionTile

	_factionTile = make([]TileItem, len(p.Original))
	copy(_factionTile, p.Original)
	for i := range _factionTile {
		_factionTile[i].Use = false
	}

	item.Items = make([]TileItem, len(p.Original))
	copy(item.Items, p.Original)
	for i := range item.Items {
		item.Items[i].Use = false
	}

	return &item
}

func GetFactionTile(value TileType) TileItem {
	for _, v := range _factionTile {
		if v.Type == value {
			return v
		}
	}

	return TileItem{}
}
