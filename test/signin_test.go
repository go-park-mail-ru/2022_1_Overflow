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

func TestSignin(t *testing.T) {

	rm := delivery.RouterManager{}
	rm.Init()

	srv := httptest.NewServer(rm.NewRouter())
	defer srv.Close()

	data := map[string]string{
		"email":    "ededededed",
		"password": "pass",
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
	rm := delivery.RouterManager{}
	rm.Init()

	srv := httptest.NewServer(rm.NewRouter())
	defer srv.Close()

	data := map[string]string{
		"email":    "ededededed",
		"password": "",
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

	if response["status"].(float64) != 1 {
		t.Errorf("Неверный статус от сервера: %v. Ожидается: %v.", response["status"].(float64), 1)
		return
	}
}
