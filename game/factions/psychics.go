package factions

import (
	"aoi/game/action"
	"aoi/game/resources"
)

type Psychics struct {
	Faction *Faction
}

func (p *Psychics) Init(tile resources.TileItem) {
	p.Faction = NewFaction("Psychics", "Psychics", resources.GetFactionTile(resources.TileFactionPsychics), tile)
}

func (p *Psychics) Income() {
	p.Faction.Income()
}

func (p *Psychics) FirstIncome() {
	p.Faction.FirstIncome()
}

func (p *Psychics) Print() {
	p.Faction.Print()
}

func (p *Psychics) GetInstance() *Faction {
	return p.Faction
}

func (p *Psychics) AdvanceShip() error {
	return p.Faction.AdvanceShip()
}

func (p *Psychics) AdvanceSpade() error {
	return p.Faction.AdvanceSpade()
}

func (p *Psychics) FirstBuild(x int, y int) {
	p.Faction.FirstBuild(x, y)
}

func (p *Psychics) Build(x int, y int, needSpade int, building resources.Building) error {
	return p.Faction.Build(x, y, needSpade, building)
}

func (p *Psychics) Upgrade(x int, y int, target resources.Building) error {
	return p.Faction.Upgrade(x, y, target)
}

func (p *Psychics) SendScholar() error {
	return p.Faction.SendScholar()
}

func (p *Psychics) SupployScholar() error {
	return p.Faction.SupployScholar()
}

func (p *Psychics) PowerAction(item action.PowerActionItem) error {
	return p.Faction.PowerAction(item)
}

func (p *Psychics) Book(item action.BookActionItem, book resources.Book) error {
	return p.Faction.Book(item, book)
}

func (p *Psychics) Bridge(x1 int, y1 int, x2 int, y2 int) error {
	return p.Faction.Bridge(x1, y1, x2, y2)
}

func (p *Psychics) ReceiveCity(item resources.CityItem) error {
	return p.Faction.ReceiveCity(item)
}

func (p *Psychics) Pass(tile resources.TileItem) (error, resources.TileItem) {
	return p.Faction.Pass(tile)
}

func (p *Psychics) Dig(dig int) error {
	return p.Faction.Dig(dig)
}

func (p *Psychics) ConvertDig(spade int) error {
	return p.Faction.ConvertDig(spade)
}

func (p *Psychics) TurnEnd() error {
	return p.Faction.TurnEnd()
}

func (p *Psychics) PalaceTile(tile resources.TileItem) error {
	return p.Faction.PalaceTile(tile)
}

func (p *Psychics) SchoolTile(tile resources.TileItem) error {
	return p.Faction.SchoolTile(tile)
}

func (p *Psychics) TileAction(category resources.TileCategory, pos int) error {
	return p.Faction.TileAction(category, pos)
}
