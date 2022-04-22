package gmail

import (
	"fmt"

	"github.com/mstolin/present-roulette/utils/clients"
)

type GMailClient struct {
	URL        string
	httpFacade clients.HTTPFacade
}

func NewGMailClient(url string) (GMailClient, error) {
	client := GMailClient{}

	if url == "" {
		return client, fmt.Errorf("service url can't be empty")
	}

	client.URL = url
	client.httpFacade = clients.NewHTTPFacade()
	return client, nil
}
