package usecase

import (
	"OverflowBackend/internal/models"
	"encoding/json"
	"fmt"
)

func (uc *UseCase) Income(data *models.Session) (parsed []byte, err error) {
	user, err := uc.db.GetUserInfoByEmail(data.Email)
	if err != nil {
		return
	}
	id := user.Id
	mails, err := uc.db.GetIncomeMails(id)
	if err != nil {
		return
	}
	parsed, err = json.Marshal(mails)
	if err != nil {
		return
	}
	return parsed, nil
}

func (uc *UseCase) Outcome(data *models.Session) (parsed []byte, err error) {
	user, err := uc.db.GetUserInfoByEmail(data.Email)
	if err != nil {
		return
	}
	id := user.Id
	mails, err := uc.db.GetIncomeMails(id)
	if err != nil {
		return
	}
	parsed, err = json.Marshal(mails)
	if err != nil {
		return
	}
	return parsed, nil
}

func (uc *UseCase) DeleteMail(data *models.Session, id int) error {
	_, err := uc.db.GetMailInfoById(id)
	if err != nil {
		return err
	}
	// тут удаление письма из БД
	return nil
}

func (uc *UseCase) ReadMail(data *models.Session, id int) error {
	mail, err := uc.db.GetMailInfoById(id)
	if err != nil {
		return err
	}
	if mail.Addressee != data.Email {
		return fmt.Errorf("Письмо не принадлежит запрашивающему пользователю.")
	}
	err = uc.db.ReadMail(mail)
	if err != nil {
		return err
	}
	return nil
}