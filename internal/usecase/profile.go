package usecase

import (
	"OverflowBackend/internal/models"
	"encoding/json"
)

func (uc *UseCase) GetInfo(data *models.Session) (userJson []byte, err error) {
	user, err := uc.db.GetUserInfoByEmail(data.Email)
	if err != nil {
		return
	}

	userJson, err = json.Marshal(user)
	if err != nil {
		return
	}
	return
}