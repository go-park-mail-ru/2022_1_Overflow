package repository

import "OverflowBackend/internal/models"

type DatabaseRepository interface {
	Create(url string) error

	GetUserInfoByUsername(username string) (models.User, error)
	GetUserInfoById(userId int32) (models.User, error)
	AddUser(user models.User) error

	ChangeUserPassword(user models.User, newPassword string) error
	ChangeUserFirstName(user models.User, newFirstName string) error
	ChangeUserLastName(user models.User, newLastName string) error

	GetIncomeMails(userId int32) ([]models.Mail, error)
	GetOutcomeMails(userId int32) ([]models.Mail, error)
	AddMail(mail models.Mail) error
	DeleteMail(mail models.Mail, username string) error
	ReadMail(mail models.Mail) error
	GetMailInfoById(mailId int32) (models.Mail, error)
}
