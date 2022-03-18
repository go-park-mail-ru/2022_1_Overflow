package repository

import "OverflowBackend/internal/models"

type DatabaseRepository interface {
	Create(url string) error

	GetUserInfoByEmail(email string) (models.User, error)
	GetUserInfoById(id int32) (models.User, error)
	AddUser(user models.User) error

	GetIncomeMails(userId int32) ([]models.Mail, error)
	GetOutcomeMails(userId int32) ([]models.Mail, error)
	AddMail(email models.Mail) error
}
