package mail

import (
	"net/smtp"

	"github.com/mstolin/present-roulette/gmail-adapter/models"
)

func (smtpClient SMTPClient) SendMail(mail *models.Mail) error {
	sender := smtpClient.Credentials.Sender
	password := smtpClient.Credentials.Password
	host := smtpClient.Host
	port := smtpClient.Port
	recipient := mail.Recipient
	subject := mail.Subject
	text := mail.Message
	auth := smtp.PlainAuth("", sender, password, host)
	message := []byte(GenerateMessage(sender, recipient, subject, text))

	if error := smtp.SendMail(host+":"+port, auth, sender, []string{recipient}, message); error != nil {
		return error
	}
	return nil
}
