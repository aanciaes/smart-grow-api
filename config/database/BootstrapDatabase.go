package database

import (
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
