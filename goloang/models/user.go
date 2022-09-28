package models

type User struct {
	Id       string
	Username string
	Password string
	Profile  Profile
}

type Profile struct {
}
