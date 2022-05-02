package mock

import (
	"OverflowBackend/proto/repository_proto"
	"OverflowBackend/proto/utils_proto"
	"context"
	"errors"
	"time"
	log "github.com/sirupsen/logrus"

	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type MockDB struct {
	user []map[string]interface{}
	mail []map[string]interface{}
	user_id int32
	mail_id int32
}

func (m *MockDB) Create(url string) error {
	m.user = []map[string]interface{}{}
	m.mail = []map[string]interface{}{}
	return nil
}

func (m *MockDB) Fill(data map[string][]map[string]interface{}) {
	m.user = data["user"]
	m.mail = data["mail"]
}

func (m *MockDB) GetUserInfoByUsername(context context.Context, request *repository_proto.GetUserInfoByUsernameRequest) (*repository_proto.ResponseUser, error) {
	log.Info("MockDB: вызов GetUserInfoByUsername")
	username := request.Username
	var user utils_proto.User
	for _, val := range m.user {
		if val["username"] == username {
			user = utils_proto.User{
				Id:        val["id"].(int32),
				FirstName: val["first_name"].(string),
				LastName:  val["last_name"].(string),
				Password:  val["password"].(string),
				Username:  username,
			}
			return &repository_proto.ResponseUser{
				User: &user,
				Response: &utils_proto.DatabaseResponse{
					Status: utils_proto.DatabaseStatus_OK,
				},
			}, nil
		}
	}
	return &repository_proto.ResponseUser{
		User: &user,
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_ERROR,
		},
	}, nil //errors.New("Пользователь не найден.")
}

func (m *MockDB) GetUserInfoById(context context.Context, request *repository_proto.GetUserInfoByIdRequest) (*repository_proto.ResponseUser, error) {
	log.Info("MockDB: вызов GetUserInfoById")
	userId := request.UserId
	var user utils_proto.User
	for _, val := range m.user {
		if val["id"] == userId {
			user = utils_proto.User{
				Id:        userId,
				FirstName: val["first_name"].(string),
				LastName:  val["last_name"].(string),
				Password:  val["password"].(string),
				Username:  val["username"].(string),
			}
			return &repository_proto.ResponseUser{
				User: &user,
				Response: &utils_proto.DatabaseResponse{
					Status: utils_proto.DatabaseStatus_OK,
				},
			}, nil
		}
	}
	return &repository_proto.ResponseUser{
		User: &user,
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_ERROR,
		},
	}, nil //errors.New("Пользователь не найден.")
}

func (m *MockDB) AddUser(context context.Context, request *repository_proto.AddUserRequest) (*utils_proto.DatabaseResponse, error) {
	log.Info("MockDB: вызов AddUser")
	user := request.User
	u := map[string]interface{}{
		"id":         m.user_id,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"username":   user.Username,
		"password":   user.Password,
	}
	m.user_id++
	m.user = append(m.user, u)
	return &utils_proto.DatabaseResponse{
		Status: utils_proto.DatabaseStatus_OK,
	}, nil
}

func (m *MockDB) ChangeUserPassword(context context.Context, request *repository_proto.ChangeForm) (*utils_proto.DatabaseResponse, error) {
	log.Info("MockDB: вызов ChangeUserPassword")
	user := request.User
	for i, val := range m.user {
		if val["username"] == user.Username {
			m.user[i]["password"] = request.Data
			return &utils_proto.DatabaseResponse{
				Status: utils_proto.DatabaseStatus_OK,
			}, nil
		}
	}
	return &utils_proto.DatabaseResponse{
		Status: utils_proto.DatabaseStatus_ERROR,
	}, errors.New("Пользователь не найден.")
}

func (m *MockDB) ChangeUserFirstName(context context.Context, request *repository_proto.ChangeForm) (*utils_proto.DatabaseResponse, error) {
	log.Info("MockDB: вызов ChangeUserFirstName")
	user := request.User
	for i, val := range m.user {
		if val["username"] == user.Username {
			m.user[i]["first_name"] = request.Data
			return &utils_proto.DatabaseResponse{
				Status: utils_proto.DatabaseStatus_OK,
			}, nil
		}
	}
	return &utils_proto.DatabaseResponse{
		Status: utils_proto.DatabaseStatus_ERROR,
	}, errors.New("Пользователь не найден.")
}

func (m *MockDB) ChangeUserLastName(context context.Context, request *repository_proto.ChangeForm) (*utils_proto.DatabaseResponse, error) {
	log.Info("MockDB: вызов ChangeUserLastName")
	user := request.User
	for i, val := range m.user {
		if val["username"] == user.Username {
			m.user[i]["last_name"] = request.Data
			return &utils_proto.DatabaseResponse{
				Status: utils_proto.DatabaseStatus_OK,
			}, nil
		}
	}
	return &utils_proto.DatabaseResponse{
		Status: utils_proto.DatabaseStatus_ERROR,
	}, errors.New("Пользователь не найден.")
}

func (m *MockDB) AddMail(context context.Context, request *repository_proto.AddMailRequest) (*utils_proto.DatabaseResponse, error) {
	log.Info("MockDB: вызов AddMail")
	mailInfo := request.Mail
	mail := map[string]interface{}{
		"id":        m.mail_id, // потому что поле не используется (пока что)
		"client_id": mailInfo.ClientId,
		"sender":    mailInfo.Sender,
		"addressee": mailInfo.Addressee,
		"date":      mailInfo.Date.AsTime(),
		"theme":     mailInfo.Theme,
		"text":      mailInfo.Text,
		"files":     mailInfo.Files,
		"read":      mailInfo.Read.Value,
	}
	m.mail_id++
	m.mail = append(m.mail, mail)
	return &utils_proto.DatabaseResponse{
		Status: utils_proto.DatabaseStatus_OK,
	}, nil
}

func (m *MockDB) DeleteMail(context context.Context, request *repository_proto.DeleteMailRequest) (*utils_proto.DatabaseResponse, error) {
	log.Info("MockDB: вызов DeleteMail")
	mail := request.Mail
	username := request.Username
	for i, val := range m.mail {
		if val["id"] == mail.Id {
			switch {
			case val["addressee"] == username:
				m.mail[i]["addressee"] = ""
			case val["sender"] == username:
				m.mail[i]["sender"] = ""
			default:
				{
					m.mail[i] = m.mail[len(m.mail)-1]
					m.mail = m.mail[:len(m.mail)-1]
				}
			}
			return &utils_proto.DatabaseResponse{
				Status: utils_proto.DatabaseStatus_OK,
			}, nil
		}
	}
	// письмо не найдено
	return &utils_proto.DatabaseResponse{
		Status: utils_proto.DatabaseStatus_ERROR,
	}, errors.New("Письмо не найдено.")
}

func (m *MockDB) ReadMail(context context.Context, request *repository_proto.ReadMailRequest) (*utils_proto.DatabaseResponse, error) {
	log.Info("MockDB: вызов ReadMail")
	mail := request.Mail
	for _, val := range m.mail {
		if val["id"].(int32) == mail.Id {
			val["read"] = true
			return &utils_proto.DatabaseResponse{
				Status: utils_proto.DatabaseStatus_OK,
			}, nil
		}
	}
	return &utils_proto.DatabaseResponse{
		Status: utils_proto.DatabaseStatus_ERROR,
	}, errors.New("Письмо не найдено.")
}

func (m *MockDB) GetMailInfoById(context context.Context, request *repository_proto.GetMailInfoByIdRequest) (*repository_proto.ResponseMail, error) {
	log.Info("MockDB: вызов GetMailInfoById")
	mailId := request.MailId
	var mail utils_proto.Mail
	for _, val := range m.mail {
		if val["id"].(int32) == mailId {
			mail = utils_proto.Mail{
				ClientId: val["client_id"].(int32),
				Sender:    val["sender"].(string),
				Addressee: val["addressee"].(string),
				Date:      timestamppb.New(val["date"].(time.Time)),
				Theme:     val["theme"].(string),
				Text:      val["text"].(string),
				Files:     val["files"].(string),
				Read:      wrapperspb.Bool(val["read"].(bool)),
			}
			return &repository_proto.ResponseMail{
				Mail: &mail,
				Response: &utils_proto.DatabaseResponse{
					Status: utils_proto.DatabaseStatus_OK,
				},
			}, nil
		}
	}
	return &repository_proto.ResponseMail{
		Mail: &mail,
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_ERROR,
		},
	}, errors.New("Письмо не найдено.")
}

func (m *MockDB) GetIncomeMails(context context.Context, request *repository_proto.GetIncomeMailsRequest) (*repository_proto.ResponseMails, error) {
	log.Info("MockDB: вызов GetIncomeMails")
	var mails []*utils_proto.Mail
	resp, err := m.GetUserInfoById(context, &repository_proto.GetUserInfoByIdRequest{UserId: request.UserId})
	if err != nil || resp.Response.Status != utils_proto.DatabaseStatus_OK {
		return &repository_proto.ResponseMails{
			Mails: mails,
			Response: &utils_proto.DatabaseResponse{
				Status: utils_proto.DatabaseStatus_ERROR,
			},
		}, errors.New("Пользователь не найден.")
	}
	user := resp.User
	for _, val := range m.mail {
		if val["addressee"] == user.Username {
			mail := utils_proto.Mail{
				ClientId: val["client_id"].(int32),
				Sender:    val["sender"].(string),
				Addressee: val["addressee"].(string),
				Date:      timestamppb.New(val["date"].(time.Time)),
				Theme:     val["theme"].(string),
				Text:      val["text"].(string),
				Files:     val["files"].(string),
				Read:      wrapperspb.Bool(val["read"].(bool)),
			}
			mails = append(mails, &mail)
		}
	}
	return &repository_proto.ResponseMails{
		Mails: mails,
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
	}, nil
}

func (m *MockDB) GetOutcomeMails(context context.Context, request *repository_proto.GetOutcomeMailsRequest) (*repository_proto.ResponseMails, error) {
	log.Info("MockDB: вызов GetOutcomeMails")
	var mails []*utils_proto.Mail
	resp, err := m.GetUserInfoById(context, &repository_proto.GetUserInfoByIdRequest{UserId: request.UserId})
	if err != nil || resp.Response.Status != utils_proto.DatabaseStatus_OK {
		return &repository_proto.ResponseMails{
			Mails: mails,
			Response: &utils_proto.DatabaseResponse{
				Status: utils_proto.DatabaseStatus_ERROR,
			},
		}, errors.New("Пользователь не найден.")
	}
	user := resp.User
	for _, val := range m.mail {
		if val["sender"] == user.Username {
			mail := utils_proto.Mail{
				ClientId: val["client_id"].(int32),
				Sender:    val["sender"].(string),
				Addressee: val["addressee"].(string),
				Date:      timestamppb.New(val["date"].(time.Time)),
				Theme:     val["theme"].(string),
				Text:      val["text"].(string),
				Files:     val["files"].(string),
				Read:      wrapperspb.Bool(val["read"].(bool)),
			}
			mails = append(mails, &mail)
		}
	}
	return &repository_proto.ResponseMails{
		Mails: mails,
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
	}, nil
}
