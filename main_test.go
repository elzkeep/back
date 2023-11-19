package main_test

import (
	"log"
	"math"
	"testing"

	"github.com/bmizerany/assert"
)

/*
func TestBuild(t *testing.T) {
	g := game.NewGame()

	g.AddFaction(&factions.Monks{})
	g.AddFaction(&factions.Lizards{})

	g.BuildStart()

	assert.Equal(t, nil, g.FirstBuild(0, 1, 6))

	assert.Equal(t, nil, g.FirstBuild(1, 0, 0))
	assert.Equal(t, nil, g.FirstBuild(1, 4, 1))

	assert.Equal(t, nil, g.FirstBuild(0, 4, 3))

	assert.Equal(t, nil, g.GetRoundTile(1, 1))
	assert.Equal(t, nil, g.GetRoundTile(0, 0))

	g.Factions[0].Print()
	g.Factions[1].Print()

	//assert.NotEqual(t, nil, g.Build(0, 4, 3))
	//assert.NotEqual(t, nil, g.Build(0, 3, 4))

	assert.Equal(t, nil, g.AdvanceShip(0))
	g.TurnEnd(0)
	assert.Equal(t, nil, g.AdvanceShip(1))
	g.TurnEnd(1)
	assert.Equal(t, nil, g.AdvanceSpade(0))
	g.TurnEnd(0)
	assert.Equal(t, nil, g.AdvanceSpade(1))
	g.TurnEnd(1)

	assert.Equal(t, nil, g.Build(0, 3, 4))
	g.TurnEnd(0)
	assert.Equal(t, nil, g.AdvanceSpade(1))
	g.TurnEnd(1)
	assert.Equal(t, nil, g.Build(0, 3, 5))
	g.TurnEnd(0)

	assert.Equal(t, nil, g.AdvanceShip(1))
	g.TurnEnd(1)
	assert.Equal(t, nil, g.Upgrade(0, 3, 5, resources.TP))

	g.TurnEnd(0)
	assert.Equal(t, nil, g.SendScholar(1, game.Banking))
	g.TurnEnd(1)
	assert.Equal(t, nil, g.SendScholar(0, game.Engineering))
	g.TurnEnd(0)
	assert.Equal(t, nil, g.SendScholar(1, game.Banking))
	g.TurnEnd(1)
	assert.Equal(t, nil, g.SendScholar(0, game.Engineering))
	g.TurnEnd(0)
	assert.Equal(t, nil, g.SendScholar(1, game.Engineering))
	g.TurnEnd(1)

	assert.Equal(t, nil, g.PowerAction(0, int(action.Coin)))
	g.TurnEnd(0)
	assert.NotEqual(t, nil, g.PowerAction(1, int(action.Coin)))
	g.TurnEnd(1)

	g.Factions[0].Print()
	g.Factions[1].Print()

	g.Map.Print()
	g.Sciences.Print()
}

func TestExtraAction(t *testing.T) {
	g := game.NewGame()

	g.AddFaction(&factions.Monks{})
	g.AddFaction(&factions.Lizards{})

	g.BuildStart()

	assert.Equal(t, nil, g.FirstBuild(0, 1, 6))

	assert.Equal(t, nil, g.FirstBuild(1, 0, 0))
	assert.Equal(t, nil, g.FirstBuild(1, 4, 1))

	assert.Equal(t, nil, g.FirstBuild(0, 4, 3))

	assert.Equal(t, nil, g.GetRoundTile(1, 1))
	assert.Equal(t, nil, g.GetRoundTile(0, 0))

	assert.Equal(t, nil, g.AdvanceShip(0))
	g.TurnEnd(0)
	assert.Equal(t, nil, g.AdvanceShip(1))
	g.TurnEnd(1)

	assert.Equal(t, nil, g.PowerAction(0, int(action.Spade)))
	assert.Equal(t, nil, g.Build(0, 3, 4))
	//assert.Equal(t, nil, g.Build(0, 3, 5))
	g.TurnEnd(0)
	assert.Equal(t, nil, g.AdvanceSpade(1))
	g.TurnEnd(1)
	assert.Equal(t, nil, g.PowerAction(0, int(action.Bridge)))
	assert.Equal(t, nil, g.Bridge(0, 4, 3, 5, 4))
	g.TurnEnd(1)

	g.Factions[0].Print()
	g.Factions[1].Print()

	g.Map.Print()
	g.Sciences.Print()
}

func TestDistance(t *testing.T) {
	g := game.NewGame()

	g.AddFaction(&factions.Monks{})
	g.AddFaction(&factions.Lizards{})

	g.BuildStart()

	assert.Equal(t, nil, g.FirstBuild(0, 1, 6))

	assert.Equal(t, nil, g.FirstBuild(1, 0, 0))
	assert.Equal(t, nil, g.FirstBuild(1, 4, 1))

	assert.Equal(t, nil, g.FirstBuild(0, 4, 3))

	log.Println("-------------------------------")
	log.Println(g.Map.CheckDistance(0, 3, 2, 2))
}

func TestCity(t *testing.T) {
	g := game.NewGame()

	g.AddFaction(&factions.Monks{})

	g.BuildStart()

	assert.Equal(t, nil, g.FirstBuild(0, 1, 6))
	assert.Equal(t, nil, g.FirstBuild(0, 4, 3))

	assert.Equal(t, nil, g.GetRoundTile(0, 0))

	assert.Equal(t, nil, g.Build(0, 2, 6))
	g.TurnEnd(0)
	assert.Equal(t, nil, g.Build(0, 2, 5))
	g.TurnEnd(0)
	assert.Equal(t, nil, g.PowerAction(0, int(action.Bridge)))
	assert.Equal(t, nil, g.Bridge(0, 2, 6, 3, 7))
	g.TurnEnd(0)
	assert.Equal(t, nil, g.Build(0, 3, 7))
	g.TurnEnd(0)
	assert.Equal(t, nil, g.Build(0, 4, 7))
	g.TurnEnd(0)
	assert.Equal(t, nil, g.Build(0, 5, 7))
	g.TurnEnd(0)
	assert.Equal(t, nil, g.Build(0, 3, 4))
	assert.Equal(t, nil, g.City(0, game.ScienceCity))
	g.TurnEnd(0)

	g.Map.Print()
	g.Sciences.Print()
}

func TestCityUpgrade(t *testing.T) {
	g := game.NewGame()

	g.AddFaction(&factions.Monks{})

	g.BuildStart()

	assert.Equal(t, nil, g.FirstBuild(0, 1, 6))
	assert.Equal(t, nil, g.FirstBuild(0, 4, 3))

	assert.Equal(t, nil, g.GetRoundTile(0, 1))

	assert.Equal(t, nil, g.Build(0, 2, 6))
	g.TurnEnd(0)
	assert.Equal(t, nil, g.Build(0, 2, 5))
	g.TurnEnd(0)
	assert.Equal(t, nil, g.PowerAction(0, int(action.Bridge)))
	assert.Equal(t, nil, g.Bridge(0, 2, 6, 3, 7))
	g.TurnEnd(0)
	assert.Equal(t, nil, g.Build(0, 3, 7))
	g.TurnEnd(0)
	assert.Equal(t, nil, g.Build(0, 4, 7))
	g.TurnEnd(0)
	assert.Equal(t, nil, g.Build(0, 5, 7))
	g.TurnEnd(0)
	assert.Equal(t, nil, g.Upgrade(0, 5, 7, resources.TP))
	assert.Equal(t, nil, g.City(0, game.ScienceCity))
	g.TurnEnd(0)

	g.Map.Print()
	g.Sciences.Print()
}
*/

func TestElo(t *testing.T) {
	elo_k := 16.0
	//first := 1000
	rank := 1

	a := 1480.0
	b := 1000.0

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

	log.Println(ret1, ret2)

	assert.Equal(t, true, ret1 == ret2)
}
