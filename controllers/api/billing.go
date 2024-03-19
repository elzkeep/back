package api

import (
	"log"
	"zkeep/controllers"
)

type BillingController struct {
	controllers.Controller
}

func (c *BillingController) Print(ids []int64) {
	log.Println(ids)
}
