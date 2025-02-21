package repositories

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

func (repo *UserRepository) Create(ctx context.Context, user *models.User) (*models.User, error) {
	q := `INSERT INTO users (username, password_hash, name, gender, birthday) VALUES (:username, :password_hash, :name, :gender, :birthday) returning id`
	namedStmt, err := repo.db.PrepareNamed(q)
	if err != nil {
		return nil, err
	}

	err = namedStmt.GetContext(ctx, &user.Id, user)

	return user, err
}

func (repo *UserRepository) Get(ctx context.Context, username string) (models.User, error) {
	var user models.User
	err := repo.db.GetContext(ctx, &user, `SELECT * FROM users WHERE username = $1`, username)
	return user, err
}

func (repo *UserRepository) GetById(ctx context.Context, id int) (models.User, error) {
	var user models.User
	err := repo.db.GetContext(ctx, &user, `SELECT * FROM users WHERE id = $1`, id)
	return user, err
}
