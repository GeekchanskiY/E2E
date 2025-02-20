package users

import (
	"encoding/json"
	"net/http"

	"finworker/internal/models"
	requests "finworker/internal/models/requests/users"
	responses "finworker/internal/models/responses/users"
	"finworker/internal/utils"
)

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
func (c *UserController) RegisterUser(w http.ResponseWriter, r *http.Request) {
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

	newUser := &models.User{
		Username: req.Username,
		Password: password,
		Name:     req.Name,
		Gender:   req.Gender,
		Birthday: req.Birthday,
	}

	newUser, err = c.userRepo.Create(r.Context(), newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newUser.Password = ""

	resp := responses.RegisterResponse{
		User: newUser,
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(&resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
