package delivery_test

import (
	"OverflowBackend/internal/delivery"
	"OverflowBackend/internal/models"
	"OverflowBackend/mocks"
	"OverflowBackend/pkg"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGetInfo(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUC := mocks.NewMockUseCaseInterface(mockCtrl)
	
	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: jar,
	}

	d := delivery.Delivery{}
	router := InitTestRouter(mockUC, &d, []string{"/profile", "/signin"}, []func(http.ResponseWriter, *http.Request){d.GetInfo, d.SignIn})

	srv := httptest.NewServer(router)
	defer srv.Close()

	url := fmt.Sprintf("%s/profile", srv.URL)

	signinForm := models.SignInForm{
		Username: "test",
		Password: "test",
	}

	info, _ := json.Marshal(models.User{
		Id: 0,
		FirstName: "test",
		LastName: "test",
		Password: "test",
		Username: "test",
	})

	mockUC.EXPECT().SignIn(signinForm).Return(pkg.NO_ERR)
	mockUC.EXPECT().GetInfo(&models.Session{Username: "test", Authenticated: true}).Return(info, pkg.NO_ERR)

	_, err, _ := Get(client, url, http.StatusUnauthorized)
	if err != nil {
		t.Error(err)
		return
	}

	err = SigninUser(client, signinForm, srv.URL)

	if err != nil {
		t.Error(err)
		return
	}

	_, err = Post(client, nil, url, http.StatusForbidden, "")
	if err != nil {
		t.Error(err)
		return
	}

	_, err, _ = Get(client, url, http.StatusOK)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestSetInfo(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUC := mocks.NewMockUseCaseInterface(mockCtrl)

	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: jar,
	}

	d := delivery.Delivery{}
	router := InitTestRouter(mockUC, &d, []string{"/profile/set", "/signin"}, []func(http.ResponseWriter, *http.Request){d.SetInfo, d.SignIn})

	srv := httptest.NewServer(router)
	defer srv.Close()

	signinForm := models.SignInForm{
		Username: "test",
		Password: "test",
	}

	url := fmt.Sprintf("%s/profile/set", srv.URL)

	data := models.SettingsForm{
		FirstName: "changed",
		LastName: "changed",
		Password: "changed",
	}

	mockUC.EXPECT().SignIn(signinForm).Return(pkg.NO_ERR)
	mockUC.EXPECT().SetInfo(&models.Session{Username: "test", Authenticated: true}, &data).Return(pkg.NO_ERR)

	err := SigninUser(client, signinForm, srv.URL)

	if err != nil {
		t.Error(err)
		return
	}

	dataJson, err := json.Marshal(data)
	if err != nil {
		t.Error(err)
		return
	}

	_, err, token := Get(client, url, http.StatusOK)
	if err != nil {
		t.Error(err)
		return
	}
	_, err = Post(client, dataJson, url, http.StatusOK, token)
	if err != nil {
		t.Error(err)
		return
	}
}
