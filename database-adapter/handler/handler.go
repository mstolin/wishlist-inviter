package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/mstolin/present-roulette/database-adapter/db"
)

var dbHandler db.DatabaseHandler

func NewHandler(db db.DatabaseHandler) http.Handler {
	dbHandler = db

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.MethodNotAllowed(methodNotAllowedHandler)
	r.NotFound(notFoundHandler)
	//r.Route("/wishlist", wishlistHandler)
	r.Route("/users", userHandler)
	//r.Route("/items", itemHandler)

	/*
		NUR GET
		/users/USER_ID/items/

		POST GET PUT DELETE
		/users/USER_ID/items/ITEM_ID
	*/

	return r
}

func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(405)
	render.Render(w, r, ErrMethodNotAllowed)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(400)
	render.Render(w, r, ErrNotFound)
}
