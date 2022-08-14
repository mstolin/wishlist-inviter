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
func (client DatabaseClient) GetItemsForUser(userId string, wantedIds []uint, accessToken string) (models.ItemList, *httpErrors.ErrorResponse) {
	items := models.ItemList{}

	url := fmt.Sprintf("%s/users/%s/items", client.URL, userId)
	res, httpErr := client.httpFacade.DoGet(url, accessToken)
	if httpErr != nil {
		return items, httpErr
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return items, httpErrors.ErrServerErrorRenderer(err)
	}

	wantedItems := filterItems(items, wantedIds)
	if len(wantedItems) <= 0 {
		return items, &httpErrors.ErrNotFound
	}

	return wantedItems, nil
}

// Filters an array of items based on their IDs
func filterItems(items []models.Item, wantedIds []uint) []models.Item {
	filteredItems := []models.Item{}
	for _, item := range items {
		if contains(item.ID, wantedIds) {
			filteredItems = append(filteredItems, item)
		}
	}
	return filteredItems
}

// Check if the wanted number is in the given array.
func contains(search uint, array []uint) bool {
	for _, id := range array {
		if search == id {
			return true
		}
	}
	return false
}
