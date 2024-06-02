package api

import (
	"zkeep/controllers"
	"zkeep/models"
)

type DashboardController struct {
	controllers.Controller
}

func (c *DashboardController) InitData() {
	conn := c.NewConnection()

	session := c.Session

	noticeManager := models.NewNoticeManager(conn)
	//companyManager := models.NewCompanyManager(conn)
	customerManager := models.NewCustomerManager(conn)
	userManager := models.NewUserManager(conn)
	reportManager := models.NewReportManager(conn)

	notices := noticeManager.Find([]interface{}{
		models.Paging(1, 4),
		models.Ordering("n_id desc"),
	})

	// companys := companyManager.Find([]interface{}{
	// 	models.Where{Column: "company", Value: session.Company, Compare: "="},
	// })

	customers := customerManager.Find([]interface{}{
		models.Where{Column: "company", Value: session.Company, Compare: "="},
		models.Where{Column: "status", Value: 1, Compare: "="},
	})

	users := userManager.Count([]interface{}{
		models.Where{Column: "company", Value: session.Company, Compare: "="},
		models.Where{Column: "status", Value: 1, Compare: "="},
	})

	reports := reportManager.Find([]interface{}{
		models.Where{Column: "company", Value: session.Company, Compare: "="},
	})

	c.Set("notices", notices)
	//c.Set("companys", companys)
	c.Set("customers", customers)
	c.Set("users", users)
	c.Set("reports", reports)
}
