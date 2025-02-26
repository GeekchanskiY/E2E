package users

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// Get godoc
//
//	@Summary		Get user
//	@Description	get user by user id
//	@Tags			users
//	@Accept			json
//	@Param			userId	path	int	true	"user id"
//	@Success		200
//	@Router			/users/{userId} [get]
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
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

	user, err := h.controller.GetUser(ctx, userIdInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}
