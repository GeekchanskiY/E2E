package utils

import (
	"context"
)

const (
	userKey = "user"
)

func BuildDefaultDataMapFromContext(ctx context.Context) map[string]interface{} {
	var data = make(map[string]interface{})

	user := ctx.Value("user")
	if user != nil {
		data[userKey] = user
	}

	return data
}
