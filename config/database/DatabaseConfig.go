package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const (
	DEV = "dev"
	STAGING = "staging"
)

type Database struct {
	connection *sql.DB
}

func ConfigDatabase () (*sql.DB, error) {
	//env := os.Getenv("LM_API_ENV")
	// Return database configuration based on environment variable

	return configStagingDatabase()
}

func configStagingDatabase () (*sql.DB, error){
	log.Printf("Starting configuration for %s environment", STAGING)

	db, _ := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/league_manager")

	// Checks for database connection errors since sql open will only validate arguments
	// without actually creating a connection
	pong := db.Ping()
	if pong != nil {
		return nil, pong
	}

	return db, nil
}

/*
func configDevDatabase () (Database, error) {
	log.Printf("Starting configuration for %s environment", DEV)
	conn := Database{}
	db, err := bolt.Open("league_manager.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
		return Database{}, errors.New("database connection failed")
	} else {
		conn.connection = db
		return conn, nil
	}
}*/