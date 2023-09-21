package game

import (
	"aoi/game/resources"
	"aoi/game/resources/city"
)

type CityType int

const (
	WorkerCity CityType = iota
	SpaceCity
	BookCity
	CoinCity
	ScienceCity
	PowerCity
	PristCity
)

type City struct {
	Items [][]city.CityItem
}

func NewCity() *City {
	var item City

	item.Items = make([][]city.CityItem, 0)

	for i := 0; i < 7; i++ {
		item.Items = append(item.Items, make([]city.CityItem, 0))
	}

	for i := 0; i < 3; i++ {
		item.Items[0] = append(item.Items[0], city.CityItem{Name: "3 worker", Receive: resources.Price{Worker: 3, VP: 4}, Use: false})
		item.Items[1] = append(item.Items[1], city.CityItem{Name: "2 space", Receive: resources.Price{Spade: 2, VP: 5}, Use: false})
		item.Items[2] = append(item.Items[2], city.CityItem{Name: "2 book", Receive: resources.Price{Book: 2, VP: 5}, Use: false})
		item.Items[3] = append(item.Items[3], city.CityItem{Name: "6 coin", Receive: resources.Price{Coin: 6, VP: 6}, Use: false})
		item.Items[4] = append(item.Items[4], city.CityItem{Name: "science", Receive: resources.Price{Science: resources.Science{Banking: 1, Law: 1, Engineering: 1, Medicine: 1}, VP: 7}, Use: false})
		item.Items[5] = append(item.Items[5], city.CityItem{Name: "8 power", Receive: resources.Price{Power: 8, VP: 8}, Use: false})
		item.Items[6] = append(item.Items[6], city.CityItem{Name: "1 prist", Receive: resources.Price{Prist: 1, VP: 8}, Use: false})
	}

	return &item
}

func (p *City) IsRemain(pos CityType) bool {
	if len(p.Items[int(pos)]) == 0 {
		return false
	}

	return true
}

func (p *City) Use(item CityType) city.CityItem {
	pos := int(item)
	tile := p.Items[pos][0]

	p.Items[pos] = p.Items[pos][1:]

	return tile
}
