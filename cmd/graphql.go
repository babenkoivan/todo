package main

import (
	"github.com/babenkoivan/todo/internal/db"
	"github.com/babenkoivan/todo/internal/graphql"
	"net/http"
)

func main() {
	db.InitConnection()
	defer db.CloseConnection()

	http.HandleFunc("/api", graphql.ProcessRequest)
	http.ListenAndServe(":8080", nil)
}
