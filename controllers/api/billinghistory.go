package api

import (
	"log"
	"zkeep/controllers"
	"zkeep/models"
	"zkeep/models/billing"
)

type BillinghistoryController struct {
	controllers.Controller
}

// @Post()
func (c *BillinghistoryController) Deposit(item *models.Billinghistory) {
	session := c.Session

	if session == nil || item == nil {
		c.Error("invalid request")
		return
	}

	db := c.NewConnection()

	conn, _ := db.Begin()
	defer conn.Rollback()

	billinghistoryManager := models.NewBillinghistoryManager(conn)
	billingManager := models.NewBillingManager(conn)

	billingItem := billingManager.Get(item.Billing)
	if billingItem == nil || billingItem.Company != session.Company {
		c.Error("billing not found")
		return
	}

	item.Company = session.Company
	item.Building = billingItem.Building
	billinghistoryManager.Insert(item)

	billingItem.Depositprice += item.Price
	if billingItem.Depositprice == billingItem.Price+billingItem.Vat {
		billingItem.Status = billing.StatusComplete
	} else {
		billingItem.Status = billing.StatusPart
	}

	billingManager.Update(billingItem)

	conn.Commit()
}

// @Post()
func (c *BillinghistoryController) Depositdelete(id int64, item *[]models.Billinghistory) {
	session := c.Session

	if session == nil || item == nil {
		c.Error("invalid request")
		return
	}

	db := c.NewConnection()

	conn, _ := db.Begin()
	defer conn.Rollback()

	billinghistoryManager := models.NewBillinghistoryManager(conn)
	billingManager := models.NewBillingManager(conn)

	billingItem := billingManager.Get(id)
	if billingItem == nil || billingItem.Company != session.Company {
		c.Error("billing not found")
		return
	}

	log.Println(item)
	for _, v := range *item {
		history := billinghistoryManager.Get(v.Id)
		if history == nil || history.Billing != id {
			c.Error("billinghistory not found")
			return
		}

		billinghistoryManager.Delete(v.Id)
	}

	billinghistorys := billinghistoryManager.Find([]interface{}{
		models.Where{Column: "billing", Value: id, Compare: "="},
	})

	total := 0
	for _, v := range billinghistorys {
		total += v.Price
	}

	log.Println("total", total)

	billingItem.Depositprice = total
	if billingItem.Depositprice == 0 {
		billingItem.Status = billing.StatusWait
	} else if billingItem.Depositprice == billingItem.Price+billingItem.Vat {
		billingItem.Status = billing.StatusComplete
	} else {
		billingItem.Status = billing.StatusPart
	}

	billingManager.Update(billingItem)

	conn.Commit()
}
