package models

import "time"

type Folder struct {
	Id        int32  `json:"id"`
	Name      string `json:"name"`
	UserId    int32  `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}