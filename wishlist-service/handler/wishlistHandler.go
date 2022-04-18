package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/mstolin/present-roulette/wishlist-service/models"
)

func wishlistHandler(router chi.Router) {
	router.Post("/", createWishlist)
}

func createWishlist(w http.ResponseWriter, r *http.Request) {
	createReq := &models.CreateReq{}

	if error := render.Bind(r, createReq); error != nil {
		fmt.Fprintf(os.Stderr, "Error: %q\n", error)
		render.Render(w, r, ErrBadRequest)
		return
	}

	// 1. Get Wishlist from vendor
	wishlist, err := scrapperClientInstance.GetWishlist(createReq)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %q\n", err)
		render.Render(w, r, ErrorRenderer(err))
		return
	}

	// 2. Save items to database
	resp, err := dbClientInstance.SaveItems(createReq.UserId, wishlist.Items)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %q\n", err)
		render.Render(w, r, ErrorRenderer(err))
		return
	}

	if error := render.Render(w, r, &resp); error != nil { // TODO better success response
		fmt.Fprintf(os.Stderr, "Error: %q\n", error)
		render.Render(w, r, ServerErrorRenderer(error))
		return
	}
}
