package factions

import (
	"aoi/game/action"
	"aoi/game/resources"
)

type Illusionists struct {
	Faction *Faction
}

func (p *Illusionists) Init(tile resources.TileItem) {
	p.Faction = NewFaction("Illusionists", "Illusionists", resources.GetFactionTile(resources.TileFactionIllusionists), tile)
}

func (p *Illusionists) Income() {
	p.Faction.Income()
}

func (p *Illusionists) FirstIncome() {
	p.Faction.FirstIncome()
}

func (p *Illusionists) Print() {
	p.Faction.Print()
}

func (p *Illusionists) GetInstance() *Faction {
	return p.Faction
}

func (p *Illusionists) AdvanceShip() error {
	return p.Faction.AdvanceShip()
}

func (p *Illusionists) AdvanceSpade() error {
	return p.Faction.AdvanceSpade()
}

func (p *Illusionists) FirstBuild(x int, y int) {
	p.Faction.FirstBuild(x, y)
}

func (p *Illusionists) Build(x int, y int, needSpade int, building resources.Building) error {
	return p.Faction.Build(x, y, needSpade, building)
}

func (p *Illusionists) Upgrade(x int, y int, target resources.Building) error {
	return p.Faction.Upgrade(x, y, target)
}

func (p *Illusionists) SendScholar() error {
	return p.Faction.SendScholar()
}

func (p *Illusionists) SupployScholar() error {
	return p.Faction.SupployScholar()
}

func (p *Illusionists) PowerAction(item action.PowerActionItem) error {
	return p.Faction.PowerAction(item)
}

func (p *Illusionists) Book(item action.BookActionItem, book resources.Book) error {
	return p.Faction.Book(item, book)
}

func (p *Illusionists) Bridge(x1 int, y1 int, x2 int, y2 int) error {
	return p.Faction.Bridge(x1, y1, x2, y2)
}

func (p *Illusionists) ReceiveCity(item resources.CityItem) error {
	return p.Faction.ReceiveCity(item)
}

func (p *Illusionists) Pass(tile resources.TileItem) (error, resources.TileItem) {
	return p.Faction.Pass(tile)
}

func (p *Illusionists) Dig(dig int) error {
	return p.Faction.Dig(dig)
}

func (p *Illusionists) ConvertDig(spade int) error {
	return p.Faction.ConvertDig(spade)
}

func (p *Illusionists) TurnEnd() error {
	return p.Faction.TurnEnd()
}

func (p *Illusionists) PalaceTile(tile resources.TileItem) error {
	return p.Faction.PalaceTile(tile)
}

func (p *Illusionists) SchoolTile(tile resources.TileItem) error {
	return p.Faction.SchoolTile(tile)
}

func (p *Illusionists) TileAction(category resources.TileCategory, pos int) error {
	return p.Faction.TileAction(category, pos)
}
