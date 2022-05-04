package mailbox

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/internal/models"
	"OverflowBackend/pkg"
	"OverflowBackend/proto/mailbox_proto"
	"OverflowBackend/proto/profile_proto"
	"OverflowBackend/proto/repository_proto"
	"OverflowBackend/proto/utils_proto"
	"context"
	"encoding/json"
	"time"

	log "github.com/sirupsen/logrus"
)

type MailBoxService struct {
	config  *config.Config
	db      repository_proto.DatabaseRepositoryClient
	profile profile_proto.ProfileClient
}

func (s *MailBoxService) Init(config *config.Config, db repository_proto.DatabaseRepositoryClient, profile profile_proto.ProfileClient) {
	s.config = config
	s.db = db
	s.profile = profile
}

func (s *MailBoxService) Income(context context.Context, request *mailbox_proto.IncomeRequest) (*mailbox_proto.ResponseMails, error) {
	log.Debug("Получение входящих писем, username = ", request.Data.Username, ", limit = ", request.Limit, ", offset = ", request.Offset)
	resp, err := s.db.GetUserInfoByUsername(context, &repository_proto.GetUserInfoByUsernameRequest{Username: request.Data.Username})
	if err != nil {
		log.Error(err)
		return &mailbox_proto.ResponseMails{
			Response: &utils_proto.JsonResponse{
				Response: pkg.DB_ERR.Bytes(),
			}, Mails: nil,
		}, err
	}
	if resp.Response.Status != utils_proto.DatabaseStatus_OK {
		return &mailbox_proto.ResponseMails{
			Response: &utils_proto.JsonResponse{
				Response: pkg.DB_ERR.Bytes(),
			}, Mails: nil,
		}, nil
	}
	var user models.User
	err = json.Unmarshal(resp.User, &user)
	if err != nil {
		return &mailbox_proto.ResponseMails{
			Response: &utils_proto.JsonResponse{
				Response: pkg.JSON_ERR.Bytes(),
			}, Mails: nil,
		}, err
	}
	resp2, err := s.db.GetIncomeMails(context, &repository_proto.GetIncomeMailsRequest{UserId: user.Id, Limit: request.Limit, Offset: request.Offset})
	if err != nil {
		log.Error(err)
		return &mailbox_proto.ResponseMails{
			Response: &utils_proto.JsonResponse{
				Response: pkg.DB_ERR.Bytes(),
			}, Mails: nil,
		}, err
	}
	if resp2.Response.Status != utils_proto.DatabaseStatus_OK {
		return &mailbox_proto.ResponseMails{
			Response: &utils_proto.JsonResponse{
				Response: pkg.DB_ERR.Bytes(),
			}, Mails: nil,
		}, nil
	}
	var mails []models.Mail
	err = json.Unmarshal(resp2.Mails, &mails)
	if err != nil {
		return &mailbox_proto.ResponseMails{
			Response: &utils_proto.JsonResponse{
				Response: pkg.JSON_ERR.Bytes(),
			}, Mails: nil,
		}, err
	}
	var mails_add []models.MailAdditional
	for _, mail := range mails {
		mail_add := models.MailAdditional{}
		mail_add.Mail = mail
		resp, err := s.profile.GetAvatar(
			context,
			&profile_proto.GetAvatarRequest{Username: mail.Sender},
		)
		if err != nil {
			return &mailbox_proto.ResponseMails{
				Response: &utils_proto.JsonResponse{
					Response: pkg.INTERNAL_ERR.Bytes(),
				}, Mails: nil,
			}, err
		}
		var response pkg.JsonResponse
		err = json.Unmarshal(resp.Response.Response, &response)
		if err != nil {
			return &mailbox_proto.ResponseMails{
				Response: &utils_proto.JsonResponse{
					Response: pkg.JSON_ERR.Bytes(),
				}, Mails: nil,
			}, err
		}
		if response != pkg.NO_ERR {
			return &mailbox_proto.ResponseMails{Response: resp.Response, Mails: nil}, nil
		}
		mail_add.AvatarUrl = resp.Url
		mails_add = append(mails_add, mail_add)
	}
	parsed, err := json.Marshal(mails_add)
	if err != nil {
		log.Error(err)
		return &mailbox_proto.ResponseMails{
			Response: &utils_proto.JsonResponse{
				Response: pkg.JSON_ERR.Bytes(),
			}, Mails: nil,
		}, err
	}
	return &mailbox_proto.ResponseMails{
		Response: &utils_proto.JsonResponse{
			Response: pkg.NO_ERR.Bytes(),
		}, Mails: parsed,
	}, nil
}

func (s *MailBoxService) Outcome(context context.Context, request *mailbox_proto.OutcomeRequest) (*mailbox_proto.ResponseMails, error) {
	log.Debug("Получение входящих писем, username = ", request.Data.Username, ", limit = ", request.Limit, ", offset = ", request.Offset)
	resp, err := s.db.GetUserInfoByUsername(context, &repository_proto.GetUserInfoByUsernameRequest{Username: request.Data.Username})
	if err != nil {
		log.Error(err)
		return &mailbox_proto.ResponseMails{
			Response: &utils_proto.JsonResponse{
				Response: pkg.DB_ERR.Bytes(),
			}, Mails: nil,
		}, err
	}
	if resp.Response.Status != utils_proto.DatabaseStatus_OK {
		return &mailbox_proto.ResponseMails{
			Response: &utils_proto.JsonResponse{
				Response: pkg.DB_ERR.Bytes(),
			}, Mails: nil,
		}, nil
	}
	var user models.User
	err = json.Unmarshal(resp.User, &user)
	if err != nil {
		return &mailbox_proto.ResponseMails{
			Response: &utils_proto.JsonResponse{
				Response: pkg.JSON_ERR.Bytes(),
			}, Mails: nil,
		}, err
	}
	resp2, err := s.db.GetOutcomeMails(context, &repository_proto.GetOutcomeMailsRequest{UserId: user.Id, Limit: request.Limit, Offset: request.Offset})
	if err != nil {
		log.Error(err)
		return &mailbox_proto.ResponseMails{
			Response: &utils_proto.JsonResponse{
				Response: pkg.DB_ERR.Bytes(),
			}, Mails: nil,
		}, err
	}
	if resp2.Response.Status != utils_proto.DatabaseStatus_OK {
		return &mailbox_proto.ResponseMails{
			Response: &utils_proto.JsonResponse{
				Response: pkg.DB_ERR.Bytes(),
			}, Mails: nil,
		}, nil
	}
	var mails []models.Mail
	err = json.Unmarshal(resp2.Mails, &mails)
	if err != nil {
		return &mailbox_proto.ResponseMails{
			Response: &utils_proto.JsonResponse{
				Response: pkg.JSON_ERR.Bytes(),
			}, Mails: nil,
		}, err
	}
	var mails_add []models.MailAdditional
	for _, mail := range mails {
		mail_add := models.MailAdditional{}
		mail_add.Mail = mail
		resp, err := s.profile.GetAvatar(
			context,
			&profile_proto.GetAvatarRequest{Username: mail.Addressee},
		)
		if err != nil {
			return &mailbox_proto.ResponseMails{
				Response: &utils_proto.JsonResponse{
					Response: pkg.INTERNAL_ERR.Bytes(),
				}, Mails: nil,
			}, err
		}
		var response pkg.JsonResponse
		err = json.Unmarshal(resp.Response.Response, &response)
		if err != nil {
			return &mailbox_proto.ResponseMails{
				Response: &utils_proto.JsonResponse{
					Response: pkg.JSON_ERR.Bytes(),
				}, Mails: nil,
			}, err
		}
		if response != pkg.NO_ERR {
			return &mailbox_proto.ResponseMails{Response: resp.Response, Mails: nil}, nil
		}
		mail_add.AvatarUrl = resp.Url
		mails_add = append(mails_add, mail_add)
	}
	parsed, err := json.Marshal(mails_add)
	if err != nil {
		log.Error(err)
		return &mailbox_proto.ResponseMails{
			Response: &utils_proto.JsonResponse{
				Response: pkg.JSON_ERR.Bytes(),
			}, Mails: nil,
		}, err
	}
	return &mailbox_proto.ResponseMails{
		Response: &utils_proto.JsonResponse{
			Response: pkg.NO_ERR.Bytes(),
		}, Mails: parsed,
	}, nil
}

func (s *MailBoxService) GetMail(context context.Context, request *mailbox_proto.GetMailRequest) (*mailbox_proto.ResponseMail, error) {
	log.Debug("Получение письма, mail_id = ", request.Id)
	resp, err := s.db.GetMailInfoById(context, &repository_proto.GetMailInfoByIdRequest{
		MailId: request.Id,
	})
	if err != nil {
		log.Error(err)
		return &mailbox_proto.ResponseMail{
			Response: &utils_proto.JsonResponse{
				Response: pkg.DB_ERR.Bytes(),
			}, Mail: nil,
		}, err
	}
	if resp.Response.Status != utils_proto.DatabaseStatus_OK {
		return &mailbox_proto.ResponseMail{
			Response: &utils_proto.JsonResponse{
				Response: pkg.DB_ERR.Bytes(),
			}, Mail: nil,
		}, nil
	}
	var mail models.Mail
	err = json.Unmarshal(resp.Mail, &mail)
	if err != nil {
		return &mailbox_proto.ResponseMail{
			Response: &utils_proto.JsonResponse{
				Response: pkg.JSON_ERR.Bytes(),
			}, Mail: nil,
		}, err
	}
	data := request.Data
	if mail.Addressee != data.Username && mail.Sender != data.Username {
		return &mailbox_proto.ResponseMail{
			Response: &utils_proto.JsonResponse{
				Response: pkg.UNAUTHORIZED_ERR.Bytes(),
			}, Mail: nil,
		}, nil
	}
	parsed, err := json.Marshal(mail)
	if err != nil {
		log.Error(err)
		return &mailbox_proto.ResponseMail{
			Response: &utils_proto.JsonResponse{
				Response: pkg.JSON_ERR.Bytes(),
			}, Mail: nil,
		}, err
	}
	return &mailbox_proto.ResponseMail{
		Response: &utils_proto.JsonResponse{
			Response: pkg.NO_ERR.Bytes(),
		}, Mail: parsed,
	}, nil
}

func (s *MailBoxService) DeleteMail(context context.Context, request *mailbox_proto.DeleteMailRequest) (*utils_proto.JsonResponse, error) {
	log.Debug("Удаление письма, id = ", request.Id)
	resp, err := s.db.GetMailInfoById(context, &repository_proto.GetMailInfoByIdRequest{
		MailId: request.Id,
	})
	if err != nil {
		log.Error(err)
		return &utils_proto.JsonResponse{
			Response: pkg.DB_ERR.Bytes(),
		}, err
	}
	if resp.Response.Status != utils_proto.DatabaseStatus_OK {
		return &utils_proto.JsonResponse{
			Response: pkg.DB_ERR.Bytes(),
		}, nil
	}
	var mail models.Mail
	err = json.Unmarshal(resp.Mail, &mail)
	if err != nil {
		return &utils_proto.JsonResponse{
			Response: pkg.JSON_ERR.Bytes(),
		}, err
	}
	data := request.Data
	if mail.Addressee != data.Username && mail.Sender != data.Username {
		return &utils_proto.JsonResponse{
			Response: pkg.UNAUTHORIZED_ERR.Bytes(),
		}, nil
	}
	respUser, err := s.db.GetUserInfoByUsername(context, &repository_proto.GetUserInfoByUsernameRequest{
		Username: data.Username,
	})
	if err != nil {
		return &utils_proto.JsonResponse{
			Response: pkg.DB_ERR.Bytes(),
		}, err
	}
	if respUser.Response.Status != utils_proto.DatabaseStatus_OK {
		return &utils_proto.JsonResponse{
			Response: pkg.DB_ERR.Bytes(),
		}, nil
	}
	var user models.User
	err = json.Unmarshal(respUser.User, &user)
	if err != nil {
		return &utils_proto.JsonResponse{
			Response: pkg.JSON_ERR.Bytes(),
		}, err
	}
	resp2, err := s.db.DeleteMail(context, &repository_proto.DeleteMailRequest{
		Mail:     resp.Mail,
		UserId:	  user.Id,
	})
	if err != nil {
		log.Error(err)
		return &utils_proto.JsonResponse{
			Response: pkg.DB_ERR.Bytes(),
		}, err
	}
	if resp2.Status != utils_proto.DatabaseStatus_OK {
		return &utils_proto.JsonResponse{
			Response: pkg.DB_ERR.Bytes(),
		}, nil
	}
	return &utils_proto.JsonResponse{
		Response: pkg.NO_ERR.Bytes(),
	}, nil
}

func (s *MailBoxService) ReadMail(context context.Context, request *mailbox_proto.ReadMailRequest) (*utils_proto.JsonResponse, error) {
	log.Debug("Прочитать письмо, id = ", request.Id)
	resp, err := s.db.GetMailInfoById(context, &repository_proto.GetMailInfoByIdRequest{
		MailId: request.Id,
	})
	if err != nil {
		log.Error(err)
		return &utils_proto.JsonResponse{
			Response: pkg.DB_ERR.Bytes(),
		}, err
	}
	if resp.Response.Status != utils_proto.DatabaseStatus_OK {
		return &utils_proto.JsonResponse{
			Response: pkg.DB_ERR.Bytes(),
		}, nil
	}
	var mail models.Mail
	err = json.Unmarshal(resp.Mail, &mail)
	if err != nil {
		return &utils_proto.JsonResponse{
			Response: pkg.JSON_ERR.Bytes(),
		}, err
	}
	data := request.Data
	if mail.Addressee != data.Username {
		return &utils_proto.JsonResponse{
			Response: pkg.UNAUTHORIZED_ERR.Bytes(),
		}, nil
	}
	resp2, err := s.db.ReadMail(context, &repository_proto.ReadMailRequest{
		Mail: resp.Mail,
		Read: request.Read,
	})
	if err != nil {
		log.Error(err)
		return &utils_proto.JsonResponse{
			Response: pkg.DB_ERR.Bytes(),
		}, err
	}
	if resp2.Status != utils_proto.DatabaseStatus_OK {
		return &utils_proto.JsonResponse{
			Response: pkg.DB_ERR.Bytes(),
		}, nil
	}
	return &utils_proto.JsonResponse{
		Response: pkg.NO_ERR.Bytes(),
	}, nil
}

func (s *MailBoxService) SendMail(context context.Context, request *mailbox_proto.SendMailRequest) (*utils_proto.JsonResponse, error) {
	log.Debug("Отправить письмо, username = ", request.Data.Username)
	resp, err := s.db.GetUserInfoByUsername(context, &repository_proto.GetUserInfoByUsernameRequest{
		Username: request.Data.Username,
	})
	if err != nil {
		log.Error(err)
		return &utils_proto.JsonResponse{
			Response: pkg.DB_ERR.Bytes(),
		}, err
	}
	if resp.Response.Status != utils_proto.DatabaseStatus_OK {
		return &utils_proto.JsonResponse{
			Response: pkg.DB_ERR.Bytes(),
		}, nil
	}
	data := request.Data
	var user models.User
	err = json.Unmarshal(resp.User, &user)
	if err != nil {
		return &utils_proto.JsonResponse{
			Response: pkg.JSON_ERR.Bytes(),
		}, err
	}
	if (user == models.User{}) {
		return &utils_proto.JsonResponse{
			Response: pkg.NO_USER_EXIST.Bytes(),
		}, nil
	}
	var form models.MailForm
	err = json.Unmarshal(request.Form, &form)
	if err != nil {
		return &utils_proto.JsonResponse{
			Response: pkg.JSON_ERR.Bytes(),
		}, err
	}
	resp2, err := s.db.GetUserInfoByUsername(context, &repository_proto.GetUserInfoByUsernameRequest{
		Username: form.Addressee,
	})
	if err != nil {
		log.Error(err)
		return &utils_proto.JsonResponse{
			Response: pkg.DB_ERR.Bytes(),
		}, err
	}
	if resp2.Response.Status != utils_proto.DatabaseStatus_OK {
		return &utils_proto.JsonResponse{
			Response: pkg.DB_ERR.Bytes(),
		}, nil
	}
	var userAddressee models.User
	err = json.Unmarshal(resp2.User, &userAddressee)
	if err != nil {
		return &utils_proto.JsonResponse{
			Response: pkg.JSON_ERR.Bytes(),
		}, err
	}
	if (userAddressee == models.User{}) {
		return &utils_proto.JsonResponse{
			Response: pkg.NO_USER_EXIST.Bytes(),
		}, nil
	}
	mail := models.Mail{
		ClientId:  user.Id,
		Sender:    data.Username,
		Addressee: form.Addressee,
		Theme:     form.Theme,
		Text:      form.Text,
		Files:     form.Files,
		Date:      time.Now(),
	}
	mailBytes, _ := json.Marshal(mail)
	resp3, err := s.db.AddMail(context, &repository_proto.AddMailRequest{
		Mail: mailBytes,
	})
	if err != nil {
		log.Error(err)
		return &utils_proto.JsonResponse{
			Response: pkg.DB_ERR.Bytes(),
		}, err
	}
	if resp3.Status != utils_proto.DatabaseStatus_OK {
		return &utils_proto.JsonResponse{
			Response: pkg.DB_ERR.Bytes(),
		}, nil
	}
	return &utils_proto.JsonResponse{
		Response: pkg.NO_ERR.Bytes(),
	}, nil
}
