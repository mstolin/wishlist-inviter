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

func (client MailClient) SendInvitation(invitation models.InvitationReq, itemId int) (models.InvitationRes, error) {
	result := models.InvitationRes{}

	jsonStr, err := json.Marshal(invitation)
	if err != nil {
		return result, err
	}

	url := fmt.Sprintf("%s/mail/send/invitation/%d", client.URL, itemId)
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

	if err := json.Unmarshal(body, &result); err != nil {
		return result, err
	}

	return result, nil
}
