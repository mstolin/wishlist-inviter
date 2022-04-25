package messages

import (
	"fmt"
	"strings"

	"github.com/mstolin/present-roulette/utils/models"
)

const invitationMsg = "Hi,\n" +
	"you have been invited to buy the following items:\n" +
	"%s\n" +
	"Cheers!"

func (factory MessageFactory) GenInvitationMsg(from string, to string, subject string, items []models.Item) string {
	// 1. Generate text
	txt := factory.genInvitationTxt(items)
	// 2. Generate mail content
	return fmt.Sprintf(mailBody, to, from, subject, txt)
}

// Generates an invitation message for multiple items.
func (factory MessageFactory) genInvitationTxt(items []models.Item) string {
	itemTxts := getItemTexts(items)
	itemStr := strings.Join(itemTxts, "\n")
	return fmt.Sprint(invitationMsg, itemStr)
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
