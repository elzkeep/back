package api

import (
	"log"
	"net/smtp"
	"zkeep/config"
	"zkeep/controllers"
)

type MailController struct {
	controllers.Controller
}

// @POST()
func (c *MailController) Index(to string, subject string, body string) {
	auth := smtp.PlainAuth("", config.Mail.User, config.Mail.Password, "smtp.gmail.com")
	from := config.Mail.Sender

	msg := "MIME-Version: 1.0\n" +
		"Content-type: text/html; charset=utf-8\n" +
		"From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\r\n\r\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587", auth, from, []string{to}, []byte(msg))
	if err != nil {
		log.Println(err)
	}
}
