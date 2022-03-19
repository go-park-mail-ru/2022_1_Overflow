package test

import (
	"OverflowBackend/internal/delivery"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSignup(t *testing.T) {
	rm := cmd.RouterManager{}
	rm.Init()

	srv := httptest.NewServer(rm.NewRouter())
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
	rm := cmd.RouterManager{}
	rm.Init()

	srv := httptest.NewServer(rm.NewRouter())
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

	var response map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&response)
	if (err != nil) {
		t.Error(err)
		return
	}
	
	if response["status"].(float64) != 1 {
		t.Errorf("Неверный статус от сервера: %v. Ожидается: %v.", response["status"].(float64), 1)
		return
	}
}

func TestEmptyForm(t *testing.T) {
	rm := cmd.RouterManager{}
	rm.Init()

	srv := httptest.NewServer(rm.NewRouter())
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

	var response map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&response)
	if (err != nil) {
		t.Error(err)
		return
	}
	
	if response["status"].(float64) != 1 {
		t.Errorf("Неверный статус от сервера: %v. Ожидается: %v.", response["status"].(float64), 1)
		return
	}
}