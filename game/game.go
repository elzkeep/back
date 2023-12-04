package game

import (
	"aoi/game/action"
	"aoi/game/color"
	"aoi/game/factions"
	"aoi/game/resources"
	"aoi/global"
	"aoi/models"
	"aoi/models/game"
	gm "aoi/models/game"
	"errors"
	"log"
	"math"
	"sort"
	"strings"
)

const (
	InitRound       = -7
	FactionRound    = -6
	RoundTileRound  = -5
	BuildRound      = -4
	TileRound       = -3
	ExtraBuildRound = -2
	SpadeRound      = -1
)

type GameType int

const (
	NoneType GameType = iota
	BasicType
	DraftBasicType
	DraftSnakeType
)

type BuildType int

const (
	NormalBuild BuildType = iota
	FlyBuild
	TunnelingBuild
)

type UndoRequest struct {
	Id      int64  `json:"id"`
	User    int    `json:"user"`
	History int64  `json:"history"`
	Status  int    `json:"status"`
	Command string `json:"command"`
	Users   []int  `json:"users"`
}

type ReplayLog struct {
	Round   int
	User    int64
	History int64
	Command string
}

type Game struct {
	Id       int64                       `json:"id"`
	Name     string                      `json:"name"`
	Type     GameType                    `json:"type"`
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
	Usernames       []string                  `json:"usernames"`
	Count           int                       `json:"count"`
	Logs            []Log                     `json:"logs"`
	Command         []string                  `json:"commands"`
	OldCommand      []string                  `json:"oldCommands"`
	Network         int                       `json:"network"`
	UndoRequest     UndoRequest               `json:"undo"`
	Illusionists    gm.Illusionists           `json:"illusionists"`
	DummyNetwork    int                       `json:"-"`
	HistoryId       int64                     `json:"-"`
	Replay          bool                      `json:"replay"`
	ReplayPos       []int                     `json:"replayPos"`
	MolesBridge     bool                      `json:"molesBridge"`
	Replays         []ReplayLog               `json:"-"`
	Mapid           int64                     `json:"-"`
}

type Log struct {
	Id   int64 `json:"id"`
	User int   `json:"user"`

	Round   int      `json:"round"`
	Command []string `json:"command"`
	VP      int      `json:"vp"`
}

func NewGame(game *models.Game) *Game {
	var item Game
	id := game.Id
	count := game.Count
	typeid := game.Type
	mapid := game.Map

	item.Id = id
	item.Count = count
	item.Type = GameType(typeid)
	item.Name = game.Name
	item.Mapid = mapid
	item.PowerActions = action.NewPowerAction()
	item.BookActions = action.NewBookAction(id)
	item.RoundTiles = resources.NewRoundTile(id, typeid)
	item.RoundBonuss = NewRoundBonus(id)
	item.PalaceTiles = resources.NewPalaceTile(id)
	item.SchoolTiles = resources.NewSchoolTile(id, count)
	item.InnovationTiles = resources.NewInnovationTile(id, count)
	item.FactionTiles = resources.NewFactionTile(id)
	item.ColorTiles = resources.NewColorTile(id)
	item.Cities = NewCity()
	item.Replay = false
	item.Replays = make([]ReplayLog, 0)
	item.Illusionists = game.Illusionists
	item.MolesBridge = false

	item.Network = 0
	network := 0
	for _, v := range item.PalaceTiles.Items {
		network += int(v.Type)
	}

	item.DummyNetwork = network%4 + 12

	item.Map = NewMap(id, mapid)
	item.Sciences = NewScience(count)
	item.Factions = make([]factions.FactionInterface, 0)

	item.Turn = make([]Turn, 0)
	item.PowerTurn = make([]Turn, 0)
	item.PassOrder = make([]int, 0)
	item.TurnOrder = make([]int, 0)
	item.Users = make([]int64, 0)
	item.Usernames = make([]string, 0)

	item.Command = make([]string, 0)
	item.OldCommand = make([]string, 0)
	item.Logs = make([]Log, 0)
	item.UndoRequest = UndoRequest{Users: make([]int, 0)}
	item.ReplayPos = make([]int, 8)

	item.Round = InitRound

	return &item
}

func (p *Game) Copy() *Game {
	var item Game

	id := p.Id
	mapid := p.Mapid
	count := p.Count

	item.Id = p.Id
	item.Count = p.Count
	item.Type = p.Type
	item.Name = p.Name
	item.Mapid = p.Mapid

	item.PowerActions = p.PowerActions.Copy()
	item.BookActions = p.BookActions.Copy()
	item.RoundTiles = p.RoundTiles.Copy()
	item.RoundBonuss = p.RoundBonuss.Copy()
	item.PalaceTiles = p.PalaceTiles.Copy()
	item.SchoolTiles = p.SchoolTiles.Copy()
	item.InnovationTiles = p.InnovationTiles.Copy()
	item.FactionTiles = p.FactionTiles.Copy()
	item.ColorTiles = p.ColorTiles.Copy()
	item.Cities = p.Cities.Copy()
	item.Replay = true
	item.Illusionists = p.Illusionists
	item.MolesBridge = false

	item.Network = 0
	network := 0
	for _, v := range item.PalaceTiles.Items {
		network += int(v.Type)
	}

	item.DummyNetwork = network%4 + 12

	item.Map = NewMap(id, mapid)
	item.Sciences = NewScience(count)
	item.Factions = make([]factions.FactionInterface, 0)

	item.Turn = make([]Turn, 0)
	item.PowerTurn = make([]Turn, 0)
	item.PassOrder = make([]int, 0)
	item.TurnOrder = make([]int, 0)
	item.Users = make([]int64, 0)
	item.Usernames = make([]string, 0)

	item.Command = make([]string, 0)
	item.OldCommand = make([]string, 0)
	item.Logs = make([]Log, 0)
	item.UndoRequest = UndoRequest{Users: make([]int, 0)}
	item.ReplayPos = make([]int, 8)

	copy(item.ReplayPos, p.ReplayPos)

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

func (p *Game) GetUserPos(user int64) int {
	for i, v := range p.Users {
		if v == user {
			return i
		}
	}

	return -1
}

func (p *Game) AddUser(user int64, name string) {
	p.Users = append(p.Users, user)
	p.Usernames = append(p.Usernames, name)
	item := &factions.Basic{}
	colorTile := p.ColorTiles.Items[0]
	item.Init(colorTile, name)
	p.Factions = append(p.Factions, item)
}

func (p *Game) UpdateDBRound(value int) {
	if p.Replay == true {
		return
	}

	conn := models.NewConnection()
	defer conn.Close()

	gameManager := models.NewGameManager(conn)
	game := gameManager.Get(p.Id)
	if int(game.Status) < value {
		gameManager.UpdateStatus(value, p.Id)
	}
}

func (p *Game) CompleteAddUser() {
	p.Round = FactionRound

	if p.Type == BasicType {
		for i := 0; i < p.Count; i++ {
			p.Turn = append(p.Turn, Turn{User: p.Count - i - 1, Type: NormalTurn})
		}
	} else if p.Type == DraftBasicType {
		for i := 0; i < p.Count; i++ {
			p.Turn = append(p.Turn, Turn{User: i, Type: NormalTurn})
		}

		for i := 0; i < p.Count; i++ {
			p.Turn = append(p.Turn, Turn{User: i, Type: NormalTurn})
		}

		for i := 0; i < p.Count; i++ {
			p.Turn = append(p.Turn, Turn{User: i, Type: NormalTurn})
		}
	} else if p.Type == DraftSnakeType {
		for i := 0; i < p.Count; i++ {
			p.Turn = append(p.Turn, Turn{User: i, Type: NormalTurn})
		}

		for i := 0; i < p.Count; i++ {
			p.Turn = append(p.Turn, Turn{User: p.Count - i - 1, Type: NormalTurn})
		}

		for i := 0; i < p.Count; i++ {
			p.Turn = append(p.Turn, Turn{User: i, Type: NormalTurn})
		}
	}

	p.UpdateDBRound(int(game.StatusFaction))
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

	item.Init(colorTile, p.Usernames[user])

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

	if p.Replay == false {
		conn := models.NewConnection()
		defer conn.Close()

		gameuserManager := models.NewGameuserManager(conn)
		gameuser := gameuserManager.GetByGameUser(p.Id, p.Users[user])
		gameuser.Faction = int(factionTile.Type)
		gameuser.Color = int(colorTile.Color)

		gameuserManager.Update(gameuser)
	}
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
		if faction.Type != resources.TileFactionOmar {
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
	p.RoundTiles.Start()

	p.Round++

	if p.Round == 6 {
		p.Network = p.DummyNetwork
	}

	p.TurnOrder = p.PassOrder
	p.PassOrder = make([]int, 0)

	if p.Round >= 1 {
		for i := range p.Factions {
			user := p.TurnOrder[i]
			faction := p.Factions[user]
			f := faction.GetInstance()
			f.Income()

			// income 계산
			f.CalulateReceive()
			p.Sciences.CalculateRoundBonus(f)
			if p.Round < 6 {
				p.Sciences.CalculateRoundEndBonus(faction, p.RoundBonuss.Items[p.Round-1])
			}

			/*
				if p.Round > 1 {
					p.Sciences.RoundBonus(f)
				}
			*/
		}
	}

	for i := range p.Factions {
		user := p.TurnOrder[i]
		faction := p.Factions[user]
		f := faction.GetInstance()

		if f.Resource.Book.Any > 0 ||
			!f.Resource.Science.IsEmpty() ||
			f.Resource.Spade > 0 ||
			f.Resource.Building != resources.None {

			turn := Turn{User: user, Type: ResourceTurn}
			p.Turn = append(p.Turn, turn)

			/*
				if p.Round > 1 {
					if f.Resource.Science.IsEmpty() {
						p.Sciences.RoundBonus(f)
					}
				}
			*/
		} else {
			if p.Round > 1 {
				p.Sciences.RoundBonus(f)
			}
		}
	}

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

func (p *Game) RoundProcess() {
	if p.Round == FactionRound {
		if p.Type != BasicType {
			p.InitDraft()

			p.Round = RoundTileRound
			p.PassOrder = make([]int, 0)

			for i := len(p.Factions) - 1; i >= 0; i-- {
				p.Turn = append(p.Turn, Turn{User: i, Type: NormalTurn})
			}

			return
		}

		p.Round = BuildRound
		p.PassOrder = make([]int, 0)

		p.UpdateDBRound(int(game.StatusNormal))

		p.BuildStart()
		return
	}

	if p.Round == RoundTileRound {
		p.Round = BuildRound
		p.PassOrder = make([]int, 0)

		p.UpdateDBRound(int(game.StatusNormal))

		p.BuildStart()
		return
	}

	if p.Round == BuildRound {
		p.Round = TileRound

		for i, v := range p.Factions {
			faction := v.GetInstance()
			if faction.Type == resources.TileFactionInventors || faction.Type == resources.TileFactionMonks {
				p.Turn = append(p.Turn, Turn{User: i, Type: TileTurn})
			}
		}

		if len(p.Turn) > 0 {
			return
		}
	}

	if p.Round == TileRound {
		p.Round = SpadeRound

		for i, v := range p.Factions {
			f := v.GetInstance()

			f.Resource.Building = resources.None

			if f.Resource.Spade > 0 {
				p.Turn = append(p.Turn, Turn{User: i, Type: SpadeTurn})
			}
		}

		if len(p.Turn) > 0 {
			return
		}
	}

	if p.Round == SpadeRound {
		p.Round = 0
		p.PassOrder = make([]int, 0)

		for i := range p.Factions {
			p.PassOrder = append(p.PassOrder, i)
		}

		science := resources.Science{Banking: 2, Law: 2, Engineering: 2, Medicine: 2}
		for i := 1; i <= 5; i++ {
			item := p.RoundBonuss.Get(i)
			science.Banking += item.Science.Banking
			science.Law += item.Science.Law
			science.Engineering += item.Science.Engineering
			science.Medicine += item.Science.Medicine
		}
		p.Sciences.Init([]int{science.Banking, science.Law, science.Engineering, science.Medicine})

		p.Start()
	} else if p.Round == 6 {
		p.Round++
		p.EndGame()
	} else {
		p.RoundEnd()
		p.Start()
	}
}

func (p *Game) Calculate(user int) {
	if p.Round >= 6 {
		for _, v := range p.Factions {
			f := v.GetInstance()
			f.CalulateVP()
		}

		p.CalculateEndGame()
	} else if p.Round > 0 {
		faction := p.Factions[user]
		f := faction.GetInstance()

		f.CalulateReceive()
		p.Sciences.CalculateRoundBonus(f)

		if p.Round < 6 {
			p.Sciences.CalculateRoundEndBonus(faction, p.RoundBonuss.Items[p.Round-1])
		}
	}
}

func (p *Game) TurnEnd(user int) error {
	if !p.IsTurn(user) {
		return errors.New("It's not a turn")
	}

	faction := p.Factions[user]
	f := faction.GetInstance()

	err := faction.TurnEnd(p.Round, p.IsNormalTurn())
	if err != nil {
		return err
	}

	if p.IsResourceTurn() && p.Round > 1 {
		p.Sciences.RoundBonus(f)
	}

	turnType := p.Turn[0].Type

	p.Turn = p.Turn[1:]

	p.Turn = append(p.PowerTurn, p.Turn...)
	p.PowerTurn = make([]Turn, 0)

	p.Map.TurnEnd()

	if p.Round > 0 {
		if user >= 0 {
			if turnType == NormalTurn {
				if !f.IsPass {
					p.Turn = append(p.Turn, Turn{User: user, Type: NormalTurn})
				}
			}
		}
	}

	if len(p.Turn) == 0 {
		p.RoundProcess()
	}

	p.MolesBridge = false

	return nil
}

func (p *Game) RoundEnd() {
	for _, v := range p.Factions {
		p.Sciences.RoundEndBonus(v, p.RoundBonuss.Items[p.Round-1])
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

func (p *Game) IsTileTurn() bool {
	if len(p.Turn) == 0 {
		return false
	}

	if p.Turn[0].Type != TileTurn {
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

func (p *Game) FirstBuild(user int, x int, y int, building resources.Building) error {
	if p.Round != BuildRound {
		return errors.New("round error")
	}

	if !p.IsTurn(user) {
		return errors.New("It's not a turn")
	}

	faction := p.Factions[user]
	f := faction.GetInstance()

	if f.Type == resources.TileFactionOmar {
		if building != resources.D && building != resources.WHITE_TOWER {
			return errors.New("wrong building")
		}

		if building == resources.D && f.Building[resources.D] >= 2 {
			return errors.New("wrong building")
		}

		if building == resources.WHITE_TOWER && f.Building[resources.WHITE_TOWER] >= 1 {
			return errors.New("wrong building")
		}
	} else {
		if building != f.FirstBuilding {
			return errors.New("wrong building")
		}

		if f.Building[building] >= 2 {
			return errors.New("max first building")
		}
	}

	color := p.Map.GetType(x, y)

	if color != f.Color {
		return errors.New("It's not a user's land")
	}

	err := faction.FirstBuild(x, y, building)
	if err != nil {
		return err
	}

	p.Map.Build(x, y, f.Color, building)

	return nil
}

func (p *Game) CheckDistance(f *factions.Faction, x int, y int, tile bool) bool {
	if p.Round < 1 {
		tile = false
	}

	flag := p.Map.CheckDistance(f.Color, f.GetShipDistance(tile), x, y)

	return flag
}

func (p *Game) IsMoles(f *factions.Faction) bool {
	if f.Type != resources.TileFactionMoles {
		return false
	}

	return true
}

func (p *Game) CheckDistanceMoles(f *factions.Faction, x int, y int) bool {
	if !p.IsMoles(f) {
		return false
	}

	if !p.Map.CheckDistanceMoles(f.Color, x, y) {
		return false
	}

	items := resources.GetGroundPosition(x, y)

	lastBuild := p.Map.LastBuild
	flag := true

	for _, position := range items {
		x := position.X
		y := position.Y

		if lastBuild.X == x && lastBuild.Y == y {
			continue
		}

		if p.Map.GetOwner(x, y) == f.Color {
			flag = false
			break
		}

		for _, v := range p.Map.BridgeList {
			if v.Color != f.Color {
				continue
			}

			if x == v.X1 && y == v.Y1 {
				if lastBuild.X == v.X2 && lastBuild.Y == v.Y2 {
					continue
				}

				if p.Map.GetOwner(v.X2, v.Y2) == f.Color {
					flag = false
					break
				}
			}

			if x == v.X2 && y == v.Y2 {
				if lastBuild.X == v.X1 && lastBuild.Y == v.Y1 {
					continue
				}

				if p.Map.GetOwner(v.X1, v.Y1) == f.Color {
					flag = false
					break
				}
			}
		}

		if flag == false {
			break
		}
	}

	return flag
}

func (p *Game) IsJump(f *factions.Faction) bool {
	if !f.CheckTile(resources.TilePalaceJump) {
		return false
	}

	return true
}

func (p *Game) CheckDistanceJump(f *factions.Faction, x int, y int) bool {
	if !p.IsJump(f) {
		return false
	}

	if !p.Map.CheckDistanceJump(f.BuildingList, x, y) {
		return false
	}

	lastBuild := p.Map.LastBuild
	flag := true

	items := resources.GetGroundPosition(x, y)

	for _, position := range items {
		x := position.X
		y := position.Y

		if lastBuild.X == x && lastBuild.Y == y {
			continue
		}

		if p.Map.GetOwner(x, y) == f.Color {
			flag = false
			break
		}

		for _, v := range p.Map.BridgeList {
			if v.Color != f.Color {
				continue
			}

			if x == v.X1 && y == v.Y1 {
				if lastBuild.X == v.X2 && lastBuild.Y == v.Y2 {
					continue
				}

				if p.Map.GetOwner(v.X2, v.Y2) == f.Color {
					flag = false
					break
				}
			}

			if x == v.X2 && y == v.Y2 {
				if lastBuild.X == v.X1 && lastBuild.Y == v.Y1 {
					continue
				}

				if p.Map.GetOwner(v.X1, v.Y1) == f.Color {
					flag = false
					break
				}
			}
		}

		if flag == false {
			break
		}
	}

	return flag
}

func (p *Game) Build(user int, x int, y int, building resources.Building, extra BuildType) error {
	if p.Round < -1 && p.Round != TileRound {
		return errors.New("round error")
	}

	if !p.IsTurn(user) {
		return errors.New("It's not a turn")
	}

	faction := p.Factions[user]
	f := faction.GetInstance()

	if !p.IsNormalTurn() && !((p.IsBuildTurn() || p.IsResourceTurn() || p.IsTileTurn()) && f.Resource.Building != resources.None) {
		return errors.New("It's not a normal turn")
	}

	riverCity := false
	if f.CheckTile(resources.TilePalaceRiverCity) && p.Map.GetType(x, y) == color.River {
		riverCity = true
	}

	if f.Action {
		if building != resources.CITY && f.Resource.Building == resources.None && f.ExtraBuild == 0 {
			return errors.New("Already completed the action")
		}
	}

	if p.Turn[0].Type == SpadeTurn {
		extra = NormalBuild
	}

	flag := false
	if extra == FlyBuild {
		flag = p.CheckDistanceJump(f, x, y)
	} else if extra == TunnelingBuild {
		flag = p.CheckDistanceMoles(f, x, y)
	} else {
		flag = p.CheckDistance(f, x, y, true)
	}

	if f.Resource.Building == resources.TP {
		if p.Map.GetType(x, y) == f.Color {
			flag = true
		}
	}

	if flag == false {
		return errors.New("ship distance error")
	}

	if riverCity {
		lists := p.Map.CheckCity(f.Color, x, y, f.TownPower)

		if len(lists) == 0 {
			return errors.New("can't create a city")
		}

		f.CityBuildingList = lists
		f.Resource.City++

		p.Map.SetOwner(x, y, f.Color)
		p.Map.SetBuilding(x, y, resources.CITY)

		return nil
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

	if extra == TunnelingBuild {
		f.Resource.Worker--
		f.VP += 4
		f.IsJump = true
	}

	if extra == FlyBuild {
		f.Resource.Prist--
		f.VP += 5
		f.IsJump = true
	}

	p.Map.Build(x, y, f.Color, building)

	if f.Resource.Building != resources.None {
		p.Map.ResetLastPosition()
	}

	lists := p.Map.CheckCity(f.Color, x, y, f.TownPower)

	if len(lists) > 0 {
		f.CityBuildingList = lists
		f.Resource.City++
	}

	buildVP := p.RoundBonuss.GetBuildVP(p.Round)

	if p.Map.IsRiverside(x, y) && (building == resources.D || building == resources.WHITE_D) {
		f.ReceiveRiverVP()
	}

	if p.Map.IsEdge(x, y) && (building == resources.D || building == resources.WHITE_D) {
		f.ReceiveEdgeVP()
	}

	if building == resources.D || building == resources.WHITE_D {
		f.ReceiveResource(resources.Price{VP: buildVP.D})
	} else if building == resources.TP || building == resources.WHITE_TP {
		f.ReceiveResource(resources.Price{VP: buildVP.TP})
	} else if building == resources.WHITE_TE {
		f.ReceiveResource(resources.Price{VP: buildVP.TE})
	} else if building == resources.WHITE_SH || building == resources.WHITE_SA {
		f.ReceiveResource(resources.Price{VP: buildVP.SHSA})
	}

	f.ReceiveResource(resources.Price{VP: buildVP.Spade * needSpade})

	if p.Round == 6 {
		buildVP = p.RoundBonuss.FinalRound.Build

		if p.Map.IsEdge(x, y) && (building == resources.D || building == resources.WHITE_D) {
			f.ReceiveResource(resources.Price{VP: buildVP.Edge})
		}

		if building == resources.D || building == resources.WHITE_D {
			f.ReceiveResource(resources.Price{VP: buildVP.D})
		} else if building == resources.TP || building == resources.WHITE_TP {
			f.ReceiveResource(resources.Price{VP: buildVP.TP})
		} else if building == resources.WHITE_TE {
			f.ReceiveResource(resources.Price{VP: buildVP.TE})
		} else if building == resources.WHITE_SH || building == resources.WHITE_SA {
			f.ReceiveResource(resources.Price{VP: buildVP.SHSA})
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

func (p *Game) Downgrade(user int, x int, y int) error {
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

	if f.Resource.Downgrade == 0 || p.Map.GetBuilding(x, y) != resources.TE {
		return errors.New("Already completed the action")
	}

	if p.Map.GetOwner(x, y) != f.Color {
		return errors.New("not owner")
	}

	err := faction.Downgrade(x, y)

	if err != nil {
		return err
	}

	p.Map.SetBuilding(x, y, resources.TP)

	buildVP := p.RoundBonuss.GetBuildVP(p.Round)

	f.ReceiveResource(resources.Price{VP: buildVP.TP})

	if p.Round == 6 {
		buildVP = p.RoundBonuss.FinalRound.Build

		f.ReceiveResource(resources.Price{VP: buildVP.TP})
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
	if f.Type == resources.TileFactionIllusionists {
		have++
	}

	if have < p.PowerActions.GetNeedPower(pos) {
		return errors.New("not enough power")
	}

	item := p.PowerActions.Action(pos)
	err := faction.PowerAction(item)
	if err != nil {
		return err
	}

	if f.CheckTile(resources.TileFactionIllusionists) {
		if p.Illusionists != gm.IllusionistsVp0 {
			vp := 3

			if p.Illusionists == gm.IllusionistsVp2 {
				vp = 2
			} else if p.Illusionists == gm.IllusionistsVp1 {
				vp = 1
			}

			if len(p.Factions) == 5 {
				vp++
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
	f := faction.GetInstance()

	roundTile := resources.TileItem{}

	if pos != -1 {
		roundTile = p.RoundTiles.Pass(pos)
	}

	err, tile := f.Pass(roundTile)

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

func (p *Game) Dig(user int, x int, y int, dig int, extra BuildType) error {
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
		extra = NormalBuild
	}

	flag := false
	if extra == FlyBuild {
		flag = p.CheckDistanceJump(f, x, y)
	} else if extra == TunnelingBuild {
		flag = p.CheckDistanceMoles(f, x, y)
	} else {
		flag = p.CheckDistance(f, x, y, tile)
	}

	if flag == false {
		return errors.New("ship distance error")
	}

	if extra == TunnelingBuild {
		f.Resource.Worker--
		f.VP += 4
		f.IsJump = true
	}

	if extra == FlyBuild {
		f.Resource.Prist--
		f.VP += 5
		f.IsJump = true
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

	flag = false
	if f.Resource.Building != resources.None && f.Resource.ConvertSpade == 0 {
		flag = true
	}

	faction.Dig(x, y, spade)
	p.Map.Dig(x, y, color.Color(change))

	if flag == true {
		p.Map.ResetLastPosition()
	}

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
	f := faction.GetInstance()

	tile := p.RoundTiles.Pass(pos)

	f.AddTile(tile)
	f.Action = true

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

	molesFlag := false
	if f.Type == resources.TileFactionMoles && p.MolesBridge == true {
		molesFlag = true
	}
	log.Println("moles flag", molesFlag)

	if y1 == y2 {
		if x1 > x2 {
			x1, x2 = x2, x1
		}
		if x2-x1 == 2 {
			y := y1
			if x1%2 == 0 {
				y--
			}

			if molesFlag == true {
				flag = true
			} else if (p.Map.GetType(x1+1, y) == color.River || p.Map.GetType(x1+1, y) == color.None) && (p.Map.GetType(x1+1, y+1) == color.River || p.Map.GetType(x1+1, y+1) == color.None) {
				flag = true
			}
		}
	} else {
		if x1%2 == 1 {
			if (x1-x2 == 1 || x1-x2 == -1) && y2-y1 == 2 {
				if molesFlag == true {
					flag = true
				} else if (p.Map.GetType(x1, y1+1) == color.River || p.Map.GetType(x1, y1+1) == color.None) && (p.Map.GetType(x2, y2-1) == color.River || p.Map.GetType(x2, y2-1) == color.None) {
					flag = true
				}
			}
		} else {
			if (x1-x2 == 1 || x1-x2 == -1) && y2-y1 == 1 {
				if molesFlag == true {
					flag = true
				} else if (p.Map.GetType(x1, y1+1) == color.River || p.Map.GetType(x1, y1+1) == color.None) && (p.Map.GetType(x2, y2-1) == color.River || p.Map.GetType(x2, y2-1) == color.None) {
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

	if p.Map.GetOwner(x1, y1) == f.Color {
		lists := p.Map.CheckCity(f.Color, x1, y1, f.TownPower)

		if len(lists) > 0 {
			f.CityBuildingList = lists
			f.Resource.City++
		}
	} else if p.Map.GetOwner(x2, y2) == f.Color {
		lists := p.Map.CheckCity(f.Color, x2, y2, f.TownPower)

		if len(lists) > 0 {
			f.CityBuildingList = lists
			f.Resource.City++
		}
	}

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

		for _, v := range p.Map.BridgeList {
			if x == v.X1 && y == v.Y1 {
				if p.Map.GetOwner(v.X2, v.Y2) == f.Color {
					powers[i] += p.Map.GetPower(v.X2, v.Y2)
				}

				break
			}

			if x == v.X2 && y == v.Y2 {
				if p.Map.GetOwner(v.X1, v.Y1) == f.Color {
					powers[i] += p.Map.GetPower(v.X1, v.Y1)
				}

				break
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
			target := p.Factions[v].GetInstance()

			max := target.Resource.Power[0]*2 + target.Resource.Power[1]
			if max > 0 {
				if power > max {
					power = max
				}

				if target.VP >= power {
					if p.Round == 6 && target.IsPass == true {
						if power == 1 {
							target.ReceivePower(1, false)
						}
					} else {
						p.PowerTurn = append(p.PowerTurn, Turn{User: v, Type: PowerTurn, Power: power, From: user})
					}
				}
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
	inc1, inc2, inc3, inc4 := p.Sciences.Receive(f, tile.Receive)

	f.IncScience(int(Banking), inc1)
	f.IncScience(int(Law), inc2)
	f.IncScience(int(Engineering), inc3)
	f.IncScience(int(Medicine), inc4)

	p.Map.AddCityBuildingList(f.CityBuildingList)
	f.CityBuildingList = make([]resources.Position, 0)

	buildVP := p.RoundBonuss.GetBuildVP(p.Round)

	f.ReceiveResource(resources.Price{VP: buildVP.City})

	p.Map.ResetLastPosition()

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

	if f.BuildAction == true && f.Resource.Building == resources.None {
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
	f := faction.GetInstance()

	err := faction.PalaceTile(tile)

	if err == nil {
		p.PalaceTiles.Setup(pos)
	}

	if tile.Type == resources.TilePalaceRiverCity {
		buildVP := p.RoundBonuss.GetBuildVP(p.Round)

		f.ReceiveResource(resources.Price{VP: buildVP.Advance * 2})
	}

	if tile.Type == resources.TilePalace6PowerCity {
		already := make([]resources.Position, 0)

		for _, v := range f.BuildingList {
			flag := false

			for _, v2 := range already {
				if v.X == v2.X && v.Y == v2.Y {
					flag = true
					break
				}
			}

			if flag == true {
				continue
			}

			lists := p.Map.CheckCity(f.Color, v.X, v.Y, f.TownPower)

			if len(lists) > 0 {
				p.Map.AddCityBuildingList(lists)
				f.Resource.City++
				already = append(already, lists...)
			} else {
				already = append(already, v)
			}
		}
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

	if err != nil {
		return err
	}

	if tile.Type == resources.TileInnovationCluster {
		buildingList := make([]resources.Position, 0)
		buildingList = append(buildingList, f.BuildingList...)

		cnt := 0
		for {
			if len(buildingList) == 0 {
				break
			}

			v := buildingList[0]
			lists := p.Map.GetBuildingList(f.Color, v.X, v.Y, make([]resources.Position, 0))
			items := resources.Unique(lists)

			for _, item := range items {
				for pos, v2 := range buildingList {
					if item.X == v2.X && item.Y == v2.Y {
						buildingList = append(buildingList[:pos], buildingList[pos+1:]...)
						break
					}
				}
			}
			cnt++
		}

		vp := 0
		if cnt >= 6 {
			vp = 18
		} else if cnt >= 5 {
			vp = 12
		} else if cnt >= 4 {
			vp = 8
		}
		f.ReceiveResource(resources.Price{VP: vp})
	}

	p.InnovationTiles.Setup(pos, index)

	if tile.Type == resources.TileInnovationUpgrade {
		buildVP := p.RoundBonuss.GetBuildVP(p.Round)

		f.ReceiveResource(resources.Price{VP: buildVP.Advance * 2})
	}

	buildVP := p.RoundBonuss.GetBuildVP(p.Round)
	f.ReceiveResource(resources.Price{VP: buildVP.Innovation})

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

	if category == resources.TileFaction && int(resources.TileSchoolPassPrist)+pos == int(resources.TileFactionMoles) {
		p.MolesBridge = true
	}

	return nil
}

func (p *Game) SchoolTile(user int, science int, level int) error {
	if !p.IsTurn(user) {
		return errors.New("It's not a turn")
	}

	if !p.IsNormalTurn() && !p.IsTileTurn() {
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

	f.IncScience(science, inc)

	p.SchoolTiles.Items[science][level].Count--

	p.Map.ResetLastPosition()

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

func (p *Game) ClearHistory(id int64, user int, str string, userid int64, round int) {
	faction := p.Factions[user]
	f := faction.GetInstance()

	p.Logs = append(p.Logs, Log{Id: p.HistoryId, User: user, VP: f.VP, Round: round, Command: p.Command})

	p.OldCommand = make([]string, 0)
	p.OldCommand = append(p.OldCommand, p.Command...)
	p.Command = make([]string, 0)

	p.Replays = append(p.Replays, ReplayLog{User: userid, Round: round, Command: str, History: id})
}

func (p *Game) AddHistory(id int64, str string, userid int64, round int) {
	if len(p.Command) == 0 {
		p.HistoryId = id
	}
	p.Command = append(p.Command, str)

	p.Replays = append(p.Replays, ReplayLog{User: userid, Round: round, Command: str, History: id})
}

func (p *Game) Undo(user int) error {
	if p.Replay == true {
		return nil
	}

	if !p.IsTurn(user) {
		return errors.New("It's not a turn")
	}

	if len(p.Command) == 0 {
		return errors.New("history empty")
	}

	game := p.Copy()
	game.Replay = false

	log.Println(p.RoundBonuss.FinalRound)
	log.Println(p.RoundBonuss.OriginalFinalRound)
	game.PowerActions.Original = p.PowerActions.Original
	game.BookActions.Original = p.BookActions.Original
	game.RoundTiles.Original = p.RoundTiles.Original
	game.RoundTiles.OriginalReserved = p.RoundTiles.OriginalReserved
	game.RoundBonuss.OriginalItems = p.RoundBonuss.OriginalItems
	game.RoundBonuss.OriginalTiles = p.RoundBonuss.OriginalTiles
	game.RoundBonuss.OriginalFinalRound = p.RoundBonuss.OriginalFinalRound

	log.Println(game.RoundBonuss.FinalRound)

	game.PalaceTiles.Original = p.PalaceTiles.Original
	game.SchoolTiles.Original = p.SchoolTiles.Original
	game.InnovationTiles.Original = p.InnovationTiles.Original
	game.FactionTiles.Original = p.FactionTiles.Original
	game.ColorTiles.Original = p.ColorTiles.Original
	game.Cities.Original = p.Cities.Original

	for i, v := range p.Users {
		game.AddUser(v, p.Usernames[i])
	}

	game.CompleteAddUser()

	if len(p.Replays) == 0 {
		return errors.New("history empty")
	}

	last := p.Replays[len(p.Replays)-1]
	if last.Command[2:] == "save" {
		return errors.New("unabled")
	}

	for _, item := range p.Replays[:len(p.Replays)-1] {
		Command(game, p.Id, item.User, item.Command, false, item.History)
	}

	conn := models.NewConnection()
	defer conn.Close()

	gamehistoryManager := models.NewGamehistoryManager(conn)
	gamehistoryManager.Delete(last.History)

	_rooms[p.Id] = game

	/*

		db := models.NewConnection()
		defer db.Close()

		conn, _ := db.Begin()
		defer conn.Rollback()

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

		conn.Commit()

		MakeGame(p.Id)
	*/

	return nil
}

type Score struct {
	User  int
	Score int
	Rank  int
}

func (p *Game) LastRoundVP(calculate bool) {
	// 미션 점수
	scores := make([]Score, 0)
	for user, v := range p.Factions {
		f := v.GetInstance()

		remain := f.BuildingList
		items := make([]resources.Position, 0)
		items = append(items, remain[0])
		remain = remain[1:]

		max := 0

		for {
			connect := make([]resources.Position, 0)
			disconnect := make([]resources.Position, 0)

			for _, b := range items {
				for _, b2 := range remain {
					distance := p.Map.CheckConnect(f.Color, f.GetShipDistance(false), b.X, b.Y, b2.X, b2.Y)

					if distance == false {
						if p.IsJump(f) {
							calc := p.Map.GetDistance(b.X, b.Y, b2.X, b2.Y)

							if calc == 2 || calc == 3 {
								distance = true
							}
						}

						if p.IsMoles(f) {
							calc := p.Map.GetDistance(b.X, b.Y, b2.X, b2.Y)

							if calc == 2 {
								distance = true
							}
						}
					}

					if distance == true {
						flag := false
						for _, v := range connect {
							if b2.X == v.X && b2.Y == v.Y {
								flag = true
								break
							}
						}

						if flag == false {
							connect = append(connect, b2)
						}

						flag = false
						for _, v := range disconnect {
							if b2.X == v.X && b2.Y == v.Y {
								flag = true
								break
							}
						}

						if flag == false {
							disconnect = append(disconnect, b2)
						}
					}
				}
			}

			remain2 := make([]resources.Position, 0)
			for _, v := range remain {
				flag := false
				for _, v2 := range disconnect {
					if v.X == v2.X && v.Y == v2.Y {
						flag = true
						break
					}
				}

				if flag == false {
					remain2 = append(remain2, v)
				}
			}

			remain = remain2

			if len(connect) == 0 {
				count := len(items)
				if count > max {
					max = count
				}

				if len(remain) == 0 {
					break
				}

				items = make([]resources.Position, 0)
				items = append(items, remain[0])
				remain = remain[1:]

				if len(remain) == 0 {
					break
				}
			} else {
				items = append(items, connect...)
			}
		}

		scores = append(scores, Score{User: user, Score: max})
	}

	if p.Count <= 2 {
		scores = append(scores, Score{User: -1, Score: p.DummyNetwork})
	}

	sort.Slice(scores, func(i, j int) bool {
		return scores[i].Score > scores[j].Score
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
				if v.User == -1 {
					continue
				}
				faction := p.Factions[v.User]
				f := faction.GetInstance()
				if calculate == true {
					f.Receive.VP += receive
				} else {
					f.ReceiveResource(resources.Price{VP: receive})
				}
			}
		}

		scores = scores[cnt:]
		receiveCount += cnt

		if receiveCount >= 3 {
			break
		}

		if receiveCount > p.Count {
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
			f := v.GetInstance()

			if f.Science[k] == 0 {
				continue
			}

			scores = append(scores, Score{User: user, Score: f.Science[k]})
		}

		if p.Count <= 2 {
			scores = append(scores, Score{User: -1, Score: p.Sciences.Value[k][color.None]})
		}

		sort.Slice(scores, func(i, j int) bool {
			return scores[i].Score > scores[j].Score
		})

		receives := []int{0, 8, 6, 4, 3, 2}
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
					if v.User == -1 {
						continue
					}
					faction := p.Factions[v.User]
					f := faction.GetInstance()
					if calculate == true {
						f.Receive.VP += receive
					} else {
						f.ReceiveResource(resources.Price{VP: receive})
					}
				}
			}

			scores = scores[cnt:]
			receiveCount += cnt

			if receiveCount >= 3 {
				break
			}

			if receiveCount > p.Count {
				break
			}

			if receiveCount == 1 {
				receives = []int{0, 4, 3, 2, 1}
			} else if receiveCount == 2 {
				receives = []int{0, 2, 1, 0}
			}
		}
	}

	// 남은 자원
	for _, v := range p.Factions {
		f := v.GetInstance()

		receive := (f.Resource.Coin + f.Resource.Worker + f.Resource.Prist + f.Resource.Book.Count() + f.Resource.Power[2] + f.Resource.Power[1]/2) / 5
		if calculate == true {
			f.Receive.VP += receive
		} else {
			f.ReceiveResource(resources.Price{VP: receive})
		}
	}
}

func (p *Game) EndGame() {
	p.LastRoundVP(false)

	ranks := make([]Score, 0)
	for i, v := range p.Factions {
		f := v.GetInstance()

		ranks = append(ranks, Score{User: i, Score: f.VP})
	}

	sort.Slice(ranks, func(i, j int) bool {
		return ranks[i].Score > ranks[j].Score
	})

	rank := 1
	ranks[0].Rank = rank
	if len(ranks) > 1 {
		for i := 1; i < len(ranks); i++ {
			if ranks[i].Score != ranks[i-1].Score {
				rank = i + 1
			}

			ranks[i].Rank = rank
		}
	}

	if p.Replay == true {
		return
	}

	db := models.NewConnection()
	defer db.Close()

	conn, _ := db.Begin()
	defer conn.Rollback()

	gameManager := models.NewGameManager(conn)
	gameuserManager := models.NewGameuserManager(conn)

	g := gameManager.Get(p.Id)

	if g.Status != gm.StatusEnd {
		g.Status = gm.StatusEnd
		g.Enddate = global.GetCurrentDatetime()
		gameManager.Update(g)

		gameusers := gameuserManager.Find([]interface{}{
			models.Where{Column: "game", Value: g.Id, Compare: "="},
		})

		for i, v := range p.Factions {
			f := v.GetInstance()

			rank := 0
			for _, v3 := range ranks {
				if i == v3.User {
					rank = v3.Rank
					break
				}
			}

			for _, v2 := range gameusers {
				if p.Users[i] == v2.User {
					v2.Score = f.VP
					v2.Rank = rank
					gameuserManager.Update(&v2)
				}
			}

		}

		conn.Commit()

		p.CalculateElo()
	}

	p.MakeReplay()
}

func (p *Game) CalculateEndGame() {
	p.LastRoundVP(true)
}

func (p *Game) Elo(a float64, b float64, rank int) (float64, float64) {
	elo_k := 16.0
	//first := 1000

	var1 := 1.0 / (1.0 + math.Pow(10, (b-a)/400.0))
	var2 := 1.0 / (1.0 + math.Pow(10, (a-b)/400.0))

	s1 := 1.0
	s2 := 0.0
	if rank == 1 {
		s1 = 1.0
		s2 = 0.0
	} else if rank == 2 {
		s1 = 0.0
		s2 = 1.0
	} else {
		s1 = 0.5
		s2 = 0.5
	}

	ret1 := elo_k * (s1 - var1)
	ret2 := elo_k * (s2 - var2)

	return ret1, ret2
}

func (p *Game) CalculateElo() {
	if p.Replay == true {
		return
	}

	db := models.NewConnection()
	defer db.Close()

	conn, _ := db.Begin()
	defer conn.Rollback()

	gameuserManager := models.NewGameuserManager(conn)
	userManager := models.NewUserManager(conn)

	items := gameuserManager.Find([]interface{}{
		models.Where{Column: "game", Value: p.Id, Compare: "="},
	})

	if len(items) <= 1 {
		return
	}

	users := make([][]int64, 0)
	elos := make(map[int64]float64)

	for _, v := range items {
		elos[v.User] = 0.0
	}

	for _, v := range items {
		for _, v2 := range items[1:] {
			if v.User == v2.User {
				continue
			}

			flag := false
			for _, k := range users {
				if (k[0] == v.User && k[1] == v2.User) || (k[0] == v2.User && k[1] == v.User) {
					flag = true
					break
				}
			}

			if flag == true {
				continue
			}

			users = append(users, []int64{v.User, v2.User})

			rank := 0
			if v.Score > v2.Score {
				rank = 1
			} else if v.Score < v2.Score {
				rank = 2
			}

			user1 := v.Extra["user"].(models.User)
			user2 := v2.Extra["user"].(models.User)
			elo1, elo2 := p.Elo(float64(user1.Elo), float64(user2.Elo), rank)

			userManager.IncreaseElo(models.Double(elo1), v.User)
			userManager.IncreaseElo(models.Double(elo2), v2.User)

			elos[v.User] += elo1
			elos[v2.User] += elo2
		}

		userManager.IncreaseCount(1, v.User)
	}

	for _, v := range items {
		v.Elo = models.Double(elos[v.User])
		gameuserManager.Update(&v)
	}

	conn.Commit()
}

func (p *Game) SelectFactionTile(user int, name string) error {
	pos := 0

	for i, v := range p.FactionTiles.Items {
		if strings.ToLower(v.Name) == name {
			pos = i
		}
	}

	tile := p.FactionTiles.Items[pos]

	faction := p.Factions[user]
	f := faction.GetInstance()

	if !f.AddTile(tile) {
		return errors.New("already been selected")
	}

	p.FactionTiles.Items[pos].Use = true
	f.Action = true

	if p.Replay == false {
		conn := models.NewConnection()
		defer conn.Close()

		gameuserManager := models.NewGameuserManager(conn)
		gameuser := gameuserManager.GetByGameUser(p.Id, p.Users[user])
		gameuser.Faction = int(tile.Type)
		gameuserManager.Update(gameuser)
	}

	return nil
}

func (p *Game) SelectColorTile(user int, name string) error {
	pos := 0

	for i, v := range p.ColorTiles.Items {
		if strings.ToLower(v.Name) == name {
			pos = i
		}
	}

	tile := p.ColorTiles.Items[pos]

	faction := p.Factions[user]
	f := faction.GetInstance()

	if !f.AddTile(tile) {
		return errors.New("already been selected")
	}

	p.ColorTiles.Items[pos].Use = true
	f.Action = true

	if p.Replay == false {
		conn := models.NewConnection()
		defer conn.Close()

		gameuserManager := models.NewGameuserManager(conn)
		gameuser := gameuserManager.GetByGameUser(p.Id, p.Users[user])
		gameuser.Color = int(tile.Color)
		gameuserManager.Update(gameuser)
	}

	return nil
}

func (p *Game) SelectPalaceTile(user int, pos int) error {
	tile := p.PalaceTiles.Items[pos]

	faction := p.Factions[user]
	f := faction.GetInstance()

	if !f.AddTile(tile) {
		return errors.New("already been selected")
	}

	p.PalaceTiles.Items[pos].Use = true
	f.Action = true

	return nil
}

func (p *Game) InitDraft() {
	p.PalaceTiles.Items = make([]resources.TileItem, 0)

	for user, v := range p.Factions {
		f := v.GetInstance()

		var factionTile resources.TileItem
		var colorTile resources.TileItem
		var palaceTile resources.TileItem

		for _, tile := range f.Tiles {
			if tile.Category == resources.TileFaction {
				factionTile = tile
			}
			if tile.Category == resources.TileColor {
				colorTile = tile
			}

			if tile.Category == resources.TilePalace {
				palaceTile = tile
			}
		}

		factionTile.Use = false
		colorTile.Use = false
		palaceTile.Use = false

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

		item.Init(colorTile, p.Usernames[user])

		p.Factions[user] = item
		f = item.GetInstance()

		f.Tiles = make([]resources.TileItem, 0)
		f.Tiles = append(f.Tiles, factionTile)
		f.Tiles = append(f.Tiles, colorTile)

		if factionTile.Type == resources.TileFactionMonks {
			f.FirstBuilding = resources.SA
		}

		palaceTile.Color = colorTile.Color

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

		p.Sciences.AddUser(f.Color, f.Science)

		p.PalaceTiles.Items = append(p.PalaceTiles.Items, palaceTile)
	}
}

func (p *Game) MakeUndo(id int64, history int64, userid int64) {
	user := p.GetUserPos(userid)

	p.UndoRequest = UndoRequest{Id: id, User: user, History: history, Status: 1, Command: "", Users: []int{user}}
}

func (p *Game) AddUndoConfirm(userid int64) {
	user := p.GetUserPos(userid)

	p.UndoRequest.Users = append(p.UndoRequest.Users, user)
}

func (p *Game) Lock() {
	Lock(p.Id)
}

func (p *Game) Unlock() {
	Unlock(p.Id)
}

func (p *Game) MakeReplay() {
	replay := make([]int, 8)

	count := 1
	round := 1
	for _, v := range p.Replays {
		if v.Round == round {
			replay[v.Round] = count
			round++
		}

		if v.Command[2:] == "save" {
			count++
		}
	}

	replay[0] = count

	p.ReplayPos = replay
}
