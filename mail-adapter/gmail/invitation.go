package gmail

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mstolin/present-roulette/mail-adapter/models"
)

func (gclient GMailClient) PostInvitation(invitation *models.Invitation) (*models.GMailResponse, error) {
	gmailResp := &models.GMailResponse{}

	url := gclient.Host + ":" + gclient.Port + "/mail/send"
	jsonStr := []byte(fmt.Sprintf(`{"recipient":"%s","subject":"%s","message":"%s"}`, invitation.Recipient, invitation.Subject, invitation.Message))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return gmailResp, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return gmailResp, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return gmailResp, err
	}

	if err := json.Unmarshal(body, gmailResp); err != nil {
		return gmailResp, err
	}

	return gmailResp, nil
}
