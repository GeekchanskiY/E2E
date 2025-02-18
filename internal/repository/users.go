package repository

import (
	"context"

	"finworker/internal/models"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) Create(ctx context.Context, user *models.User) error {
	_, err := repo.db.NamedExecContext(
		ctx,
		`INSERT INTO users (username, password_hash, name, gender, age, birthday) VALUES (:username, :password_hash, :name, :gender, :age, :birthday)`,
		user)
	return err
}
