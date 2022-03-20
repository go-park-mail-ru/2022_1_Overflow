package models

// Пользователь
// @Description Структура пользователя
type User struct {
	Id        int32
	FirstName string
	LastName  string
	Password  string
	Email     string
}
