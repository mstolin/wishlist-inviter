package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	"github.com/mstolin/wishlist-inviter/gmail-adapter/mail"
	"github.com/mstolin/wishlist-inviter/utils/httpErrors"
	"github.com/mstolin/wishlist-inviter/utils/httpMiddleware"
)

var tokenAuth *jwtauth.JWTAuth
var smtpClientInstance mail.SMTPClient

func NewHandler(signKey string, smtpClient mail.SMTPClient) http.Handler {
	tokenAuth = jwtauth.New("HS256", []byte(signKey), nil)
	smtpClientInstance = smtpClient
	return newRouter()
}

func newRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.NotFound(notFoundHandler)
	r.MethodNotAllowed(methodNotAllowedHandler)
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(httpMiddleware.JSONAuthenticator)
		r.Route("/mail", mailHandler)
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
