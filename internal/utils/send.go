package utils

import (
	"encoding/json"
	"net/http"
)

const (
	ErrorBadRequestBody = "Bad Request body. Make sure it is properly formatted json"
)

func Send(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func Unauthorized(w http.ResponseWriter) {
	// The error message is intentionally vague to protect against guessing attacks
	Send(w, http.StatusUnauthorized, GenericResult{Success: false, Data: "Unauthorized"})
}

func OK(w http.ResponseWriter, data interface{}) {
	Send(w, http.StatusOK, data)
}

func BadRequest(w http.ResponseWriter, data interface{}) {
	Send(w, http.StatusBadRequest, GenericResult{Success: false, Data: data})
}

func InternalServerError(w http.ResponseWriter, data interface{}) {
	Send(w, http.StatusInternalServerError, GenericResult{Success: false, Data: data})
}

func NotFound(w http.ResponseWriter, data interface{}) {
	Send(w, http.StatusNotFound, data)
}
