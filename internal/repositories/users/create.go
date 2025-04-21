package users

import (
	"context"

	"finworker/internal/models"
)

func (repo *Repository) Create(ctx context.Context, user *models.User) (*models.User, error) {
	q := `INSERT INTO users (username, password_hash, name, gender, birthday) VALUES (:username, :password_hash, :name, :gender, :birthday) returning id`

	namedStmt, err := repo.db.PrepareNamed(q)
	if err != nil {
		return nil, err
	}

	err = namedStmt.GetContext(ctx, &user.ID, user)

	return user, err
}
