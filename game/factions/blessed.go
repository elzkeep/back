package factions

import (
	"aoi/game/action"
	"aoi/game/resources"
)

type Blessed struct {
	Faction *Faction
}

func (p *Blessed) Init(tile resources.TileItem) {
	p.Faction = NewFaction("Blessed", "Blessed", resources.GetFactionTile(resources.TileFactionBlessed), tile)
}

func (p *Blessed) Income() {
	p.Faction.Income()
}

func (p *Blessed) FirstIncome() {
	p.Faction.FirstIncome()
}

func (p *Blessed) Print() {
	p.Faction.Print()
}

func (p *Blessed) GetInstance() *Faction {
	return p.Faction
}

func (p *Blessed) AdvanceShip() error {
	return p.Faction.AdvanceShip()
}

func (p *Blessed) AdvanceSpade() error {
	return p.Faction.AdvanceSpade()
}

func (p *Blessed) FirstBuild(x int, y int) {
	p.Faction.FirstBuild(x, y)
}

func (p *Blessed) Build(x int, y int, needSpade int) error {
	return p.Faction.Build(x, y, needSpade)
}

func (p *Blessed) Upgrade(x int, y int, target resources.Building) error {
	return p.Faction.Upgrade(x, y, target)
}

func (p *Blessed) SendScholar() error {
	return p.Faction.SendScholar()
}

func (p *Blessed) SupployScholar() error {
	return p.Faction.SupployScholar()
}

func (p *Blessed) PowerAction(item action.PowerActionItem) error {
	return p.Faction.PowerAction(item)
}

func (p *Blessed) Book(item action.BookActionItem) error {
	return p.Faction.Book(item)
}

func (p *Blessed) Bridge(x1 int, y1 int, x2 int, y2 int) error {
	return p.Faction.Bridge(x1, y1, x2, y2)
}

func (p *Blessed) PassIncome() {
	p.Faction.PassIncome()
}

func (p *Blessed) ReceiveCity(item resources.CityItem) error {
	return p.Faction.ReceiveCity(item)
}

func (p *Blessed) Pass(tile resources.TileItem) (error, resources.TileItem) {
	return p.Faction.Pass(tile)
}

func (p *Blessed) Dig(dig int) error {
	return p.Faction.Dig(dig)
}

func (p *Blessed) ConvertDig(spade int) error {
	return p.Faction.ConvertDig(spade)
}

func (p *Blessed) TurnEnd() error {
	return p.Faction.TurnEnd()
}

func (p *Blessed) PalaceTile(tile resources.TileItem) error {
	return p.Faction.PalaceTile(tile)
}

func (p *Blessed) SchoolTile(tile resources.TileItem) error {
	return p.Faction.SchoolTile(tile)
}

func (p *Blessed) TileAction(category resources.TileCategory, pos int) error {
	return p.Faction.TileAction(category, pos)
}
