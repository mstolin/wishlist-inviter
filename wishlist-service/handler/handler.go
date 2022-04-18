package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/mstolin/present-roulette/wishlist-service/clients"
)

var scrapperClientInstance clients.ScrapperClient
var dbClientInstance clients.DatabaseClient

func NewHandler(scrapperClient clients.ScrapperClient, dbClient clients.DatabaseClient) http.Handler {
	scrapperClientInstance = scrapperClient
	dbClientInstance = dbClient
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(render.SetContentType(render.ContentTypeJSON))
	router.MethodNotAllowed(methodNotAllowedHandler)
	router.NotFound(notFoundHandler)
	router.Route("/wishlist", wishlistHandler)
	return router
}

func methodNotAllowedHandler(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(405)
	render.Render(writer, request, ErrMethodNotAllowed)
}

func notFoundHandler(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(400)
	render.Render(writer, request, ErrNotFound)
}
