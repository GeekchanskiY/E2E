package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"finworker/internal/models"
	requests "finworker/internal/models/requests/users"
	responses "finworker/internal/models/responses/users"
	"finworker/internal/utils"
)

// GetUser godoc
//
//	@Summary		Get user
//	@Description	get user by user id
//	@Tags			users
//	@Accept			json
//	@Param			userId	path	int	true	"user id"
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
//	@Param			user	body		requests.RegisterRequest	true	"user id"
//	@Success		201		{object}	responses.RegisterResponse	"user registered"
//	@Failure		400		{string}	string						"test"
//	@Router			/users/register [post]
func (c *Controller) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var req requests.RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = req.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	password, err := utils.HashPassword(req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newUser := models.User{
		Username: req.Username,
		Password: password,
		Name:     req.Name,
		Gender:   req.Gender,
		Birthday: req.Birthday,
	}

	newId, err := c.repo.Users.Create(r.Context(), newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newUser.Id = newId
	newUser.Password = ""

	resp := responses.RegisterResponse{
		User: &newUser,
	}
	err = json.NewEncoder(w).Encode(&resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
