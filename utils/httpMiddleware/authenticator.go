package httpMiddleware

import (
	"net/http"

	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	"github.com/lestrrat-go/jwx/jwt"
	"github.com/mstolin/present-roulette/utils/httpErrors"
)

func JSONAuthenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, _, err := jwtauth.FromContext(r.Context())
		if err != nil {
			render.Render(w, r, httpErrors.ErrUnauthorizedRenderer(err))
			return
		}
		if token == nil || jwt.Validate(token) != nil {
			render.Render(w, r, &httpErrors.ErrUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
