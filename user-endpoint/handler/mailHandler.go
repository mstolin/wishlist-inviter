package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/mstolin/present-roulette/utils/errors"
	"github.com/mstolin/present-roulette/utils/models"
)

const ITEM_ID_KEY = "itemId"

func mailHandler(r chi.Router) {
	r.Route("/invitations", func(r chi.Router) {
		r.Post("/", sendInvitation)
	})
}

func sendInvitation(w http.ResponseWriter, r *http.Request) {
	invitationReq := models.Invitation{}
	if err := render.Bind(r, &invitationReq); err != nil {
		render.Render(w, r, errors.ErrBadRequest)
		return
	}

	resp, err := mailClientInstance.SendInvitation(invitationReq)
	if err != nil {
		render.Render(w, r, errors.ServerErrorRenderer(err))
		return
	}

	if err := render.Render(w, r, &resp); err != nil { // TODO better success response
		render.Render(w, r, errors.ServerErrorRenderer(err))
		return
	}
}
