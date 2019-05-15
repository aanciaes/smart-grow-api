package security

import (
	"github.com/aanciaes/smart-grow-api/config/database"
	"github.com/aanciaes/smart-grow-api/model"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(username, password string) (model.User, bool) {
	db := database.Conn.Connection

	rows, err := db.Query("select * from users where username = ?", username);
	if err != nil {
		return model.User{}, false
	}

	defer rows.Close()

	var user model.User

	rows.Next()
	err = rows.Scan(&user.Id, &user.Name, &user.Hash, &user.IsAdmin)
	if err != nil {
		log.Println(err)
		return model.User{}, false
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Hash), []byte(password))
	return user, err == nil
}