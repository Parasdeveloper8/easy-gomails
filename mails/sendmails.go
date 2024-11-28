package easymails

import (
	"log"

	"gopkg.in/gomail.v2"
)

func SendMail(senderMail, emailPassword, receivermail, subject, body string) error {
	from := senderMail
	password := emailPassword
	smtpHost := "smtp.gmail.com"
	smtpPort := 587

	m := gomail.NewMessage()
	m.SetHeader("From", senderMail)
	m.SetHeader("To", receivermail)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer(smtpHost, smtpPort, from, password)
	if err := d.DialAndSend(m); err != nil {
		log.Fatal(err)
	}
	return d.DialAndSend(m)
}
