package api

import (
	"zkeep/controllers"
	"zkeep/global"
)

type SmsController struct {
	controllers.Controller
}

func (c *SmsController) Index(to string, message string) {
	global.SendSMS(to, message)
}
