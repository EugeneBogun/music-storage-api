package components

import (
    "net/http"
    "encoding/json"
)

type Response struct {
    Status string `json:"status"`
    Message string `json:"message"`
}

func SendResponse(w http.ResponseWriter, r interface{}, status int)  {
    jsonMessage, err := json.Marshal(r)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    w.Write(jsonMessage)
}

func SendErrorResponse(w http.ResponseWriter, m string)  {
    response := Response{
        Status: "Error",
        Message: m,
    }
    SendResponse(w, response, http.StatusBadRequest)
}

func SendSuccessResponse(w http.ResponseWriter, m string)  {
    response := Response{
        Status: "Success",
        Message: m,
    }

    SendResponse(w, response, http.StatusOK)
}
