package db

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

var Connection *sql.DB

func InitConnection(user, password, address, name string) {
	cfg := mysql.Config{
		User:      user,
		Passwd:    password,
		Net:       "tcp",
		Addr:      address,
		DBName:    name,
		ParseTime: true,
	}

	Connection, _ = sql.Open("mysql", cfg.FormatDSN())
}

func CloseConnection() {
	Connection.Close()
}
