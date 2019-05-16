package database

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
)

func BootstrapDatabase() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = DEV
	}

	db := Conn.Connection

	if env == DEV {
		createTablesForDev(db)

		_, err := db.Exec("insert into temperature_readings (reading) values (?)", 23)
		if err != nil {
			log.Fatal(err)
		}
		_, err = db.Exec("insert into temperature_readings (reading) values (?)", 24)
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
	/*rows, err := db.Query("select * from temperature_readings")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var (
		id int
		reading string
	)
	for rows.Next(){
		err:=rows.Scan(&id, &reading)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id)
		fmt.Println(reading)
	}*/
}

func createTablesForDev (db *sql.DB) {

	_, err := db.Exec("create table if not exists temperature_readings (id INTEGER PRIMARY KEY AUTOINCREMENT, reading text)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("create table if not exists users (id INTEGER PRIMARY KEY AUTOINCREMENT, username text UNIQUE, password text, isAdmin boolean)")
	if err != nil {
		log.Fatal(err)
	}
}

func createTablesForProd (db *sql.DB) {

	_, err := db.Exec("create table if not exists temperature_readings (id INTEGER PRIMARY KEY AUTO_INCREMENT, reading text)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("create table if not exists users (id INTEGER PRIMARY KEY AUTO_INCREMENT, username text UNIQUE, password text, isAdmin boolean)")
	if err != nil {
		log.Fatal(err)
	}
}