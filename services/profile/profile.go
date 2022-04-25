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
func (s *ProfileService) GetInfo(context context.Context, request *profile_proto.GetInfoRequest) (*profile_proto.GetInfoResponse, error) {
	user, err := s.db.GetUserInfoByUsername(context, &repository_proto.GetUserInfoByUsernameRequest{Username: request.Data.Username})
	log.Info("Получение информации о пользователе: ", request.Data.Username)
	if err != nil {
		log.Error(err)
		return &profile_proto.GetInfoResponse{Response: &pkg.DB_ERR, Data: nil}, err
	}

	userJson, err := json.Marshal(user)
	if err != nil {
		log.Error(err)
		return  &profile_proto.GetInfoResponse{Response: &pkg.JSON_ERR, Data: nil}, err
	}
	return &profile_proto.GetInfoResponse{Response: &pkg.NO_ERR, Data: userJson}, nil
}

// Установка аватарки пользователя.
func (s *ProfileService) SetAvatar(context context.Context, request *profile_proto.SetAvatarRequest) (*utils_proto.JsonResponse, error) {
	log.Info("Установка аватарки")
	// создание папки с файлами, если она не существует
	if err := os.MkdirAll(s.config.Server.Static.Dir, os.ModePerm); err != nil {
		log.Error(err)
		return pkg.CreateJsonErr(pkg.STATUS_UNKNOWN, "Ошибка создания папки."), err
	}
	// хеширование имени пользователя и имени аватара, приведение к соответствующему формату
	hashedUser := pkg.HashString(request.Data.Username)
	format := hashedUser + "_" + pkg.HashString(request.Avatar.Name) + filepath.Ext(request.Avatar.Name)
	matches, e := filepath.Glob(filepath.Join(s.config.Server.Static.Dir, hashedUser+"_*"))
	if e != nil {
		log.Error(e)
		return pkg.CreateJsonErr(pkg.STATUS_UNKNOWN, "Ошибка поиска файла."), e
	}
	path := filepath.Join(s.config.Server.Static.Dir, format)
	err := os.WriteFile(path, request.Avatar.File, 0644)
	if err != nil {
		log.Error(err)
		return pkg.CreateJsonErr(pkg.STATUS_UNKNOWN, "Ошибка записи в файл."), err
	}
	// если остались старые файлы (старые автарки), то удаляем их
	if len(matches) > 0 {
		for _, match := range matches {
			os.Remove(match)
		}
	}
	return &pkg.NO_ERR, nil
}

// Установка настроек пользователя.
func (s *ProfileService) SetInfo(context context.Context, request *profile_proto.SetInfoRequest) (*utils_proto.JsonResponse, error) {
	data := request.Data
	settings := request.Form
	log.Debug("Установка настроек пользователя: ", data.Username)
	if (proto.Equal(settings, &profile_proto.SettingsForm{})) {
		return &pkg.NO_ERR, nil
	}
	resp, err := s.db.GetUserInfoByUsername(context, &repository_proto.GetUserInfoByUsernameRequest{Username: data.Username})
	if err != nil {
		log.Error(err)
		return &pkg.DB_ERR, err
	}
	if resp.Response.Status != utils_proto.DatabaseStatus_OK {
		return &pkg.DB_ERR, err
	}
	user := resp.User
	var resp2 *utils_proto.DatabaseResponse
	if settings.FirstName != "" {
		resp2, err = s.db.ChangeUserFirstName(
			context,
			&repository_proto.ChangeForm{User: user, Data: settings.FirstName},
		)
		if err != nil {
			log.Error(err)
			return &pkg.DB_ERR, err
		}
		if resp2.Status != utils_proto.DatabaseStatus_OK {
			return &pkg.DB_ERR, nil
		}
	}
	if settings.LastName != "" {
		resp2, err = s.db.ChangeUserLastName(
			context,
			&repository_proto.ChangeForm{User: user, Data: settings.LastName},
		)
		if err != nil {
			log.Error(err)
			return &pkg.DB_ERR, err
		}
		if resp2.Status != utils_proto.DatabaseStatus_OK {
			return &pkg.DB_ERR, nil
		}
	}
	if settings.Password != "" {
		resp2, err = s.db.ChangeUserPassword(
			context,
			&repository_proto.ChangeForm{User: user, Data: pkg.HashPassword(settings.Password)},
		)
		if err != nil {
			log.Error(err)
			return &pkg.DB_ERR, err
		}
		if resp2.Status != utils_proto.DatabaseStatus_OK {
			return &pkg.DB_ERR, nil
		}
	}
	return &pkg.NO_ERR, nil
}

// Получение ссылки на аватарку пользователя.
func (s *ProfileService) GetAvatar(context context.Context, request *profile_proto.GetAvatarRequest) (*profile_proto.GetAvatarResponse, error) {
	matches, e := filepath.Glob(filepath.Join(s.config.Server.Static.Dir, pkg.HashString(request.Username)+"_*"))
	if e != nil {
		log.Error(e)
		return &profile_proto.GetAvatarResponse{
			Response: pkg.CreateJsonErr(pkg.STATUS_UNKNOWN, "Ошибка поиска файла."),
			Url: "",
		}, e
	}
	if len(matches) == 0 {
		avatarUrl := pkg.JoinURL(s.config.Server.Static.Handle, "dummy.png")
		return &profile_proto.GetAvatarResponse{
			Response: &pkg.NO_ERR,
			Url: avatarUrl,
		}, nil
	}
	if len(matches) > 1 {
		return &profile_proto.GetAvatarResponse{
			Response: pkg.CreateJsonErr(pkg.STATUS_UNKNOWN, "Найдены дубликаты файлов."),
			Url: "",
		}, nil
	}
	avatarUrl := pkg.JoinURL(s.config.Server.Static.Handle, filepath.Base(matches[0]))
	return &profile_proto.GetAvatarResponse{
		Response: &pkg.NO_ERR,
		Url: avatarUrl,
	}, nil
}
