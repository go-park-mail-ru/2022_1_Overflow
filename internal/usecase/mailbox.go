package usecase

import (
	"OverflowBackend/internal/models"
	"encoding/json"
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

