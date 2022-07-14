package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	//"github.com/mercadolibre/go-meli-toolkit/gomelipass"
)

var conn *sql.DB

func GetInstance() *sql.DB {
	if conn == nil {
		var dbHost string
		var connectionString string
		//linhas 17 a 21: trocar x pelo seu numero de feature ou copiar e colar direto dos snippets
		// dbUsername := "bgow1s48x_WPROD"
		// dbPassword := gomelipass.GetEnv("DB_MYSQL_DESAENV05_BGOW1S48x_BGOW1S48x_WPROD")
		// dbHost = gomelipass.GetEnv("DB_MYSQL_DESAENV05_BGOW1S48x_BGOW1S48x_ENDPOINT")
		// dbName := "bgow1s48x"
		// connectionString = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", dbUsername, dbPassword, dbHost, dbName)
		if dbHost == "" {
			connectionString = fmt.Sprintf(
				"%s:%s@tcp(%s:%s)/%s?charset=utf8",
				os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"),
				os.Getenv("DB_PORT"), os.Getenv("DB_NAME"),
			)
		}
		conn, err := sql.Open("mysql", connectionString)
		if err != nil {
			log.Fatal(err)
		}
		return conn
	}
	return conn
}
