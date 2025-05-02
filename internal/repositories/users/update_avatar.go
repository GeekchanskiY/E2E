package users

import (
	"context"
)

// UpdateAvatar returns a previous avatar path if existed
func (repo *repository) UpdateAvatar(ctx context.Context, userID int64, avatar string) (string, error) {
	var (
		previous string

		err error
	)

	err = repo.db.GetContext(ctx, &previous, `SELECT avatar FROM users WHERE id = $1`, userID)
	if err != nil {
		return "", err
	}

	_, err = repo.db.ExecContext(ctx, `UPDATE users SET avatar = $1 WHERE id = $2 RETURNING avatar`, avatar, userID)
	return previous, err
}
