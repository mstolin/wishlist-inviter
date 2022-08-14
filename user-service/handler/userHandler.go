package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	"github.com/mstolin/present-roulette/utils/httpErrors"
)

const USER_ID_KEY = "userId"

func userHandler(r chi.Router) {
	// unrestricted endpoints
	r.Group(func(r chi.Router) {
		// registration not be restricted
		r.Post("/", createUser)
	})

	// restricted endpoints
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Route("/{userId}", func(r chi.Router) {
			r.Use(userCtx)
			r.Get("/", getUser)
			r.Delete("/", deleteUser)

			r.Route("/items", itemHandler)
		})
	})
}

func userCtx(nxt http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId := chi.URLParam(r, USER_ID_KEY)
		if userId == "" {
			render.Render(w, r, httpErrors.ErrBadRequestRenderer(fmt.Errorf("user ID is required")))
			return
		}

		ctx := context.WithValue(r.Context(), USER_ID_KEY, userId)
		nxt.ServeHTTP(w, r.WithContext(ctx))
	})
}

func createUser(w http.ResponseWriter, r *http.Request) {
	accessToken := r.Header.Get("Authorization")
	resp, httpErr := dbClientInstance.CreateUser(accessToken)
	if httpErr != nil {
		render.Render(w, r, httpErr)
		return
	}

	if error := render.Render(w, r, &resp); error != nil {
		render.Render(w, r, httpErrors.ErrServerErrorRenderer(error))
		return
	}
}

func getUser(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(USER_ID_KEY).(string)
	accessToken := r.Header.Get("Authorization")
	user, httpErr := dbClientInstance.GetUser(userId, accessToken)
	if httpErr != nil {
		render.Render(w, r, httpErr)
		return
	}
	if err := render.Render(w, r, &user); err != nil {
		render.Render(w, r, httpErrors.ErrServerErrorRenderer(err))
		return
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(USER_ID_KEY).(string)
	accessToken := r.Header.Get("Authorization")
	user, httpErr := dbClientInstance.DeleteUser(userId, accessToken)
	if httpErr != nil {
		render.Render(w, r, httpErr)
		return
	}
	if err := render.Render(w, r, &user); err != nil {
		render.Render(w, r, httpErrors.ErrServerErrorRenderer(err))
		return
	}
}
