package game

import (
	"errors"
	"fmt"
	"math"
	"strings"

	"aoi/game/color"
	"aoi/game/resources"
)

type Map struct {
	Index      int                        `json:"index"`
	Type       int                        `json:"type"`
	Mx         int                        `json:"-"`
	Width      int                        `json:"width"`
	Height     int                        `json:"height"`
	Data       [][]Mapitem                `json:"data"`
	BridgeList []resources.BridgePosition `json:"bridge"`
	AnnexList  []resources.Position       `json:"annex"`
	CityList   []resources.Position       `json:"city"`
	LastBuild  resources.Position         `json:"lastBuild"`
	LastDig    resources.Position         `json:"lastDig"`
}

func NewMap(typeid int64) *Map {
	var item Map

	item.BridgeList = make([]resources.BridgePosition, 0)
	item.AnnexList = make([]resources.Position, 0)
	item.CityList = make([]resources.Position, 0)
	item.Width = 9
	item.Height = 13
	item.Type = 2
	item.Mx = 1
	item.LastBuild = resources.Position{X: -1, Y: -1}
	item.LastDig = resources.Position{X: -1, Y: -1}

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

	if typeid == 1 {
		maps = []string{
			"GWYBXSGWRX",
			"KSKRXKBSXG",
			"XBWGXXYXXY",
			"XRYXGXXWBK",
			"XXXXKBSYRW",
			"XSBXRWXGSY",
			"XXKWXXXXXX",
			"XYGRYGSKGX",
			"XRBWSKBYRW",
		}

		item.Width = 9
		item.Height = 10
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

	item.Index = 0
	return &item
}

func (p *Map) SetType(x int, y int, value color.Color) {
	p.Data[x+p.Mx][y+1].Type = value
	p.Index++
}

func (p *Map) GetType(x int, y int) color.Color {
	if x+p.Mx < 0 {
		return color.None
	}

	if x+p.Mx >= p.Width+2 {
		return color.None
	}

	if y+1 < 0 {
		return color.None
	}

	if y+1 >= p.Height+2 {
		return color.None
	}

	return p.Data[x+p.Mx][y+1].Type
}

func (p *Map) SetOwner(x int, y int, value color.Color) {
	p.Data[x+p.Mx][y+1].Owner = value
	p.Index++
}

func (p *Map) GetOwner(x int, y int) color.Color {
	if x+p.Mx < 0 {
		return color.None
	}

	if x+p.Mx >= p.Width+2 {
		return color.None
	}

	if y+1 < 0 {
		return color.None
	}

	if y+1 >= p.Height+2 {
		return color.None
	}

	return p.Data[x+p.Mx][y+1].Owner
}

func (p *Map) SetBuilding(x int, y int, value resources.Building) {
	p.Data[x+p.Mx][y+1].Building = value
	p.Index++
}

func (p *Map) GetBuilding(x int, y int) resources.Building {
	if x+p.Mx < 0 {
		return resources.None
	}

	if x+p.Mx >= p.Width+2 {
		return resources.None
	}

	if y+1 < 0 {
		return resources.None
	}

	if y+1 >= p.Height+2 {
		return resources.None
	}

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

func (p *Map) IsRiverside(x int, y int) bool {
	items := resources.GetGroundPosition(x, y)

	for _, v := range items {
		if p.GetType(v.X, v.Y) == color.River {
			return true
		}
	}

	return false
}

func (p *Map) IsEdge(x int, y int) bool {
	if p.GetType(x-1, y) == color.None {
		return true
	}

	if p.GetType(x+1, y) == color.None {
		return true
	}

	if p.GetType(x, y-1) == color.None {
		return true
	}

	if p.GetType(x, y+1) == color.None {
		return true
	}

	return false
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
	p.Index++

	p.LastBuild = resources.Position{X: x, Y: y}
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

	if p.GetType(x1, y1) == color.River || p.GetType(x2, y2) == color.River || p.GetType(x1, y1) == color.None || p.GetType(x2, y2) == color.None {
		return errors.New("can't build")
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

		if position.X == p.LastBuild.X && position.Y == p.LastBuild.Y {
			continue
		}

		if owner == user {
			return true
		}
	}

	for _, v := range p.BridgeList {
		if v.Color != user {
			continue
		}

		if x == v.X1 && y == v.Y1 {
			if v.X2 == p.LastBuild.X && v.Y2 == p.LastBuild.Y {
				continue
			}

			if p.GetOwner(v.X2, v.Y2) == user {
				return true
			}
		}

		if x == v.X2 && y == v.Y2 {
			if v.X1 == p.LastBuild.X && v.Y1 == p.LastBuild.Y {
				continue
			}

			if p.GetOwner(v.X1, v.Y1) == user {
				return true
			}
		}
	}

	return p.FindRiver(user, x, y, 1, distance+1)
}

func (p *Map) FindRiver(user color.Color, x int, y int, distance int, maxDistance int) bool {
	if distance >= maxDistance {
		return false
	}

	positions := resources.GetGroundPosition(x, y)

	for _, v := range positions {
		if p.GetType(v.X, v.Y) != color.River {
			continue
		}

		items := resources.GetGroundPosition(v.X, v.Y)
		for _, item := range items {
			if item.X == p.LastBuild.X && item.Y == p.LastBuild.Y {
				continue
			}

			if p.GetType(item.X, item.Y) == user {
				return true
			}
		}

		if p.FindRiver(user, v.X, v.Y, distance+1, maxDistance) {
			return true
		}
	}

	return false
}

func (p *Map) CheckCity(user color.Color, x int, y int, power int) []resources.Position {
	lists := p.GetBuildingList(user, x, y, make([]resources.Position, 0))

	items := resources.Unique(lists)

	total := 0
	count := len(items)
	needCount := 4

	for _, v := range items {
		for _, v2 := range p.CityList {
			if v.X == v2.X && v.Y == v2.Y {
				return make([]resources.Position, 0)
			}
		}

		if v.Building == resources.WHITE_MT {
			needCount = 2
		} else if v.Building == resources.SA || v.Building == resources.WHITE_SA {
			if needCount > 3 {
				needCount = 3
			}
		}

		for _, a := range p.AnnexList {
			if v.X != a.X || v.Y != a.Y {
				continue
			}

			if a.Color != user {
				continue
			}

			total++
			count++
		}

		total += v.Building.Power()
	}

	if total >= power && count >= needCount {
		return items
	} else {
		return make([]resources.Position, 0)
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
}

func (p *Map) AddCityBuildingList(list []resources.Position) {
	p.CityList = append(p.CityList, list...)
}

func (p *Map) CheckAnnex(user color.Color, x int, y int) error {
	if p.GetOwner(x, y) != user {
		return errors.New("not owner")
	}

	for _, v := range p.AnnexList {
		if x == v.X && y == v.Y {
			return errors.New("already")
		}
	}

	return nil
}

func (p *Map) Annex(user color.Color, x int, y int) error {
	p.AnnexList = append(p.AnnexList, resources.Position{X: x, Y: y, Color: user})

	return nil
}

func (p *Map) CheckDistanceMoles(user color.Color, x int, y int) bool {
	dx := 0
	if x%2 == 1 {
		dx = 1
	}

	items := make([]resources.Position, 0)

	items = append(items, resources.Position{X: x - 2, Y: y - 1})
	items = append(items, resources.Position{X: x - 2, Y: y - 0})
	items = append(items, resources.Position{X: x - 2, Y: y + 1})

	items = append(items, resources.Position{X: x - 1, Y: y - 2 + dx})
	items = append(items, resources.Position{X: x - 1, Y: y + 1 + dx})

	items = append(items, resources.Position{X: x + 0, Y: y - 2})
	items = append(items, resources.Position{X: x + 0, Y: y + 2})

	items = append(items, resources.Position{X: x + 1, Y: y - 2 + dx})
	items = append(items, resources.Position{X: x + 1, Y: y + 1 + dx})

	items = append(items, resources.Position{X: x + 2, Y: y - 1})
	items = append(items, resources.Position{X: x + 2, Y: y - 0})
	items = append(items, resources.Position{X: x + 2, Y: y + 1})

	for _, v := range items {
		if p.LastBuild.X == v.X && p.LastBuild.Y == v.Y {
			continue
		}

		if p.GetOwner(v.X, v.Y) == user {
			return true
		}
	}

	return false
}

func (p *Map) CheckDistanceJump(items []resources.Position, x int, y int) bool {
	for _, v := range items {
		if p.LastBuild.X == v.X && p.LastBuild.Y == v.Y {
			continue
		}

		distance := p.GetDistance(x, y, v.X, v.Y)

		if distance == 2 || distance == 3 {
			return true
		}
	}

	return false
}

func (p *Map) CheckSolo(user color.Color, x int, y int) bool {
	items := resources.GetGroundPosition(x, y)

	for _, v := range items {
		owner := p.GetOwner(v.X, v.Y)
		if owner != color.None && owner != color.River && owner != user {
			return false
		}
	}

	for _, v := range p.BridgeList {
		if x == v.X1 && y == v.Y1 {
			owner := p.GetOwner(v.X2, v.Y2)
			if owner != color.None && owner != color.River && owner != user {
				return false
			}
		}

		if x == v.X2 && y == v.Y2 {
			owner := p.GetOwner(v.X1, v.Y1)
			if owner != color.None && owner != color.River && owner != user {
				return false
			}
		}
	}

	return true
}

func (p *Map) CheckConnect(user color.Color, distance int, x int, y int, x2 int, y2 int) bool {
	positions := resources.GetGroundPosition(x, y)

	for _, position := range positions {
		owner := p.GetOwner(position.X, position.Y)

		if x2 == position.X && y2 == position.Y && owner == user {
			return true
		}
	}

	for _, v := range p.BridgeList {
		if v.Color != user {
			continue
		}

		if x == v.X1 && y == v.Y1 {
			if x2 == v.X2 && y2 == v.Y2 && p.GetOwner(v.X2, v.Y2) == user {
				return true
			}
		}

		if x == v.X2 && y == v.Y2 {
			if x2 == v.X1 && y2 == v.Y1 && p.GetOwner(v.X1, v.Y1) == user {
				return true
			}
		}
	}

	return p.FindRiverConnect(user, x, y, x2, y2, 1, distance+1)
}

func (p *Map) FindRiverConnect(user color.Color, x int, y int, x2, y2, distance int, maxDistance int) bool {
	if distance >= maxDistance {
		return false
	}

	positions := resources.GetGroundPosition(x, y)

	for _, v := range positions {
		if p.GetType(v.X, v.Y) != color.River {
			continue
		}

		items := resources.GetGroundPosition(v.X, v.Y)
		for _, item := range items {
			if item.X == x2 && item.Y == y2 && p.GetType(item.X, item.Y) == user {
				return true
			}

		}

		if p.FindRiverConnect(user, v.X, v.Y, x2, y2, distance+1, maxDistance) {
			return true
		}
	}

	return false
}

func (p *Map) TurnEnd() {
	p.LastBuild = resources.Position{X: -1, Y: -1}
	p.LastDig = resources.Position{X: -1, Y: -1}
}

func (p *Map) Dig(x int, y int, value color.Color) {
	p.SetType(x, y, color.Color(value))

	if p.LastDig.X == -1 {
		p.LastDig = resources.Position{X: x, Y: y}
	}
}
