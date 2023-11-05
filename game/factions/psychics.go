package factions

import (
	"aoi/game/action"
	. "aoi/game/resources"
	"errors"
)

type Psychics struct {
	Faction
}

func (p *Psychics) Init(tile TileItem) {
	p.InitFaction("Psychics", "Psychics", GetFactionTile(TileFactionPsychics), tile)
}

func (p *Psychics) GetInstance() *Faction {
	return &p.Faction
}

func (p *Psychics) Print() {
	p.Faction.Print()
}

func (p *Psychics) FirstIncome() {
}

func (p *Psychics) Income() {
}

func (p *Psychics) GetScience(pos int) int {
	return p.Faction.GetScience(pos)
}

func (p *Psychics) FirstBuild(x int, y int) error {
	return p.Faction.FirstBuild(x, y)
}

func (p *Psychics) Build(x int, y int, needSpade int, building Building) error {
	return p.Faction.Build(x, y, needSpade, building)
}

func (p *Psychics) Upgrade(x int, y int, target Building) error {
	return p.Faction.Upgrade(x, y, target)
}

func (p *Psychics) AdvanceShip() error {
	return p.Faction.AdvanceShip()
}

func (p *Psychics) AdvanceSpade() error {
	return p.Faction.AdvanceSpade()
}

func (p *Psychics) SendScholar() error {
	return p.Faction.SendScholar()
}

func (p *Psychics) SupployScholar() error {
	return p.Faction.SupployScholar()
}

func (p *Psychics) PowerAction(item action.PowerActionItem) error {
	return p.Faction.PowerAction(item)
}

func (p *Psychics) Book(item action.BookActionItem, book Book) error {
	return p.Faction.Book(item, book)
}

func (p *Psychics) Bridge(x1 int, y1 int, x2 int, y2 int) error {
	return p.Faction.Bridge(x1, y1, x2, y2)
}

func (p *Psychics) Pass(tile TileItem) (error, TileItem) {
	return p.Faction.Pass(tile)
}

func (p *Psychics) ReceiveCity(item CityItem) error {
	return p.Faction.ReceiveCity(item)
}

func (p *Psychics) Dig(dig int) error {
	return p.Faction.Dig(dig)
}

func (p *Psychics) TurnEnd() error {
	return p.Faction.TurnEnd()
}

func (p *Psychics) PalaceTile(tile TileItem) error {
	return p.Faction.PalaceTile(tile)
}

func (p *Psychics) SchoolTile(tile TileItem, science int) error {
	return p.Faction.SchoolTile(tile, science)
}

func (p *Psychics) RoundTile(tile TileItem) error {
	return p.Faction.RoundTile(tile)
}

func (p *Psychics) TileAction(category TileCategory, pos int) error {
	err := p.Faction.TileAction(category, pos)

	if err != nil {
		return err
	}

	var tile *TileItem

	find := -1
	for i, v := range p.Tiles {
		if v.Category == category && v.Type == TileType(pos) {
			find = i
			break
		}
	}

	if find == -1 {
		return errors.New("not found")
	}

	tile = &p.Tiles[find]

	if tile.Type == TileFactionPsychics {
		p.Action = false
	}

	return err
}
