package resources

import "aoi/game/color"

type BridgePosition struct {
	X1    int         `json:"x1"`
	Y1    int         `json:"y1"`
	X2    int         `json:"x2"`
	Y2    int         `json:"y2"`
	Color color.Color `json:"color"`
}
