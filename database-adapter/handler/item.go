package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/mstolin/present-roulette/database-adapter/models"
)

const ITEM_ID_KEY = "itemId"

func itemHandler(r chi.Router) {
	r.Post("/", addItems)
	r.Get("/", getItems)
	r.Route("/{itemId}", func(r chi.Router) {
		r.Use(itemCtx)
		r.Get("/", getItem)
		r.Put("/", updateItem)
		r.Delete("/", deleteItem)
	})
}

func itemCtx(nxt http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		itemId := chi.URLParam(r, ITEM_ID_KEY)
		if itemId == "" {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("item ID is required")))
			return
		}

		// convert id to int
		id, err := strconv.Atoi(itemId)
		if err != nil {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("invalid item ID")))
		}

		ctx := context.WithValue(r.Context(), ITEM_ID_KEY, id)
		nxt.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getItems(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(USER_ID_KEY).(string)
	user, err := dbHandler.GetUserById(userId)
	if err != nil {
		render.Render(w, r, ErrNotFound)
		return
	}

	itemLst := models.ItemList{Items: user.Items}
	if err := render.Render(w, r, &itemLst); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

func addItems(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(USER_ID_KEY).(string)
	user, err := dbHandler.GetUserById(userId)
	if err != nil {
		render.Render(w, r, ErrNotFound)
		return
	}

	itemLst := &models.ItemList{}
	if err := render.Bind(r, itemLst); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	if err := dbHandler.AddItemsToUser(user, itemLst.Items); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}

	if err := render.Render(w, r, itemLst); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

func getItem(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(USER_ID_KEY).(string)
	user, err := dbHandler.GetUserById(userId)
	if err != nil {
		render.Render(w, r, ErrNotFound)
		return
	}

	itemId := r.Context().Value(ITEM_ID_KEY).(int)
	item, err := dbHandler.GetItemByUser(user, itemId)
	if err != nil {
		render.Render(w, r, ErrNotFound)
		return
	}

	if err := render.Render(w, r, &item); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	// TODO Before update, check if request only has allowed data
	userId := r.Context().Value(USER_ID_KEY).(string)
	user, err := dbHandler.GetUserById(userId)
	if err != nil {
		render.Render(w, r, ErrNotFound)
		return
	}

	item := &models.Item{}
	if err := render.Bind(r, item); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	itemId := r.Context().Value(ITEM_ID_KEY).(int)
	update, err := dbHandler.UpdateItemByUser(user, itemId, *item)
	if err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}

	if err := render.Render(w, r, &update); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(USER_ID_KEY).(string)
	user, err := dbHandler.GetUserById(userId)
	if err != nil {
		render.Render(w, r, ErrNotFound)
		return
	}

	itemId := r.Context().Value(ITEM_ID_KEY).(int)
	item, err := dbHandler.DeleteItem(user, itemId)
	if err != nil {
		render.Render(w, r, ErrNotFound)
		return
	}

	if err := render.Render(w, r, &item); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}
