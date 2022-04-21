package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/mstolin/present-roulette/mail-adapter/models"
	"github.com/mstolin/present-roulette/utils/errors"
)

func send(router chi.Router) {
	router.Post("/invitation", invitationHandler)
}

func invitationHandler(writer http.ResponseWriter, request *http.Request) {
	invitation := &models.Invitation{}

	if error := render.Bind(request, invitation); error != nil {
		render.Render(writer, request, errors.ErrBadRequest)
		return
	}

	gmailResp, error := gmailClientInstance.PostInvitation(invitation)
	if error != nil {
		render.Render(writer, request, errors.ErrorRenderer(error))
		return
	}

	if error := render.Render(writer, request, gmailResp); error != nil {
		render.Render(writer, request, errors.ServerErrorRenderer(error))
		return
	}
}
