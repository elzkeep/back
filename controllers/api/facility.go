package api

import (
	"zkeep/controllers"
	"zkeep/global"
	"zkeep/models"
)

type FacilityController struct {
	controllers.Controller
}

func CalculateTotalweight(id int64) {
	conn := models.NewConnection()
	defer conn.Close()

	buildingManager := models.NewBuildingManager(conn)
	facilityManager := models.NewFacilityManager(conn)

	items := facilityManager.Find([]interface{}{
		models.Where{Column: "building", Value: id, Compare: "="},
	})

	var total float64 = 0
	for _, v := range items {
		if v.Category == 10 {
			total += float64(global.Atoi(v.Value2))
		}

		if v.Category == 20 {
			total += float64(global.Atoi(v.Value3))
		}

		if v.Category == 30 {
			total += float64(global.Atoi(v.Value6))
		}
	}

	buildingManager.UpdateTotalweight(models.Double(total), id)
}

func (c *FacilityController) Post_Insert(item *models.Facility) {
	CalculateTotalweight(item.Id)
}

func (c *FacilityController) Post_Update(item *models.Facility) {
	CalculateTotalweight(item.Id)
}

func (c *FacilityController) Post_Delete(item *models.Facility) {
	CalculateTotalweight(item.Id)
}
