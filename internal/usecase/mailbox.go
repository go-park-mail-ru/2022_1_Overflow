package usecase

import (
	"OverflowBackend/internal/models"
	"OverflowBackend/pkg"
	"encoding/json"
	"time"

	log "github.com/sirupsen/logrus"
)

func (uc *UseCase) Income(data *models.Session) ([]byte, pkg.JsonResponse) {
	user, err := uc.db.GetUserInfoByUsername(data.Username)
	log.Debug("Получение входящих писем, username = ", data.Username)
	if err != nil {
		log.Error(err)
		return nil, pkg.DB_ERR
	}
	id := user.Id
	mails, err := uc.db.GetIncomeMails(id)
	if err != nil {
		log.Error(err)
		return nil, pkg.DB_ERR
	}
	var mails_add []models.MailAdditional
	for _, mail := range mails {
		mail_add := models.MailAdditional{}
		mail_add.Mail = mail
		avatarUrl, resp := uc.GetAvatar(mail.Sender)
		if resp != pkg.NO_ERR {
			return nil, resp
		}
		mail_add.AvatarUrl = avatarUrl
		mails_add = append(mails_add, mail_add)
	}
	parsed, err := json.Marshal(mails_add)
	if err != nil {
		log.Error(err)
		return nil, pkg.JSON_ERR
	}
	return parsed, pkg.NO_ERR
}

func (uc *UseCase) Outcome(data *models.Session) ([]byte, pkg.JsonResponse) {
	user, err := uc.db.GetUserInfoByUsername(data.Username)
	log.Debug("Получение исходящих писем, username = ", data.Username)
	if err != nil {
		log.Error(err)
		return nil, pkg.DB_ERR
	}
	id := user.Id
	mails, err := uc.db.GetOutcomeMails(id)
	if err != nil {
		log.Error(err)
		return nil, pkg.DB_ERR
	}
	var mails_add []models.MailAdditional
	for _, mail := range mails {
		mail_add := models.MailAdditional{}
		mail_add.Mail = mail
		avatarUrl, resp := uc.GetAvatar(mail.Addressee)
		if resp != pkg.NO_ERR {
			return nil, resp
		}
		mail_add.AvatarUrl = avatarUrl
		mails_add = append(mails_add, mail_add)
	}
	parsed, err := json.Marshal(mails_add)
	if err != nil {
		log.Error(err)
		return nil, pkg.JSON_ERR
	}
	return parsed, pkg.NO_ERR
}

func (uc *UseCase) GetMail(data *models.Session, mail_id int32) ([]byte, pkg.JsonResponse) {
	log.Debug("Получение письма, mail_id = ", mail_id)
	mail, err := uc.db.GetMailInfoById(mail_id)
	if err != nil {
		log.Error(err)
		return nil, pkg.DB_ERR
	}
	if mail.Addressee != data.Username && mail.Sender != data.Username {
		return nil, pkg.UNAUTHORIZED_ERR
	}
	parsed, err := json.Marshal(mail)
	if err != nil {
		log.Error(err)
		return nil, pkg.JSON_ERR
	}
	return parsed, pkg.NO_ERR
}

func (uc *UseCase) DeleteMail(data *models.Session, id int32) pkg.JsonResponse {
	mail, err := uc.db.GetMailInfoById(id)
	log.Debug("Удаление письма, id = ", id)
	if err != nil {
		log.Error(err)
		return pkg.DB_ERR
	}
	if mail.Addressee != data.Username && mail.Sender != data.Username {
		return pkg.UNAUTHORIZED_ERR
	}
	err = uc.db.DeleteMail(mail, data.Username)
	if err != nil {
		log.Error(err)
		return pkg.DB_ERR
	}
	return pkg.NO_ERR
}

func (uc *UseCase) ReadMail(data *models.Session, id int32) pkg.JsonResponse {
	log.Debug("Прочитать письмо, id = ", id)
	mail, err := uc.db.GetMailInfoById(id)
	if err != nil {
		log.Error(err)
		return pkg.DB_ERR
	}
	if mail.Addressee != data.Username {
		return pkg.UNAUTHORIZED_ERR
	}
	err = uc.db.ReadMail(mail)
	if err != nil {
		log.Error(err)
		return pkg.DB_ERR
	}
	return pkg.NO_ERR
}

func (uc *UseCase) SendMail(data *models.Session, form models.MailForm) pkg.JsonResponse {
	log.Debug("Отправить письмо, username = ", data.Username)
	user, err := uc.db.GetUserInfoByUsername(data.Username)
	if err != nil {
		log.Error(err)
		return pkg.DB_ERR
	}
	if (user == models.User{}) {
		return pkg.DB_ERR
	}
	userAddressee, err := uc.db.GetUserInfoByUsername(form.Addressee)
	if err != nil {
		log.Error(err)
		return pkg.DB_ERR
	}
	if (userAddressee == models.User{}) {
		return pkg.NO_USER_EXIST
	}
	mail := models.Mail{
		Client_id: user.Id,
		Sender:    data.Username,
		Addressee: form.Addressee,
		Theme:     form.Theme,
		Text:      form.Text,
		Files:     form.Files,
		Date:      time.Now(),
	}
	err = uc.db.AddMail(mail)
	if err != nil {
		log.Error(err)
		return pkg.DB_ERR
	}
	return pkg.NO_ERR
}
