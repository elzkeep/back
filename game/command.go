package game

import (
	"aoi/game/resources"
	"aoi/global"
	"log"
	"strings"
)

func ConvertPosition(str string) (int, int) {
	return int(str[0]) - 'A', global.Atoi(str[1:]) - 1
}

func ConvertBuilding(str string) resources.Building {
	if str == "D" {
		return resources.D
	} else if str == "TP" {
		return resources.TP
	} else if str == "TE" {
		return resources.TE
	} else if str == "SH" {
		return resources.SH
	} else if str == "SA" {
		return resources.SA
	} else if str == "WHITE_D" {
		return resources.WHITE_D
	} else if str == "WHITE_TP" {
		return resources.WHITE_TP
	} else if str == "WHITE_TE" {
		return resources.WHITE_TE
	} else if str == "WHITE_SH" {
		return resources.WHITE_SH
	} else if str == "WHITE_SA" {
		return resources.WHITE_SA
	} else if str == "WHITE_MT" {
		return resources.WHITE_MT
	}

	return resources.None
}

func ConvertScience(str string) ScienceType {
	if str == "banking" {
		return Banking
	} else if str == "law" {
		return Law
	} else if str == "engineering" {
		return Engineering
	} else if str == "medicine" {
		return Medicine
	}

	return Banking
}

/*

	PowerAction(item action.PowerActionItem) error
	Book(item action.BookActionItem) error
	Bridge(x1 int, y1 int, x2 int, y2 int) error
	Pass(tile *RoundTileItem) error
	ReceiveCity(item city.CityItem) error
*/

func Command(p *Game, str string) error {
	log.Println("Command", str)
	strs := strings.Split(str, " ")

	user := global.Atoi(strs[0])
	cmd := strs[1]

	var err error

	if cmd == "build" {
		x, y := ConvertPosition(strs[2])

		if p.Round == BuildRound {
			err = p.FirstBuild(user, x, y)
		} else {
			target := ConvertBuilding(strs[3])

			err = p.Build(user, x, y, target)
		}
	} else if cmd == "dig" {
		x, y := ConvertPosition(strs[2])
		dig := global.Atoi(strs[3])

		err = p.Dig(user, x, y, dig)
	} else if cmd == "upgrade" {
		x, y := ConvertPosition(strs[2])
		target := ConvertBuilding(strs[3])

		err = p.Upgrade(user, x, y, target)
	} else if cmd == "advance" {
		if strs[2] == "ship" {
			err = p.AdvanceShip(user)
		} else if strs[2] == "dig" {
			err = p.AdvanceSpade(user)
		}
	} else if cmd == "send" {
		target := ConvertScience(strs[2])
		err = p.SendScholar(user, target)
	} else if cmd == "supploy" {
		target := ConvertScience(strs[2])
		err = p.SupployScholar(user, target)
	} else if cmd == "science" {
		target := ConvertScience(strs[2])
		level := global.Atoi(strs[3])
		err = p.Science(user, target, level)
	} else if cmd == "action" {
		action := strs[2][:2]

		if action == "AC" {
			pos := global.Atoi(strs[2][3:]) - 1

			if pos >= len(p.PowerActions.Items) {
				err = p.BookAction(user, pos-len(p.PowerActions.Items))
			} else {
				err = p.PowerAction(user, pos)
			}
		} else if action == "PA" {
			pos := global.Atoi(strs[2][6:]) - 1
			err = p.TileAction(user, resources.TilePalace, pos)
		} else if action == "RO" {
			pos := global.Atoi(strs[2][5:]) + int(resources.TilePalaceVp)
			err = p.TileAction(user, resources.TileRound, pos)
		} else if action == "SC" {
			pos := global.Atoi(strs[2][6:]) + int(resources.TileRoundCoin)

			pos += int(resources.TileRoundCoin)
			err = p.TileAction(user, resources.TileSchool, pos)
		} else if action == "IN" {
			pos := global.Atoi(strs[2][10:]) - 1
			err = p.TileAction(user, resources.TileInnovation, pos)
		}
	} else if cmd == "pass" {
		pos := global.Atoi(strs[2])

		err = p.Pass(user, pos)
	} else if cmd == "roundtile" {
		pos := global.Atoi(strs[2])

		err = p.GetRoundTile(user, pos)
	} else if cmd == "transform" {
		x, y := ConvertPosition(strs[2])
		dig := global.Atoi(strs[3])

		err = p.Dig(user, x, y, dig)
	} else if cmd == "spade" {
		dig := global.Atoi(strs[2])

		err = p.ConvertDig(user, dig)
	} else if cmd == "palacetile" {
		pos := global.Atoi(strs[2])

		p.PalaceTile(user, pos)
	} else if cmd == "schooltile" {
		science := ConvertScience(strs[2])
		level := 3 - global.Atoi(strs[3])

		p.SchoolTile(user, int(science), level)
	} else if cmd == "bridge" {
		x, y := ConvertPosition(strs[2])
		x2, y2 := ConvertPosition(strs[3])

		p.Bridge(user, x, y, x2, y2)
	} else if cmd == "city" {
		pos := global.Atoi(strs[2]) - 1

		p.City(user, resources.CityType(pos))
	} else if cmd == "save" {
		p.TurnEnd(user)
	}

	if err != nil {
		log.Println(err)
	}

	p.Map.Index++

	return err
}
