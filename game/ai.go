package game

import (
	"aoi/game/resources"
	"fmt"
	"math/rand"
	"strings"
)

func GetPositionName(x int, y int) string {
	letters := []string{
		"A",
		"B",
		"C",
		"D",
		"E",
		"F",
		"G",
		"H",
		"I",
		"J",
		"K",
		"L",
		"M",
		"N",
	}

	return fmt.Sprintf("%v%v", letters[x], y+1)
}

func AICommand(p *Game, user int) {
	if p.Id != 32 {
		return
	}
	if p.Round == FactionRound {
		AIFaction(p, user)
	} else if p.Round == BuildRound {
		AIBuild(p, user)
	} else if p.Round == TileRound {

	} else if p.Round == SpadeRound {
		AISpade(p, user)
	} else if p.IsScienceTurn() || p.IsBookTurn() || p.IsResourceTurn() {
		AIResource(p, user)
	} else if p.IsPowerTurn() {
		command := fmt.Sprintf("%v leech %v power", user, p.Turn[0].Power)
		AIRunCommand(p, user, command)
	} else {
		AIAction(p, user)
	}
}

func AIRunCommand(p *Game, user int, str string) {
	Command(p, p.Id, 1, str, true)
	command := fmt.Sprintf("%v save", user)
	Command(p, p.Id, 1, command, true)
}

func AIRunOnlyCommand(p *Game, user int, str string) {
	Command(p, p.Id, 1, str, true)
}

func AIRunSave(p *Game, user int) {
	command := fmt.Sprintf("%v save", user)
	Command(p, p.Id, 1, command, true)
}

func AIFaction(p *Game, user int) {
	factions := []resources.TileType{
		resources.TileFactionIllusionists,
		resources.TileFactionGoblins,
		resources.TileFactionBlessed,
		resources.TileFactionNavigators,
	}

	find := false
	for _, v := range factions {
		for _, tile := range p.FactionTiles.Items {
			if tile.Use == true {
				continue
			}

			if tile.Type == v {
				command := fmt.Sprintf("%v faction %v", user, strings.ToLower(tile.Name))
				AIRunCommand(p, user, command)
				find = true
				break
			}
		}

		if find == true {
			break
		}
	}
}

func AIBuild(p *Game, user int) {
	faction := p.Factions[user]
	f := faction.GetInstance()

	positions := [][]string{
		{},
		{},
		{"H4", "E9"},
		{"H5", "E8"},
		{"E6", "B7"},
		{"E5", "H8"},
		{"E7", "H7"},
		{"D5", "F8"},
		{"D8", "F6"},
	}

	pos := positions[int(f.Color)][f.Building[resources.D]]

	command := fmt.Sprintf("%v build %v D", user, pos)
	AIRunCommand(p, user, command)
}

func AISpade(p *Game, user int) {
	faction := p.Factions[user]
	f := faction.GetInstance()

	for i := 1; i <= 3; i++ {
		for _, v := range f.BuildingList {
			positions := resources.GetGroundPosition(v.X, v.Y)

			for _, position := range positions {
				need := p.Map.GetNeedSpade(position.X, position.Y, f.Color)
				if need == i {
					command := fmt.Sprintf("%v dig %v 1", user, GetPositionName(position.X, position.Y))
					AIRunCommand(p, user, command)
					return
				}
			}
		}
	}
}

func AIResource(p *Game, user int) {
	faction := p.Factions[user]
	f := faction.GetInstance()

	roundBonus := p.RoundBonuss.Get(p.Round)

	science := ""
	if roundBonus.Science.Banking > 0 {
		science = "banking"
	} else if roundBonus.Science.Law > 0 {
		science = "law"
	} else if roundBonus.Science.Engineering > 0 {
		science = "engineering"
	} else if roundBonus.Science.Medicine > 0 {
		science = "medicine"
	}

	for i := 0; i < f.Resource.Science.Any; i++ {
		command := fmt.Sprintf("%v science %v 1", user, science)
		AIRunOnlyCommand(p, user, command)
	}

	book := []string{"banking", "law", "engineering", "medicine"}
	for i := 0; i < f.Resource.Book.Any; i++ {
		command := fmt.Sprintf("%v book %v 1", user, book[rand.Intn(4)])
		AIRunOnlyCommand(p, user, command)
	}

	AIRunSave(p, user)
}

func AIAction(p *Game, user int) {
	faction := p.Factions[user]
	f := faction.GetInstance()
	other := p.Factions[1-user].GetInstance()

	if p.Round == 1 {
		if p.PowerActions.Items[3].Use == false {
			if f.Resource.Power[2]+f.Resource.Power[1]/2 >= 4 {
				command := fmt.Sprintf("%v action ACT4", user)
				AIRunCommand(p, user, command)
				return
			}
		}

		if f.Building[resources.TP] == 0 && f.Building[resources.TE] == 0 {
			for _, v := range f.BuildingList {
				if v.Building != resources.D {
					continue
				}

				positions := resources.GetGroundPosition(v.X, v.Y)

				for _, position := range positions {
					if p.Map.GetOwner(position.X, position.Y) == other.Color {
						command := fmt.Sprintf("%v upgrade %v TP", user, GetPositionName(v.X, v.Y))
						AIRunCommand(p, user, command)
						return
					}
				}
			}
			for _, v := range f.BuildingList {
				if v.Building != resources.D {
					continue
				}

				command := fmt.Sprintf("%v upgrade %v TP", user, GetPositionName(v.X, v.Y))
				AIRunCommand(p, user, command)
				return
			}
		} else if f.Building[resources.TP] == 1 && f.Building[resources.TE] == 0 {
			book := []string{"banking", "law", "engineering", "medicine"}

			for _, v := range f.BuildingList {
				if v.Building != resources.TP {
					continue
				}

				for i, s := range p.SchoolTiles.Items {
					for i2, s2 := range s {
						if s2.Tile.Type == resources.TileSchoolPower || s2.Tile.Type == resources.TileSchoolWorker {
							command := fmt.Sprintf("%v upgrade %v TE", user, GetPositionName(v.X, v.Y))
							AIRunOnlyCommand(p, user, command)
							command = fmt.Sprintf("%v schooltile %v %v", user, book[i], i2+1)
							AIRunOnlyCommand(p, user, command)
							AIRunSave(p, user)
							return
						}
					}
				}
			}
		} else {
			for _, v := range f.Tiles {
				if v.Use == true {
					continue
				}

				if v.Type == resources.TileSchoolPower {
					command := fmt.Sprintf("%v action SCHOOL10", user)
					AIRunCommand(p, user, command)
					return
				} else if v.Type == resources.TileRoundSpade {
					command := fmt.Sprintf("%v action ROUND5", user)
					AIRunOnlyCommand(p, user, command)

					for _, v := range f.BuildingList {
						positions := resources.GetGroundPosition(v.X, v.Y)

						for _, position := range positions {
							need := p.Map.GetNeedSpade(position.X, position.Y, f.Color)
							if need == 1 {
								command := fmt.Sprintf("%v build %v D", user, GetPositionName(position.X, position.Y))
								AIRunCommand(p, user, command)
								return
							}
						}
					}

					for _, v := range f.BuildingList {
						positions := resources.GetGroundPosition(v.X, v.Y)

						for _, position := range positions {
							need := p.Map.GetNeedSpade(position.X, position.Y, f.Color)
							if need == 2 {
								command := fmt.Sprintf("%v dig %v 1", user, GetPositionName(position.X, position.Y))
								AIRunCommand(p, user, command)
								return
							}
						}
					}

					AIRunSave(p, user)

					return
				} else if v.Type == resources.TileRoundScienceCube {
					command := fmt.Sprintf("%v action ROUND7", user)
					AIRunOnlyCommand(p, user, command)

					roundBonus := p.RoundBonuss.Get(p.Round)

					science := ""
					if roundBonus.Science.Banking > 0 {
						science = "banking"
					} else if roundBonus.Science.Law > 0 {
						science = "law"
					} else if roundBonus.Science.Engineering > 0 {
						science = "engineering"
					} else if roundBonus.Science.Medicine > 0 {
						science = "medicine"
					}

					for i := 0; i < f.Resource.Science.Any; i++ {
						command := fmt.Sprintf("%v science %v 1", user, science)
						AIRunOnlyCommand(p, user, command)
					}

					AIRunSave(p, user)

					return
				} else if v.Type == resources.TileRoundSpade {
					command := fmt.Sprintf("%v action SCHOOL10", user)
					AIRunCommand(p, user, command)
					return
				}
			}
		}
	}
}
