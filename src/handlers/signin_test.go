package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

var handler SigninHandler

func TestSignin(t *testing.T) {
	srv := httptest.NewServer(handler.Handlers())
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
	srv := httptest.NewServer(handler.Handlers())
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
