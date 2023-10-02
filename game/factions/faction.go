package factions

import (
	"aoi/game/action"
	"aoi/game/color"
	. "aoi/game/resources"
	"aoi/game/resources/city"
	"errors"
	"fmt"

	"log"
	"math"
)

type FactionInterface interface {
	GetInstance() *Faction

	Init()
	Print()
	Income()
	PassIncome()

	FirstBuild(x int, y int)
	Build(x int, y int, needSpade int) error
	Upgrade(x int, y int, target Building) error
	AdvanceShip() error
	AdvanceSpade() error
	SendScholar() error
	SupployScholar() error
	PowerAction(item action.PowerActionItem) error
	Book(item action.BookActionItem) error
	Bridge(x1 int, y1 int, x2 int, y2 int) error
	Pass(tile *TileItem) error
	ReceiveCity(item city.CityItem) error
	Dig(dig int) error
	ConvertDig(spade int) error
	TurnEnd() error
	PalaceTile(tile TileItem) error
	SchoolTile(tile TileItem) error
	TileAction(category TileCategory, pos int) error
}

type Faction struct {
	Name              string           `json:"name"`
	Ename             string           `json:"ename"`
	Color             color.Color      `json:"color"`
	Resource          Resource         `json:"resource"`
	Tiles             []TileItem       `json:"tiles"`
	MaxBuilding       [6]int           `json:"maxBuilding"`
	Building          [6]int           `json:"building"`
	BuildingPower     [6]int           `json:"buildingPower"`
	BuildingList      []Position       `json:"buildingList"`
	BridgeList        []BridgePosition `json:"bridgeList"`
	Price             [6]Price         `json:"price"`
	AdvanceShipPrice  Price            `json:"advanceShipPrice"`
	AdvanceSpadePrice Price            `json:"advanceSpadePrice"`
	Incomes           [][]Price        `json:"incomes"`
	Point             int              `json:"point"`
	TownPower         int              `json:"townPower"`
	Spade             int              `json:"spade"`
	MaxSpade          int              `json:"maxSpade"`
	Ship              int              `json:"ship"`
	MaxShip           int              `json:"maxShip"`
	MaxPrist          int              `json:"maxPrist"`
	Science           []int            `json:"science"`
	Key               int              `json:"key"`
	RoundTile         *TileItem        `json:"roundTile"`
	Action            bool             `json:"action"`
	MaxBridge         int              `json:"maxBridge"`
	ExtraBuild        int              `json:"extraBuild"`
	VP                int              `json:"vp"`
	City              int              `json:"city"`
	IsPass            bool             `json:"-"`
}

func NewFaction(name string, ename string, color color.Color) *Faction {
	var item Faction

	item.Name = name
	item.Ename = ename
	item.Color = color

	log.Printf("color = %v\n", color.ToString())

	item.Resource.Coin = 15
	item.Resource.Worker = 4
	item.Resource.Power = [3]int{2, 4, 0}
	item.VP = 20
	item.Spade = 0
	item.MaxSpade = 2
	item.Ship = 0
	item.MaxShip = 3
	item.MaxPrist = 7
	item.Science = []int{0, 0, 2, 0}
	item.MaxBridge = 3
	item.Key = 0
	item.Action = false
	item.ExtraBuild = 0
	item.RoundTile = &TileItem{}
	item.IsPass = false

	item.Tiles = make([]TileItem, 0)
	item.MaxBuilding = [6]int{0, 9, 4, 3, 1, 1}
	item.Building = [6]int{0, 0, 0, 0, 0, 0}
	item.BuildingPower = [6]int{0, 1, 2, 2, 3, 3}
	item.Price = [6]Price{
		Price{Worker: 0, Coin: 0},
		Price{Worker: 1, Coin: 2},
		Price{Worker: 2, Coin: 3},
		Price{Worker: 3, Coin: 5},
		Price{Worker: 4, Coin: 6},
		Price{Worker: 6, Coin: 6},
	}

	item.AdvanceShipPrice = Price{Prist: 1, Coin: 4}
	item.AdvanceSpadePrice = Price{Prist: 1, Worker: 1, Coin: 5}

	item.Incomes = [][]Price{
		[]Price{Price{}},
		[]Price{Price{Worker: 1}, Price{Worker: 1}, Price{Worker: 1}, Price{Worker: 1}, Price{Worker: 1}, Price{}, Price{Worker: 1}, Price{Worker: 1}, Price{Worker: 1}, Price{Worker: 1}},
		[]Price{Price{}, Price{Coin: 2, Power: 1}, Price{Coin: 2, Power: 1}, Price{Coin: 2, Power: 2}, Price{Coin: 2, Power: 2}},
		[]Price{Price{}, Price{Prist: 1}, Price{Prist: 1}, Price{Prist: 1}},
		[]Price{Price{}, Price{Power: 4}},
		[]Price{Price{}, Price{Prist: 1}},
	}

	item.Point = 10
	item.TownPower = 7

	item.BuildingList = make([]Position, 0)
	item.BridgeList = make([]BridgePosition, 0)

	item.Resource.Coin = 100
	item.Resource.Worker = 100
	item.Resource.Prist = 7
	item.Resource.Book = 100
	item.Resource.Power = [3]int{0, 0, 12}

	return &item
}

func (p *Faction) GetShipDistance() int {
	return p.Ship + p.RoundTile.Ship + 1
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
			c = p.Resource.Power[0]
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
	log.Println("ReceiveResource", receive)
	p.Resource.Coin += receive.Coin
	p.Resource.Worker += receive.Worker
	p.Resource.Prist += receive.Prist
	p.Resource.Book += receive.Book
	p.Resource.Spade += receive.Spade

	p.Resource.City += receive.City
	p.Resource.VP += receive.VP

	p.Resource.Bridge += receive.Bridge

	p.Resource.TpUpgrade += receive.TpUpgrade

	if p.Resource.Bridge > p.MaxBridge {
		p.Resource.Bridge = p.MaxBridge
	}

	if p.Resource.Prist > p.MaxPrist {
		p.Resource.Prist = p.MaxPrist
	}

	if receive.Spade > 0 {
		p.ExtraBuild = 1
	}

	p.ReceivePower(receive.Power, false)
}

func (p *Faction) Income() {
	power := 0

	// 종족판

	for i, v := range p.Incomes {
		for j := 0; j <= p.Building[i]; j++ {
			p.ReceiveResource(v[j])
		}
	}

	p.ReceiveResource(p.RoundTile.Receive)
	// 패스타일
	// science
	// sh 타일
	// te 타일
	// innovation 타일

	p.ReceivePower(power, false)

	if p.Resource.Prist > p.MaxPrist {
		p.Resource.Prist = p.MaxPrist
	}

	if p.Resource.Bridge > p.MaxBridge {
		p.Resource.Bridge = p.MaxBridge
	}

	p.IsPass = false
}

func (p *Faction) Burn(count int) bool {
	if p.Resource.Power[1] < count*2 {
		return false
	}

	p.Resource.Power[1] -= count * 2
	p.Resource.Power[2] += count

	return true
}

func (p *Faction) UsePower(value int) {
	if p.Resource.Power[2] < value {
		if p.Burn(value) == false {
			return
		}
	}

	p.Resource.Power[0] += value
	p.Resource.Power[2] -= value
}

/*

	getTile(pos: int) {
		var tile = _research.tiles[pos]
		p.tiles.push(clone(tile))


		var item = tile.receive

		if (item != null) {
			if ('worker' in item)
				p.resource.worker += item.worker
			if ('coin' in item)
				p.resource.coin += item.coin
			if ('knowledge' in item)
				p.resource.knowledge += item.knowledge
			if ('qic' in item)
				p.resource.qic += item.qic
			if ('power' in item)
				p.receivePower(item.power)
			if ('token' in item)
				p.resource.power[1] += item.token
			if ('point' in item)
				p.point += item.point
			if ('gaiaPoint' in item)
				p.gaiaPoint += item.gaiaPoint
			if ('buildingPower' in item) {
				for (var i = 0 i < item.buildingPower.length i++) {
					p.buildingPower[i] += item.buildingPower[i]
				}
			}
		}

		p.onResourceChange()
	}


	getTokenCount(): int {
		return p.resource.power[1] + p.resource.power[2] + p.resource.power[3]
	}

	getTransTokenCount(): int {
		var value = [0, 6, 6, 4, 3, 3]

		return value[p.research[ResearchType.FORMER]]
	}

	removeToken(value) {
		var remain = value

		for (var i = 1 i <= 3 i++) {
			if (remain <= p.resource.power[i]) {
				p.resource.power[i] -= remain
				break
			}

			remain -= p.resource.power[i]
			p.resource.power[i] = 0
		}

		p.onResourceChange()
	}

	transforming(x, y): boolean {
		var need = p.getTransTokenCount()
		if (p.resource.former <= 0 || p.getTokenCount() < need)
			return false

		p.resource.former--
		p.removeToken(need)
		p.resource.power[0] += need

		p.onResourceChange()

		return true
	}

	newRound() {
		p.receiveIncome()

		var value = p.resource.power[0]
		p.resource.power[1] += value
		p.resource.power[0] = 0

		p.onResourceChange()
	}
}
*/

func (p *Faction) Print() {
	extraShip := ""
	if p.RoundTile.Ship > 0 {
		extraShip = fmt.Sprintf("+%v", p.RoundTile.Ship)
	}

	log.Printf("%v: %v C, %v W, %v/%v P, %v B, %v/%v/%v pw, dig level: %v/%v, ship level: %v%v/%v\n",
		p.Ename,
		p.Resource.Coin,
		p.Resource.Worker,
		p.Resource.Prist, p.MaxPrist,
		p.Resource.Book,
		p.Resource.Power[0], p.Resource.Power[1], p.Resource.Power[2],
		p.Spade, p.MaxSpade, p.Ship, extraShip, p.MaxShip)
}

func (p *Faction) UsePrice(need Price) {
	p.Resource.Coin -= need.Coin
	p.Resource.Worker -= need.Worker
	p.Resource.Prist -= need.Prist
	p.Resource.Spade -= need.Spade
	p.Resource.Bridge -= need.Bridge

	if p.Resource.Power[2] >= need.Power {
		p.Resource.Power[2] -= need.Power
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
		log.Println(err)
		return err
	}

	p.UsePrice(p.AdvanceShipPrice)

	p.Ship++

	p.Action = true

	p.Print()

	return nil
}

func (p *Faction) AdvanceSpade() error {
	log.Println("Advance Spade:")
	if p.Spade == p.MaxSpade {
		return errors.New("max spade")
	}

	err := CheckResource(p.Resource, p.AdvanceSpadePrice)

	if err != nil {
		log.Println(err)
		return err
	}

	p.UsePrice(p.AdvanceSpadePrice)

	p.Spade++

	p.Action = true

	p.Print()

	return nil
}

func (p *Faction) FirstBuild(x int, y int) {
	p.Building[D]++
	p.BuildingList = append(p.BuildingList, Position{X: x, Y: y, Building: D})
}

func (p *Faction) Build(x int, y int, needSpade int) error {
	if p.Action {
		if p.ExtraBuild == 0 {
			return errors.New("Already completed the action")
		} else {
			p.ExtraBuild--
		}
	}

	if p.MaxBuilding[D] <= p.Building[D] {
		return errors.New("full building")
	}

	err := CheckResource(p.Resource, p.Price[D])
	if err != nil {
		log.Println(err)
		return err
	}

	if p.Resource.Worker < p.GetWorkerForSpade()*(needSpade-p.Resource.Spade)+p.Price[D].Worker {
		return errors.New("not enough spade")
	}

	log.Println(p.Resource.Spade, needSpade)

	if p.Resource.Spade >= needSpade {
		p.Resource.Spade -= needSpade
	} else {
		p.Resource.Worker -= p.GetWorkerForSpade() * (needSpade - p.Resource.Spade)
		p.Resource.Spade = 0
	}

	p.UsePrice(p.Price[D])

	p.Building[D]++
	p.BuildingList = append(p.BuildingList, Position{X: x, Y: y, Building: D})

	p.Action = true

	p.Print()

	return nil
}

func (p *Faction) Upgrade(x int, y int, target Building) error {
	current := None

	for _, item := range p.BuildingList {
		if item.X == x && item.Y == y {
			current = item.Building
			if current == None || current == SH || current == SA {
				return errors.New("can not upgrade")
			}
		}
	}

	if current == None {
		return errors.New("not found building")
	}

	log.Println("max", p.MaxBuilding[target])
	log.Println("current", p.Building[target])
	if p.MaxBuilding[target] <= p.Building[target] {
		log.Println("this this")
		return errors.New("full building")
	}

	log.Println("nonono")

	err := CheckResource(p.Resource, p.Price[target])
	if err != nil {
		log.Println(err)
		return err
	}

	p.Building[current]--
	p.Building[target]++

	p.UsePrice(p.Price[target])

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

	p.Action = true

	p.Print()

	return nil
}

func (p *Faction) SendScholar() error {
	p.Resource.Prist--
	p.MaxPrist--

	p.Action = true

	p.Print()

	return nil
}

func (p *Faction) SupployScholar() error {
	p.Resource.Prist--

	p.Action = true

	p.Print()

	return nil
}

func (p *Faction) PowerAction(item action.PowerActionItem) error {
	p.UsePower(item.Power)
	p.ReceiveResource(item.Receive)

	p.Action = true

	p.Print()

	return nil
}

func (p *Faction) Book(item action.BookActionItem) error {
	p.Resource.Book -= item.Book
	p.ReceiveResource(item.Receive)

	p.Action = true

	p.Print()

	return nil
}

func (p *Faction) Bridge(x1 int, y1 int, x2 int, y2 int) error {
	if p.Resource.Bridge == 0 {
		log.Println("not have bridge")
		return errors.New("not have bridge")
	}

	p.Resource.Bridge--
	p.MaxBridge--

	p.BridgeList = append(p.BridgeList, BridgePosition{X1: x1, Y1: y1, X2: x2, Y2: y2})

	p.Print()

	return nil
}

func (p *Faction) Pass(tile *TileItem) error {
	p.PassIncome()

	if p.RoundTile != nil {
		p.RoundTile.Use = false
	}

	p.RoundTile = tile
	p.RoundTile.Use = false

	p.Resource.Spade = 0
	p.Resource.Bridge = 0
	p.Resource.TpUpgrade = 0

	p.ReceiveResource(p.RoundTile.Receive)

	for i, v := range p.Tiles {
		p.ReceiveResource(v.Receive)
		p.Tiles[i].Use = false
	}

	p.IsPass = true

	p.Action = true

	p.Print()

	return nil
}

func (p *Faction) PassIncome() {
}

func (p *Faction) ReceiveCity(item city.CityItem) error {
	p.ReceiveResource(item.Receive)
	p.Resource.City--
	p.City++
	p.Key++

	return nil
}

func (p *Faction) Dig(dig int) error {
	if p.Resource.Spade < dig {
		return errors.New("not enough spade")
	}

	p.Resource.Spade -= dig

	return nil
}

func (p *Faction) ConvertDig(spade int) error {
	if p.Action == true && p.ExtraBuild == 0 {
		return errors.New("already action end")
	}

	if p.Resource.Worker < p.GetWorkerForSpade()*spade {
		return errors.New("not enough worker")
	}

	p.Resource.Worker -= p.GetWorkerForSpade() * spade
	p.Resource.Spade += spade

	return nil
}

func (p *Faction) TurnEnd() error {
	p.Action = false
	p.Resource.Spade = 0
	p.Resource.Science = Science{}

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
	p.Tiles = append(p.Tiles, tile)

	p.ReceiveResource(tile.Once)

	p.Resource.PalaceTile--

	return nil
}

func (p *Faction) SchoolTile(tile TileItem) error {
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

	return nil
}

func (p *Faction) TileAction(category TileCategory, pos int) error {
	var tile *TileItem

	if category == TileRound {
		if p.RoundTile.Type != TileType(pos) {
			return errors.New("not found")
		}

		tile = p.RoundTile
	} else {
		find := -1
		for i, v := range p.Tiles {
			if v.Category == category && v.Type == TileType(pos) {
				find = i
				break
			}
		}

		if find == -1 {
			return errors.New("not found")
		}

		tile = &p.Tiles[find]
	}

	if tile.Use == true {
		return errors.New("already")
	}

	p.ReceiveResource(tile.Action)

	tile.Use = true
	return nil
}
