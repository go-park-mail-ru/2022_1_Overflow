package handlers

import (
	"general"
	"fmt"
	"testing"
	"net/url"
	"net/http"
	"net/http/httptest"
	"encoding/json"
	"github.com/gorilla/mux"
)

func TestSignup(t *testing.T) {
	r := mux.NewRouter()
	var handler SigninHandler
	handler.Init(r, nil)

	srv := httptest.NewServer(general.SetupCORS(r))
	defer srv.Close()

	data := url.Values{
		"last_name":             {"John"},
		"first_name":            {"Doe"},
		"email":                 {"ededededed"},
		"password":              {"pass"},
		"password_confirmation": {"pass"},
	}
	r, err := http.PostForm(fmt.Sprintf("%s/signup", srv.URL), data)
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
	
	if response["status"].(float64) != 0 {
		t.Errorf("Неверный статус от сервера.")
		return
	}
}

func TestBadPassword(t *testing.T) {
	mux := http.NewServeMux()
	var handler SigninHandler
	handler.Init(mux, nil)

	srv := httptest.NewServer(general.SetupCORS(mux))
	defer srv.Close()

	data := url.Values{
		"last_name":             {"John"},
		"first_name":            {"Doe"},
		"email":                 {"ededededed"},
		"password":              {"pass"},
		"password_confirmation": {"passd"},
	}
	r, err := http.PostForm(fmt.Sprintf("%s/signup", srv.URL), data)
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
	
	if response["status"].(float64) != 2 {
		t.Errorf("Неверный статус от сервера: %v. Ожидается: %v.", response["status"].(float64), 2)
		return
	}
}

func TestEmptyForm(t *testing.T) {
	mux := http.NewServeMux()
	var handler SigninHandler
	handler.Init(mux, nil)

	srv := httptest.NewServer(general.SetupCORS(mux))
	defer srv.Close()

	data := url.Values{
		"last_name":             {""},
		"first_name":            {""},
		"email":                 {"ededededed"},
		"password":              {"pass"},
		"password_confirmation": {"pass"},
	}
	r, err := http.PostForm(fmt.Sprintf("%s/signup", srv.URL), data)
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
	
	if response["status"].(float64) != 2 {
		t.Errorf("Неверный статус от сервера: %v. Ожидается: %v.", response["status"].(float64), 2)
		return
	}
}