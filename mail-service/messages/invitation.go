package messages

import (
	"fmt"
	"strings"

	"github.com/mstolin/present-roulette/utils/models"
)

const invitationMsg = "Hi,\nyou have been invited to buy the following items:\n\n%s\n\nCheers!"

const invitationSubject = "Someones sharing his wishlist"

func (factory MessageFactory) GenInvitationMail(from string, to string, items []models.Item) string {
	// 1. Generate text
	msg := factory.genInvitationMsg(items)
	// 2. Generate mail content
	content := fmt.Sprintf(mailBody, to, from, invitationSubject, msg)
	return content
}

// Generates an invitation message for multiple items.
func (factory MessageFactory) genInvitationMsg(items []models.Item) string {
	itemTxts := genItemTxt(items)
	content := fmt.Sprintf(invitationMsg, itemTxts)
	return content
}

// Returns an array of Item string representations.
func genItemTxt(items []models.Item) string {
	texts := []string{}
	for _, item := range items {
		texts = append(texts, generateTextFromItem(item))
	}
	return strings.Join(texts, "\n")
}

// Returns a string representation of an Item.
func generateTextFromItem(item models.Item) string {
	url := fmt.Sprintf("https://www.amazon.com/dp/%s/", item.VendorID)
	return fmt.Sprintf("  - %s, %.2f€ (%s)", item.Name, item.Price, url)
}
