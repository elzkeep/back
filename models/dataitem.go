package models

type Dataitem struct {
	Report int64  `json:"report"`
	Data   Data   `json:"data"`
	Items  []Item `json:"items"`
}
