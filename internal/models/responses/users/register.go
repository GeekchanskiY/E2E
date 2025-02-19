package users

import (
	"finworker/internal/models"
)

type RegisterResponse struct {
	User  *models.User `json:"user,omitempty"`
	Error string       `json:"error,omitempty"`
}
