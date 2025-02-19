package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"finworker/internal/models/requests/users"
	"github.com/go-chi/chi/v5"
)

// GetUser godoc
//
//	@Summary		Get user
//	@Description	get user by user id
//	@Tags			users
//	@Accept			json
//	@Param			userId	path int	true	"user id"
//	@Success		200
//	@Router			/users/{userId} [get]
func (c *Controller) GetUser(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")
	if userId == "" {
		http.Error(w, "User Id is required", http.StatusBadRequest)

		return
	}

	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	user, err := c.repo.Users.GetById(r.Context(), userIdInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}

// RegisterUser godoc
//
//	@Summary		Register user
//	@Description	register user
//	@Tags			users
//	@Accept			json
//	@Param			user	body users.RegisterRequest	true	"user id"
//	@Success		200
//	@Router			/users/register [post]
func (c *Controller) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var req users.RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
