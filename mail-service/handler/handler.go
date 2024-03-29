package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	"github.com/mstolin/wishlist-inviter/mail-service/database"
	"github.com/mstolin/wishlist-inviter/mail-service/gmail"
	"github.com/mstolin/wishlist-inviter/mail-service/messages"
	"github.com/mstolin/wishlist-inviter/utils/httpErrors"
	"github.com/mstolin/wishlist-inviter/utils/httpMiddleware"
)

var tokenAuth *jwtauth.JWTAuth
var gmailClientInstance gmail.GMailClient
var dbClientInstance database.DatabaseClient
var msgFactoryInstance messages.MessageFactory

func NewHandler(signKey string, gmailClient gmail.GMailClient, dbClient database.DatabaseClient, msgFactory messages.MessageFactory) http.Handler {
	tokenAuth = jwtauth.New("HS256", []byte(signKey), nil)
	gmailClientInstance = gmailClient
	dbClientInstance = dbClient
	msgFactoryInstance = msgFactory
	return newRouter()
}

func newRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.MethodNotAllowed(methodNotAllowedHandler)
	r.NotFound(notFoundHandler)
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(httpMiddleware.JSONAuthenticator)
		r.Route("/invitations", invitationHandler)
	})
	return r
}

func notFoundHandler(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(404)
	render.Render(writer, request, &httpErrors.ErrNotFound)
}

func methodNotAllowedHandler(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(405)
	render.Render(writer, request, &httpErrors.ErrMethodNotAllowed)
}
