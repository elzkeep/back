package factions

import (
	"aoi/game/action"
	. "aoi/game/resources"
)

type Goblins struct {
	Faction
}

func (p *Goblins) Init(tile TileItem) {
	p.InitFaction("Goblins", "Goblins", GetFactionTile(TileFactionGoblins), tile)
}

func (p *Goblins) GetInstance() *Faction {
	return &p.Faction
}

func (p *Goblins) Print() {
	p.Faction.Print()
}

func (p *Goblins) FirstIncome() {
	p.Faction.FirstIncome()
}

func (p *Goblins) Income() {
	p.Faction.Income()
}

func (p *Goblins) GetScience(pos int) int {
	return p.Faction.GetScience(pos)
}

func (p *Goblins) FirstBuild(x int, y int) error {
	return p.Faction.FirstBuild(x, y)
}

func (p *Goblins) Build(x int, y int, needSpade int, building Building) error {
	return p.Faction.Build(x, y, needSpade, building)
}

func (p *Goblins) Upgrade(x int, y int, target Building, extra int) error {
	return p.Faction.Upgrade(x, y, target, extra)
}

func (p *Goblins) AdvanceShip() error {
	return p.Faction.AdvanceShip()
}

func (p *Goblins) AdvanceSpade() error {
	return p.Faction.AdvanceSpade()
}

func (p *Goblins) SendScholar(pos int, inc int) error {
	return p.Faction.SendScholar(pos, inc)
}

func (p *Goblins) SupployScholar(pos int, inc int) error {
	return p.Faction.SupployScholar(pos, inc)
}

func (p *Goblins) PowerAction(item action.PowerActionItem) error {
	return p.Faction.PowerAction(item)
}

func (p *Goblins) Book(item action.BookActionItem, book Book) error {
	return p.Faction.Book(item, book)
}

func (p *Goblins) Bridge(x1 int, y1 int, x2 int, y2 int) error {
	return p.Faction.Bridge(x1, y1, x2, y2)
}

func (p *Goblins) Pass(tile TileItem) (error, TileItem) {
	return p.Faction.Pass(tile)
}

func (p *Goblins) ReceiveCity(item CityItem) error {
	return p.Faction.ReceiveCity(item)
}

func (p *Goblins) Dig(x int, y int, dig int) error {
	err := p.Faction.Dig(x, y, dig)
	if err == nil {
		p.ReceiveResource(Price{Coin: dig * 2})
	}

	return err
}

func (p *Goblins) TurnEnd(round int) error {
	return p.Faction.TurnEnd(round)
}

func (p *Goblins) PalaceTile(tile TileItem) error {
	return p.Faction.PalaceTile(tile)
}

func (p *Goblins) SchoolTile(tile TileItem, science int) error {
	return p.Faction.SchoolTile(tile, science)
}

func (p *Goblins) RoundTile(tile TileItem) error {
	return p.Faction.RoundTile(tile)
}

func (p *Goblins) TileAction(category TileCategory, pos int) error {
	return p.Faction.TileAction(category, pos)
}
