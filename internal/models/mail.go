package models

import "time"

// Структура письма
type Mail struct {
	Client_id int32		`json:"id"`
	Sender    string    `json:"sender"`
	Addressee string    `json:"addressee"`
	Theme     string    `json:"theme"`
	Text      string    `json:"text"`
	Files     string    `json:"files"`
	Date      time.Time `json:"date"`
}
