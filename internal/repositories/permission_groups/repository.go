package permission_groups

import (
	"context"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	"finworker/internal/models"
)

type Repository interface {
	Create(ctx context.Context, group *models.PermissionGroup) (*models.PermissionGroup, error)
	GetByID(ctx context.Context, id int64) (*models.PermissionGroup, error)
	GetByName(ctx context.Context, name string) (*models.PermissionGroup, error)
	GetUserEditGroups(ctx context.Context, userID int64) (permissionGroups []*models.PermissionGroup, err error)
	GetUserGroups(ctx context.Context, userID int64) (permissionGroups []*models.PermissionGroupWithRole, err error)
}

type repository struct {
	log *zap.Logger

	db *sqlx.DB
}

func New(db *sqlx.DB, log *zap.Logger) Repository {
	return &repository{
		log: log,
		db:  db,
	}
}
