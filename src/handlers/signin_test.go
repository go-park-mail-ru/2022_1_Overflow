package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"general"
	"github.com/gorilla/mux"
)

func TestSignin(t *testing.T) {

	r := mux.NewRouter()
	var handler SigninHandler
	handler.Init(r, nil)

	srv := httptest.NewServer(general.SetupCORS(r))
	defer srv.Close()

	data := url.Values{
		"email":    {"ededededed"},
		"password": {"pass"},
	}
	r, err := http.PostForm(fmt.Sprintf("%s/Signin", srv.URL), data)
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
		t.Errorf("Неверный статус от сервера.")
		return
	}
}

func TestBadSignin(t *testing.T) {
	mux := http.NewServeMux()
	var handler SigninHandler
	handler.Init(mux, nil)

	srv := httptest.NewServer(general.SetupCORS(mux))
	defer srv.Close()

	data := url.Values{
		"email":    {"ededededed"},
		"password": {"pass"},
	}
	r, err := http.PostForm(fmt.Sprintf("%s/Signin", srv.URL), data)
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

	if response["status"].(float64) != 2 {
		t.Errorf("Неверный статус от сервера.")
		return
	}
}
