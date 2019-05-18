package model

type LoginForm struct {
	Username string `json:username`
	Password string `json:password`
}

type RegisterForm struct {
	Username string `json:username`
	Password string `json:password`
	ConfirmPassword string `json:confirmPassword`
	IsAdmin bool `json:isAdmin`
}

type User struct {
	Id       int
	Name string
	Hash     string
	IsAdmin  bool
}

type TemperatureReading struct {
	Id int
	Date string
	Reading float32
}