package factions

import (
	"aoi/game/action"
	. "aoi/game/resources"
)

type Basic struct {
	Faction
}

func (p *Basic) Init(tile TileItem) {
	p.InitFaction("Basic", "Basic", GetFactionTile(TileFactionBlessed), tile)
}

func (p *Basic) GetInstance() *Faction {
	return &p.Faction
}

func (p *Basic) Print() {
	p.Faction.Print()
}

func (p *Basic) FirstIncome() {
	p.Faction.FirstIncome()
}

func (p *Basic) Income() {
	p.Faction.Income()
}

func (p *Basic) GetScience(pos int) int {
	return p.Faction.GetScience(pos) + 3
}

func (p *Basic) FirstBuild(x int, y int, building Building) error {
	return p.Faction.FirstBuild(x, y, building)
}

func (p *Basic) Build(x int, y int, needSpade int, building Building) error {
	return p.Faction.Build(x, y, needSpade, building)
}

func (p *Basic) Upgrade(x int, y int, target Building, extra int) error {
	return p.Faction.Upgrade(x, y, target, extra)
}

func (p *Basic) AdvanceShip() error {
	return p.Faction.AdvanceShip()
}

func (p *Basic) AdvanceSpade() error {
	return p.Faction.AdvanceSpade()
}

func (p *Basic) SendScholar(pos int, inc int) error {
	return p.Faction.SendScholar(pos, inc)
}

func (p *Basic) SupployScholar(pos int, inc int) error {
	return p.Faction.SupployScholar(pos, inc)
}

func (p *Basic) PowerAction(item action.PowerActionItem) error {
	return p.Faction.PowerAction(item)
}

func (p *Basic) Book(item action.BookActionItem, book Book) error {
	return p.Faction.Book(item, book)
}

func (p *Basic) Bridge(x1 int, y1 int, x2 int, y2 int) error {
	return p.Faction.Bridge(x1, y1, x2, y2)
}

func (p *Basic) Pass(tile TileItem) (error, TileItem) {
	return p.Faction.Pass(tile)
}

func (p *Basic) ReceiveCity(item CityItem) error {
	return p.Faction.ReceiveCity(item)
}

func (p *Basic) Dig(x int, y int, dig int) error {
	return p.Faction.Dig(x, y, dig)
}

func (p *Basic) TurnEnd(round int) error {
	return p.Faction.TurnEnd(round)
}

func (p *Basic) PalaceTile(tile TileItem) error {
	return p.Faction.PalaceTile(tile)
}

func (p *Basic) SchoolTile(tile TileItem, science int) error {
	return p.Faction.SchoolTile(tile, science)
}

func (p *Basic) RoundTile(tile TileItem) error {
	return p.Faction.RoundTile(tile)
}

func (p *Basic) TileAction(category TileCategory, pos int) error {
	return p.Faction.TileAction(category, pos)
}
