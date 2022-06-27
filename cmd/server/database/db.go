package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func GetInstance() *sql.DB {
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
