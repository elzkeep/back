package api

import (
	"zkeep/controllers"
)

type UploadController struct {
	controllers.Controller
}

// @POST()
func (c *UploadController) Index() {
	path := c.Get("path")
	originalfilename, filename := c.GetUpload(path, "file")

	c.Set("filename", filename)
	c.Set("originalfilename", originalfilename)
}
