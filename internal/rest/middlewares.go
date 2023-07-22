package rest

import (
	"github.com/babenkoivan/todo/internal/users"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authenticate(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	user, err := users.Authenticate(token)

	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	setUser(ctx, user)
	ctx.Next()
}
