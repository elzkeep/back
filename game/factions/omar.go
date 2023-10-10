package factions

import (
	"aoi/game/action"
	"aoi/game/resources"
)

type Omar struct {
	Faction *Faction
}

func (p *Omar) Init(tile resources.TileItem) {
	p.Faction = NewFaction("Omar", "Omar", resources.GetFactionTile(resources.TileFactionOmar), tile)
}

func (p *Omar) Income() {
	p.Faction.Income()
}

func (p *Omar) FirstIncome() {
	p.Faction.FirstIncome()
}

func (p *Omar) Print() {
	p.Faction.Print()
}

func (p *Omar) GetInstance() *Faction {
	return p.Faction
}

func (p *Omar) AdvanceShip() error {
	return p.Faction.AdvanceShip()
}

func (p *Omar) AdvanceSpade() error {
	return p.Faction.AdvanceSpade()
}

func (p *Omar) FirstBuild(x int, y int) {
	p.Faction.FirstBuild(x, y)
}

func (p *Omar) Build(x int, y int, needSpade int) error {
	return p.Faction.Build(x, y, needSpade)
}

func (p *Omar) Upgrade(x int, y int, target resources.Building) error {
	return p.Faction.Upgrade(x, y, target)
}

func (p *Omar) SendScholar() error {
	return p.Faction.SendScholar()
}

func (p *Omar) SupployScholar() error {
	return p.Faction.SupployScholar()
}

func (p *Omar) PowerAction(item action.PowerActionItem) error {
	return p.Faction.PowerAction(item)
}

func (p *Omar) Book(item action.BookActionItem) error {
	return p.Faction.Book(item)
}

func (p *Omar) Bridge(x1 int, y1 int, x2 int, y2 int) error {
	return p.Faction.Bridge(x1, y1, x2, y2)
}

func (p *Omar) PassIncome() {
	p.Faction.PassIncome()
}

func (p *Omar) ReceiveCity(item resources.CityItem) error {
	return p.Faction.ReceiveCity(item)
}

func (p *Omar) Pass(tile resources.TileItem) (error, resources.TileItem) {
	return p.Faction.Pass(tile)
}

func (p *Omar) Dig(dig int) error {
	return p.Faction.Dig(dig)
}

func (p *Omar) ConvertDig(spade int) error {
	return p.Faction.ConvertDig(spade)
}

func (p *Omar) TurnEnd() error {
	return p.Faction.TurnEnd()
}

func (p *Omar) PalaceTile(tile resources.TileItem) error {
	return p.Faction.PalaceTile(tile)
}

func (p *Omar) SchoolTile(tile resources.TileItem) error {
	return p.Faction.SchoolTile(tile)
}

func (p *Omar) TileAction(category resources.TileCategory, pos int) error {
	return p.Faction.TileAction(category, pos)
}
