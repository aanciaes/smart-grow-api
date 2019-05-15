package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

const (
	DEV = "dev"
	PROD = "prod"
)

type Database struct {
	Connection *sql.DB
}

var Conn = Database{nil}

func ConfigDatabase () (*Database, error) {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = DEV
	}

	// Return database configuration based on environment variable
	if env == DEV {
		return configDevDatabase()
	} else {
		return configProdDatabase()
	}
}

func configProdDatabase () (*Database, error){
	log.Printf("Starting configuration for %s environment", PROD)

	db, _ := sql.Open("mysql", "root:smartgrow@tcp(mysql-db:3306)/smartgrow")

	// Checks for database connection errors since sql open will only validate arguments
	// without actually creating a connection
	pong := db.Ping()
	if pong != nil {
		return nil, pong
	}

	Conn.Connection = db

	return &Database{db}, nil
}


func configDevDatabase () (*Database, error) {
	log.Printf("Starting configuration for %s environment", DEV)

	db, _ := sql.Open("sqlite3", ":memory:")

	// Checks for database connection errors since sql open will only validate arguments
	// without actually creating a connection
	pong := db.Ping()
	if pong != nil {
		return nil, pong
	}

	Conn.Connection = db

	return &Database{db}, nil
}