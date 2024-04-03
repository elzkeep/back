package api

import (
	"log"
	"zkeep/controllers"
	"zkeep/global"
)

type SmsController struct {
	controllers.Controller
}

// @POST()
func (c *SmsController) Index(to string, message string) {
	log.Println("SmsController.Index")
	global.SendSMS(to, message)
}
