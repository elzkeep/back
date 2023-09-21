package game

import (
	"aoi/game/color"
	"aoi/game/resources"
)

type Mapitem struct {
	Type     color.Color        `json:"type"`
	Owner    color.Color        `json:"owner"`
	Building resources.Building `json:"building"`
}

func NewMapitem() *Mapitem {
	var item Mapitem

	item.Type = color.Empty
	item.Owner = color.Empty
	item.Building = resources.None

	return &item
}
