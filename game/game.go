package game

import (
	"aoi/game/action"
	"aoi/game/color"
	"aoi/game/factions"
	"aoi/game/resources"
	"errors"
	"log"
)

const (
	BuildRound     = -1
	RoundTileRound = 0
)

type Game struct {
	Map      *Map                        `json:"map"`
	Sciences *Science                    `json:"sciences"`
	Factions []factions.FactionInterface `json:"factions"`

	PowerActions *action.PowerAction   `json:"powerActions"`
	BookActions  *action.BookAction    `json:"bookActions"`
	RoundTiles   *resources.RoundTile  `json:"roundTiles"`
	RoundBonuss  *RoundBonus           `json:"roundBonuss"`
	PalaceTiles  *resources.PalaceTile `json:"palaceTiles"`
	SchoolTiles  *resources.SchoolTile `json:"schoolTiles"`
	Cities       *City                 `json:"cities"`
	Turn         []Turn                `json:"turn"`
	PowerTurn    []Turn                `json:"powerTurn"`
	Round        int                   `json:"round"`
}

type TurnType int

const (
	NormalTurn TurnType = iota
	PowerTurn
	ScienceTurn
)

type Turn struct {
	User    int
	Type    TurnType
	From    int
	Power   int
	Science resources.Science
}

func NewGame() *Game {

	var item Game
	item.PowerActions = action.NewPowerAction()
	item.BookActions = action.NewBookAction()
	item.RoundTiles = resources.NewRoundTile()
	item.RoundBonuss = NewRoundBonus()
	item.PalaceTiles = resources.NewPalaceTile()
	item.SchoolTiles = resources.NewSchoolTile()
	item.Cities = NewCity()

	item.Map = NewMap()
	item.Sciences = NewScience()
	item.Factions = make([]factions.FactionInterface, 0)

	item.Turn = make([]Turn, 0)
	item.PowerTurn = make([]Turn, 0)

	item.Round = -1

	return &item
}

func (p *Game) InitGame() {
	count := len(p.Factions)
	p.PalaceTiles.Init(count)
	p.SchoolTiles.Init(count)
}

func (p *Game) AddFaction(item factions.FactionInterface) {
	item.Init()

	p.Factions = append(p.Factions, item)

	faction := item.GetInstance()
	p.Sciences.AddUser(faction.Color, faction.Science)
}

func (p *Game) IsTurn(user int) bool {
	if len(p.Turn) == 0 {
		return false
	}

	if p.Turn[0].User != user {
		return false
	}

	return true
}

func (p *Game) IsNormalTurn(user int) bool {
	if len(p.Turn) == 0 {
		return false
	}

	if p.Turn[0].User != user {
		return false
	}

	if p.Turn[0].Type != NormalTurn {
		return false
	}

	return true
}

func (p *Game) BuildStart() {
	p.RoundTiles.Init(len(p.Factions))

	for i, _ := range p.Factions {
		p.Turn = append(p.Turn, Turn{User: i, Type: NormalTurn})
	}

	for i, _ := range p.Factions {
		p.Turn = append(p.Turn, Turn{User: len(p.Factions) - i - 1, Type: NormalTurn})
	}

	log.Println(p.Turn)
}

func (p *Game) Start() {
	log.Println("Start")
	p.RoundTiles.Start()

	p.Round++

	for i, v := range p.Factions {
		faction := v.GetInstance()
		v.Income()
		p.Sciences.RoundBonus(faction)
		p.RoundBonus(faction)
		p.Turn = append(p.Turn, Turn{User: i, Type: NormalTurn})
	}

	p.PowerActions.Start()
	p.BookActions.Start()

	log.Println("round", p.Round)

}

func (p *Game) TurnEnd(user int) {
	log.Println("TurnEnd")
	faction := p.Factions[user].GetInstance()
	faction.TurnEnd()

	turnType := p.Turn[0].Type

	p.Turn = p.Turn[1:]

	p.Turn = append(p.PowerTurn, p.Turn...)

	if p.Round > 0 {
		if turnType != PowerTurn && user >= 0 {
			if !faction.IsPass {
				p.Turn = append(p.Turn, Turn{User: user, Type: NormalTurn})
			}
		}
	}

	if len(p.Turn) == 0 {
		if p.Round == BuildRound {
			for i, _ := range p.Factions {
				p.Turn = append(p.Turn, Turn{User: len(p.Factions) - i - 1, Type: NormalTurn})
			}

			p.Round = RoundTileRound
		} else if p.Round == RoundTileRound {
			p.Start()
		} else {
			p.RoundEnd()
			p.Start()
		}
	}
}

func (p *Game) RoundEnd() {
	for _, v := range p.Factions {
		faction := v.GetInstance()
		p.Sciences.RoundEndBonus(faction, p.RoundBonuss.Items[p.Round])
	}
}

func (p *Game) IsPowerTurn() bool {
	if len(p.Turn) == 0 {
		return false
	}

	if p.Turn[0].Type != PowerTurn {
		return false
	}

	return true
}

func (p *Game) IsScienceTurn() bool {
	if len(p.Turn) == 0 {
		return false
	}

	if p.Turn[0].Type != ScienceTurn {
		return false
	}

	return true
}

func (p *Game) FirstBuild(user int, x int, y int) error {
	if p.Round != BuildRound {
		log.Println("round error : FirstBuild")
		return errors.New("round error : FirstBuild")
	}

	if !p.IsTurn(user) {
		log.Println("It's not a turn", p.Turn, user)
		return errors.New("It's not a turn")
	}

	faction := p.Factions[user].GetInstance()

	color := p.Map.GetType(x, y)

	if color != faction.Color {
		log.Printf("color error : map = %v, faction = %v\n", color.ToString(), faction.Color.ToString())
		return errors.New("It's not a user's land")
	}

	log.Println("build success")
	faction.FirstBuild(x, y)

	p.Map.Build(x, y, faction.Color, resources.D)

	p.TurnEnd(user)

	return nil
}

func (p *Game) Build(user int, x int, y int) error {
	if p.Round < 1 {
		log.Println("round error")
		return errors.New("round error")
	}

	if !p.IsTurn(user) || p.IsPowerTurn() || p.IsScienceTurn() {
		log.Println("It's not a turn", p.Turn, user)
		return errors.New("It's not a turn")
	}

	faction := p.Factions[user].GetInstance()

	flag := p.Map.CheckDistance(faction.Color, faction.GetShipDistance(), x, y)

	if flag == false {
		log.Println("ship distance error")
		return errors.New("ship distance error")
	}

	err := p.Map.CheckBuild(x, y, faction.Color, faction.GetSpadeCount())

	if err != nil {
		log.Println("map check build error")
		log.Println(err)
		return err
	}

	needSpade := p.Map.GetNeedSpade(x, y, faction.Color)
	if needSpade < 0 {
		needSpade = 0
	}

	err = faction.Build(x, y, needSpade)
	if err != nil {
		log.Println(err)
		return err
	}

	p.Map.Build(x, y, faction.Color, resources.D)

	if p.Map.CheckCity(faction.Color, x, y) == true {
		faction.Resource.City++
	}

	p.PowerDiffusion(user, x, y)

	return nil
}

func (p *Game) Upgrade(user int, x int, y int, target resources.Building) error {
	log.Println("Upgrade:")

	if p.Round < 1 {
		log.Println("round error")
		return errors.New("round error")
	}

	if !p.IsTurn(user) || p.IsPowerTurn() {
		log.Println("It's not a turn", p.Turn, user)
		return errors.New("It's not a turn")
	}

	faction := p.Factions[user].GetInstance()
	if faction.Action {
		return errors.New("Already completed the action")
	}

	log.Println(p.Map.GetOwner(x, y))
	log.Println(faction.Color)
	if p.Map.GetOwner(x, y) != faction.Color {
		log.Println("not owner")
		return errors.New("not owner")
	}

	err := faction.Upgrade(x, y, target)

	if err != nil {
		return err
	}

	p.Map.SetBuilding(x, y, target)

	if p.Map.CheckCity(faction.Color, x, y) == true {
		faction.Resource.City++
	}

	p.PowerDiffusion(user, x, y)

	return nil
}

func (p *Game) PowerAction(user int, pos int) error {
	if p.Round < 1 {
		log.Println("round error")
		return errors.New("round error")
	}

	if !p.IsTurn(user) || p.IsPowerTurn() {
		log.Println("It's not a turn", p.Turn, user)
		return errors.New("It's not a turn")
	}

	faction := p.Factions[user].GetInstance()
	if faction.Action {
		return errors.New("Already completed the action")
	}

	if p.PowerActions.IsUse(pos) {
		log.Println("already")
		return errors.New("already")
	}

	have := faction.GetHavePowerCount()
	if have < p.PowerActions.GetNeedPower(pos) {
		log.Println("not enough power")
		return errors.New("not enough power")
	}

	item := p.PowerActions.Action(pos)
	faction.PowerAction(item)

	return nil
}

func (p *Game) BookAction(user int, pos int) error {
	if p.Round < 1 {
		log.Println("round error")
		return errors.New("round error")
	}

	if !p.IsTurn(user) || p.IsPowerTurn() {
		log.Println("It's not a turn", p.Turn, user)
		return errors.New("It's not a turn")
	}

	faction := p.Factions[user].GetInstance()
	if faction.Action {
		return errors.New("Already completed the action")
	}

	if pos < 0 || pos > 2 {
		log.Println("not found")
		return errors.New("not found")
	}

	if p.BookActions.IsUse(pos) {
		log.Println("already")
		return errors.New("already")
	}

	have := faction.Resource.Book
	if have < p.BookActions.GetNeedBook(pos) {
		log.Println("not enough book")
		return errors.New("not enough book")
	}

	item := p.BookActions.Action(pos)
	faction.Book(item)

	return nil
}

func (p *Game) AdvanceShip(user int) error {
	if p.Round < 1 {
		log.Println("round error")
		return errors.New("round error")
	}

	if !p.IsTurn(user) || p.IsPowerTurn() {
		log.Println("It's not a turn", p.Turn, user)
		return errors.New("It's not a turn")
	}

	faction := p.Factions[user].GetInstance()
	if faction.Action {
		return errors.New("Already completed the action")
	}

	err := faction.AdvanceShip()
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (p *Game) AdvanceSpade(user int) error {
	if p.Round < 1 {
		log.Println("round error")
		return errors.New("round error")
	}

	if !p.IsTurn(user) || p.IsPowerTurn() {
		log.Println("It's not a turn", p.Turn, user)
		return errors.New("It's not a turn")
	}

	faction := p.Factions[user].GetInstance()
	if faction.Action {
		return errors.New("Already completed the action")
	}

	err := faction.AdvanceSpade()
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (p *Game) SendScholar(user int, pos ScienceType) error {
	if p.Round < 1 {
		log.Println("round error")
		return errors.New("round error")
	}

	if !p.IsTurn(user) || p.IsPowerTurn() {
		log.Println("It's not a turn", p.Turn, user)
		return errors.New("It's not a turn")
	}

	log.Println("pos", pos)

	faction := p.Factions[user].GetInstance()
	if faction.Action {
		return errors.New("Already completed the action")
	}

	if faction.Resource.Prist == 0 {
		return errors.New("not enough prist")
	}

	p.Sciences.Send(faction, pos)

	faction.SendScholar()

	return nil
}

func (p *Game) SupployScholar(user int, pos ScienceType) error {
	if p.Round < 1 {
		log.Println("round error")
		return errors.New("round error")
	}

	if !p.IsTurn(user) || p.IsPowerTurn() {
		log.Println("It's not a turn", p.Turn, user)
		return errors.New("It's not a turn")
	}

	faction := p.Factions[user].GetInstance()
	if faction.Action {
		return errors.New("Already completed the action")
	}

	if faction.Resource.Prist == 0 {
		return errors.New("not enough prist")
	}

	p.Sciences.Supploy(faction, pos)

	faction.SupployScholar()

	return nil
}

func (p *Game) Pass(user int, pos int) error {
	if p.Round < 1 {
		log.Println("round error")
		return errors.New("round error")
	}

	if !p.IsTurn(user) || p.IsPowerTurn() {
		log.Println("It's not a turn", p.Turn, user)
		return errors.New("It's not a turn")
	}

	faction := p.Factions[user].GetInstance()
	roundTile := p.RoundTiles.Pass(pos)
	faction.Pass(roundTile)

	p.TurnEnd(user)

	return nil
}

func (p *Game) FactionAction() {
}

func (p *Game) Dig(user int, x int, y int, dig int) error {
	if p.Round < 1 {
		log.Println("round error")
		return errors.New("round error")
	}

	if !p.IsTurn(user) || p.IsPowerTurn() {
		log.Println("It's not a turn", p.Turn, user)
		return errors.New("It's not a turn")
	}

	faction := p.Factions[user].GetInstance()

	if p.Map.GetType(x, y) == color.River {
		log.Println("can't spade")
		return errors.New("can't spade")
	}

	if faction.Resource.Spade < dig {
		log.Println("not have spade")
		return errors.New("not have spade")
	}

	needSpade := p.Map.GetNeedSpade(x, y, faction.Color)
	if needSpade == 0 {
		log.Println("need not spade")
		return errors.New("need not spade")
	}

	spade := 0
	var change color.Color
	if dig >= needSpade {
		change = faction.Color
		spade = needSpade
	} else {
		target := p.Map.GetType(x, y)

		diff := faction.Color - target

		change := 0
		if diff > 0 {
			change = int(target) + dig
		} else {
			change = int(target) - dig
			if change <= int(color.River) {
				change += 7
			}
		}

		spade = dig
	}

	faction.Dig(spade)
	p.Map.SetType(x, y, color.Color(change))

	return nil
}

func (p *Game) GetRoundTile(user int, tile int) error {
	if p.Round != RoundTileRound {
		log.Println("round error : GetRoundTile")
		return errors.New("round error : GetRoundTile")
	}

	if !p.IsTurn(user) {
		log.Println("It's not a turn", p.Turn, user)
		return errors.New("It's not a turn")
	}

	faction := p.Factions[user].GetInstance()
	p.RoundTiles.Items[tile].Use = true
	faction.RoundTile = &p.RoundTiles.Items[tile]

	p.TurnEnd(user)

	return nil
}

func (p *Game) Bridge(user int, x1 int, y1 int, x2 int, y2 int) error {
	if p.Round < 1 {
		log.Println("round error")
		return errors.New("round error")
	}

	if !p.IsTurn(user) || p.IsPowerTurn() {
		log.Println("It's not a turn", p.Turn, user)
		return errors.New("It's not a turn")
	}

	/*
			0, 0, 2, 0
			1, 0, 0, 2
			1, 0, 2, 2
		    0, 0, 1, 1
			2, 0, 1, 1
	*/

	flag := false

	if y1 > y2 {
		x1, x2 = x2, x1
		y1, y2 = y2, y1
	}

	if y1 == y2 {
		if x1 > x2 {
			x1, x2 = x2, x1
		}
		if x2-x1 == 2 {
			y := y1
			if x1%2 == 0 {
				y--
			}

			if p.Map.GetType(x1+1, y) == color.River && p.Map.GetType(x1+1, y+1) == color.River {
				flag = true
			}
		}
	} else {
		if x1%2 == 1 {
			if (x1-x2 == 1 || x1-x2 == -1) && y2-y1 == 2 {
				if p.Map.GetType(x1, y1+1) == color.River && p.Map.GetType(x2, y2-1) == color.River {
					flag = true
				}
			}
		} else {
			log.Println("1")
			if (x1-x2 == 1 || x1-x2 == -1) && y2-y1 == 1 {
				log.Println("2", x1, y1, x2, y2)
				if p.Map.GetType(x1, y1+1) == color.River && p.Map.GetType(x2, y2-1) == color.River {
					flag = true
				}
			}
		}

	}

	if flag == false {
		log.Println("Locations that cannot be built")
		return errors.New("Locations that cannot be built")
	}

	faction := p.Factions[user].GetInstance()

	err := p.Map.CheckBridge(faction.Color, x1, y1, x2, y2)
	if err != nil {
		log.Println(err)
		return err
	}

	err = faction.Bridge(x1, y1, x2, y2)
	if err != nil {
		return err
	}

	p.Map.Bridge(faction.Color, x1, y1, x2, y2)

	return nil
}

func (p *Game) PowerDiffusion(user int, x int, y int) {
	powers := make([]int, len(p.Factions))

	positions := resources.GetGroundPosition(x, y)

	for i, v := range p.Factions {
		if i == user {
			continue
		}

		faction := v.GetInstance()

		for _, position := range positions {
			owner := p.Map.GetOwner(x+position.X, y+position.Y)

			if owner == faction.Color {
				powers[i] += p.Map.GetPower(x+position.X, y+position.Y)
			}
		}
	}

	users := make([]int, 0)
	for i := user + 1; i < len(p.Factions); i++ {
		users = append(users, i)
	}

	for i := 0; i < user; i++ {
		users = append(users, i)
	}

	for _, v := range users {
		if powers[v] > 0 {
			p.PowerTurn = append(p.PowerTurn, Turn{User: v, Type: PowerTurn, Power: powers[v], From: user})
		}
	}
}

func (p *Game) PowerConfirm(user int, confirm bool) error {
	if p.Round < 1 {
		log.Println("round error")
		return errors.New("round error")
	}

	if !p.IsTurn(user) || !p.IsPowerTurn() {
		log.Println("It's not a turn", p.Turn, user)
		return errors.New("It's not a turn")
	}

	if confirm == true {
		faction := p.Factions[user].GetInstance()
		faction.ReceivePower(p.Turn[0].Power, true)
	}

	p.TurnEnd(user)

	return nil
}

func (p *Game) City(user int, city CityType) error {
	if p.Round < 1 {
		log.Println("round error")
		return errors.New("round error")
	}

	if !p.IsTurn(user) {
		log.Println("It's not a turn", p.Turn, user)
		return errors.New("It's not a turn")
	}

	faction := p.Factions[user].GetInstance()
	if faction.Resource.City == 0 {
		log.Println("have not city")
		return errors.New("have not city")
	}

	if !p.Cities.IsRemain(city) {
		log.Println("not remain city")
		return errors.New("not remain city")
	}

	tile := p.Cities.Use(city)
	faction.ReceiveCity(tile)
	p.Sciences.Receive(faction, tile.Receive)

	return nil
}

func (p *Game) Science(user int, resource resources.Price) error {
	if p.Round < 1 {
		log.Println("round error")
		return errors.New("round error")
	}

	if !p.IsTurn(user) || !p.IsScienceTurn() {
		log.Println("It's not a turn", p.Turn, user)
		return errors.New("It's not a turn")
	}

	faction := p.Factions[user].GetInstance()
	if faction.Resource.Science.Any == 0 && faction.Resource.Science.Single == 0 {
		log.Println("have not science")
		return errors.New("have not science")
	}

	return nil
}

func (p *Game) RoundBonus(faction *factions.Faction) {
}

func (p *Game) ConvertDig(user int, spade int) error {
	if p.Round < 1 {
		log.Println("round error")
		return errors.New("round error")
	}

	if !p.IsTurn(user) {
		log.Println("It's not a turn", p.Turn, user)
		return errors.New("It's not a turn")
	}

	faction := p.Factions[user].GetInstance()
	faction.ConvertDig(spade)
	return nil
}

func (p *Game) PalaceTile(user int, pos int) error {
	if p.Round < 1 {
		log.Println("round error")
		return errors.New("round error")
	}

	if !p.IsTurn(user) {
		log.Println("It's not a turn", p.Turn, user)
		return errors.New("It's not a turn")
	}

	log.Println("palace", len(p.PalaceTiles.Items), pos)
	if pos >= len(p.PalaceTiles.Items) {
		log.Println("not found tile")
		return errors.New("not found tile")
	}

	tile := p.PalaceTiles.Items[pos]

	if tile.Use == true {
		return errors.New("already select")
	}

	faction := p.Factions[user].GetInstance()
	err := faction.PalaceTile(tile)

	if err == nil {
		p.PalaceTiles.Setup(pos)
	}
	return nil
}

func (p *Game) TileAction(user int, category resources.TileCategory, pos int) error {
	if p.Round < 1 {
		log.Println("round error")
		return errors.New("round error")
	}

	if !p.IsTurn(user) || p.IsPowerTurn() {
		log.Println("It's not a turn", p.Turn, user)
		return errors.New("It's not a turn")
	}

	faction := p.Factions[user].GetInstance()
	if faction.Action {
		return errors.New("Already completed the action")
	}

	err := faction.TileAction(category, pos)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (p *Game) SchoolTile(user int, science int, level int) error {
	if p.Round < 1 {
		log.Println("round error")
		return errors.New("round error")
	}

	if !p.IsTurn(user) {
		log.Println("It's not a turn", p.Turn, user)
		return errors.New("It's not a turn")
	}

	if p.SchoolTiles.Items[science][level].Count == 0 {
		log.Println("tile not remain")
		return errors.New("tile not remain")
	}

	tile := p.SchoolTiles.Items[science][level].Tile

	faction := p.Factions[user].GetInstance()
	err := faction.SchoolTile(tile)

	if err != nil {
		log.Println(err)
		return err
	}

	faction.ReceiveResource(resources.Price{Book: level})
	p.Sciences.Action(faction, ScienceType(science), 3-level)

	p.SchoolTiles.Items[science][level].Count--

	return nil
}
