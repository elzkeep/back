package factions

import (
	"aoi/game/action"
	. "aoi/game/resources"
)

type Omar struct {
	Faction
}

func (p *Omar) Init(tile TileItem) {
	p.InitFaction("Omar", "Omar", GetFactionTile(TileFactionOmar), tile)
}

func (p *Omar) GetInstance() *Faction {
	return &p.Faction
}

func (p *Omar) Print() {
	p.Faction.Print()
}

func (p *Omar) FirstIncome() {
}

func (p *Omar) Income() {
}

func (p *Omar) GetScience(pos int) int {
	return p.Faction.GetScience(pos)
}

func (p *Omar) FirstBuild(x int, y int) {
}

func (p *Omar) Build(x int, y int, needSpade int, building Building) error {
	return p.Faction.Build(x, y, needSpade, building)
}

func (p *Omar) Upgrade(x int, y int, target Building) error {
	return p.Faction.Upgrade(x, y, target)
}

func (p *Omar) AdvanceShip() error {
	return p.Faction.AdvanceShip()
}

func (p *Omar) AdvanceSpade() error {
	return p.Faction.AdvanceSpade()
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

func (p *Omar) Book(item action.BookActionItem, book Book) error {
	return p.Faction.Book(item, book)
}

func (p *Omar) Bridge(x1 int, y1 int, x2 int, y2 int) error {
	return p.Faction.Bridge(x1, y1, x2, y2)
}

func (p *Omar) Pass(tile TileItem) (error, TileItem) {
	return p.Faction.Pass(tile)
}

func (p *Omar) ReceiveCity(item CityItem) error {
	return p.Faction.ReceiveCity(item)
}

func (p *Omar) Dig(dig int) error {
	return p.Faction.Dig(dig)
}

func (p *Omar) TurnEnd() error {
	return p.Faction.TurnEnd()
}

func (p *Omar) PalaceTile(tile TileItem) error {
	return p.Faction.PalaceTile(tile)
}

func (p *Omar) SchoolTile(tile TileItem, science int) error {
	return p.Faction.SchoolTile(tile, science)
}

func (p *Omar) RoundTile(tile TileItem) error {
	return p.Faction.RoundTile(tile)
}

func (p *Omar) TileAction(category TileCategory, pos int) error {
	return p.Faction.TileAction(category, pos)
}
