package m

import "time"

// Письмо
// @Description Структура письма
type Mail struct {
	Id		  int32		`json:"id"`
	Client_id int32		`json:"client_id"`
	Sender    string    `json:"sender"`
	Addressee string    `json:"addressee"`
	Theme     string    `json:"theme"`
	Text      string    `json:"text"`
	Files     string    `json:"files"`
	Date      time.Time `json:"date"`
	Read	  bool		`json:"read"`
}

type MailAdditional struct {
	Mail Mail			`json:"mail"`
	AvatarUrl string	`json:"sender_avatar"`
}