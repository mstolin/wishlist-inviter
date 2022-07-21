package clients

import (
	"encoding/json"
	"fmt"

	"github.com/mstolin/present-roulette/utils/clients"
	"github.com/mstolin/present-roulette/utils/httpErrors"
	"github.com/mstolin/present-roulette/utils/models"
)

type MailClient struct {
	URL        string
	httpFacade clients.HTTPFacade
}

func NewMailClient(url string) (MailClient, error) {
	client := MailClient{}
	if url == "" {
		return client, fmt.Errorf("mail service URL is not defined or empty")
	} else {
		client.URL = url
	}

	client.httpFacade = clients.NewHTTPFacade()
	return client, nil
}

// Sends an request to the mail service to send an invitation
func (client MailClient) SendInvitation(invitation models.Invitation) (models.Mail, *httpErrors.ErrorResponse) {
	mailRes := models.Mail{}
	jsonStr, err := json.Marshal(invitation)
	if err != nil {
		return mailRes, httpErrors.ErrBadRequestRenderer(err)
	}

	url := fmt.Sprintf("%s/invitations", client.URL)
	res, httpErr := client.httpFacade.DoPost(url, jsonStr)
	if httpErr != nil {
		return mailRes, httpErr
	}

	if err := json.Unmarshal(res, &mailRes); err != nil {
		return mailRes, httpErrors.ErrBadRequestRenderer(err)
	}
	return mailRes, nil
}
