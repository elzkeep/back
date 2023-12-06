package game

import (
	"aoi/game/action"
	"aoi/game/resources"
	"aoi/global"
	"aoi/models"
	"aoi/models/game"
	"aoi/models/gameundo"
	"errors"
	"log"
	"math/rand"
	"sync"
)

var _rooms map[int64]*Game
var _mutex map[int64]*sync.Mutex

func init() {
	_mutex = make(map[int64]*sync.Mutex)
	_rooms = make(map[int64]*Game)
}

func SetGame(id int64, game *Game) {
	_rooms[id] = game
}

func MakeGame(id int64) {
	conn := models.NewConnection()
	defer conn.Close()

	gameManager := models.NewGameManager(conn)
	gameuserManager := models.NewGameuserManager(conn)
	gamehistoryManager := models.NewGamehistoryManager(conn)
	gameundoManager := models.NewGameundoManager(conn)
	gameundoitemManager := models.NewGameundoitemManager(conn)

	gameItem := gameManager.Get(id)

	g := NewGame(gameItem)
	SetGame(id, g)

	gameusers := gameuserManager.Find([]interface{}{
		models.Where{Column: "game", Value: id, Compare: "="},
		models.Ordering("gu_order"),
	})

	if gameItem.Count == len(gameusers) {
		for _, gameuser := range gameusers {
			user := gameuser.Extra["user"].(models.User)
			g.AddUser(gameuser.User, user.Name)
		}

		g.CompleteAddUser()

		historys := gamehistoryManager.Find([]interface{}{
			models.Where{Column: "game", Value: id, Compare: "="},
			models.Ordering("gh_id"),
		})

		for _, history := range historys {
			err := Command(g, history.Game, history.User, history.Command, false, history.Id)
			if err != nil {
				log.Println(err)
			}
		}

		if g.Round > 0 && g.Round <= 6 {
			for user := range g.Factions {
				g.Calculate(user)
			}
		}

		if gameItem.Status != game.StatusEnd {
			gameundos := gameundoManager.Find([]interface{}{
				models.Where{Column: "game", Value: id, Compare: "="},
				models.Where{Column: "status", Value: gameundo.StatusWait, Compare: "="},
			})

			for _, gameundo := range gameundos {
				user := g.GetUserPos(gameundo.User)
				g.UndoRequest = UndoRequest{Id: gameundo.Id, User: user, History: gameundo.Gamehistory, Status: 1, Command: "", Users: make([]int, 0)}

				gameundoitems := gameundoitemManager.Find([]interface{}{
					models.Where{Column: "gameundo", Value: gameundo.Id, Compare: "="},
				})

				for _, gameundoitem := range gameundoitems {
					g.AddUndoConfirm(gameundoitem.User)
				}
			}
		}
	}
}
func Init() {

	conn := models.NewConnection()
	defer conn.Close()

	gameManager := models.NewGameManager(conn)
	gameuserManager := models.NewGameuserManager(conn)
	gamehistoryManager := models.NewGamehistoryManager(conn)
	gameundoManager := models.NewGameundoManager(conn)
	gameundoitemManager := models.NewGameundoitemManager(conn)

	games := gameManager.Find([]interface{}{models.Where{Column: "status", Value: game.StatusEnd, Compare: "<>"}})

	for _, v := range games {
		g := NewGame(&v)
		SetGame(v.Id, g)

		gameusers := gameuserManager.Find([]interface{}{
			models.Where{Column: "game", Value: v.Id, Compare: "="},
			models.Ordering("gu_order"),
		})

		if v.Count == len(gameusers) {
			for _, gameuser := range gameusers {
				user := gameuser.Extra["user"].(models.User)
				g.AddUser(gameuser.User, user.Name)
			}

			g.CompleteAddUser()

			historys := gamehistoryManager.Find([]interface{}{
				models.Where{Column: "game", Value: v.Id, Compare: "="},
				models.Ordering("gh_id"),
			})

			for _, history := range historys {
				err := Command(g, history.Game, history.User, history.Command, false, history.Id)
				if err != nil {
					log.Println(err)
				}
			}

			if g.Round > 0 && g.Round <= 6 {
				for user, _ := range g.Factions {
					g.Calculate(user)
				}
			}

			if v.Status != game.StatusEnd {
				gameundos := gameundoManager.Find([]interface{}{
					models.Where{Column: "game", Value: v.Id, Compare: "="},
					models.Where{Column: "status", Value: gameundo.StatusWait, Compare: "="},
				})

				for _, gameundo := range gameundos {
					user := g.GetUserPos(gameundo.User)
					g.UndoRequest = UndoRequest{Id: gameundo.Id, User: user, History: gameundo.Gamehistory, Status: 1, Command: "", Users: make([]int, 0)}

					gameundoitems := gameundoitemManager.Find([]interface{}{
						models.Where{Column: "gameundo", Value: gameundo.Id, Compare: "="},
					})

					for _, gameundoitem := range gameundoitems {
						g.AddUndoConfirm(gameundoitem.User)
					}
				}
			}

			//for i, v := range g.Users {
			//	if v == 1 {
			//		if g.IsTurn(i) {
			//			AICommand(g, i)
			//		}
			//	}
			//}

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
	item.Join = 0
	gameManager.Insert(item)

	id := gameManager.GetIdentity()
	item.Id = id

	count := item.Count

	if count == 1 {
		count = 2
	}

	factionCount := 7
	colorCount := 7
	roundCount := 10

	if GameType(item.Type) != BasicType {
		factionCount = count + 1
		colorCount = 6
		roundCount = count + 3
	}

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

		for {
			rand.Shuffle(len(items), func(i, j int) { items[i], items[j] = items[j], items[i] })

			if item.Illusionists == game.IllusionistsBan {
				find := false

				for _, v := range items[:factionCount] {
					if v == int(resources.TileFactionIllusionists) {
						find = true
						break
					}
				}

				if find == true {
					continue
				}
			}

			for i, v := range items[:factionCount] {
				var tile models.Gametile

				tile.Type = int(resources.TileFaction)
				tile.Number = v
				tile.Order = i + 1
				tile.Game = id

				gametileManager.Insert(&tile)
			}

			break
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
		for i, v := range items[:colorCount] {
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

		for i, v := range items[:roundCount] {

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

		for i, v := range items[:count+1] {
			var tile models.Gametile

			tile.Type = int(resources.TilePalace)
			tile.Number = v
			tile.Order = i + 1
			tile.Game = id

			gametileManager.Insert(&tile)
		}

		if item.Type == int(BasicType) {
			var tile models.Gametile

			tile.Type = int(resources.TilePalace)
			tile.Number = int(resources.TilePalaceVp)
			tile.Order = count + 1 + 1
			tile.Game = id

			gametileManager.Insert(&tile)
		}
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

		sciences := []resources.Science{
			{Law: 3},
			{Banking: 3},
			{Law: 3},
			{Medicine: 4},
			{Banking: 1},
			{Medicine: 2},
			{Banking: 2},
			{Engineering: 1},
			{Medicine: 3},
			{Engineering: 4},
			{Engineering: 3},
			{Law: 2},
		}

		for {
			rand.Shuffle(len(items), func(i, j int) { items[i], items[j] = items[j], items[i] })

			if items[4] == 7 || items[5] == 7 {
				continue
			}

			science := resources.Science{Banking: 0, Law: 0, Engineering: 0, Medicine: 0}
			for _, v := range items[:6] {
				item := sciences[v]
				if item.Banking > 0 {
					science.Banking++
				} else if item.Law > 0 {
					science.Law++
				} else if item.Engineering > 0 {
					science.Engineering++
				} else if item.Medicine > 0 {
					science.Medicine++
				}
			}

			if science.Banking >= 3 ||
				science.Law >= 3 ||
				science.Engineering >= 3 ||
				science.Medicine >= 3 {
				continue
			}

			break
		}

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

		pos := 0
		for {
			pos = rand.Intn(len(finalRound))

			if (items[5] == 0 || items[5] == 1) && (pos == 0 || pos == 3) {
				continue
			}

			if (items[5] == 2 || items[5] == 3) && pos == 1 {
				continue
			}

			if items[5] == 4 && pos == 2 {
				continue
			}

			break
		}

		tile.Type = int(resources.TileRoundBonus)
		tile.Number = finalRound[pos]
		tile.Order = 6 + 1
		tile.Game = id

		gametileManager.Insert(&tile)
	}

	{
		items := []int{
			int(resources.TileInnovationSpade),
			int(resources.TileInnovationTP),
			int(resources.TileInnovationPrist),
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

		for i, v := range items[:counts[count-1]] {
			var tile models.Gametile

			tile.Type = int(resources.TileInnovation)
			tile.Number = v
			tile.Order = i + 1
			tile.Game = id

			gametileManager.Insert(&tile)
		}
	}

	conn.Commit()

	g := NewGame(item)
	SetGame(id, g)

	Join(user, id)
	/*
		if g.Count == 1 {
			Join(1, id)

			for i, v := range g.Users {
				if v == 1 {
					if g.IsTurn(i) {
						AICommand(g, i)
					}
				}
			}
		}
	*/
}

func Lock(id int64) {
	_, exists := _mutex[id]
	if !exists {
		_mutex[id] = &sync.Mutex{}
	}

	_mutex[id].Lock()
}

func Unlock(id int64) {
	_mutex[id].Unlock()
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

	item := gameManager.Get(id)

	if item.Status != game.StatusReady {
		return errors.New("status error")
	}

	if gameuserManager.CountByGameUser(id, user) > 0 {
		return errors.New("already")
	}

	count := gameuserManager.CountByGame(id)
	if count >= item.Count {
		return errors.New("full")
	}

	var gameuser models.Gameuser
	gameuser.User = user
	gameuser.Game = id

	gameuserManager.Insert(&gameuser)
	count++

	item.Join++
	gameManager.Update(item)

	items := gameuserManager.FindByGame(id)

	if len(items) == item.Count {
		gameManager.UpdateStatus(int(game.StatusFaction), id)

		rand.Shuffle(len(items), func(i, j int) { items[i], items[j] = items[j], items[i] })

		for i, v := range items {
			v.Order = i + 1
			gameuserManager.Update(&v)

			user := v.Extra["user"].(models.User)
			g.AddUser(v.User, user.Name)
		}

		g.CompleteAddUser()
	}

	msg := global.Notify{Id: id, Title: "join"}
	global.SendNotify(msg)

	return nil
}

func Get(id int64) *Game {
	return _rooms[id]
}

func Replay(id int64, pos int) *Game {
	old := _rooms[id]

	game := old.Copy()

	for i, v := range old.Users {
		game.AddUser(v, old.Usernames[i])
	}

	game.CompleteAddUser()

	if pos > 0 {
		end := 0
		for _, item := range old.Replays {
			Command(game, id, item.User, item.Command, false, 0)

			if item.Command[2:] == "save" {
				end++

			}

			if pos == end {
				break
			}
		}
	}

	return game
}
