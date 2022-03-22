package delivery

import (
	"OverflowBackend/cmd"
	"OverflowBackend/internal/config"
	"OverflowBackend/internal/models"
	"OverflowBackend/internal/repository/mock"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var defConf = config.TestConfig()

func TestSignin(t *testing.T) {

	db := mock.MockDB{}
	db.Create("test")
	db.AddUser(models.User{
		Id:        0,
		FirstName: "test",
		LastName:  "test",
		Email:     "test",
		Password:  "test",
	})

	rm := cmd.RouterManager{}
	rm.Init(&db, defConf)

	srv := httptest.NewServer(rm.NewRouter(defConf.Server.Port))
	defer srv.Close()

	data := map[string]string{
		"email":    "test",
		"password": "test",
	}
	dataJson, _ := json.Marshal(data)
	r, err := http.Post(fmt.Sprintf("%s/signin", srv.URL), "application/json", bytes.NewBuffer(dataJson))
	if err != nil {
		t.Error(err)
		return
	}

	if r.StatusCode != 200 {
		t.Errorf("Неверный статус от сервера: %v. Ожидается: %v.", r.StatusCode, 200)
		return
	}
}

func TestBadSignin(t *testing.T) {
	db := mock.MockDB{}
	db.Create("test")
	db.AddUser(models.User{
		Id:        0,
		FirstName: "test",
		LastName:  "test",
		Email:     "test",
		Password:  "test",
	})

	rm := cmd.RouterManager{}
	rm.Init(&db, defConf)

	srv := httptest.NewServer(rm.NewRouter(defConf.Server.Port))
	defer srv.Close()

	data := map[string]string{
		"email":    "test",
		"password": "pass",
	}
	dataJson, _ := json.Marshal(data)
	r, err := http.Post(fmt.Sprintf("%s/signin", srv.URL), "application/json", bytes.NewBuffer(dataJson))
	if err != nil {
		t.Error(err)
		return
	}

	if r.StatusCode != 500 {
		t.Errorf("Неверный статус от сервера: %v. Ожидается: %v.", r.StatusCode, 500)
		return
	}
}

func TestSignup(t *testing.T) {
	db := mock.MockDB{}
	db.Create("test")

	rm := cmd.RouterManager{}
	rm.Init(&db, defConf)

	srv := httptest.NewServer(rm.NewRouter(defConf.Server.Port))
	defer srv.Close()

	data := map[string]string{
		"last_name":             "John",
		"first_name":            "Doe",
		"email":                 "ededededed",
		"password":              "pass",
		"password_confirmation": "pass",
	}
	dataJson, _ := json.Marshal(data)
	r, err := http.Post(fmt.Sprintf("%s/signup", srv.URL), "application/json", bytes.NewBuffer(dataJson))
	if (err != nil) {
		t.Error(err)
		return
	}

	if r.StatusCode != 200 {
		t.Errorf("Неверный статус от сервера.")
		return
	}
}

func TestBadPassword(t *testing.T) {
	db := mock.MockDB{}
	db.Create("test")

	rm := cmd.RouterManager{}
	rm.Init(&db, defConf)

	srv := httptest.NewServer(rm.NewRouter(defConf.Server.Port))
	defer srv.Close()

	data := map[string]string{
		"last_name":             "John",
		"first_name":            "Doe",
		"email":                 "ededededed",
		"password":              "pass",
		"password_confirmation": "passd",
	}
	dataJson, _ := json.Marshal(data)
	r, err := http.Post(fmt.Sprintf("%s/signup", srv.URL), "application/json", bytes.NewBuffer(dataJson))

	if (err != nil) {
		t.Error(err)
		return
	}

	if r.StatusCode != 500 {
		t.Errorf("Неверный статус от сервера: %v. Ожидается: %v.", r.StatusCode, 500)
		return
	}
}

func TestEmptyForm(t *testing.T) {
	db := mock.MockDB{}
	db.Create("test")

	rm := cmd.RouterManager{}
	rm.Init(&db, defConf)

	srv := httptest.NewServer(rm.NewRouter(defConf.Server.Port))
	defer srv.Close()

	data := map[string]string{
		"last_name":             "",
		"first_name":            "",
		"email":                 "ededededed",
		"password":              "pass",
		"password_confirmation": "passd",
	}
	dataJson, _ := json.Marshal(data)
	r, err := http.Post(fmt.Sprintf("%s/signup", srv.URL), "application/json", bytes.NewBuffer(dataJson))
	if (err != nil) {
		t.Error(err)
		return
	}

	if r.StatusCode != 500 {
		t.Errorf("Неверный статус от сервера: %v. Ожидается: %v.", r.StatusCode, 500)
		return
	}
}
