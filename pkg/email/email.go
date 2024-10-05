package email

import (
	"net/smtp"
)

type Email struct {
	auth     smtp.Auth
	username string
	host     string
}

func NewEmail(username, password, host string) *Email {
	auth := smtp.PlainAuth("", username, password, host)
	return &Email{
		auth:     auth,
		username: username,
		host:     host,
	}
}

func (e *Email) Send(to []string, msg []byte) error {
	return smtp.SendMail(e.host+":587", e.auth, e.username, to, msg)
}
