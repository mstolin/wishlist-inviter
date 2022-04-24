package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/mstolin/present-roulette/utils/errors"
)

const WISHLIST_ID_KEY = "wishlistId"

func itemHandler(r chi.Router) {
	r.Route("/amazon", func(r chi.Router) {
		r.Route("/wishlist/{wishlistId}", func(r chi.Router) {
			r.Use(wishlistCtx)
			r.Get("/", importWishlist)
		})
	})
}

func wishlistCtx(nxt http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wishlistId := chi.URLParam(r, WISHLIST_ID_KEY)
		if wishlistId == "" {
			render.Render(w, r, errors.ErrBadRequest)
			// TODO Mach uberall so
			// render.Render(w, r, errors.ErrorRenderer(fmt.Errorf("wishlist ID is required")))
			return
		}

		// convert id to int
		id, err := strconv.Atoi(wishlistId)
		if err != nil {
			render.Render(w, r, errors.ErrorRenderer(fmt.Errorf("invalid wishlist ID")))
		}

		ctx := context.WithValue(r.Context(), WISHLIST_ID_KEY, id)
		nxt.ServeHTTP(w, r.WithContext(ctx))
	})
}

func importWishlist(w http.ResponseWriter, r *http.Request) {
	wishlistId := ""

	itemLst, err := wishlistClientInstance.ImportAmazonWishlist(wishlistId)
	if err != nil {
		render.Render(w, r, errors.ServerErrorRenderer(err))
		return
	}

	if err := render.Render(w, r, &itemLst); err != nil {
		render.Render(w, r, errors.ServerErrorRenderer(err))
		return
	}
}