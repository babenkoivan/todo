package main

import (
	"github.com/babenkoivan/todo/internal/db"
	"github.com/babenkoivan/todo/internal/rest"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitConnection()
	defer db.CloseConnection()

	router := gin.Default()
	router.POST("/api/users", rest.CreateUser)
	router.POST("/api/users/login", rest.LoginUser)
	router.POST("/api/todos", rest.Authenticate, rest.CreateTodo)
	router.GET("/api/todos", rest.Authenticate, rest.GetTodos)
	router.Run()
}
