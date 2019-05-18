package persistence

import (
	"github.com/aanciaes/smart-grow-api/config/database"
	"github.com/aanciaes/smart-grow-api/model"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

const (
	insertQuery       = "INSERT INTO users (username, password, isAdmin) VALUES (?, ?, ?)"
	lastTemperature   = "SELECT * FROM temperature_readings ORDER BY dateOf DESC LIMIT ?"
	createTemperature = "INSERT INTO temperature_readings (reading, dateOf) VALUES (?, ?)"
)

func RegisterUser(registerForm model.RegisterForm) error {
	db := database.Conn.Connection

	bytes, err := bcrypt.GenerateFromPassword([]byte(registerForm.Password), 14)
	hashedPassword := string(bytes)
	_, err = db.Exec(insertQuery, registerForm.Username, hashedPassword, registerForm.IsAdmin)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func GetTemperature(numberOfReadings int64) ([]model.TemperatureReading, error) {
	db := database.Conn.Connection

	rows, err := db.Query(lastTemperature, numberOfReadings)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var rst = make([] model.TemperatureReading, 0)

	var (
		id      int
		reading float32
		dateOf  string
	)

	for rows.Next() {
		err = rows.Scan(&id, &reading, &dateOf)
		if err != nil {
			return [] model.TemperatureReading{}, err
		}

		rst = append(rst, model.TemperatureReading{Id:id, Reading:reading, Date:dateOf})
	}

	return rst, nil
}

func CreateTemperatureReading(reading float32) error {
	db := database.Conn.Connection

	_, err := db.Exec(createTemperature, reading, time.Now().String())
	if err != nil {
		return err
	}

	return nil
}
