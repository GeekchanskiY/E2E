package users

import (
	"time"
)

type RegisterRequest struct {
	// Username, which will be displayed and unique.
	Username string `json:"username"`

	// Password, which will be used with secret
	Password string `json:"password"`

	// Name. just to be displayed
	Name string `json:"name"`

	// Gender. `male`/`female`.
	Gender string `json:"gender"`

	// Birthday. Age must be > 18.
	Birthday time.Time `json:"birthday"`
}
