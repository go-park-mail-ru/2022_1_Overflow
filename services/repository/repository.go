package repository

import (
	"OverflowBackend/proto/repository_proto"
	"OverflowBackend/proto/utils_proto"
	"context"
)

type DatabaseRepository interface {
	Create(url string) error // Инициализировать подключение к БД по ссылке url

	GetUserInfoByUsername(context.Context, *repository_proto.GetUserInfoByUsernameRequest) (*repository_proto.ResponseUser, error) // Получить информацию о пользователе по его логину (имени пользователя)
	GetUserInfoById(context.Context, *repository_proto.GetUserInfoByIdRequest) (*repository_proto.ResponseUser, error)             // Получить информацию о пользователе по его id

	AddUser(context.Context, *repository_proto.AddUserRequest) (*utils_proto.DatabaseResponse, error) // Добавить пользователя

	ChangeUserPassword(context.Context, *repository_proto.ChangeForm) (*utils_proto.DatabaseResponse, error)  // Изменить пароль пользователя
	ChangeUserFirstName(context.Context, *repository_proto.ChangeForm) (*utils_proto.DatabaseResponse, error) // Изменить имя пользователя
	ChangeUserLastName(context.Context, *repository_proto.ChangeForm) (*utils_proto.DatabaseResponse, error)  // Изменить фамилию пользователя

	GetIncomeMails(context.Context, *repository_proto.GetIncomeMailsRequest) (*repository_proto.ResponseMails, error)   // Получить входящие письма
	GetOutcomeMails(context.Context, *repository_proto.GetOutcomeMailsRequest) (*repository_proto.ResponseMails, error) // Получить исходящие письма
	AddMail(context.Context, *repository_proto.AddMailRequest) (*utils_proto.DatabaseResponse, error)                   // Добавить письмо
	DeleteMail(context.Context, *repository_proto.DeleteMailRequest) (*utils_proto.DatabaseResponse, error)             // Удалить письмо
	ReadMail(context.Context, *repository_proto.ReadMailRequest) (*utils_proto.DatabaseResponse, error)                 // Прочитать письмо
	GetMailInfoById(context.Context, *repository_proto.GetMailInfoByIdRequest) (*repository_proto.ResponseMail, error)  // Получить информацию о письме по его id
}

type FolderRepositoryInteface interface {
	GetFolderById(context.Context, *repository_proto.GetFolderByIdRequest)       // Получить объект папки по id папки
	GetFolderByName(context.Context, *repository_proto.GetFolderByNameRequest)   // Получить объект папки по имени пользователя и названию папки
	GetFolderMail(context.Context, *repository_proto.GetFolderMailRequest)       // Получить письма, содержащиеся в папке
	DeleteFolder(context.Context, *repository_proto.DeleteFolderRequest)         // Удалить папку
	AddFolder(context.Context, *repository_proto.AddFolderRequest)               // Добавить папку
	ChangeFolderName(context.Context, *repository_proto.ChangeFolderNameRequest) // Изменить название папки
	AddMailToFolder(context.Context, *repository_proto.AddMailToFolderRequest)   // Добавить письмо в папку
}
