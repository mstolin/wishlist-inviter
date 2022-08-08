package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/mstolin/present-roulette/utils/httpErrors"
	"github.com/mstolin/present-roulette/utils/models"
)

func mailHandler(router chi.Router) {
	router.Post("/", send)
}

func send(w http.ResponseWriter, r *http.Request) {
	mail := models.Mail{}

	if err := render.Bind(r, &mail); err != nil {
		render.Render(w, r, httpErrors.ErrBadRequestRenderer(err))
		return
	}

	successMsg, err := smtpClientInstance.SendMail(mail)
	if err != nil {
		render.Render(w, r, httpErrors.ErrServerErrorRenderer(err))
		return
	}

	if err := render.Render(w, r, &successMsg); err != nil {
		render.Render(w, r, httpErrors.ErrServerErrorRenderer(err))
		return
	}
}
