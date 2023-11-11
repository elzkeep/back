package game

import (
	"aoi/game/action"
	"aoi/game/resources"
	"aoi/models"
	"aoi/models/game"
	"errors"
	"log"
	"math/rand"
	"sync"
)

var _rooms map[int64]*Game
var _mutex sync.Mutex

func init() {
	_mutex = sync.Mutex{}
	_rooms = make(map[int64]*Game, 0)
}

func Init() {
	log.Println("room Init")
	conn := models.NewConnection()
	defer conn.Close()

	gameManager := models.NewGameManager(conn)
	gameuserManager := models.NewGameuserManager(conn)
	gamehistoryManager := models.NewGamehistoryManager(conn)

	games := gameManager.Find(nil)

	for _, v := range games {
		g := NewGame(v.Id, v.Count)
		_rooms[v.Id] = g

		gameusers := gameuserManager.Find([]interface{}{
			models.Where{Column: "game", Value: v.Id, Compare: "="},
			models.Ordering("gu_order"),
		})

		if v.Count == len(gameusers) {
			for _, gameuser := range gameusers {
				g.AddUser(gameuser.User)
			}

			g.CompleteAddUser()

			historys := gamehistoryManager.Find([]interface{}{
				models.Where{Column: "game", Value: v.Id, Compare: "="},
				models.Ordering("gh_id"),
			})

			for _, history := range historys {
				Command(g, history.Game, history.User, history.Command, false)
				//Command(g, history.Game, history.User, fmt.Sprintf("%v save", history.Command[:1]), false)
			}

			log.Println(g.Command)
		}
	}
}

func Make(user int64, item *models.Game) {
	db := models.NewConnection()
	defer db.Close()

	conn, _ := db.Begin()
	defer conn.Rollback()

	gameManager := models.NewGameManager(conn)
	gametileManager := models.NewGametileManager(conn)

	item.Status = game.StatusReady
	gameManager.Insert(item)

	id := gameManager.GetIdentity()

	{
		items := []int{
			int(resources.TileFactionBlessed),
			int(resources.TileFactionFelines),
			int(resources.TileFactionGoblins),
			int(resources.TileFactionIllusionists),
			int(resources.TileFactionInventors),
			int(resources.TileFactionLizards),
			int(resources.TileFactionMoles),
			int(resources.TileFactionMonks),
			int(resources.TileFactionNavigators),
			int(resources.TileFactionOmar),
			int(resources.TileFactionPhilosophers),
			int(resources.TileFactionPsychics),
		}

		rand.Shuffle(len(items), func(i, j int) { items[i], items[j] = items[j], items[i] })
		for i, v := range items[:7] {
			var tile models.Gametile

			tile.Type = int(resources.TileFaction)
			tile.Number = v
			tile.Order = i + 1
			tile.Game = id

			gametileManager.Insert(&tile)
		}
	}

	{
		items := []int{
			int(resources.TileColorRed),
			int(resources.TileColorYellow),
			int(resources.TileColorBrown),
			int(resources.TileColorBlack),
			int(resources.TileColorBlue),
			int(resources.TileColorGreen),
			int(resources.TileColorGray),
		}

		rand.Shuffle(len(items), func(i, j int) { items[i], items[j] = items[j], items[i] })
		for i, v := range items {
			var tile models.Gametile

			tile.Type = int(resources.TileColor)
			tile.Number = v
			tile.Order = i + 1
			tile.Game = id

			gametileManager.Insert(&tile)
		}
	}

	{
		items := []int{
			int(resources.TileRoundEdgeVP),
			int(resources.TileRoundPristVP),
			int(resources.TileRoundTpVP),
			int(resources.TileRoundShVP),
			int(resources.TileRoundSpade),
			int(resources.TileRoundBridge),
			int(resources.TileRoundScienceCube),
			int(resources.TileRoundSchoolScienceCoin),
			int(resources.TileRoundPower),
			int(resources.TileRoundCoin),
		}

		rand.Shuffle(len(items), func(i, j int) { items[i], items[j] = items[j], items[i] })
		for i, v := range items[:7] {
			var tile models.Gametile

			tile.Type = int(resources.TileRound)
			tile.Number = v
			tile.Order = i + 1
			tile.Game = id

			gametileManager.Insert(&tile)
		}
	}

	{
		items := []int{
			int(resources.TilePalaceWorker),
			int(resources.TilePalaceSpade),
			int(resources.TilePalaceDowngrade),
			int(resources.TilePalaceTpUpgrade),
			int(resources.TilePalaceSchoolTile),
			int(resources.TilePalaceScience),
			int(resources.TilePalaceSchoolVp),
			int(resources.TilePalace6PowerCity),
			int(resources.TilePalaceJump),
			int(resources.TilePalacePower),
			int(resources.TilePalaceCity),
			int(resources.TilePalaceDVp),
			int(resources.TilePalaceTpVp),
			int(resources.TilePalaceRiverCity),
			int(resources.TilePalaceBridge),
			int(resources.TilePalaceTpBuild),
		}

		rand.Shuffle(len(items), func(i, j int) { items[i], items[j] = items[j], items[i] })
		for i, v := range items[:item.Count+1] {
			var tile models.Gametile

			tile.Type = int(resources.TilePalace)
			tile.Number = v
			tile.Order = i + 1
			tile.Game = id

			gametileManager.Insert(&tile)
		}

		var tile models.Gametile

		tile.Type = int(resources.TilePalace)
		tile.Number = int(resources.TilePalaceVp)
		tile.Order = item.Count + 1 + 1
		tile.Game = id

		gametileManager.Insert(&tile)
	}

	{
		items := []int{
			int(resources.TileSchoolWorker),
			int(resources.TileSchoolSpade),
			int(resources.TileSchoolPrist),
			int(resources.TileSchoolEdgeVP),
			int(resources.TileSchoolCoin),
			int(resources.TileSchoolAnnex),
			int(resources.TileSchoolNeutral),
			int(resources.TileSchoolBook),
			int(resources.TileSchoolVP),
			int(resources.TileSchoolPower),
			int(resources.TileSchoolPassCity),
			int(resources.TileSchoolPassPrist),
		}

		rand.Shuffle(len(items), func(i, j int) { items[i], items[j] = items[j], items[i] })
		for i, v := range items {
			var tile models.Gametile

			tile.Type = int(resources.TileSchool)
			tile.Number = v
			tile.Order = i + 1
			tile.Game = id

			gametileManager.Insert(&tile)
		}
	}

	{
		items := []int{
			int(action.Power5),
			int(action.Science),
			int(action.Coin6),
			int(action.TpUpgrade),
			int(action.TpVP),
			int(action.Spade3),
		}

		rand.Shuffle(len(items), func(i, j int) { items[i], items[j] = items[j], items[i] })

		for i, v := range items[:3] {
			var tile models.Gametile

			tile.Type = int(resources.TileBookAction)
			tile.Number = v
			tile.Order = i + 1
			tile.Game = id

			gametileManager.Insert(&tile)
		}
	}

	{
		/*
			items := []int{
				int(DVP),
				int(DVP),
				int(TpVP),
				int(TpVP),
				int(TeVP),
				int(ShSaVP),
				int(ShSaVP),
				int(SpadeVP),
				int(ScienceVP),
				int(CityVP),
				int(AdvanceVP),
				int(InnovationVP),
			}
		*/

		items := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

		rand.Shuffle(len(items), func(i, j int) { items[i], items[j] = items[j], items[i] })

		for i, v := range items[:6] {
			var tile models.Gametile

			tile.Type = int(resources.TileRoundBonus)
			tile.Number = v
			tile.Order = i + 1
			tile.Game = id

			gametileManager.Insert(&tile)
		}

		/*
			finalRound := []int{
				int(DVP),
				int(TpVP),
				int(TeVP),
				int(EdgeVP),
			}
		*/

		finalRound := []int{0, 1, 2, 3}

		var tile models.Gametile

		tile.Type = int(resources.TileRoundBonus)
		tile.Number = finalRound[rand.Intn(len(finalRound))]
		tile.Order = 6 + 1
		tile.Game = id

		gametileManager.Insert(&tile)
	}

	{
		items := []int{
			int(resources.TileInnovationKind),
			int(resources.TileInnovationCount),
			int(resources.TileInnovationSchool),
			int(resources.TileInnovationCity),
			int(resources.TileInnovationScience),
			int(resources.TileInnovationCluster),
			int(resources.TileInnovationD),
			int(resources.TileInnovationUpgrade),
			int(resources.TileInnovationBridge),
			int(resources.TileInnovationFreeD),
			int(resources.TileInnovationFreeTP),
			int(resources.TileInnovationFreeSchool),
			int(resources.TileInnovationFreeSA),
			int(resources.TileInnovationFreeSH),
			int(resources.TileInnovationFreeMT),
		}

		rand.Shuffle(len(items), func(i, j int) { items[i], items[j] = items[j], items[i] })

		counts := []int{4, 6, 8, 10, 12}

		for i, v := range items[:counts[item.Count-1]] {
			var tile models.Gametile

			tile.Type = int(resources.TileInnovation)
			tile.Number = v
			tile.Order = i + 1
			tile.Game = id

			gametileManager.Insert(&tile)
		}
	}

	conn.Commit()

	g := NewGame(id, item.Count)
	_rooms[id] = g

	Join(user, id)
}

func Lock() {
	_mutex.Lock()
}

func Unlock() {
	_mutex.Unlock()
}

func Join(user int64, id int64) error {
	g := Get(id)

	if g == nil {
		return errors.New("not found game")
	}

	conn := models.NewConnection()
	defer conn.Close()

	gameManager := models.NewGameManager(conn)
	gameuserManager := models.NewGameuserManager(conn)

	Lock()

	item := gameManager.Get(id)

	if item.Status != game.StatusReady {
		Unlock()
		return errors.New("status error")
	}

	if gameuserManager.CountByGameUser(id, user) > 0 {
		Unlock()
		return errors.New("already")
	}

	count := gameuserManager.CountByGame(id)
	if count >= item.Count {
		Unlock()
		return errors.New("full")
	}

	var gameuser models.Gameuser
	gameuser.User = user
	gameuser.Game = id

	gameuserManager.Insert(&gameuser)
	count++

	items := gameuserManager.FindByGame(id)

	if len(items) == item.Count {
		log.Println("join complete")
		gameManager.UpdateStatus(int(game.StatusFaction), id)

		rand.Shuffle(len(items), func(i, j int) { items[i], items[j] = items[j], items[i] })

		for i, v := range items {
			v.Order = i + 1
			gameuserManager.Update(&v)

			g.AddUser(v.User)
		}

		g.CompleteAddUser()
	} else {
		log.Println("join not complete")
	}

	Unlock()

	return nil
}

func Get(id int64) *Game {
	return _rooms[id]
}
