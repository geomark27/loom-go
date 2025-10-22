package helpers

import (
	"encoding/json"
	"net/http"
)

// Response estructura estándar de respuesta
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// RespondJSON envía una respuesta JSON
func RespondJSON(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// RespondSuccess envía una respuesta exitosa estructurada
func RespondSuccess(w http.ResponseWriter, data interface{}, message string) {
	RespondJSON(w, Response{
		Status:  "success",
		Message: message,
		Data:    data,
	}, http.StatusOK)
}

// RespondError envía una respuesta de error estructurada
func RespondError(w http.ResponseWriter, err error, status int) {
	RespondJSON(w, Response{
		Status: "error",
		Error:  err.Error(),
	}, status)
}

// RespondCreated envía una respuesta 201 Created
func RespondCreated(w http.ResponseWriter, data interface{}, message string) {
	RespondJSON(w, Response{
		Status:  "success",
		Message: message,
		Data:    data,
	}, http.StatusCreated)
}

// RespondNoContent envía 204 No Content
func RespondNoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

// RespondBadRequest envía 400 Bad Request
func RespondBadRequest(w http.ResponseWriter, message string) {
	RespondJSON(w, Response{
		Status: "error",
		Error:  message,
	}, http.StatusBadRequest)
}

// RespondUnauthorized envía 401 Unauthorized
func RespondUnauthorized(w http.ResponseWriter, message string) {
	RespondJSON(w, Response{
		Status: "error",
		Error:  message,
	}, http.StatusUnauthorized)
}

// RespondForbidden envía 403 Forbidden
func RespondForbidden(w http.ResponseWriter, message string) {
	RespondJSON(w, Response{
		Status: "error",
		Error:  message,
	}, http.StatusForbidden)
}

// RespondNotFound envía 404 Not Found
func RespondNotFound(w http.ResponseWriter, message string) {
	RespondJSON(w, Response{
		Status: "error",
		Error:  message,
	}, http.StatusNotFound)
}

// RespondInternalError envía 500 Internal Server Error
func RespondInternalError(w http.ResponseWriter, err error) {
	RespondJSON(w, Response{
		Status: "error",
		Error:  "Internal server error",
	}, http.StatusInternalServerError)
}
