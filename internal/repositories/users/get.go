package users

import (
	"context"

	"finworker/internal/models"
)

func (repo *Repository) Get(ctx context.Context, id int) (models.User, error) {
	var user models.User

	return user, repo.db.GetContext(ctx, &user, `SELECT * FROM users WHERE id = $1`, id)
}
