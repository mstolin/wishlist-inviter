package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/mstolin/present-roulette/utils/errors"
	"github.com/mstolin/present-roulette/utils/models"
)

func invitationHandler(r chi.Router) {
	r.Post("/", sendInvitation)
}

// Send invitation endpoint
func sendInvitation(w http.ResponseWriter, r *http.Request) {
	invitation := models.Invitation{}
	if error := render.Bind(r, &invitation); error != nil {
		render.Render(w, r, &errors.ErrBadRequest)
		return
	}

	// Get all items
	items, err := getItemsForUser(invitation.UserId, invitation.Items)
	if err != nil {
		render.Render(w, r, errors.ErrBadRequestRenderer(err))
		return
	}

	// Generate invitation
	invitationMail := genInvitationMail(invitation.Subject, invitation.Recipient, items)

	// Send Invitation
	resp, err := gmailClientInstance.SendInvitation(invitationMail)
	if err != nil {
		render.Render(w, r, errors.ErrBadRequestRenderer(err))
		return
	}

	if err := render.Render(w, r, &resp); err != nil {
		render.Render(w, r, errors.ErrServerErrorRenderer(err))
		return
	}
}

// Returns an array of all requested user items
func getItemsForUser(userId string, wantedIds []uint) ([]models.Item, error) {
	itemLst, err := dbClientInstance.GetItemsForUser(userId)
	if err != nil {
		return []models.Item{}, err
	}
	return filterItems(itemLst.Items, wantedIds), nil
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

// Creates an instance of an Invitation.
func genInvitationMail(subject string, recipient string, items []models.Item) models.Mail {
	message := msgFactoryInstance.GenInvitationMsg("", recipient, subject, items)
	return models.Mail{
		Subject:   subject,
		Recipient: recipient,
		Body:      message,
	}
}
