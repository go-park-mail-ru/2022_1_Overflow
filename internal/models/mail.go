package models

import (
	"time"
)

type Mail struct {
	Id        int32     `json:"id"`
	Sender    string    `json:"sender" validate:"max=45"`
	Addressee string    `json:"addressee" validate:"max=45"`
	Theme     string    `json:"theme"`
	Text      string    `json:"text"`
	Files     string    `json:"files"`
	Date      time.Time `json:"date"`
	Read      bool      `json:"read"`
}

type MailAdditional struct {
	Mail      Mail   `json:"mail"`
	AvatarUrl string `json:"avatar_url"`
}

type MailAddList struct {
	Amount int              `json:"amount"`
	Mails  []MailAdditional `json:"mails"`
}

type MailList struct {
	Amount int    `json:"amount"`
	Mails  []Mail `json:"mails"`
}

type Attach struct {
	Filename    string `json:"filename"`
	Payload     []byte `json:"payload"`
	PayloadSize int64  `json:"payload_size"`
}
