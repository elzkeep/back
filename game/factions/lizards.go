package factions

import (
	"aoi/game/action"
	"aoi/game/resources"
)

type Lizards struct {
	Faction *Faction
}

func (p *Lizards) Init(tile resources.TileItem) {
	p.Faction = NewFaction("Lizards", "Lizards", resources.GetFactionTile(resources.TileFactionLizards), tile)
}

func (p *Lizards) Income() {
	p.Faction.Income()
}

func (p *Lizards) FirstIncome() {
	p.Faction.FirstIncome()
}

func (p *Lizards) Print() {
	p.Faction.Print()
}

func (p *Lizards) GetInstance() *Faction {
	return p.Faction
}

func (p *Lizards) AdvanceShip() error {
	return p.Faction.AdvanceShip()
}

func (p *Lizards) AdvanceSpade() error {
	return p.Faction.AdvanceSpade()
}

func (p *Lizards) FirstBuild(x int, y int) {
	p.Faction.FirstBuild(x, y)
}

func (p *Lizards) Build(x int, y int, needSpade int, building resources.Building) error {
	return p.Faction.Build(x, y, needSpade, building)
}

func (p *Lizards) Upgrade(x int, y int, target resources.Building) error {
	return p.Faction.Upgrade(x, y, target)
}

func (p *Lizards) SendScholar() error {
	return p.Faction.SendScholar()
}

func (p *Lizards) SupployScholar() error {
	return p.Faction.SupployScholar()
}

func (p *Lizards) PowerAction(item action.PowerActionItem) error {
	return p.Faction.PowerAction(item)
}

func (p *Lizards) Book(item action.BookActionItem) error {
	return p.Faction.Book(item)
}

func (p *Lizards) Bridge(x1 int, y1 int, x2 int, y2 int) error {
	return p.Faction.Bridge(x1, y1, x2, y2)
}

func (p *Lizards) PassIncome() {
	p.Faction.PassIncome()
}

func (p *Lizards) ReceiveCity(item resources.CityItem) error {
	return p.Faction.ReceiveCity(item)
}

func (p *Lizards) Pass(tile resources.TileItem) (error, resources.TileItem) {
	return p.Faction.Pass(tile)
}

func (p *Lizards) Dig(dig int) error {
	return p.Faction.Dig(dig)
}

func (p *Lizards) ConvertDig(spade int) error {
	return p.Faction.ConvertDig(spade)
}

func (p *Lizards) TurnEnd() error {
	return p.Faction.TurnEnd()
}

func (p *Lizards) PalaceTile(tile resources.TileItem) error {
	return p.Faction.PalaceTile(tile)
}

func (p *Lizards) SchoolTile(tile resources.TileItem) error {
	return p.Faction.SchoolTile(tile)
}

func (p *Lizards) TileAction(category resources.TileCategory, pos int) error {
	return p.Faction.TileAction(category, pos)
}
