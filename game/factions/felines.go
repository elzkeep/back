package factions

import (
	"aoi/game/action"
	"aoi/game/resources"
)

type Felines struct {
	Faction *Faction
}

func (p *Felines) Init(tile resources.TileItem) {
	p.Faction = NewFaction("Felines", "Felines", resources.GetFactionTile(resources.TileFactionFelines), tile)
}

func (p *Felines) Income() {
	p.Faction.Income()
}

func (p *Felines) FirstIncome() {
	p.Faction.FirstIncome()
}

func (p *Felines) Print() {
	p.Faction.Print()
}

func (p *Felines) GetInstance() *Faction {
	return p.Faction
}

func (p *Felines) AdvanceShip() error {
	return p.Faction.AdvanceShip()
}

func (p *Felines) AdvanceSpade() error {
	return p.Faction.AdvanceSpade()
}

func (p *Felines) FirstBuild(x int, y int) {
	p.Faction.FirstBuild(x, y)
}

func (p *Felines) Build(x int, y int, needSpade int, building resources.Building) error {
	return p.Faction.Build(x, y, needSpade, building)
}

func (p *Felines) Upgrade(x int, y int, target resources.Building) error {
	return p.Faction.Upgrade(x, y, target)
}

func (p *Felines) SendScholar() error {
	return p.Faction.SendScholar()
}

func (p *Felines) SupployScholar() error {
	return p.Faction.SupployScholar()
}

func (p *Felines) PowerAction(item action.PowerActionItem) error {
	return p.Faction.PowerAction(item)
}

func (p *Felines) Book(item action.BookActionItem) error {
	return p.Faction.Book(item)
}

func (p *Felines) Bridge(x1 int, y1 int, x2 int, y2 int) error {
	return p.Faction.Bridge(x1, y1, x2, y2)
}

func (p *Felines) PassIncome() {
	p.Faction.PassIncome()
}

func (p *Felines) ReceiveCity(item resources.CityItem) error {
	return p.Faction.ReceiveCity(item)
}

func (p *Felines) Pass(tile resources.TileItem) (error, resources.TileItem) {
	return p.Faction.Pass(tile)
}

func (p *Felines) Dig(dig int) error {
	return p.Faction.Dig(dig)
}

func (p *Felines) ConvertDig(spade int) error {
	return p.Faction.ConvertDig(spade)
}

func (p *Felines) TurnEnd() error {
	return p.Faction.TurnEnd()
}

func (p *Felines) PalaceTile(tile resources.TileItem) error {
	return p.Faction.PalaceTile(tile)
}

func (p *Felines) SchoolTile(tile resources.TileItem) error {
	return p.Faction.SchoolTile(tile)
}

func (p *Felines) TileAction(category resources.TileCategory, pos int) error {
	return p.Faction.TileAction(category, pos)
}
