package clients

import (
	"encoding/json"
	"fmt"

	"github.com/mstolin/present-roulette/user-service/models"
	"github.com/mstolin/present-roulette/utils/clients"
)

type DatabaseClient struct {
	URL        string
	httpFacade clients.HTTPFacade
}

// Constructs a new DatabaseClient instance.
func NewDatabaseClient(url string) (DatabaseClient, error) {
	client := DatabaseClient{}
	if url == "" {
		return client, fmt.Errorf("DATABASE_URL can't be mepty")
	}

	client.URL = url
	client.httpFacade = clients.NewHTTPFacade()
	return client, nil
}

// Sends a POST request to create a new User without any items.
func (client DatabaseClient) CreateUser() (models.User, error) {
	user := models.User{}

	jsonStr, err := json.Marshal(`{"items": []}`) // empty user with no items
	if err != nil {
		return user, err
	}

	url := fmt.Sprintf("%s/users", client.URL)
	res, err := client.httpFacade.DoPost(url, jsonStr)
	if err != nil {
		return user, err
	}

	if err := json.Unmarshal(res, &user); err != nil {
		return user, err
	}
	return user, nil
}

// Sends a GET request to retrieve a specific user given its ID.
func (client DatabaseClient) GetUser(userId string) (models.User, error) {
	user := models.User{}

	url := fmt.Sprintf("%s/users/%s", client.URL, userId)
	res, err := client.httpFacade.DoGet(url)
	if err != nil {
		return user, err
	}

	if err := json.Unmarshal(res, &user); err != nil {
		return user, err
	}
	return user, nil
}

// Sends a DELETE request to delete a specific user.
func (client DatabaseClient) DeleteUser(userId string) (models.User, error) {
	user := models.User{}

	url := fmt.Sprintf("%s/users/%s", client.URL, userId)
	res, err := client.httpFacade.DoDelete(url)
	if err != nil {
		return user, err
	}

	if err := json.Unmarshal(res, &user); err != nil {
		return user, err
	}
	return user, nil
}

// Sends a GET request to retrieve all items of a specific user.
func (client DatabaseClient) GetItemsByUser(userId string) (models.ItemList, error) {
	list := models.ItemList{}

	url := fmt.Sprintf("%s/users/%s/items", client.URL, userId)
	res, err := client.httpFacade.DoGet(url)
	if err != nil {
		return list, err
	}

	if err := json.Unmarshal(res, &list); err != nil {
		return list, err
	}
	return list, nil
}

// Sends a POST request to add given Items to a given User.
func (client DatabaseClient) AddItemsToUser(userId string, items models.ItemList) (models.ItemList, error) {
	addedItems := models.ItemList{}
	jsonStr, err := json.Marshal(items)
	if err != nil {
		return addedItems, err
	}

	url := fmt.Sprintf("%s/users/%s/items", client.URL, userId)
	res, err := client.httpFacade.DoPost(url, jsonStr)
	if err != nil {
		return addedItems, err
	}

	if err := json.Unmarshal(res, &addedItems); err != nil {
		return addedItems, err
	}
	return addedItems, nil
}

// Sends a GET request to retrieve a specific Item of a User.
func (client DatabaseClient) GetItemByUser(userId string, itemId int) (models.Item, error) {
	item := models.Item{}

	url := fmt.Sprintf("%s/users/%s/items/%d", client.URL, userId, itemId)
	res, err := client.httpFacade.DoGet(url)
	if err != nil {
		return item, err
	}

	if err := json.Unmarshal(res, &item); err != nil {
		return item, err
	}
	return item, nil
}

// Sends a PUT request to update a specific Item of a User.
func (client DatabaseClient) UpdateItemByUser(userId string, itemId int, update models.Item) (models.Item, error) {
	updatedItem := models.Item{}
	jsonStr, err := json.Marshal(update)
	if err != nil {
		return updatedItem, err
	}

	url := fmt.Sprintf("%s/users/%s/items/%d", client.URL, userId, itemId)
	res, err := client.httpFacade.DoPut(url, jsonStr)
	if err != nil {
		return updatedItem, err
	}

	if err := json.Unmarshal(res, &updatedItem); err != nil {
		return updatedItem, err
	}
	return updatedItem, nil
}

// Sends a DELETE request to delete a specific user.
func (client DatabaseClient) DeleteItemByUser(userId string, itemId int) (models.Item, error) {
	deletedItem := models.Item{}

	url := fmt.Sprintf("%s/users/%s/items/%d", client.URL, userId, itemId)
	res, err := client.httpFacade.DoDelete(url)
	if err != nil {
		return deletedItem, err
	}

	if err := json.Unmarshal(res, &deletedItem); err != nil {
		return deletedItem, err
	}
	return deletedItem, nil
}
