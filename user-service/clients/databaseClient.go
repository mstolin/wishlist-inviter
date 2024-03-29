package clients

import (
	"encoding/json"
	"fmt"

	"github.com/mstolin/wishlist-inviter/utils/clients"
	"github.com/mstolin/wishlist-inviter/utils/httpErrors"
	"github.com/mstolin/wishlist-inviter/utils/models"
)

type DatabaseClient struct {
	URL        string
	httpFacade clients.HTTPFacade
}

// Constructs a new DatabaseClient instance.
func NewDatabaseClient(url string) (DatabaseClient, error) {
	client := DatabaseClient{}
	if url == "" {
		return client, fmt.Errorf("DATABASE_ADAPTER can't be empty")
	}

	client.URL = url
	client.httpFacade = clients.NewHTTPFacade()
	return client, nil
}

// Sends a POST request to create a new User without any items.
func (client DatabaseClient) CreateUser(accessToken string) (models.User, *httpErrors.ErrorResponse) {
	user := models.User{}

	url := fmt.Sprintf("%s/users", client.URL)
	res, httpErr := client.httpFacade.DoPost(url, accessToken, []byte{}) // send nothing
	if httpErr != nil {
		return user, httpErr
	}

	if err := json.Unmarshal(res, &user); err != nil {
		return user, httpErrors.ErrServerErrorRenderer(err)
	}

	return user, nil
}

// Sends a GET request to validate a user.
func (client DatabaseClient) VerifyUser(userId string) (models.UserVerification, *httpErrors.ErrorResponse) {
	var userVerification models.UserVerification

	url := fmt.Sprintf("%s/users/verify/%s", client.URL, userId)
	res, httpErr := client.httpFacade.DoGet(url, "")
	if httpErr != nil {
		return userVerification, httpErr
	}

	if err := json.Unmarshal(res, &userVerification); err != nil {
		return userVerification, httpErrors.ErrServerErrorRenderer(err)
	}
	return userVerification, nil
}

// Sends a GET request to retrieve a specific user given its ID.
func (client DatabaseClient) GetUser(userId, accessToken string) (models.User, *httpErrors.ErrorResponse) {
	user := models.User{}

	url := fmt.Sprintf("%s/users/%s", client.URL, userId)
	res, httpErr := client.httpFacade.DoGet(url, accessToken)
	if httpErr != nil {
		return user, httpErr
	}

	if err := json.Unmarshal(res, &user); err != nil {
		return user, httpErrors.ErrServerErrorRenderer(err)
	}
	return user, nil
}

// Sends a DELETE request to delete a specific user.
func (client DatabaseClient) DeleteUser(userId, accessToken string) (models.User, *httpErrors.ErrorResponse) {
	user := models.User{}

	url := fmt.Sprintf("%s/users/%s", client.URL, userId)
	res, httpErr := client.httpFacade.DoDelete(url, accessToken)
	if httpErr != nil {
		return user, httpErr
	}

	if err := json.Unmarshal(res, &user); err != nil {
		return user, httpErrors.ErrServerErrorRenderer(err)
	}
	return user, nil
}

// Sends a GET request to retrieve all items of a specific user.
func (client DatabaseClient) GetItemsByUser(userId, accessToken string) (models.ItemList, *httpErrors.ErrorResponse) {
	list := models.ItemList{}

	url := fmt.Sprintf("%s/users/%s/items", client.URL, userId)
	res, httpErr := client.httpFacade.DoGet(url, accessToken)
	if httpErr != nil {
		return list, httpErr
	}

	if err := json.Unmarshal(res, &list); err != nil {
		return list, httpErrors.ErrServerErrorRenderer(err)
	}
	return list, nil
}

// Sends a POST request to add given Items to a given User.
func (client DatabaseClient) AddItemsToUser(userId string, items models.ItemList, accessToken string) (models.ItemList, *httpErrors.ErrorResponse) {
	addedItems := models.ItemList{}
	jsonStr, err := json.Marshal(items)
	if err != nil {
		return addedItems, httpErrors.ErrServerErrorRenderer(err)
	}

	url := fmt.Sprintf("%s/users/%s/items", client.URL, userId)
	res, httpErr := client.httpFacade.DoPost(url, accessToken, jsonStr)
	if httpErr != nil {
		return addedItems, httpErr
	}

	if err := json.Unmarshal(res, &addedItems); err != nil {
		return addedItems, httpErrors.ErrServerErrorRenderer(err)
	}
	return addedItems, nil
}

// Sends a GET request to retrieve a specific Item of a User.
func (client DatabaseClient) GetItemByUser(userId string, itemId int, accessToken string) (models.Item, *httpErrors.ErrorResponse) {
	item := models.Item{}

	url := fmt.Sprintf("%s/users/%s/items/%d", client.URL, userId, itemId)
	res, httpErr := client.httpFacade.DoGet(url, accessToken)
	if httpErr != nil {
		return item, httpErr
	}

	if err := json.Unmarshal(res, &item); err != nil {
		return item, httpErrors.ErrServerErrorRenderer(err)
	}
	return item, nil
}

// Sends a PUT request to update a specific Item of a User.
func (client DatabaseClient) UpdateItemByUser(userId string, itemId int, update models.ItemUpdate, accessToken string) (models.Item, *httpErrors.ErrorResponse) {
	updatedItem := models.Item{}
	jsonStr, err := json.Marshal(update)
	if err != nil {
		return updatedItem, httpErrors.ErrServerErrorRenderer(err)
	}

	url := fmt.Sprintf("%s/users/%s/items/%d", client.URL, userId, itemId)
	res, httpErr := client.httpFacade.DoPut(url, accessToken, jsonStr)
	if httpErr != nil {
		return updatedItem, httpErr
	}

	if err := json.Unmarshal(res, &updatedItem); err != nil {
		return updatedItem, httpErrors.ErrServerErrorRenderer(err)
	}
	return updatedItem, nil
}

// Sends a DELETE request to delete a specific user.
func (client DatabaseClient) DeleteItemByUser(userId string, itemId int, accessToken string) (models.Item, *httpErrors.ErrorResponse) {
	deletedItem := models.Item{}

	url := fmt.Sprintf("%s/users/%s/items/%d", client.URL, userId, itemId)
	res, httpErr := client.httpFacade.DoDelete(url, accessToken)
	if httpErr != nil {
		return deletedItem, httpErr
	}

	if err := json.Unmarshal(res, &deletedItem); err != nil {
		return deletedItem, httpErrors.ErrServerErrorRenderer(err)
	}
	return deletedItem, nil
}
