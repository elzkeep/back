package api

import (
	"aoi/controllers"
	"aoi/global"
)

type SmsController struct {
	controllers.Controller
}

func (c *SmsController) Index(to string, message string) {
	global.SendSMS(to, message)
}
