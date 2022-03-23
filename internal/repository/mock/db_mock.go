package mock

import (
	"OverflowBackend/internal/models"
	"errors"
	"time"
)

type MockDB struct{
	user []map[string]interface{}
	mail []map[string]interface{}
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

func (m *MockDB) GetUserInfoByEmail(email string) (models.User, error) {
	for _, val := range m.user {
		if val["email"] == email {
			return models.User{
				Id: val["id"].(int32),
				FirstName: val["first_name"].(string),
				LastName: val["last_name"].(string),
				Password: val["password"].(string),
				Email: email,
			}, nil
		}
	}
	return models.User{}, errors.New("Пользователь не найден.")
}

func (m *MockDB) GetUserInfoById(userId int32) (models.User, error) {
	for _, val := range m.user {
		if val["id"] == userId {
			return models.User{
				Id: userId,
				FirstName: val["first_name"].(string),
				LastName: val["last_name"].(string),
				Password: val["password"].(string),
				Email: val["email"].(string),
			}, nil
		}
	}
	return models.User{}, errors.New("Пользователь не найден.")
}

func (m *MockDB) AddUser(user models.User) error {
	u := map[string]interface{} {
		"id": user.Id,
		"first_name": user.FirstName,
		"last_name": user.LastName,
		"email": user.Email,
		"password": user.Password,
	}
	m.user = append(m.user, u)
	return nil
}

func (m *MockDB) ChangeUserPassword(user models.User, newPassword string) error {
	for i, val := range m.user {
		if val["email"] == user.Email {
			m.user[i]["password"] = newPassword
			return nil
		}
	}
	return errors.New("Пользователь не найден.")
}

func (m *MockDB) AddMail(email models.Mail) error {
	mail := map[string]interface{} {
		"id": 0, // потому что поле не используется (пока что)
		"client_id": email.Client_id,
		"sender": email.Sender,
		"addressee": email.Addressee,
		"date": email.Date,
		"theme": email.Theme,
		"text": email.Text,
		"files": email.Files,
		"read": email.Read,
	}
	m.mail = append(m.mail, mail)
	return nil
}

func (m *MockDB) DeleteMail(email models.Mail) error {
	for i, val := range m.mail {
		if val["id"] == email.Id {
			m.mail[i] = m.mail[len(m.mail)-1]
			m.mail = m.mail[:len(m.mail)-1]
			return nil
		}
	}
	return nil
}

func (m *MockDB) ReadMail(email models.Mail) error {
	for _, val := range m.mail {
		if val["id"] == email.Id {
			val["read"] = true
			return nil
		}
	}
	return errors.New("Письмо не найдено.")
}

func (m *MockDB) GetMailInfoById(mailId int) (models.Mail, error) {
	for _, val := range m.mail {
		if val["id"] == mailId {
			mail := models.Mail{
				Client_id: val["client_id"].(int32),
				Sender: val["sender"].(string),
				Addressee: val["addressee"].(string),
				Date: val["date"].(time.Time),
				Theme: val["theme"].(string),
				Text: val["text"].(string),
				Files: val["files"].(string),
				Read: val["read"].(bool),
			}
			return mail, nil
		}
	}
	return models.Mail{}, errors.New("Письмо не найдено.")
}

func (m *MockDB) GetIncomeMails(userId int32) ([]models.Mail, error) {
	var mails []models.Mail
	user, err := m.GetUserInfoById(userId)
	if err != nil {
		return mails, err
	}
	for _, val := range m.mail {
		if val["addressee"] == user.Email {
			mail := models.Mail{
				Client_id: val["client_id"].(int32),
				Sender: val["sender"].(string),
				Addressee: val["addressee"].(string),
				Date: val["date"].(time.Time),
				Theme: val["theme"].(string),
				Text: val["text"].(string),
				Files: val["files"].(string),
				Read: val["read"].(bool),
			}
			mails = append(mails, mail)
		}
	}
	return mails, nil
}

func (m *MockDB) GetOutcomeMails(userId int32) ([]models.Mail, error) {
	var mails []models.Mail
	user, err := m.GetUserInfoById(userId)
	if err != nil {
		return mails, err
	}
	for _, val := range m.mail {
		if val["sender"] == user.Email {
			mail := models.Mail{
				Client_id: val["client_id"].(int32),
				Sender: val["sender"].(string),
				Addressee: val["addressee"].(string),
				Date: val["date"].(time.Time),
				Theme: val["theme"].(string),
				Text: val["text"].(string),
				Files: val["files"].(string),
				Read: val["read"].(bool),
			}
			mails = append(mails, mail)
		}
	}
	return mails, nil
}