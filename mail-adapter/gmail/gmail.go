package gmail

import (
	"fmt"

	"github.com/mstolin/present-roulette/utils/clients"
)

type GMailClient struct {
	Host       string
	Port       string
	httpFacade clients.HTTPFacade
}

func NewGMailClient(host, port string) (GMailClient, error) {
	client := GMailClient{}

	if host == "" || port == "" {
		return client, fmt.Errorf("host or port can't be empty")
	}

	client.Host = host
	client.Port = port
	client.httpFacade = clients.NewHTTPFacade()
	return client, nil
}
