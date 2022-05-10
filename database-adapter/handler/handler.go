package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/mstolin/present-roulette/database-adapter/db"
	"github.com/mstolin/present-roulette/utils/httpErrors"
)

var dbHandler db.DatabaseHandler

func NewHandler(db db.DatabaseHandler) http.Handler {
	dbHandler = db

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.NotFound(notFoundHandler)
	r.Route("/users", userHandler)

	return r
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	render.Render(w, r, &httpErrors.ErrNotFound)
}
