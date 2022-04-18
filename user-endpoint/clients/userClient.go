package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mstolin/present-roulette/user-endpoint/models"
)

type UserClient struct {
	URL string
}

func NewUserClient(url string) (UserClient, error) {
	client := UserClient{}
	if url == "" {
		return client, fmt.Errorf("user service URL is not defined or empty")
	} else {
		client.URL = url
	}

	return client, nil
}

func (client UserClient) createUser() (*models.User, error) {
	user := &models.User{}

	url := client.URL + "/mail/send"
	jsonStr := []byte("{}")
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return user, err
	}
	req.Header.Set("Content-Type", "application/json")

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return user, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return user, err
	}

	if err := json.Unmarshal(body, user); err != nil {
		return user, err
	}

	return user, nil
}
