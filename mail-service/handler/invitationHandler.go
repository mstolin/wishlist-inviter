package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/mstolin/present-roulette/utils/errors"
	"github.com/mstolin/present-roulette/utils/models"
)

func invitationHandler(r chi.Router) {
	r.Post("/send", sendInvitation)
}

// Send invitation endpoint
func sendInvitation(w http.ResponseWriter, r *http.Request) {
	invitationReq := models.InvitationReq{}
	if error := render.Bind(r, &invitationReq); error != nil {
		render.Render(w, r, errors.ErrBadRequest)
		return
	}

	// Get all items
	items, err := getItemsForUser(invitationReq.UserId, invitationReq.Items)
	if err != nil {
		render.Render(w, r, errors.ErrorRenderer(err))
		return
	}

	// Generate invitation
	invitation := generateInvitation(invitationReq.Subject, invitationReq.Recipient, items)

	// Send Invitation
	gmailResp, err := gmailClientInstance.SendInvitation(invitation)
	if err != nil {
		render.Render(w, r, errors.ErrorRenderer(err))
		return
	}

	if err := render.Render(w, r, &gmailResp); err != nil {
		render.Render(w, r, errors.ServerErrorRenderer(err))
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

// Creates an instance if Invitation.
func generateInvitation(subject string, recipient string, items []models.Item) models.Invitation {
	message := msgFactoryInstance.GenerateInvitationMessage(items)
	return models.Invitation{
		Subject:   subject,
		Recipient: recipient,
		Message:   message,
	}
}
