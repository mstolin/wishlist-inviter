package gmail

import (
	"encoding/json"
	"fmt"

	"github.com/mstolin/present-roulette/utils/models"
)

// Sends POST request to send an invitation mail.
func (gClient GMailClient) SendInvitation(invitation models.Mail) (models.Mail, error) {
	gmailResp := models.Mail{}

	url := fmt.Sprintf("%s/mail", gClient.URL)
	jsonStr, err := json.Marshal(fmt.Sprintf(`{"recipient":"%s","subject":"%s","body":"%s"}`, invitation.Recipient, invitation.Subject, invitation.Body))
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
