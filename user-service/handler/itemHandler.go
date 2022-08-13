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

func getItems(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(USER_ID_KEY).(string)
	itemLst, err := dbClientInstance.GetItemsByUser(userId)
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

func addItems(w http.ResponseWriter, r *http.Request) {
	itemLst := models.ItemList{}
	if err := render.Bind(r, &itemLst); err != nil {
		render.Render(w, r, httpErrors.ErrBadRequestRenderer(err))
		return
	}

	userId := r.Context().Value(USER_ID_KEY).(string)
	items, err := dbClientInstance.AddItemsToUser(userId, itemLst)
	if err != nil {
		render.Render(w, r, err)
		return
	}

	itemLstRenderer := models.NewItemResponseListRenderer(items)
	if err := render.RenderList(w, r, itemLstRenderer); err != nil {
		render.Render(w, r, httpErrors.ErrServerErrorRenderer(err))
		return
	}
}

func getItem(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(USER_ID_KEY).(string)
	itemId := r.Context().Value(ITEM_ID_KEY).(int)
	item, err := dbClientInstance.GetItemByUser(userId, itemId)
	if err != nil {
		render.Render(w, r, err)
		return
	}

	if err := render.Render(w, r, &item); err != nil {
		render.Render(w, r, httpErrors.ErrServerErrorRenderer(err))
		return
	}
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	update := models.ItemUpdate{}
	if err := render.Bind(r, &update); err != nil {
		render.Render(w, r, httpErrors.ErrBadRequestRenderer(err))
		return
	}

	userId := r.Context().Value(USER_ID_KEY).(string)
	itemId := r.Context().Value(ITEM_ID_KEY).(int)
	updatedItem, err := dbClientInstance.UpdateItemByUser(userId, itemId, update)
	if err != nil {
		render.Render(w, r, err)
		return
	}
	if err := render.Render(w, r, &updatedItem); err != nil {
		render.Render(w, r, httpErrors.ErrServerErrorRenderer(err))
		return
	}
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(USER_ID_KEY).(string)
	itemId := r.Context().Value(ITEM_ID_KEY).(int)
	item, err := dbClientInstance.DeleteItemByUser(userId, itemId)
	if err != nil {
		render.Render(w, r, err)
		return
	}
	if err := render.Render(w, r, &item); err != nil {
		render.Render(w, r, httpErrors.ErrServerErrorRenderer(err))
		return
	}
}
