package usecase

import (
	"OverflowBackend/internal/models"
	"OverflowBackend/pkg"
	"encoding/json"
	"log"
	"time"
)

func (uc *UseCase) Income(data *models.Session) ([]byte, pkg.JsonResponse) {
	user, err := uc.db.GetUserInfoByEmail(data.Email)
	if err != nil {
		log.Println(err)
		return nil, pkg.DB_ERR
	}
	id := user.Id
	mails, err := uc.db.GetIncomeMails(id)
	if err != nil {
		log.Println(err)
		return nil, pkg.DB_ERR
	}
	parsed, err := json.Marshal(mails)
	if err != nil {
		log.Println(err)
		return nil, pkg.JSON_ERR
	}
	return parsed, pkg.NO_ERR
}

func (uc *UseCase) Outcome(data *models.Session) ([]byte, pkg.JsonResponse) {
	user, err := uc.db.GetUserInfoByEmail(data.Email)
	if err != nil {
		log.Println(err)
		return nil, pkg.DB_ERR
	}
	id := user.Id
	mails, err := uc.db.GetIncomeMails(id)
	if err != nil {
		log.Println(err)
		return nil, pkg.DB_ERR
	}
	parsed, err := json.Marshal(mails)
	if err != nil {
		log.Println(err)
		return nil, pkg.JSON_ERR
	}
	return parsed, pkg.NO_ERR
}

func (uc *UseCase) DeleteMail(data *models.Session, id int) pkg.JsonResponse {
	mail, err := uc.db.GetMailInfoById(id)
	if err != nil {
		log.Println(err)
		return pkg.DB_ERR
	}
	if mail.Addressee != data.Email && mail.Sender != data.Email {
		return pkg.UNAUTHORIZED_ERR
	}
	err = uc.db.DeleteMail(mail)
	if err != nil {
		log.Println(err)
		return pkg.DB_ERR
	}
	return pkg.NO_ERR
}

func (uc *UseCase) ReadMail(data *models.Session, id int) pkg.JsonResponse {
	mail, err := uc.db.GetMailInfoById(id)
	if err != nil {
		log.Println(err)
		return pkg.DB_ERR
	}
	if mail.Addressee != data.Email {
		return pkg.UNAUTHORIZED_ERR
	}
	err = uc.db.ReadMail(mail)
	if err != nil {
		log.Println(err)
		return pkg.DB_ERR
	}
	return pkg.NO_ERR
}

func (uc *UseCase) SendMail(data *models.Session, form models.MailForm) pkg.JsonResponse {
	user, err := uc.db.GetUserInfoByEmail(data.Email)
	if err != nil {
		log.Println(err)
		return pkg.DB_ERR
	}
	mail := models.Mail {
		Client_id: user.Id,
		Sender: data.Email,
		Addressee: form.Addressee,
		Theme: form.Theme,
		Text: form.Text,
		Files: form.Files,
		Date: time.Now(),
	}
	err = uc.db.AddMail(mail)
	if err != nil {
		log.Println(err)
		return pkg.DB_ERR
	}
	return pkg.NO_ERR
}