package gmail

import (
	"encoding/json"
	"fmt"

	"github.com/mstolin/present-roulette/mail-adapter/models"
)

// Sends POST request to send an invitation mail.
func (gClient GMailClient) PostInvitation(invitation models.Invitation) (models.GMailResponse, error) {
	gmailResp := models.GMailResponse{}

	url := fmt.Sprintf("%s:%s/mail/send", gClient.Host, gClient.Port)
	jsonStr, err := json.Marshal(fmt.Sprintf(`{"recipient":"%s","subject":"%s","message":"%s"}`, invitation.Recipient, invitation.Subject, invitation.Message))
	if err != nil {
		return gmailResp, err
	}
	res, err := gClient.httpFacade.DoPost(url, jsonStr)
	if err != nil {
		return gmailResp, err
	}

	if err := json.Unmarshal(res, &gmailResp); err != nil {
		return gmailResp, err
	}
	return gmailResp, nil
}
