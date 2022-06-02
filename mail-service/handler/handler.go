package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/mstolin/present-roulette/mail-service/database"
	"github.com/mstolin/present-roulette/mail-service/gmail"
	"github.com/mstolin/present-roulette/mail-service/messages"
	"github.com/mstolin/present-roulette/utils/httpErrors"
)

var gmailClientInstance gmail.GMailClient
var dbClientInstance database.DatabaseClient
var msgFactoryInstance messages.MessageFactory

func NewHandler(gmailClient gmail.GMailClient, dbClient database.DatabaseClient, msgFactory messages.MessageFactory) http.Handler {
	gmailClientInstance = gmailClient
	dbClientInstance = dbClient
	msgFactoryInstance = msgFactory

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.NotFound(notFoundHandler)
	r.Route("/invitations", invitationHandler)
	return r
}

func notFoundHandler(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(404)
	render.Render(writer, request, &httpErrors.ErrNotFound)
}
