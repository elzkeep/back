package resources

import (
	"aoi/models"
)

type RoundTile struct {
	Items            []TileItem `json:"items"`
	Reserved         []TileItem `json:"reserved"`
	Original         []TileItem `json:"-"`
	OriginalReserved []TileItem `json:"-"`
}

func NewRoundTile(id int64, typeid int) *RoundTile {
	var item RoundTile

	item.Items = make([]TileItem, 0)
	item.Reserved = make([]TileItem, 0)

	tiles := []TileItem{
		{Category: TileRound, Type: TileRoundEdgeVP, Name: "side VP", Build: BuildVP{River: 2}, Ship: 1, Use: false},
		{Category: TileRound, Type: TileRoundPristVP, Name: "P VP", Receive: Price{Prist: 1}, Build: BuildVP{Prist: 2}, Use: false},
		{Category: TileRound, Type: TileRoundTpVP, Name: "TP VP", Receive: Price{Power: 3}, Build: BuildVP{TP: 3}, Use: false},
		{Category: TileRound, Type: TileRoundShVP, Name: "SH/SA VP", Receive: Price{Worker: 1}, Pass: Price{ShVP: 4}, Use: false},
		{Category: TileRound, Type: TileRoundSpade, Name: "spd", Receive: Price{Book: Book{Any: 1}}, Action: Price{Spade: 1}, Use: false},
		{Category: TileRound, Type: TileRoundBridge, Name: "bridge", Receive: Price{Book: Book{Any: 1}}, Action: Price{Bridge: 1}, Use: false},
		{Category: TileRound, Type: TileRoundScienceCube, Name: "1 science", Receive: Price{Worker: 2}, Action: Price{Science: Science{Single: 1}}, Use: false},
		{Category: TileRound, Type: TileRoundSchoolScienceCoin, Name: "te science", Receive: Price{Coin: 4}, Pass: Price{Science: Science{Any: 1}}, Use: false},
		{Category: TileRound, Type: TileRoundPower, Name: "4PW", Receive: Price{Coin: 2, Power: 4}, Use: false},
		{Category: TileRound, Type: TileRoundCoin, Name: "6C", Receive: Price{Coin: 6}, Use: false},
	}

	conn := models.NewConnection()
	defer conn.Close()

	gametileManager := models.NewGametileManager(conn)
	items := gametileManager.Find([]interface{}{
		models.Where{Column: "game", Value: id, Compare: "="},
		models.Where{Column: "type", Value: int(TileRound), Compare: "="},
		models.Ordering("gt_order"),
	})

	if id <= 10 {
		for _, v := range items {
			for _, tile := range tiles {
				if v.Number == int(tile.Type) {
					item.Reserved = append(item.Reserved, tile)
				}
			}
		}

		for _, tile := range tiles {
			flag := false

			for _, v := range items {
				if v.Number == int(tile.Type) {
					flag = true
				}
			}

			if flag == true {
				item.Items = append(item.Items, tile)
			}
		}
	} else {
		if typeid != 1 {
			for _, v := range items {
				for _, tile := range tiles {
					if v.Number == int(tile.Type) {
						item.Items = append(item.Items, tile)
					}
				}
			}
		} else {
			pos := 0
			for _, v := range items {
				for _, tile := range tiles {
					if v.Number == int(tile.Type) {
						if pos < 7 {
							item.Reserved = append(item.Reserved, tile)
						} else {
							item.Items = append(item.Items, tile)
						}
						pos++
					}
				}
			}
		}
	}

	item.Original = make([]TileItem, len(item.Items))
	item.OriginalReserved = make([]TileItem, len(item.Reserved))
	copy(item.Original, item.Items)
	copy(item.OriginalReserved, item.Reserved)

	return &item
}

func (p *RoundTile) Copy() *RoundTile {
	var item RoundTile

	item.Items = make([]TileItem, len(p.Original))
	copy(item.Items, p.Original)
	for i := range item.Items {
		item.Items[i].Use = false
	}

	item.Reserved = make([]TileItem, len(p.OriginalReserved))
	copy(item.Reserved, p.OriginalReserved)
	for i := range item.Reserved {
		item.Reserved[i].Use = false
	}

	return &item
}

func (p *RoundTile) BuildStart() {
	p.Reserved = make([]TileItem, 0)
}

func (p *RoundTile) Start() {
	for i := range p.Items {
		p.Items[i].Use = false
		p.Items[i].Coin++
	}
}

func (p *RoundTile) Pass(pos int) TileItem {
	ret := p.Items[pos]

	p.Items = append(p.Items[:pos], p.Items[pos+1:]...)

	return ret
}

func (p *RoundTile) Add(tile TileItem) {
	p.Items = append(p.Items, tile)
}
