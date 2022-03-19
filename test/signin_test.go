package test

import (
	"OverflowBackend/cmd"
	"OverflowBackend/internal/models"
	"OverflowBackend/internal/repository/mock"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSignin(t *testing.T) {

	db := mock.MockDB{}
	db.Create("test")
	db.AddUser(models.User{
		Id: 0,
		FirstName: "test",
		LastName: "test",
		Email: "test",
		Password: "test",
	})

	rm := cmd.RouterManager{}
	rm.Init(&db)

	srv := httptest.NewServer(rm.NewRouter())
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
		t.Errorf("Неверный статус от сервера.")
		return
	}
}

func TestBadSignin(t *testing.T) {
	db := mock.MockDB{}
	db.Create("test")
	db.AddUser(models.User{
		Id: 0,
		FirstName: "test",
		LastName: "test",
		Email: "test",
		Password: "test",
	})

	rm := cmd.RouterManager{}
	rm.Init(&db)

	srv := httptest.NewServer(rm.NewRouter())
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

	var response map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		t.Error(err)
		return
	}

	if response["status"].(float64) != 0 {
		t.Errorf("Неверный статус от сервера: %v. Ожидается: %v.", response["status"].(float64), 1)
		return
	}
}
