package action

import "aoi/game/resources"

type BookActionItem struct {
	Name    string
	Book    int
	Receive resources.Price
	Use     bool
}
