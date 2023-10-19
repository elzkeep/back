package resources

var _factionTiles []TileItem

func init() {
	_factionTiles = []TileItem{
		TileItem{Category: TileFaction, Type: TileFactionBlessed, Name: "Blessed", Once: Price{Science: Science{Banking: 1, Law: 1, Engineering: 1, Medicine: 1}}, Use: false},
		TileItem{Category: TileFaction, Type: TileFactionFelines, Name: "Felines", Once: Price{Science: Science{Banking: 1, Medicine: 1}}, Use: false},
		TileItem{Category: TileFaction, Type: TileFactionGoblins, Name: "Goblins", Once: Price{Science: Science{Banking: 1, Engineering: 1}, Worker: 1}, Use: false},
		TileItem{Category: TileFaction, Type: TileFactionIllusionists, Name: "Illusionists", Once: Price{Science: Science{Medicine: 2}}, Use: false},
		TileItem{Category: TileFaction, Type: TileFactionInventors, Name: "Inventors", Once: Price{Tile: 1}, Use: false},
		TileItem{Category: TileFaction, Type: TileFactionLizards, Name: "Lizards", Once: Price{Science: Science{Any: 2}}, Use: false},
		TileItem{Category: TileFaction, Type: TileFactionMoles, Name: "Moles", Once: Price{Science: Science{Engineering: 2}}, Use: false},
		TileItem{Category: TileFaction, Type: TileFactionMonks, Name: "Monks", Once: Price{Science: Science{Law: 1}}, Use: false},
		TileItem{Category: TileFaction, Type: TileFactionNavigators, Name: "Navigators", Once: Price{Science: Science{Law: 3}}, Build: BuildVP{River: 2}, Use: false},
		TileItem{Category: TileFaction, Type: TileFactionOmar, Name: "Omar", Once: Price{Science: Science{Banking: 1, Engineering: 1}, Building: WHITE_TOWER}, Receive: Price{Coin: 2, Power: 2}, Use: false},
		TileItem{Category: TileFaction, Type: TileFactionPhilosophers, Name: "Philosophers", Once: Price{Science: Science{Banking: 2}}, Action: Price{Book: Book{Any: 1}}, Use: false},
		TileItem{Category: TileFaction, Type: TileFactionPsychics, Name: "Psychics", Once: Price{Science: Science{Banking: 1, Medicine: 1}, Worker: 1}, Action: Price{Power: 5}, Use: false},
	}

	/*
		rand.Shuffle(len(items), func(i, j int) { items[i], items[j] = items[j], items[i] })

		for i := 0; i < count+3; i++ {
			p.Items = append(p.Items, items[i])
		}
	*/
}

func GetFactionTile(value TileType) TileItem {
	pos := int(value) - int(TileSchoolPassPrist) - 1

	return _factionTiles[pos]
}
