package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mstolin/present-roulette/user-endpoint/models"
)

type MailClient struct {
	URL string
}

func NewMailClient(url string) (MailClient, error) {
	client := MailClient{}
	if url == "" {
		return client, fmt.Errorf("mail service URL is not defined or empty")
	} else {
		client.URL = url
	}

	return client, nil
}

func (client MailClient) sendInvitation(invitation *models.Invitation) (*models.Invitation, error) {
	result := &models.Invitation{}

	url := client.URL + "/mail/send"
	jsonStr := []byte(fmt.Sprintf(`{"recipient":"%s","subject":"%s","message":"%s"}`, invitation.Recipient, invitation.Subject, invitation.Message))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return result, err
	}
	req.Header.Set("Content-Type", "application/json")

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	if err := json.Unmarshal(body, result); err != nil {
		return result, err
	}

	return result, nil
}
