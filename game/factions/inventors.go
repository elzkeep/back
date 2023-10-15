package factions

import (
	"aoi/game/action"
	"aoi/game/resources"
)

type Inventors struct {
	Faction *Faction
}

func (p *Inventors) Init(tile resources.TileItem) {
	p.Faction = NewFaction("Inventors", "Inventors", resources.GetFactionTile(resources.TileFactionInventors), tile)
}

func (p *Inventors) Income() {
	p.Faction.Income()
}

func (p *Inventors) FirstIncome() {
	p.Faction.FirstIncome()
}

func (p *Inventors) Print() {
	p.Faction.Print()
}

func (p *Inventors) GetInstance() *Faction {
	return p.Faction
}

func (p *Inventors) AdvanceShip() error {
	return p.Faction.AdvanceShip()
}

func (p *Inventors) AdvanceSpade() error {
	return p.Faction.AdvanceSpade()
}

func (p *Inventors) FirstBuild(x int, y int) {
	p.Faction.FirstBuild(x, y)
}

func (p *Inventors) Build(x int, y int, needSpade int, building resources.Building) error {
	return p.Faction.Build(x, y, needSpade, building)
}

func (p *Inventors) Upgrade(x int, y int, target resources.Building) error {
	return p.Faction.Upgrade(x, y, target)
}

func (p *Inventors) SendScholar() error {
	return p.Faction.SendScholar()
}

func (p *Inventors) SupployScholar() error {
	return p.Faction.SupployScholar()
}

func (p *Inventors) PowerAction(item action.PowerActionItem) error {
	return p.Faction.PowerAction(item)
}

func (p *Inventors) Book(item action.BookActionItem) error {
	return p.Faction.Book(item)
}

func (p *Inventors) Bridge(x1 int, y1 int, x2 int, y2 int) error {
	return p.Faction.Bridge(x1, y1, x2, y2)
}

func (p *Inventors) PassIncome() {
	p.Faction.PassIncome()
}

func (p *Inventors) ReceiveCity(item resources.CityItem) error {
	return p.Faction.ReceiveCity(item)
}

func (p *Inventors) Pass(tile resources.TileItem) (error, resources.TileItem) {
	return p.Faction.Pass(tile)
}

func (p *Inventors) Dig(dig int) error {
	return p.Faction.Dig(dig)
}

func (p *Inventors) ConvertDig(spade int) error {
	return p.Faction.ConvertDig(spade)
}

func (p *Inventors) TurnEnd() error {
	return p.Faction.TurnEnd()
}

func (p *Inventors) PalaceTile(tile resources.TileItem) error {
	return p.Faction.PalaceTile(tile)
}

func (p *Inventors) SchoolTile(tile resources.TileItem) error {
	return p.Faction.SchoolTile(tile)
}

func (p *Inventors) TileAction(category resources.TileCategory, pos int) error {
	return p.Faction.TileAction(category, pos)
}
