package repository

import (
	"OverflowBackend/proto/repository_proto"
	"OverflowBackend/proto/utils_proto"
	"context"
)

type DatabaseRepository interface {
	Create(url string) error

	GetUserInfoByUsername(context.Context, *repository_proto.GetUserInfoByUsernameRequest) (*repository_proto.ResponseUser, error)
	GetUserInfoById(context.Context, *repository_proto.GetUserInfoByIdRequest) (*repository_proto.ResponseUser, error)

	AddUser(context.Context, *repository_proto.AddUserRequest) (*utils_proto.DatabaseResponse, error)
	
	ChangeUserPassword(context.Context, *repository_proto.ChangeForm) (*utils_proto.DatabaseResponse, error)
	ChangeUserFirstName(context.Context, *repository_proto.ChangeForm) (*utils_proto.DatabaseResponse, error)
	ChangeUserLastName(context.Context, *repository_proto.ChangeForm) (*utils_proto.DatabaseResponse, error)

	GetIncomeMails(context.Context, *repository_proto.GetIncomeMailsRequest) (*repository_proto.ResponseMails, error)
	GetOutcomeMails(context.Context, *repository_proto.GetOutcomeMailsRequest) (*repository_proto.ResponseMails, error)
	AddMail(context.Context, *repository_proto.AddMailRequest) (*utils_proto.DatabaseResponse, error)
	DeleteMail(context.Context, *repository_proto.DeleteMailRequest) (*utils_proto.DatabaseResponse, error)
	ReadMail(context.Context, *repository_proto.ReadMailRequest) (*utils_proto.DatabaseResponse, error)
	GetMailInfoById(context.Context, *repository_proto.GetMailInfoByIdRequest) (*repository_proto.ResponseMail, error)
}