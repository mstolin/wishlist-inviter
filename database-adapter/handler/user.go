package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/mstolin/present-roulette/utils/errors"
	"github.com/mstolin/present-roulette/utils/models"
)

const USER_ID_KEY = "userId"

func userHandler(r chi.Router) {
	r.Post("/", createUser)
	r.Route("/{userId}", func(r chi.Router) {
		r.Use(userCtx)
		r.Get("/", getUser)
		r.Delete("/", deleteUser)

		r.Route("/items", itemHandler)
	})
}

func userCtx(nxt http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId := chi.URLParam(r, USER_ID_KEY)
		if userId == "" {
			render.Render(w, r, errors.ErrorRenderer(fmt.Errorf("user ID is required")))
			return
		}

		ctx := context.WithValue(r.Context(), USER_ID_KEY, userId)
		nxt.ServeHTTP(w, r.WithContext(ctx))
	})
}

func createUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}

	if err := render.Bind(r, user); err != nil {
		render.Render(w, r, errors.ErrBadRequest)
		return
	}

	model, err := dbHandler.CreateUser(user)
	if err != nil {
		render.Render(w, r, errors.ErrorRenderer(err))
		return
	}

	if err := render.Render(w, r, &model); err != nil {
		render.Render(w, r, errors.ServerErrorRenderer(err))
		return
	}
}

func getUser(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(USER_ID_KEY).(string)
	user, err := dbHandler.GetUserById(id)
	if err != nil {
		render.Render(w, r, errors.ErrNotFound)
		return
	}

	log.Default().Printf("LEN: %d \n", len(user.Items))

	if err := render.Render(w, r, &user); err != nil {
		render.Render(w, r, errors.ServerErrorRenderer(err))
		return
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(USER_ID_KEY).(string)
	user, err := dbHandler.DeleteUserById(id)
	if err != nil {
		render.Render(w, r, errors.ErrNotFound)
		return
	}

	if err := render.Render(w, r, &user); err != nil {
		render.Render(w, r, errors.ServerErrorRenderer(err))
		return
	}
}
