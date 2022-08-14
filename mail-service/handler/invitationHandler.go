package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/mstolin/present-roulette/utils/httpErrors"
	"github.com/mstolin/present-roulette/utils/models"
)

func invitationHandler(r chi.Router) {
	r.Post("/", sendInvitation)
}

// Send invitation endpoint
func sendInvitation(w http.ResponseWriter, r *http.Request) {
	invitation := models.Invitation{}
	if err := render.Bind(r, &invitation); err != nil {
		render.Render(w, r, httpErrors.ErrBadRequestRenderer(err))
		return
	}

	accessToken := r.Header.Get("Authorization")

	// Get all items
	items, httpErr := dbClientInstance.GetItemsForUser(invitation.UserId, invitation.Items, accessToken)
	if httpErr != nil {
		render.Render(w, r, httpErr)
		return
	}

	// Generate invitation
	invitationMail := msgFactoryInstance.GenInvitationMail(invitation.Recipient, items)

	// Send Invitation
	resp, httpErr := gmailClientInstance.SendInvitation(invitationMail, accessToken)
	if httpErr != nil {
		render.Render(w, r, httpErr)
		return
	}

	if err := render.Render(w, r, &resp); err != nil {
		render.Render(w, r, httpErrors.ErrServerErrorRenderer(err))
		return
	}
}
