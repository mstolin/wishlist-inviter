package mail

import (
	"errors"
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

func NewSMTPClient(host, port, sender, password string) (SMTPClient, error) {
	client := SMTPClient{}

	if host == "" || sender == "" {
		return client, errors.New("Host and sender can`t be empty")
	} else {
		client.Host = host
		client.Port = port
		client.Credentials = SMTPAuthCredentials{Sender: sender, Password: password}
		return client, nil
	}
}
