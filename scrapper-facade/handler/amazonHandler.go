package handler

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/mstolin/present-roulette/utils/errors"
)

const PARAM_WISHLIST_ID = "wishlistId"

func amazonHandler(r chi.Router) {
	r.Route("/wishlist/{wishlistId}", func(r chi.Router) {
		r.Use(wishlistCtx)
		r.Get("/", scrapWishlist) // GET /wishlist/amazon/1234
	})
}

func wishlistCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if whishlistId := chi.URLParam(r, PARAM_WISHLIST_ID); whishlistId == "" {
			render.Render(w, r, errors.ErrBadRequest)
			return
		} else {
			ctx := context.WithValue(r.Context(), PARAM_WISHLIST_ID, whishlistId)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}

func scrapWishlist(w http.ResponseWriter, r *http.Request) {
	whishlistId := r.Context().Value(PARAM_WISHLIST_ID).(string)

	resp, error := scrapperFacadeInstance.ScrapAmazonWishlist(whishlistId)
	if error != nil {
		render.Render(w, r, errors.ErrorRenderer(error))
		return
	}

	if error := render.Render(w, r, &resp); error != nil {
		render.Render(w, r, errors.ServerErrorRenderer(error))
		return
	}
}
