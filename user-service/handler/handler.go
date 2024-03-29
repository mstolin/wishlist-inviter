package handler

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	"github.com/mstolin/wishlist-inviter/user-service/clients"
	"github.com/mstolin/wishlist-inviter/utils/httpErrors"
	"github.com/mstolin/wishlist-inviter/utils/models"
)

var tokenAuth *jwtauth.JWTAuth
var dbClientInstance clients.DatabaseClient

func NewHandler(signKey string, dbClient clients.DatabaseClient) http.Handler {
	tokenAuth = jwtauth.New("HS256", []byte(signKey), nil)
	dbClientInstance = dbClient
	return newRouter()
}

func newRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.MethodNotAllowed(methodNotAllowedHandler)
	r.NotFound(notFoundHandler)
	r.Route("/users", userHandler)
	r.Post("/auth", auth)
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

func auth(w http.ResponseWriter, r *http.Request) {
	var authObj models.AuthObj
	if err := render.Bind(r, &authObj); err != nil {
		render.Render(w, r, httpErrors.ErrServerErrorRenderer(err))
		return
	}

	// create claims/access token
	claims := map[string]interface{}{}
	jwtauth.SetIssuedNow(claims)
	jwtauth.SetExpiryIn(claims, 86400*time.Second) // Expires in one day
	_, accessToken, err := tokenAuth.Encode(claims)
	if err != nil {
		render.Render(w, r, httpErrors.ErrServerErrorRenderer(err))
		return
	}

	// check if user is valid
	userVerification, httpErr := dbClientInstance.VerifyUser(authObj.UserId)
	if httpErr != nil || !userVerification.IsVerified {
		render.Render(w, r, &httpErrors.ErrNotAcceptable)
		return
	}

	// return token
	atObj := models.AccessToken{AccessToken: accessToken}
	if err := render.Render(w, r, &atObj); err != nil {
		render.Render(w, r, httpErrors.ErrServerErrorRenderer(err))
		return
	}
}
