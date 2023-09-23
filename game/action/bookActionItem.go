package action

import "aoi/game/resources"

type BookActionItem struct {
	Type    BookActionType  `json:"type"`
	Name    string          `json:"name"`
	Book    int             `json:"book"`
	Receive resources.Price `json:"receive"`
	Use     bool            `json:"use"`
}
