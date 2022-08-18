package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	"github.com/mstolin/wishlist-inviter/scrapper-facade/scrapper"
	"github.com/mstolin/wishlist-inviter/utils/httpErrors"
	"github.com/mstolin/wishlist-inviter/utils/httpMiddleware"
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
		r.Use(httpMiddleware.JSONAuthenticator)
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
