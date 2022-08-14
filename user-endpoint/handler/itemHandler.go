package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/mstolin/present-roulette/utils/httpErrors"
)

const WISHLIST_ID_KEY = "wishlistId"

func itemHandler(r chi.Router) {
	r.Route("/amazon", func(r chi.Router) {
		r.Route("/wishlists/{wishlistId}", func(r chi.Router) {
			r.Use(wishlistCtx)
			r.Get("/", importWishlist)
		})
	})
}

func wishlistCtx(nxt http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wishlistId := chi.URLParam(r, WISHLIST_ID_KEY)
		if wishlistId == "" {
			render.Render(w, r, httpErrors.ErrBadRequestRenderer(fmt.Errorf("wishlist ID is required")))
			return
		}

		ctx := context.WithValue(r.Context(), WISHLIST_ID_KEY, wishlistId)
		nxt.ServeHTTP(w, r.WithContext(ctx))
	})
}

func importWishlist(w http.ResponseWriter, r *http.Request) {
	wishlistId := r.Context().Value(WISHLIST_ID_KEY).(string)
	accessToken := r.Header.Get("Authorization")

	wishlist, err := scrapperFacadeInstance.ImportAmazonWishlist(wishlistId, accessToken)
	if err != nil {
		render.Render(w, r, err)
		return
	}

	if err := render.Render(w, r, &wishlist); err != nil {
		render.Render(w, r, httpErrors.ErrServerErrorRenderer(err))
		return
	}
}
