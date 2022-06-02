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

/*type ErrorResponse struct {
	Err        error  `json:"-"`
	StatusCode int    `json:"-"`
	StatusText string `json:"status_text"`
	Message    string `json:"message"`
}*/

/*var (
	ErrMethodNotAllowed = &ErrorResponse{StatusCode: 405, Message: "Method not allowed"}
	ErrNotFound         = &ErrorResponse{StatusCode: 404, Message: "Resource not found"}
	ErrBadRequest       = &ErrorResponse{StatusCode: 400, Message: "Bad request"}
)*/

var (
	ErrMethodNotAllowed = ErrorDetail{Status: 405, Err: "Method Not Allowed", Message: "The target resource doesn't support this method."}
	ErrDetailNotFound   = ErrorDetail{Status: 404, Err: "Not Found", Message: "The requested resource is not available."}
	ErrDetailBadRequest = ErrorDetail{Status: 400, Err: "Bad Request", Message: "Received request is invalid."}
)
var (
	ErrNotFound   = ErrorResponse{Detail: ErrDetailNotFound}
	ErrBadRequest = ErrorResponse{Detail: ErrDetailBadRequest}
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
			Status:  400,
			Err:     "Bad Request",
			Message: err.Error(),
		},
	}
}

func ErrNotFoundRenderer(err error) *ErrorResponse {
	return &ErrorResponse{
		Detail: ErrorDetail{
			Status:  404,
			Err:     "Not Found",
			Message: err.Error(),
		},
	}
}

func ErrServerErrorRenderer(err error) *ErrorResponse {
	return &ErrorResponse{
		Detail: ErrorDetail{
			Status:  500,
			Err:     "Internal Server Error",
			Message: err.Error(),
		},
	}
}

/*func ErrorRenderer(err error) *ErrorResponse {
	return &ErrorResponse{
		Err:        err,
		StatusCode: 400,
		StatusText: "Bad request",
		Message:    err.Error(),
	}
}
func ServerErrorRenderer(err error) *ErrorResponse {
	return &ErrorResponse{
		Err:        err,
		StatusCode: 500,
		StatusText: "Internal server error",
		Message:    err.Error(),
	}
}*/
