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
	AddMail(context.Context, *repository_proto.AddMailRequest) (*utils_proto.DatabaseExtendResponse, error)             // Добавить письмо
	UpdateMail(context.Context, *repository_proto.UpdateMailRequest) (*utils_proto.DatabaseResponse, error) // Обновить данные письма
	DeleteMail(context.Context, *repository_proto.DeleteMailRequest) (*utils_proto.DatabaseResponse, error)             // Удалить письмо
	ReadMail(context.Context, *repository_proto.ReadMailRequest) (*utils_proto.DatabaseResponse, error)                 // Прочитать письмо
	GetMailInfoById(context.Context, *repository_proto.GetMailInfoByIdRequest) (*repository_proto.ResponseMail, error)  // Получить информацию о письме по его id
	CountUnread(ctx context.Context, request *repository_proto.CountUnreadRequest) (*repository_proto.ResponseCountUnread, error)

	AttachRepositoryInterface
	FolderRepositoryInterface
}

type AttachRepositoryInterface interface {
	AddAttachLink(ctx context.Context, request *repository_proto.AddAttachLinkRequest) (*repository_proto.Nothing, error)
	CheckAttachLink(ctx context.Context, request *repository_proto.GetAttachRequest) (*repository_proto.Nothing, error)
	ListAttaches(ctx context.Context, request *repository_proto.GetAttachRequest) (*repository_proto.ResponseAttaches, error)
	CheckAttachPermission(ctx context.Context, request *repository_proto.AttachPermissionRequest) (*repository_proto.ResponseAttachPermission, error)
}

type FolderRepositoryInterface interface {
	GetFolderById(context.Context, *repository_proto.GetFolderByIdRequest) (*repository_proto.ResponseFolder, error)                  // Получить объект папки по id папки
	GetFolderByName(context.Context, *repository_proto.GetFolderByNameRequest) (*repository_proto.ResponseFolder, error)              // Получить объект папки по имени пользователя и названию папки
	GetFoldersByUser(context.Context, *repository_proto.GetFoldersByUserRequest) (*repository_proto.ResponseFolders, error)           // Получить папки, принадлежащие пользователю
	GetFolderMail(context.Context, *repository_proto.GetFolderMailRequest) (*repository_proto.ResponseMails, error)                   // Получить письма, содержащиеся в папке
	DeleteFolder(context.Context, *repository_proto.DeleteFolderRequest) (*utils_proto.DatabaseResponse, error)                       // Удалить папку
	AddFolder(context.Context, *repository_proto.AddFolderRequest) (*utils_proto.DatabaseResponse, error)                             // Добавить папку
	ChangeFolderName(context.Context, *repository_proto.ChangeFolderNameRequest) (*utils_proto.DatabaseResponse, error)               // Изменить название папки
	AddMailToFolderById(context.Context, *repository_proto.AddMailToFolderByIdRequest) (*utils_proto.DatabaseResponse, error)         // Добавить письмо в папку по ID
	AddMailToFolderByObject(context.Context, *repository_proto.AddMailToFolderByObjectRequest) (*utils_proto.DatabaseResponse, error) // Добавить письмо в папку по объекту письма
	DeleteFolderMail(context.Context, *repository_proto.DeleteFolderMailRequest) (*utils_proto.DatabaseResponse, error)               // Удалить письмо из папки
	MoveFolderMail(context.Context, *repository_proto.MoveFolderMailRequest) (*utils_proto.DatabaseResponse, error)                   // Переместить письмо из одной папки в другую
	IsMailMoved(context.Context, *repository_proto.IsMailMovedRequest) (*repository_proto.ResponseIsMoved, error) // Является ли данное письмо перемещенным в папку
}
