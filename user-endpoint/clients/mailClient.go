package clients

import (
	"encoding/json"
	"fmt"

	"github.com/mstolin/wishlist-inviter/utils/clients"
	"github.com/mstolin/wishlist-inviter/utils/httpErrors"
	"github.com/mstolin/wishlist-inviter/utils/models"
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
func (client MailClient) SendInvitation(invitation models.Invitation, accessToken string) (models.SuccessMessage, *httpErrors.ErrorResponse) {
	mailRes := models.SuccessMessage{}
	jsonStr, err := json.Marshal(invitation)
	if err != nil {
		return mailRes, httpErrors.ErrBadRequestRenderer(err)
	}

	url := fmt.Sprintf("%s/invitations", client.URL)
	res, httpErr := client.httpFacade.DoPost(url, accessToken, jsonStr)
	if httpErr != nil {
		return mailRes, httpErr
	}

	if err := json.Unmarshal(res, &mailRes); err != nil {
		return mailRes, httpErrors.ErrBadRequestRenderer(err)
	}
	return mailRes, nil
}
