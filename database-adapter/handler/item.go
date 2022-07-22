package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
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
		r.Delete("/", deleteItem)
		r.Put("/", updateItem)
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
			render.Render(w, r, httpErrors.ErrBadRequestRenderer(fmt.Errorf("invalid item ID")))
		}

		ctx := context.WithValue(r.Context(), ITEM_ID_KEY, id)
		nxt.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getItems(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(USER_ID_KEY).(string)
	user, err := dbHandler.GetUserById(userId)
	if err != nil {
		render.Render(w, r, &httpErrors.ErrNotFound)
		return
	}

	itemLstRenderer := models.NewItemResponseListRenderer(user.Items)
	if err := render.RenderList(w, r, itemLstRenderer); err != nil {
		render.Render(w, r, httpErrors.ErrServerErrorRenderer(err))
		return
	}
}

func addItems(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(USER_ID_KEY).(string)
	user, err := dbHandler.GetUserById(userId)
	if err != nil {
		render.Render(w, r, &httpErrors.ErrNotFound)
		return
	}

	itemLst := models.ItemList{}
	if err := render.Bind(r, &itemLst); err != nil {
		render.Render(w, r, httpErrors.ErrBadRequestRenderer(err))
		return
	}

	if err := dbHandler.AddItemsToUser(user, itemLst); err != nil {
		render.Render(w, r, httpErrors.ErrBadRequestRenderer(err))
		return
	}

	// response all items
	getItems(w, r)
}

func getItem(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(USER_ID_KEY).(string)
	user, err := dbHandler.GetUserById(userId)
	if err != nil {
		render.Render(w, r, httpErrors.ErrNotFoundRenderer(fmt.Errorf("user with id %s not found", userId)))
		return
	}

	itemId := r.Context().Value(ITEM_ID_KEY).(int)
	item, err := dbHandler.GetItemByUser(user, itemId)
	if err != nil {
		render.Render(w, r, httpErrors.ErrNotFoundRenderer(fmt.Errorf("item with id %d not found", itemId)))
		return
	}

	if err := render.Render(w, r, &item); err != nil {
		render.Render(w, r, httpErrors.ErrServerErrorRenderer(err))
		return
	}
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(USER_ID_KEY).(string)
	user, err := dbHandler.GetUserById(userId)
	if err != nil {
		render.Render(w, r, httpErrors.ErrNotFoundRenderer(fmt.Errorf("user with id %s not found", userId)))
		return
	}

	itemId := r.Context().Value(ITEM_ID_KEY).(int)
	item, err := dbHandler.DeleteItem(user, itemId)
	if err != nil {
		render.Render(w, r, httpErrors.ErrNotFoundRenderer(fmt.Errorf("item with id %d not found", itemId)))
		return
	}

	if err := render.Render(w, r, &item); err != nil {
		render.Render(w, r, httpErrors.ErrServerErrorRenderer(err))
		return
	}
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(USER_ID_KEY).(string)
	user, err := dbHandler.GetUserById(userId)
	if err != nil {
		render.Render(w, r, httpErrors.ErrNotFoundRenderer(fmt.Errorf("user with id %s not found", userId)))
		return
	}

	update := models.ItemUpdate{}
	if err := render.Bind(r, &update); err != nil {
		render.Render(w, r, httpErrors.ErrBadRequestRenderer(err))
		return
	}

	fmt.Printf("UPDATE - NAME: %s, PRICE: %f, ID: %s, BOUGHT?: %t\n",
		update.Name, update.Price, update.VendorID, update.HasBeenBaught)

	itemId := r.Context().Value(ITEM_ID_KEY).(int)
	item, err := dbHandler.UpdateItemByUser(user, itemId, update)
	if err != nil {
		render.Render(w, r, httpErrors.ErrServerErrorRenderer(err))
		return
	}

	if err := render.Render(w, r, &item); err != nil {
		render.Render(w, r, httpErrors.ErrServerErrorRenderer(err))
		return
	}
}
