package delivery

import (
	"OverflowBackend/internal/models"
	"OverflowBackend/internal/repository/mock"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"testing"
)

func TestGetInfo(t *testing.T) {
	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: jar,
	}

	db := mock.MockDB{}
	db.Create("test")
	createTestUsers(&db)

	d := Delivery{}
	router := InitTestRouter(&db, &d, []string{"/profile", "/signin"}, []func(http.ResponseWriter, *http.Request){d.GetInfo, d.SignIn})

	srv := httptest.NewServer(router)
	defer srv.Close()

	url := fmt.Sprintf("%s/profile", srv.URL)

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

	_, err = TestGet(client, url, http.StatusOK)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestSetInfo(t *testing.T) {
	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: jar,
	}

	db := mock.MockDB{}
	db.Create("test")
	createTestUsers(&db)

	d := Delivery{}
	router := InitTestRouter(&db, &d, []string{"/profile/set", "/signin"}, []func(http.ResponseWriter, *http.Request){d.SetInfo, d.SignIn})

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

	url := fmt.Sprintf("%s/profile/set", srv.URL)

	data := map[string]string{
		"last_name": "",
		"first_name": "",
		"password": "changed",
	}

	dataJson, err := json.Marshal(data)
	if err != nil {
		t.Error(err)
		return
	}

	_, err = TestPost(client, dataJson, url, http.StatusOK)
	if err != nil {
		t.Error(err)
		return
	}

	/*

	var dataNew map[string]string

	err = json.NewDecoder(r.Body).Decode(&dataNew)

	if err != nil {
		t.Error(err)
		return
	}

	if dataNew["first_name"] != data["first_name"] || dataNew["last_name"] != data["last_name"] {
		t.Errorf("Несоответствие данных в SetInfo.")
		return
	}
	*/
}
