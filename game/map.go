package game

import (
	"errors"
	"fmt"
	"log"
	"math"
	"strings"

	"aoi/game/color"
	"aoi/game/resources"
)

type Map struct {
	Type       int                        `json:"type"`
	Mx         int                        `json:"-"`
	Width      int                        `json:"width"`
	Height     int                        `json:"height"`
	Data       [][]Mapitem                `json:"data"`
	BridgeList []resources.BridgePosition `json:"bridge"`
}

func NewMap() *Map {
	var item Map

	item.BridgeList = make([]resources.BridgePosition, 0)
	item.Width = 9
	item.Height = 13
	item.Type = 2
	item.Mx = 1
	width := item.Width

	if item.Type == 2 {
		item.Mx = 2
		width++
	}

	item.Data = make([][]Mapitem, item.Width+2)

	for i := 0; i < item.Width+2; i++ {
		line := make([]Mapitem, item.Height+2)
		item.Data[i] = line
	}

	for i := 0; i < item.Width+2; i++ {
		for j := 0; j < item.Height+2; j++ {
			item.Data[i][j].Owner = color.None
			item.Data[i][j].Building = resources.None
		}
	}

	maps := []string{
		"GXSKBRWXSKBYX",
		"WXXGSGYXGWSXK",
		"RSKXWBKXXRXXR",
		"YWRXYRXKXXBWG",
		"BGBYXXXGWSRYB",
		"KXXXSWXYBXKSR",
		"SXBKXGBXXXXXX",
		"XYSXRKYRKSGKX",
		"GWRXYWBSGWRYB",
	}

	y := 0
	for _, v := range maps {
		line := strings.Split(v, "")

		x := 0

		for _, v := range line {
			col := color.None

			if v == "R" {
				col = color.Red
			} else if v == "Y" {
				col = color.Yellow
			} else if v == "B" {
				col = color.Brown
			} else if v == "K" {
				col = color.Black
			} else if v == "S" {
				col = color.Blue
			} else if v == "G" {
				col = color.Green
			} else if v == "W" {
				col = color.Gray
			} else if v == "X" {
				col = color.River
			}

			item.SetType(y, x, col)
			x++
		}

		y++
	}

	return &item
}

func (p *Map) SetType(x int, y int, value color.Color) {
	p.Data[x+p.Mx][y+1].Type = value
}

func (p *Map) GetType(x int, y int) color.Color {
	return p.Data[x+p.Mx][y+1].Type
}

func (p *Map) SetOwner(x int, y int, value color.Color) {
	log.Println("SetOwner", x, y)
	p.Data[x+p.Mx][y+1].Owner = value
}

func (p *Map) GetOwner(x int, y int) color.Color {
	return p.Data[x+p.Mx][y+1].Owner
}

func (p *Map) SetBuilding(x int, y int, value resources.Building) {
	p.Data[x+p.Mx][y+1].Building = value
}

func (p *Map) GetBuilding(x int, y int) resources.Building {
	return p.Data[x+p.Mx][y+1].Building
}

func (p *Map) GetPower(x int, y int) int {
	building := p.GetBuilding(x, y)
	return building.Power()
}

func (p *Map) abs(value int) int {
	return int(math.Abs(float64(value)))
}

func (p *Map) GetDistance(x1 int, y1 int, x2 int, y2 int) int {
	if p.Type == 2 {
		x1++
		x2++
	}

	dx := x2 - x1
	cnt := 0

	x := 0

	cnt = p.abs(dx) + 1
	if dx < 0 {
		if x2%2 == 0 {
			cnt--
		}

		if x1%2 == 1 {
			cnt--
		}

		x = y1 - (y2 - (cnt / 2))
	} else {
		if x1%2 == 0 {
			cnt--
		}

		if x2%2 == 1 {
			cnt--
		}

		x = y1 - (y2 + (cnt / 2))
	}

	y := 0

	cnt = p.abs(dx) + 1
	if dx < 0 {
		if x2%2 == 1 {
			cnt--
		}

		if x1%2 == 0 {
			cnt--
		}

		y = y1 - (y2 + (cnt / 2))
	} else {
		if x1%2 == 1 {
			cnt--
		}

		if x2%2 == 0 {
			cnt--
		}

		y = y1 - (y2 - (cnt / 2))
	}

	z := x2 - x1
	distance := (p.abs(x) + p.abs(y) + p.abs(z)) / 2

	return distance
}

func (p *Map) CheckBuild(x int, y int, colorval color.Color, spade int) error {
	typeval := p.GetType(x, y)

	if typeval <= color.River {
		return errors.New("river")
	}

	if p.GetOwner(x, y) != color.None {
		return errors.New("already")
	}

	need := int(math.Abs(float64(typeval) - float64(colorval)))
	if need > 3 {
		need = 7 - need
	}

	log.Printf("need spade = %v\n", need)

	if spade >= need {
		return nil
	} else {
		return errors.New("need spade")
	}
}

func (p *Map) GetNeedSpade(x int, y int, colorval color.Color) int {
	typeval := p.GetType(x, y)

	if typeval <= color.River {
		return 0
	}

	need := int(math.Abs(float64(typeval) - float64(colorval)))
	if need > 3 {
		need = 7 - need
	}

	return need
}

func (p *Map) Build(x int, y int, color color.Color, building resources.Building) {
	p.Data[x+p.Mx][y+1].Type = color
	p.Data[x+p.Mx][y+1].Owner = color
	p.Data[x+p.Mx][y+1].Building = building
}

func (p *Map) Print() {
	str := "\n"

	width := p.Width
	if p.Type == 2 {
		width--
	}
	for j := 0; j < p.Height+2; j++ {
		if j == 0 || j == p.Height+1 {
			if j == 0 {
				for i := 0; i < width; i++ {
					str += fmt.Sprintf("|--------")
				}

				str += "|\n"

				for i := 0; i < width; i++ {
					str += fmt.Sprintf("|    %2v  ", i-1)
				}

				str += "|\n"

			}

			for i := 0; i < width; i++ {
				str += fmt.Sprintf("|--------")
			}

			str += "|\n"
			continue
		}

		for i := 0; i < p.Width; i++ {
			if p.Type == 2 && i == 1 {
				continue
			}

			if i%2 == 0 {
				if i == 0 {
					str += fmt.Sprintf("|    %2v  ", j-1)
				} else {
					item := p.Data[i][j]
					c := fmt.Sprintf("%v %4s", item.Owner.ToString()[:1], item.Building.ToString())
					str += fmt.Sprintf("| %v ", item.Type.ToStringBackground(c))
				}
			} else {
				str += fmt.Sprintf("|--------")
			}
		}

		str += "|\n"

		for i := 0; i < p.Width; i++ {
			if p.Type == 2 && i == 1 {
				continue
			}

			if i%2 == 0 {
				str += fmt.Sprintf("|--------")
			} else {
				item := p.Data[i][j]
				c := fmt.Sprintf("%v %4s", item.Owner.ToString()[:1], item.Building.ToString())
				str += fmt.Sprintf("| %v ", item.Type.ToStringBackground(c))
			}
		}

		str += "|\n"
	}

	fmt.Println(str)
}

func (p *Map) CheckBridge(user color.Color, x1 int, y1 int, x2 int, y2 int) error {
	if p.GetOwner(x1, y1) != user && p.GetOwner(x2, y2) != user {
		return errors.New("not owner")
	}

	for _, v := range p.BridgeList {
		if (x1 == v.X1 && x2 == v.X2 && y1 == v.Y1 && y2 == v.Y2) ||
			(x1 == v.X2 && x2 == v.X1 && y1 == v.Y2 && y2 == v.Y1) {
			return errors.New("already")
		}
	}

	return nil
}

func (p *Map) Bridge(user color.Color, x1 int, y1 int, x2 int, y2 int) error {
	p.BridgeList = append(p.BridgeList, resources.BridgePosition{X1: x1, Y1: y1, X2: x2, Y2: y2, Color: user})

	return nil
}

func (p *Map) CheckDistance(user color.Color, distance int, x int, y int) bool {
	positions := resources.GetGroundPosition(x, y)

	for _, position := range positions {
		owner := p.GetOwner(position.X, position.Y)

		if owner == user {
			return true
		}
	}

	for _, v := range p.BridgeList {
		if x == v.X1 && y == v.Y1 {
			if p.GetOwner(v.X2, v.Y2) == user {
				return true
			}
		}

		if x == v.X2 && y == v.Y2 {
			if p.GetOwner(v.X1, v.Y1) == user {
				return true
			}
		}
	}

	return p.FindRiver(user, x, y, 1, distance)
}

func (p *Map) FindRiver(user color.Color, x int, y int, distance int, maxDistance int) bool {
	if distance == maxDistance {
		return false
	}

	positions := resources.GetGroundPosition(x, y)

	for _, v := range positions {
		if p.GetType(v.X, v.Y) != color.River {
			continue
		}

		log.Println(v.X, v.Y)

		items := resources.GetGroundPosition(v.X, v.Y)
		for _, item := range items {
			if p.GetType(item.X, item.Y) == user {
				return true
			}
		}

		p.FindRiver(user, v.X, v.Y, distance+1, maxDistance)
	}

	return false
}

func (p *Map) CheckCity(user color.Color, x int, y int) bool {
	lists := p.GetBuildingList(user, x, y, make([]resources.Position, 0))

	items := resources.Unique(lists)

	total := 0
	for _, v := range items {
		total += v.Building.Power()
	}

	log.Println("+++++++++++++++++++++++++++++++++")
	log.Println("total :", total)
	log.Println(items)
	log.Println("---------------------------------")

	if total >= 7 {
		return true
	} else {
		return false
	}
}

func (p *Map) GetBuildingList(user color.Color, x int, y int, lists []resources.Position) []resources.Position {
	lists = append(lists, resources.Position{X: x, Y: y, Building: p.GetBuilding(x, y)})
	grounds := resources.GetGroundPosition(x, y)

	for _, v := range p.BridgeList {
		if x == v.X1 && y == v.Y1 {
			if p.GetOwner(v.X2, v.Y2) == user {
				grounds = append(grounds, resources.Position{X: v.X2, Y: v.Y2, Building: p.GetBuilding(v.X2, v.Y2)})
			}
		}

		if x == v.X2 && y == v.Y2 {
			if p.GetOwner(v.X1, v.Y1) == user {
				grounds = append(grounds, resources.Position{X: v.X1, Y: v.Y1, Building: p.GetBuilding(v.X1, v.Y1)})
			}
		}
	}

	for _, ground := range grounds {
		if p.GetOwner(ground.X, ground.Y) != user {
			continue
		}

		flag := false
		for _, v := range lists {
			if ground.X == v.X && ground.Y == v.Y {
				flag = true
				break
			}
		}

		if flag == true {
			continue
		}

		lists = p.GetBuildingList(user, ground.X, ground.Y, lists)
	}

	return lists
	// 주위 6개 좌표에서 내것
	// 연결된 다리
	//
	// 목록에 대해서 겹지는 것 제외
	// 파워 합계
	// 연결된 것중 이미 마을이면 마울 불가능

}
