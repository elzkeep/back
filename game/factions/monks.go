package factions

import (
	"aoi/game/action"
	"aoi/game/resources"
	"log"
)

type Monks struct {
	Faction *Faction
}

func (p *Monks) Init(tile resources.TileItem) {
	p.Faction = NewFaction("Monks", "Monks", resources.GetFactionTile(resources.TileFactionMonks), tile)
	p.Faction.FirstBuilding = resources.SA
}

func (p *Monks) Income() {
	p.Faction.Income()
}

func (p *Monks) FirstIncome() {
	p.Faction.FirstIncome()
}

func (p *Monks) Print() {
	p.Faction.Print()
}

func (p *Monks) GetInstance() *Faction {
	return p.Faction
}

func (p *Monks) AdvanceShip() error {
	return p.Faction.AdvanceShip()
}

func (p *Monks) AdvanceSpade() error {
	return p.Faction.AdvanceSpade()
}

func (p *Monks) FirstBuild(x int, y int) {
	log.Println("monks firstbuild")
	p.Faction.FirstBuild(x, y)
}

func (p *Monks) Build(x int, y int, needSpade int) error {
	return p.Faction.Build(x, y, needSpade)
}

func (p *Monks) Upgrade(x int, y int, target resources.Building) error {
	return p.Faction.Upgrade(x, y, target)
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

func (p *Monks) Book(item action.BookActionItem) error {
	return p.Faction.Book(item)
}

func (p *Monks) Bridge(x1 int, y1 int, x2 int, y2 int) error {
	return p.Faction.Bridge(x1, y1, x2, y2)
}

func (p *Monks) PassIncome() {
	p.Faction.PassIncome()
}

func (p *Monks) ReceiveCity(item resources.CityItem) error {
	return p.Faction.ReceiveCity(item)
}

func (p *Monks) Pass(tile resources.TileItem) (error, resources.TileItem) {
	return p.Faction.Pass(tile)
}

func (p *Monks) Dig(dig int) error {
	return p.Faction.Dig(dig)
}

func (p *Monks) ConvertDig(spade int) error {
	return p.Faction.ConvertDig(spade)
}

func (p *Monks) TurnEnd() error {
	return p.Faction.TurnEnd()
}

func (p *Monks) PalaceTile(tile resources.TileItem) error {
	return p.Faction.PalaceTile(tile)
}

func (p *Monks) SchoolTile(tile resources.TileItem) error {
	return p.Faction.SchoolTile(tile)
}

func (p *Monks) TileAction(category resources.TileCategory, pos int) error {
	return p.Faction.TileAction(category, pos)
}
