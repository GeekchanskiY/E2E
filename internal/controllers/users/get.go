package users

import (
	"context"

	"finworker/internal/models"
)

func (c *Controller) GetUser(ctx context.Context, userID int) (*models.User, error) {
	user, err := c.userRepo.Get(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
