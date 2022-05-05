package profile_test

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/internal/models"
	"OverflowBackend/pkg"
	"OverflowBackend/proto/profile_proto"
	"OverflowBackend/proto/repository_proto"
	"OverflowBackend/proto/utils_proto"
	"OverflowBackend/services/profile"
	"context"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"testing"

	"bou.ke/monkey"
	"github.com/golang/mock/gomock"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func InitTestUseCase(ctrl *gomock.Controller) (*repository_proto.MockDatabaseRepositoryClient, *profile.ProfileService) {
	monkey.Patch(os.WriteFile, func(name string, data []byte, perm fs.FileMode) error {return nil})
	monkey.Patch(os.MkdirAll, func(path string, perm fs.FileMode) error { return nil })
	monkey.Patch(filepath.Glob, func(pattern string) (matches []string, err error) {return []string{}, nil})
	log.SetLevel(log.FatalLevel)
	db := repository_proto.NewMockDatabaseRepositoryClient(ctrl)
	uc := profile.ProfileService{}
	uc.Init(config.TestConfig(), db)
	return db, &uc
}

func TestGetInfo(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB, uc := InitTestUseCase(mockCtrl)

	user := models.User{
		Id:        0,
		Firstname: "test",
		Lastname:  "test",
		Password:  "test",
		Username:  "test",
	}
	userBytes, _ := json.Marshal(user)

	profileInfo := models.ProfileInfo{
		Id: user.Id,
		Firstname: user.Firstname,
		Lastname: user.Lastname,
		Username: user.Username,
	}

	session := utils_proto.Session {
		Username:      user.Username,
		Authenticated: wrapperspb.Bool(true),
	}

	mockDB.EXPECT().GetUserInfoByUsername(context.Background(), &repository_proto.GetUserInfoByUsernameRequest{
		Username: session.Username,
	}).Return(&repository_proto.ResponseUser{
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
		User: userBytes,
	}, nil)

	var response pkg.JsonResponse
	resp, err := uc.GetInfo(context.Background(), &profile_proto.GetInfoRequest{
		Data: &session,
	})
	json_err := json.Unmarshal(resp.Response.Response, &response)
	if err != nil || json_err != nil || response != pkg.NO_ERR {
		t.Errorf("Неверный ответ от UseCase.")
		return
	}

	profileInfoResp := models.ProfileInfo{}

	err = json.Unmarshal(resp.Data, &profileInfoResp)
	if err != nil {
		t.Error(err)
		return
	}

	if profileInfoResp != profileInfo {
		t.Errorf("Информация о пользователе не соответствует заданной. Получено: %v, ожидается: %v.", profileInfoResp, profileInfo)
		return
	}
}

func TestSetInfo(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB, uc := InitTestUseCase(mockCtrl)

	user := models.User{
		Id:        0,
		Firstname: "test",
		Lastname:  "test",
		Password:  "test",
		Username:  "test",
	}
	userBytes, _ := json.Marshal(user)

	settings := models.ProfileSettingsForm{
		Firstname: user.Firstname+"test",
		Lastname:  user.Lastname+"test",
	}
	settingsBytes, _ := json.Marshal(settings)

	session := utils_proto.Session{
		Username:      user.Username,
		Authenticated: wrapperspb.Bool(true),
	}

	mockDB.EXPECT().GetUserInfoByUsername(context.Background(), &repository_proto.GetUserInfoByUsernameRequest{
		Username: session.Username,
	}).Return(&repository_proto.ResponseUser{
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
		User: userBytes,
	}, nil)
	mockDB.EXPECT().ChangeUserFirstName(context.Background(), &repository_proto.ChangeForm{
		User: userBytes,
		Data: settings.Firstname,
	}).Return(&utils_proto.DatabaseResponse{
		Status: utils_proto.DatabaseStatus_OK,
	}, nil)
	mockDB.EXPECT().ChangeUserLastName(context.Background(), &repository_proto.ChangeForm{
		User: userBytes,
		Data: settings.Lastname,
	}).Return(&utils_proto.DatabaseResponse{
		Status: utils_proto.DatabaseStatus_OK,
	}, nil)

	resp, err := uc.SetInfo(context.Background(), &profile_proto.SetInfoRequest{
		Data: &session,
		Form: settingsBytes,
	})
	var response pkg.JsonResponse
	json_err := json.Unmarshal(resp.Response, &response)
	if err != nil || json_err != nil || response != pkg.NO_ERR {
		t.Errorf("Неверный ответ от UseCase.")
		return
	}
}

func TestSetAvatar(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	_, uc := InitTestUseCase(mockCtrl)
	
	session := utils_proto.Session{
		Username:      "test",
		Authenticated: wrapperspb.Bool(true),
	}

	avatar := models.Avatar{
		Name:      "avatar",
		Username: session.Username,
		File:   []byte{10, 10, 10, 10},
	}
	avatarBytes, _ := json.Marshal(avatar)

	var response pkg.JsonResponse
	resp, err := uc.SetAvatar(context.Background(), &profile_proto.SetAvatarRequest{
		Data: &session,
		Avatar: avatarBytes,
	})
	json_err := json.Unmarshal(resp.Response, &response)
	if err != nil || json_err != nil || response != pkg.NO_ERR {
		t.Errorf("Неверный ответ от UseCase.")
		return
	}

	monkey.Patch(os.MkdirAll, func(path string, perm fs.FileMode) error { return fmt.Errorf("Ошибка.") })
	resp, _ = uc.SetAvatar(context.Background(), &profile_proto.SetAvatarRequest{
		Data: &session,
		Avatar: avatarBytes,
	})
	json_err = json.Unmarshal(resp.Response, &response)
	if json_err != nil || response.Status != pkg.STATUS_UNKNOWN {
		t.Errorf("Неверный ответ от UseCase.")
		return
	}
}


func TestChangePassword(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB, uc := InitTestUseCase(mockCtrl)

	user := models.User{
		Id:        0,
		Firstname: "test",
		Lastname:  "test",
		Password:  pkg.HashPassword("test"),
		Username:  "test",
	}
	userBytes, _ := json.Marshal(user)

	session := utils_proto.Session {
		Username:      user.Username,
		Authenticated: wrapperspb.Bool(true),
	}

	form := models.ChangePasswordForm{
		OldPassword: "test",
		NewPassword: "test2",
		NewPasswordConf: "test2",
	}

	mockDB.EXPECT().GetUserInfoByUsername(context.Background(), &repository_proto.GetUserInfoByUsernameRequest{
		Username: session.Username,
	}).Return(&repository_proto.ResponseUser{
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
		User: userBytes,
	}, nil)
	mockDB.EXPECT().ChangeUserPassword(context.Background(), &repository_proto.ChangeForm{
		User: userBytes,
		Data: pkg.HashPassword(form.NewPassword),
	}).Return(&utils_proto.DatabaseResponse{
		Status: utils_proto.DatabaseStatus_OK,
	}, nil)

	var response pkg.JsonResponse
	resp, err := uc.ChangePassword(context.Background(), &profile_proto.ChangePasswordRequest{
		Data: &session,
		PasswordOld: form.OldPassword,
		PasswordNew: form.NewPassword,
	})
	json_err := json.Unmarshal(resp.Response, &response)
	if err != nil || json_err != nil || response != pkg.NO_ERR {
		t.Errorf("Неверный ответ от UseCase.")
		return
	}
}