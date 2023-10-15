package factions

import (
	"aoi/game/action"
	"aoi/game/resources"
)

type Goblins struct {
	Faction *Faction
}

func (p *Goblins) Init(tile resources.TileItem) {
	p.Faction = NewFaction("Goblins", "Goblins", resources.GetFactionTile(resources.TileFactionGoblins), tile)
}

func (p *Goblins) Income() {
	p.Faction.Income()
}

func (p *Goblins) FirstIncome() {
	p.Faction.FirstIncome()
}

func (p *Goblins) Print() {
	p.Faction.Print()
}

func (p *Goblins) GetInstance() *Faction {
	return p.Faction
}

func (p *Goblins) AdvanceShip() error {
	return p.Faction.AdvanceShip()
}

func (p *Goblins) AdvanceSpade() error {
	return p.Faction.AdvanceSpade()
}

func (p *Goblins) FirstBuild(x int, y int) {
	p.Faction.FirstBuild(x, y)
}

func (p *Goblins) Build(x int, y int, needSpade int, building resources.Building) error {
	return p.Faction.Build(x, y, needSpade, building)
}

func (p *Goblins) Upgrade(x int, y int, target resources.Building) error {
	return p.Faction.Upgrade(x, y, target)
}

func (p *Goblins) SendScholar() error {
	return p.Faction.SendScholar()
}

func (p *Goblins) SupployScholar() error {
	return p.Faction.SupployScholar()
}

func (p *Goblins) PowerAction(item action.PowerActionItem) error {
	return p.Faction.PowerAction(item)
}

func (p *Goblins) Book(item action.BookActionItem) error {
	return p.Faction.Book(item)
}

func (p *Goblins) Bridge(x1 int, y1 int, x2 int, y2 int) error {
	return p.Faction.Bridge(x1, y1, x2, y2)
}

func (p *Goblins) PassIncome() {
	p.Faction.PassIncome()
}

func (p *Goblins) ReceiveCity(item resources.CityItem) error {
	return p.Faction.ReceiveCity(item)
}

func (p *Goblins) Pass(tile resources.TileItem) (error, resources.TileItem) {
	return p.Faction.Pass(tile)
}

func (p *Goblins) Dig(dig int) error {
	err := p.Faction.Dig(dig)

	if err != nil {
		p.Faction.ReceiveResource(resources.Price{Coin: 2})
	}

	return err
}

func (p *Goblins) ConvertDig(spade int) error {
	return p.Faction.ConvertDig(spade)
}

func (p *Goblins) TurnEnd() error {
	return p.Faction.TurnEnd()
}

func (p *Goblins) PalaceTile(tile resources.TileItem) error {
	return p.Faction.PalaceTile(tile)
}

func (p *Goblins) SchoolTile(tile resources.TileItem) error {
	return p.Faction.SchoolTile(tile)
}

func (p *Goblins) TileAction(category resources.TileCategory, pos int) error {
	return p.Faction.TileAction(category, pos)
}
