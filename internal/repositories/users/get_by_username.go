package users

import (
	"context"

	"finworker/internal/models"
)

func (repo *Repository) GetByUsername(ctx context.Context, username string) (models.User, error) {
	var user models.User
	err := repo.db.GetContext(ctx, &user, `SELECT * FROM users WHERE username = $1`, username)
	return user, err
}
