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
//	@Description	Registers user and creates permission group for him.
//	@Tags			users
//	@Accept			json
//	@Param			user	body		requests.RegisterRequest	true	"user id"
//	@Success		201		{object}	responses.RegisterResponse	"user registered"
//	@Failure		400		{string}	string						"test"
//	@Router			/users/register [post]
func (c *UserController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	c.logger.Info("UserController.RegisterUser")

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

	bank, err := c.bankRepo.GetByName(req.PreferredBankName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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

	// sets password to empty to avoid sending password hash
	newUser.Password = ""

	// creating permission group for user
	permissionGroup, err := c.permissionGroupRepo.Create(r.Context(), &models.PermissionGroup{
		Name: req.Username,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userPermission, err := c.userPermissionRepo.Create(r.Context(), &models.UserPermission{
		PermissionGroupId: permissionGroup.Id,
		UserId:            newUser.Id,
		Level:             models.AccessLevelOwner,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	wallet, err := c.walletRepo.Create(r.Context(), &models.Wallet{
		Name:              req.Username + "_salary",
		Description:       "Salary wallet",
		PermissionGroupId: permissionGroup.Id,
		Currency:          models.CurrencyBYN,
		BankId:            bank.Id,
		IsSalary:          true,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := responses.RegisterResponse{
		User:            newUser,
		PermissionGroup: permissionGroup,
		UserPermission:  userPermission,
		Wallet:          wallet,
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(&resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
