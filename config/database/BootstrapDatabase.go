package database

import (
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


	_, err := db.Exec("create table if not exists temperature_readings (id INTEGER AUTO_INCREMENT PRIMARY KEY, reading text)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("create table if not exists users (id INTEGER AUTO_INCREMENT PRIMARY KEY, username text, password text, isAdmin BIT)")
	if err != nil {
		log.Fatal(err)
	}

	if env == DEV {
		_, err = db.Exec("insert into temperature_readings (reading) values (?)", 23)
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
