package core

import (
	model "MailingReports/MailingReports/Resources/Utilities/Models"
	"net/smtp"
	"os"
)

var (
	host     = os.Getenv("EMAIL_HOST")
	username = os.Getenv("EMAiL_USERNAME")
	password = os.Getenv("EMAIL_PASSWORD")
	//portNumber = os.Getenv("EMAIL_PORT")
)

func New() *model.Sender {
	auth := smtp.PlainAuth("", username, password, host)
	return &model.Sender{auth}
}

func NewMessage(s, b string) *model.Message {
	return &model.Message{Subject: s, Body: b, Attachments: make(map[string][]byte)}
}
