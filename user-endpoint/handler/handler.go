package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/mstolin/present-roulette/user-endpoint/clients"
	"github.com/mstolin/present-roulette/utils/errors"
)

var wishlistClientInstance clients.WishlistClient
var mailClientInstance clients.MailClient
var userClientInstance clients.UserClient

func NewHandler(userClient clients.UserClient, mailClient clients.MailClient, wishlistClient clients.WishlistClient) http.Handler {
	wishlistClientInstance = wishlistClient
	mailClientInstance = mailClient
	userClientInstance = userClient

	return newRouter()
}

func newRouter() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(render.SetContentType(render.ContentTypeJSON))
	router.MethodNotAllowed(methodNotAllowedHandler)
	router.NotFound(notFoundHandler)
	router.Route("/users", userHandler)
	router.Route("/items", itemHandler)
	router.Route("/mail", mailHandler)
	return router
}

func methodNotAllowedHandler(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(405)
	render.Render(writer, request, errors.ErrMethodNotAllowed)
}

func notFoundHandler(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(400)
	render.Render(writer, request, errors.ErrNotFound)
}
