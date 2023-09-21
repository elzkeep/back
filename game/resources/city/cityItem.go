package city

import "aoi/game/resources"

type CityItem struct {
	Name    string
	VP      int
	Receive resources.Price
	Use     bool
}
