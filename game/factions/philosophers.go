package factions

import (
	"aoi/game/action"
	"aoi/game/resources"
)

type Philosophers struct {
	Faction *Faction
}

func (p *Philosophers) Init(tile resources.TileItem) {
	p.Faction = NewFaction("Philosophers", "Philosophers", resources.GetFactionTile(resources.TileFactionPhilosophers), tile)
}

func (p *Philosophers) Income() {
	p.Faction.Income()
}

func (p *Philosophers) FirstIncome() {
	p.Faction.FirstIncome()
}

func (p *Philosophers) Print() {
	p.Faction.Print()
}

func (p *Philosophers) GetInstance() *Faction {
	return p.Faction
}

func (p *Philosophers) AdvanceShip() error {
	return p.Faction.AdvanceShip()
}

func (p *Philosophers) AdvanceSpade() error {
	return p.Faction.AdvanceSpade()
}

func (p *Philosophers) FirstBuild(x int, y int) {
	p.Faction.FirstBuild(x, y)
}

func (p *Philosophers) Build(x int, y int, needSpade int, building resources.Building) error {
	return p.Faction.Build(x, y, needSpade, building)
}

func (p *Philosophers) Upgrade(x int, y int, target resources.Building) error {
	return p.Faction.Upgrade(x, y, target)
}

func (p *Philosophers) SendScholar() error {
	return p.Faction.SendScholar()
}

func (p *Philosophers) SupployScholar() error {
	return p.Faction.SupployScholar()
}

func (p *Philosophers) PowerAction(item action.PowerActionItem) error {
	return p.Faction.PowerAction(item)
}

func (p *Philosophers) Book(item action.BookActionItem, book resources.Book) error {
	return p.Faction.Book(item, book)
}

func (p *Philosophers) Bridge(x1 int, y1 int, x2 int, y2 int) error {
	return p.Faction.Bridge(x1, y1, x2, y2)
}

func (p *Philosophers) ReceiveCity(item resources.CityItem) error {
	return p.Faction.ReceiveCity(item)
}

func (p *Philosophers) Pass(tile resources.TileItem) (error, resources.TileItem) {
	return p.Faction.Pass(tile)
}

func (p *Philosophers) Dig(dig int) error {
	return p.Faction.Dig(dig)
}

func (p *Philosophers) ConvertDig(spade int) error {
	return p.Faction.ConvertDig(spade)
}

func (p *Philosophers) TurnEnd() error {
	return p.Faction.TurnEnd()
}

func (p *Philosophers) PalaceTile(tile resources.TileItem) error {
	return p.Faction.PalaceTile(tile)
}

func (p *Philosophers) SchoolTile(tile resources.TileItem) error {
	return p.Faction.SchoolTile(tile)
}

func (p *Philosophers) TileAction(category resources.TileCategory, pos int) error {
	return p.Faction.TileAction(category, pos)
}
