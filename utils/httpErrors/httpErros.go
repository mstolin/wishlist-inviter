package httpErrors

import (
	"net/http"

	"github.com/go-chi/render"
)

type ErrorResponse struct {
	Detail ErrorDetail `json:"error"`
}

type ErrorDetail struct {
	Status  int    `json:"status"`
	Err     string `json:"error"`
	Message string `json:"message"`
}

var (
	ErrDetailMethodNotAllowed = ErrorDetail{Status: http.StatusMethodNotAllowed, Err: http.StatusText(http.StatusMethodNotAllowed), Message: "The target resource doesn't support this method."}
	ErrDetailNotFound         = ErrorDetail{Status: http.StatusNotFound, Err: http.StatusText(http.StatusNotFound), Message: "The requested resource is not available."}
	ErrDetailUnauthorized     = ErrorDetail{Status: http.StatusUnauthorized, Err: http.StatusText(http.StatusUnauthorized), Message: "You are not authorized to access this resource."}
	ErrDetailBadRequest       = ErrorDetail{Status: http.StatusBadRequest, Err: http.StatusText(http.StatusBadRequest), Message: "Received request is invalid."}
)
var (
	//ErrMethodNotAllowed = ErrorResponse{Detail: ErrMethodNotAllowed}
	ErrNotFound     = ErrorResponse{Detail: ErrDetailNotFound}
	ErrUnauthorized = ErrorResponse{Detail: ErrDetailUnauthorized}
	ErrBadRequest   = ErrorResponse{Detail: ErrDetailBadRequest}
)

func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.Detail.Status)
	return nil
}

func (e *ErrorDetail) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func ErrBadRequestRenderer(err error) *ErrorResponse {
	return &ErrorResponse{
		Detail: ErrorDetail{
			Status:  http.StatusBadRequest,
			Err:     http.StatusText(http.StatusBadRequest),
			Message: err.Error(),
		},
	}
}

func ErrNotFoundRenderer(err error) *ErrorResponse {
	return &ErrorResponse{
		Detail: ErrorDetail{
			Status:  http.StatusNotFound,
			Err:     http.StatusText(http.StatusNotFound),
			Message: err.Error(),
		},
	}
}

func ErrServerErrorRenderer(err error) *ErrorResponse {
	return &ErrorResponse{
		Detail: ErrorDetail{
			Status:  http.StatusInternalServerError,
			Err:     http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
		},
	}
}

func ErrUnauthorizedRenderer(err error) *ErrorResponse {
	return &ErrorResponse{
		Detail: ErrorDetail{
			Status:  http.StatusUnauthorized,
			Err:     http.StatusText(http.StatusUnauthorized),
			Message: err.Error(),
		},
	}
}
