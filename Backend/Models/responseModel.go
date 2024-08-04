package models

import (
	"encoding/json"
	"net/http"
)

/* -------------------------------- Types ----------------------------------*/

// Response model
type Response struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data,omitempty"`
	Error   string                 `json:"error,omitempty"`
}

/* -------------------------------- Helpers ----------------------------------*/

// Crafting a successful JSON response based on status code, message, and response data
func (res *Response) OkResponse(w http.ResponseWriter, status int, message string, data map[string]interface{}) {
	res.Status = status
	res.Message = message
	res.Data = data

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(res)
}

// Crafting a failed JSON response based on status code, message, and error
func (res *Response) BadResponse(w http.ResponseWriter, status int, message string, err error) {
	res.Status = status
	res.Message = message
	res.Error = err.Error()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(res)
}
