package factions

import (
	"aoi/game/action"
	"aoi/game/color"
	. "aoi/game/resources"
	"errors"
	"sort"

	"math"
)

type FactionInterface interface {
	GetInstance() *Faction

	Init(tile TileItem, name string)
	Print()
	FirstIncome()
	Income()
	GetScience(pos int) int

	FirstBuild(x int, y int, building Building) error
	Build(x int, y int, needSpade int, building Building) error
	Upgrade(x int, y int, target Building, extra int) error
	Downgrade(x int, y int) error
	AdvanceShip() error
	AdvanceSpade() error
	SendScholar(pos int, inc int) error
	SupployScholar(pos int, inc int) error
	PowerAction(item action.PowerActionItem) error
	Book(item action.BookActionItem, book Book) error
	Bridge(x1 int, y1 int, x2 int, y2 int) error
	Pass(tile TileItem) (error, TileItem)
	ReceiveCity(item CityItem) error
	Dig(x int, y int, dig int) error
	TurnEnd(round int, pass bool) error
	PalaceTile(tile TileItem) error
	SchoolTile(tile TileItem, science int) error
	InnovationTile(tile TileItem, book Book) error
	RoundTile(tile TileItem) error
	TileAction(category TileCategory, pos int) error
}

type Faction struct {
	Name              string           `json:"name"`
	Ename             string           `json:"ename"`
	Color             color.Color      `json:"color"`
	Resource          Resource         `json:"resource"`
	Tiles             []TileItem       `json:"tiles"`
	MaxBuilding       [13]int          `json:"maxBuilding"`
	Building          [13]int          `json:"building"`
	BuildingPower     [13]int          `json:"buildingPower"`
	BuildingList      []Position       `json:"buildingList"`
	BridgeList        []BridgePosition `json:"bridgeList"`
	AnnexList         []Position       `json:"AnnexList"`
	Price             [13]Price        `json:"price"`
	AdvanceShipPrice  Price            `json:"advanceShipPrice"`
	AdvanceSpadePrice Price            `json:"advanceSpadePrice"`
	Incomes           [][]Price        `json:"incomes"`
	Receive           Price            `json:"receive"`
	Point             int              `json:"point"`
	TownPower         int              `json:"townPower"`
	Spade             int              `json:"spade"`
	DigPosition       []Position       `json:"digPosition"`
	MaxSpade          int              `json:"maxSpade"`
	Ship              int              `json:"ship"`
	MaxShip           int              `json:"maxShip"`
	MaxPrist          int              `json:"maxPrist"`
	Science           []int            `json:"science"`
	Key               int              `json:"key"`
	Action            bool             `json:"action"`
	BuildAction       bool             `json:"buildAction"`
	MaxBridge         int              `json:"maxBridge"`
	ExtraBuild        int              `json:"extraBuild"`
	VP                int              `json:"vp"`
	City              int              `json:"city"`
	CityBuildingList  []Position       `json:"cityBuildingList"`
	Cities            []CityItem       `json:"cities"`
	Type              TileType         `json:"type"`
	IsPass            bool             `json:"isPass"`
	FirstBuilding     Building         `json:"-"`
}

func (item *Faction) InitFaction(name string, ename string, factionTile TileItem, colorTile TileItem) {
	item.Name = name
	item.Ename = ename
	item.Type = factionTile.Type

	item.Resource.Coin = 15
	item.Resource.Worker = 3
	item.Resource.Power = [3]int{5, 7, 0}
	item.Resource.Building = None
	item.VP = 20
	item.Spade = 0
	item.MaxSpade = 2
	item.Ship = 0
	item.MaxShip = 3
	item.MaxPrist = 7
	item.Science = []int{0, 0, 0, 0}
	item.MaxBridge = 3
	item.Key = 0
	item.Action = false
	item.ExtraBuild = 0
	item.IsPass = false
	item.FirstBuilding = D

	item.DigPosition = make([]Position, 0)
	item.Cities = make([]CityItem, 0)
	item.Tiles = make([]TileItem, 0)
	if ename != "" {
		colorTile.Use = false
		factionTile.Use = false

		item.Color = colorTile.Color
		item.Tiles = append(item.Tiles, colorTile)
		item.Tiles = append(item.Tiles, factionTile)
	} else {
		item.Color = color.None
	}

	item.MaxBuilding = [13]int{0, 9, 4, 3, 1, 1, 1, 2, 1, 1, 1, 1, 1}
	item.Building = [13]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	item.BuildingPower = [13]int{0, 1, 2, 2, 3, 3, 1, 2, 2, 2, 3, 3, 4}
	item.Price = [13]Price{
		{Worker: 0, Coin: 0},
		{Worker: 1, Coin: 2},
		{Worker: 2, Coin: 3},
		{Worker: 3, Coin: 5},
		{Worker: 4, Coin: 6},
		{Worker: 5, Coin: 8},
		{Worker: 0, Coin: 0},
		{Worker: 0, Coin: 0},
		{Worker: 0, Coin: 0},
		{Worker: 0, Coin: 0},
		{Worker: 0, Coin: 0},
		{Worker: 0, Coin: 0},
		{Worker: 0, Coin: 0},
	}

	item.AdvanceShipPrice = Price{Prist: 1, Coin: 4}
	item.AdvanceSpadePrice = Price{Prist: 1, Worker: 1, Coin: 5}

	item.Incomes = [][]Price{
		{Price{}},
		{Price{Worker: 1}, Price{Worker: 1}, Price{Worker: 1}, Price{Worker: 1}, Price{Worker: 1}, Price{}, Price{Worker: 1}, Price{Worker: 1}, Price{Worker: 1}, Price{Worker: 1}},
		{Price{}, Price{Coin: 2, Power: 1}, Price{Coin: 2, Power: 1}, Price{Coin: 2, Power: 2}, Price{Coin: 2, Power: 2}},
		{Price{}, Price{Prist: 1}, Price{Prist: 1}, Price{Prist: 1}},
		{Price{}, Price{Power: 0}},
		{Price{}, Price{Prist: 1}},
	}

	if colorTile.Color == color.Gray {
		item.Incomes[D][0] = Price{Worker: 1, Coin: 2}
		item.Incomes[TP][1] = Price{Coin: 3, Power: 1}
	} else if colorTile.Color == color.Brown {
		item.AdvanceSpadePrice = Price{Prist: 1, Worker: 1, Coin: 1}
	}

	item.Point = 10
	item.TownPower = 7

	item.BuildingList = make([]Position, 0)
	item.BridgeList = make([]BridgePosition, 0)
	item.AnnexList = make([]Position, 0)

	/*
		item.Resource.Coin = 100
		item.Resource.Worker = 100
		item.Resource.Prist = 7
		item.Resource.Book = Book{Banking: 5, Law: 5, Engineering: 5, Medicine: 5}

		item.Resource.Power = [3]int{0, 0, 12}

		if colorTile.Type == TileColorGray {
			item.Incomes[D][0] = Price{Worker: 1, Coin: 2}
			item.Incomes[TP][1] = Price{Coin: 3, Power: 1}
		} else if colorTile.Type == TileColorBrown {
			item.AdvanceSpadePrice = Price{Prist: 1, Worker: 1, Coin: 1}
		}
	*/
}

func (p *Faction) GetShipDistance(tile bool) int {
	ship := 0
	if tile == true {
		for _, v := range p.Tiles {
			ship += v.Ship
		}
	}

	return p.Ship + ship
}

func (p *Faction) GetWorkerForSpade() int {
	if p.Spade == 2 {
		return 1
	} else if p.Spade == 1 {
		return 2
	} else {
		return 3
	}
}

func (p *Faction) GetSpadeCount() int {
	c := p.GetWorkerForSpade()

	workerSpade := int(math.Floor(float64(p.Resource.Worker) / float64(c)))
	spade := p.Resource.Spade

	return workerSpade + spade
}

func (p *Faction) ReceivePower(value int, vp bool) {
	remain := value
	c := 0
	total := 0

	for i := 0; i <= 1; i++ {
		if p.Resource.Power[i] < remain {
			c = p.Resource.Power[i]
		} else {
			c = remain
		}

		total += c

		p.Resource.Power[i+1] += c
		p.Resource.Power[i] -= c

		remain -= c

		if remain <= 0 {
			break
		}
	}

	if vp == true && total > 1 {
		p.VP = p.VP - (total - 1)
	}
}

func (p *Faction) GetHavePowerCount() int {
	power := p.Resource.Power[2]
	burnPower := int(math.Floor(float64(p.Resource.Power[1]) / 2.0))

	return power + burnPower
}

func (p *Faction) ReceiveResource(receive Price) {
	p.Resource.Coin += receive.Coin
	p.Resource.Worker += receive.Worker
	p.Resource.Prist += receive.Prist
	p.Resource.Spade += receive.Spade

	p.Resource.Book.Any += receive.Book.Any
	p.Resource.Book.Banking += receive.Book.Banking
	p.Resource.Book.Law += receive.Book.Law
	p.Resource.Book.Engineering += receive.Book.Engineering
	p.Resource.Book.Medicine += receive.Book.Medicine

	p.Resource.Science.Any += receive.Science.Any
	p.Resource.Science.Single += receive.Science.Single

	p.Resource.Annex += receive.Annex

	p.Resource.City += receive.City
	p.VP += receive.VP

	p.Resource.Bridge += receive.Bridge

	p.Resource.TpUpgrade += receive.TpUpgrade
	p.Resource.Downgrade += receive.Downgrade

	p.Resource.SchoolTile += receive.Tile

	if receive.Building != None {
		p.Resource.Building = receive.Building
		p.BuildAction = false
	}

	if p.Resource.Bridge > p.MaxBridge {
		p.Resource.Bridge = p.MaxBridge
	}

	if p.Resource.Prist > p.MaxPrist {
		p.Resource.Prist = p.MaxPrist
	}

	if receive.ShipUpgrade == 1 {
		if p.Color == color.Blue {
			if p.Ship == 1 {
				p.ReceiveResource(Price{VP: 3})
			} else if p.Ship == 2 {
				p.ReceiveResource(Price{Book: Book{Any: 2}})
			}
		} else {
			if p.Ship == 0 {
				p.ReceiveResource(Price{VP: 2})
			} else if p.Ship == 1 {
				p.ReceiveResource(Price{Book: Book{Any: 2}})
			} else if p.Ship == 2 {
				p.ReceiveResource(Price{VP: 4})
			}
		}
	} else if receive.ShipUpgrade == 2 {
		if p.Color == color.Blue {
			if p.Ship == 1 {
				p.ReceiveResource(Price{VP: 3})
				p.ReceiveResource(Price{Book: Book{Any: 2}})
			} else if p.Ship == 2 {
				p.ReceiveResource(Price{Book: Book{Any: 2}})
			}
		} else {
			if p.Ship == 0 {
				p.ReceiveResource(Price{VP: 2})
				p.ReceiveResource(Price{Book: Book{Any: 2}})
			} else if p.Ship == 1 {
				p.ReceiveResource(Price{Book: Book{Any: 2}})
				p.ReceiveResource(Price{VP: 4})
			} else if p.Ship == 2 {
				p.ReceiveResource(Price{VP: 4})
			}
		}
	}

	p.Ship += receive.ShipUpgrade
	if p.Ship > p.MaxShip {
		p.Ship = p.MaxShip
	}

	p.Spade += receive.SpadeUpgrade

	if p.Spade > p.MaxSpade {
		p.Spade = p.MaxSpade
	}

	p.VP += receive.DVP * (p.Building[D] + p.Building[WHITE_D])
	p.VP += receive.TpVP * (p.Building[TP] + p.Building[WHITE_TP])
	p.VP += receive.TeVP * (p.Building[TE] + p.Building[WHITE_TE])
	p.VP += receive.ShVP * (p.Building[SH] + p.Building[SA] + p.Building[WHITE_SH] + p.Building[WHITE_SA])
	p.VP += receive.CityVP * p.City

	if receive.ScienceVP > 0 {
		min := 999
		for _, v := range p.Science {
			if v < min {
				min = v
			}
		}
		p.VP += min
	}

	if receive.Spade > 0 {
		p.ExtraBuild = 1
	}

	p.Resource.Power[2] += receive.Token
	p.ReceivePower(receive.Power, false)
}

func (p *Faction) FirstIncome() {
	for i, v := range p.Incomes {
		for j := 0; j <= p.Building[i]; j++ {
			p.ReceiveResource(v[j])
		}
	}

	for _, v := range p.Tiles {
		if v.Category != TileRound {
			p.ReceiveResource(v.Receive)
		}
	}

	if p.Resource.Prist > p.MaxPrist {
		p.Resource.Prist = p.MaxPrist
	}

	if p.Resource.Bridge > p.MaxBridge {
		p.Resource.Bridge = p.MaxBridge
	}
}

func (p *Faction) Income() {
	for i, v := range p.Incomes {
		for j := 0; j <= p.Building[i]; j++ {
			p.ReceiveResource(v[j])
		}
	}

	for _, v := range p.Tiles {
		p.ReceiveResource(v.Receive)
	}

	if p.Resource.Prist > p.MaxPrist {
		p.Resource.Prist = p.MaxPrist
	}

	if p.Resource.Bridge > p.MaxBridge {
		p.Resource.Bridge = p.MaxBridge
	}

	p.IsPass = false
}

func (p *Faction) Burn(count int) error {
	if p.Resource.Power[1] < count*2 {
		return errors.New("not enough power")
	}

	p.Resource.Power[1] -= count * 2
	p.Resource.Power[2] += count

	return nil
}

func (p *Faction) ConvertPower(convert Price) error {
	need := convert.Coin + convert.Worker*3 + convert.Prist*5 + convert.Book.Any*5

	err := p.UsePower(need)
	if err != nil {
		return err
	}

	p.ReceiveResource(convert)

	return nil
}

func (p *Faction) UsePower(value int) error {
	if p.Resource.Power[2] < value && p.Resource.Power[2]+p.Resource.Power[1]/2 < value {
		return errors.New("not enough power")
	}

	if p.Resource.Power[2] >= value {
		p.Resource.Power[2] -= value
		p.Resource.Power[0] += value
	} else {
		value -= p.Resource.Power[2]
		p.Resource.Power[0] += p.Resource.Power[2]
		p.Resource.Power[2] = 0
		p.Resource.Power[1] -= value * 2
		p.Resource.Power[0] += value
	}

	return nil
}

func (p *Faction) Print() {
	/*
		extraShip := ""

		log.Printf("%v: %v C, %v W, %v/%v P, %v/%v/%v/%v B, %v/%v/%v pw, dig level: %v/%v, ship level: %v%v/%v\n",
			p.Ename,
			p.Resource.Coin,
			p.Resource.Worker,
			p.Resource.Prist, p.MaxPrist,
			p.Resource.Book.Banking, p.Resource.Book.Law, p.Resource.Book.Engineering, p.Resource.Book.Medicine,
			p.Resource.Power[0], p.Resource.Power[1], p.Resource.Power[2],
			p.Spade, p.MaxSpade, p.Ship, extraShip, p.MaxShip)
	*/
}

func (p *Faction) UsePrice(need Price) {
	p.Resource.Coin -= need.Coin
	p.Resource.Worker -= need.Worker
	p.Resource.Prist -= need.Prist
	p.Resource.Spade -= need.Spade
	p.Resource.Bridge -= need.Bridge

	p.Resource.Book.Banking -= need.Book.Banking
	p.Resource.Book.Law -= need.Book.Law
	p.Resource.Book.Engineering -= need.Book.Engineering
	p.Resource.Book.Medicine -= need.Book.Medicine

	if p.Resource.Power[2] >= need.Power {
		p.Resource.Power[2] -= need.Power
		p.Resource.Power[0] += need.Power
	} else {
		need.Power -= p.Resource.Power[2]
		p.Resource.Power[2] = 0
		p.Resource.Power[1] -= need.Power * 2
		p.Resource.Power[0] += need.Power
	}

	p.Print()
}

func (p *Faction) AdvanceShip() error {
	if p.Ship == p.MaxShip {
		return errors.New("max ship")
	}

	err := CheckResource(p.Resource, p.AdvanceShipPrice)

	if err != nil {
		return err
	}

	p.UsePrice(p.AdvanceShipPrice)

	if p.Color == color.Blue {
		if p.Ship == 1 {
			p.ReceiveResource(Price{VP: 3})
		} else if p.Ship == 2 {
			p.ReceiveResource(Price{Book: Book{Any: 2}})
		}
	} else {
		if p.Ship == 0 {
			p.ReceiveResource(Price{VP: 2})
		} else if p.Ship == 1 {
			p.ReceiveResource(Price{Book: Book{Any: 2}})
		} else if p.Ship == 2 {
			p.ReceiveResource(Price{VP: 4})
		}
	}

	p.Ship++

	p.Action = true

	p.Print()

	return nil
}

func (p *Faction) AdvanceSpade() error {
	if p.Spade == p.MaxSpade {
		return errors.New("max spade")
	}

	err := CheckResource(p.Resource, p.AdvanceSpadePrice)

	if err != nil {
		return err
	}

	p.UsePrice(p.AdvanceSpadePrice)

	if p.Spade == 0 {
		p.ReceiveResource(Price{Book: Book{Any: 2}})
	} else if p.Spade == 1 {
		p.ReceiveResource(Price{VP: 6})
	}

	p.Spade++

	p.Action = true

	p.Print()

	return nil
}

func (p *Faction) FirstBuild(x int, y int, building Building) error {
	if p.Action {
		return errors.New("Already completed the action")
	}

	p.Building[building]++
	p.BuildingList = append(p.BuildingList, Position{X: x, Y: y, Building: building})

	if building == SA {
		p.Resource.SchoolTile++
	}

	p.Action = true

	return nil
}

func (p *Faction) Build(x int, y int, needSpade int, building Building) error {
	if p.Action {
		if p.ExtraBuild == 0 {
			if p.Resource.Building == None {
				return errors.New("Already completed the action")
			} else {
				if p.Resource.Building != building {
					return errors.New("not match building")
				}
			}
		} else {
			if len(p.DigPosition) > 0 {
				if p.DigPosition[0].X != x || p.DigPosition[0].Y != y {
					return errors.New("must build first dig position")
				}
			}

			p.ExtraBuild--
		}
	}

	if p.MaxBuilding[building] <= p.Building[building] {
		return errors.New("full building")
	}

	if building == D {
		err := CheckResource(p.Resource, p.Price[D])
		if err != nil {
			return err
		}
	}

	if p.Resource.Spade >= needSpade {
		if p.Resource.Spade-p.Resource.ConvertSpade >= needSpade {
			p.Resource.Spade -= needSpade
		} else {
			p.Resource.Spade = 0
		}
	} else {
		if p.Resource.Worker < p.GetWorkerForSpade()*(needSpade-p.Resource.Spade) {
			return errors.New("not enough spade")
		}

		p.Resource.Worker -= p.GetWorkerForSpade() * (needSpade - p.Resource.Spade)
	}

	if building >= WHITE_D {
		p.Resource.Spade = 0
	}

	p.Resource.ConvertSpade = 0

	if building == D {
		p.UsePrice(p.Price[D])
		p.ReceiveDVP()
	} else if building == WHITE_D {
		p.ReceiveDVP()
	} else if building == WHITE_TP {
		p.ReceiveTpVP()
	} else if building == WHITE_TE {
		p.ReceiveTeVP()
	} else if building == WHITE_SH || building == WHITE_SA {
		p.ReceiveShsaVP()
	}

	if p.Type == TileFactionGoblins {
		p.ReceiveResource(Price{Coin: needSpade * 2})
	}

	p.Building[building]++

	p.BuildingList = append(p.BuildingList, Position{X: x, Y: y, Building: building})

	p.Resource.Building = None

	p.Action = true
	p.BuildAction = true

	p.Print()

	return nil
}

func (p *Faction) ReceiveDVP() {
	for _, v := range p.Tiles {
		p.VP += v.Build.D
	}
}

func (p *Faction) ReceiveTpVP() {
	for _, v := range p.Tiles {
		p.VP += v.Build.TP
	}
}

func (p *Faction) ReceiveTeVP() {
	for _, v := range p.Tiles {
		p.VP += v.Build.TE
	}
}

func (p *Faction) ReceiveShsaVP() {
	for _, v := range p.Tiles {
		p.VP += v.Build.SHSA
	}
}

func (p *Faction) ReceiveEdgeVP() {
	for _, v := range p.Tiles {
		p.VP += v.Build.Edge
	}
}

func (p *Faction) ReceiveRiverVP() {
	for _, v := range p.Tiles {
		p.VP += v.Build.River
	}
}

func (p *Faction) ReceivePristVP() {
	for _, v := range p.Tiles {
		p.VP += v.Build.Prist
	}
}

func (p *Faction) Upgrade(x int, y int, target Building, extra int) error {
	current := None

	for _, item := range p.BuildingList {
		if item.X == x && item.Y == y {
			current = item.Building
			if current == None || current == SH || current == SA || int(current) >= int(WHITE_D) {
				return errors.New("can not upgrade")
			}
		}
	}

	if current == None {
		return errors.New("not found building")
	}

	if p.MaxBuilding[target] <= p.Building[target] {
		return errors.New("full building")
	}

	if target == TP && p.Resource.TpUpgrade > 0 {
	} else {
		price := p.Price[target]
		price.Coin += extra

		err := CheckResource(p.Resource, price)
		if err != nil {
			return err
		}
		p.UsePrice(price)
	}

	p.Building[current]--
	p.Building[target]++

	for i, item := range p.BuildingList {
		if item.X == x && item.Y == y {
			// 요소 삭제
			p.BuildingList = append(p.BuildingList[:i], p.BuildingList[i+1:]...)
			break
		}
	}

	p.BuildingList = append(p.BuildingList, Position{X: x, Y: y, Building: target})

	if target == SH {
		p.Resource.PalaceTile++
	} else if target == TE || target == SA {
		p.Resource.SchoolTile++
	}

	if target == TP {
		p.ReceiveTpVP()
	} else if target == SH || target == SA {
		p.ReceiveShsaVP()
	}

	p.Action = true
	p.ResetResource()

	p.Print()

	return nil
}

func (p *Faction) Downgrade(x int, y int) error {
	current := None

	for _, item := range p.BuildingList {
		if item.X == x && item.Y == y {
			current = item.Building
			if current != TE {
				return errors.New("can not upgrade")
			}
		}
	}

	if current == None {
		return errors.New("not found building")
	}

	if p.MaxBuilding[TP] <= p.Building[TP] {
		return errors.New("full building")
	}

	p.Building[TE]--
	p.Building[TP]++

	for i, item := range p.BuildingList {
		if item.X == x && item.Y == y {
			// 요소 삭제
			p.BuildingList = append(p.BuildingList[:i], p.BuildingList[i+1:]...)
			break
		}
	}

	p.BuildingList = append(p.BuildingList, Position{X: x, Y: y, Building: TP})

	p.ReceiveTpVP()

	p.Resource.Downgrade--
	p.Action = true
	p.ResetResource()

	p.Print()

	return nil
}

func (p *Faction) SendScholar(pos int, inc int) error {
	p.Science[pos] += inc
	p.Resource.Prist--
	p.MaxPrist--

	p.Action = true
	p.ResetResource()

	p.ReceivePristVP()

	p.Print()

	return nil
}

func (p *Faction) SupployScholar(pos int, inc int) error {
	p.Science[pos] += inc
	p.Resource.Prist--

	p.Action = true
	p.ResetResource()

	p.ReceivePristVP()

	p.Print()

	return nil
}

func (p *Faction) PowerAction(item action.PowerActionItem) error {
	err := p.UsePower(item.Power)
	if err != nil {
		return err
	}

	p.ResetResource()

	p.ReceiveResource(item.Receive)

	p.Action = true

	p.Print()

	return nil
}

func (p *Faction) Book(item action.BookActionItem, book Book) error {
	if p.Resource.Book.Banking < book.Banking ||
		p.Resource.Book.Law < book.Law ||
		p.Resource.Book.Engineering < book.Engineering ||
		p.Resource.Book.Medicine < book.Medicine {
		return errors.New("not enough book")
	}

	p.Resource.Book.Banking -= book.Banking
	p.Resource.Book.Law -= book.Law
	p.Resource.Book.Engineering -= book.Engineering
	p.Resource.Book.Medicine -= book.Medicine

	p.ReceiveResource(item.Receive)

	p.Action = true

	p.Print()

	return nil
}

func (p *Faction) Bridge(x1 int, y1 int, x2 int, y2 int) error {
	if p.Resource.Bridge == 0 {
		return errors.New("not have bridge")
	}

	p.Resource.Bridge--
	p.MaxBridge--

	p.BridgeList = append(p.BridgeList, BridgePosition{X1: x1, Y1: y1, X2: x2, Y2: y2})

	p.Print()

	return nil
}

func (p *Faction) ResetResource() {
	p.Resource.Science = Science{}
	p.Resource.Spade = 0
	p.Resource.ConvertSpade = 0
	p.Resource.Bridge = 0
	p.Resource.TpUpgrade = 0
	p.Resource.Building = None
	p.ExtraBuild = 0
	p.DigPosition = make([]Position, 0)
}

func (p *Faction) Pass(tile TileItem) (error, TileItem) {
	p.ResetResource()

	for i, v := range p.Tiles {
		if v.Type == TileRoundSchoolScienceCoin {
			p.Resource.Science.Any += p.Building[TE] + p.Building[WHITE_TE]
			v.Pass.Science.Any = 0
		}

		if v.Category == TileRound {
			p.ReceiveResource(v.Pass)
		}
		p.Tiles[i].Use = false

	}

	var ret TileItem

	for i, v := range p.Tiles {
		if v.Category == TileRound {
			ret = v
			p.Tiles = append(p.Tiles[:i], p.Tiles[i+1:]...)
			break
		}
	}

	p.Resource.Coin += tile.Coin
	tile.Coin = 0

	if tile.Name != "" {
		p.Tiles = append(p.Tiles, tile)
	}

	p.IsPass = true
	p.Action = true

	p.Print()

	return nil, ret
}

func (p *Faction) ReceiveCity(item CityItem) error {
	p.Cities = append(p.Cities, item)

	p.ResetResource()

	p.ReceiveResource(item.Receive)
	p.Resource.City--
	p.City++
	p.Key++

	return nil
}

func (p *Faction) Dig(x int, y int, dig int) error {
	if p.Resource.Spade < dig {
		return errors.New("not enough spade")
	}

	p.Resource.Spade -= dig

	p.DigPosition = append(p.DigPosition, Position{X: x, Y: y})

	return nil
}

func (p *Faction) ConvertDig(spade int) error {
	if p.Action == true && p.ExtraBuild == 0 && p.Resource.Building == None {
		return errors.New("already action end")
	}

	if p.Resource.Worker < p.GetWorkerForSpade()*spade {
		return errors.New("not enough worker")
	}

	p.Resource.Worker -= p.GetWorkerForSpade() * spade
	p.Resource.Spade += spade
	p.Resource.ConvertSpade += spade

	return nil
}

func (p *Faction) TurnEnd(round int, pass bool) error {
	if p.Action == false {
		return errors.New("must action")
	}

	if p.Resource.Downgrade > 0 {
		return errors.New("did not act on the downgrade")
	}

	if pass == true && p.IsPass == true {
		for i, v := range p.Tiles {
			if v.Category == TileRound {
				continue
			}

			p.ReceiveResource(v.Pass)
			p.Tiles[i].Use = false
		}
	}

	p.Action = false
	p.BuildAction = false

	if round > 0 {
		p.ResetResource()
	}

	return nil
}

func (p *Faction) PalaceTile(tile TileItem) error {
	if p.Resource.PalaceTile == 0 {
		return errors.New("not have palace tile")
	}

	for _, v := range p.Tiles {
		if v.Category == TilePalace {
			return errors.New("already")
		}
	}

	tile.Color = color.None
	p.Tiles = append(p.Tiles, tile)

	p.ReceiveResource(tile.Once)

	p.Resource.PalaceTile--

	if tile.Type == TilePalace6PowerCity {
		p.TownPower--
	}

	p.Action = true

	return nil
}

func (p *Faction) SchoolTile(tile TileItem, science int) error {
	if p.Resource.SchoolTile == 0 {
		return errors.New("not have school tile")
	}

	for _, v := range p.Tiles {
		if v.Type == tile.Type {
			return errors.New("already")
		}
	}

	p.Tiles = append(p.Tiles, tile)

	p.ReceiveResource(tile.Once)

	p.Resource.SchoolTile--

	p.Action = true

	return nil
}

func (p *Faction) RoundTile(tile TileItem) error {
	for _, v := range p.Tiles {
		if v.Type == tile.Type {
			return errors.New("already")
		}
	}

	p.Tiles = append(p.Tiles, tile)

	p.ReceiveResource(tile.Once)

	return nil
}

func (p *Faction) TileAction(category TileCategory, pos int) error {
	var tile *TileItem

	tilePos := 0

	if category == TilePalace {
		tilePos = pos
	} else if category == TileRound {
		tilePos = int(TilePalaceVp) + pos
	} else if category == TileSchool {
		tilePos = int(TileRoundCoin) + pos
	} else if category == TileFaction {
		tilePos = int(TileSchoolPassPrist) + pos
	} else if category == TileColor {
		tilePos = int(TileFactionPsychics) + pos
	} else if category == TileInnovation {
		tilePos = int(TileColorRed) + pos
	}

	find := -1
	for i, v := range p.Tiles {
		if v.Category == category && v.Type == TileType(tilePos) {
			find = i
			break
		}
	}

	if find == -1 {
		return errors.New("not found")
	}

	tile = &p.Tiles[find]

	if tile.Use == true {
		return errors.New("already")
	}

	p.ReceiveResource(tile.Action)

	tile.Use = true

	p.Action = true

	return nil
}

func (p *Faction) Convert(source Price, target Price) error {
	if source.Prist > 0 {
		if p.Resource.Prist < source.Prist || source.Prist < target.Worker+target.Coin {
			return errors.New("not enough prist")
		}
	} else if source.Worker > 0 {
		if p.Resource.Worker < source.Worker || source.Worker < target.Coin {
			return errors.New("not enough prist")
		}
	} else if source.Book.Count() > 0 {
		if p.Resource.Book.Count() < source.Book.Count() || source.Book.Count() < target.Coin {
			return errors.New("not enough prist")
		}
	} else if source.Power > 0 {
		need := target.Coin + target.Worker*3 + target.Prist*5 + target.Book.Any*5
		if p.Resource.Power[2] < source.Power || source.Power < need {
			return errors.New("not enough power")
		}
	}

	p.UsePrice(source)
	p.ReceiveResource(target)

	return nil
}

func (p *Faction) Annex(x int, y int) error {
	if p.Action == true {
		return errors.New("already action end")
	}

	if p.Resource.Annex == 0 {
		return errors.New("not have annex")
	}

	p.Resource.Annex--

	p.AnnexList = append(p.AnnexList, Position{X: x, Y: y})

	p.ResetResource()
	p.Action = true

	p.Print()

	return nil
}

func (p *Faction) GetScience(pos int) int {
	return p.Science[pos]
}

func (p *Faction) IncScience(pos int, level int) {
	p.Science[pos] += level
}

func (p *Faction) InnovationTile(tile TileItem, book Book) error {
	count := 0

	for _, v := range p.Tiles {
		if v.Category == TileInnovation {
			count++
		}

		if v.Type == tile.Type {
			return errors.New("already")
		}
	}

	need := 5
	if count >= 3 {
		return errors.New("full")
	} else if count == 2 {
		need += 2
	} else if count == 1 {
		if p.Color != color.Red {
			need++
		}
	}

	if book.Count() < need {
		return errors.New("not enough book")
	}

	price := Price{Book: book}
	if p.Building[SH] == 0 && p.Building[WHITE_SH] == 0 {
		if p.Resource.Coin < 5 {
			return errors.New("not enough coin")
		}

		price.Coin = 5
	}

	p.Tiles = append(p.Tiles, tile)

	p.ResetResource()
	p.ReceiveResource(tile.Once)
	p.UsePrice(price)

	if tile.Type == TileInnovationKind {
		cnt := 0
		for _, v := range p.Building[:len(p.Building)-1] {
			if v > 0 {
				cnt++
			}
		}

		p.ReceiveResource(Price{Science: Science{Any: cnt}})
	} else if tile.Type == TileInnovationCount {
		cnt := 0
		for _, v := range p.Building {
			cnt += v
		}

		vp := 0
		if cnt >= 11 {
			vp = 18
		} else if cnt >= 9 {
			vp = 12
		} else if cnt >= 7 {
			vp = 8
		}
		p.ReceiveResource(Price{VP: vp})
	} else if tile.Type == TileInnovationScience {
		items := make([]int, 0)
		for _, v := range p.Science {
			items = append(items, v)
		}

		sort.Sort(sort.Reverse(sort.IntSlice(items)))

		p.ReceiveResource(Price{VP: items[0] + items[1]})
	} else if tile.Type == TileInnovationCluster {
	} else if tile.Type == TileInnovationBridge {
		vps := []int{0, 8, 12, 18}
		p.ReceiveResource(Price{VP: vps[len(p.BridgeList)]})
	}

	p.Action = true

	return nil
}

func (p *Faction) ReceiveIncome(receive Price) {
	p.Receive.Coin += receive.Coin
	p.Receive.Worker += receive.Worker
	p.Receive.Prist += receive.Prist

	p.Receive.Book.Any += receive.Book.Any
	p.Receive.Book.Banking += receive.Book.Banking
	p.Receive.Book.Law += receive.Book.Law
	p.Receive.Book.Engineering += receive.Book.Engineering
	p.Receive.Book.Medicine += receive.Book.Medicine

	p.Receive.Science.Any += receive.Science.Any
	p.Receive.Science.Single += receive.Science.Single

	if p.Receive.Prist > p.MaxPrist {
		p.Receive.Prist = p.MaxPrist
	}

	p.Receive.Power += receive.Power
}

func (p *Faction) CalulateReceive() {
	p.Receive = Price{}
	p.Receive.VP = p.VP

	for i, v := range p.Incomes {
		for j := 0; j <= p.Building[i]; j++ {
			p.ReceiveIncome(v[j])
		}
	}

	for _, v := range p.Tiles {
		if p.IsPass == false && v.Category == TileRound {
			continue
		}

		p.ReceiveIncome(v.Receive)
	}
}

func (p *Faction) CalulateVP() {
	p.Receive = Price{}
	p.Receive.VP = p.VP

	if p.IsPass == true {
		return
	}

	for _, v := range p.Tiles {
		receive := v.Pass

		p.Receive.VP += receive.VP
		p.Receive.VP += receive.DVP * (p.Building[D] + p.Building[WHITE_D])
		p.Receive.VP += receive.TpVP * (p.Building[TP] + p.Building[WHITE_TP])
		p.Receive.VP += receive.TeVP * (p.Building[TE] + p.Building[WHITE_TE])
		p.Receive.VP += receive.ShVP * (p.Building[SH] + p.Building[SA] + p.Building[WHITE_SH] + p.Building[WHITE_SA])
		p.Receive.VP += receive.CityVP * p.City

		if receive.ScienceVP > 0 {
			min := 999
			for _, v := range p.Science {
				if v < min {
					min = v
				}
			}
			p.Receive.VP += min
		}
	}
}

func (p *Faction) CheckTile(tile TileType) bool {
	for _, v := range p.Tiles {
		if v.Type == tile {
			return true
		}
	}

	return false
}

func (p *Faction) AddTile(tile TileItem) bool {
	for _, v := range p.Tiles {
		if v.Category == tile.Category {
			return false
		}

		if v.Type == tile.Type {
			return false
		}
	}

	p.Tiles = append(p.Tiles, tile)

	return true
}
