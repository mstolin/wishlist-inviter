package gmail

import (
	"errors"
)

type GMailClient struct {
	Host string
	Port string
}

func NewGMailClient(host, port string) (GMailClient, error) {
	client := GMailClient{}

	if host == "" || port == "" {
		return client, errors.New("Host or port can`t be empty")
	} else {
		client.Host = host
		client.Port = port
		return client, nil
	}
}
