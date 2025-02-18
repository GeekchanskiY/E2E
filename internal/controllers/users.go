package controllers

import (
	"encoding/json"
	"net/http"
)

func (c *Controller) GetUser(w http.ResponseWriter, r *http.Request) {
	user, err := c.repo.Users.GetById(r.Context(), 1)
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
