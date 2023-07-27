package graphql

import (
	"context"
	"github.com/babenkoivan/todo/internal/users"
)

const userContextKey = "user"

func setUser(ctx context.Context, user *users.User) context.Context {
	return context.WithValue(ctx, userContextKey, user)
}

func getUser(ctx context.Context) (*users.User, bool) {
	user, ok := ctx.Value(userContextKey).(*users.User)
	return user, ok
}
