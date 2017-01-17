package components

import (
	"encoding/json"
	"net/http"
)

// Response model.
type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// SendResponse send response in json format.
func SendResponse(w http.ResponseWriter, r interface{}, status int) {
	jsonMessage, err := json.Marshal(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonMessage)
}

// SendErrorResponse send response with error status.
func SendErrorResponse(w http.ResponseWriter, m string) {
	response := Response{
		Status:  "Error",
		Message: m,
	}
	SendResponse(w, response, http.StatusBadRequest)
}

// SendSuccessResponse send response with success status.
func SendSuccessResponse(w http.ResponseWriter, m string) {
	response := Response{
		Status:  "Success",
		Message: m,
	}

	SendResponse(w, response, http.StatusOK)
}
