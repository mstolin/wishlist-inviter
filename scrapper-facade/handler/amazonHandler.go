package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/mstolin/wishlist-inviter/utils/httpErrors"
)

const PARAM_WISHLIST_ID = "wishlistId"

func amazonHandler(r chi.Router) {
	r.Route("/wishlists/{wishlistId}", func(r chi.Router) {
		r.Use(wishlistCtx)
		r.Get("/", scrapWishlist) // GET /wishlist/amazon/1234
	})
}

func wishlistCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if whishlistId := chi.URLParam(r, PARAM_WISHLIST_ID); whishlistId == "" {
			render.Render(w, r, httpErrors.ErrBadRequestRenderer(fmt.Errorf("wishlist_id is required")))
			return
		} else {
			ctx := context.WithValue(r.Context(), PARAM_WISHLIST_ID, whishlistId)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}

func scrapWishlist(w http.ResponseWriter, r *http.Request) {
	whishlistId := r.Context().Value(PARAM_WISHLIST_ID).(string)
	accessToken := r.Header.Get("Authorization")

	resp, httpErr := scrapperFacadeInstance.ScrapAmazonWishlist(whishlistId, accessToken)
	if httpErr != nil {
		render.Render(w, r, httpErr)
		return
	}

	if err := render.Render(w, r, &resp); err != nil {
		render.Render(w, r, httpErrors.ErrServerErrorRenderer(err))
		return
	}
}
