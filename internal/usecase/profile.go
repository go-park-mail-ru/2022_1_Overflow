package usecase

import (
	"OverflowBackend/internal/models"
	"encoding/json"
	"os"
	"path/filepath"
)

// Получение настроек пользователя.
func (uc *UseCase) GetInfo(data *models.Session) (settingsJson []byte, err error) {
	user, err := uc.db.GetUserInfoByEmail(data.Email)
	if err != nil {
		return
	}

	settings := models.SettingsForm{}
	settings.User = user

	settingsJson, err = json.Marshal(settings)
	if err != nil {
		return
	}
	return
}

// Установка аватарки пользователя.
func (uc *UseCase) SetAvatar(data *models.Session, avatar *models.Avatar) error {
	format := data.Email + "_" + avatar.Name
	dirPath := filepath.Join("..", "static")
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return err
	}
	path := filepath.Join(dirPath, format)
	err := os.WriteFile(path, avatar.Content, 0644)
	if (err != nil) {
		return err
	}
	return nil
}

// Установка настроек пользователя.
func (uc *UseCase) SetInfo(settings *models.SettingsForm) error {
	// пока нет доступа к изменению в БД
	return nil
}