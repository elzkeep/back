package resources

import (
	"errors"
	"log"
)

type Science struct {
	Any         int `json:"any"`
	Single      int `json:"single"`
	Banking     int `json:"banking"`
	Law         int `json:"law"`
	Engineering int `json:"engineering"`
	Medicine    int `json:"medicine"`
}

type BookType int

const (
	BookBanking BookType = iota
	BookLaw
	BookEngineering
	BookMedicine
)

type Book struct {
	Any         int `json:"any"`
	Banking     int `json:"banking"`
	Law         int `json:"law"`
	Engineering int `json:"engineering"`
	Medicine    int `json:"medicine"`
}

func (p *Book) Count() int {
	return p.Any + p.Banking + p.Law + p.Engineering + p.Medicine
}

func (p *Book) IsEmpty() bool {
	if p.Any+p.Banking+p.Law+p.Engineering+p.Medicine == 0 {
		return true
	} else {
		return false
	}
}

func (p *Book) Print() {
	log.Printf("Any: %v, Banking: %v, Law: %v, Engineering: %v, Medicine: %v\n", p.Any, p.Banking, p.Law, p.Engineering, p.Medicine)
}

func (p *Science) IsEmpty() bool {
	if p.Any != 0 {
		return false
	}

	if p.Single != 0 {
		return false
	}

	if p.Banking != 0 {
		return false
	}

	if p.Law != 0 {
		return false
	}

	if p.Engineering != 0 {
		return false
	}

	if p.Medicine != 0 {
		return false
	}

	return true
}

type BuildVP struct {
	D          int `json:"d"`
	TP         int `json:"tp"`
	TE         int `json:"te"`
	SHSA       int `json:"shsa"`
	Spade      int `json:"spade"`
	Science    int `json:"science"`
	City       int `json:"city"`
	Advance    int `json:"advance"`
	Innovation int `json:"innovation"`
	Edge       int `json:"edge"`
	River      int `json:"river"`
	Prist      int `json:"prist"`
}

type Price struct {
	Coin        int      `json:"coin"`
	Worker      int      `json:"worker"`
	Prist       int      `json:"prist"`
	Power       int      `json:"power"`
	Spade       int      `json:"spade"`
	Bridge      int      `json:"bridge"`
	Book        Book     `json:"book"`
	TpUpgrade   int      `json:"tpUpgrade"`
	TpVP        int      `json:"tpVp"`
	City        int      `json:"city"`
	Science     Science  `json:"science"`
	VP          int      `json:"vp"`
	Downgrade   int      `json:"downgrade"`
	Tile        int      `json:"tile"`
	ShVP        int      `json:"shVp"`
	TeVP        int      `json:"teVp"`
	ShipUpgrade int      `json:"shipUpgrade"`
	Building    Building `json:"building"`
}

type Resource struct {
	Coin       int      `json:"coin"`
	Worker     int      `json:"worker"`
	Prist      int      `json:"prist"`
	Power      [3]int   `json:"power"`
	Spade      int      `json:"spade"`
	Bridge     int      `json:"bridge"`
	Book       Book     `json:"book"`
	TpUpgrade  int      `json:"tpUpgrade"`
	TpVP       int      `json:"tpVp"`
	City       int      `json:"city"`
	Science    Science  `json:"science"`
	VP         int      `json:"vp"`
	Downgrade  int      `json:"downgrade"`
	PalaceTile int      `json:"palaceTile"`
	SchoolTile int      `json:"schoolTile"`
	ShVP       int      `json:"shVp"`
	TeVP       int      `json:"teVp"`
	Building   Building `json:"building"`
}

func CheckResource(have Resource, need Price) error {
	if have.Coin < need.Coin {
		return errors.New("not enough coin")
	}

	if have.Worker < need.Worker {
		return errors.New("not enough worker")
	}

	if have.Prist < need.Prist {
		return errors.New("not enough prist")
	}

	if have.Spade < need.Spade {
		return errors.New("not enough spade")
	}

	if have.Bridge < need.Bridge {
		return errors.New("not enough bridge")
	}

	if have.Book.Count() < need.Book.Count() {
		return errors.New("not enough book")
	}

	if have.Power[2]+have.Power[1]/2 < need.Power {
		return errors.New("not enough power")
	}

	return nil
}
