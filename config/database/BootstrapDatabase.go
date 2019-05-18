package database

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"time"
)

func BootstrapDatabase() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = DEV
	}

	db := Conn.Connection

	if env == DEV {
		createTablesForDev(db)

		_, err := db.Exec("insert into temperature_readings (reading, dateOf) values (?, ?)", 23, time.Now().String())
		if err != nil {
			log.Fatal(err)
		}
		_, err = db.Exec("insert into temperature_readings (reading, dateOf) values (?, ?)", 24, time.Now().String())
		if err != nil {
			log.Fatal(err)
		}

		bytes, err := bcrypt.GenerateFromPassword([]byte("admin"), 14)
		hashedPassword := string(bytes)
		_, err = db.Exec("insert into users (username, password, isAdmin) values (?, ?, ?)", "admin", hashedPassword, 1)
		if err != nil {
			log.Fatal(err)
		}

		bytes, err = bcrypt.GenerateFromPassword([]byte("miguel"), 14)
		hashedPassword = string(bytes)
		_, err = db.Exec("insert into users (username, password, isAdmin) values (?, ?, ?)", "miguel", hashedPassword, 0)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		createTablesForProd(db)
	}
}

func createTablesForDev (db *sql.DB) {

	_, err := db.Exec("create table if not exists temperature_readings (id INTEGER PRIMARY KEY AUTOINCREMENT, reading text, dateOf text)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("create table if not exists users (id INTEGER PRIMARY KEY AUTOINCREMENT, username text UNIQUE, password text, isAdmin boolean)")
	if err != nil {
		log.Fatal(err)
	}
}

func createTablesForProd (db *sql.DB) {

	_, err := db.Exec("create table if not exists temperature_readings (id INTEGER PRIMARY KEY AUTO_INCREMENT, reading text, dateOf text)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("create table if not exists users (id INTEGER PRIMARY KEY AUTO_INCREMENT, username VARCHAR (255) UNIQUE, password text, isAdmin boolean)")
	if err != nil {
		log.Fatal(err)
	}
}