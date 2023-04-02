package domain

import (
	"net/http"
)

const (
	RecordNotFound      = "data not found"
	InternalServerError = "internal server error"
	Unauthorized        = "unauthorized"
	BadRequest          = "invalid request body"
)

type ErrorResponses struct {
	Message string `json:"message"`
}

func ErrorResponse(err string) (code int, errMessage string) {

	switch err {
	case RecordNotFound:
		return http.StatusNotFound, RecordNotFound
	case InternalServerError:
		return http.StatusInternalServerError, InternalServerError
	case Unauthorized:
		return http.StatusUnauthorized, Unauthorized
	case BadRequest:
		return http.StatusBadRequest, BadRequest
	default:
		return http.StatusNotFound, err
	}

}
