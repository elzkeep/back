package resources

type Position struct {
	X        int      `json:"x"`
	Y        int      `json:"y"`
	Building Building `json:"building"`
}

func GetGroundPosition(x int, y int) []Position {
	/*
		        0,1  0,2
			  1,0 1,1  1,2
			    2,1, 2,2

		         1,0 1,1
			   2.0 2.1 2.2
			     2,0 2,2
	*/

	dx := 0
	if x%2 == 1 {
		dx = 1
	}

	positions := []Position{{X: x - 1, Y: y - 1 + dx}, {X: x - 1, Y: y + 0 + dx}, {X: x + 0, Y: y - 1}, {X: x + 0, Y: y + 1}, {X: x + 1, Y: y - 1 + dx}, {X: x + 1, Y: y + 0 + dx}}

	return positions
}

func Unique(lists []Position) []Position {
	items := make([]Position, 0)
	for _, v := range lists {
		flag := false
		for _, item := range items {
			if v.X == item.X && v.Y == item.Y {
				flag = true
				break
			}
		}

		if flag == true {
			continue
		}

		items = append(items, v)
	}

	return items
}
