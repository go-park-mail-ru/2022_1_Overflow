package usecase_test

import (
	"OverflowBackend/internal/models"
	"OverflowBackend/pkg"
	"encoding/json"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
)

func TestIncome(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB, uc := InitTestUseCase(mockCtrl)

	mail := models.Mail{
		Id: 0,
		Client_id: 0,
		Sender: "test",
		Addressee: "test2",
		Theme: "test",
		Text: "test",
		Files: "files",
		Date: time.Now(),
		Read: false,
	}

	mailAdd := models.MailAdditional{
		Mail: mail,
		AvatarUrl: "/static/dummy.png",
	}

	mails := []models.Mail{
		mail, mail, mail, mail, mail,
	}

	mockDB.EXPECT().GetUserInfoByUsername("test").Return(models.User{
		Id:        0,
		FirstName: "test",
		LastName:  "test",
		Username:  "test",
		Password:  "test",
	}, nil)
	mockDB.EXPECT().GetIncomeMails(int32(0)).Return(mails, nil)

	resp, r := uc.Income(&models.Session{Username: "test", Authenticated: true})
	if r != pkg.NO_ERR {
		t.Errorf("Неверный ответ от UseCase.")
		return
	}

	mailsUC := make([]models.MailAdditional, 5)

	err := json.Unmarshal(resp, &mailsUC)
	if err != nil {
		t.Error(err)
		return
	}

	if len(mailsUC) != 5 {
		t.Errorf("Неверное количество сообщений. Получено: %v, ожидалось: %v", len(mailsUC), 5)
		return
	}

	mailsUC[0].Mail.Date = mailAdd.Mail.Date
	if mailsUC[0] != mailAdd {
		t.Errorf("Сообщение не соответствует ожидаемому. Получено: %v, ожидается: %v.", mailsUC[0], mailAdd)
	}
}

func TestOutcome(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB, uc := InitTestUseCase(mockCtrl)

	mail := models.Mail{
		Id: 0,
		Client_id: 0,
		Sender: "test",
		Addressee: "test2",
		Theme: "test",
		Text: "test",
		Files: "files",
		Date: time.Now(),
		Read: false,
	}

	mailAdd := models.MailAdditional{
		Mail: mail,
		AvatarUrl: "/static/dummy.png",
	}

	mails := []models.Mail{
		mail, mail, mail, mail, mail,
	}

	mockDB.EXPECT().GetUserInfoByUsername("test").Return(models.User{
		Id:        0,
		FirstName: "test",
		LastName:  "test",
		Username:  "test",
		Password:  "test",
	}, nil)
	mockDB.EXPECT().GetOutcomeMails(int32(0)).Return(mails, nil)

	resp, r := uc.Outcome(&models.Session{Username: "test", Authenticated: true})
	if r != pkg.NO_ERR {
		t.Errorf("Неверный ответ от UseCase.")
		return
	}

	mailsUC := make([]models.MailAdditional, 5)

	err := json.Unmarshal(resp, &mailsUC)
	if err != nil {
		t.Error(err)
		return
	}

	if len(mailsUC) != 5 {
		t.Errorf("Неверное количество сообщений. Получено: %v, ожидалось: %v", len(mailsUC), 5)
		return
	}

	mailsUC[0].Mail.Date = mailAdd.Mail.Date
	if mailsUC[0] != mailAdd {
		t.Errorf("Сообщение не соответствует ожидаемому. Получено: %v, ожидается: %v.", mailsUC[0], mailAdd)
	}
}

func TestGetMail(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB, uc := InitTestUseCase(mockCtrl)

	mail := models.Mail{
		Id:        0,
		Client_id: 0,
		Sender:    "test",
		Addressee: "test2",
		Theme:     "test",
		Text:      "test",
		Files:     "files",
		Date:      time.Now(),
		Read:      false,
	}

	mockDB.EXPECT().GetMailInfoById(int32(0)).Return(mail, nil)

	mailBytes, r := uc.GetMail(&models.Session{Username: "test", Authenticated: true}, int32(0))
	if r != pkg.NO_ERR {
		t.Errorf("Неверный ответ от UseCase.")
		return
	}

	mailUC := models.Mail{}

	err := json.Unmarshal(mailBytes, &mailUC)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestDeleteMail(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB, uc := InitTestUseCase(mockCtrl)

	mail := models.Mail{
		Id:        0,
		Client_id: 0,
		Sender:    "test",
		Addressee: "test2",
		Theme:     "test",
		Text:      "test",
		Files:     "files",
		Date:      time.Now(),
		Read:      false,
	}

	session := models.Session{
		Username:      "test",
		Authenticated: true,
	}

	mockDB.EXPECT().GetMailInfoById(int32(0)).Return(mail, nil)
	mockDB.EXPECT().DeleteMail(mail, session.Username).Return(nil)

	r := uc.DeleteMail(&session, int32(0))
	if r != pkg.NO_ERR {
		t.Errorf("Неверный ответ от UseCase.")
		return
	}
}

func TestReadMail(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB, uc := InitTestUseCase(mockCtrl)

	mail := models.Mail{
		Id:        0,
		Client_id: 0,
		Sender:    "test",
		Addressee: "test2",
		Theme:     "test",
		Text:      "test",
		Files:     "files",
		Date:      time.Now(),
		Read:      false,
	}

	session := models.Session{
		Username:      "test2",
		Authenticated: true,
	}

	mockDB.EXPECT().GetMailInfoById(int32(0)).Return(mail, nil)
	mockDB.EXPECT().ReadMail(mail).Return(nil)

	r := uc.ReadMail(&session, int32(0))
	if r != pkg.NO_ERR {
		t.Errorf("Неверный ответ от UseCase.")
		return
	}
}

func TestSendMail(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB, uc := InitTestUseCase(mockCtrl)

	session := models.Session{
		Username:      "test",
		Authenticated: true,
	}

	user := models.User{
		Id:        0,
		FirstName: "test",
		LastName:  "test",
		Password:  "test",
		Username:  session.Username,
	}

	mail := models.Mail{
		Id:        0,
		Client_id: 0,
		Sender:    user.Username,
		Addressee: "test2",
		Theme:     "test",
		Text:      "test",
		Files:     "files",
		Date:      time.Now(),
		Read:      false,
	}

	form := models.MailForm{
		Addressee: mail.Addressee,
		Theme:     mail.Theme,
		Text:      mail.Text,
		Files:     mail.Files,
	}

	mockDB.EXPECT().GetUserInfoByUsername(session.Username).Return(user, nil)
	mockDB.EXPECT().AddMail(mail).Return(nil)

	r := uc.SendMail(&session, form)
	if r != pkg.NO_ERR {
		t.Errorf("Неверный ответ от UseCase.")
		return
	}
}

func TestForwardMail(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB, uc := InitTestUseCase(mockCtrl)

	mail1 := models.Mail{
		Id:        int32(0),
		Client_id: 0,
		Sender:    "test",
		Addressee: "test2",
		Theme:     "test",
		Text:      "test",
		Files:     "test",
		Date:      time.Now(),
		Read:      false,
	}

	user := models.User{
		Id:        0,
		FirstName: "test2",
		LastName:  "test2",
		Username:  "test2",
		Password:  "test2",
	}

	session := models.Session{
		Username:      user.Username,
		Authenticated: true,
	}

	form := models.MailForm{
		Addressee: "test3",
		Theme:     "test",
		Text:      "test",
		Files:     "files",
	}

	mailWrapped := pkg.MailWrapper(form, mail1)

	mail2 := models.Mail{
		Sender: user.Username,
		Addressee: mailWrapped.Addressee,
		Theme: mailWrapped.Theme,
		Text: mailWrapped.Text,
		Files: mailWrapped.Files,
		Date: time.Now(),
	}

	mockDB.EXPECT().GetMailInfoById(mail1.Id).Return(mail1, nil)
	mockDB.EXPECT().GetUserInfoByUsername(user.Username).Return(user, nil)
	mockDB.EXPECT().AddMail(mail2).Return(nil)

	r := uc.ForwardMail(&session, form, int32(0))
	if r != pkg.NO_ERR {
		t.Errorf("Неверный ответ от UseCase.")
		return
	}
}

func TestRespondMail(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB, uc := InitTestUseCase(mockCtrl)

	session := models.Session{
		Username: "test",
		Authenticated: true,
	}

	user := models.User{
		Id:        0,
		FirstName: "test",
		LastName:  "test",
		Username:  "test",
		Password:  "test",
	}

	mail1 := models.Mail{
		Id: int32(0),
		Client_id: 0,
		Sender: "test2",
		Addressee: "test",
		Theme: "test",
		Text: "test",
		Files: "files",
		Date: time.Now(),
		Read: false,
	}

	form := models.MailForm{
		Addressee: mail1.Sender,
		Theme: "response",
		Text: "test",
		Files: "files",
	}

	mailWrapped := pkg.MailWrapper(form, mail1)

	mail2 := models.Mail{
		Sender: user.Username,
		Addressee: mailWrapped.Addressee,
		Theme: mailWrapped.Theme,
		Text: mailWrapped.Text,
		Files: mailWrapped.Files,
		Date: time.Now(),
	}

	mockDB.EXPECT().GetMailInfoById(mail1.Id).Return(mail1, nil)
	mockDB.EXPECT().GetUserInfoByUsername(user.Username).Return(user, nil)
	mockDB.EXPECT().AddMail(mail2).Return(nil)

	r := uc.RespondMail(&session, form, mail1.Id)
	if r != pkg.NO_ERR {
		t.Errorf("Неверный ответ от UseCase.")
		return
	}
}
