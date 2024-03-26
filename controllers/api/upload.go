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
	path := c.Get("path")
	originalfilename, filename := c.GetUpload(path, "file")

	log.Println(filename)
	c.Set("filename", filename)
	c.Set("originalfilename", originalfilename)
}
