package game

var rooms map[int64]*Game

func init() {
	rooms = make(map[int64]*Game, 0)
}

func Make(id int64) {
	g := NewGame()

	rooms[id] = g

	//g.AddFaction(&factions.Monks{}, resources.GetColorTile(color.Yellow))
	//g.AddFaction(&factions.Lizards{})

	//g.InitGame(3)

	//g.BuildStart()
	//g.FirstBuild(0, 3, 4)
	//g.FirstBuild(0, 4, 11)
	//g.GetRoundTile(0, 0)
	/*
		log.Println(g.FirstBuild(0, 3, 4))
		g.TurnEnd(0)

		log.Println(g.GetRoundTile(0, 0))
		g.TurnEnd(0)
	*/
	/*
		log.Println(g.AdvanceShip(0))
		g.TurnEnd(0)
	*/
}

func Get(id int64) *Game {
	return rooms[id]
}
