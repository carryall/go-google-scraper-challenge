package errors

import (
	"errors"
	"net/http"
)

var (
	ErrInvalidRequest      = errors.New("invalid_request")
	ErrInvalidCredentials  = errors.New("invalid_credentials")
	ErrInvalidAuthClient   = errors.New("invalid_authentication_client")
	ErrUnauthorizedUser    = errors.New("unauthorized_user")
	ErrUnProcessableEntity = errors.New("unprocessable_entity")
	ErrNotFound            = errors.New("not_found")
	ErrServerError         = errors.New("server_error")
)

var Titles = map[error]string{
	ErrInvalidRequest:      "Invalid Request",
	ErrInvalidCredentials:  "Invalid Creadentials",
	ErrInvalidAuthClient:   "Invalid Authentication Client",
	ErrUnauthorizedUser:    "Unauthorized User",
	ErrUnProcessableEntity: "Unprocessable Entity",
	ErrNotFound:            "Not Found",
	ErrServerError:         "Internal Server Error",
}

var Descriptions = map[error]string{
	ErrInvalidRequest:      "The request is missing a required parameter, contains an invalid parameter value or mulformed",
	ErrInvalidCredentials:  "The username or password is invalid",
	ErrInvalidAuthClient:   "Authentication client is invalid",
	ErrUnauthorizedUser:    "The user is not authorized",
	ErrUnProcessableEntity: "The request is unprocessable",
	ErrNotFound:            "Cannot find the resource you are looking for",
	ErrServerError:         "Something went wrong",
}

var StatusCodes = map[error]int{
	ErrInvalidRequest:      http.StatusBadRequest,
	ErrInvalidCredentials:  http.StatusUnauthorized,
	ErrInvalidAuthClient:   http.StatusUnauthorized,
	ErrUnauthorizedUser:    http.StatusUnauthorized,
	ErrUnProcessableEntity: http.StatusUnprocessableEntity,
	ErrNotFound:            http.StatusNotFound,
	ErrServerError:         http.StatusInternalServerError,
}
