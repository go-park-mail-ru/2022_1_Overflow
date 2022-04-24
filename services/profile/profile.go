package profile

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/pkg"
	"OverflowBackend/proto/profile_proto"
	"OverflowBackend/proto/repository_proto"
	"OverflowBackend/proto/utils_proto"
	"context"
	"encoding/json"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
	log "github.com/sirupsen/logrus"
)

type ProfileService struct {
	config *config.Config
	db repository_proto.DatabaseRepositoryClient
}

func (s *ProfileService) Init(config *config.Config, db repository_proto.DatabaseRepositoryClient) {
	s.config = config
	s.db = db
}

// Получение информации о пользователе.
func (s *ProfileService) GetInfo(data *utils_proto.Session) *profile_proto.GetInfoResponse {
	user, err := s.db.GetUserInfoByUsername(context.Background(), &repository_proto.GetUserInfoByUsernameRequest{Username: data.Username})
	log.Info("Получение информации о пользователе: ", data.Username)
	if err != nil {
		log.Error(err)
		return &profile_proto.GetInfoResponse{Response: &pkg.DB_ERR, Data: nil}
	}

	userJson, err := json.Marshal(user)
	if err != nil {
		log.Error(err)
		return  &profile_proto.GetInfoResponse{Response: &pkg.JSON_ERR, Data: nil}
	}
	return &profile_proto.GetInfoResponse{Response: &pkg.NO_ERR, Data: userJson}
}

// Установка аватарки пользователя.
func (s *ProfileService) SetAvatar(request *profile_proto.SetAvatarRequest) *utils_proto.JsonResponse {
	log.Info("Установка аватарки")
	// создание папки с файлами, если она не существует
	if err := os.MkdirAll(s.config.Server.Static.Dir, os.ModePerm); err != nil {
		log.Error(err)
		return pkg.CreateJsonErr(pkg.STATUS_UNKNOWN, "Ошибка создания папки.")
	}
	// хеширование имени пользователя и имени аватара, приведение к соответствующему формату
	hashedUser := pkg.HashString(request.Data.Username)
	format := hashedUser + "_" + pkg.HashString(request.Avatar.Name) + filepath.Ext(request.Avatar.Name)
	matches, e := filepath.Glob(filepath.Join(s.config.Server.Static.Dir, hashedUser+"_*"))
	if e != nil {
		log.Error(e)
		return pkg.CreateJsonErr(pkg.STATUS_UNKNOWN, "Ошибка поиска файла.")
	}
	path := filepath.Join(s.config.Server.Static.Dir, format)
	err := os.WriteFile(path, request.Avatar.File, 0644)
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
	return &pkg.NO_ERR
}

// Установка настроек пользователя.
func (s *ProfileService) SetInfo(request *profile_proto.SetInfoRequest) *utils_proto.JsonResponse {
	data := request.Data
	settings := request.Form
	log.Debug("Установка настроек пользователя: ", data.Username)
	if (proto.Equal(settings, &profile_proto.SettingsForm{})) {
		return &pkg.NO_ERR
	}
	resp, err := s.db.GetUserInfoByUsername(context.Background(), &repository_proto.GetUserInfoByUsernameRequest{Username: data.Username})
	if err != nil {
		log.Error(err)
		return &pkg.DB_ERR
	}
	if resp.Response.Status != utils_proto.DatabaseStatus_OK {
		return &pkg.DB_ERR
	}
	user := resp.User
	var resp2 *utils_proto.DatabaseResponse
	if settings.FirstName != "" {
		resp2, err = s.db.ChangeUserFirstName(
			context.Background(),
			&repository_proto.ChangeForm{User: user, Data: settings.FirstName},
		)
		if err != nil {
			log.Error(err)
			return &pkg.DB_ERR
		}
		if resp2.Status != utils_proto.DatabaseStatus_OK {
			return &pkg.DB_ERR
		}
	}
	if settings.LastName != "" {
		resp2, err = s.db.ChangeUserLastName(
			context.Background(),
			&repository_proto.ChangeForm{User: user, Data: settings.LastName},
		)
		if err != nil {
			log.Error(err)
			return &pkg.DB_ERR
		}
		if resp2.Status != utils_proto.DatabaseStatus_OK {
			return &pkg.DB_ERR
		}
	}
	if settings.Password != "" {
		resp2, err = s.db.ChangeUserPassword(
			context.Background(),
			&repository_proto.ChangeForm{User: user, Data: pkg.HashPassword(settings.Password)},
		)
		if err != nil {
			log.Error(err)
			return &pkg.DB_ERR
		}
	}
	return &pkg.NO_ERR
}

// Получение ссылки на аватарку пользователя.
func (s *ProfileService) GetAvatar(request *profile_proto.GetAvatarRequest) *profile_proto.GetAvatarResponse {
	matches, e := filepath.Glob(filepath.Join(s.config.Server.Static.Dir, pkg.HashString(request.Username)+"_*"))
	if e != nil {
		log.Error(e)
		return &profile_proto.GetAvatarResponse{
			Response: pkg.CreateJsonErr(pkg.STATUS_UNKNOWN, "Ошибка поиска файла."),
			Url: "",
		}
	}
	if len(matches) == 0 {
		avatarUrl := pkg.JoinURL(s.config.Server.Static.Handle, "dummy.png")
		return &profile_proto.GetAvatarResponse{
			Response: &pkg.NO_ERR,
			Url: avatarUrl,
		}
	}
	if len(matches) > 1 {
		return &profile_proto.GetAvatarResponse{
			Response: pkg.CreateJsonErr(pkg.STATUS_UNKNOWN, "Найдены дубликаты файлов."),
			Url: "",
		}
	}
	avatarUrl := pkg.JoinURL(s.config.Server.Static.Handle, filepath.Base(matches[0]))
	return &profile_proto.GetAvatarResponse{
		Response: &pkg.NO_ERR,
		Url: avatarUrl,
	}
}
