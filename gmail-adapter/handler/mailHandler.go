package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/mstolin/present-roulette/utils/errors"
	"github.com/mstolin/present-roulette/utils/models"
)

func mailHandler(router chi.Router) {
	router.Post("/", send)
}

func send(w http.ResponseWriter, r *http.Request) {
	mail := models.Mail{}

	if err := render.Bind(r, &mail); err != nil {
		render.Render(w, r, errors.ErrBadRequest)
		return
	}

	if err := smtpClientInstance.SendMail(mail); err != nil {
		render.Render(w, r, errors.ErrorRenderer(err))
		return
	}

	if err := render.Render(w, r, &mail); err != nil {
		render.Render(w, r, errors.ServerErrorRenderer(err))
		return
	}
}
