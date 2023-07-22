package rest

import (
	"github.com/babenkoivan/todo/internal/users"
	"github.com/gin-gonic/gin"
)

const userContextKey = "user"

func setUser(ctx *gin.Context, user *users.User) {
	ctx.Set(userContextKey, user)
}

func getUser(ctx *gin.Context) *users.User {
	return ctx.MustGet(userContextKey).(*users.User)
}
