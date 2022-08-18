package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/mstolin/wishlist-inviter/utils/httpErrors"
	"github.com/mstolin/wishlist-inviter/utils/models"
)

func mailHandler(r chi.Router) {
	r.Route("/invitations", func(r chi.Router) {
		r.Post("/", sendInvitation)
	})
}

func sendInvitation(w http.ResponseWriter, r *http.Request) {
	invitationReq := models.Invitation{}
	if err := render.Bind(r, &invitationReq); err != nil {
		render.Render(w, r, httpErrors.ErrBadRequestRenderer(err))
		return
	}

	accessToken := r.Header.Get("Authorization")
	resp, err := mailClientInstance.SendInvitation(invitationReq, accessToken)
	if err != nil {
		render.Render(w, r, err)
		return
	}

	if err := render.Render(w, r, &resp); err != nil {
		render.Render(w, r, httpErrors.ErrServerErrorRenderer(err))
		return
	}
}
