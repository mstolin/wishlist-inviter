package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/mstolin/present-roulette/mail-adapter/gmail"
	"github.com/mstolin/present-roulette/utils/errors"
)

var gmailClientInstance gmail.GMailClient

func NewHandler(gmailClient gmail.GMailClient) http.Handler {
	gmailClientInstance = gmailClient
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.MethodNotAllowed(methodNotAllowedHandler)
	r.NotFound(notFoundHandler)
	r.Route("/invitation", invitationHandler)
	return r
}

func methodNotAllowedHandler(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(405)
	render.Render(writer, request, errors.ErrMethodNotAllowed)
}

func notFoundHandler(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(400)
	render.Render(writer, request, errors.ErrNotFound)
}
