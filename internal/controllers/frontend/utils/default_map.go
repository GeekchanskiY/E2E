package utils

import (
	"context"

	"finworker/internal/config"
)

const (
	userKey = "username"
)

func BuildDefaultDataMapFromContext(ctx context.Context) map[string]interface{} {
	var data = make(map[string]interface{})

	user := ctx.Value(config.UsernameContextKey)
	if user != nil {
		data[userKey] = user
	}

	return data
}
