package models

import "time"

type Mail struct {
	Id        int32     `json:"id"`
	ClientId  int32     `json:"client_id"`
	Sender    string    `json:"sender" validate:"max=45"`
	Addressee string    `json:"addressee" validate:"max=45"`
	Theme     string    `json:"theme"`
	Text      string    `json:"text"`
	Files     string    `json:"files"`
	Date      time.Time `json:"date"`
	Read      bool      `json:"read"`
}

type MailAdditional struct {
	Mail Mail `json:"mail"`
	AvatarUrl string `json:"avatar_url"`
}