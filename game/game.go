package game

import (
	"aoi/game/action"
	"aoi/game/color"
	"aoi/game/factions"
	"aoi/game/resources"
	"aoi/models"
	"aoi/models/game"
	"errors"
	"log"
	"sort"
	"strings"
)

const (
	InitRound      = -3
	FactionRound   = -2
	BuildRound     = -1
	RoundTileRound = 0
)

type Game struct {
	Id       int64                       `json:"id"`
	Map      *Map                        `json:"map"`
	Sciences *Science                    `json:"sciences"`
	Factions []factions.FactionInterface `json:"factions"`

	PowerActions    *action.PowerAction       `json:"powerActions"`
	BookActions     *action.BookAction        `json:"bookActions"`
	RoundTiles      *resources.RoundTile      `json:"roundTiles"`
	RoundBonuss     *RoundBonus               `json:"roundBonuss"`
	PalaceTiles     *resources.PalaceTile     `json:"palaceTiles"`
	SchoolTiles     *resources.SchoolTile     `json:"schoolTiles"`
	InnovationTiles *resources.InnovationTile `json:"innovationTiles"`
	FactionTiles    *resources.FactionTile    `json:"factionTiles"`
	ColorTiles      *resources.ColorTile      `json:"colorTiles"`
	Cities          *City                     `json:"cities"`
	Turn            []Turn                    `json:"turn"`
	PowerTurn       []Turn                    `json:"powerTurn"`
	Round           int                       `json:"round"`
	PassOrder       []int                     `json:"passOrder"`
	TurnOrder       []int                     `json:"turnOrder"`
	Users           []int64                   `json:"users"`
	Count           int                       `json:"count"`
	Command         []string                  `json:"commands"`
	OldCommand      []string                  `json:"oldCommands"`
}

type TurnType int

const (
	NormalTurn TurnType = iota
	PowerTurn
	ScienceTurn
	SpadeTurn
	BookTurn
	TileTurn
	BuildTurn
	ResourceTurn
)

type Turn struct {
	User    int               `json:"user"`
	Type    TurnType          `json:"type"`
	From    int               `json:"from"`
	Power   int               `json:"power"`
	Science resources.Science `json:"science"`
}

func (p *Turn) Print() {
	titles := []string{"Normal", "Power", "Science", "Spade", "Book", "Tile", "Build", "Resource"}
	log.Printf("user = %v, type = %v\n", p.User, titles[int(p.Type)])
}

func NewGame(id int64, count int) *Game {
	var item Game
	item.Id = id
	item.PowerActions = action.NewPowerAction()
	item.BookActions = action.NewBookAction(id)
	item.RoundTiles = resources.NewRoundTile(id)
	item.RoundBonuss = NewRoundBonus(id)
	item.PalaceTiles = resources.NewPalaceTile(id)
	item.SchoolTiles = resources.NewSchoolTile(id, count)
	item.InnovationTiles = resources.NewInnovationTile(id, count)
	item.FactionTiles = resources.NewFactionTile(id)
	item.ColorTiles = resources.NewColorTile(id)
	item.Cities = NewCity()

	item.Map = NewMap(count)
	item.Sciences = NewScience()
	item.Factions = make([]factions.FactionInterface, 0)

	item.Turn = make([]Turn, 0)
	item.PowerTurn = make([]Turn, 0)
	item.PassOrder = make([]int, 0)
	item.TurnOrder = make([]int, 0)
	item.Users = make([]int64, 0)

	item.Command = make([]string, 0)
	item.OldCommand = make([]string, 0)

	item.Count = count

	item.Round = InitRound

	return &item
}

func (p *Game) CheckUser(id int64, user int) bool {
	count := len(p.Users)

	if user < 0 || user > count-1 {
		return false
	}

	if p.Users[user] != id {
		return false
	}

	return true
}

func (p *Game) AddUser(user int64) {
	p.Users = append(p.Users, user)
	item := &factions.Basic{}
	colorTile := p.ColorTiles.Items[0]
	item.Init(colorTile)
	p.Factions = append(p.Factions, item)
}

func (p *Game) UpdateDBRound(value int) {
	conn := models.NewConnection()
	defer conn.Close()

	gameManager := models.NewGameManager(conn)
	gameManager.UpdateStatus(value, p.Id)
}

func (p *Game) CompleteAddUser() {
	p.Round = FactionRound

	for i := 0; i < p.Count; i++ {
		p.Turn = append(p.Turn, Turn{User: p.Count - i - 1, Type: NormalTurn})
	}

	p.UpdateDBRound(int(game.StatusFaction))
}

func (p *Game) AddFaction(item factions.FactionInterface, tile resources.TileItem) {
	item.Init(tile)

	p.Factions = append(p.Factions, item)

	faction := item.GetInstance()
	p.Sciences.AddUser(faction.Color, faction.Science)
}

func (p *Game) SelectFaction(user int, name string) {
	pos := 0

	for i, v := range p.FactionTiles.Items {
		if strings.ToLower(v.Name) == name {
			pos = i
		}
	}

	factionTile := p.FactionTiles.Items[pos]
	colorTile := p.ColorTiles.Items[pos]
	roundTile := p.RoundTiles.Reserved[pos]

	var item factions.FactionInterface

	if factionTile.Type == resources.TileFactionBlessed {
		item = &factions.Blessed{}
	} else if factionTile.Type == resources.TileFactionFelines {
		item = &factions.Felines{}
	} else if factionTile.Type == resources.TileFactionGoblins {
		item = &factions.Goblins{}
	} else if factionTile.Type == resources.TileFactionIllusionists {
		item = &factions.Illusionists{}
	} else if factionTile.Type == resources.TileFactionInventors {
		item = &factions.Inventors{}
	} else if factionTile.Type == resources.TileFactionLizards {
		item = &factions.Lizards{}
	} else if factionTile.Type == resources.TileFactionMoles {
		item = &factions.Moles{}
	} else if factionTile.Type == resources.TileFactionMonks {
		item = &factions.Monks{}
	} else if factionTile.Type == resources.TileFactionNavigators {
		item = &factions.Navigators{}
	} else if factionTile.Type == resources.TileFactionOmar {
		item = &factions.Omar{}
	} else if factionTile.Type == resources.TileFactionPhilosophers {
		item = &factions.Philosophers{}
	} else if factionTile.Type == resources.TileFactionPsychics {
		item = &factions.Psychics{}
	}

	item.Init(colorTile)

	p.Factions[user] = item

	f := item.GetInstance()

	f.ReceiveResource(factionTile.Once)
	f.ReceiveResource(colorTile.Once)

	f.IncScience(int(Banking), p.Sciences.Action(f, Banking, factionTile.Once.Science.Banking))
	f.IncScience(int(Law), p.Sciences.Action(f, Law, factionTile.Once.Science.Law))
	f.IncScience(int(Engineering), p.Sciences.Action(f, Engineering, factionTile.Once.Science.Engineering))
	f.IncScience(int(Medicine), p.Sciences.Action(f, Medicine, factionTile.Once.Science.Medicine))

	f.IncScience(int(Banking), p.Sciences.Action(f, Banking, colorTile.Once.Science.Banking))
	f.IncScience(int(Law), p.Sciences.Action(f, Law, colorTile.Once.Science.Law))
	f.IncScience(int(Engineering), p.Sciences.Action(f, Engineering, colorTile.Once.Science.Engineering))
	f.IncScience(int(Medicine), p.Sciences.Action(f, Medicine, colorTile.Once.Science.Medicine))

	item.RoundTile(roundTile)
	p.Sciences.AddUser(f.Color, f.Science)

	p.FactionTiles.Items[pos].Use = true
	f.Action = true
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
	p.RoundTiles.BuildStart()

	for i, v := range p.Factions {
		faction := v.GetInstance()
		if faction.FirstBuilding == resources.SA {
			continue
		}
		p.Turn = append(p.Turn, Turn{User: i, Type: NormalTurn})
	}

	for i := len(p.Factions) - 1; i >= 0; i-- {
		v := p.Factions[i]
		faction := v.GetInstance()
		if faction.FirstBuilding == resources.SA {
			continue
		}
		p.Turn = append(p.Turn, Turn{User: i, Type: NormalTurn})
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

	if p.Round >= 1 {
		for i := range p.Factions {
			user := p.TurnOrder[i]
			faction := p.Factions[user]
			f := faction.GetInstance()
			faction.Income()

			p.Sciences.RoundBonus(faction.GetInstance())

			if p.Round > 1 {
				roundBonus := p.RoundBonuss.Get(p.Round - 1)

				log.Println(roundBonus)

				log.Println("banking", roundBonus.Science.Banking)
				log.Println("law", roundBonus.Science.Law)
				log.Println("engineering", roundBonus.Science.Engineering)
				log.Println("medicine", roundBonus.Science.Medicine)
				count := 0
				if roundBonus.Science.Banking > 0 {
					count = faction.GetScience(0) / roundBonus.Science.Banking
				} else if roundBonus.Science.Law > 0 {
					count = faction.GetScience(1) / roundBonus.Science.Law
				} else if roundBonus.Science.Engineering > 0 {
					count = faction.GetScience(2) / roundBonus.Science.Engineering
				} else if roundBonus.Science.Medicine > 0 {
					log.Println("faction value", faction.GetScience(3))
					count = faction.GetScience(3) / roundBonus.Science.Medicine
				}

				log.Println("round count", count)
				roundBonus.Receive.Prist *= count
				roundBonus.Receive.Power *= count
				roundBonus.Receive.Book.Any *= count
				roundBonus.Receive.Spade *= count
				roundBonus.Receive.Coin *= count
				roundBonus.Receive.Worker *= count

				f.ReceiveResource(roundBonus.Receive)
			}
		}
	}

	for i := range p.Factions {
		user := p.TurnOrder[i]
		faction := p.Factions[user].GetInstance()

		if faction.Resource.Book.Any > 0 ||
			!faction.Resource.Science.IsEmpty() ||
			faction.Resource.Spade > 0 ||
			faction.Resource.Building != resources.None {

			turn := Turn{User: user, Type: ResourceTurn}
			p.Turn = append(p.Turn, turn)
		}
	}

	/*
		for i := range p.Factions {
			user := p.TurnOrder[i]
			faction := p.Factions[user].GetInstance()

			if !faction.Resource.Science.IsEmpty() {
				turn := []Turn{{User: user, Type: ScienceTurn}}
				p.Turn = append(turn, p.Turn...)
			}
		}

		for i := range p.Factions {
			user := p.TurnOrder[i]
			faction := p.Factions[user].GetInstance()

			if faction.Resource.Spade > 0 {
				turn := []Turn{{User: user, Type: SpadeTurn}}
				p.Turn = append(turn, p.Turn...)
			}
		}

		for i := range p.Factions {
			user := p.TurnOrder[i]
			faction := p.Factions[user].GetInstance()

			if faction.Resource.Building != resources.None {
				turn := []Turn{{User: user, Type: BuildTurn}}
				p.Turn = append(turn, p.Turn...)
			}
		}
	*/

	for i := range p.Factions {
		user := p.TurnOrder[i]

		p.Turn = append(p.Turn, Turn{User: user, Type: NormalTurn})
	}

	for _, v := range p.Turn {
		v.Print()
	}

	p.PowerActions.Start()
	p.BookActions.Start()
}

func (p *Game) TurnEnd(user int) error {
	if !p.IsTurn(user) {
		return errors.New("It's not a turn")
	}

	faction := p.Factions[user]
	f := faction.GetInstance()

	if f.Action == false {
		return errors.New("must action")
	}

	turnType := p.Turn[0].Type

	p.Turn = p.Turn[1:]

	p.Turn = append(p.PowerTurn, p.Turn...)
	p.PowerTurn = make([]Turn, 0)

	faction.TurnEnd(p.Round)

	if p.Round > 0 {
		if user >= 0 {
			if turnType == NormalTurn {
				if !f.IsPass {
					p.Turn = append(p.Turn, Turn{User: user, Type: NormalTurn})
				}
			}
		}
	} else {
		//p.PassOrder = append(p.PassOrder, user)
	}

	if len(p.Turn) == 0 {
		if p.Round == FactionRound {
			p.Round = BuildRound
			p.PassOrder = make([]int, 0)

			p.UpdateDBRound(int(game.StatusNormal))

			p.BuildStart()
		} else if p.Round == BuildRound {
			p.Round = 0
			p.PassOrder = make([]int, 0)
			for i, _ := range p.Factions {
				p.PassOrder = append(p.PassOrder, i)
			}

			log.Println("start 1 round")
			p.Start()
		} else {
			p.RoundEnd()

			if p.Round == 6 {
				p.Round++
				p.EndGame()
			} else {
				p.Start()
			}
		}
	}

	log.Println("round", p.Round)
	for _, v := range p.Turn {
		v.Print()
	}

	return nil
}

func (p *Game) RoundEnd() {
	for _, v := range p.Factions {
		faction := v.GetInstance()
		p.Sciences.RoundEndBonus(faction, p.RoundBonuss.Items[p.Round-1])
	}
}

func (p *Game) IsResourceTurn() bool {
	if len(p.Turn) == 0 {
		return false
	}

	if p.Turn[0].Type != ResourceTurn {
		return false
	}

	return true
}

func (p *Game) IsBuildTurn() bool {
	if len(p.Turn) == 0 {
		return false
	}

	if p.Turn[0].Type != BuildTurn {
		return false
	}

	return true
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
		return errors.New("round error : FirstBuild")
	}

	if !p.IsTurn(user) {
		return errors.New("It's not a turn")
	}

	faction := p.Factions[user]

	color := p.Map.GetType(x, y)
	f := faction.GetInstance()
	if color != f.Color {
		return errors.New("It's not a user's land")
	}

	err := faction.FirstBuild(x, y)
	if err != nil {
		return err
	}

	p.Map.Build(x, y, f.Color, f.FirstBuilding)

	return nil
}

func (p *Game) Build(user int, x int, y int, building resources.Building) error {
	log.Println("build", p.Round)
	if p.Round < 1 {
		return errors.New("round error")
	}

	if !p.IsTurn(user) {
		return errors.New("It's not a turn")
	}

	faction := p.Factions[user]
	f := faction.GetInstance()

	if !p.IsNormalTurn() && !((p.IsBuildTurn() || p.IsResourceTurn()) && f.Resource.Building != resources.None) {
		return errors.New("It's not a normal turn")
	}

	if f.Action {
		if f.Resource.Building == resources.None && f.ExtraBuild == 0 {
			return errors.New("Already completed the action")
		}
	}

	flag := p.Map.CheckDistance(f.Color, f.GetShipDistance(true), x, y)

	molesBuild := false

	if flag == false {
		if f.Name == "Moles" {
			if (f.Resource.Building == resources.None && f.Resource.Worker > 0) || (f.Resource.Worker > 0 && f.Resource.Coin >= 2) {
				flag = p.Map.CheckDistanceMoles(f.Color, x, y)

				items := resources.GetGroundPosition(x, y)

				flag := false
				for _, position := range items {
					x := position.X
					y := position.Y

					if p.Map.GetOwner(x, y) == f.Color {
						flag = true
						break
					}

					for _, v := range p.Map.BridgeList {
						if v.Color != f.Color {
							continue
						}

						if x == v.X1 && y == v.Y1 {
							if p.Map.GetOwner(v.X2, v.Y2) == f.Color {
								flag = true
								break
							}
						}

						if x == v.X2 && y == v.Y2 {
							if p.Map.GetOwner(v.X1, v.Y1) == f.Color {
								flag = true
								break
							}
						}
					}

					if flag == true {
						break
					}
				}

				if flag == false {
					molesBuild = true
				}
			}
		}
	}

	if f.Resource.Building == resources.TP {
		if p.Map.GetType(x, y) == f.Color {
			flag = true
		}
	}

	if flag == false {
		return errors.New("ship distance error")
	}

	err := p.Map.CheckBuild(x, y, f.Color, f.GetSpadeCount())

	if err != nil {
		return err
	}

	needSpade := p.Map.GetNeedSpade(x, y, f.Color)
	if needSpade < 0 {
		needSpade = 0
	}

	err = faction.Build(x, y, needSpade, building)
	if err != nil {
		return err
	}

	if molesBuild == true {
		if f.Resource.Building != resources.TP {
			f.Resource.Worker--
			f.VP += 4
		}
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
	if p.Round < 1 {
		return errors.New("round error")
	}

	if !p.IsTurn(user) {
		return errors.New("It's not a turn")
	}

	if !p.IsNormalTurn() {
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
		return errors.New("not owner")
	}

	extra := 0

	if f.Resource.TpUpgrade == 0 && target == resources.TP {
		// 생업 체크
		log.Println("check solor", p.Map.CheckSolo(f.Color, x, y))
		if p.Map.CheckSolo(f.Color, x, y) {
			extra = 3
		}
	}

	err := faction.Upgrade(x, y, target, extra)

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
		return errors.New("round error")
	}

	if !p.IsTurn(user) {
		return errors.New("It's not a turn")
	}

	if !p.IsNormalTurn() {
		return errors.New("It's not a normal turn")
	}

	faction := p.Factions[user]
	f := faction.GetInstance()
	if f.Action {
		return errors.New("Already completed the action")
	}

	if p.PowerActions.IsUse(pos) {
		return errors.New("already")
	}

	have := f.GetHavePowerCount()
	if have < p.PowerActions.GetNeedPower(pos) {
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
		return errors.New("round error")
	}

	if !p.IsTurn(user) {
		return errors.New("It's not a turn")
	}

	if !p.IsNormalTurn() {
		return errors.New("It's not a normal turn")
	}

	faction := p.Factions[user]
	f := faction.GetInstance()
	if f.Action {
		return errors.New("Already completed the action")
	}

	if pos < 0 || pos > 2 {
		return errors.New("not found")
	}

	if p.BookActions.IsUse(pos) {
		return errors.New("already")
	}

	have := f.Resource.Book.Count()
	if have < p.BookActions.GetNeedBook(pos) {
		return errors.New("not enough book")
	}

	item := p.BookActions.Action(pos)
	faction.Book(item, book)

	return nil
}

func (p *Game) AdvanceShip(user int) error {
	if p.Round < 1 {
		return errors.New("round error")
	}

	if !p.IsTurn(user) {
		return errors.New("It's not a turn")
	}

	if !p.IsNormalTurn() {
		return errors.New("It's not a normal turn")
	}

	faction := p.Factions[user]
	f := faction.GetInstance()
	if f.Action {
		return errors.New("Already completed the action")
	}

	err := faction.AdvanceShip()
	if err != nil {
		return err
	}

	buildVP := p.RoundBonuss.GetBuildVP(p.Round)

	f.ReceiveResource(resources.Price{VP: buildVP.Advance})

	return nil
}

func (p *Game) AdvanceSpade(user int) error {
	if p.Round < 1 {
		return errors.New("round error")
	}

	if !p.IsTurn(user) {
		return errors.New("It's not a turn")
	}

	if !p.IsNormalTurn() {
		return errors.New("It's not a normal turn")
	}

	faction := p.Factions[user]
	f := faction.GetInstance()
	if f.Action {
		return errors.New("Already completed the action")
	}

	err := faction.AdvanceSpade()
	if err != nil {
		return err
	}

	buildVP := p.RoundBonuss.GetBuildVP(p.Round)

	f.ReceiveResource(resources.Price{VP: buildVP.Advance})

	return nil
}

func (p *Game) SendScholar(user int, pos ScienceType) error {
	if p.Round < 1 {
		return errors.New("round error")
	}

	if !p.IsTurn(user) {
		return errors.New("It's not a turn")
	}

	if !p.IsNormalTurn() {
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

	faction.SendScholar(int(pos), inc)

	buildVP := p.RoundBonuss.GetBuildVP(p.Round)
	f.ReceiveResource(resources.Price{VP: buildVP.Science * inc})

	return nil
}

func (p *Game) SupployScholar(user int, pos ScienceType) error {
	if p.Round < 1 {
		return errors.New("round error")
	}

	if !p.IsTurn(user) {
		return errors.New("It's not a turn")
	}

	if !p.IsNormalTurn() {
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

	faction.SupployScholar(int(pos), inc)

	buildVP := p.RoundBonuss.GetBuildVP(p.Round)
	f.ReceiveResource(resources.Price{VP: buildVP.Science * inc})

	return nil
}

func (p *Game) Pass(user int, pos int) error {
	if p.Round < 1 {
		return errors.New("round error")
	}

	if !p.IsTurn(user) {
		return errors.New("It's not a turn")
	}

	if !p.IsNormalTurn() {
		return errors.New("It's not a normal turn")
	}

	faction := p.Factions[user]
	roundTile := p.RoundTiles.Pass(pos)
	err, tile := faction.Pass(roundTile)

	if err != nil {
		p.RoundTiles.Add(roundTile)

		return err
	}

	p.RoundTiles.Add(tile)
	p.PassOrder = append(p.PassOrder, user)

	return nil
}

func (p *Game) FactionAction() {
}

func (p *Game) Dig(user int, x int, y int, dig int) error {
	if !p.IsTurn(user) {
		return errors.New("It's not a turn")
	}

	if !p.IsNormalTurn() && !p.IsSpadeTurn() && !p.IsResourceTurn() {
		return errors.New("It's not a normal turn")
	}

	faction := p.Factions[user]
	f := faction.GetInstance()

	if p.Map.GetType(x, y) == color.River {
		return errors.New("can't spade")
	}

	if f.Resource.Spade < dig {
		return errors.New("not have spade")
	}

	needSpade := p.Map.GetNeedSpade(x, y, f.Color)
	if needSpade == 0 {
		return errors.New("need not spade")
	}

	tile := true

	if p.IsSpadeTurn() || p.IsResourceTurn() {
		tile = false
	}

	flag := p.Map.CheckDistance(f.Color, f.GetShipDistance(tile), x, y)

	if flag == false {
		return errors.New("ship distance error")
	}

	spade := 0
	change := 0
	if dig >= needSpade {
		change = int(f.Color)
		spade = needSpade
	} else {
		target := p.Map.GetType(x, y)
		diff := f.Color - target

		if diff > 0 {
			if diff >= 4 {
				change = int(target) - dig
			} else {
				change = int(target) + dig
			}
		} else {
			if diff <= -4 {
				change = int(target) + dig
			} else {
				change = int(target) - dig
			}
		}

		if change <= int(color.River) {
			change += 7
		} else if change > int(color.Gray) {
			change -= 7
		}

		spade = dig
	}

	faction.Dig(x, y, spade)
	p.Map.SetType(x, y, color.Color(change))

	buildVP := p.RoundBonuss.GetBuildVP(p.Round)

	f.ReceiveResource(resources.Price{VP: buildVP.Spade * spade})

	if f.Action == false && p.Round > 0 {
		f.ExtraBuild++
	}

	f.Action = true

	return nil
}

func (p *Game) GetRoundTile(user int, pos int) error {
	if p.Round != RoundTileRound {
		return errors.New("round error : GetRoundTile")
	}

	if !p.IsTurn(user) {
		return errors.New("It's not a turn")
	}

	if !p.IsNormalTurn() {
		return errors.New("It's not a normal turn")
	}

	faction := p.Factions[user]
	tile := p.RoundTiles.Pass(pos)

	faction.RoundTile(tile)

	p.TurnEnd(user)

	return nil
}

func (p *Game) Bridge(user int, x1 int, y1 int, x2 int, y2 int) error {
	if p.Round < 1 {
		return errors.New("round error")
	}

	if !p.IsTurn(user) {
		return errors.New("It's not a turn")
	}

	if !p.IsNormalTurn() {
		return errors.New("It's not a normal turn")
	}

	faction := p.Factions[user]
	f := faction.GetInstance()

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

			if f.Ename == "Moles" {
				flag = true
			} else if p.Map.GetType(x1+1, y) == color.River && p.Map.GetType(x1+1, y+1) == color.River {
				flag = true
			}
		}
	} else {
		if x1%2 == 1 {
			if (x1-x2 == 1 || x1-x2 == -1) && y2-y1 == 2 {
				if f.Ename == "Moles" {
					flag = true
				} else if p.Map.GetType(x1, y1+1) == color.River && p.Map.GetType(x2, y2-1) == color.River {
					flag = true
				}
			}
		} else {
			if (x1-x2 == 1 || x1-x2 == -1) && y2-y1 == 1 {
				if f.Ename == "Moles" {
					flag = true
				} else if p.Map.GetType(x1, y1+1) == color.River && p.Map.GetType(x2, y2-1) == color.River {
					flag = true
				}
			}
		}

	}

	if flag == false {
		return errors.New("Locations that cannot be built")
	}

	err := p.Map.CheckBridge(f.Color, x1, y1, x2, y2)
	if err != nil {
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
			owner := p.Map.GetOwner(position.X, position.Y)

			if owner == f.Color {
				powers[i] += p.Map.GetPower(position.X, position.Y)
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
		power := powers[v]
		if power > 0 {
			target := p.Factions[user].GetInstance()
			if target.VP >= power {
				p.PowerTurn = append(p.PowerTurn, Turn{User: v, Type: PowerTurn, Power: power, From: user})
			}
		}
	}
}

func (p *Game) PowerConfirm(user int, confirm bool) error {
	if p.Round < 1 {
		return errors.New("round error")
	}

	if !p.IsTurn(user) || !p.IsPowerTurn() {
		return errors.New("It's not a turn")
	}

	faction := p.Factions[user]
	f := faction.GetInstance()

	if confirm == true {
		f.ReceivePower(p.Turn[0].Power, true)
	}

	f.Action = true

	return nil
}

func (p *Game) City(user int, city resources.CityType) error {
	if p.Round < 1 {
		return errors.New("round error")
	}

	if !p.IsTurn(user) {
		return errors.New("It's not a turn")
	}

	if !p.IsNormalTurn() {
		return errors.New("It's not a normal turn")
	}

	faction := p.Factions[user]
	f := faction.GetInstance()

	if f.Resource.City == 0 {
		return errors.New("have not city")
	}

	if !p.Cities.IsRemain(city) {
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
		return errors.New("round error")
	}

	if !p.IsTurn(user) {
		return errors.New("It's not a turn")
	}

	if !p.IsNormalTurn() && !p.IsScienceTurn() && !p.IsResourceTurn() {
		return errors.New("It's not a normal turn")
	}

	faction := p.Factions[user]
	f := faction.GetInstance()

	if f.Resource.Science.Any == 0 && f.Resource.Science.Single == 0 {
		return errors.New("have not science")
	}

	if level > 1 {
		if f.Resource.Science.Single < level {
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
	f.IncScience(int(pos), inc)

	buildVP := p.RoundBonuss.GetBuildVP(p.Round)
	f.ReceiveResource(resources.Price{VP: buildVP.Science * inc})

	if p.IsScienceTurn() || p.IsResourceTurn() {
		f.Action = true
	}

	return nil
}

func (p *Game) Book(user int, pos resources.BookType, count int) error {
	if !p.IsTurn(user) {
		return errors.New("It's not a turn")
	}

	if !p.IsNormalTurn() && !p.IsBookTurn() && !p.IsResourceTurn() {
		return errors.New("It's not a normal turn")
	}

	faction := p.Factions[user]
	f := faction.GetInstance()

	if f.Resource.Book.Any == 0 {
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

	if p.IsBookTurn() || p.IsResourceTurn() {
		f.Action = true
	}

	return nil
}

func (p *Game) ConvertDig(user int, spade int) error {
	if p.Round < 1 {
		return errors.New("round error")
	}

	if !p.IsTurn(user) {
		return errors.New("It's not a turn")
	}

	if !p.IsNormalTurn() && !p.IsBuildTurn() {
		return errors.New("It's not a normal turn")
	}

	faction := p.Factions[user]
	f := faction.GetInstance()

	if f.Resource.Spade+spade > 3 {
		return errors.New("over")
	}

	if f.BuildAction == true {
		return errors.New("already build")
	}

	f.ConvertDig(spade)
	return nil
}

func (p *Game) PalaceTile(user int, pos int) error {
	if p.Round < 1 {
		return errors.New("round error")
	}

	if !p.IsTurn(user) {
		return errors.New("It's not a turn")
	}

	if !p.IsNormalTurn() {
		return errors.New("It's not a normal turn")
	}

	if pos >= len(p.PalaceTiles.Items) {
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

func (p *Game) InnovationTile(user int, pos int, index int, book resources.Book) error {
	if p.Round < 1 {
		return errors.New("round error")
	}

	if !p.IsTurn(user) {
		return errors.New("It's not a turn")
	}

	if !p.IsNormalTurn() {
		return errors.New("It's not a normal turn")
	}

	if pos >= len(p.InnovationTiles.Items) {
		return errors.New("not found tile")
	}

	if index >= len(p.InnovationTiles.Items[pos]) {
		return errors.New("not found tile")
	}

	tile := p.InnovationTiles.GetTile(pos, index)

	if tile.Use == true {
		return errors.New("already select")
	}

	faction := p.Factions[user]
	f := faction.GetInstance()
	err := faction.InnovationTile(tile, book)

	if err == nil {
		p.InnovationTiles.Setup(pos, index)
	}

	f.Action = true

	return nil
}

func (p *Game) TileAction(user int, category resources.TileCategory, pos int) error {
	if p.Round < 1 {
		return errors.New("round error")
	}

	if !p.IsTurn(user) {
		return errors.New("It's not a turn")
	}

	if !p.IsNormalTurn() {
		return errors.New("It's not a normal turn")
	}

	faction := p.Factions[user]
	f := faction.GetInstance()

	if f.Action {
		return errors.New("Already completed the action")
	}

	err := faction.TileAction(category, pos)
	if err != nil {
		return err
	}

	return nil
}

func (p *Game) SchoolTile(user int, science int, level int) error {
	if !p.IsTurn(user) {
		return errors.New("It's not a turn")
	}

	if !p.IsNormalTurn() {
		return errors.New("It's not a normal turn")
	}

	if p.SchoolTiles.Items[science][level].Count == 0 {
		return errors.New("tile not remain")
	}

	tile := p.SchoolTiles.Items[science][level].Tile

	faction := p.Factions[user]
	f := faction.GetInstance()
	err := faction.SchoolTile(tile, science)

	if err != nil {
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

	if p.Round > 0 {
		buildVP := p.RoundBonuss.GetBuildVP(p.Round)
		f.ReceiveResource(resources.Price{VP: buildVP.Science * inc})
	}

	p.SchoolTiles.Items[science][level].Count--

	return nil
}

func (p *Game) Burn(user int, count int) error {
	if p.Round < 1 {
		return errors.New("round error")
	}

	if !p.IsTurn(user) {
		return errors.New("It's not a turn")
	}

	if !p.IsNormalTurn() {
		return errors.New("It's not a normal turn")
	}

	faction := p.Factions[user]
	f := faction.GetInstance()
	return f.Burn(count)
}

func (p *Game) Convert(user int, source resources.Price, target resources.Price) error {
	if p.Round < 1 {
		return errors.New("round error")
	}

	if !p.IsTurn(user) {
		return errors.New("It's not a turn")
	}

	if !p.IsNormalTurn() {
		return errors.New("It's not a normal turn")
	}

	faction := p.Factions[user]
	f := faction.GetInstance()
	return f.Convert(source, target)
}

func (p *Game) Annex(user int, x int, y int) error {
	if p.Round < 1 {
		return errors.New("round error")
	}

	if !p.IsTurn(user) {
		return errors.New("It's not a turn")
	}

	if !p.IsNormalTurn() {
		return errors.New("It's not a normal turn")
	}

	faction := p.Factions[user]
	f := faction.GetInstance()

	err := p.Map.CheckAnnex(f.Color, x, y)
	if err != nil {
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

func (p *Game) ClearHistory() {
	p.OldCommand = make([]string, 0)
	p.OldCommand = append(p.OldCommand, p.Command...)
	p.Command = make([]string, 0)
}

func (p *Game) AddHistory(str string) {
	p.Command = append(p.Command, str)
}

func (p *Game) Undo(user int) error {
	if !p.IsTurn(user) {
		return errors.New("It's not a turn")
	}

	if len(p.Command) == 0 {
		return errors.New("history empty")
	}
	conn := models.NewConnection()
	defer conn.Close()

	gamehistoryManager := models.NewGamehistoryManager(conn)
	items := gamehistoryManager.Find([]interface{}{
		models.Where{Column: "game", Value: p.Id, Compare: "="},
		models.Ordering("gh_id desc"),
	})

	if len(items) == 0 {
		return errors.New("history empty")
	}

	last := items[0]

	strs := strings.Split(last.Command, " ")
	if strs[1] == "save" {
		return errors.New("unabled")
	}

	gamehistoryManager.Delete(last.Id)

	MakeGame(p.Id, p.Count)

	return nil
}

type Score struct {
	User  int
	Score int
}

func (p *Game) EndGame() {
	log.Println("EndGame *************************")

	scores := make([]Score, 0)
	// 미션 점수
	for user, v := range p.Factions {
		faction := v.GetInstance()

		total := 0
		for _, d := range faction.Building {
			total += d
		}

		scores = append(scores, Score{User: user, Score: total})
	}

	sort.Slice(scores, func(i, j int) bool {
		return scores[i].Score < scores[j].Score
	})

	receives := []int{0, 18, 15, 12, 9, 7}
	receiveCount := 0

	for i := 0; i < 3; i++ {
		if len(scores) == 0 {
			break
		}

		score := scores[0].Score

		cnt := 0
		for _, v := range scores {
			if v.Score == score {
				cnt++
			}
		}

		receive := receives[cnt]

		for _, v := range scores {
			if v.Score == score {
				faction := p.Factions[v.User]
				f := faction.GetInstance()
				f.ReceiveResource(resources.Price{VP: receive})
			}
		}

		scores = scores[cnt:]
		receiveCount += cnt

		if receiveCount >= 3 {
			break
		}

		if receiveCount == 1 {
			receives = []int{0, 12, 9, 6, 4}
		} else if receiveCount == 2 {
			receives = []int{0, 6, 3, 2}
		}
	}

	// 과학 점수
	for k := 0; k < 4; k++ {
		scores = make([]Score, 0)

		for user, v := range p.Factions {
			faction := v.GetInstance()

			scores = append(scores, Score{User: user, Score: faction.Science[k]})
		}

		sort.Slice(scores, func(i, j int) bool {
			return scores[i].Score < scores[j].Score
		})

		receives := []int{0, 8, 4, 2}
		receiveCount := 0

		for i := 0; i < 3; i++ {
			if len(scores) == 0 {
				break
			}

			score := scores[0].Score

			cnt := 0
			for _, v := range scores {
				if v.Score == score {
					cnt++
				}
			}

			receive := receives[cnt]

			for _, v := range scores {
				if v.Score == score {
					faction := p.Factions[v.User]
					f := faction.GetInstance()
					f.ReceiveResource(resources.Price{VP: receive})
				}
			}

			scores = scores[cnt:]
			receiveCount += cnt

			if receiveCount >= 3 {
				break
			}

			if receiveCount == 1 {
				receives = []int{0, 4, 2, 1}
			} else if receiveCount == 2 {
				receives = []int{0, 2, 1, 0}
			}
		}
	}
	// 등수 계산
}
