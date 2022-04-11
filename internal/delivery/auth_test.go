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

func TestSignin(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUC := mocks.NewMockUseCaseInterface(mockCtrl)

	d := delivery.Delivery{}
	router := InitTestRouter(mockUC, &d, []string{"/signin"}, []func(http.ResponseWriter, *http.Request){d.SignIn})
	d.Init(mockUC, DefConf)

	data := models.SignInForm{
		Username: "test",
		Password: "test",
	}

	mockUC.EXPECT().SignIn(data).Return(pkg.NO_ERR)

	srv := httptest.NewServer(router)
	defer srv.Close()

	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: jar,
	}

	dataJson, _ := json.Marshal(data)
	_, err, token := Get(client, fmt.Sprintf("%s/signin", srv.URL), http.StatusMethodNotAllowed)
	if err != nil {
		t.Error(err)
		return
	}
	_, err = Post(client, dataJson, fmt.Sprintf("%s/signin", srv.URL), http.StatusOK, token, "")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestBadSignin(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUC := mocks.NewMockUseCaseInterface(mockCtrl)

	d := delivery.Delivery{}
	router := InitTestRouter(mockUC, &d, []string{"/signin"}, []func(http.ResponseWriter, *http.Request){d.SignIn})
	d.Init(mockUC, DefConf)

	data := models.SignInForm{
		Username: "test",
		Password: "pass",
	}

	mockUC.EXPECT().SignIn(data).Return(pkg.WRONG_CREDS_ERR)

	srv := httptest.NewServer(router)
	defer srv.Close()

	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: jar,
	}

	dataJson, _ := json.Marshal(data)
	_, err, token := Get(client, fmt.Sprintf("%s/signin", srv.URL), http.StatusMethodNotAllowed)
	if err != nil {
		t.Error(err)
		return
	}
	_, err = Post(client, dataJson, fmt.Sprintf("%s/signin", srv.URL), http.StatusBadRequest, token, "")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestSignup(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUC := mocks.NewMockUseCaseInterface(mockCtrl)

	d := delivery.Delivery{}
	router := InitTestRouter(mockUC, &d, []string{"/signup"}, []func(http.ResponseWriter, *http.Request){d.SignUp})
	d.Init(mockUC, DefConf)

	data := models.SignUpForm{
		LastName:     "John",
		FirstName:    "Doe",
		Username:     "ededededed",
		Password:     "pass",
		PasswordConf: "pass",
	}

	mockUC.EXPECT().SignUp(data).Return(pkg.NO_ERR)

	srv := httptest.NewServer(router)
	defer srv.Close()

	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: jar,
	}

	dataJson, _ := json.Marshal(data)
	_, err, token := Get(client, fmt.Sprintf("%s/signup", srv.URL), http.StatusMethodNotAllowed)
	if err != nil {
		t.Error(err)
		return
	}
	_, err = Post(client, dataJson, fmt.Sprintf("%s/signup", srv.URL), http.StatusOK, token, "")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestBadPassword(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUC := mocks.NewMockUseCaseInterface(mockCtrl)

	d := delivery.Delivery{}
	router := InitTestRouter(mockUC, &d, []string{"/signup"}, []func(http.ResponseWriter, *http.Request){d.SignUp})
	d.Init(mockUC, DefConf)

	data := models.SignUpForm{
		LastName:     "John",
		FirstName:    "Doe",
		Username:     "ededededed",
		Password:     "pass",
		PasswordConf: "passd",
	}

	mockUC.EXPECT().SignUp(data).Return(pkg.CreateJsonErr(pkg.STATUS_BAD_VALIDATION, ""))

	srv := httptest.NewServer(router)
	defer srv.Close()

	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: jar,
	}

	dataJson, _ := json.Marshal(data)

	_, err, token := Get(client, fmt.Sprintf("%s/signup", srv.URL), http.StatusMethodNotAllowed)
	if err != nil {
		t.Error(err)
		return
	}
	_, err = Post(client, dataJson, fmt.Sprintf("%s/signup", srv.URL), http.StatusBadRequest, token, "")

	if err != nil {
		t.Error(err)
		return
	}
}

func TestEmptyForm(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUC := mocks.NewMockUseCaseInterface(mockCtrl)

	d := delivery.Delivery{}
	router := InitTestRouter(mockUC, &d, []string{"/signup"}, []func(http.ResponseWriter, *http.Request){d.SignUp})
	d.Init(mockUC, DefConf)

	data := models.SignUpForm{
		LastName:     "",
		FirstName:    "",
		Username:     "ededededed",
		Password:     "pass",
		PasswordConf: "passd",
	}

	mockUC.EXPECT().SignUp(data).Return(pkg.CreateJsonErr(pkg.STATUS_BAD_VALIDATION, ""))

	srv := httptest.NewServer(router)
	defer srv.Close()

	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: jar,
	}

	dataJson, _ := json.Marshal(data)
	_, err, token := Get(client, fmt.Sprintf("%s/signup", srv.URL), http.StatusMethodNotAllowed)
	if err != nil {
		t.Error(err)
		return
	}
	_, err = Post(client, dataJson, fmt.Sprintf("%s/signup", srv.URL), http.StatusBadRequest, token, "")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestSignout(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUC := mocks.NewMockUseCaseInterface(mockCtrl)

	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: jar,
	}

	d := delivery.Delivery{}
	router := InitTestRouter(mockUC, &d, []string{"/logout", "/signin"}, []func(http.ResponseWriter, *http.Request){d.SignOut, d.SignIn})

	srv := httptest.NewServer(router)
	defer srv.Close()

	url := fmt.Sprintf("%s/logout", srv.URL)
	
	form := models.SignInForm{
		Username: "test",
		Password: "test",
	}

	mockUC.EXPECT().SignIn(form).Return(pkg.NO_ERR)

	err := SigninUser(client, form, srv.URL)
	if err != nil {
		t.Error(err)
		return
	}
	
	_, err, token := Get(client, url, http.StatusMethodNotAllowed)
	if err != nil {
		t.Error(err)
		return
	}
	_, err = Post(client, nil, url, http.StatusOK, token, "")
	if err != nil {
		t.Error(err)
		return
	}
}