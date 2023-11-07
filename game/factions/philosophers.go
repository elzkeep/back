package factions

import (
	"aoi/game/action"
	. "aoi/game/resources"
)

type Philosophers struct {
	Faction
}

func (p *Philosophers) Init(tile TileItem) {
	p.InitFaction("Philosophers", "Philosophers", GetFactionTile(TileFactionPhilosophers), tile)
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

func (p *Philosophers) FirstBuild(x int, y int) error {
	return p.Faction.FirstBuild(x, y)
}

func (p *Philosophers) Build(x int, y int, needSpade int, building Building) error {
	return p.Faction.Build(x, y, needSpade, building)
}

func (p *Philosophers) Upgrade(x int, y int, target Building) error {
	return p.Faction.Upgrade(x, y, target)
}

func (p *Philosophers) AdvanceShip() error {
	return p.Faction.AdvanceShip()
}

func (p *Philosophers) AdvanceSpade() error {
	return p.Faction.AdvanceSpade()
}

func (p *Philosophers) SendScholar() error {
	return p.Faction.SendScholar()
}

func (p *Philosophers) SupployScholar() error {
	return p.Faction.SupployScholar()
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

func (p *Philosophers) Dig(dig int) error {
	return p.Faction.Dig(dig)
}

func (p *Philosophers) TurnEnd(round int) error {
	return p.Faction.TurnEnd(round)
}

func (p *Philosophers) PalaceTile(tile TileItem) error {
	return p.Faction.PalaceTile(tile)
}

func (p *Philosophers) SchoolTile(tile TileItem, science int) error {
	if science == 0 {
		tile.Once.Book.Banking++
	} else if science == 1 {
		tile.Once.Book.Law++
	} else if science == 2 {
		tile.Once.Book.Engineering++
	} else if science == 3 {
		tile.Once.Book.Medicine++
	}

	return p.Faction.SchoolTile(tile, science)
}

func (p *Philosophers) RoundTile(tile TileItem) error {
	return p.Faction.RoundTile(tile)
}

func (p *Philosophers) TileAction(category TileCategory, pos int) error {
	return p.Faction.TileAction(category, pos)
}
