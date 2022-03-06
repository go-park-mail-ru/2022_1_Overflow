package main

import (
	"fmt"
	"testing"
	"net/url"
	"net/http"
	"net/http/httptest"
	"encoding/json"
)

func TestSignup(t *testing.T) {
	srv := httptest.NewServer(handlers())
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

func TestBadSignup(t *testing.T) {
	srv := httptest.NewServer(handlers())
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
		t.Errorf("Неверный статус от сервера.")
		return
	}
}