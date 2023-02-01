package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Connect(user, password, name string) {
	datasource := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", user, password, name)

	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", datasource)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	log.Println("DB connected!")

	DBInit()

}

func GetRecords() {
	// TODO
}
