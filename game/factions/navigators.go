package factions

import (
	"aoi/game/action"
	"aoi/game/resources"
)

type Navigators struct {
	Faction *Faction
}

func (p *Navigators) Init(tile resources.TileItem) {
	p.Faction = NewFaction("Navigators", "Navigators", resources.GetFactionTile(resources.TileFactionNavigators), tile)
}

func (p *Navigators) Income() {
	p.Faction.Income()
}

func (p *Navigators) FirstIncome() {
	p.Faction.FirstIncome()
}

func (p *Navigators) Print() {
	p.Faction.Print()
}

func (p *Navigators) GetInstance() *Faction {
	return p.Faction
}

func (p *Navigators) AdvanceShip() error {
	return p.Faction.AdvanceShip()
}

func (p *Navigators) AdvanceSpade() error {
	return p.Faction.AdvanceSpade()
}

func (p *Navigators) FirstBuild(x int, y int) {
	p.Faction.FirstBuild(x, y)
}

func (p *Navigators) Build(x int, y int, needSpade int) error {
	return p.Faction.Build(x, y, needSpade)
}

func (p *Navigators) Upgrade(x int, y int, target resources.Building) error {
	return p.Faction.Upgrade(x, y, target)
}

func (p *Navigators) SendScholar() error {
	return p.Faction.SendScholar()
}

func (p *Navigators) SupployScholar() error {
	return p.Faction.SupployScholar()
}

func (p *Navigators) PowerAction(item action.PowerActionItem) error {
	return p.Faction.PowerAction(item)
}

func (p *Navigators) Book(item action.BookActionItem) error {
	return p.Faction.Book(item)
}

func (p *Navigators) Bridge(x1 int, y1 int, x2 int, y2 int) error {
	return p.Faction.Bridge(x1, y1, x2, y2)
}

func (p *Navigators) PassIncome() {
	p.Faction.PassIncome()
}

func (p *Navigators) ReceiveCity(item resources.CityItem) error {
	return p.Faction.ReceiveCity(item)
}

func (p *Navigators) Pass(tile resources.TileItem) (error, resources.TileItem) {
	return p.Faction.Pass(tile)
}

func (p *Navigators) Dig(dig int) error {
	return p.Faction.Dig(dig)
}

func (p *Navigators) ConvertDig(spade int) error {
	return p.Faction.ConvertDig(spade)
}

func (p *Navigators) TurnEnd() error {
	return p.Faction.TurnEnd()
}

func (p *Navigators) PalaceTile(tile resources.TileItem) error {
	return p.Faction.PalaceTile(tile)
}

func (p *Navigators) SchoolTile(tile resources.TileItem) error {
	return p.Faction.SchoolTile(tile)
}

func (p *Navigators) TileAction(category resources.TileCategory, pos int) error {
	return p.Faction.TileAction(category, pos)
}
