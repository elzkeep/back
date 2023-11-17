package factions

import (
	"aoi/game/action"
	. "aoi/game/resources"
)

type Philosophers struct {
	Faction
}

func (p *Philosophers) Init(tile TileItem, name string) {
	p.InitFaction(name, "Philosophers", GetFactionTile(TileFactionPhilosophers), tile)
}

func (p *Philosophers) GetInstance() *Faction {
	return &p.Faction
}

func (p *Philosophers) Print() {
	p.Faction.Print()
}

func (p *Philosophers) FirstIncome() {
	p.Faction.FirstIncome()
}

func (p *Philosophers) Income() {
	p.Faction.Income()
}

func (p *Philosophers) GetScience(pos int) int {
	return p.Faction.GetScience(pos)
}

func (p *Philosophers) FirstBuild(x int, y int, building Building) error {
	return p.Faction.FirstBuild(x, y, building)
}

func (p *Philosophers) Build(x int, y int, needSpade int, building Building) error {
	return p.Faction.Build(x, y, needSpade, building)
}

func (p *Philosophers) Upgrade(x int, y int, target Building, extra int) error {
	return p.Faction.Upgrade(x, y, target, extra)
}

func (p *Philosophers) AdvanceShip() error {
	return p.Faction.AdvanceShip()
}

func (p *Philosophers) AdvanceSpade() error {
	return p.Faction.AdvanceSpade()
}

func (p *Philosophers) SendScholar(pos int, inc int) error {
	return p.Faction.SendScholar(pos, inc)
}

func (p *Philosophers) SupployScholar(pos int, inc int) error {
	return p.Faction.SupployScholar(pos, inc)
}

func (p *Philosophers) PowerAction(item action.PowerActionItem) error {
	return p.Faction.PowerAction(item)
}

func (p *Philosophers) Book(item action.BookActionItem, book Book) error {
	return p.Faction.Book(item, book)
}

func (p *Philosophers) Bridge(x1 int, y1 int, x2 int, y2 int) error {
	return p.Faction.Bridge(x1, y1, x2, y2)
}

func (p *Philosophers) Pass(tile TileItem) (error, TileItem) {
	return p.Faction.Pass(tile)
}

func (p *Philosophers) ReceiveCity(item CityItem) error {
	return p.Faction.ReceiveCity(item)
}

func (p *Philosophers) Dig(x int, y int, dig int) error {
	return p.Faction.Dig(x, y, dig)
}

func (p *Philosophers) TurnEnd(round int) error {
	return p.Faction.TurnEnd(round)
}

func (p *Philosophers) PalaceTile(tile TileItem) error {
	return p.Faction.PalaceTile(tile)
}

func (p *Philosophers) SchoolTile(tile TileItem, science int) error {
	err := p.Faction.SchoolTile(tile, science)

	if err != nil {
		return err
	}

	price := Price{}
	if science == 0 {
		price.Book.Banking = 1
	} else if science == 1 {
		price.Book.Law = 1
	} else if science == 2 {
		price.Book.Engineering = 1
	} else if science == 3 {
		price.Book.Medicine = 1
	}

	p.Faction.ReceiveResource(price)

	return nil
}

func (p *Philosophers) RoundTile(tile TileItem) error {
	return p.Faction.RoundTile(tile)
}

func (p *Philosophers) TileAction(category TileCategory, pos int) error {
	return p.Faction.TileAction(category, pos)
}
