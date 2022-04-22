package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/mstolin/present-roulette/mail-service/models"
	"github.com/mstolin/present-roulette/utils/errors"
)

func invitationHandler(r chi.Router) {
	r.Post("/send", sendInvitation)
}

// Send invitation endpoint
func sendInvitation(w http.ResponseWriter, r *http.Request) {
	invitation := models.Invitation{}

	if error := render.Bind(r, &invitation); error != nil {
		render.Render(w, r, errors.ErrBadRequest)
		return
	}

	gmailResp, error := gmailClientInstance.PostInvitation(invitation)
	if error != nil {
		render.Render(w, r, errors.ErrorRenderer(error))
		return
	}

	if error := render.Render(w, r, &gmailResp); error != nil {
		render.Render(w, r, errors.ServerErrorRenderer(error))
		return
	}
}
