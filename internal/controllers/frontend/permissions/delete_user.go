package permissions

import (
	"context"

	"go.uber.org/zap"
)

func (c *controller) DeleteUser(ctx context.Context, username string, permissionGroupID int64) error {
	c.logger.Debug("frontend.add_user_to_permission_group.controller.form", zap.String("event", "got request"))

	return c.userPermissionRepo.Delete(ctx, username, permissionGroupID)
}
