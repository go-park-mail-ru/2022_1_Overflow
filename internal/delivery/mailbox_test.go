package delivery

import (
	"OverflowBackend/internal/models"
	"OverflowBackend/internal/repository/mock"
	"OverflowBackend/pkg"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"testing"
)


func TestSend(t *testing.T) {
	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: jar,
	}

	db := mock.MockDB{}
	db.Create("test")
	createTestUsers(&db)
	d := Delivery{}
	router := InitTestRouter(&db, &d, []string{"/mail/send", "/signin"}, []func(http.ResponseWriter, *http.Request){d.SendMail, d.SignIn})

	srv := httptest.NewServer(router)
	defer srv.Close()

	sendUrl := fmt.Sprintf("%s/mail/send", srv.URL)

	data := models.MailForm{
		Addressee: "test2",
		Theme: "test",
		Text: "test",
		Files: "test",
	}

	dataJson, _ := json.Marshal(data)
	
	_, err := TestPost(client, dataJson, sendUrl, http.StatusUnauthorized)
	if err != nil {
		t.Error(err)
		return
	}

	// ==============================================

	err = SigninUser(client, models.SignInForm{
		Username: "test",
		Password: "test",
	}, srv.URL)

	if err != nil {
		t.Error(err)
		return
	}

	_, err = TestPost(client, dataJson, sendUrl, http.StatusOK)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestIncome (t *testing.T) {
	mailNum := 5

	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: jar,
	}

	db := mock.MockDB{}
	db.Create("test")
	createTestUsers(&db)
	PrepareMails(&db, mailNum)

	d := Delivery{}
	router := InitTestRouter(&db, &d, []string{"/mail/income", "/signin"}, []func(http.ResponseWriter, *http.Request){d.Income, d.SignIn})

	srv := httptest.NewServer(router)
	defer srv.Close()

	url := fmt.Sprintf("%s/mail/income", srv.URL)

	_, err := TestGet(client, url, http.StatusUnauthorized)
	if err != nil {
		t.Error(err)
		return
	}

	err = SigninUser(client, models.SignInForm{
		Username: "test2",
		Password: "test2",
	}, srv.URL)

	if err != nil {
		t.Error(err)
		return
	}

	_, err = TestPost(client, nil, url, http.StatusMethodNotAllowed)
	if err != nil {
		t.Error(err)
		return
	}

	r, err := TestGet(client, url, http.StatusOK)
	if err != nil {
		t.Error(err)
		return
	}

	var mails []models.Mail

	err = json.NewDecoder(r.Body).Decode(&mails)

	if err != nil {
		t.Error(err)
		return
	}

	if len(mails) != mailNum {
		t.Errorf("Количество сообщений не соответствует ожидаемому. Получено: %v, ожидается: %v.", len(mails), mailNum)
		return
	}
}

func TestOutcome (t *testing.T) {
	mailNum := 5

	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: jar,
	}

	db := mock.MockDB{}
	db.Create("test")
	createTestUsers(&db)
	PrepareMails(&db, mailNum)

	d := Delivery{}
	router := InitTestRouter(&db, &d, []string{"/mail/outcome", "/signin"}, []func(http.ResponseWriter, *http.Request){d.Outcome, d.SignIn})

	srv := httptest.NewServer(router)
	defer srv.Close()

	url := fmt.Sprintf("%s/mail/outcome", srv.URL)

	_, err := TestGet(client, url, http.StatusUnauthorized)
	if err != nil {
		t.Error(err)
		return
	}

	err = SigninUser(client, models.SignInForm{
		Username: "test",
		Password: "test",
	}, srv.URL)

	if err != nil {
		t.Error(err)
		return
	}

	_, err = TestPost(client, nil, url, http.StatusMethodNotAllowed)
	if err != nil {
		t.Error(err)
		return
	}

	r, err := TestGet(client, url, http.StatusOK)
	if err != nil {
		t.Error(err)
		return
	}

	var mails []models.Mail

	err = json.NewDecoder(r.Body).Decode(&mails)

	if err != nil {
		t.Error(err)
		return
	}

	if len(mails) != mailNum {
		t.Errorf("Количество сообщений не соответствует ожидаемому. Получено: %v, ожидается: %v.", len(mails), mailNum)
		return
	}
}

func TestRead(t *testing.T) {
	mailNum := 5

	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: jar,
	}

	db := mock.MockDB{}
	db.Create("test")
	createTestUsers(&db)
	PrepareMails(&db, mailNum)

	d := Delivery{}
	router := InitTestRouter(&db, &d, []string{"/mail/read", "/signin"}, []func(http.ResponseWriter, *http.Request){d.ReadMail, d.SignIn})

	srv := httptest.NewServer(router)
	defer srv.Close()

	err := SigninUser(client, models.SignInForm{
		Username: "test2",
		Password: "test2",
	}, srv.URL)

	if err != nil {
		t.Error(err)
		return
	}

	url := fmt.Sprintf("%s/mail/read?id=0", srv.URL)

	r, err := TestPost(client, nil, url, http.StatusOK)
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

	mail, err := db.GetMailInfoById(0)
	if err != nil {
		t.Error(err)
		return
	}

	if mail.Read != true {
		t.Errorf("Письмо не было прочитано.")
		return
	}
}

func TestForward(t *testing.T) {
	mailNum := 5

	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: jar,
	}

	db := mock.MockDB{}
	db.Create("test")
	createTestUsers(&db)
	PrepareMails(&db, mailNum)

	d := Delivery{}
	router := InitTestRouter(&db, &d, []string{"/mail/forward", "/signin"}, []func(http.ResponseWriter, *http.Request){d.ForwardMail, d.SignIn})

	srv := httptest.NewServer(router)
	defer srv.Close()

	err := SigninUser(client, models.SignInForm{
		Username: "test",
		Password: "test",
	}, srv.URL)

	if err != nil {
		t.Error(err)
		return
	}

	url := fmt.Sprintf("%s/mail/forward?mail_id=0&username=test2", srv.URL)

	_, err = TestPost(client, nil, url, http.StatusOK)
	if err != nil {
		t.Error(err)
		return
	}

	user, err := db.GetUserInfoByUsername("test")
	if err != nil {
		t.Error(err)
		return
	}
	mails, err := db.GetOutcomeMails(user.Id)

	if err != nil {
		t.Error(err)
		return
	}

	mailNum++
	if len(mails) != mailNum {
		t.Errorf("Количество сообщений не соответствует ожидаемому. Получено: %v, ожидается: %v.", len(mails), mailNum)
		return
	}
}

func TestDelete(t *testing.T) {
	mailNum := 5

	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: jar,
	}

	db := mock.MockDB{}
	db.Create("test")
	createTestUsers(&db)
	PrepareMails(&db, mailNum)

	d := Delivery{}
	router := InitTestRouter(&db, &d, []string{"/mail/delete", "/signin"}, []func(http.ResponseWriter, *http.Request){d.DeleteMail, d.SignIn})

	srv := httptest.NewServer(router)
	defer srv.Close()

	err := SigninUser(client, models.SignInForm{
		Username: "test",
		Password: "test",
	}, srv.URL)

	if err != nil {
		t.Error(err)
		return
	}

	url := fmt.Sprintf("%s/mail/delete?id=0", srv.URL)

	_, err = TestPost(client, nil, url, http.StatusOK)
	if err != nil {
		t.Error(err)
		return
	}

	jar, _ = cookiejar.New(nil)
	client.Jar = jar

	err = SigninUser(client, models.SignInForm{
		Username: "test2",
		Password: "test2",
	}, srv.URL)

	if err != nil {
		t.Error(err)
		return
	}

	_, err = TestPost(client, nil, url, http.StatusOK)
	if err != nil {
		t.Error(err)
		return
	}

	user, err := db.GetUserInfoByUsername("test")
	if err != nil {
		t.Error(err)
		return
	}
	mails, err := db.GetOutcomeMails(user.Id)

	if err != nil {
		t.Error(err)
		return
	}

	mailNum--
	if len(mails) != mailNum {
		t.Errorf("Количество сообщений не соответствует ожидаемому. Получено: %v, ожидается: %v.", len(mails), mailNum)
		return
	}
}
