package repository

import (
	"OverflowBackend/proto/repository_proto"
	"OverflowBackend/proto/utils_proto"
)

type DatabaseRepository interface {
	Create(url string) error

	GetUserInfoByUsername(*repository_proto.GetUserInfoByUsernameRequest) *repository_proto.ResponseUser
	GetUserInfoById(*repository_proto.GetUserInfoByIdRequest) *repository_proto.ResponseUser

	AddUser(*repository_proto.AddUserRequest) *utils_proto.DatabaseResponse
	
	ChangeUserPassword(*repository_proto.ChangeForm) *utils_proto.DatabaseResponse
	ChangeUserFirstName(*repository_proto.ChangeForm) *utils_proto.DatabaseResponse
	ChangeUserLastName(*repository_proto.ChangeForm) *utils_proto.DatabaseResponse

	GetIncomeMails(*repository_proto.GetIncomeMailsRequest) *repository_proto.ResponseMails
	GetOutcomeMails(*repository_proto.GetOutcomeMailsRequest) *repository_proto.ResponseMails
	AddMail(*repository_proto.AddMailRequest) *utils_proto.DatabaseResponse
	DeleteMail(*repository_proto.DeleteMailRequest) *utils_proto.DatabaseResponse
	ReadMail(*repository_proto.ReadMailRequest) *utils_proto.DatabaseResponse
	GetMailInfoById(*repository_proto.GetMailInfoByIdRequest) repository_proto.ResponseMail
}