package factions

import (
	"aoi/game/action"
	. "aoi/game/resources"
)

type Navigators struct {
	Faction
}

func (p *Navigators) Init(tile TileItem) {
	p.InitFaction("Navigators", "Navigators", GetFactionTile(TileFactionNavigators), tile)
}

func (p *Navigators) GetInstance() *Faction {
	return &p.Faction
}

func (p *Navigators) Print() {
	p.Faction.Print()
}

func (p *Navigators) FirstIncome() {
}

func (p *Navigators) Income() {
}

func (p *Navigators) GetScience(pos int) int {
	return p.Faction.GetScience(pos)
}

func (p *Navigators) FirstBuild(x int, y int) error {
	return p.Faction.FirstBuild(x, y)
}

func (p *Navigators) Build(x int, y int, needSpade int, building Building) error {
	return p.Faction.Build(x, y, needSpade, building)
}

func (p *Navigators) Upgrade(x int, y int, target Building) error {
	return p.Faction.Upgrade(x, y, target)
}

func (p *Navigators) AdvanceShip() error {
	return p.Faction.AdvanceShip()
}

func (p *Navigators) AdvanceSpade() error {
	return p.Faction.AdvanceSpade()
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

func (p *Navigators) Book(item action.BookActionItem, book Book) error {
	return p.Faction.Book(item, book)
}

func (p *Navigators) Bridge(x1 int, y1 int, x2 int, y2 int) error {
	return p.Faction.Bridge(x1, y1, x2, y2)
}

func (p *Navigators) Pass(tile TileItem) (error, TileItem) {
	return p.Faction.Pass(tile)
}

func (p *Navigators) ReceiveCity(item CityItem) error {
	return p.Faction.ReceiveCity(item)
}

func (p *Navigators) Dig(dig int) error {
	return p.Faction.Dig(dig)
}

func (p *Navigators) TurnEnd(round int) error {
	return p.Faction.TurnEnd(round)
}

func (p *Navigators) PalaceTile(tile TileItem) error {
	return p.Faction.PalaceTile(tile)
}

func (p *Navigators) SchoolTile(tile TileItem, science int) error {
	return p.Faction.SchoolTile(tile, science)
}

func (p *Navigators) RoundTile(tile TileItem) error {
	return p.Faction.RoundTile(tile)
}

func (p *Navigators) TileAction(category TileCategory, pos int) error {
	return p.Faction.TileAction(category, pos)
}
