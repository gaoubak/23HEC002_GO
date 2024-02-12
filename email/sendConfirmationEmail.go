package email

import (
	"strconv"

	"gopkg.in/gomail.v2"
)

// SendConfirmationEmail sends a confirmation email using the specified SMTP configuration
func SendConfirmationEmail(to, subject, body, smtpServer, smtpPort, smtpUser, smtpPassword string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", smtpUser)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	port, err := strconv.Atoi(smtpPort)
	if err != nil {
		return err
	}

	d := gomail.NewDialer(smtpServer, port, smtpUser, smtpPassword)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
