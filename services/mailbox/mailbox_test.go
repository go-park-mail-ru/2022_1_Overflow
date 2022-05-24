package mailbox_test

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/internal/models"
	"OverflowBackend/pkg"
	"OverflowBackend/proto/mailbox_proto"
	"OverflowBackend/proto/profile_proto"
	"OverflowBackend/proto/repository_proto"
	"OverflowBackend/proto/utils_proto"
	"OverflowBackend/services/mailbox"
	"context"
	"github.com/mailru/easyjson"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/golang/mock/gomock"
	log "github.com/sirupsen/logrus"
)

func InitTestUseCase(ctrl *gomock.Controller) (*repository_proto.MockDatabaseRepositoryClient, *profile_proto.MockProfileClient, *mailbox.MailBoxService) {
	current := time.Now()
	monkey.Patch(time.Now, func() time.Time { return current })
	log.SetLevel(log.FatalLevel)
	db := repository_proto.NewMockDatabaseRepositoryClient(ctrl)
	profile := profile_proto.NewMockProfileClient(ctrl)
	uc := mailbox.MailBoxService{}
	uc.Init(config.TestConfig(), db, profile)
	return db, profile, &uc
}

func TestIncome(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB, _, uc := InitTestUseCase(mockCtrl)

	mails := models.MailList{
		Amount: 0,
		Mails:  []models.Mail{},
	}
	mailsBytes, _ := easyjson.Marshal(mails)

	user := models.User{
		Id:        0,
		Firstname: "test",
		Lastname:  "test",
		Username:  "test",
		Password:  "test",
	}
	userBytes, _ := easyjson.Marshal(user)

	session := utils_proto.Session{
		Username:      user.Username,
		Authenticated: true,
	}

	var limit int32 = 10
	var offset int32 = 0

	mockDB.EXPECT().GetUserInfoByUsername(context.Background(), &repository_proto.GetUserInfoByUsernameRequest{
		Username: user.Username,
	}).Return(&repository_proto.ResponseUser{
		User: userBytes,
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
	}, nil)

	mockDB.EXPECT().GetIncomeMails(context.Background(), &repository_proto.GetIncomeMailsRequest{
		UserId: user.Id,
		Limit:  limit,
		Offset: offset,
	}).Return(&repository_proto.ResponseMails{
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
		Mails: mailsBytes,
	}, nil)

	resp, err := uc.Income(context.Background(), &mailbox_proto.IncomeRequest{
		Data:   &session,
		Limit:  limit,
		Offset: offset,
	})
	var response pkg.JsonResponse
	json_err := easyjson.Unmarshal(resp.Response.Response, &response)
	if err != nil || json_err != nil || response != pkg.NO_ERR {
		t.Errorf("Неверный ответ от UseCase.")
		return
	}
}

func TestOutcome(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB, _, uc := InitTestUseCase(mockCtrl)

	mails := models.MailList{
		Amount: 0,
		Mails:  []models.Mail{},
	}
	mailsBytes, _ := easyjson.Marshal(mails)

	user := models.User{
		Id:        0,
		Firstname: "test",
		Lastname:  "test",
		Username:  "test",
		Password:  "test",
	}
	userBytes, _ := easyjson.Marshal(user)

	session := utils_proto.Session{
		Username:      user.Username,
		Authenticated: true,
	}

	var limit int32 = 10
	var offset int32 = 0

	mockDB.EXPECT().GetUserInfoByUsername(context.Background(), &repository_proto.GetUserInfoByUsernameRequest{
		Username: user.Username,
	}).Return(&repository_proto.ResponseUser{
		User: userBytes,
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
	}, nil)

	mockDB.EXPECT().GetOutcomeMails(context.Background(), &repository_proto.GetOutcomeMailsRequest{
		UserId: user.Id,
		Limit:  limit,
		Offset: offset,
	}).Return(&repository_proto.ResponseMails{
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
		Mails: mailsBytes,
	}, nil)

	resp, err := uc.Outcome(context.Background(), &mailbox_proto.OutcomeRequest{
		Data:   &session,
		Limit:  limit,
		Offset: offset,
	})
	var response pkg.JsonResponse
	json_err := easyjson.Unmarshal(resp.Response.Response, &response)
	if err != nil || json_err != nil || response != pkg.NO_ERR {
		t.Errorf("Неверный ответ от UseCase.")
		return
	}
}

func TestGetMail(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB, _, uc := InitTestUseCase(mockCtrl)

	user := models.User{
		Id:        0,
		Firstname: "test",
		Lastname:  "test",
		Username:  "test",
		Password:  "test",
	}
	//userBytes, _ := easyjson.Marshal(user)

	session := utils_proto.Session{
		Username:      user.Username,
		Authenticated: true,
	}

	mail := models.Mail{
		Id:        0,
		Sender:    session.Username,
		Addressee: "test2",
		Theme:     "test",
		Text:      "test",
		Files:     "files",
		Date:      time.Now(),
		Read:      false,
	}
	mailBytes, _ := easyjson.Marshal(mail)

	var mailId int32 = mail.Id

	mockDB.EXPECT().GetMailInfoById(context.Background(), &repository_proto.GetMailInfoByIdRequest{
		MailId: mailId,
	}).Return(&repository_proto.ResponseMail{
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
		Mail: mailBytes,
	}, nil)

	resp, err := uc.GetMail(context.Background(), &mailbox_proto.GetMailRequest{
		Data: &session,
		Id:   mailId,
	})
	var response pkg.JsonResponse
	var respMail models.Mail
	json_err := easyjson.Unmarshal(resp.Response.Response, &response)
	json_mail_err := easyjson.Unmarshal(resp.Mail, &respMail)
	if err != nil || json_err != nil || json_mail_err != nil || response != pkg.NO_ERR {
		t.Errorf("Неверный ответ от UseCase.")
		return
	}
}

func TestDeleteMail(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB, _, uc := InitTestUseCase(mockCtrl)

	user := models.User{
		Id:        0,
		Firstname: "test",
		Lastname:  "test",
		Username:  "test",
		Password:  "test",
	}
	userBytes, _ := easyjson.Marshal(user)

	session := utils_proto.Session{
		Username:      user.Username,
		Authenticated: true,
	}

	mail := models.Mail{
		Id:        0,
		Sender:    session.Username,
		Addressee: "test2",
		Theme:     "test",
		Text:      "test",
		Files:     "files",
		Date:      time.Now(),
		Read:      false,
	}
	mailBytes, _ := easyjson.Marshal(mail)

	var mailId int32 = mail.Id

	mockDB.EXPECT().GetMailInfoById(context.Background(), &repository_proto.GetMailInfoByIdRequest{
		MailId: mailId,
	}).Return(&repository_proto.ResponseMail{
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
		Mail: mailBytes,
	}, nil)

	mockDB.EXPECT().GetUserInfoByUsername(context.Background(), &repository_proto.GetUserInfoByUsernameRequest{
		Username: user.Username,
	}).Return(&repository_proto.ResponseUser{
		User: userBytes,
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
	}, nil)

	mockDB.EXPECT().DeleteMail(context.Background(), &repository_proto.DeleteMailRequest{
		Mail:   mailBytes,
		UserId: user.Id,
	}).Return(&utils_proto.DatabaseResponse{
		Status: utils_proto.DatabaseStatus_OK,
	}, nil)

	resp, err := uc.DeleteMail(context.Background(), &mailbox_proto.DeleteMailRequest{
		Data: &session,
		Id:   mailId,
	})
	var response pkg.JsonResponse
	json_err := easyjson.Unmarshal(resp.Response, &response)
	if err != nil || json_err != nil || response != pkg.NO_ERR {
		t.Errorf("Неверный ответ от UseCase.")
		return
	}
}

func TestReadMail(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB, _, uc := InitTestUseCase(mockCtrl)

	user := models.User{
		Id:        0,
		Firstname: "test",
		Lastname:  "test",
		Username:  "test",
		Password:  "test",
	}
	//userBytes, _ := easyjson.Marshal(user)

	session := utils_proto.Session{
		Username:      user.Username,
		Authenticated: true,
	}

	mail := models.Mail{
		Id:        0,
		Sender:    "test2",
		Addressee: session.Username,
		Theme:     "test",
		Text:      "test",
		Files:     "files",
		Date:      time.Now(),
		Read:      false,
	}
	mailBytes, _ := easyjson.Marshal(mail)

	var mailId int32 = mail.Id
	read := true

	mockDB.EXPECT().GetMailInfoById(context.Background(), &repository_proto.GetMailInfoByIdRequest{
		MailId: mailId,
	}).Return(&repository_proto.ResponseMail{
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
		Mail: mailBytes,
	}, nil)

	mockDB.EXPECT().ReadMail(context.Background(), &repository_proto.ReadMailRequest{
		Mail: mailBytes,
		Read: read,
	}).Return(&utils_proto.DatabaseResponse{
		Status: utils_proto.DatabaseStatus_OK,
	}, nil)

	resp, err := uc.ReadMail(context.Background(), &mailbox_proto.ReadMailRequest{
		Data: &session,
		Id:   mailId,
		Read: read,
	})

	var response pkg.JsonResponse
	json_err := easyjson.Unmarshal(resp.Response, &response)
	if err != nil || json_err != nil || response != pkg.NO_ERR {
		t.Errorf("Неверный ответ от UseCase.")
		return
	}
}

func TestSendMail(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB, _, uc := InitTestUseCase(mockCtrl)

	user := models.User{
		Id:        0,
		Firstname: "test",
		Lastname:  "test",
		Username:  "test",
		Password:  "test",
	}
	userBytes, _ := easyjson.Marshal(user)

	session := utils_proto.Session{
		Username:      user.Username,
		Authenticated: true,
	}

	mail := models.Mail{
		Id:        0,
		Sender:    session.Username,
		Addressee: session.Username,
		Theme:     "test",
		Text:      "test",
		Files:     "files",
		Date:      time.Now(),
		Read:      false,
	}
	mailBytes, _ := easyjson.Marshal(mail)

	form := models.MailForm{
		Addressee: user.Username,
		Theme:     mail.Theme,
		Text:      mail.Text,
		Files:     mail.Files,
	}
	formBytes, _ := easyjson.Marshal(form)

	mockDB.EXPECT().GetUserInfoByUsername(context.Background(), &repository_proto.GetUserInfoByUsernameRequest{
		Username: user.Username,
	}).Return(&repository_proto.ResponseUser{
		User: userBytes,
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
	}, nil)

	mockDB.EXPECT().GetUserInfoByUsername(context.Background(), &repository_proto.GetUserInfoByUsernameRequest{
		Username: user.Username,
	}).Return(&repository_proto.ResponseUser{
		User: userBytes,
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
	}, nil)

	mockDB.EXPECT().AddMail(context.Background(), &repository_proto.AddMailRequest{
		Mail: mailBytes,
	}).Return(&utils_proto.DatabaseExtendResponse{
		Status: utils_proto.DatabaseStatus_OK,
		Param:  "1",
	}, nil)

	resp, err := uc.SendMail(context.Background(), &mailbox_proto.SendMailRequest{
		Data: &session,
		Form: formBytes,
	})

	var response pkg.JsonResponse
	json_err := easyjson.Unmarshal(resp.Response, &response)
	if err != nil || json_err != nil || response != pkg.NO_ERR {
		t.Errorf("Неверный ответ от UseCase.")
		return
	}
}
