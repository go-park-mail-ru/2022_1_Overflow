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

	user, err := uc.db.GetUserInfoByUsername(data.Username)
	log.Info("Получение информации о пользователе: ", data.Username)
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
	// создание папки с файлами, если она не существует
	if err := os.MkdirAll(uc.config.Server.Static.Dir, os.ModePerm); err != nil {
		log.Error(err)
		return pkg.CreateJsonErr(pkg.STATUS_UNKNOWN, "Ошибка создания папки.")
	}
	// хеширование имени пользователя и имени аватара, приведение к соответствующему формату
	hashedUser := pkg.HashString(data.Username)
	format := hashedUser + "_" + pkg.HashString(avatar.Name) + filepath.Ext(avatar.Name)
	matches, e := filepath.Glob(filepath.Join(uc.config.Server.Static.Dir, hashedUser+"_*"))
	if e != nil {
		log.Error(e)
		return pkg.CreateJsonErr(pkg.STATUS_UNKNOWN, "Ошибка поиска файла.")
	}
	path := filepath.Join(uc.config.Server.Static.Dir, format)
	err := os.WriteFile(path, avatar.Content, 0644)
	if err != nil {
		log.Error(err)
		return pkg.CreateJsonErr(pkg.STATUS_UNKNOWN, "Ошибка записи в файл.")
	}
	// если остались старые файлы (старые автарки), то удаляем их
	if len(matches) > 0 {
		for _, match := range matches {
			os.Remove(match)
		}
	}
	return pkg.NO_ERR
}

// Установка настроек пользователя.
func (uc *UseCase) SetInfo(data *models.Session, settings *models.SettingsForm) pkg.JsonResponse {
	log.Debug("Установка настроек пользователя: ", data.Username)
	if (*settings == models.SettingsForm{}) {
		return pkg.NO_ERR
	}
	user, err := uc.db.GetUserInfoByUsername(data.Username)
	if err != nil {
		log.Error(err)
		return pkg.DB_ERR
	}
	if settings.FirstName != "" {
		err = uc.db.ChangeUserFirstName(user, settings.FirstName)
		if err != nil {
			log.Error(err)
			return pkg.DB_ERR
		}
	}
	if settings.LastName != "" {
		err = uc.db.ChangeUserLastName(user, settings.LastName)
		if err != nil {
			log.Error(err)
			return pkg.DB_ERR
		}
	}
	if settings.Password != "" {
		err = uc.db.ChangeUserPassword(user, pkg.HashPassword(settings.Password))
		if err != nil {
			log.Error(err)
			return pkg.DB_ERR
		}
	}
	return pkg.NO_ERR
}

// Получение ссылки на аватарку пользователя.
func (uc *UseCase) GetAvatar(username string) (string, pkg.JsonResponse) {
	matches, e := filepath.Glob(filepath.Join(uc.config.Server.Static.Dir, pkg.HashString(username)+"_*"))
	if e != nil {
		log.Error(e)
		return "", pkg.CreateJsonErr(pkg.STATUS_UNKNOWN, "Ошибка поиска файла.")
	}
	if len(matches) == 0 {
		avatarUrl := pkg.JoinURL(uc.config.Server.Static.Handle, "dummy.png")
		return avatarUrl, pkg.NO_ERR
	}
	if len(matches) > 1 {
		return "", pkg.CreateJsonErr(pkg.STATUS_UNKNOWN, "Найдены дубликаты файлов.")
	}
	avatarUrl := pkg.JoinURL(uc.config.Server.Static.Handle, filepath.Base(matches[0]))
	return avatarUrl, pkg.NO_ERR
}
