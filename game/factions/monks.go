package factions

import (
	"aoi/game/action"
	. "aoi/game/resources"
)

type Monks struct {
	Faction
}

func (p *Monks) Init(tile TileItem) {
	p.InitFaction("Monks", "Monks", GetFactionTile(TileFactionMonks), tile)
	p.FirstBuilding = SA
}

func (p *Monks) GetInstance() *Faction {
	return &p.Faction
}

func (p *Monks) Print() {
	p.Faction.Print()
}

func (p *Monks) FirstIncome() {
	p.Faction.FirstIncome()
}

func (p *Monks) Income() {
	p.Faction.Income()
}

func (p *Monks) GetScience(pos int) int {
	return p.Faction.GetScience(pos)
}

func (p *Monks) FirstBuild(x int, y int) error {
	return p.Faction.FirstBuild(x, y)
}

func (p *Monks) Build(x int, y int, needSpade int, building Building) error {
	return p.Faction.Build(x, y, needSpade, building)
}

func (p *Monks) Upgrade(x int, y int, target Building) error {
	return p.Faction.Upgrade(x, y, target)
}

func (p *Monks) AdvanceShip() error {
	return p.Faction.AdvanceShip()
}

func (p *Monks) AdvanceSpade() error {
	return p.Faction.AdvanceSpade()
}

func (p *Monks) SendScholar() error {
	return p.Faction.SendScholar()
}

func (p *Monks) SupployScholar() error {
	return p.Faction.SupployScholar()
}

func (p *Monks) PowerAction(item action.PowerActionItem) error {
	return p.Faction.PowerAction(item)
}

func (p *Monks) Book(item action.BookActionItem, book Book) error {
	return p.Faction.Book(item, book)
}

func (p *Monks) Bridge(x1 int, y1 int, x2 int, y2 int) error {
	return p.Faction.Bridge(x1, y1, x2, y2)
}

func (p *Monks) Pass(tile TileItem) (error, TileItem) {
	return p.Faction.Pass(tile)
}

func (p *Monks) ReceiveCity(item CityItem) error {
	return p.Faction.ReceiveCity(item)
}

func (p *Monks) Dig(dig int) error {
	return p.Faction.Dig(dig)
}

func (p *Monks) TurnEnd(round int) error {
	return p.Faction.TurnEnd(round)
}

func (p *Monks) PalaceTile(tile TileItem) error {
	return p.Faction.PalaceTile(tile)
}

func (p *Monks) SchoolTile(tile TileItem, science int) error {
	return p.Faction.SchoolTile(tile, science)
}

func (p *Monks) RoundTile(tile TileItem) error {
	return p.Faction.RoundTile(tile)
}

func (p *Monks) TileAction(category TileCategory, pos int) error {
	return p.Faction.TileAction(category, pos)
}
