package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/mstolin/present-roulette/utils/errors"
	"github.com/mstolin/present-roulette/utils/models"
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
			render.Render(w, r, errors.ErrorRenderer(fmt.Errorf("item ID is required")))
			return
		}

		// convert id to int
		id, err := strconv.Atoi(itemId)
		if err != nil {
			render.Render(w, r, errors.ErrorRenderer(fmt.Errorf("invalid item ID")))
		}

		ctx := context.WithValue(r.Context(), ITEM_ID_KEY, id)
		nxt.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getItems(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(USER_ID_KEY).(string)
	itemLst, err := dbClientInstance.GetItemsByUser(userId)
	if err != nil {
		render.Render(w, r, errors.ServerErrorRenderer(err))
		return
	}

	if err := render.Render(w, r, &itemLst); err != nil {
		render.Render(w, r, errors.ServerErrorRenderer(err))
		return
	}
}

func addItems(w http.ResponseWriter, r *http.Request) {
	itemLst := models.ItemList{}
	if err := render.Bind(r, &itemLst); err != nil {
		render.Render(w, r, errors.ErrBadRequest)
		return
	}

	userId := r.Context().Value(USER_ID_KEY).(string)
	items, err := dbClientInstance.AddItemsToUser(userId, itemLst)
	if err != nil {
		render.Render(w, r, errors.ErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, &items); err != nil {
		render.Render(w, r, errors.ServerErrorRenderer(err))
		return
	}
}

func getItem(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(USER_ID_KEY).(string)
	itemId := r.Context().Value(ITEM_ID_KEY).(int)
	item, err := dbClientInstance.GetItemByUser(userId, itemId)
	if err != nil {
		render.Render(w, r, errors.ErrNotFound)
		return
	}

	if err := render.Render(w, r, &item); err != nil {
		render.Render(w, r, errors.ServerErrorRenderer(err))
		return
	}
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	// TODO Before update, check if request only has allowed data
	item := models.Item{}
	if err := render.Bind(r, &item); err != nil {
		render.Render(w, r, errors.ErrBadRequest)
		return
	}

	userId := r.Context().Value(USER_ID_KEY).(string)
	itemId := r.Context().Value(ITEM_ID_KEY).(int)
	update, err := dbClientInstance.UpdateItemByUser(userId, itemId, item)
	if err != nil {
		render.Render(w, r, errors.ErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, &update); err != nil {
		render.Render(w, r, errors.ServerErrorRenderer(err))
		return
	}
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(USER_ID_KEY).(string)
	itemId := r.Context().Value(ITEM_ID_KEY).(int)
	item, err := dbClientInstance.DeleteItemByUser(userId, itemId)
	if err != nil {
		render.Render(w, r, errors.ErrNotFound)
		return
	}
	if err := render.Render(w, r, &item); err != nil {
		render.Render(w, r, errors.ServerErrorRenderer(err))
		return
	}
}
