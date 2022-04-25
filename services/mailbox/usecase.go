package mailbox

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/proto/mailbox_proto"
	"OverflowBackend/proto/profile_proto"
	"OverflowBackend/proto/repository_proto"
	"OverflowBackend/proto/utils_proto"
)

type MailBoxServiceInterface interface {
	Init(config *config.Config, db repository_proto.DatabaseRepositoryClient, profile profile_proto.ProfileClient)
	Income(data *utils_proto.Session) *mailbox_proto.ResponseMails
	Outcome(data *utils_proto.Session) *mailbox_proto.ResponseMails
	GetMail(request *mailbox_proto.GetMailRequest) *mailbox_proto.ResponseMail
	DeleteMail(request *mailbox_proto.DeleteMailRequest) *utils_proto.JsonResponse
	ReadMail(request *mailbox_proto.ReadMailRequest) *utils_proto.JsonResponse
	SendMail(request *mailbox_proto.SendMailRequest) *utils_proto.JsonResponse
}