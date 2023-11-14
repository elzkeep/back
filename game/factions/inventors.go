package factions

import (
	"aoi/game/action"
	. "aoi/game/resources"
)

type Inventors struct {
	Faction
}

func (p *Inventors) Init(tile TileItem) {
	p.InitFaction("Inventors", "Inventors", GetFactionTile(TileFactionInventors), tile)
}

func (p *Inventors) GetInstance() *Faction {
	return &p.Faction
}

func (p *Inventors) Print() {
	p.Faction.Print()
}

func (p *Inventors) FirstIncome() {
	p.Faction.FirstIncome()
}

func (p *Inventors) Income() {
	p.Faction.Income()
}

func (p *Inventors) GetScience(pos int) int {
	return p.Faction.GetScience(pos)
}

func (p *Inventors) FirstBuild(x int, y int, building Building) error {
	return p.Faction.FirstBuild(x, y, building)
}

func (p *Inventors) Build(x int, y int, needSpade int, building Building) error {
	return p.Faction.Build(x, y, needSpade, building)
}

func (p *Inventors) Upgrade(x int, y int, target Building, extra int) error {
	return p.Faction.Upgrade(x, y, target, extra)
}

func (p *Inventors) AdvanceShip() error {
	return p.Faction.AdvanceShip()
}

func (p *Inventors) AdvanceSpade() error {
	return p.Faction.AdvanceSpade()
}

func (p *Inventors) SendScholar(pos int, inc int) error {
	return p.Faction.SendScholar(pos, inc)
}

func (p *Inventors) SupployScholar(pos int, inc int) error {
	return p.Faction.SupployScholar(pos, inc)
}

func (p *Inventors) PowerAction(item action.PowerActionItem) error {
	return p.Faction.PowerAction(item)
}

func (p *Inventors) Book(item action.BookActionItem, book Book) error {
	return p.Faction.Book(item, book)
}

func (p *Inventors) Bridge(x1 int, y1 int, x2 int, y2 int) error {
	return p.Faction.Bridge(x1, y1, x2, y2)
}

func (p *Inventors) Pass(tile TileItem) (error, TileItem) {
	return p.Faction.Pass(tile)
}

func (p *Inventors) ReceiveCity(item CityItem) error {
	return p.Faction.ReceiveCity(item)
}

func (p *Inventors) Dig(x int, y int, dig int) error {
	return p.Faction.Dig(x, y, dig)
}

func (p *Inventors) TurnEnd(round int) error {
	return p.Faction.TurnEnd(round)
}

func (p *Inventors) PalaceTile(tile TileItem) error {
	return p.Faction.PalaceTile(tile)
}

func (p *Inventors) SchoolTile(tile TileItem, science int) error {
	return p.Faction.SchoolTile(tile, science)
}

func (p *Inventors) RoundTile(tile TileItem) error {
	return p.Faction.RoundTile(tile)
}

func (p *Inventors) TileAction(category TileCategory, pos int) error {
	return p.Faction.TileAction(category, pos)
}
