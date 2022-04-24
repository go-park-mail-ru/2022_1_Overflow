package profile_test

import (
	"OverflowBackend/internal/models"
	"OverflowBackend/pkg"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"testing"

	"bou.ke/monkey"
	"github.com/golang/mock/gomock"
)

func TestGetInfo(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB, uc := InitTestUseCase(mockCtrl)

	session := models.Session{
		Username:      "test",
		Authenticated: true,
	}

	user := models.User{
		Id:        0,
		FirstName: "test",
		LastName:  "test",
		Password:  "test",
		Username:  "test",
	}

	mockDB.EXPECT().GetUserInfoByUsername(user.Username).Return(user, nil)

	resp, r := uc.GetInfo(&session)
	if r != pkg.NO_ERR {
		t.Errorf("Неверный ответ от UseCase.")
		return
	}

	userUC := models.User{}

	err := json.Unmarshal(resp, &userUC)
	if err != nil {
		t.Error(err)
		return
	}

	if userUC != user {
		t.Errorf("Информация о пользователе не соответствует заданной. Получено: %v, ожидается: %v.", userUC, user)
		return
	}
}

func TestSetInfo(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB, uc := InitTestUseCase(mockCtrl)

	session := models.Session{
		Username:      "test",
		Authenticated: true,
	}

	user := models.User{
		Id:        0,
		FirstName: "test",
		LastName:  "test",
		Password:  "test",
		Username:  "test",
	}

	settings := models.SettingsForm{
		FirstName: "test2",
		LastName:  "test2",
		Password:  "test2",
	}

	mockDB.EXPECT().GetUserInfoByUsername(user.Username).Return(user, nil)
	mockDB.EXPECT().ChangeUserFirstName(user, settings.FirstName).Return(nil)
	mockDB.EXPECT().ChangeUserLastName(user, settings.LastName).Return(nil)
	mockDB.EXPECT().ChangeUserPassword(user, settings.Password).Return(nil)

	r := uc.SetInfo(&session, &settings)
	if r != pkg.NO_ERR {
		t.Errorf("Неверный ответ от UseCase.")
		return
	}
}

func TestSetAvatar(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	_, uc := InitTestUseCase(mockCtrl)

	session := models.Session{
		Username:      "test",
		Authenticated: true,
	}

	avatar := models.Avatar{
		Name:      "avatar",
		UserEmail: session.Username,
		Content:   []byte{10, 10, 10, 10},
	}

	r := uc.SetAvatar(&session, &avatar)
	if r != pkg.NO_ERR {
		t.Errorf("Неверный ответ от UseCase.")
		return
	}

	monkey.Patch(os.MkdirAll, func(path string, perm fs.FileMode) error { return fmt.Errorf("Ошибка.") })

	r = uc.SetAvatar(&session, &avatar)
	if r.Status != pkg.STATUS_UNKNOWN {
		t.Errorf("Неверный возвращенный статус JSON ответа.")
		return
	}
}
