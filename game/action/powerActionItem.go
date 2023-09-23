package action

import "aoi/game/resources"

type PowerActionItem struct {
	Type    PowerActionType `json:"type"`
	Name    string          `json:"name"`
	Power   int             `json:"power"`
	Receive resources.Price `json:"receive"`
	Use     bool            `json:"use"`
}
