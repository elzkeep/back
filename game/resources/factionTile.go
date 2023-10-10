package resources

var _factionTiles []TileItem

func init() {
	_factionTiles = []TileItem{
		TileItem{Category: TileFaction, Type: TileFactionBlessed, Name: "side VP", Once: Price{Science: Science{Banking: 1, Law: 1, Engineering: 1, Medicine: 1}}, Use: false},
		TileItem{Category: TileFaction, Type: TileFactionFelines, Name: "P VP", Once: Price{Science: Science{Banking: 1, Medicine: 1}}, Use: false},
		TileItem{Category: TileFaction, Type: TileFactionGoblins, Name: "TP VP", Once: Price{Science: Science{Banking: 1, Engineering: 1}, Worker: 1}, Use: false},
		TileItem{Category: TileFaction, Type: TileFactionIllusionists, Name: "SH/SA VP", Once: Price{Science: Science{Medicine: 2}}, Use: false},
		TileItem{Category: TileFaction, Type: TileFactionInventors, Name: "spd", Once: Price{Tile: 1}, Use: false},
		TileItem{Category: TileFaction, Type: TileFactionLizards, Name: "bridge", Once: Price{Science: Science{Any: 2}}, Use: false},
		TileItem{Category: TileFaction, Type: TileFactionMoles, Name: "1 science", Once: Price{Science: Science{Engineering: 2}}, Use: false},
		TileItem{Category: TileFaction, Type: TileFactionMonks, Name: "te science", Once: Price{Science: Science{Law: 1}}, Use: false},
		TileItem{Category: TileFaction, Type: TileFactionNavigators, Name: "4PW", Once: Price{Science: Science{Law: 3}}, Build: BuildVP{River: 2}, Use: false},
		TileItem{Category: TileFaction, Type: TileFactionOmar, Name: "6C", Once: Price{Science: Science{Banking: 1, Engineering: 1}}, Receive: Price{Coin: 2, Power: 2}, Use: false},
		TileItem{Category: TileFaction, Type: TileFactionPhilosophers, Name: "6C", Once: Price{Science: Science{Banking: 2}}, Action: Price{Book: 1}, Use: false},
		TileItem{Category: TileFaction, Type: TileFactionPsychics, Name: "6C", Once: Price{Science: Science{Banking: 1, Medicine: 1}, Worker: 1}, Action: Price{Power: 5}, Use: false},
	}

	/*
		rand.Shuffle(len(items), func(i, j int) { items[i], items[j] = items[j], items[i] })

		for i := 0; i < count+3; i++ {
			p.Items = append(p.Items, items[i])
		}
	*/
}

func GetFactionTile(value TileType) TileItem {
	pos := int(value) - int(TileSchoolPassPrist)

	return _factionTiles[pos]
}
