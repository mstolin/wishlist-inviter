package clients

import (
	"encoding/json"
	"fmt"

	"github.com/mstolin/present-roulette/utils/clients"
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
func (client MailClient) SendInvitation(invitationReq models.InvitationReq) (models.Invitation, error) {
	invitationRes := models.Invitation{}
	jsonStr, err := json.Marshal(invitationReq)
	if err != nil {
		return invitationRes, err
	}

	url := fmt.Sprintf("%s/invitation", client.URL)
	res, err := client.httpFacade.DoPost(url, jsonStr)
	if err != nil {
		return invitationRes, err
	}

	if err := json.Unmarshal(res, &invitationRes); err != nil {
		return invitationRes, err
	}
	return invitationRes, nil
}
