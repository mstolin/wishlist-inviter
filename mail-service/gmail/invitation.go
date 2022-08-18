package gmail

import (
	"encoding/json"
	"fmt"

	"github.com/mstolin/wishlist-inviter/utils/httpErrors"
	"github.com/mstolin/wishlist-inviter/utils/models"
)

// Sends POST request to send an invitation mail.
func (gClient GMailClient) SendInvitation(invitation models.Mail, accessToken string) (models.SuccessMessage, *httpErrors.ErrorResponse) {
	var successMsg models.SuccessMessage

	url := fmt.Sprintf("%s/mail", gClient.URL)
	jsonData, err := json.Marshal(invitation)
	if err != nil {
		return successMsg, httpErrors.ErrServerErrorRenderer(err)
	}

	res, httpErr := gClient.httpFacade.DoPost(url, accessToken, jsonData)
	if httpErr != nil {
		return successMsg, httpErr
	}

	if err := json.Unmarshal(res, &successMsg); err != nil {
		return successMsg, httpErrors.ErrServerErrorRenderer(err)
	}
	return successMsg, nil
}
