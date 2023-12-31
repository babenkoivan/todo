package main

import (
	"fmt"
	"github.com/babenkoivan/todo/internal/db"
	"github.com/babenkoivan/todo/internal/graphql"
	"net/http"
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

	http.HandleFunc("/api", graphql.ProcessRequest)
	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("APP_PORT")), nil)
}
