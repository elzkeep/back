package factions

import (
	"aoi/game/action"
	. "aoi/game/resources"
)

type Moles struct {
	Faction
}

func (p *Moles) Init(tile TileItem) {
	p.InitFaction("Moles", "Moles", GetFactionTile(TileFactionMoles), tile)
}

func (p *Moles) GetInstance() *Faction {
	return &p.Faction
}

func (p *Moles) Print() {
	p.Faction.Print()
}

func (p *Moles) FirstIncome() {
	p.Faction.FirstIncome()
}

func (p *Moles) Income() {
	p.Faction.Income()
}

func (p *Moles) GetScience(pos int) int {
	return p.Faction.GetScience(pos)
}

func (p *Moles) FirstBuild(x int, y int) error {
	return p.Faction.FirstBuild(x, y)
}

func (p *Moles) Build(x int, y int, needSpade int, building Building) error {
	return p.Faction.Build(x, y, needSpade, building)
}

func (p *Moles) Upgrade(x int, y int, target Building) error {
	return p.Faction.Upgrade(x, y, target)
}

func (p *Moles) AdvanceShip() error {
	return p.Faction.AdvanceShip()
}

func (p *Moles) AdvanceSpade() error {
	return p.Faction.AdvanceSpade()
}

func (p *Moles) SendScholar() error {
	return p.Faction.SendScholar()
}

func (p *Moles) SupployScholar() error {
	return p.Faction.SupployScholar()
}

func (p *Moles) PowerAction(item action.PowerActionItem) error {
	return p.Faction.PowerAction(item)
}

func (p *Moles) Book(item action.BookActionItem, book Book) error {
	return p.Faction.Book(item, book)
}

func (p *Moles) Bridge(x1 int, y1 int, x2 int, y2 int) error {
	return p.Faction.Bridge(x1, y1, x2, y2)
}

func (p *Moles) Pass(tile TileItem) (error, TileItem) {
	return p.Faction.Pass(tile)
}

func (p *Moles) ReceiveCity(item CityItem) error {
	return p.Faction.ReceiveCity(item)
}

func (p *Moles) Dig(dig int) error {
	return p.Faction.Dig(dig)
}

func (p *Moles) TurnEnd(round int) error {
	return p.Faction.TurnEnd(round)
}

func (p *Moles) PalaceTile(tile TileItem) error {
	return p.Faction.PalaceTile(tile)
}

func (p *Moles) SchoolTile(tile TileItem, science int) error {
	return p.Faction.SchoolTile(tile, science)
}

func (p *Moles) RoundTile(tile TileItem) error {
	return p.Faction.RoundTile(tile)
}

func (p *Moles) TileAction(category TileCategory, pos int) error {
	return p.Faction.TileAction(category, pos)
}
