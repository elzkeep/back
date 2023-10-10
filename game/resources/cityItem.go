package resources

type CityType int

const (
	WorkerCity CityType = iota
	SpadeCity
	BookCity
	CoinCity
	ScienceCity
	PowerCity
	PristCity
)

type CityItem struct {
	Type    CityType `json:"type"`
	Name    string   `json:"name"`
	Receive Price    `json:"receive"`
	Use     bool     `json:"use"`
}
