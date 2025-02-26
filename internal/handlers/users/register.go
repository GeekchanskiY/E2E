package users

import (
	"encoding/json"
	"net/http"

	requests "finworker/internal/models/requests/users"
)

// Register godoc
//
//	@Summary		Register user
//	@Description	Registers user and creates permission group for him.
//	@Tags			users
//	@Accept			json
//	@Param			user	body		requests.RegisterRequest	true	"user id"
//	@Success		201		{object}	responses.RegisterResponse	"user registered"
//	@Failure		400		{string}	string						"test"
//	@Router			/users/register [post]
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var req requests.RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := h.controller.RegisterUser(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
