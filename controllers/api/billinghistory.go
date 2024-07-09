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

	db := c.NewConnection()

	conn, _ := db.Begin()
	defer conn.Rollback()

	billinghistoryManager := models.NewBillinghistoryManager(conn)
	billingManager := models.NewBillingManager(conn)

	item.Company = session.Company
	billinghistoryManager.Insert(item)

	billingItem := billingManager.Get(item.Billing)
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
	db := c.NewConnection()

	conn, _ := db.Begin()
	defer conn.Rollback()

	billinghistoryManager := models.NewBillinghistoryManager(conn)
	billingManager := models.NewBillingManager(conn)

	log.Println(item)
	for _, v := range *item {
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

	billingItem := billingManager.Get(id)
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
