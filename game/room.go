package game

import (
	"aoi/game/factions"
	"log"
)

var rooms map[int64]*Game

func init() {
	rooms = make(map[int64]*Game, 0)
}

func Make(id int64) {
	g := NewGame()

	rooms[id] = g

	g.AddFaction(&factions.Monks{})
	//g.AddFaction(&factions.Lizards{})

	g.BuildStart()

	log.Println(g.FirstBuild(0, 3, 4))
	g.TurnEnd(0)
	log.Println(g.GetRoundTile(0, 0))
	g.TurnEnd(0)

	g.Start()

	log.Println(g.AdvanceShip(0))
	g.TurnEnd(0)
}

func Get(id int64) *Game {
	return rooms[id]
}
