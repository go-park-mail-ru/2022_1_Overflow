package delivery_test

import (
	"OverflowBackend/internal/delivery"
	"OverflowBackend/internal/models"
	"OverflowBackend/internal/repository/mock"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"testing"
)

func TestSignin(t *testing.T) {
	db := mock.MockDB{}
	db.Create("test")
	err := db.AddUser(models.User{
		Id:        0,
		FirstName: "test",
		LastName:  "test",
		Username:  "test",
		Password:  "test",
	})
	if err != nil {
		t.Error(err)
		return
	}

	d := delivery.Delivery{}
	router := InitTestRouter(&db, &d, []string{"/signin"}, []func(http.ResponseWriter, *http.Request){d.SignIn})

	srv := httptest.NewServer(router)
	defer srv.Close()

	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: jar,
	}

	data := map[string]string{
		"username": "test",
		"password": "test",
	}
	dataJson, _ := json.Marshal(data)
	_, err, token := Get(client, fmt.Sprintf("%s/signin", srv.URL), http.StatusOK)
	if err != nil {
		t.Error(err)
		return
	}
	_, err = Post(client, dataJson, fmt.Sprintf("%s/signin", srv.URL), http.StatusOK, token)
	if err != nil {
		t.Error(err)
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
		Username:  "test",
		Password:  "test",
	})

	d := delivery.Delivery{}
	router := InitTestRouter(&db, &d, []string{"/signin"}, []func(http.ResponseWriter, *http.Request){d.SignIn})

	srv := httptest.NewServer(router)
	defer srv.Close()

	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: jar,
	}

	data := map[string]string{
		"username": "test",
		"password": "pass",
	}
	dataJson, _ := json.Marshal(data)
	_, err, token := Get(client, fmt.Sprintf("%s/signin", srv.URL), http.StatusOK)
	if err != nil {
		t.Error(err)
		return
	}
	_, err = Post(client, dataJson, fmt.Sprintf("%s/signin", srv.URL), http.StatusBadRequest, token)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestSignup(t *testing.T) {
	db := mock.MockDB{}
	db.Create("test")

	d := delivery.Delivery{}
	router := InitTestRouter(&db, &d, []string{"/signup"}, []func(http.ResponseWriter, *http.Request){d.SignUp})

	srv := httptest.NewServer(router)
	defer srv.Close()

	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: jar,
	}

	data := map[string]string{
		"last_name":             "John",
		"first_name":            "Doe",
		"username":              "ededededed",
		"password":              "pass",
		"password_confirmation": "pass",
	}
	dataJson, _ := json.Marshal(data)
	_, err, token := Get(client, fmt.Sprintf("%s/signup", srv.URL), http.StatusOK)
	if err != nil {
		t.Error(err)
		return
	}
	_, err = Post(client, dataJson, fmt.Sprintf("%s/signup", srv.URL), http.StatusOK, token)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestBadPassword(t *testing.T) {
	db := mock.MockDB{}
	db.Create("test")

	d := delivery.Delivery{}
	router := InitTestRouter(&db, &d, []string{"/signup"}, []func(http.ResponseWriter, *http.Request){d.SignUp})

	srv := httptest.NewServer(router)
	defer srv.Close()

	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: jar,
	}

	data := map[string]string{
		"last_name":             "John",
		"first_name":            "Doe",
		"username":              "ededededed",
		"password":              "pass",
		"password_confirmation": "passd",
	}

	dataJson, _ := json.Marshal(data)

	_, err, token := Get(client, fmt.Sprintf("%s/signup", srv.URL), http.StatusOK)
	if err != nil {
		t.Error(err)
		return
	}
	_, err = Post(client, dataJson, fmt.Sprintf("%s/signup", srv.URL), http.StatusBadRequest, token)

	if err != nil {
		t.Error(err)
		return
	}
}

func TestEmptyForm(t *testing.T) {
	db := mock.MockDB{}
	db.Create("test")

	d := delivery.Delivery{}
	router := InitTestRouter(&db, &d, []string{"/signup"}, []func(http.ResponseWriter, *http.Request){d.SignUp})

	srv := httptest.NewServer(router)
	defer srv.Close()

	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: jar,
	}

	data := map[string]string{
		"last_name":             "",
		"first_name":            "",
		"username":              "ededededed",
		"password":              "pass",
		"password_confirmation": "passd",
	}
	dataJson, _ := json.Marshal(data)
	_, err, token := Get(client, fmt.Sprintf("%s/signup", srv.URL), http.StatusOK)
	if err != nil {
		t.Error(err)
		return
	}
	_, err = Post(client, dataJson, fmt.Sprintf("%s/signup", srv.URL), http.StatusBadRequest, token)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestSignout(t *testing.T) {
	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: jar,
	}

	db := mock.MockDB{}
	db.Create("test")
	createTestUsers(&db)
	d := delivery.Delivery{}
	router := InitTestRouter(&db, &d, []string{"/logout", "/signin"}, []func(http.ResponseWriter, *http.Request){d.SignOut, d.SignIn})

	srv := httptest.NewServer(router)
	defer srv.Close()

	url := fmt.Sprintf("%s/logout", srv.URL)
	
	form := models.SignInForm{
		Username: "test",
		Password: "test",
	}

	err := SigninUser(client, form, srv.URL)
	if err != nil {
		t.Error(err)
		return
	}
	
	_, err, token := Get(client, url, http.StatusOK)
	if err != nil {
		t.Error(err)
		return
	}
	_, err = Post(client, nil, url, http.StatusOK, token)
	if err != nil {
		t.Error(err)
		return
	}

	_, err, token = Get(client, url, http.StatusUnauthorized)
	if err != nil {
		t.Error(err)
		return
	}
	_, err = Post(client, nil, url, http.StatusUnauthorized, token)
	if err != nil {
		t.Error(err)
		return
	}
}