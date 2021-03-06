package mailbox

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/proto/mailbox_proto"
	"OverflowBackend/proto/profile_proto"
	"OverflowBackend/proto/repository_proto"
	"OverflowBackend/proto/utils_proto"
	"context"
)

type MailBoxServiceInterface interface {
	Init(config *config.Config, db repository_proto.DatabaseRepositoryClient, profile profile_proto.ProfileClient)
	Income(context context.Context, data *mailbox_proto.IncomeRequest) (*mailbox_proto.ResponseMails, error)
	Outcome(context context.Context, data *mailbox_proto.OutcomeRequest) (*mailbox_proto.ResponseMails, error)
	GetMail(context context.Context, request *mailbox_proto.GetMailRequest) (*mailbox_proto.ResponseMail, error)
	DeleteMail(context context.Context, request *mailbox_proto.DeleteMailRequest) (*utils_proto.JsonResponse, error)
	ReadMail(context context.Context, request *mailbox_proto.ReadMailRequest) (*utils_proto.JsonResponse, error)
	SendMail(context context.Context, request *mailbox_proto.SendMailRequest) (*utils_proto.JsonExtendResponse, error)
	CountUnread(context context.Context, request *mailbox_proto.CountUnreadRequest) (*mailbox_proto.ResponseCountUnread, error)
}
