package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/mstolin/present-roulette/user-endpoint/clients"
	"github.com/mstolin/present-roulette/utils/httpErrors"
)

var scrapperFacadeInstance clients.ScrapperFacadeClient
var mailClientInstance clients.MailClient
var userClientInstance clients.UserClient

func NewHandler(userClient clients.UserClient, mailClient clients.MailClient, scrapperFacade clients.ScrapperFacadeClient) http.Handler {
	scrapperFacadeInstance = scrapperFacade
	mailClientInstance = mailClient
	userClientInstance = userClient

	return newRouter()
}

func newRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	r.MethodNotAllowed(methodNotAllowedHandler)
	r.NotFound(notFoundHandler)

	r.Route("/users", userHandler)
	r.Route("/items", itemHandler)
	r.Route("/mail", mailHandler)

	return r
}

func methodNotAllowedHandler(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(405)
	render.Render(writer, request, &httpErrors.ErrMethodNotAllowed)
}

func notFoundHandler(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(400)
	render.Render(writer, request, &httpErrors.ErrNotFound)
}
