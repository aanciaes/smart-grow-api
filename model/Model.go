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

type ReadingsForm struct {
	Reading float32
}

type Readings struct {
	Id int
	Date string
	Reading float32
}

type RoutineForm struct {
	Datetime string `json:datetime`
	Output string `json:output`
}

type Routine struct {
	Id int
	Datetime string
	Output string
}

type DeleteRoutineForm struct {
	Id int
}