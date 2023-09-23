package resources

import (
	"errors"
)

type Science struct {
	Any         int `json:"any"`
	Single      int `json:"single"`
	Banking     int `json:"banking"`
	Law         int `json:"law"`
	Engineering int `json:"engineering"`
	Medicine    int `json:"medicine"`
}

type BuildVP struct {
	D          int
	TP         int
	TE         int
	SHSA       int
	Spade      int
	Science    int
	City       int
	Advance    int
	Innovation int
	Edge       int
	River      int
}

type Price struct {
	Coin      int     `json:"coin"`
	Worker    int     `json:"worker"`
	Prist     int     `json:"prist"`
	Power     int     `json:"power"`
	Spade     int     `json:"spade"`
	Bridge    int     `json:"bridge"`
	Book      int     `json:"book"`
	TpUpgrade int     `json:"tpUpgrade"`
	TpVP      int     `json:"tpVP"`
	Science   Science `json:"science"`
	VP        int     `json:"vp"`
}

type Resource struct {
	Coin      int     `json:"coin"`
	Worker    int     `json:"worker"`
	Prist     int     `json:"prist"`
	Power     [3]int  `json:"power"`
	Spade     int     `json:"spade"`
	Bridge    int     `json:"bridge"`
	Book      int     `json:"book"`
	TpUpgrade int     `json:"tpupgrade"`
	TpVP      int     `json:"tpvp"`
	City      int     `json:"city"`
	Science   Science `json:"science"`
	VP        int     `json:"vp"`
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

	if have.Book < need.Book {
		return errors.New("not enough book")
	}

	if have.Power[2]+have.Power[1]/2 < need.Power {
		return errors.New("not enough power")
	}

	return nil
}
