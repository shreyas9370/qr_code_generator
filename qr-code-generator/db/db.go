package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Init() {
	var err error

	DB, err = sql.Open("mysql", "root:Shreyas9370@@tcp(localhost:3306)/qrcode")
	if err != nil {
		log.Fatal(err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal(err)
	}
}
