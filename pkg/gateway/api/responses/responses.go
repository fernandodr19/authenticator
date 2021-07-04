package responses

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/fernandodr19/library/pkg/domain/entities/accounts"
	accounts_uc "github.com/fernandodr19/library/pkg/domain/usecases/accounts"
)

// Response represents an API response
type Response struct {
	Status  int
	Error   error
	Payload interface{}
	headers map[string]string
}

// Headers list response headers
func (r *Response) Headers() map[string]string {
	return r.headers
}

// SetHeader set response header
func (r *Response) SetHeader(key, value string) {
	if r.headers == nil {
		r.headers = make(map[string]string)
	}

	r.headers[key] = value
}

// ErrorPayload represents error response payload
type ErrorPayload struct {
	Type  string `json:"type"`
	Title string `json:"title"`
}

// shared
var (
	ErrInternalServerError = ErrorPayload{Type: "error:internal_server_error", Title: "Internal Server Error"}
	ErrInvalidBody         = ErrorPayload{Type: "error:invalid_body", Title: "Invalid body"}
	ErrInvalidAuth         = ErrorPayload{Type: "error:invalid_auth", Title: "Invalid authorization"}
	ErrInvalidParams       = ErrorPayload{Type: "error:invalid_parameters", Title: "Invalid query parameters"}
	ErrNotImplemented      = ErrorPayload{Type: "error:not_implemented", Title: "Not implemented"}
)

// accounts
var (
	ErrInvalidEmail           = ErrorPayload{Type: "error:invalid_email", Title: "Invalid email"}
	ErrInvalidPassword        = ErrorPayload{Type: "error:invalid_password", Title: "Invalid password"}
	ErrEmailAlreadyRegistered = ErrorPayload{Type: "error:already_registered", Title: "Email already registered"}
)

// ErrorResponse maps response error
func ErrorResponse(err error) Response {
	switch {
	case errors.Is(err, accounts.ErrInvalidEmail):
		return BadRequest(err, ErrInvalidEmail)
	case errors.Is(err, accounts.ErrInvalidPassword):
		return BadRequest(err, ErrInvalidPassword)
	case errors.Is(err, accounts_uc.ErrNotImplemented):
		return NotImplemented(err)
	case errors.Is(err, accounts_uc.ErrEmailAlreadyRegistered):
		return Conflict(err, ErrEmailAlreadyRegistered)
	default:
		return InternalServerError(err)
	}
}

// InternalServerError 500
func InternalServerError(err error) Response {
	return Response{
		Status:  http.StatusInternalServerError,
		Error:   err,
		Payload: ErrInternalServerError,
	}
}

// NotImplemented 501
func NotImplemented(err error) Response {
	return Response{
		Status:  http.StatusNotImplemented,
		Error:   err,
		Payload: ErrNotImplemented,
	}
}

// BadRequest 400
func BadRequest(err error, payload ErrorPayload) Response {
	return genericError(http.StatusBadRequest, err, payload)
}

// UnprocessableEntity 422
func UnprocessableEntity(err error, payload ErrorPayload) Response {
	return genericError(http.StatusUnprocessableEntity, err, payload)
}

// Conflict 409
func Conflict(err error, payload ErrorPayload) Response {
	return genericError(http.StatusConflict, err, payload)
}

// NotFound 404
func NotFound(err error, payload ErrorPayload) Response {
	return genericError(http.StatusNotFound, err, payload)
}

func genericError(status int, err error, payload ErrorPayload) Response {
	return Response{
		Status:  status,
		Error:   err,
		Payload: payload,
	}
}

// NoContent 204
func NoContent() Response {
	return Response{
		Status: http.StatusNoContent,
	}
}

// OK 200
func OK(payload interface{}) Response {
	return Response{
		Status:  http.StatusOK,
		Payload: payload,
	}
}

// Created 201
func Created(payload interface{}) Response {
	return Response{
		Status:  http.StatusCreated,
		Payload: payload,
	}
}

// Accepted 202
func Accepted(payload interface{}) Response {
	return Response{
		Status:  http.StatusAccepted,
		Payload: payload,
	}
}

// SendJSON responds requests based on
func SendJSON(w http.ResponseWriter, payload interface{}, statusCode int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if payload == nil { // Blank body is not valid JSON.
		return nil
	}

	switch p := payload.(type) {
	case string:
		if p == "" {
			return nil
		}
	}

	return json.NewEncoder(w).Encode(payload)
}
