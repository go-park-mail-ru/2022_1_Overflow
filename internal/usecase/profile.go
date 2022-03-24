package usecase

import (
	"OverflowBackend/internal/models"
	"OverflowBackend/pkg"
	"encoding/json"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

// Получение информации о пользователе.
func (uc *UseCase) GetInfo(data *models.Session) ([]byte, pkg.JsonResponse) {

	user, err := uc.db.GetUserInfoByEmail(data.Email)
	log.Info("Получение информации о пользователе: ", data.Email)
	if err != nil {
		log.Error(err)
		return nil, pkg.DB_ERR
	}

	userJson, err := json.Marshal(user)
	if err != nil {
		log.Error(err)
		return nil, pkg.JSON_ERR
	}
	return userJson, pkg.NO_ERR
}

// Установка аватарки пользователя.
func (uc *UseCase) SetAvatar(data *models.Session, avatar *models.Avatar) pkg.JsonResponse {
	log.Info("Установка аватарки")
	format := data.Email + "_" + avatar.Name
	if err := os.MkdirAll(uc.config.Server.Static.Dir, os.ModePerm); err != nil {
		log.Error(err)
		return pkg.CreateJsonErr(pkg.STATUS_UNKNOWN, "Ошибка создания папки.")
	}
	path := filepath.Join(uc.config.Server.Static.Dir, format)
	err := os.WriteFile(path, avatar.Content, 0644)
	if err != nil {
		log.Error(err)
		return pkg.CreateJsonErr(pkg.STATUS_UNKNOWN, "Ошибка записи в файл.")
	}
	return pkg.NO_ERR
}

// Установка настроек пользователя.
func (uc *UseCase) SetInfo(data *models.Session, settings *models.SettingsForm) pkg.JsonResponse {

	if settings.FirstName != "" {
		return pkg.NOT_IMPLEMENTED_ERR
	}
	if settings.LastName != "" {
		return pkg.NOT_IMPLEMENTED_ERR
	}
	if settings.Password != "" {
		log.Info("Установка настроек пользователя: ", data.Email)
		user, err := uc.db.GetUserInfoByEmail(data.Email)
		if err != nil {
			log.Error(err)
			return pkg.DB_ERR
		}
		err = uc.db.ChangeUserPassword(user, pkg.HashPassword(settings.Password))
		if err != nil {
			log.Error(err)
			return pkg.DB_ERR
		}
	}
	return pkg.NO_ERR
}

// Получение ссылки на аватарку пользователя.
func (uc *UseCase) GetAvatar(data *models.Session) (string, pkg.JsonResponse) {
	matches, e := filepath.Glob(filepath.Join(uc.config.Server.Static.Dir, data.Email+"_*"))
	if e != nil {
		log.Error(e)
		return "", pkg.CreateJsonErr(pkg.STATUS_UNKNOWN, "Ошибка поиска файла.")
	}
	if len(matches) == 0 {
		avatarUrl := uc.config.Server.Static.Handle + "dummy.png"
		return avatarUrl, pkg.NO_ERR
	}
	if len(matches) > 1 {
		return "", pkg.CreateJsonErr(pkg.STATUS_UNKNOWN, "Найдены дубликаты файлов.")
	}
	avatarUrl := uc.config.Server.Static.Handle + filepath.Base(matches[0])
	return avatarUrl, pkg.NO_ERR
}
