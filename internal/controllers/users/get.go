package users

import (
	"context"

	"finworker/internal/models"
)

func (c *Controller) GetUser(ctx context.Context, userId int) (*models.User, error) {
	user, err := c.userRepo.Get(ctx, userId)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
