package resources

import (
	"errors"
)

type Science struct {
	Any         int
	Single      int
	Banking     int
	Law         int
	Engineering int
	Medicine    int
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
	Coin      int
	Worker    int
	Prist     int
	Power     int
	Spade     int
	Bridge    int
	Book      int
	TpUpgrade int
	TpVP      int
	Science   Science
	VP        int
}

type Resource struct {
	Coin      int
	Worker    int
	Prist     int
	Power     [3]int
	Spade     int
	Bridge    int
	Book      int
	TpUpgrade int
	TpVP      int
	City      int
	Science   Science
	VP        int
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
