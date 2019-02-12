package models

import (
	"fmt"
	"log"
	"database/sql"
	"os"
	_ "github.com/lib/pq"
)

const (
	USER    = "postgres"
	PASS    = "@root"
	DBNAME  = "blockcoin"
	SSLMODE = "disable"
)

const DEVELOPMENT = true

func Connect() *sql.DB {
	var db *sql.DB = nil
	var err error = nil
	if DEVELOPMENT {
		URL := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", USER, PASS, DBNAME, SSLMODE)
		db, err = sql.Open("postgres", URL)
	} else {
		db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	}
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}
	return db
}

func TestConnection() {
	con := Connect()
	defer con.Close()
	err := con.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Database connected!")
}