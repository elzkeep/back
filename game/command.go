package game

import (
	"aoi/game/resources"
	"aoi/global"
	"fmt"
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

func ConvertBook(str string) resources.Book {
	var book resources.Book
	temp := strings.Split(str, " ")

	for i := 0; i < len(temp)/2; i++ {
		name := temp[i*2]
		count := global.Atoi(temp[i*2+1])

		if name == "banking" {
			book.Banking += count
		} else if name == "law" {
			book.Law += count
		} else if name == "engineering" {
			book.Engineering += count
		} else if name == "medicine" {
			book.Medicine += count
		}
	}

	return book
}

func ConvertBookType(str string) resources.BookType {
	if str == "banking" {
		return resources.BookBanking
	} else if str == "law" {
		return resources.BookLaw
	} else if str == "engineering" {
		return resources.BookEngineering
	}

	return resources.BookMedicine
}

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
	} else if cmd == "book" {
		target := ConvertBookType(strs[2])
		level := global.Atoi(strs[3])
		err = p.Book(user, target, level)
	} else if cmd == "action" {
		action := strs[2][:2]

		if action == "AC" {
			pos := global.Atoi(strs[2][3:]) - 1

			if pos >= len(p.PowerActions.Items) {
				temp := strings.Split(str, "book ")
				book := ConvertBook(temp[1])

				err = p.BookAction(user, pos-len(p.PowerActions.Items), book)
			} else {
				err = p.PowerAction(user, pos)
			}
		} else if action == "PA" {
			pos := global.Atoi(strs[2][6:]) - 1
			err = p.TileAction(user, resources.TilePalace, pos)
		} else if action == "RO" {
			pos := global.Atoi(strs[2][5:]) + int(resources.TilePalaceVp)
			log.Println("ROUND", pos)
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

		err = p.PalaceTile(user, pos)
	} else if cmd == "schooltile" {
		science := ConvertScience(strs[2])
		level := 3 - global.Atoi(strs[3])

		err = p.SchoolTile(user, int(science), level)
	} else if cmd == "bridge" {
		x, y := ConvertPosition(strs[2])
		x2, y2 := ConvertPosition(strs[3])

		err = p.Bridge(user, x, y, x2, y2)
	} else if cmd == "city" {
		pos := global.Atoi(strs[2]) - 1

		err = p.City(user, resources.CityType(pos))
	} else if cmd == "burn" {
		count := global.Atoi(strs[2]) - 1

		err = p.Burn(user, count)
	} else if cmd == "convert" {
		category := strs[2]
		source := resources.Price{}
		target := resources.Price{}

		pos := 0

		if category == "book" {
			book := ConvertBook(fmt.Sprintf("%v %v", strs[3], strs[4]))
			source.Book = book

			pos = 6
		} else {
			count := global.Atoi(strs[3])

			if category == "power" {
				source.Power = count
			} else if category == "prist" {
				source.Prist = count
			} else if category == "worker" {
				source.Worker = count
			}

			pos = 5
		}

		targetCategory := strs[pos]
		targetCount := global.Atoi(strs[pos+1])

		if targetCategory == "book" {
			target.Book.Any = targetCount
		} else if targetCategory == "prist" {
			target.Prist = targetCount
		} else if targetCategory == "worker" {
			target.Worker = targetCount
		} else if targetCategory == "coin" {
			target.Coin = targetCount
		}

		err = p.Convert(user, source, target)
	} else if cmd == "annex" {
		x, y := ConvertPosition(strs[2])

		err = p.Annex(user, x, y)
	} else if cmd == "save" {
		p.TurnEnd(user)
	}

	p.Map.Index++

	return err
}
