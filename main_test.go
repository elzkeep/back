package main_test

import (
	"aoi/game"
	"aoi/game/action"
	"aoi/game/factions"
	"aoi/game/resources"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
