package api

import (
	"log"
	"zkeep/controllers"
)

type UploadController struct {
	controllers.Controller
}

// @POST()
func (c *UploadController) Index() {
	log.Println("UPLOAD ----------------")
	_, filename := c.GetUpload("", "")

	log.Println(filename)
	c.Set("filename", filename)
}
