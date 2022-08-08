package handler

import (
	"net/http"

	"github.com/go-chi/chi"
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

	// Get all items
	items, httpErr := dbClientInstance.GetItemsForUser(invitation.UserId, invitation.Items)
	if httpErr != nil {
		render.Render(w, r, httpErr)
		return
	}

	// Generate invitation
	invitationMail := msgFactoryInstance.GenInvitationMail(invitation.Recipient, items)

	// Send Invitation
	resp, httpErr := gmailClientInstance.SendInvitation(invitationMail)
	if httpErr != nil {
		render.Render(w, r, httpErr)
		return
	}

	if err := render.Render(w, r, &resp); err != nil {
		render.Render(w, r, httpErrors.ErrServerErrorRenderer(err))
		return
	}
}
