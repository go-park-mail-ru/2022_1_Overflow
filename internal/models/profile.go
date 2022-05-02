package models

type Avatar struct {
	Name string
	Username string
	File []byte
}

type ProfileInfo struct {
	Id	int32 `json:"id"`
	Firstname string `json:"first_name" validate:"max=45"`
	Lastname string `json:"last_name" validate:"max=45"`
	Username string `json:"username" validate:"max=45"`
}