package database

import (
	"encoding/json"
	"fmt"

	"github.com/mstolin/present-roulette/utils/clients"
	"github.com/mstolin/present-roulette/utils/httpErrors"
	"github.com/mstolin/present-roulette/utils/models"
)

type DatabaseClient struct {
	URL        string
	httpFacade clients.HTTPFacade
}

// Construct a new DatabaseClient instance
func NewDatabaseClient(url string) (DatabaseClient, error) {
	client := DatabaseClient{}
	if url == "" {
		return client, fmt.Errorf("database service URL can't be empty")
	}
	client.URL = url
	client.httpFacade = clients.NewHTTPFacade()
	return client, nil
}

// Requests all items of a specific user
func (client DatabaseClient) GetItemsForUser(userId string) (models.ItemList, *httpErrors.ErrorResponse) {
	items := models.ItemList{}

	url := fmt.Sprintf("%s/users/%s/items", client.URL, userId)
	res, err := client.httpFacade.DoGet(url)
	if err != nil {
		return items, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return items, httpErrors.ErrServerErrorRenderer(err)
	}
	return items, nil
}
