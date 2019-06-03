package persistence

import (
	"fmt"
	"github.com/aanciaes/smart-grow-api/config/database"
	"github.com/aanciaes/smart-grow-api/model"
	"golang.org/x/crypto/bcrypt"
	"log"
	"sort"
	"strconv"
	"time"
)

const (
	insertQuery       = "INSERT INTO users (username, password, isAdmin) VALUES (?, ?, ?)"
	lastTemperatureDesc   = "SELECT * FROM temperature_readings ORDER BY dateOf DESC LIMIT ?"
	lastHumidityDesc   = "SELECT * FROM humidity_readings ORDER BY dateOf DESC LIMIT ?"
	lastLightDesc   = "SELECT * FROM light_readings ORDER BY dateOf DESC LIMIT ?"
	lastSoilDesc   = "SELECT * FROM soil_readings ORDER BY dateOf DESC LIMIT ?"
	createTemperature = "INSERT INTO temperature_readings (reading, dateOf) VALUES (?, ?)"
	createHumidity = "INSERT INTO humidity_readings (reading, dateOf) VALUES (?, ?)"
	createLight = "INSERT INTO light_readings (reading, dateOf) VALUES (?, ?)"
	createSoil = "INSERT INTO soil_readings (reading, dateOf) VALUES (?, ?)"

	routineChecker = "SELECT * FROM routines WHERE datetime < ?"
	createRoutine =  "INSERT INTO routines (motor, datetime) VALUES (?, ?)"
	deleteRoutine = "DELETE FROM routines WHERE id = ?"
	getRoutines = "SELECT * FROM routines"
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

func GetTemperature(numberOfReadings int64, asc bool) ([]model.Readings, error) {
	db := database.Conn.Connection

	rows, err := db.Query(lastTemperatureDesc, numberOfReadings)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var rst = make([] model.Readings, 0)

	var (
		id      int
		reading float32
		dateOf  string
	)

	for rows.Next() {
		err = rows.Scan(&id, &reading, &dateOf)
		if err != nil {
			return [] model.Readings{}, err
		}

		rst = append(rst, model.Readings{Id:id, Reading:reading, Date:dateOf})
	}

	if asc {
		sort.Slice(rst, func(i, j int) bool {
			return rst[i].Date < rst[j].Date
		})

		return rst, nil
	}
	return rst, nil
}

func GetHumidity(numberOfReadings int64, asc bool) ([]model.Readings, error) {
	db := database.Conn.Connection

	rows, err := db.Query(lastHumidityDesc, numberOfReadings)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var rst = make([] model.Readings, 0)

	var (
		id      int
		reading float32
		dateOf  string
	)

	for rows.Next() {
		err = rows.Scan(&id, &reading, &dateOf)
		if err != nil {
			return [] model.Readings{}, err
		}

		rst = append(rst, model.Readings{Id:id, Reading:reading, Date:dateOf})
	}

	if asc {
		sort.Slice(rst, func(i, j int) bool {
			return rst[i].Date < rst[j].Date
		})

		return rst, nil
	}
	return rst, nil
}

func GetLight(numberOfReadings int64, asc bool) ([]model.Readings, error) {
	db := database.Conn.Connection

	rows, err := db.Query(lastLightDesc, numberOfReadings)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var rst = make([] model.Readings, 0)

	var (
		id      int
		reading float32
		dateOf  string
	)

	for rows.Next() {
		err = rows.Scan(&id, &reading, &dateOf)
		if err != nil {
			return [] model.Readings{}, err
		}

		rst = append(rst, model.Readings{Id:id, Reading:reading, Date:dateOf})
	}

	if asc {
		sort.Slice(rst, func(i, j int) bool {
			return rst[i].Date < rst[j].Date
		})

		return rst, nil
	}
	return rst, nil
}

func GetSoil(numberOfReadings int64, asc bool) ([]model.Readings, error) {
	db := database.Conn.Connection

	rows, err := db.Query(lastSoilDesc, numberOfReadings)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var rst = make([] model.Readings, 0)

	var (
		id      int
		reading float32
		dateOf  string
	)

	for rows.Next() {
		err = rows.Scan(&id, &reading, &dateOf)
		if err != nil {
			return [] model.Readings{}, err
		}

		rst = append(rst, model.Readings{Id:id, Reading:reading, Date:dateOf})
	}

	if asc {
		sort.Slice(rst, func(i, j int) bool {
			return rst[i].Date < rst[j].Date
		})

		return rst, nil
	}
	return rst, nil
}

func CreateTemperatureReading(reading float32) error {
	db := database.Conn.Connection

	loc, _ := time.LoadLocation("Europe/Lisbon")
	_, err := db.Exec(createTemperature, reading, time.Now().In(loc).Unix())
	if err != nil {
		return err
	}

	return nil
}

func CreateHumidityReading(reading float32) error {
	db := database.Conn.Connection

	loc, _ := time.LoadLocation("Europe/Lisbon")
	_, err := db.Exec(createHumidity, reading, time.Now().In(loc).Unix())
	if err != nil {
		return err
	}

	return nil
}

func CreateLightReading(reading float32) error {
	db := database.Conn.Connection

	loc, _ := time.LoadLocation("Europe/Lisbon")
	_, err := db.Exec(createLight, reading, time.Now().In(loc).Unix())
	if err != nil {
		return err
	}

	return nil
}

func CreateSoilReading(reading float32) error {
	db := database.Conn.Connection

	loc, _ := time.LoadLocation("Europe/Lisbon")
	_, err := db.Exec(createSoil, reading, time.Now().In(loc).Unix())
	if err != nil {
		return err
	}

	return nil
}

func checkRoutines () (model.Routine, error) {
	db := database.Conn.Connection

	loc, _ := time.LoadLocation("Europe/Lisbon")
	rows, err := db.Query(routineChecker, time.Now().In(loc).Format("2006-01-02 03:04:05"))
	if err != nil {
		return model.Routine{}, err
	}

	var (
		id      int
		datetime string
		output  string
	)

	defer rows.Close()
	rows.Next()

	err = rows.Scan(&id, &output, &datetime)
	if err != nil {
		return model.Routine{}, err
	}

	return model.Routine{Id:id, Datetime:datetime, Output:output}, nil
}

func CreateRoutine (routine model.RoutineForm) error {
	db := database.Conn.Connection

	datetimeString,_ := strconv.ParseInt(routine.Datetime, 10, 64)
	datetime := time.Unix(datetimeString, 0)

	loc, _ := time.LoadLocation("Europe/Lisbon")
	_, err := db.Exec(createRoutine, routine.Output, datetime.In(loc).Format("2006-01-02 03:04:05"))
	if err != nil {
		return err
	}

	return nil
}

func DeleteRoutine (id int) {
	db := database.Conn.Connection
	db.Exec(deleteRoutine, id)
}

func CheckRoutines () {
	for true {
		time.Sleep(10 * time.Second)

		routine, err := checkRoutines()

		if err != nil {
			//fmt.Printf("Err: %s\n", err)
		} else {
			fmt.Printf("Routine for motor: %s with time: %s\n", routine.Output, routine.Datetime)
			DeleteRoutine(routine.Id)

			//TODO: Request to arduino
		}
	}
}

func GetRoutines () ([]model.Routine, error) {
	db := database.Conn.Connection

	rows, err := db.Query(getRoutines)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var rst = make([] model.Routine, 0)

	var (
		id      int
		output string
		datetime  string
	)

	for rows.Next() {
		err = rows.Scan(&id, &output, &datetime)
		if err != nil {
			return [] model.Routine{}, err
		}

		rst = append(rst, model.Routine{Id:id, Output:output, Datetime:datetime})
	}

	return rst, nil
}