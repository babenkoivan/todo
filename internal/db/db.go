package db

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"os"
)

var Connection *sql.DB

func InitConnection() {
	cfg := mysql.Config{
		User:      os.Getenv("DB_USER"),
		Passwd:    os.Getenv("DB_PASSWORD"),
		Net:       "tcp",
		Addr:      fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
		DBName:    os.Getenv("DB_NAME"),
		ParseTime: true,
	}

	Connection, _ = sql.Open("mysql", cfg.FormatDSN())
}

func CloseConnection() {
	Connection.Close()
}
