package messages

import (
	"fmt"
	"strings"

	"github.com/mstolin/present-roulette/utils/models"
)

type MessageFactory struct {
	DatabaseService string
}

const message = "Hi,\n" +
	"you have been invited to buy the following items:\n" +
	"%s\n" +
	"Cheers!"

// Generates an invitation message for multiple items.
func (factory MessageFactory) GenerateInvitationMessage(items []models.Item) string {
	texts := getItemTexts(items)
	itemStr := strings.Join(texts, "\n")
	return fmt.Sprintf(message, itemStr)
}

// Returns an array of Item string representations.
func getItemTexts(items []models.Item) []string {
	texts := []string{}
	for _, item := range items {
		texts = append(texts, generateTextFromitem(item))
	}
	return texts
}

// Returns a string representation of an Item.
func generateTextFromitem(item models.Item) string {
	return fmt.Sprintf("%s, %.2f€", item.Name, item.Price)
}
