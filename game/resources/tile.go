package resources

type TileCategory int

const (
	TilePalace TileCategory = iota
	TileSchool
	TileInnovation
	TileRound
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
	TilePalaceCity
	TilePalaceDVp
	TilePalaceTpVp
	TilePalaceRiverCity
	TilePalaceBridge
	TilePalaceTpBuild
	TilePalaceVp

	TileRoundSideVP
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
	TileSchoolSideVP
	TileSchoolCoin
	TileSchoolAnnex
	TileSchoolNeutral
	TileSchoolBook
	TileSchoolVP
	TileSchoolPower
	TileSchoolPassCity
	TileSchoolPassPrist
)

type TileItem struct {
	Type     TileType     `json:"type"`
	Category TileCategory `json:"category"`
	Name     string       `json:"name"`
	Receive  Price        `json:"receive"`
	Action   Price        `json:"action"`
	Once     Price        `json:"once"`
	Pass     Price        `json:"pass"`
	Use      bool         `json:"use"`
	Ship     int          `json:"ship"`
}

type Tile struct {
	Items []TileItem `json:"items"`
}

func NewTile() *Tile {
	var item Tile

	item.Items = make([]TileItem, 0)

	return &item
}

/*
func (p *Tile) GetTile(pos TileType) *TileItem {
	return &p.Items[pos]
}

func (p *Tile) Setup(pos int) {
	p.Items = append(p.Items[:pos], p.Items[pos+1:]...)
}
*/
