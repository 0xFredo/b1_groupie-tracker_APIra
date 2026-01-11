package utils

import (
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, status int) {
	w.WriteHeader(status)

	var message string
	switch status {
	case http.StatusNotFound:
		message = "404 - Page Not Found"
	case http.StatusBadRequest:
		message = "400 - Bad Request"
	case http.StatusInternalServerError:
		message = "500 - Internal Server Error"
	default:
		message = "Error"
	}

	w.Write([]byte(message))
}
