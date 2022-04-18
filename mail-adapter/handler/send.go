package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/mstolin/present-roulette/mail-adapter/models"
)

func send(router chi.Router) {
	router.Post("/invitation", invitationHandler)
}

func invitationHandler(writer http.ResponseWriter, request *http.Request) {
	invitation := &models.Invitation{}

	if error := render.Bind(request, invitation); error != nil {
		fmt.Fprintf(os.Stderr, "Error: %q\n", error)
		render.Render(writer, request, ErrBadRequest)
		return
	}

	gmailResp, error := gmailClientInstance.PostInvitation(invitation)
	if error != nil {
		fmt.Fprintf(os.Stderr, "Error: %q\n", error)
		render.Render(writer, request, ErrorRenderer(error))
		return
	}

	if error := render.Render(writer, request, gmailResp); error != nil {
		fmt.Fprintf(os.Stderr, "Error: %q\n", error)
		render.Render(writer, request, ServerErrorRenderer(error))
		return
	}
}
