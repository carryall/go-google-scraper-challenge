package constants

import "net/http"

var Errors = map[int]string{
	http.StatusUnauthorized:        "Unauthorized",
	http.StatusBadRequest:          "Bad Request",
	http.StatusUnprocessableEntity: "Unprocessable Entity",
	http.StatusInternalServerError: "Internal Server Error",
}
