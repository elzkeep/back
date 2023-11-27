package factions

import (
	"aoi/game/action"
	. "aoi/game/resources"
)

type Lizards struct {
	Faction
}

func (p *Lizards) Init(tile TileItem, name string) {
	p.InitFaction(name, "Lizards", GetFactionTile(TileFactionLizards), tile)
}

func (p *Lizards) GetInstance() *Faction {
	return &p.Faction
}

func (p *Lizards) Print() {
	p.Faction.Print()
}

func (p *Lizards) FirstIncome() {
	p.Faction.FirstIncome()
}

func (p *Lizards) Income() {
	p.Faction.Income()
}

func (p *Lizards) GetScience(pos int) int {
	return p.Faction.GetScience(pos)
}

func (p *Lizards) FirstBuild(x int, y int, building Building) error {
	return p.Faction.FirstBuild(x, y, building)
}

func (p *Lizards) Build(x int, y int, needSpade int, building Building) error {
	return p.Faction.Build(x, y, needSpade, building)
}

func (p *Lizards) Upgrade(x int, y int, target Building, extra int) error {
	return p.Faction.Upgrade(x, y, target, extra)
}

func (p *Lizards) Downgrade(x int, y int) error {
	return p.Faction.Downgrade(x, y)
}

func (p *Lizards) AdvanceShip() error {
	return p.Faction.AdvanceShip()
}

func (p *Lizards) AdvanceSpade() error {
	return p.Faction.AdvanceSpade()
}

func (p *Lizards) SendScholar(pos int, inc int) error {
	return p.Faction.SendScholar(pos, inc)
}

func (p *Lizards) SupployScholar(pos int, inc int) error {
	return p.Faction.SupployScholar(pos, inc)
}

func (p *Lizards) PowerAction(item action.PowerActionItem) error {
	return p.Faction.PowerAction(item)
}

func (p *Lizards) Book(item action.BookActionItem, book Book) error {
	return p.Faction.Book(item, book)
}

func (p *Lizards) Bridge(x1 int, y1 int, x2 int, y2 int) error {
	return p.Faction.Bridge(x1, y1, x2, y2)
}

func (p *Lizards) Pass(tile TileItem) (error, TileItem) {
	return p.Faction.Pass(tile)
}

func (p *Lizards) ReceiveCity(item CityItem) error {
	err := p.Faction.ReceiveCity(item)

	if err == nil {
		p.Faction.ReceiveResource(Price{Spade: 1, Building: D})
	}

	return err
}

func (p *Lizards) Dig(x int, y int, dig int) error {
	return p.Faction.Dig(x, y, dig)
}

func (p *Lizards) TurnEnd(round int, pass bool) error {
	return p.Faction.TurnEnd(round, pass)
}

func (p *Lizards) PalaceTile(tile TileItem) error {
	return p.Faction.PalaceTile(tile)
}

func (p *Lizards) SchoolTile(tile TileItem, science int) error {
	return p.Faction.SchoolTile(tile, science)
}

func (p *Lizards) RoundTile(tile TileItem) error {
	return p.Faction.RoundTile(tile)
}

func (p *Lizards) TileAction(category TileCategory, pos int) error {
	return p.Faction.TileAction(category, pos)
}
