package httpErrors

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	ErrBadRequest         = "bad request"
	ErrEmailAlreadyExists = "user with given email already exists"
	ErrNoSuchUser         = "user not found"
	ErrWrongCredentials   = "wrong Credentials"
	ErrNotFound           = "not Found"
	ErrUnauthorized       = "unauthorized"
	ErrForbidden          = "forbidden"
	ErrBadQueryParams     = "invalid query params"
)

type RestErrorResponse struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
	Detail string `json:"detail"`
}

func parseErrors(err error) RestErrorResponse {
	switch err.Error() {
	case ErrBadRequest:
		return RestErrorResponse{Status: http.StatusBadRequest, Error: ErrBadRequest, Detail: err.Error()}
	case ErrEmailAlreadyExists:
		return RestErrorResponse{Status: http.StatusBadRequest, Error: ErrBadRequest, Detail: err.Error()}
	case ErrNoSuchUser:
		return RestErrorResponse{Status: http.StatusNotFound, Error: ErrNoSuchUser, Detail: err.Error()}
	case ErrWrongCredentials:
		return RestErrorResponse{Status: http.StatusUnauthorized, Error: ErrWrongCredentials, Detail: err.Error()}
	case ErrNotFound:
		return RestErrorResponse{Status: http.StatusNotFound, Error: ErrNotFound, Detail: err.Error()}
	case ErrUnauthorized:
		return RestErrorResponse{Status: http.StatusUnauthorized, Error: ErrUnauthorized, Detail: err.Error()}
	case ErrForbidden:
		return RestErrorResponse{Status: http.StatusForbidden, Error: ErrForbidden, Detail: err.Error()}
	case ErrBadQueryParams:
		return RestErrorResponse{Status: http.StatusBadRequest, Error: ErrBadQueryParams, Detail: err.Error()}
	default:
		return RestErrorResponse{Status: http.StatusInternalServerError, Error: "Internal Server Error", Detail: err.Error()}
	}
}

func WriteJsonResponse(w http.ResponseWriter, err error) {
	// Parse the error to create the corresponding RestErrorResponse
	restError := parseErrors(err)

	// Marshal the RestErrorResponse to JSON
	jsonError, marshalErr := json.Marshal(restError)
	if marshalErr != nil {
		fmt.Println(marshalErr)
		return
	}

	// Set the Content-Type header
	w.Header().Set("Content-Type", "application/json")

	// Write the HTTP status code
	w.WriteHeader(restError.Status)

	// Write the JSON error response to the response writer
	w.Write(jsonError)
}
