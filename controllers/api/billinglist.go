package api

import (
	"zkeep/controllers"
	"zkeep/models"
)

type BillinglistController struct {
	controllers.Controller
}

func (c *BillinglistController) InitData() {
	conn := c.NewConnection()

	session := c.Session

	companyManager := models.NewCompanyManager(conn)
	company := companyManager.Get(session.Company)

	c.Set("company", company)

	customercompanyManager := models.NewCustomercompanyManager(conn)
	customercompanys := customercompanyManager.Find([]interface{}{
		models.Where{Column: "company", Value: session.Company, Compare: "="},
		models.Ordering("c_name"),
	})

	companys := make([]models.Company, 0)
	for _, v := range customercompanys {
		company := v.Extra["company"].(models.Company)
		companys = append(companys, company)
	}

	c.Set("companys", companys)
}
