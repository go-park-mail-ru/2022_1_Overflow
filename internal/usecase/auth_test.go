package usecase_test

import (
	"OverflowBackend/internal/models"
	"OverflowBackend/pkg"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestSignIn(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB, uc := InitTestUseCase(mockCtrl)
	
	form := models.SignInForm{
		Username: "test",
		Password: "test",
	}

	mockDB.EXPECT().GetUserInfoByUsername("test").Return(models.User{
		Id: 0,
		FirstName: "test",
		LastName: "test",
		Username: "test",
		Password: "test",
	}, nil)

	r := uc.SignIn(form)
	if r != pkg.NO_ERR {
		t.Errorf("Неверный ответ от UseCase.")
		return
	}

	mockDB.EXPECT().GetUserInfoByUsername("test").Return(models.User{}, nil)
	r = uc.SignIn(form)
	if r != pkg.WRONG_CREDS_ERR {
		t.Errorf("Неверный ответ от UseCase.")
		return
	}

	mockDB.EXPECT().GetUserInfoByUsername("test").Return(models.User{
		Id: 0,
		FirstName: "test",
		LastName: "test",
		Username: "test",
		Password: "pass",
	}, nil)
	r = uc.SignIn(form)
	if r != pkg.WRONG_CREDS_ERR {
		t.Errorf("Неверный ответ от UseCase.")
		return
	}
}

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