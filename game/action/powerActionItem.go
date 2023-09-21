package action

import "aoi/game/resources"

type PowerActionItem struct {
	Name    string
	Power   int
	Receive resources.Price
	Use     bool
}
