package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func invitationHandler(router chi.Router) {
	router.Post("/", sendInvitation)
}

func sendInvitation(w http.ResponseWriter, r *http.Request) {
	// Create User
	resp, err := dbClientInstance.CreateUser()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %q\n", err)
		render.Render(w, r, ErrorRenderer(err))
		return
	}

	if error := render.Render(w, r, resp); error != nil { // TODO better success response
		fmt.Fprintf(os.Stderr, "Error: %q\n", error)
		render.Render(w, r, ServerErrorRenderer(error))
		return
	}
}
