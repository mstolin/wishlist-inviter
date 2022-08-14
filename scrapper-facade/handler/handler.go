package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	"github.com/mstolin/present-roulette/scrapper-facade/scrapper"
	"github.com/mstolin/present-roulette/utils/httpErrors"
)

var tokenAuth *jwtauth.JWTAuth
var scrapperFacadeInstance scrapper.ScrapperFacade

func NewHandler(signKey string, scrapperFacade scrapper.ScrapperFacade) http.Handler {
	tokenAuth = jwtauth.New("HS256", []byte(signKey), nil)
	scrapperFacadeInstance = scrapperFacade
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
		r.Use(jwtauth.Authenticator)
		r.Route("/amazon", amazonHandler)
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
