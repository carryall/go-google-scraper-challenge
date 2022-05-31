package constants

import "net/http"

var Errors = map[int]string{
	http.StatusUnauthorized:        "Unauthorized",
	http.StatusBadRequest:          "Bad Request",
	http.StatusUnprocessableEntity: "Unprocessable Entity",
	http.StatusInternalServerError: "Internal Server Error",
}

const (
	ERROR_CODE_MALFORM_REQUEST     = "malformed_request"
	ERROR_CODE_INVALID_PARAM       = "invalid_param"
	ERROR_CODE_INVALID_CREDENTIALS = "invalid_credentials"
)
