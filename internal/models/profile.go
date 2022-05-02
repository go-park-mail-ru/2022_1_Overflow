package models

type Avatar struct {
	Name string
	Username string
	File []byte
}

type ProfileInfo struct {
	Id	int32 `json:"id"`
	Firstname string `json:"first_name"`
	Lastname string `json:"last_name"`
	Username string `json:"username"`
}