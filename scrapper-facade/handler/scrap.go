package handler

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func scrapWishlist(r chi.Router) {
	r.Route("/amazon/{whishlistID}", func(r chi.Router) {
		r.Use(whishlistCtx)
		r.Get("/", amazonHandler) // GET /wishlist/amazon/1234
	})
}

func whishlistCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if whishlistID := chi.URLParam(r, "whishlistID"); whishlistID == "" {
			render.Render(w, r, ErrBadRequest)
			return
		} else {
			ctx := context.WithValue(r.Context(), "whishlistID", whishlistID)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}

func amazonHandler(w http.ResponseWriter, r *http.Request) {
	whishlistId := r.Context().Value("whishlistID").(string)

	resp, error := scrapperFacadeInstance.ScrapAmazonWishlist(whishlistId)
	if error != nil {
		fmt.Fprintf(os.Stderr, "Error: %q\n", error)
		render.Render(w, r, ErrorRenderer(error))
		return
	}

	if error := render.Render(w, r, resp); error != nil {
		fmt.Fprintf(os.Stderr, "Error: %q\n", error)
		render.Render(w, r, ServerErrorRenderer(error))
		return
	}
}
