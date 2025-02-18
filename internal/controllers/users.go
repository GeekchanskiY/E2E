package controllers

import (
	"encoding/json"
	"net/http"

	"finworker/internal/models/requests/users"
)

func (c *Controller) GetUser(w http.ResponseWriter, r *http.Request) {
	var req users.GetUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if req.Username == nil && req.Id == nil {
		http.Error(w, "Required fields missing", http.StatusBadRequest)
		return
	}
}
