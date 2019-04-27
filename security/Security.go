package security

import (
	"github.com/aanciaes/smart-grow-api/config/database"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(username, password string) bool {
	db := database.Conn.Connection

	rows, err := db.Query("select * from users where username = ?", username);
	if err != nil {
		return false
	}

	defer rows.Close()

	var (
		id       int
		name string
		hash     string
		isAdmin  int
	)

	rows.Next()
	err = rows.Scan(&id, &name, &hash, &isAdmin)
	if err != nil {
		log.Println(err)
		return false
	}
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
