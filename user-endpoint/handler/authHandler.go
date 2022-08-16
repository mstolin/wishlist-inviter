package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/mstolin/present-roulette/utils/httpErrors"
	"github.com/mstolin/present-roulette/utils/models"
)

func authHandler(r chi.Router) {
	r.Post("/", auth)
}

func auth(w http.ResponseWriter, r *http.Request) {
	var authObj models.AuthObj
	if err := render.Bind(r, &authObj); err != nil {
		render.Render(w, r, httpErrors.ErrServerErrorRenderer(err))
		return
	}

	accessToken, httpErr := authClientInstance.Authenticate(authObj)
	if httpErr != nil {
		render.Render(w, r, httpErr)
		return
	}

	if err := render.Render(w, r, &accessToken); err != nil {
		render.Render(w, r, httpErrors.ErrServerErrorRenderer(err))
		return
	}
}
