package profile

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/internal/models"
	"OverflowBackend/pkg"
	"OverflowBackend/proto/profile_proto"
	"OverflowBackend/proto/repository_proto"
	"OverflowBackend/proto/utils_proto"
	"context"
	"github.com/mailru/easyjson"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

type ProfileService struct {
	config *config.Config
	db     repository_proto.DatabaseRepositoryClient
}

func (s *ProfileService) Init(config *config.Config, db repository_proto.DatabaseRepositoryClient) {
	s.config = config
	s.db = db
}

// Получение информации о пользователе.
func (s *ProfileService) GetInfo(context context.Context, request *profile_proto.GetInfoRequest) (*profile_proto.GetInfoResponse, error) {
	resp, err := s.db.GetUserInfoByUsername(context, &repository_proto.GetUserInfoByUsernameRequest{Username: request.Data.Username})
	log.Info("Получение информации о пользователе: ", request.Data.Username)
	if err != nil {
		log.Error(err)
		return &profile_proto.GetInfoResponse{
			Response: &utils_proto.JsonResponse{
				Response: pkg.DB_ERR.Bytes(),
			},
			Data: nil,
		}, err
	}
	var user models.User
	err = easyjson.Unmarshal(resp.User, &user)
	if err != nil {
		return &profile_proto.GetInfoResponse{
			Response: &utils_proto.JsonResponse{
				Response: pkg.JSON_ERR.Bytes(),
			},
			Data: nil,
		}, err
	}
	var info models.ProfileInfo
	info.Id = user.Id
	info.Firstname = user.Firstname
	info.Lastname = user.Lastname
	info.Username = user.Username
	infoBytes, _ := easyjson.Marshal(info)
	return &profile_proto.GetInfoResponse{
		Response: &utils_proto.JsonResponse{
			Response: pkg.NO_ERR.Bytes(),
		},
		Data: infoBytes,
	}, nil
}

// Установка аватарки пользователя.
func (s *ProfileService) SetAvatar(context context.Context, request *profile_proto.SetAvatarRequest) (*utils_proto.JsonResponse, error) {
	log.Info("Установка аватарки")
	// создание папки с файлами, если она не существует
	if err := os.MkdirAll(s.config.Server.Static.Dir, os.ModePerm); err != nil {
		log.Error(err)
		return &utils_proto.JsonResponse{
			Response: pkg.CreateJsonErr(pkg.STATUS_UNKNOWN, "Ошибка создания папки.").Bytes(),
		}, err
	}
	var avatar models.Avatar
	err := easyjson.Unmarshal(request.Avatar, &avatar)
	if err != nil {
		return &utils_proto.JsonResponse{
			Response: pkg.JSON_ERR.Bytes(),
		}, err
	}
	// хеширование имени пользователя и имени аватара, приведение к соответствующему формату
	hashedUser := pkg.HashString(request.Data.Username)
	format := hashedUser + "_" + pkg.HashString(avatar.Name) + filepath.Ext(avatar.Name)
	matches, err := filepath.Glob(filepath.Join(s.config.Server.Static.Dir, hashedUser+"_*"))
	if err != nil {
		log.Error(err)
		return &utils_proto.JsonResponse{
			Response: pkg.CreateJsonErr(pkg.STATUS_UNKNOWN, "Ошибка поиска файла.").Bytes(),
		}, err
	}
	path := filepath.Join(s.config.Server.Static.Dir, format)
	err = os.WriteFile(path, avatar.File, 0644)
	if err != nil {
		log.Error(err)
		return &utils_proto.JsonResponse{
			Response: pkg.CreateJsonErr(pkg.STATUS_UNKNOWN, "Ошибка записи в файл.").Bytes(),
		}, err
	}
	// если остались старые файлы (старые автарки), то удаляем их
	if len(matches) > 0 {
		for _, match := range matches {
			os.Remove(match)
		}
	}
	return &utils_proto.JsonResponse{
		Response: pkg.NO_ERR.Bytes(),
	}, nil
}

// Установка настроек пользователя.
func (s *ProfileService) SetInfo(context context.Context, request *profile_proto.SetInfoRequest) (*utils_proto.JsonResponse, error) {
	data := request.Data
	log.Debug("Установка настроек пользователя: ", data.Username)
	var settings models.ProfileSettingsForm
	err := easyjson.Unmarshal(request.Form, &settings)
	if err != nil {
		return &utils_proto.JsonResponse{
			Response: pkg.JSON_ERR.Bytes(),
		}, err
	}
	if (settings == models.ProfileSettingsForm{}) {
		return &utils_proto.JsonResponse{
			Response: pkg.NO_ERR.Bytes(),
		}, nil
	}
	resp, err := s.db.GetUserInfoByUsername(context, &repository_proto.GetUserInfoByUsernameRequest{Username: data.Username})
	if err != nil {
		log.Error(err)
		return &utils_proto.JsonResponse{
			Response: pkg.DB_ERR.Bytes(),
		}, err
	}
	if resp.Response.Status != utils_proto.DatabaseStatus_OK {
		return &utils_proto.JsonResponse{
			Response: pkg.DB_ERR.Bytes(),
		}, nil
	}
	user := resp.User
	var resp2 *utils_proto.DatabaseResponse
	if settings.Firstname != "" {
		resp2, err = s.db.ChangeUserFirstName(
			context,
			&repository_proto.ChangeForm{User: user, Data: settings.Firstname},
		)
		if err != nil {
			log.Error(err)
			return &utils_proto.JsonResponse{
				Response: pkg.DB_ERR.Bytes(),
			}, err
		}
		if resp2.Status != utils_proto.DatabaseStatus_OK {
			return &utils_proto.JsonResponse{
				Response: pkg.DB_ERR.Bytes(),
			}, nil
		}
	}
	if settings.Lastname != "" {
		resp2, err = s.db.ChangeUserLastName(
			context,
			&repository_proto.ChangeForm{User: user, Data: settings.Lastname},
		)
		if err != nil {
			log.Error(err)
			return &utils_proto.JsonResponse{
				Response: pkg.DB_ERR.Bytes(),
			}, err
		}
		if resp2.Status != utils_proto.DatabaseStatus_OK {
			return &utils_proto.JsonResponse{
				Response: pkg.DB_ERR.Bytes(),
			}, nil
		}
	}
	return &utils_proto.JsonResponse{
		Response: pkg.NO_ERR.Bytes(),
	}, nil
}

// Получение ссылки на аватарку пользователя.
func (s *ProfileService) GetAvatar(context context.Context, request *profile_proto.GetAvatarRequest) (*profile_proto.GetAvatarResponse, error) {
	matches, e := filepath.Glob(filepath.Join(s.config.Server.Static.Dir, pkg.HashString(request.Username)+"_*"))
	if e != nil {
		log.Error(e)
		return &profile_proto.GetAvatarResponse{
			Response: &utils_proto.JsonResponse{
				Response: pkg.CreateJsonErr(pkg.STATUS_UNKNOWN, "Ошибка поиска файла.").Bytes(),
			},
			Url: "",
		}, e
	}
	if len(matches) == 0 {
		avatarUrl := pkg.JoinURL(s.config.Server.Static.Handle, request.DummyName+".png")
		return &profile_proto.GetAvatarResponse{
			Response: &utils_proto.JsonResponse{
				Response: pkg.NO_ERR.Bytes(),
			},
			Url: avatarUrl,
		}, nil
	}
	if len(matches) > 1 {
		return &profile_proto.GetAvatarResponse{
			Response: &utils_proto.JsonResponse{
				Response: pkg.CreateJsonErr(pkg.STATUS_UNKNOWN, "Найдены дубликаты файлов.").Bytes(),
			},
			Url: "",
		}, nil
	}
	avatarUrl := pkg.JoinURL(s.config.Server.Static.Handle, filepath.Base(matches[0]))
	return &profile_proto.GetAvatarResponse{
		Response: &utils_proto.JsonResponse{
			Response: pkg.NO_ERR.Bytes(),
		},
		Url: avatarUrl,
	}, nil
}

func (s *ProfileService) ChangePassword(context context.Context, request *profile_proto.ChangePasswordRequest) (*utils_proto.JsonResponse, error) {
	log.Debug("Изменение пароля ползователя, username = ", request.Data.Username, ", password_old = ", request.PasswordOld, ", password_new = ", request.PasswordNew)
	resp, err := s.db.GetUserInfoByUsername(context, &repository_proto.GetUserInfoByUsernameRequest{Username: request.Data.Username})
	if err != nil {
		log.Error(err)
		return &utils_proto.JsonResponse{
			Response: pkg.DB_ERR.Bytes(),
		}, err
	}
	if resp.Response.Status != utils_proto.DatabaseStatus_OK {
		return &utils_proto.JsonResponse{
			Response: pkg.DB_ERR.Bytes(),
		}, nil
	}
	userBytes := resp.User
	var user models.User
	err = easyjson.Unmarshal(userBytes, &user)
	if err != nil {
		return &utils_proto.JsonResponse{
			Response: pkg.JSON_ERR.Bytes(),
		}, err
	}
	if user.Password != pkg.HashPassword(request.PasswordOld) {
		return &utils_proto.JsonResponse{
			Response: pkg.CreateJsonErr(pkg.STATUS_WRONG_CREDS, "Неверный пароль.").Bytes(),
		}, nil
	}
	resp3, err := s.db.ChangeUserPassword(context, &repository_proto.ChangeForm{
		User: userBytes,
		Data: pkg.HashPassword(request.PasswordNew),
	})
	if err != nil {
		return &utils_proto.JsonResponse{
			Response: pkg.DB_ERR.Bytes(),
		}, err
	}
	if resp3.Status != utils_proto.DatabaseStatus_OK {
		return &utils_proto.JsonResponse{
			Response: pkg.DB_ERR.Bytes(),
		}, nil
	}
	return &utils_proto.JsonResponse{
		Response: pkg.NO_ERR.Bytes(),
	}, nil
}
