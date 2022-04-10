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
	"time"

	"github.com/golang/mock/gomock"
)


func TestSend(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUC := mocks.NewMockUseCaseInterface(mockCtrl)

	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: jar,
	}

	d := delivery.Delivery{}
	router := InitTestRouter(mockUC, &d, []string{"/mail/send", "/signin"}, []func(http.ResponseWriter, *http.Request){d.SendMail, d.SignIn})

	srv := httptest.NewServer(router)
	defer srv.Close()

	sendUrl := fmt.Sprintf("%s/mail/send", srv.URL)

	signinForm := models.SignInForm{
		Username: "test",
		Password: "test",
	}

	data := models.MailForm{
		Addressee: "test2",
		Theme: "test",
		Text: "test",
		Files: "test",
	}

	mockUC.EXPECT().SignIn(signinForm).Return(pkg.NO_ERR)
	mockUC.EXPECT().SendMail(&models.Session{Username: "test", Authenticated: true}, data).Return(pkg.NO_ERR)

	dataJson, _ := json.Marshal(data)
	
	_, err := Post(client, dataJson, sendUrl, http.StatusForbidden, "")
	if err != nil {
		t.Error(err)
		return
	}

	// ==============================================

	err = SigninUser(client, signinForm, srv.URL)

	if err != nil {
		t.Error(err)
		return
	}

	_, err, token := Get(client, sendUrl, http.StatusOK)
	if err != nil {
		t.Error(err)
		return
	}
	_, err = Post(client, dataJson, sendUrl, http.StatusOK, token)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestIncome(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUC := mocks.NewMockUseCaseInterface(mockCtrl)

	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: jar,
	}

	d := delivery.Delivery{}
	router := InitTestRouter(mockUC, &d, []string{"/mail/income", "/signin"}, []func(http.ResponseWriter, *http.Request){d.Income, d.SignIn})

	srv := httptest.NewServer(router)
	defer srv.Close()

	url := fmt.Sprintf("%s/mail/income", srv.URL)

	signinForm := models.SignInForm{
		Username: "test2",
		Password: "test2",
	}

	mails, _ := json.Marshal([]models.MailAdditional{})

	mockUC.EXPECT().SignIn(signinForm).Return(pkg.NO_ERR)
	mockUC.EXPECT().Income(&models.Session{Username:"test2", Authenticated: true}).Return(mails, pkg.NO_ERR)

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

func TestOutcome(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUC := mocks.NewMockUseCaseInterface(mockCtrl)

	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: jar,
	}

	d := delivery.Delivery{}
	router := InitTestRouter(mockUC, &d, []string{"/mail/outcome", "/signin"}, []func(http.ResponseWriter, *http.Request){d.Outcome, d.SignIn})

	srv := httptest.NewServer(router)
	defer srv.Close()

	url := fmt.Sprintf("%s/mail/outcome", srv.URL)

	signinForm := models.SignInForm{
		Username: "test",
		Password: "test",
	}

	mails, _ := json.Marshal([]models.MailAdditional{})

	mockUC.EXPECT().SignIn(signinForm).Return(pkg.NO_ERR)
	mockUC.EXPECT().Outcome(&models.Session{Username:"test", Authenticated: true}).Return(mails, pkg.NO_ERR)

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

func TestRead(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUC := mocks.NewMockUseCaseInterface(mockCtrl)

	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: jar,
	}

	d := delivery.Delivery{}
	router := InitTestRouter(mockUC, &d, []string{"/mail/read", "/signin"}, []func(http.ResponseWriter, *http.Request){d.ReadMail, d.SignIn})

	srv := httptest.NewServer(router)
	defer srv.Close()

	signinForm := models.SignInForm{
		Username: "test2",
		Password: "test2",
	}

	mockUC.EXPECT().SignIn(signinForm).Return(pkg.NO_ERR)
	mockUC.EXPECT().ReadMail(&models.Session{Username: "test2", Authenticated: true}, int32(0)).Return(pkg.NO_ERR)

	err := SigninUser(client, signinForm, srv.URL)

	if err != nil {
		t.Error(err)
		return
	}

	url := fmt.Sprintf("%s/mail/read?id=0", srv.URL)

	_, err, token := Get(client, url, http.StatusOK)
	if err != nil {
		t.Error(err)
		return
	}
	r, err := Post(client, nil, url, http.StatusOK, token)
	if err != nil {
		t.Error(err)
		return
	}

	var resp pkg.JsonResponse

	err = json.NewDecoder(r.Body).Decode(&resp)

	if err != nil {
		t.Error(err)
		return
	}

	if resp.Status != pkg.STATUS_OK {
		t.Errorf("Статус тела ответа не соответствует ожидаемому. Получено: %v, ожидается: %v.", resp.Status, pkg.STATUS_OK)
		return
	}
}

func TestForward(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUC := mocks.NewMockUseCaseInterface(mockCtrl)

	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: jar,
	}

	d := delivery.Delivery{}
	router := InitTestRouter(mockUC, &d, []string{"/mail/forward", "/signin"}, []func(http.ResponseWriter, *http.Request){d.ForwardMail, d.SignIn})

	srv := httptest.NewServer(router)
	defer srv.Close()

	url := fmt.Sprintf("%s/mail/forward?mail_id=0", srv.URL)

	data := models.MailForm {
		Addressee: "test2",
		Theme: "forwarded",
		Text: "hello",
		Files: "files",
	}

	signinForm := models.SignInForm{
		Username: "test",
		Password: "test",
	}

	mockUC.EXPECT().SignIn(signinForm).Return(pkg.NO_ERR)
	mockUC.EXPECT().ForwardMail(&models.Session{Username: "test", Authenticated: true}, data, int32(0)).Return(pkg.NO_ERR)

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

func TestDelete(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUC := mocks.NewMockUseCaseInterface(mockCtrl)

	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: jar,
	}

	d := delivery.Delivery{}
	router := InitTestRouter(mockUC, &d, []string{"/mail/delete", "/signin"}, []func(http.ResponseWriter, *http.Request){d.DeleteMail, d.SignIn})

	srv := httptest.NewServer(router)
	defer srv.Close()

	url := fmt.Sprintf("%s/mail/delete?id=0", srv.URL)

	signinForm :=  models.SignInForm{
		Username: "test",
		Password: "test",
	}

	mockUC.EXPECT().SignIn(signinForm).Return(pkg.NO_ERR)
	mockUC.EXPECT().DeleteMail(&models.Session{Username: "test", Authenticated: true}, int32(0)).Return(pkg.NO_ERR)

	err := SigninUser(client, signinForm, srv.URL)

	if err != nil {
		t.Error(err)
		return
	}

	_, err, token := Get(client, url, http.StatusOK)
	if err != nil {
		t.Error(err)
		return
	}

	_, err = Post(client, nil, url, http.StatusOK, token)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestGetMail(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUC := mocks.NewMockUseCaseInterface(mockCtrl)

	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: jar,
	}

	d := delivery.Delivery{}
	router := InitTestRouter(mockUC, &d, []string{"/mail/get", "/signin"}, []func(http.ResponseWriter, *http.Request){d.GetMail, d.SignIn})

	srv := httptest.NewServer(router)
	defer srv.Close()

	url := fmt.Sprintf("%s/mail/get?id=0", srv.URL)

	signinForm :=  models.SignInForm{
		Username: "test",
		Password: "test",
	}

	mail := models.Mail{
		Id:        0,
		Client_id: 0,
		Sender:    "test",
		Addressee: "test2",
		Theme:     "test",
		Text:      "test",
		Files:     "files",
		Date:      time.Now(),
		Read:      false,
	}

	mailBytes, _ := json.Marshal(mail)

	mockUC.EXPECT().SignIn(signinForm).Return(pkg.NO_ERR)
	mockUC.EXPECT().GetMail(&models.Session{Username: "test", Authenticated: true}, int32(0)).Return(mailBytes, pkg.NO_ERR)

	err := SigninUser(client, signinForm, srv.URL)

	if err != nil {
		t.Error(err)
		return
	}

	r, err, _ := Get(client, url, http.StatusOK)
	if err != nil {
		t.Error(err)
		return
	}

	var resp pkg.JsonResponse

	err = json.NewDecoder(r.Body).Decode(&resp)

	if err != nil {
		t.Error(err)
		return
	}

	if resp.Status != pkg.STATUS_OK {
		t.Errorf("Статус тела ответа не соответствует ожидаемому. Получено: %v, ожидается: %v.", resp.Status, pkg.STATUS_OK)
		return
	}
}
