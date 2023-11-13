package factions

import (
	"aoi/game/action"
	. "aoi/game/resources"
)

type Blessed struct {
	Faction
}

func (p *Blessed) Init(tile TileItem) {
	p.InitFaction("Blessed", "Blessed", GetFactionTile(TileFactionBlessed), tile)
}

func (p *Blessed) GetInstance() *Faction {
	return &p.Faction
}

func (p *Blessed) Print() {
	p.Faction.Print()
}

func (p *Blessed) FirstIncome() {
	p.Faction.FirstIncome()
}

func (p *Blessed) Income() {
	p.Faction.Income()
}

func (p *Blessed) GetScience(pos int) int {
	return p.Faction.GetScience(pos) + 3
}

func (p *Blessed) FirstBuild(x int, y int) error {
	return p.Faction.FirstBuild(x, y)
}

func (p *Blessed) Build(x int, y int, needSpade int, building Building) error {
	return p.Faction.Build(x, y, needSpade, building)
}

func (p *Blessed) Upgrade(x int, y int, target Building, extra int) error {
	return p.Faction.Upgrade(x, y, target, extra)
}

func (p *Blessed) AdvanceShip() error {
	return p.Faction.AdvanceShip()
}

func (p *Blessed) AdvanceSpade() error {
	return p.Faction.AdvanceSpade()
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

func (p *Blessed) Book(item action.BookActionItem, book Book) error {
	return p.Faction.Book(item, book)
}

func (p *Blessed) Bridge(x1 int, y1 int, x2 int, y2 int) error {
	return p.Faction.Bridge(x1, y1, x2, y2)
}

func (p *Blessed) Pass(tile TileItem) (error, TileItem) {
	return p.Faction.Pass(tile)
}

func (p *Blessed) ReceiveCity(item CityItem) error {
	return p.Faction.ReceiveCity(item)
}

func (p *Blessed) Dig(x int, y int, dig int) error {
	return p.Faction.Dig(x, y, dig)
}

func (p *Blessed) TurnEnd(round int) error {
	return p.Faction.TurnEnd(round)
}

func (p *Blessed) PalaceTile(tile TileItem) error {
	return p.Faction.PalaceTile(tile)
}

func (p *Blessed) SchoolTile(tile TileItem, science int) error {
	return p.Faction.SchoolTile(tile, science)
}

func (p *Blessed) RoundTile(tile TileItem) error {
	return p.Faction.RoundTile(tile)
}

func (p *Blessed) TileAction(category TileCategory, pos int) error {
	return p.Faction.TileAction(category, pos)
}
