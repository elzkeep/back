package factions

import (
	"aoi/game/action"
	. "aoi/game/resources"
)

type Felines struct {
	Faction
}

func (p *Felines) Init(tile TileItem, name string) {
	p.InitFaction(name, "Felines", GetFactionTile(TileFactionFelines), tile)
}

func (p *Felines) GetInstance() *Faction {
	return &p.Faction
}

func (p *Felines) Print() {
	p.Faction.Print()
}

func (p *Felines) FirstIncome() {
	p.Faction.FirstIncome()
}

func (p *Felines) Income() {
	p.Faction.Income()
}

func (p *Felines) GetScience(pos int) int {
	return p.Faction.GetScience(pos)
}

func (p *Felines) FirstBuild(x int, y int, building Building) error {
	return p.Faction.FirstBuild(x, y, building)
}

func (p *Felines) Build(x int, y int, needSpade int, building Building) error {
	return p.Faction.Build(x, y, needSpade, building)
}

func (p *Felines) Upgrade(x int, y int, target Building, extra int) error {
	return p.Faction.Upgrade(x, y, target, extra)
}

func (p *Felines) AdvanceShip() error {
	return p.Faction.AdvanceShip()
}

func (p *Felines) AdvanceSpade() error {
	return p.Faction.AdvanceSpade()
}

func (p *Felines) SendScholar(pos int, inc int) error {
	return p.Faction.SendScholar(pos, inc)
}

func (p *Felines) SupployScholar(pos int, inc int) error {
	return p.Faction.SupployScholar(pos, inc)
}

func (p *Felines) PowerAction(item action.PowerActionItem) error {
	return p.Faction.PowerAction(item)
}

func (p *Felines) Book(item action.BookActionItem, book Book) error {
	return p.Faction.Book(item, book)
}

func (p *Felines) Bridge(x1 int, y1 int, x2 int, y2 int) error {
	return p.Faction.Bridge(x1, y1, x2, y2)
}

func (p *Felines) Pass(tile TileItem) (error, TileItem) {
	return p.Faction.Pass(tile)
}

func (p *Felines) ReceiveCity(item CityItem) error {
	err := p.Faction.ReceiveCity(item)

	if err == nil {
		p.Faction.ReceiveResource(Price{Book: Book{Any: 1}, Science: Science{Any: 3}})
	}

	return err
}

func (p *Felines) Dig(x int, y int, dig int) error {
	return p.Faction.Dig(x, y, dig)
}

func (p *Felines) TurnEnd(round int) error {
	return p.Faction.TurnEnd(round)
}

func (p *Felines) PalaceTile(tile TileItem) error {
	return p.Faction.PalaceTile(tile)
}

func (p *Felines) SchoolTile(tile TileItem, science int) error {
	return p.Faction.SchoolTile(tile, science)
}

func (p *Felines) RoundTile(tile TileItem) error {
	return p.Faction.RoundTile(tile)
}

func (p *Felines) TileAction(category TileCategory, pos int) error {
	return p.Faction.TileAction(category, pos)
}
