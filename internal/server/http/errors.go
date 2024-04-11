package http

import (
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/go-openapi/errors"
	"github.com/krivenkov/order/internal/server/http/auth"
	"github.com/krivenkov/order/internal/server/http/models"
	"github.com/krivenkov/pkg/ptr"
)

// DefaultHTTPCode is used when the error Code cannot be used as an HTTP code.
var DefaultHTTPCode = http.StatusUnprocessableEntity

func serveError(rw http.ResponseWriter, r *http.Request, err error) {
	rw.Header().Set("Content-Type", "application/json")

	switch e := err.(type) {
	case auth.ErrInvalidGrand:
		rw.WriteHeader(http.StatusUnauthorized)
		_, _ = rw.Write(errorAsJSON(errors.New(http.StatusUnauthorized, e.Description)))
	case errors.Error:
		value := reflect.ValueOf(e)
		if value.Kind() == reflect.Ptr && value.IsNil() {
			rw.WriteHeader(http.StatusInternalServerError)
			_, _ = rw.Write(errorAsJSON(errors.New(http.StatusInternalServerError, "Unknown error")))
			return
		}
		rw.WriteHeader(asHTTPCode(int(e.Code())))
		if r == nil || r.Method != http.MethodHead {
			_, _ = rw.Write(errorAsJSON(e))
		}
	case nil:
		rw.WriteHeader(http.StatusInternalServerError)
		_, _ = rw.Write(errorAsJSON(errors.New(http.StatusInternalServerError, "Unknown error")))
	default:
		rw.WriteHeader(http.StatusInternalServerError)
		if r == nil || r.Method != http.MethodHead {
			_, _ = rw.Write(errorAsJSON(errors.New(http.StatusInternalServerError, err.Error())))
		}
	}
}

func errorAsJSON(err errors.Error) []byte {
	var errorCode string

	switch err.Code() {
	case http.StatusUnauthorized:
		errorCode = models.ErrorErrorInvalidGrant
	case http.StatusUnprocessableEntity,
		http.StatusMethodNotAllowed:
		errorCode = models.ErrorErrorInvalidRequest
	case http.StatusForbidden:
		errorCode = models.ErrorErrorAccessDenied
	case http.StatusNotFound:
		errorCode = models.ErrorErrorNotFound
	default:
		errorCode = models.ErrorErrorServerError
	}

	b, _ := json.Marshal(models.Error{
		Error:            ptr.Pointer(errorCode),
		ErrorDescription: ptr.Pointer(err.Error()),
	})
	return b
}

func asHTTPCode(input int) int {
	if input >= 600 {
		return DefaultHTTPCode
	}
	return input
}
