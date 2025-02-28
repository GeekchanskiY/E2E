package users

import (
	"errors"
	"time"
	"unicode"
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

	// PreferredBankName is used to created initial salary wallet
	PreferredBankName string `json:"preferred_bank_name"`

	// Salary may be 0, then you'll need to manually set your salary every month
	Salary int `json:"salary"`

	// SalaryCurrency
	SalaryCurrency string `json:"salary_currency"`

	// SalaryDate may be zero, then you'll need to confirm achieving your salary
	SalaryDate time.Time `json:"salary_date"`
}

func (req *RegisterRequest) Validate() error {
	var errs []error
	if req.Username == "" {
		errs = append(errs, errors.New("username is required"))
	}

	if req.Gender != "male" && req.Gender != "female" {
		errs = append(errs, errors.New("gender must be either male or female"))
	}

	if req.Birthday.IsZero() {
		errs = append(errs, errors.New("birthday is required"))
	}

	if time.Now().Sub(req.Birthday).Hours()/24/365.25 < 18 {
		errs = append(errs, errors.New("birthday must be older than 18 years"))
	}

	if req.Password == "" {
		errs = append(errs, errors.New("password is required"))
	}

	if len(req.Password) < 8 {
		errs = append(errs, errors.New("password must be at least 8 characters"))
	}

	uppercaseFound := false
	for _, s := range req.Password {
		if unicode.IsUpper(s) {
			uppercaseFound = true
		}
	}

	if !uppercaseFound {
		errs = append(errs, errors.New("password must contain at least one uppercase character"))
	}

	return errors.Join(errs...)
}
