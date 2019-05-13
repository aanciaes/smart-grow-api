package model

type LoginForm struct {
	Username string `json:username`
	Password string `json:password`
}

type User struct {
	Id       int
	Name string
	Hash     string
	IsAdmin  int
}