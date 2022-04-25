package mailbox

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/pkg"
	"OverflowBackend/proto/mailbox_proto"
	"OverflowBackend/proto/profile_proto"
	"OverflowBackend/proto/repository_proto"
	"OverflowBackend/proto/utils_proto"
	"context"
	"encoding/json"

	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type MailBoxService struct {
	config *config.Config
	db repository_proto.DatabaseRepositoryClient
	profile profile_proto.ProfileClient
}

func (s *MailBoxService) Init(config *config.Config, db repository_proto.DatabaseRepositoryClient, profile profile_proto.ProfileClient) {
	s.config = config
	s.db = db
	s.profile = profile
}

func (s *MailBoxService) Income(context context.Context, data *utils_proto.Session) (*mailbox_proto.ResponseMails, error) {
	resp, err := s.db.GetUserInfoByUsername(context, &repository_proto.GetUserInfoByUsernameRequest{Username: data.Username})
	log.Debug("Получение входящих писем, username = ", data.Username)
	if err != nil {
		log.Error(err)
		return &mailbox_proto.ResponseMails{Response: &pkg.DB_ERR, Mails: nil}, err
	}
	if resp.Response.Status != utils_proto.DatabaseStatus_OK {
		return &mailbox_proto.ResponseMails{Response: &pkg.DB_ERR, Mails: nil}, nil
	}
	user := resp.User
	resp2, err := s.db.GetIncomeMails(context, &repository_proto.GetIncomeMailsRequest{UserId: user.Id})
	if err != nil {
		log.Error(err)
		return &mailbox_proto.ResponseMails{Response: &pkg.DB_ERR, Mails: nil}, err
	}
	if resp2.Response.Status != utils_proto.DatabaseStatus_OK {
		return &mailbox_proto.ResponseMails{Response: &pkg.DB_ERR, Mails: nil}, nil
	}
	mails := resp2.Mails
	var mails_add []*utils_proto.MailAdditional
	for _, mail := range mails {
		mail_add := utils_proto.MailAdditional{}
		mail_add.Mail = mail
		resp, err := s.profile.GetAvatar(
			context,
			&profile_proto.GetAvatarRequest{Username: mail.Sender},
		)
		if err != nil {
			return &mailbox_proto.ResponseMails{Response: &pkg.INTERNAL_ERR, Mails: nil}, err
		}
		if !proto.Equal(resp.Response, &pkg.NO_ERR) {
			return &mailbox_proto.ResponseMails{Response: resp.Response, Mails: nil}, nil
		}
		mail_add.AvatarUrl = resp.Url
		mails_add = append(mails_add, &mail_add)
	}
	parsed, err := json.Marshal(mails_add)
	if err != nil {
		log.Error(err)
		return &mailbox_proto.ResponseMails{Response: &pkg.JSON_ERR, Mails: nil}, err
	}
	return &mailbox_proto.ResponseMails{Response: &pkg.NO_ERR, Mails: parsed}, nil
}

func (s *MailBoxService) Outcome(context context.Context, data *utils_proto.Session) (*mailbox_proto.ResponseMails, error) {
	resp, err := s.db.GetUserInfoByUsername(context, &repository_proto.GetUserInfoByUsernameRequest{Username: data.Username})
	log.Debug("Получение исходящих писем, username = ", data.Username)
	if err != nil {
		log.Error(err)
		return &mailbox_proto.ResponseMails{Response: &pkg.DB_ERR, Mails: nil}, err
	}
	if resp.Response.Status != utils_proto.DatabaseStatus_OK {
		return &mailbox_proto.ResponseMails{Response: &pkg.DB_ERR, Mails: nil}, nil
	}
	user := resp.User
	resp2, err := s.db.GetOutcomeMails(context, &repository_proto.GetOutcomeMailsRequest{UserId: user.Id})
	if err != nil {
		log.Error(err)
		return &mailbox_proto.ResponseMails{Response: &pkg.DB_ERR, Mails: nil}, err
	}
	if resp2.Response.Status != utils_proto.DatabaseStatus_OK {
		return &mailbox_proto.ResponseMails{Response: &pkg.DB_ERR, Mails: nil}, nil
	}
	mails := resp2.Mails
	var mails_add []*utils_proto.MailAdditional
	for _, mail := range mails {
		mail_add := utils_proto.MailAdditional{}
		mail_add.Mail = mail
		resp, err := s.profile.GetAvatar(
			context,
			&profile_proto.GetAvatarRequest{Username: mail.Addressee},
		)
		if err != nil {
			return &mailbox_proto.ResponseMails{Response: &pkg.INTERNAL_ERR, Mails: nil}, err
		}
		if !proto.Equal(resp.Response, &pkg.NO_ERR) {
			return &mailbox_proto.ResponseMails{Response: resp.Response, Mails: nil}, nil
		}
		mail_add.AvatarUrl = resp.Url
		mails_add = append(mails_add, &mail_add)
	}
	parsed, err := json.Marshal(mails_add)
	if err != nil {
		log.Error(err)
		return &mailbox_proto.ResponseMails{Response: &pkg.JSON_ERR, Mails: nil}, err
	}
	return &mailbox_proto.ResponseMails{Response: &pkg.NO_ERR, Mails: parsed}, nil
}

func (s *MailBoxService) GetMail(context context.Context, request *mailbox_proto.GetMailRequest) (*mailbox_proto.ResponseMail, error) {
	log.Debug("Получение письма, mail_id = ", request.Id)
	resp, err := s.db.GetMailInfoById(context, &repository_proto.GetMailInfoByIdRequest{
		MailId: request.Id,
	})
	if err != nil {
		log.Error(err)
		return &mailbox_proto.ResponseMail{Response: &pkg.DB_ERR, Mail: nil}, err
	}
	if resp.Response.Status != utils_proto.DatabaseStatus_OK {
		return &mailbox_proto.ResponseMail{Response: &pkg.DB_ERR, Mail: nil}, nil
	}
	mail := resp.Mail
	data := request.Data
	if mail.Addressee != data.Username && mail.Sender != data.Username {
		return &mailbox_proto.ResponseMail{Response: &pkg.UNAUTHORIZED_ERR, Mail: nil}, nil
	}
	parsed, err := json.Marshal(mail)
	if err != nil {
		log.Error(err)
		return &mailbox_proto.ResponseMail{Response: &pkg.JSON_ERR, Mail: nil}, err
	}
	return &mailbox_proto.ResponseMail{Response: &pkg.NO_ERR, Mail: parsed}, nil
}

func (s *MailBoxService) DeleteMail(context context.Context, request *mailbox_proto.DeleteMailRequest) (*utils_proto.JsonResponse, error) {
	log.Debug("Удаление письма, id = ", request.Id)
	resp, err := s.db.GetMailInfoById(context, &repository_proto.GetMailInfoByIdRequest{
		MailId: request.Id,
	})
	if err != nil {
		log.Error(err)
		return &pkg.DB_ERR, err
	}
	if resp.Response.Status != utils_proto.DatabaseStatus_OK {
		return &pkg.DB_ERR, nil
	}
	mail := resp.Mail
	data := request.Data
	if mail.Addressee != data.Username && mail.Sender != data.Username {
		return &pkg.UNAUTHORIZED_ERR, nil
	}
	resp2, err := s.db.DeleteMail(context, &repository_proto.DeleteMailRequest{
		Mail: mail,
		Username: data.Username,
	})
	if err != nil {
		log.Error(err)
		return &pkg.DB_ERR, err
	}
	if resp2.Status != utils_proto.DatabaseStatus_OK {
		return &pkg.DB_ERR, nil
	}
	return &pkg.NO_ERR, nil
}

func (s *MailBoxService) ReadMail(context context.Context, request *mailbox_proto.ReadMailRequest) (*utils_proto.JsonResponse, error) {
	log.Debug("Прочитать письмо, id = ", request.Id)
	resp, err := s.db.GetMailInfoById(context, &repository_proto.GetMailInfoByIdRequest{
		MailId: request.Id,
	})
	if err != nil {
		log.Error(err)
		return &pkg.DB_ERR, err
	}
	if resp.Response.Status != utils_proto.DatabaseStatus_OK {
		return &pkg.DB_ERR, nil
	}
	mail := resp.Mail
	data := request.Data
	if mail.Addressee != data.Username {
		return &pkg.UNAUTHORIZED_ERR, nil
	}
	resp2, err := s.db.ReadMail(context, &repository_proto.ReadMailRequest{
		Mail: mail,
	})
	if err != nil {
		log.Error(err)
		return &pkg.DB_ERR, err
	}
	if resp2.Status != utils_proto.DatabaseStatus_OK {
		return &pkg.DB_ERR, nil
	}
	return &pkg.NO_ERR, nil
}

func (s *MailBoxService) SendMail(context context.Context, request *mailbox_proto.SendMailRequest) (*utils_proto.JsonResponse, error) {
	log.Debug("Отправить письмо, username = ", request.Data.Username)
	resp, err := s.db.GetUserInfoByUsername(context, &repository_proto.GetUserInfoByUsernameRequest{
		Username: request.Data.Username,
	})
	if err != nil {
		log.Error(err)
		return &pkg.DB_ERR, err
	}
	if resp.Response.Status != utils_proto.DatabaseStatus_OK {
		return &pkg.DB_ERR, nil
	}
	data := request.Data
	user := resp.User
	if (proto.Equal(user, &utils_proto.User{})) {
		return &pkg.DB_ERR, nil
	}
	resp2, err := s.db.GetUserInfoByUsername(context, &repository_proto.GetUserInfoByUsernameRequest{
		Username: request.Form.Addressee,
	})
	if err != nil {
		log.Error(err)
		return &pkg.DB_ERR, err
	}
	if resp2.Response.Status != utils_proto.DatabaseStatus_OK {
		return &pkg.DB_ERR, nil
	}
	userAddressee := resp2.User
	if (proto.Equal(userAddressee, &utils_proto.User{})) {
		return &pkg.NO_USER_EXIST, nil
	}
	form := request.Form
	mail := utils_proto.Mail{
		ClientId: user.Id,
		Sender:    data.Username,
		Addressee: form.Addressee,
		Theme:     form.Theme,
		Text:      form.Text,
		Files:     form.Files,
		Date:      timestamppb.Now(),
	}
	resp3, err := s.db.AddMail(context, &repository_proto.AddMailRequest{
		Mail: &mail,
	})
	if err != nil {
		log.Error(err)
		return &pkg.DB_ERR, err
	}
	if resp3.Status != utils_proto.DatabaseStatus_OK {
		return &pkg.DB_ERR, nil
	}
	return &pkg.NO_ERR, nil
}
