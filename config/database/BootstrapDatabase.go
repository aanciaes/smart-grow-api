package database

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func BootstrapDatabase() {
	db := Conn.Connection

	_, err := db.Exec("create table temperature_readings (id INTEGER PRIMARY KEY, reading varchar)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("insert into temperature_readings (reading) values (?)", 23)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("insert into temperature_readings (reading) values (?)", 24)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("create table users (id INTEGER PRIMARY KEY, username varchar, password varchar, isAdmin number)")
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
