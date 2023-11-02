package resources

import "aoi/game/color"

type TileCategory int

const (
	TilePalace TileCategory = iota
	TileSchool
	TileInnovation
	TileRound
	TileFaction
	TileColor
	TileRoundBonus
	TileBookAction
)

type TileType int

const (
	TilePalaceWorker TileType = iota
	TilePalaceSpade
	TilePalaceDowngrade
	TilePalaceTpUpgrade
	TilePalaceSchoolTile
	TilePalaceScience
	TilePalaceSchoolVp
	TilePalace6PowerCity
	TilePalaceJump
	TilePalacePower
	TilePalaceCity
	TilePalaceDVp
	TilePalaceTpVp
	TilePalaceRiverCity
	TilePalaceBridge
	TilePalaceTpBuild
	TilePalaceVp

	TileRoundEdgeVP
	TileRoundPristVP
	TileRoundTpVP
	TileRoundShVP
	TileRoundSpade
	TileRoundBridge
	TileRoundScienceCube
	TileRoundSchoolScienceCoin
	TileRoundPower
	TileRoundCoin

	TileSchoolWorker
	TileSchoolSpade
	TileSchoolPrist
	TileSchoolEdgeVP
	TileSchoolCoin
	TileSchoolAnnex
	TileSchoolNeutral
	TileSchoolBook
	TileSchoolVP
	TileSchoolPower
	TileSchoolPassCity
	TileSchoolPassPrist

	TileFactionBlessed
	TileFactionFelines
	TileFactionGoblins
	TileFactionIllusionists
	TileFactionInventors
	TileFactionLizards
	TileFactionMoles
	TileFactionMonks
	TileFactionNavigators
	TileFactionOmar
	TileFactionPhilosophers
	TileFactionPsychics

	TileColorYellow
	TileColorGreen
	TileColorBlue
	TileColorGray
	TileColorBrown
	TileColorBlack
	TileColorRed

	TileInnovationKind
	TileInnovationCount
	TileInnovationSchool
	TileInnovationCity
	TileInnovationScience
	TileInnovationCluster
	TileInnovationD
	TileInnovationUpgrade
	TileInnovationBridge
	TileInnovationFreeD
	TileInnovationFreeTP
	TileInnovationFreeSchool
	TileInnovationFreeSA
	TileInnovationFreeSH
	TileInnovationFreeMT
)

type TileItem struct {
	Type     TileType     `json:"type"`
	Category TileCategory `json:"category"`
	Name     string       `json:"name"`
	Receive  Price        `json:"receive"`
	Action   Price        `json:"action"`
	Once     Price        `json:"once"`
	Pass     Price        `json:"pass"`
	Build    BuildVP      `json:"build"`
	Use      bool         `json:"use"`
	Ship     int          `json:"ship"`
	Coin     int          `json:"coin"`
	Color    color.Color  `json:"color"`
}

type Tile struct {
	Items []TileItem `json:"items"`
}

func NewTile() *Tile {
	var item Tile

	item.Items = make([]TileItem, 0)

	return &item
}
