package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var conn *sql.DB

func GetInstance() *sql.DB {
	if conn == nil {
		dat := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?parseTime=true",
			os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"), os.Getenv("DB_NAME"),
		)
		conn, err := sql.Open("mysql", dat)
		if err != nil {
			log.Fatal("failed: ", err.Error())
		}
		return conn
	}
	return conn
}
