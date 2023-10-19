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
	PassOrder    []int                 `json:"PassOrder"`
	TurnOrder    []int                 `json:"turnOrder"`
}

type TurnType int

const (
	NormalTurn TurnType = iota
	PowerTurn
	ScienceTurn
	SpadeTurn
	BookTurn
)

type Turn struct {
	User    int
	Type    TurnType
	From    int
	Power   int
	Science resources.Science
}

func (p *Turn) Print() {
	titles := []string{"Normal", "Power", "Science", "Spade", "Book"}
	log.Printf("user = %v, type = %v\n", p.User, titles[int(p.Type)])
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

	item.PassOrder = make([]int, 0)
	item.PassOrder = make([]int, 0)

	item.Round = -1

	return &item
}

func (p *Game) InitGame() {
	count := len(p.Factions)
	p.PalaceTiles.Init(count)
	p.SchoolTiles.Init(count)

	for i := 0; i < count; i++ {
		p.PassOrder = append(p.PassOrder, i)
	}
}

func (p *Game) AddFaction(item factions.FactionInterface, tile resources.TileItem) {
	item.Init(tile)

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

func (p *Game) BuildStart() {
	p.RoundTiles.Init(len(p.Factions))

	for i, v := range p.Factions {
		faction := v.GetInstance()
		if faction.FirstBuilding == resources.SA {
			continue
		}
		p.Turn = append(p.Turn, Turn{User: i, Type: NormalTurn})
	}

	for i, v := range p.Factions {
		faction := v.GetInstance()
		if faction.FirstBuilding == resources.SA {
			continue
		}
		p.Turn = append(p.Turn, Turn{User: len(p.Factions) - i - 1, Type: NormalTurn})
	}

	for i, v := range p.Factions {
		faction := v.GetInstance()
		if faction.FirstBuilding == resources.SA {
			p.Turn = append(p.Turn, Turn{User: i, Type: NormalTurn})
		}
	}
}

func (p *Game) Start() {
	log.Println("Start")
	p.RoundTiles.Start()

	p.Round++

	p.TurnOrder = p.PassOrder
	p.PassOrder = make([]int, 0)

	for i, _ := range p.Factions {
		user := p.TurnOrder[i]
		faction := p.Factions[user]
		f := faction.GetInstance()
		faction.Income()

		p.Sciences.RoundBonus(faction.GetInstance())

		roundBonus := p.RoundBonuss.Get(p.Round)

		count := 0
		if roundBonus.Science.Banking > 0 {
			count = faction.GetScience(0) / roundBonus.Science.Banking
		} else if roundBonus.Science.Law > 0 {
			count = faction.GetScience(1) / roundBonus.Science.Law
		} else if roundBonus.Science.Engineering > 0 {
			count = faction.GetScience(2) / roundBonus.Science.Engineering
		} else if roundBonus.Science.Medicine > 0 {
			count = faction.GetScience(3) / roundBonus.Science.Medicine
		}

		roundBonus.Receive.Prist *= count
		roundBonus.Receive.Power *= count
		roundBonus.Receive.Book.Any *= count
		roundBonus.Receive.Spade *= count
		roundBonus.Receive.Coin *= count
		roundBonus.Receive.Worker *= count

		f.ReceiveResource(roundBonus.Receive)
	}

	for i, _ := range p.Factions {
		user := p.TurnOrder[i]
		faction := p.Factions[user].GetInstance()

		if faction.Resource.Book.Any > 0 {
			turn := []Turn{Turn{User: user, Type: BookTurn}}
			p.Turn = append(turn, p.Turn...)
		}
	}

	for i, _ := range p.Factions {
		user := p.TurnOrder[i]
		faction := p.Factions[user].GetInstance()

		if !faction.Resource.Science.IsEmpty() {
			turn := []Turn{Turn{User: user, Type: ScienceTurn}}
			p.Turn = append(turn, p.Turn...)
		}
	}

	for i, _ := range p.Factions {
		user := p.TurnOrder[i]
		faction := p.Factions[user].GetInstance()

		if faction.Resource.Spade > 0 {
			log.Println("Add Spade turn")
			turn := []Turn{Turn{User: user, Type: SpadeTurn}}
			p.Turn = append(turn, p.Turn...)
		}
	}

	for i, _ := range p.Factions {
		user := p.TurnOrder[i]

		p.Turn = append(p.Turn, Turn{User: user, Type: NormalTurn})
	}

	for _, v := range p.Turn {
		v.Print()
	}

	p.PowerActions.Start()
	p.BookActions.Start()

	log.Println("round", p.Round)

}

func (p *Game) TurnEnd(user int) {
	log.Println("TurnEnd")
	faction := p.Factions[user]
	faction.TurnEnd()

	turnType := p.Turn[0].Type

	p.Turn = p.Turn[1:]

	p.Turn = append(p.PowerTurn, p.Turn...)

	for _, v := range p.Turn {
		v.Print()
	}

	if p.Round > 0 {
		if user >= 0 {
			if turnType == NormalTurn {
				if !faction.GetInstance().IsPass {
					p.Turn = append(p.Turn, Turn{User: user, Type: NormalTurn})
				}
			}
		}
	}

	if len(p.Turn) == 0 {
		log.Println("turn len == 0")
		if p.Round == BuildRound {
			for i, _ := range p.Factions {
				p.Turn = append(p.Turn, Turn{User: len(p.Factions) - i - 1, Type: NormalTurn})
			}

			p.Round = RoundTileRound
		} else if p.Round == RoundTileRound {
			for _, v := range p.Factions {
				v.FirstIncome()
			}

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

func (p *Game) IsNormalTurn() bool {
	if len(p.Turn) == 0 {
		return false
	}

	if p.Turn[0].Type != NormalTurn {
		return false
	}

	return true
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

func (p *Game) IsSpadeTurn() bool {
	if len(p.Turn) == 0 {
		return false
	}

	if p.Turn[0].Type != SpadeTurn {
		return false
	}

	return true
}

func (p *Game) IsBookTurn() bool {
	if len(p.Turn) == 0 {
		return false
	}

	if p.Turn[0].Type != BookTurn {
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

	faction := p.Factions[user]

	color := p.Map.GetType(x, y)
	f := faction.GetInstance()
	if color != f.Color {
		log.Printf("color error : map = %v, faction = %v\n", color.ToString(), f.Color.ToString())
		return errors.New("It's not a user's land")
	}

	log.Println("build success")
	faction.FirstBuild(x, y)

	p.Map.Build(x, y, f.Color, f.FirstBuilding)

	p.TurnEnd(user)

	return nil
}

func (p *Game) Build(user int, x int, y int, building resources.Building) error {
	if p.Round < 1 {
		log.Println("round error")
		return errors.New("round error")
	}

	if !p.IsTurn(user) {
		log.Println("It's not a turn", p.Turn, user)
		return errors.New("It's not a turn")
	}

	if !p.IsNormalTurn() {
		log.Println("It's not a normal turn")
		return errors.New("It's not a normal turn")
	}

	faction := p.Factions[user]
	f := faction.GetInstance()
	if f.Action {
		if f.Resource.Building == resources.None && f.ExtraBuild == 0 {
			return errors.New("Already completed the action")
		}
	}

	flag := p.Map.CheckDistance(f.Color, f.GetShipDistance(), x, y)

	if f.Resource.Building == resources.TP {
		if p.Map.GetType(x, y) == f.Color {
			flag = true
		}
	}

	log.Println("flag", flag)
	if flag == false {
		log.Println("ship distance error")
		return errors.New("ship distance error")
	}

	err := p.Map.CheckBuild(x, y, f.Color, f.GetSpadeCount())

	if err != nil {
		log.Println("map check build error")
		log.Println(err)
		return err
	}

	needSpade := p.Map.GetNeedSpade(x, y, f.Color)
	if needSpade < 0 {
		needSpade = 0
	}

	err = faction.Build(x, y, needSpade, building)
	if err != nil {
		log.Println(err)
		return err
	}

	p.Map.Build(x, y, f.Color, building)

	lists := p.Map.CheckCity(f.Color, x, y, f.TownPower)

	if len(lists) > 0 {
		f.CityBuildingList = lists
		f.Resource.City++
	}

	buildVP := p.RoundBonuss.GetBuildVP(p.Round)

	if p.Map.IsRiverside(x, y) && building == resources.D {
		f.ReceiveRiverVP()
	}

	if p.Map.IsEdge(x, y) && building == resources.D {
		f.ReceiveEdgeVP()

		if p.Round == 6 {
			f.ReceiveResource(resources.Price{VP: buildVP.Edge})
		}
	}

	if building == resources.D {
		f.ReceiveResource(resources.Price{VP: buildVP.D})
	} else if building == resources.TP {
		f.ReceiveResource(resources.Price{VP: buildVP.TP})
	}

	if p.Round == 6 {
		buildVP = p.RoundBonuss.FinalRound.Build

		if building == resources.D {
			f.ReceiveResource(resources.Price{VP: buildVP.D})
		} else if building == resources.TP {
			f.ReceiveResource(resources.Price{VP: buildVP.TP})
		}
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

	if !p.IsTurn(user) {
		log.Println("It's not a turn", p.Turn, user)
		return errors.New("It's not a turn")
	}

	if !p.IsNormalTurn() {
		log.Println("It's not a normal turn")
		return errors.New("It's not a normal turn")
	}

	faction := p.Factions[user]
	f := faction.GetInstance()
	if f.Action {
		if f.Resource.TpUpgrade == 0 || target != resources.TP {
			return errors.New("Already completed the action")
		}
	}

	if p.Map.GetOwner(x, y) != f.Color {
		log.Println("not owner")
		return errors.New("not owner")
	}

	err := faction.Upgrade(x, y, target)

	if err != nil {
		return err
	}

	p.Map.SetBuilding(x, y, target)

	lists := p.Map.CheckCity(f.Color, x, y, f.TownPower)

	if len(lists) > 0 {
		f.CityBuildingList = lists
		f.Resource.City++
	}

	buildVP := p.RoundBonuss.GetBuildVP(p.Round)

	if target == resources.TP {
		f.ReceiveResource(resources.Price{VP: buildVP.TP})
	} else if target == resources.TE {
		f.ReceiveResource(resources.Price{VP: buildVP.TE})
	} else if target == resources.SA || target == resources.SH {
		f.ReceiveResource(resources.Price{VP: buildVP.SHSA})
	}

	if p.Round == 6 {
		buildVP = p.RoundBonuss.FinalRound.Build

		if target == resources.TP {
			f.ReceiveResource(resources.Price{VP: buildVP.TP})
		} else if target == resources.TE {
			f.ReceiveResource(resources.Price{VP: buildVP.TE})
		} else if target == resources.SA || target == resources.SH {
			f.ReceiveResource(resources.Price{VP: buildVP.SHSA})
		}
	}

	p.PowerDiffusion(user, x, y)

	return nil
}

func (p *Game) PowerAction(user int, pos int) error {
	if p.Round < 1 {
		log.Println("round error")
		return errors.New("round error")
	}

	if !p.IsTurn(user) {
		log.Println("It's not a turn", p.Turn, user)
		return errors.New("It's not a turn")
	}

	if !p.IsNormalTurn() {
		log.Println("It's not a normal turn")
		return errors.New("It's not a normal turn")
	}

	faction := p.Factions[user]
	f := faction.GetInstance()
	if f.Action {
		return errors.New("Already completed the action")
	}

	if p.PowerActions.IsUse(pos) {
		log.Println("already")
		return errors.New("already")
	}

	have := f.GetHavePowerCount()
	if have < p.PowerActions.GetNeedPower(pos) {
		log.Println("not enough power")
		return errors.New("not enough power")
	}

	item := p.PowerActions.Action(pos)
	err := faction.PowerAction(item)
	if err != nil {
		return err
	}

	for _, v := range f.Tiles {
		if v.Type == resources.TileFactionIllusionists {
			vp := 3

			if len(p.Factions) == 5 {
				vp = 4
			}

			f.ReceiveResource(resources.Price{VP: vp})
		}
	}

	return nil
}

func (p *Game) BookAction(user int, pos int, book resources.Book) error {
	if p.Round < 1 {
		log.Println("round error")
		return errors.New("round error")
	}

	if !p.IsTurn(user) {
		log.Println("It's not a turn", p.Turn, user)
		return errors.New("It's not a turn")
	}

	if !p.IsNormalTurn() {
		log.Println("It's not a normal turn")
		return errors.New("It's not a normal turn")
	}

	faction := p.Factions[user]
	f := faction.GetInstance()
	if f.Action {
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

	have := f.Resource.Book.Count()
	if have < p.BookActions.GetNeedBook(pos) {
		log.Println("not enough book")
		return errors.New("not enough book")
	}

	item := p.BookActions.Action(pos)
	faction.Book(item, book)

	return nil
}

func (p *Game) AdvanceShip(user int) error {
	if p.Round < 1 {
		log.Println("round error")
		return errors.New("round error")
	}

	if !p.IsTurn(user) {
		log.Println("It's not a turn", p.Turn, user)
		return errors.New("It's not a turn")
	}

	if !p.IsNormalTurn() {
		log.Println("It's not a normal turn")
		return errors.New("It's not a normal turn")
	}

	faction := p.Factions[user]
	f := faction.GetInstance()
	if f.Action {
		return errors.New("Already completed the action")
	}

	err := faction.AdvanceShip()
	if err != nil {
		log.Println(err)
		return err
	}

	buildVP := p.RoundBonuss.GetBuildVP(p.Round)

	f.ReceiveResource(resources.Price{VP: buildVP.Advance})

	return nil
}

func (p *Game) AdvanceSpade(user int) error {
	if p.Round < 1 {
		log.Println("round error")
		return errors.New("round error")
	}

	if !p.IsTurn(user) {
		log.Println("It's not a turn", p.Turn, user)
		return errors.New("It's not a turn")
	}

	if !p.IsNormalTurn() {
		log.Println("It's not a normal turn")
		return errors.New("It's not a normal turn")
	}

	faction := p.Factions[user]
	f := faction.GetInstance()
	if f.Action {
		return errors.New("Already completed the action")
	}

	err := faction.AdvanceSpade()
	if err != nil {
		log.Println(err)
		return err
	}

	buildVP := p.RoundBonuss.GetBuildVP(p.Round)

	f.ReceiveResource(resources.Price{VP: buildVP.Advance})

	return nil
}

func (p *Game) SendScholar(user int, pos ScienceType) error {
	if p.Round < 1 {
		log.Println("round error")
		return errors.New("round error")
	}

	if !p.IsTurn(user) {
		log.Println("It's not a turn", p.Turn, user)
		return errors.New("It's not a turn")
	}

	if !p.IsNormalTurn() {
		log.Println("It's not a normal turn")
		return errors.New("It's not a normal turn")
	}

	faction := p.Factions[user]
	f := faction.GetInstance()
	if f.Action {
		return errors.New("Already completed the action")
	}

	if f.Resource.Prist == 0 {
		return errors.New("not enough prist")
	}

	inc := p.Sciences.Send(f, pos)

	faction.SendScholar()

	buildVP := p.RoundBonuss.GetBuildVP(p.Round)
	f.ReceiveResource(resources.Price{VP: buildVP.Science * inc})

	return nil
}

func (p *Game) SupployScholar(user int, pos ScienceType) error {
	if p.Round < 1 {
		log.Println("round error")
		return errors.New("round error")
	}

	if !p.IsTurn(user) {
		log.Println("It's not a turn", p.Turn, user)
		return errors.New("It's not a turn")
	}

	if !p.IsNormalTurn() {
		log.Println("It's not a normal turn")
		return errors.New("It's not a normal turn")
	}

	faction := p.Factions[user]
	f := faction.GetInstance()
	if f.Action {
		return errors.New("Already completed the action")
	}

	if f.Resource.Prist == 0 {
		return errors.New("not enough prist")
	}

	inc := p.Sciences.Supploy(f, pos)

	faction.SupployScholar()

	buildVP := p.RoundBonuss.GetBuildVP(p.Round)
	f.ReceiveResource(resources.Price{VP: buildVP.Science * inc})

	return nil
}

func (p *Game) Pass(user int, pos int) error {
	if p.Round < 1 {
		log.Println("round error")
		return errors.New("round error")
	}

	if !p.IsTurn(user) {
		log.Println("It's not a turn", p.Turn, user)
		return errors.New("It's not a turn")
	}

	if !p.IsNormalTurn() {
		log.Println("It's not a normal turn")
		return errors.New("It's not a normal turn")
	}

	faction := p.Factions[user]
	roundTile := p.RoundTiles.Pass(pos)
	err, tile := faction.Pass(roundTile)

	if err != nil {
		p.RoundTiles.Add(roundTile)

		log.Println(err)
		return err
	}

	p.RoundTiles.Add(tile)
	p.PassOrder = append(p.PassOrder, user)

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

	if !p.IsTurn(user) {
		log.Println("It's not a turn", p.Turn, user)
		return errors.New("It's not a turn")
	}

	if !p.IsNormalTurn() && !p.IsSpadeTurn() {
		log.Println("It's not a normal turn")
		return errors.New("It's not a normal turn")
	}

	faction := p.Factions[user]
	f := faction.GetInstance()

	if p.Map.GetType(x, y) == color.River {
		log.Println("can't spade")
		return errors.New("can't spade")
	}

	if f.Resource.Spade < dig {
		log.Println("not have spade")
		return errors.New("not have spade")
	}

	needSpade := p.Map.GetNeedSpade(x, y, f.Color)
	if needSpade == 0 {
		log.Println("need not spade")
		return errors.New("need not spade")
	}

	spade := 0
	var change color.Color
	if dig >= needSpade {
		change = f.Color
		spade = needSpade
	} else {
		target := p.Map.GetType(x, y)

		diff := f.Color - target

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

	buildVP := p.RoundBonuss.GetBuildVP(p.Round)

	f.ReceiveResource(resources.Price{VP: buildVP.Spade * spade})

	return nil
}

func (p *Game) GetRoundTile(user int, pos int) error {
	if p.Round != RoundTileRound {
		log.Println("round error : GetRoundTile")
		return errors.New("round error : GetRoundTile")
	}

	if !p.IsTurn(user) {
		log.Println("It's not a turn", p.Turn, user)
		return errors.New("It's not a turn")
	}

	if !p.IsNormalTurn() {
		log.Println("It's not a normal turn")
		return errors.New("It's not a normal turn")
	}

	faction := p.Factions[user]
	tile := p.RoundTiles.Pass(pos)
	log.Println("tile name", tile.Name)
	faction.RoundTile(tile)

	p.TurnEnd(user)

	return nil
}

func (p *Game) Bridge(user int, x1 int, y1 int, x2 int, y2 int) error {
	if p.Round < 1 {
		log.Println("round error")
		return errors.New("round error")
	}

	if !p.IsTurn(user) {
		log.Println("It's not a turn", p.Turn, user)
		return errors.New("It's not a turn")
	}

	if !p.IsNormalTurn() {
		log.Println("It's not a normal turn")
		return errors.New("It's not a normal turn")
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

	faction := p.Factions[user]
	f := faction.GetInstance()

	err := p.Map.CheckBridge(f.Color, x1, y1, x2, y2)
	if err != nil {
		log.Println(err)
		return err
	}

	err = faction.Bridge(x1, y1, x2, y2)
	if err != nil {
		return err
	}

	p.Map.Bridge(f.Color, x1, y1, x2, y2)

	return nil
}

func (p *Game) PowerDiffusion(user int, x int, y int) {
	powers := make([]int, len(p.Factions))

	positions := resources.GetGroundPosition(x, y)

	for i, v := range p.Factions {
		if i == user {
			continue
		}

		faction := v
		f := faction.GetInstance()

		for _, position := range positions {
			owner := p.Map.GetOwner(x+position.X, y+position.Y)

			if owner == f.Color {
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
		faction := p.Factions[user]
		f := faction.GetInstance()
		f.ReceivePower(p.Turn[0].Power, true)
	}

	p.TurnEnd(user)

	return nil
}

func (p *Game) City(user int, city resources.CityType) error {
	if p.Round < 1 {
		log.Println("round error")
		return errors.New("round error")
	}

	if !p.IsTurn(user) {
		log.Println("It's not a turn", p.Turn, user)
		return errors.New("It's not a turn")
	}

	if !p.IsNormalTurn() {
		log.Println("It's not a normal turn")
		return errors.New("It's not a normal turn")
	}

	faction := p.Factions[user]
	f := faction.GetInstance()

	if f.Resource.City == 0 {
		log.Println("have not city")
		return errors.New("have not city")
	}

	if !p.Cities.IsRemain(city) {
		log.Println("not remain city")
		return errors.New("not remain city")
	}

	tile := p.Cities.Use(city)
	faction.ReceiveCity(tile)
	p.Sciences.Receive(f, tile.Receive)

	p.Map.AddCityBuildingList(f.CityBuildingList)
	f.CityBuildingList = make([]resources.Position, 0)

	buildVP := p.RoundBonuss.GetBuildVP(p.Round)

	f.ReceiveResource(resources.Price{VP: buildVP.City})

	return nil
}

func (p *Game) Science(user int, pos ScienceType, level int) error {
	if p.Round < 1 {
		log.Println("round error")
		return errors.New("round error")
	}

	if !p.IsTurn(user) {
		log.Println("It's not a turn", p.Turn, user)
		return errors.New("It's not a turn")
	}

	if !p.IsNormalTurn() && !p.IsScienceTurn() {
		log.Println("It's not a normal turn")
		return errors.New("It's not a normal turn")
	}

	faction := p.Factions[user]
	f := faction.GetInstance()

	if f.Resource.Science.Any == 0 && f.Resource.Science.Single == 0 {
		log.Println("have not science")
		return errors.New("have not science")
	}

	if level > 1 {
		if f.Resource.Science.Single < level {
			log.Println("not enough science")
			return errors.New("not enough science")
		}

		f.Resource.Science.Single -= level
	} else {
		if f.Resource.Science.Single > 0 {
			f.Resource.Science.Single -= level
		} else {
			f.Resource.Science.Any -= level
		}
	}

	inc := p.Sciences.Action(f, pos, level)

	buildVP := p.RoundBonuss.GetBuildVP(p.Round)
	f.ReceiveResource(resources.Price{VP: buildVP.Science * inc})

	return nil
}

func (p *Game) Book(user int, pos resources.BookType, count int) error {
	if p.Round < 1 {
		log.Println("round error")
		return errors.New("round error")
	}

	if !p.IsTurn(user) {
		log.Println("It's not a turn", p.Turn, user)
		return errors.New("It's not a turn")
	}

	if !p.IsNormalTurn() && !p.IsBookTurn() {
		log.Println("It's not a normal turn")
		return errors.New("It's not a normal turn")
	}

	faction := p.Factions[user]
	f := faction.GetInstance()

	if f.Resource.Book.Any == 0 {
		log.Println("have not book")
		return errors.New("have not book")
	}

	var book resources.Book
	if pos == resources.BookBanking {
		book.Banking = count
	} else if pos == resources.BookLaw {
		book.Law = count
	} else if pos == resources.BookEngineering {
		book.Engineering = count
	} else if pos == resources.BookMedicine {
		book.Medicine = count
	}

	f.ReceiveResource(resources.Price{Book: book})
	f.Resource.Book.Any -= count

	return nil
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

	if !p.IsNormalTurn() {
		log.Println("It's not a normal turn")
		return errors.New("It's not a normal turn")
	}

	faction := p.Factions[user]
	f := faction.GetInstance()
	f.ConvertDig(spade)
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

	if !p.IsNormalTurn() {
		log.Println("It's not a normal turn")
		return errors.New("It's not a normal turn")
	}

	log.Println("palace", len(p.PalaceTiles.Items), pos)
	if pos >= len(p.PalaceTiles.Items) {
		log.Println("not found tile")
		return errors.New("not found tile")
	}

	tile := p.PalaceTiles.GetTile(pos)

	if tile.Use == true {
		return errors.New("already select")
	}

	faction := p.Factions[user]
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

	if !p.IsTurn(user) {
		log.Println("It's not a turn", p.Turn, user)
		return errors.New("It's not a turn")
	}

	if !p.IsNormalTurn() {
		log.Println("It's not a normal turn")
		return errors.New("It's not a normal turn")
	}

	faction := p.Factions[user]
	f := faction.GetInstance()

	if f.Action {
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

	if !p.IsNormalTurn() {
		log.Println("It's not a normal turn")
		return errors.New("It's not a normal turn")
	}

	if p.SchoolTiles.Items[science][level].Count == 0 {
		log.Println("tile not remain")
		return errors.New("tile not remain")
	}

	tile := p.SchoolTiles.Items[science][level].Tile

	faction := p.Factions[user]
	f := faction.GetInstance()
	err := faction.SchoolTile(tile, science)

	if err != nil {
		log.Println(err)
		return err
	}

	var book resources.Book
	if science == 0 {
		book.Banking = level
	} else if science == 1 {
		book.Law = level
	} else if science == 2 {
		book.Engineering = level
	} else if science == 3 {
		book.Medicine = level
	}

	f.ReceiveResource(resources.Price{Book: book})
	inc := p.Sciences.Action(f, ScienceType(science), 3-level)

	buildVP := p.RoundBonuss.GetBuildVP(p.Round)
	f.ReceiveResource(resources.Price{VP: buildVP.Science * inc})

	p.SchoolTiles.Items[science][level].Count--

	return nil
}

func (p *Game) Burn(user int, count int) error {
	if p.Round < 1 {
		log.Println("round error")
		return errors.New("round error")
	}

	if !p.IsTurn(user) {
		log.Println("It's not a turn", p.Turn, user)
		return errors.New("It's not a turn")
	}

	if !p.IsNormalTurn() {
		log.Println("It's not a normal turn")
		return errors.New("It's not a normal turn")
	}

	faction := p.Factions[user]
	f := faction.GetInstance()
	return f.Burn(count)
}

func (p *Game) Convert(user int, source resources.Price, target resources.Price) error {
	if p.Round < 1 {
		log.Println("round error")
		return errors.New("round error")
	}

	if !p.IsTurn(user) {
		log.Println("It's not a turn", p.Turn, user)
		return errors.New("It's not a turn")
	}

	if !p.IsNormalTurn() {
		log.Println("It's not a normal turn")
		return errors.New("It's not a normal turn")
	}

	faction := p.Factions[user]
	f := faction.GetInstance()
	return f.Convert(source, target)
}

func (p *Game) Annex(user int, x int, y int) error {
	if p.Round < 1 {
		log.Println("round error")
		return errors.New("round error")
	}

	if !p.IsTurn(user) {
		log.Println("It's not a turn", p.Turn, user)
		return errors.New("It's not a turn")
	}

	if !p.IsNormalTurn() {
		log.Println("It's not a normal turn")
		return errors.New("It's not a normal turn")
	}

	faction := p.Factions[user]
	f := faction.GetInstance()

	err := p.Map.CheckAnnex(f.Color, x, y)
	if err != nil {
		log.Println(err)
		return err
	}

	err = f.Annex(x, y)
	if err != nil {
		return err
	}

	p.Map.Annex(f.Color, x, y)

	lists := p.Map.CheckCity(f.Color, x, y, f.TownPower)

	if len(lists) > 0 {
		f.CityBuildingList = lists
		f.Resource.City++
	}

	return nil
}
