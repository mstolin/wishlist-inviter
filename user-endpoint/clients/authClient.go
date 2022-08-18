package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mstolin/wishlist-inviter/utils/httpErrors"
	"github.com/mstolin/wishlist-inviter/utils/models"
)

type AuthClient struct {
	URL string
}

func NewAuthClient(url string) (AuthClient, error) {
	client := AuthClient{}
	if url == "" {
		return client, fmt.Errorf("auth service URL can't be empty")
	} else {
		client.URL = url
		return client, nil
	}
}

// Sends a request to the authenticate endpoint
func (client AuthClient) Authenticate(authObj models.AuthObj) (models.AccessToken, *httpErrors.ErrorResponse) {
	var accessToken models.AccessToken

	jsonStr, err := json.Marshal(authObj)
	if err != nil {
		return accessToken, httpErrors.ErrServerErrorRenderer(err)
	}

	url := fmt.Sprintf("%s/auth", client.URL)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		return accessToken, httpErrors.ErrServerErrorRenderer(err)
	}
	defer resp.Body.Close()

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return accessToken, httpErrors.ErrServerErrorRenderer(err)
	}

	// A non 2xx status does not cause an error
	if resp.StatusCode != http.StatusOK {
		errResp := httpErrors.ErrorResponse{}
		if err := json.Unmarshal(res, &errResp); err != nil {
			return accessToken, httpErrors.ErrServerErrorRenderer(err)
		}
		return accessToken, &errResp
	}

	if err := json.Unmarshal(res, &accessToken); err != nil {
		return accessToken, httpErrors.ErrServerErrorRenderer(err)
	}

	return accessToken, nil
}
