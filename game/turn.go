package game

import "aoi/game/resources"

type TurnType int

const (
	NormalTurn TurnType = iota
	PowerTurn
	ScienceTurn
	SpadeTurn
	BookTurn
	TileTurn
	BuildTurn
	ResourceTurn
)

type Turn struct {
	User    int               `json:"user"`
	Type    TurnType          `json:"type"`
	From    int               `json:"from"`
	Power   int               `json:"power"`
	Science resources.Science `json:"science"`
}

func (p *Turn) Print() {
	/*
		titles := []string{"Normal", "Power", "Science", "Spade", "Book", "Tile", "Build", "Resource"}
		log.Printf("user = %v, type = %v\n", p.User, titles[int(p.Type)])
	*/
}
