package models

type Avatar struct {
	Name string `json:"name"`
	Username string `json:"username"`
	File []byte `json:"file"`
}

type ProfileInfo struct {
	Id	int32 `json:"id"`
	Firstname string `json:"first_name" validate:"max=45"`
	Lastname string `json:"last_name" validate:"max=45"`
	Username string `json:"username" validate:"max=45"`
}