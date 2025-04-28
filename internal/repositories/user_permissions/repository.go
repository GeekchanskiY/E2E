package user_permissions

import (
	"context"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	"finworker/internal/models"
)

type Repository interface {
	Create(ctx context.Context, permission *models.UserPermission) (*models.UserPermission, error)
	Delete(ctx context.Context, username string, permissionGroupID int64) error
	GetForGroup(ctx context.Context, groupID int64) ([]*models.UserPermissionExtended, error)
}

type repository struct {
	db  *sqlx.DB
	log *zap.Logger
}

func New(db *sqlx.DB, log *zap.Logger) Repository {
	return &repository{
		log: log,
		db:  db,
	}
}
