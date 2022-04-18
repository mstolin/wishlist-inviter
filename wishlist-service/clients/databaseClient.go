package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mstolin/present-roulette/wishlist-service/models"
)

type DatabaseClient struct {
	URL string
}

func NewDatabaseClient(url string) (DatabaseClient, error) {
	client := DatabaseClient{}

	if url == "" {
		return client, fmt.Errorf("DATABASE_URL is undefined")
	} else {
		client.URL = url
		return client, nil
	}
}

func (client DatabaseClient) SaveItems(userId string, items []models.Item) (models.ItemList, error) {
	response := models.ItemList{}

	jsonStr, err := json.Marshal(items)
	if err != nil {
		return response, err
	}

	url := client.URL + "/users/" + userId + "/items"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return response, err
	}
	req.Header.Set("Content-Type", "application/json")

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return response, err
	}

	return response, nil
}
