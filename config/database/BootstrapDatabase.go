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
		createEnvironmentForDev(db)
	} else {
		createEnvironmentForProd(db)
	}
}

func createEnvironmentForDev(db *sql.DB) {
	_, err := db.Exec("DROP DATABASE IF EXISTS smartgrow")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("CREATE DATABASE smartgrow")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("USE smartgrow")
	if err != nil {
		log.Fatal(err)
	}

	// Creates all necessary tables
	createTables(db)

	_, err = db.Exec("insert into temperature_readings (reading, dateOf) values (?, ?)", 23, time.Now().Unix())
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("insert into temperature_readings (reading, dateOf) values (?, ?)", 24, time.Now().Unix())
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

	/*_, err = db.Exec("insert into routines (motor, datetime) values (?, ?)", "water", time.Now().UTC().Format("2006-01-02 03:04:05"))
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("insert into routines (motor, datetime) values (?, ?)", "water", time.Now().Add(time.Minute * 1).UTC().Format("2006-01-02 03:04:05"))
	if err != nil {
		log.Fatal(err)
	}*/
}

func createEnvironmentForProd(db *sql.DB) {
	createTables(db)

	//creates admin user
	bytes, err := bcrypt.GenerateFromPassword([]byte("admin"), 14)
	hashedPassword := string(bytes)
	_, err = db.Exec("insert into users (username, password, isAdmin) values (?, ?, ?)", "admin", hashedPassword, 1)
	if err != nil {
		log.Fatal(err)
	}
}

func createTables(db *sql.DB) {
	_, err := db.Exec("create table if not exists temperature_readings (id INTEGER PRIMARY KEY AUTO_INCREMENT, reading text, dateOf text)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("create table if not exists users (id INTEGER PRIMARY KEY AUTO_INCREMENT, username VARCHAR (255) UNIQUE, password text, isAdmin boolean)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("create table if not exists humidity_readings (id INTEGER PRIMARY KEY AUTO_INCREMENT, reading text, dateOf text)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("create table if not exists light_readings (id INTEGER PRIMARY KEY AUTO_INCREMENT, reading text, dateOf text)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("create table if not exists soil_readings (id INTEGER PRIMARY KEY AUTO_INCREMENT, reading text, dateOf text)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("create table if not exists routines (id INTEGER PRIMARY KEY AUTO_INCREMENT, motor text, datetime timestamp)")
	if err != nil {
		log.Fatal(err)
	}
}
