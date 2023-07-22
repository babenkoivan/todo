package rest

import (
	"github.com/babenkoivan/todo/internal/todos"
	"github.com/babenkoivan/todo/internal/users"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(ctx *gin.Context) {
	var user *users.User

	if err := ctx.BindJSON(&user); err != nil {
		return
	}

	user.Create()
	ctx.JSON(http.StatusCreated, user)
}

func LoginUser(ctx *gin.Context) {
	var user *users.User

	if err := ctx.BindJSON(&user); err != nil {
		return
	}

	token, err := user.Login()

	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func CreateTodo(ctx *gin.Context) {
	var todo *todos.Todo

	if err := ctx.BindJSON(&todo); err != nil {
		return
	}

	todo.User = getUser(ctx)
	todo.Create()
	ctx.JSON(http.StatusOK, todo)
}

func GetTodos(ctx *gin.Context) {
	user := getUser(ctx)
	ctx.JSON(http.StatusOK, todos.GetAll(user))
}
