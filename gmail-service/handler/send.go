package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/mstolin/present-roulette/gmail-service/models"
)

func send(router chi.Router) {
	router.Post("/", sendHandler)
}

func sendHandler(w http.ResponseWriter, r *http.Request) {
	mail := &models.Mail{}

	if error := render.Bind(r, mail); error != nil {
		fmt.Fprintf(os.Stderr, "Error: %q\n", error)
		render.Render(w, r, ErrBadRequest)
		return
	}

	if error := smtpClientInstance.SendMail(mail); error != nil {
		fmt.Fprintf(os.Stderr, "Error: %q\n", error)
		render.Render(w, r, ErrorRenderer(error))
		return
	}

	if error := render.Render(w, r, mail); error != nil {
		fmt.Fprintf(os.Stderr, "Error: %q\n", error)
		render.Render(w, r, ServerErrorRenderer(error))
	}
}
