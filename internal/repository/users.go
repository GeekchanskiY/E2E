package repository

import (
	"context"

	"github.com/jmoiron/sqlx"

	"finworker/internal/models"
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

func (repo *UserRepository) Get(ctx context.Context, username string) (*models.User, error) {
	var user *models.User
	err := repo.db.SelectContext(ctx, user, `SELECT * FROM users WHERE username = ?`, username)
	return user, err
}

func (repo *UserRepository) GetById(ctx context.Context, id int) (*models.User, error) {
	var user *models.User
	err := repo.db.SelectContext(ctx, user, `SELECT * FROM users WHERE id = ?`, id)
	return user, err
}
