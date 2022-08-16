package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	"github.com/mstolin/present-roulette/database-adapter/db"
	"github.com/mstolin/present-roulette/utils/httpErrors"
)

var tokenAuth *jwtauth.JWTAuth
var dbHandler db.DatabaseHandler

func NewHandler(signKey string, db db.DatabaseHandler) http.Handler {
	tokenAuth = jwtauth.New("HS256", []byte(signKey), nil)
	dbHandler = db
	return newRouter()
}

func newRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.NotFound(notFoundHandler)
	r.MethodNotAllowed(methodNotAllowedHandler)
	r.Route("/users", userHandler)

	return r
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	render.Render(w, r, &httpErrors.ErrNotFound)
}

func methodNotAllowedHandler(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(405)
	render.Render(writer, request, &httpErrors.ErrMethodNotAllowed)
}
