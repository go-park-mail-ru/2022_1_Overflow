package repository

import "OverflowBackend/internal/models"

type DatabaseRepository interface {
	Create(url string) error

	GetUserInfoByEmail(email string) (models.User, error)
	GetUserInfoById(userId int32) (models.User, error)
	AddUser(user models.User) error
	ChangeUserPassword(user models.User, newPassword string) error

	GetIncomeMails(userId int32) ([]models.Mail, error)
	GetOutcomeMails(userId int32) ([]models.Mail, error)
	AddMail(email models.Mail) error
	DeleteMail(email models.Mail, userEmail string) error
	ReadMail(email models.Mail) error
	GetMailInfoById(mailId int) (models.Mail, error)
}
