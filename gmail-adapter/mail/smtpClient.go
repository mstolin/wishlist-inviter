package mail

import (
	"errors"
	"fmt"
	"net/smtp"

	"github.com/mstolin/present-roulette/utils/models"
)

type SMTPAuthCredentials struct {
	Sender   string
	Password string
}

type SMTPClient struct {
	Host        string
	Port        string
	Credentials SMTPAuthCredentials
}

// Construct a new SMTP client
func NewSMTPClient(host, port, sender, password string) (SMTPClient, error) {
	client := SMTPClient{}

	if host == "" || sender == "" {
		return client, errors.New("host and sender can`t be empty")
	} else {
		client.Host = host
		client.Port = port
		client.Credentials = SMTPAuthCredentials{Sender: sender, Password: password}
		return client, nil
	}
}

// Sends a mail using the SMTP protocol
func (client SMTPClient) SendMail(mail models.Mail) error {
	server := fmt.Sprintf("%s:%s", client.Host, client.Port)
	auth := smtp.PlainAuth("", client.Credentials.Sender, client.Credentials.Password, client.Host)
	body := []byte(mail.Body)

	err := smtp.SendMail(server, auth, client.Credentials.Sender, []string{mail.Recipient}, body)
	if err != nil {
		return err
	}
	return nil
}
