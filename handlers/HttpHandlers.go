package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/aanciaes/smart-grow-api/model"
	"github.com/aanciaes/smart-grow-api/persistence"
	"github.com/aanciaes/smart-grow-api/security"
	"net/http"
	"strconv"
)

func Login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var login model.LoginForm
	err := decoder.Decode(&login)

	if err == nil {
		user, authorized := security.CheckPasswordHash(login.Username, login.Password)
		if authorized {
			token, err := generateJwt(user)

			if err == nil {
				var _, err = fmt.Fprintf(w, token)
				if err != nil {
					_, _ = fmt.Fprintf(w, "An error occured: %d", err)
				}
			} else {
				_, _ = fmt.Fprintf(w, "An error occured: %d", err)
			}
		} else {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		}
	} else {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func Register (w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var register model.RegisterForm
	err := decoder.Decode(&register)

	if err == nil {
		if register.Password == register.ConfirmPassword {
			err := persistence.RegisterUser(register); if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			} else {
				w.WriteHeader(http.StatusCreated)
			}
		} else {
			http.Error(w, "Passwords don't match", http.StatusBadRequest)
		}
	} else {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func GetTemperature (w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("readings")
	var numberOfReadings int64
	var asc bool

	if param != "" {
		numberOfReadings, _ = strconv.ParseInt(param, 10, 64)
		asc = true
	} else {
		numberOfReadings = 1
		asc = false
	}

	reading, err := persistence.GetTemperature(numberOfReadings, asc)

	if err == nil {
		w.Header().Set("Content-Type", "application/json")

		encoder := json.NewEncoder(w)
		encoder.Encode(reading)
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetHumidity (w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("readings")
	var numberOfReadings int64
	var asc bool

	if param != "" {
		numberOfReadings, _ = strconv.ParseInt(param, 10, 64)
		asc = true
	} else {
		numberOfReadings = 1
		asc = false
	}

	reading, err := persistence.GetHumidity(numberOfReadings, asc)

	if err == nil {
		w.Header().Set("Content-Type", "application/json")

		encoder := json.NewEncoder(w)
		encoder.Encode(reading)
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetLight (w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("readings")
	var numberOfReadings int64
	var asc bool

	if param != "" {
		numberOfReadings, _ = strconv.ParseInt(param, 10, 64)
		asc = true
	} else {
		numberOfReadings = 1
		asc = false
	}

	reading, err := persistence.GetLight(numberOfReadings, asc)

	if err == nil {
		w.Header().Set("Content-Type", "application/json")

		encoder := json.NewEncoder(w)
		encoder.Encode(reading)
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetSoil (w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("readings")
	var numberOfReadings int64
	var asc bool

	if param != "" {
		numberOfReadings, _ = strconv.ParseInt(param, 10, 64)
		asc = true
	} else {
		numberOfReadings = 1
		asc = false
	}

	reading, err := persistence.GetSoil(numberOfReadings, asc)

	if err == nil {
		w.Header().Set("Content-Type", "application/json")

		encoder := json.NewEncoder(w)
		encoder.Encode(reading)
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CreateTemperature (w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var reading model.ReadingsForm
	err := decoder.Decode(&reading)

	if err == nil {
		err = persistence.CreateTemperatureReading(reading.Reading)

		if err == nil {
			w.WriteHeader(http.StatusCreated)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func CreateHumidity (w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var reading model.ReadingsForm
	err := decoder.Decode(&reading)

	if err == nil {
		err = persistence.CreateHumidityReading(reading.Reading)

		if err == nil {
			w.WriteHeader(http.StatusCreated)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func CreateLight (w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var reading model.ReadingsForm
	err := decoder.Decode(&reading)

	if err == nil {
		err = persistence.CreateLightReading(reading.Reading)

		if err == nil {
			w.WriteHeader(http.StatusCreated)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func CreateSoil (w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var reading model.ReadingsForm
	err := decoder.Decode(&reading)

	if err == nil {
		err = persistence.CreateSoilReading(reading.Reading)

		if err == nil {
			w.WriteHeader(http.StatusCreated)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func TurnOnOffLight (w http.ResponseWriter, r *http.Request) {
	_ = r.URL.Query().Get("status")

	w.WriteHeader(http.StatusNoContent)
}

func WaterPlants (w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

func CreateRoutine (w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var routine model.RoutineForm
	err := decoder.Decode(&routine)

	if err == nil {
		err = persistence.CreateRoutine(routine)

		if err == nil {
			w.WriteHeader(http.StatusCreated)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func GetRoutines (w http.ResponseWriter, r *http.Request) {
	reading, err := persistence.GetRoutines()

	if err == nil {
		w.Header().Set("Content-Type", "application/json")

		encoder := json.NewEncoder(w)
		encoder.Encode(reading)
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}