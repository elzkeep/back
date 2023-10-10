package city

import "aoi/game/resources"

type CityItem struct {
	Name    string
	Type    CityType
	VP      int
	Receive resources.Price
	Use     bool
}
