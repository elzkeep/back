package api

import (
	"zkeep/controllers"
)

type DownloadController struct {
	controllers.Controller
}

func (c *DownloadController) File(id int64) {
	/*
		conn := c.NewConnection()

		manager := models.NewFileManager(conn)
		item := manager.Get(id)

		fullFilename := fmt.Sprintf("%v/%v", config.UploadPath, item.Filename)
		c.Download(fullFilename, item.Originalfilename)
	*/
}
