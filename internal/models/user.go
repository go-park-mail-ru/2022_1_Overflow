package models

type User struct {
	Id int32 `json:"id"`
	Firstname string `json:"first_name"`
	Lastname string `json:"last_name"`
	Username string `json:"username"`
	Password string `json:"password"`
}