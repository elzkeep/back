package game

import (
	"aoi/game/color"
	"aoi/game/factions"
	"aoi/game/resources"
	"fmt"
)

type ScienceType int

const (
	Banking ScienceType = iota
	Law
	Engineering
	Medicine
)

type Science struct {
	Count     [][]color.Color       `json:"count"`
	Value     []map[color.Color]int `json:"value"`
	UserCount int                   `json:"-"`
}

func NewScience(count int) *Science {
	var item Science

	item.UserCount = count
	item.Count = make([][]color.Color, 4)
	item.Value = make([]map[color.Color]int, 4)

	for i := 0; i < 4; i++ {
		item.Count[i] = make([]color.Color, 0)
		item.Value[i] = make(map[color.Color]int)
	}

	return &item
}

func (p *Science) AddUser(user color.Color, value []int) {
	for i, v := range value {
		p.Value[i][user] = v
	}
}

func (p *Science) Send(user *factions.Faction, pos ScienceType) int {
	step := 2

	count := len(p.Count[pos])
	if p.UserCount > 2 {
		if count == 0 {
			step = 3
		} else if count >= 4 {
			step = 1
		}
	} else {
		if count == 1 {
			step = 3
		} else if count >= 4 {
			step = 1
		}
	}

	if count < 4 {
		p.Count[pos] = append(p.Count[pos], user.Color)
	}

	return p.Action(user, pos, step)
}

func (p *Science) Supploy(user *factions.Faction, pos ScienceType) int {
	return p.Action(user, pos, 1)
}

func (p *Science) Action(user *factions.Faction, pos ScienceType, step int) int {
	if step == 0 {
		return 0
	}

	if p.Value[pos][user.Color] >= 12 {
		return 0
	}

	top := false
	for _, v := range p.Value[pos] {
		if v >= 12 {
			top = true
			break
		}
	}

	inc := 0

	for i := 0; i < step; i++ {
		if p.Value[pos][user.Color] == 7 {
			if user.Key == 0 {
				return inc
			}
			user.Key--

			for i, v := range user.Cities {
				if v.Use == true {
					user.Cities[i].Use = false
					break
				}
			}
		}

		if p.Value[pos][user.Color] == 11 {
			if top == true {
				break
			}
		}

		p.Value[pos][user.Color]++

		value := p.Value[pos][user.Color]
		if value == 3 {
			user.ReceivePower(1, false)
		} else if value == 5 {
			user.ReceivePower(2, false)
		} else if value == 7 {
			user.ReceivePower(2, false)
		} else if value == 12 {
			user.ReceivePower(3, false)
		}

		inc++

		if p.Value[pos][user.Color] == 12 {
			break
		}
	}

	return inc
}

func (p *Science) Receive(user *factions.Faction, resource resources.Price) (int, int, int, int) {
	inc1 := p.Action(user, Banking, resource.Science.Banking)
	inc2 := p.Action(user, Law, resource.Science.Law)
	inc3 := p.Action(user, Engineering, resource.Science.Engineering)
	inc4 := p.Action(user, Medicine, resource.Science.Medicine)

	return inc1, inc2, inc3, inc4
}

func (p *Science) Print() {
	for i := 12; i >= 0; i-- {
		fmt.Println("|------------|------------|------------|------------|")
		fmt.Printf("|")
		for j := 0; j < 4; j++ {
			values := p.Value[j]

			str := ""

			flag := 0
			for k, v := range values {
				if v == i {
					str += k.ToShortString() + " "
					flag++
				}
			}

			for k := 0; k < 4-flag; k++ {
				str += "   "
			}

			fmt.Printf("%v|", str)
		}
		fmt.Printf("\n")
	}

	fmt.Println("|------------|------------|------------|------------|")
}

func (p *Science) RoundBonus(user *factions.Faction) {
	for i := 0; i < 4; i++ {
		if p.Value[i][user.Color] < 9 {
			continue
		}

		if i == 0 {
			user.ReceiveResource(resources.Price{Coin: 3})
		} else if i == 1 {
			user.ReceiveResource(resources.Price{Power: 6})
		} else if i == 2 {
			user.ReceiveResource(resources.Price{Worker: 1})
		} else if i == 3 {
			user.ReceiveResource(resources.Price{VP: 3})
		}
	}
}

func (p *Science) RoundEndBonus(user factions.FactionInterface, tile RoundBonusItem) {
	f := user.GetInstance()

	if tile.Science.Banking > 0 {
		value := user.GetScience(int(Banking)) / tile.Science.Banking

		for i := 0; i < value; i++ {
			f.ReceiveResource(tile.Receive)
		}
	}

	if tile.Science.Law > 0 {
		value := user.GetScience(int(Law)) / tile.Science.Law

		for i := 0; i < value; i++ {
			f.ReceiveResource(tile.Receive)
		}
	}

	if tile.Science.Engineering > 0 {
		value := user.GetScience(int(Engineering)) / tile.Science.Engineering

		for i := 0; i < value; i++ {
			f.ReceiveResource(tile.Receive)
		}
	}

	if tile.Science.Medicine > 0 {
		value := user.GetScience(int(Medicine)) / tile.Science.Medicine

		for i := 0; i < value; i++ {
			f.ReceiveResource(tile.Receive)
		}
	}
}

func (p *Science) CalculateRoundEndBonus(user factions.FactionInterface, tile RoundBonusItem) {
	f := user.GetInstance()

	if tile.Science.Banking > 0 {
		value := user.GetScience(int(Banking)) / tile.Science.Banking

		for i := 0; i < value; i++ {
			f.ReceiveIncome(tile.Receive)
		}
	}

	if tile.Science.Law > 0 {
		value := user.GetScience(int(Law)) / tile.Science.Law

		for i := 0; i < value; i++ {
			f.ReceiveIncome(tile.Receive)
		}
	}

	if tile.Science.Engineering > 0 {
		value := user.GetScience(int(Engineering)) / tile.Science.Engineering

		for i := 0; i < value; i++ {
			f.ReceiveIncome(tile.Receive)
		}
	}

	if tile.Science.Medicine > 0 {
		value := user.GetScience(int(Medicine)) / tile.Science.Medicine

		for i := 0; i < value; i++ {
			f.ReceiveIncome(tile.Receive)
		}
	}
}

func (p *Science) CalculateRoundBonus(user *factions.Faction) {
	for i := 0; i < 4; i++ {
		if p.Value[i][user.Color] < 9 {
			continue
		}

		if i == 0 {
			user.ReceiveIncome(resources.Price{Coin: 3})
		} else if i == 1 {
			user.ReceiveIncome(resources.Price{Power: 6})
		} else if i == 2 {
			user.ReceiveIncome(resources.Price{Worker: 1})
		} else if i == 3 {
			user.ReceiveIncome(resources.Price{VP: 3})
		}
	}
}

func (p *Science) Init(value []int) {
	if p.UserCount > 2 {
		return
	}

	for i := 0; i < 4; i++ {
		p.Count[i] = append(p.Count[i], color.None)
	}

	p.AddUser(color.None, value)
}
