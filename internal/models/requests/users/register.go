package users

import (
	"time"
)

type RegisterRequest struct {
	Username string    `json:"username"`
	Password string    `json:"password"`
	Name     string    `json:"name"`
	Gender   string    `json:"gender"`
	Birthday time.Time `json:"birthday"`
	Secret   string    `json:"secret"`
}
