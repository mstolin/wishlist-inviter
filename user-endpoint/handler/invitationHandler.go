package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/mstolin/present-roulette/user-endpoint/models"
)

const ITEM_ID_KEY = "itemId"

func invitationHandler(r chi.Router) {
	r.Route("/{itemId}", func(r chi.Router) {
		r.Use(itemCtx)
		r.Post("/", sendInvitation)
	})
}

func itemCtx(nxt http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		itemId := chi.URLParam(r, ITEM_ID_KEY)
		if itemId == "" {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("item ID is required")))
			return
		}

		// convert id to int
		id, err := strconv.Atoi(itemId)
		if err != nil {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("invalid item ID")))
		}

		ctx := context.WithValue(r.Context(), ITEM_ID_KEY, id)
		nxt.ServeHTTP(w, r.WithContext(ctx))
	})
}

func sendInvitation(w http.ResponseWriter, r *http.Request) {
	invitationReq := models.InvitationReq{}
	if err := render.Bind(r, &invitationReq); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	itemId := r.Context().Value(ITEM_ID_KEY).(int)
	resp, err := mailClientInstance.SendInvitation(invitationReq, itemId)
	if err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}

	if err := render.Render(w, r, &resp); err != nil { // TODO better success response
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}
