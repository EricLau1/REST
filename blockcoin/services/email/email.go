package email

import (
	"net/smtp"
)

type Email struct {
	To 		string
	Subject string
	Message string
}

const (
	host   = "smtp.gmail.com"
	port   = "587"
	sender = "eric.devtt@gmail.com"
	pass   = "PASSWORD EMAIL SENDER"
)

var auth smtp.Auth

func InitializeEmailAuth() {
	auth = smtp.PlainAuth("", sender, pass, host)
}

func parse(email Email) []byte {
	return []byte("To: " + email.To + "\r\n" +
				  "Subject: " + email.Subject + "\r\n" + 
				  "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n" + 
				  "\n" + email.Message)
}

func Send(email Email) error {
	receiver := email.To
	msg := parse(email)
	err := smtp.SendMail(
		host + ":" + port,
		auth,
		sender,
		[]string{receiver},
		msg)
	if err != nil {
		return err
	}
	return nil
}