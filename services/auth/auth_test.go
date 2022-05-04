package auth_test

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/internal/models"
	"OverflowBackend/pkg"
	"OverflowBackend/proto/auth_proto"
	"OverflowBackend/proto/repository_proto"
	"OverflowBackend/proto/utils_proto"
	"OverflowBackend/services/auth"
	"context"
	"encoding/json"
	"testing"

	"github.com/golang/mock/gomock"
	log "github.com/sirupsen/logrus"
)

func InitTestUseCase(ctrl *gomock.Controller) (*repository_proto.MockDatabaseRepositoryClient, *auth.AuthService) {
	log.SetLevel(log.FatalLevel)
	db := repository_proto.NewMockDatabaseRepositoryClient(ctrl)
	uc := auth.AuthService{}
	uc.Init(config.TestConfig(), db)
	return db, &uc
}

func TestSignIn(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB, uc := InitTestUseCase(mockCtrl)
	
	form := models.SignInForm{
		Username: "test",
		Password: "test",
	}
	formBytes, _ := json.Marshal(form)

	user := models.User{
		Id: 0,
		Firstname: "test",
		Lastname: "test",
		Username: form.Username,
		Password: pkg.HashPassword(form.Password),
	}
	userBytes, _ := json.Marshal(user)

	var response pkg.JsonResponse

	mockDB.EXPECT().GetUserInfoByUsername(context.Background(), &repository_proto.GetUserInfoByUsernameRequest{
		Username: form.Username,
	}).Return(&repository_proto.ResponseUser{
		User: userBytes,
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
	}, nil)

	resp, err := uc.SignIn(context.Background(), &auth_proto.SignInRequest{
		Form: formBytes,
	})
	json_err := json.Unmarshal(resp.Response, &response)
	if err != nil || json_err != nil || response != pkg.NO_ERR {
		t.Errorf("Неверный ответ от UseCase.")
		return
	}

	userBytes, _ = json.Marshal(models.User{})
	mockDB.EXPECT().GetUserInfoByUsername(context.Background(), &repository_proto.GetUserInfoByUsernameRequest{
		Username: form.Username,
	}).Return(&repository_proto.ResponseUser{
		User: userBytes,
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
	}, nil)
	resp, _ = uc.SignIn(context.Background(), &auth_proto.SignInRequest{
		Form: formBytes,
	})
	json_err = json.Unmarshal(resp.Response, &response)
	if json_err != nil || response != pkg.WRONG_CREDS_ERR {
		t.Errorf("Неверный ответ от UseCase.")
		return
	}

	user.Password = "123"
	userBytes, _ = json.Marshal(user)
	mockDB.EXPECT().GetUserInfoByUsername(context.Background(), &repository_proto.GetUserInfoByUsernameRequest{
		Username: form.Username,
	}).Return(&repository_proto.ResponseUser{
		User: userBytes,
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
	}, nil)
	resp, _ = uc.SignIn(context.Background(), &auth_proto.SignInRequest{
		Form: formBytes,
	})
	json_err = json.Unmarshal(resp.Response, &response)
	if json_err != nil || response != pkg.WRONG_CREDS_ERR {
		t.Errorf("Неверный ответ от UseCase.")
		return
	}
}

/*
func TestSignUp(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB, uc := InitTestUseCase(mockCtrl)

	form := models.SignUpForm{
		FirstName: "test",
		LastName: "test",
		Username: "test",
		Password: "test",
		PasswordConf: "test",
	}

	mockDB.EXPECT().GetUserInfoByUsername(form.Username).Return(models.User{}, nil)
	mockDB.EXPECT().AddUser(models.User{
		Id: 0,
		FirstName: form.FirstName,
		LastName: form.LastName,
		Username: form.Username,
		Password: form.Password,
	}).Return(nil)

	r := uc.SignUp(form)
	if r != pkg.NO_ERR {
		t.Errorf("Неверный ответ от UseCase.")
		return
	}
}
*/