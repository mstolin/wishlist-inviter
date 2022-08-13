package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/mstolin/present-roulette/user-service/clients"
	"github.com/mstolin/present-roulette/utils/httpErrors"
)

var dbClientInstance clients.DatabaseClient

func NewHandler(dbClient clients.DatabaseClient) http.Handler {
	dbClientInstance = dbClient
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(render.SetContentType(render.ContentTypeJSON))
	router.MethodNotAllowed(methodNotAllowedHandler)
	router.NotFound(notFoundHandler)
	router.Route("/users", userHandler)
	return router
}

func methodNotAllowedHandler(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(405)
	render.Render(writer, request, &httpErrors.ErrMethodNotAllowed)
}

func notFoundHandler(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(400)
	render.Render(writer, request, &httpErrors.ErrNotFound)
}
