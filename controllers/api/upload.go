package api

import (
	"zkeep/controllers"
)

type UploadController struct {
	controllers.Controller
}

// @POST()
func (c *UploadController) Index() {
	_, filename := c.GetUpload("", "file")

	c.Set("filename", filename)
}
