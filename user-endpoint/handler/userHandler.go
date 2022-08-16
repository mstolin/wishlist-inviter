package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/mstolin/present-roulette/utils/httpErrors"
	"github.com/mstolin/present-roulette/utils/models"
)

const USER_ID_KEY = "userId"
const ITEM_ID_KEY = "itemId"

func userHandler(r chi.Router) {
	r.Post("/", createUser)
	r.Route("/{userId}", func(r chi.Router) {
		r.Use(userCtx)
		r.Get("/", getUser)
		r.Delete("/", deleteUser)

		r.Route("/items", func(r chi.Router) {
			r.Get("/", getUserItems)
			r.Post("/", addUserItems)

			r.Route("/{itemId}", func(r chi.Router) {
				r.Use(itemCtx)
				r.Get("/", getUserItem)
				r.Put("/", updateUserItem)
				r.Delete("/", deleteUserItem)
			})
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

func itemCtx(nxt http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		itemId := chi.URLParam(r, ITEM_ID_KEY)
		if itemId == "" {
			render.Render(w, r, httpErrors.ErrBadRequestRenderer(fmt.Errorf("item ID is required")))
			return
		}

		// convert id to int
		id, err := strconv.Atoi(itemId)
		if err != nil {
			render.Render(w, r, httpErrors.ErrServerErrorRenderer(fmt.Errorf("invalid item ID")))
		}

		ctx := context.WithValue(r.Context(), ITEM_ID_KEY, id)
		nxt.ServeHTTP(w, r.WithContext(ctx))
	})
}

func createUser(w http.ResponseWriter, r *http.Request) {
	accessToken := r.Header.Get("Authorization")
	resp, err := userClientInstance.CreateEmptyUser(accessToken)
	if err != nil {
		render.Render(w, r, err)
		return
	}

	if err := render.Render(w, r, &resp); err != nil {
		render.Render(w, r, httpErrors.ErrServerErrorRenderer(err))
		return
	}
}

func getUser(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(USER_ID_KEY).(string)
	accessToken := r.Header.Get("Authorization")

	user, err := userClientInstance.GetUser(userId, accessToken)
	if err != nil {
		render.Render(w, r, err)
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

	user, err := userClientInstance.DeleteUser(userId, accessToken)
	if err != nil {
		render.Render(w, r, err)
		return
	}
	if err := render.Render(w, r, &user); err != nil {
		render.Render(w, r, httpErrors.ErrServerErrorRenderer(err))
		return
	}
}

func getUserItems(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(USER_ID_KEY).(string)
	accessToken := r.Header.Get("Authorization")

	itemLst, err := userClientInstance.GetUserItems(userId, accessToken)
	if err != nil {
		render.Render(w, r, err)
		return
	}

	itemLstRenderer := models.NewItemResponseListRenderer(itemLst)
	if err := render.RenderList(w, r, itemLstRenderer); err != nil {
		render.Render(w, r, httpErrors.ErrServerErrorRenderer(err))
		return
	}
}

func addUserItems(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(USER_ID_KEY).(string)

	itemLstReq := models.ItemList{}
	if err := render.Bind(r, &itemLstReq); err != nil {
		render.Render(w, r, httpErrors.ErrBadRequestRenderer(err))
		return
	}

	accessToken := r.Header.Get("Authorization")
	itemLst, err := userClientInstance.AddUserItems(userId, itemLstReq, accessToken)
	if err != nil {
		render.Render(w, r, err)
		return
	}

	itemLstRenderer := models.NewItemResponseListRenderer(itemLst)
	if err := render.RenderList(w, r, itemLstRenderer); err != nil {
		render.Render(w, r, httpErrors.ErrServerErrorRenderer(err))
		return
	}
}

func getUserItem(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(USER_ID_KEY).(string)
	itemId := r.Context().Value(ITEM_ID_KEY).(int)
	accessToken := r.Header.Get("Authorization")

	item, err := userClientInstance.GetItem(userId, itemId, accessToken)
	if err != nil {
		render.Render(w, r, err)
		return
	}

	if err := render.Render(w, r, &item); err != nil {
		render.Render(w, r, httpErrors.ErrServerErrorRenderer(err))
		return
	}
}

func deleteUserItem(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(USER_ID_KEY).(string)
	itemId := r.Context().Value(ITEM_ID_KEY).(int)
	accessToken := r.Header.Get("Authorization")

	item, err := userClientInstance.DeleteItem(userId, itemId, accessToken)
	if err != nil {
		render.Render(w, r, err)
		return
	}

	if err := render.Render(w, r, &item); err != nil {
		render.Render(w, r, httpErrors.ErrServerErrorRenderer(err))
		return
	}
}

func updateUserItem(w http.ResponseWriter, r *http.Request) {
	update := models.ItemUpdate{}
	if err := render.Bind(r, &update); err != nil {
		render.Render(w, r, httpErrors.ErrBadRequestRenderer(err))
		return
	}

	userId := r.Context().Value(USER_ID_KEY).(string)
	itemId := r.Context().Value(ITEM_ID_KEY).(int)
	accessToken := r.Header.Get("Authorization")

	item, err := userClientInstance.UpdateItem(userId, itemId, update, accessToken)
	if err != nil {
		render.Render(w, r, err)
		return
	}

	if err := render.Render(w, r, &item); err != nil {
		render.Render(w, r, httpErrors.ErrServerErrorRenderer(err))
		return
	}
}
