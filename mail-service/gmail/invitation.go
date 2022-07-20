package gmail

import (
	"encoding/json"
	"fmt"

	"github.com/mstolin/present-roulette/utils/httpErrors"
	"github.com/mstolin/present-roulette/utils/models"
)

// Sends POST request to send an invitation mail.
func (gClient GMailClient) SendInvitation(invitation models.Mail) (models.Mail, *httpErrors.ErrorResponse) {
	gmailResp := models.Mail{}

	url := fmt.Sprintf("%s/mail", gClient.URL)
	jsonData, err := json.Marshal(invitation)

	if err != nil {
		return gmailResp, httpErrors.ErrServerErrorRenderer(err)
	}

	res, httpErr := gClient.httpFacade.DoPost(url, jsonData)
	if httpErr != nil {
		return gmailResp, httpErr
	}

	if err := json.Unmarshal(res, &gmailResp); err != nil {
		return gmailResp, httpErrors.ErrServerErrorRenderer(err)
	}
	return gmailResp, nil
}
