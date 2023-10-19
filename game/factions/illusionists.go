package factions

import (
	"aoi/game/action"
	. "aoi/game/resources"
)

type Illusionists struct {
	Faction
}

func (p *Illusionists) Init(tile TileItem) {
	p.InitFaction("Illusionists", "Illusionists", GetFactionTile(TileFactionIllusionists), tile)
}

func (p *Illusionists) GetInstance() *Faction {
	return &p.Faction
}

func (p *Illusionists) Print() {
	p.Faction.Print()
}

func (p *Illusionists) FirstIncome() {
}

func (p *Illusionists) Income() {
}

func (p *Illusionists) GetScience(pos int) int {
	return p.Faction.GetScience(pos)
}

func (p *Illusionists) FirstBuild(x int, y int) {
}

func (p *Illusionists) Build(x int, y int, needSpade int, building Building) error {
	return p.Faction.Build(x, y, needSpade, building)
}

func (p *Illusionists) Upgrade(x int, y int, target Building) error {
	return p.Faction.Upgrade(x, y, target)
}

func (p *Illusionists) AdvanceShip() error {
	return p.Faction.AdvanceShip()
}

func (p *Illusionists) AdvanceSpade() error {
	return p.Faction.AdvanceSpade()
}

func (p *Illusionists) SendScholar() error {
	return p.Faction.SendScholar()
}

func (p *Illusionists) SupployScholar() error {
	return p.Faction.SupployScholar()
}

func (p *Illusionists) PowerAction(item action.PowerActionItem) error {
	item.Power--
	return p.Faction.PowerAction(item)
}

func (p *Illusionists) Book(item action.BookActionItem, book Book) error {
	return p.Faction.Book(item, book)
}

func (p *Illusionists) Bridge(x1 int, y1 int, x2 int, y2 int) error {
	return p.Faction.Bridge(x1, y1, x2, y2)
}

func (p *Illusionists) Pass(tile TileItem) (error, TileItem) {
	return p.Faction.Pass(tile)
}

func (p *Illusionists) ReceiveCity(item CityItem) error {
	return p.Faction.ReceiveCity(item)
}

func (p *Illusionists) Dig(dig int) error {
	return p.Faction.Dig(dig)
}

func (p *Illusionists) TurnEnd() error {
	return p.Faction.TurnEnd()
}

func (p *Illusionists) PalaceTile(tile TileItem) error {
	return p.Faction.PalaceTile(tile)
}

func (p *Illusionists) SchoolTile(tile TileItem, science int) error {
	return p.Faction.SchoolTile(tile, science)
}

func (p *Illusionists) RoundTile(tile TileItem) error {
	return p.Faction.RoundTile(tile)
}

func (p *Illusionists) TileAction(category TileCategory, pos int) error {
	return p.Faction.TileAction(category, pos)
}
