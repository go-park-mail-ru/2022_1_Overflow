package test

import (
	response "OverflowBackend/src/response"
	handlers "OverflowBackend/src/handlers"

	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestSignin(t *testing.T) {

	router := mux.NewRouter()
	var handler handlers.SigninHandler
	handler.Init(router, nil)

	srv := httptest.NewServer(response.SetupCORS(router))
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
	router := mux.NewRouter()
	var handler handlers.SigninHandler
	handler.Init(router, nil)

	srv := httptest.NewServer(response.SetupCORS(router))
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
