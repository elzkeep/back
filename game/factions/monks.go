package factions

import (
	"aoi/game/action"
	"aoi/game/color"
	"aoi/game/resources"
	"aoi/game/resources/city"
)

type Monks struct {
	Faction *Faction
}

func (p *Monks) Init() {
	p.Faction = NewFaction("몽크", "Monks", color.Yellow)
}

func (p *Monks) Income() {
	p.Faction.Income()
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

func (p *Monks) ReceiveCity(item city.CityItem) error {
	return p.Faction.ReceiveCity(item)
}

func (p *Monks) Pass(tile *resources.RoundTileItem) error {
	return p.Faction.Pass(tile)
}

func (p *Monks) Dig(dig int) error {
	return p.Faction.Dig(dig)
}

func (p *Monks) ConvertDig(spade int) error {
	return p.Faction.ConvertDig(spade)
}
