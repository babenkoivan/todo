package main

import (
	"fmt"
	"github.com/babenkoivan/todo/internal/db"
	"github.com/babenkoivan/todo/internal/rest"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	db.InitConnection(
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
		os.Getenv("DB_NAME"),
	)
	defer db.CloseConnection()

	router := gin.Default()
	router.POST("/api/users", rest.CreateUser)
	router.POST("/api/users/login", rest.LoginUser)
	router.POST("/api/todos", rest.Authenticate, rest.CreateTodo)
	router.GET("/api/todos", rest.Authenticate, rest.GetTodos)
	router.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}
