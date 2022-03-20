package usecase

import (
	"OverflowBackend/internal/models"
	"OverflowBackend/pkg"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Получение информации о пользователе.
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

// Установка аватарки пользователя.
func (uc *UseCase) SetAvatar(data *models.Session, avatar *models.Avatar) error {
	format := data.Email + "_" + avatar.Name
	if err := os.MkdirAll(uc.config.Server.Static.Dir, os.ModePerm); err != nil {
		return err
	}
	path := filepath.Join(uc.config.Server.Static.Dir, format)
	err := os.WriteFile(path, avatar.Content, 0644)
	if (err != nil) {
		return err
	}
	return nil
}

// Установка настроек пользователя.
func (uc *UseCase) SetInfo(data *models.Session, settings *models.SettingsForm) error {
	if settings.FirstName != "" {
		return fmt.Errorf("Не имплементировано.")
	}
	if settings.LastName != "" {
		return fmt.Errorf("Не имплементировано.")
	}
	if settings.Password != "" {
		user, err := uc.db.GetUserInfoByEmail(data.Email)
		if err != nil {
			return err
		}
		err = uc.db.ChangeUserPassword(user, pkg.HashPassword(settings.Password))
		if err != nil {
			return err
		}
	}
	return nil
}

// Получение ссылки на аватарку пользователя.
func (uc *UseCase) GetAvatar(data *models.Session) (avatarUrl string, err error) {
	matches, err := filepath.Glob(filepath.Join(uc.config.Server.Static.Dir, data.Email+"_*"))
	if err != nil {
		return
	}
	if len(matches) == 0 {
		avatarUrl = uc.config.Server.Static.Handle + "dummy.png"
		return
	}
	if len(matches) > 1 {
		err = fmt.Errorf("Найдены дубликаты файлов.")
		return
	}
	avatarUrl = uc.config.Server.Static.Handle + filepath.Base(matches[0])
	return
}
