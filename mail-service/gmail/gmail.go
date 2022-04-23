package gmail

import (
	"fmt"

	"github.com/mstolin/present-roulette/utils/clients"
)

type GMailClient struct {
	URL        string
	httpFacade clients.HTTPFacade
}

func NewGMailClient(gmailServiceUrl string) (GMailClient, error) {
	client := GMailClient{}
	if gmailServiceUrl == "" {
		return client, fmt.Errorf("gmail service url can't be empty")
	}
	client.URL = gmailServiceUrl
	client.httpFacade = clients.NewHTTPFacade()
	return client, nil
}
