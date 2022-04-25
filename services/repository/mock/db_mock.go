package mock

import (
	"OverflowBackend/proto/repository_proto"
	"OverflowBackend/proto/utils_proto"
	"time"

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

func (m *MockDB) GetUserInfoByUsername(request *repository_proto.GetUserInfoByUsernameRequest) *repository_proto.ResponseUser {
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
			}
		}
	}
	return &repository_proto.ResponseUser{
		User: &user,
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_ERROR,
		},
	}
}

func (m *MockDB) GetUserInfoById(request *repository_proto.GetUserInfoByIdRequest) *repository_proto.ResponseUser{
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
			}
		}
	}
	return &repository_proto.ResponseUser{
		User: &user,
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_ERROR,
		},
	}
}

func (m *MockDB) AddUser(request *repository_proto.AddUserRequest) *utils_proto.DatabaseResponse {
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
	}
}

func (m *MockDB) ChangeUserPassword(request *repository_proto.ChangeForm) *utils_proto.DatabaseResponse {
	user := request.User
	for i, val := range m.user {
		if val["username"] == user.Username {
			m.user[i]["password"] = request.Data
			return &utils_proto.DatabaseResponse{
				Status: utils_proto.DatabaseStatus_OK,
			}
		}
	}
	return &utils_proto.DatabaseResponse{
		Status: utils_proto.DatabaseStatus_ERROR,
	}
}

func (m *MockDB) ChangeUserFirstName(request *repository_proto.ChangeForm) *utils_proto.DatabaseResponse {
	user := request.User
	for i, val := range m.user {
		if val["username"] == user.Username {
			m.user[i]["first_name"] = request.Data
			return &utils_proto.DatabaseResponse{
				Status: utils_proto.DatabaseStatus_OK,
			}
		}
	}
	return &utils_proto.DatabaseResponse{
		Status: utils_proto.DatabaseStatus_ERROR,
	}
}

func (m *MockDB) ChangeUserLastName(request *repository_proto.ChangeForm) *utils_proto.DatabaseResponse {
	user := request.User
	for i, val := range m.user {
		if val["username"] == user.Username {
			m.user[i]["last_name"] = request.Data
			return &utils_proto.DatabaseResponse{
				Status: utils_proto.DatabaseStatus_OK,
			}
		}
	}
	return &utils_proto.DatabaseResponse{
		Status: utils_proto.DatabaseStatus_ERROR,
	}
}

func (m *MockDB) AddMail(request *repository_proto.AddMailRequest) *utils_proto.DatabaseResponse {
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
	}
}

func (m *MockDB) DeleteMail(request *repository_proto.DeleteMailRequest) *utils_proto.DatabaseResponse {
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
			}
		}
	}
	// письмо не найдено
	return &utils_proto.DatabaseResponse{
		Status: utils_proto.DatabaseStatus_ERROR,
	}
}

func (m *MockDB) ReadMail(request *repository_proto.ReadMailRequest) *utils_proto.DatabaseResponse {
	mail := request.Mail
	for _, val := range m.mail {
		if val["id"].(int32) == mail.Id {
			val["read"] = true
			return &utils_proto.DatabaseResponse{
				Status: utils_proto.DatabaseStatus_OK,
			}
		}
	}
	return &utils_proto.DatabaseResponse{
		Status: utils_proto.DatabaseStatus_ERROR,
	}
}

func (m *MockDB) GetMailInfoById(request *repository_proto.GetMailInfoByIdRequest) *repository_proto.ResponseMail {
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
			}
		}
	}
	return &repository_proto.ResponseMail{
		Mail: &mail,
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_ERROR,
		},
	}
}

func (m *MockDB) GetIncomeMails(request *repository_proto.GetIncomeMailsRequest) *repository_proto.ResponseMails {
	var mails []*utils_proto.Mail
	resp := m.GetUserInfoById(&repository_proto.GetUserInfoByIdRequest{UserId: request.UserId})
	if resp.Response.Status != utils_proto.DatabaseStatus_OK {
		return &repository_proto.ResponseMails{
			Mails: mails,
			Response: &utils_proto.DatabaseResponse{
				Status: utils_proto.DatabaseStatus_ERROR,
			},
		}
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
	}
}

func (m *MockDB) GetOutcomeMails(request *repository_proto.GetOutcomeMailsRequest) *repository_proto.ResponseMails {
	var mails []*utils_proto.Mail
	resp := m.GetUserInfoById(&repository_proto.GetUserInfoByIdRequest{UserId: request.UserId})
	if resp.Response.Status != utils_proto.DatabaseStatus_OK {
		return &repository_proto.ResponseMails{
			Mails: mails,
			Response: &utils_proto.DatabaseResponse{
				Status: utils_proto.DatabaseStatus_ERROR,
			},
		}
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
	}
}
