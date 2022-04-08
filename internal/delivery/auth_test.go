package delivery_test

import (
	"OverflowBackend/internal/delivery"
	"OverflowBackend/internal/models"
	"OverflowBackend/internal/repository/mock"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"strconv"
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

	data := map[string]string{
		"username": "test",
		"password": "test",
	}
	dataJson, _ := json.Marshal(data)
	r, err := http.Post(fmt.Sprintf("%s/signin", srv.URL), "application/json", bytes.NewBuffer(dataJson))
	if err != nil {
		t.Error(err)
		return
	}

	if r.StatusCode != http.StatusOK {
		t.Errorf("Неверный статус от сервера: %v. Ожидается: %v.", r.StatusCode, http.StatusOK)
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

	data := map[string]string{
		"username": "test",
		"password": "pass",
	}
	dataJson, _ := json.Marshal(data)
	r, err := http.Post(fmt.Sprintf("%s/signin", srv.URL), "application/json", bytes.NewBuffer(dataJson))
	if err != nil {
		t.Error(err)
		return
	}

	if r.StatusCode != http.StatusBadRequest {
		t.Errorf("Неверный статус от сервера: %v. Ожидается: %v.", r.StatusCode, http.StatusBadRequest)
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

	data := map[string]string{
		"last_name":             "John",
		"first_name":            "Doe",
		"username":              "ededededed",
		"password":              "pass",
		"password_confirmation": "pass",
	}
	dataJson, _ := json.Marshal(data)
	r, err := http.Post(fmt.Sprintf("%s/signup", srv.URL), "application/json", bytes.NewBuffer(dataJson))
	if err != nil {
		t.Error(err)
		return
	}

	if r.StatusCode != http.StatusOK {
		t.Errorf("Неверный статус от сервера: %v. Ожидается: %v.", r.StatusCode, http.StatusOK)
		return
	}
}

func TestMultiAuth(t *testing.T) {
	numUsers := 50

	db := mock.MockDB{}
	db.Create("test")

	d := delivery.Delivery{}
	router := InitTestRouter(&db, &d, []string{"/signin", "/signup"}, []func(http.ResponseWriter, *http.Request){d.SignIn, d.SignUp})

	srv := httptest.NewServer(router)
	defer srv.Close()

	for i := 0; i < numUsers; i++ {
		user := strconv.Itoa(i)
		signup := func() {
			data := map[string]string{
				"last_name":             "John",
				"first_name":            "Doe",
				"username":              user,
				"password":              "pass",
				"password_confirmation": "pass",
			}
			dataJson, _ := json.Marshal(data)
			log.Println("Регистрирую пользователя", user)
			r, err := http.Post(fmt.Sprintf("%s/signup", srv.URL), "application/json", bytes.NewBuffer(dataJson))
			if err != nil {
				t.Error(err)
				return
			}
			defer r.Body.Close()

			if r.StatusCode != http.StatusOK {
				t.Errorf("Неверный статус от сервера: %v. Ожидается: %v.", r.StatusCode, http.StatusOK)
				return
			}
		}
		signup()

		signin := func() {
			form2 := models.SignInForm{
				Username: user,
				Password: "pass",
			}

			formJson, _ := json.Marshal(form2)
			log.Println("Вхожу под логином", user)
			r, err := http.Post(fmt.Sprintf("%s/signin", srv.URL), "application/json", bytes.NewBuffer(formJson))
			if err != nil {
				t.Error(err)
				return
			}
			defer r.Body.Close()

			if r.StatusCode != http.StatusOK {
				t.Errorf("Неверный статус от сервера: %v. Ожидается: %v.", r.StatusCode, http.StatusOK)
				return
			}
		}
		signin()
	}
}

func TestBadPassword(t *testing.T) {
	db := mock.MockDB{}
	db.Create("test")

	d := delivery.Delivery{}
	router := InitTestRouter(&db, &d, []string{"/signup"}, []func(http.ResponseWriter, *http.Request){d.SignUp})

	srv := httptest.NewServer(router)
	defer srv.Close()

	data := map[string]string{
		"last_name":             "John",
		"first_name":            "Doe",
		"username":              "ededededed",
		"password":              "pass",
		"password_confirmation": "passd",
	}
	dataJson, _ := json.Marshal(data)
	r, err := http.Post(fmt.Sprintf("%s/signup", srv.URL), "application/json", bytes.NewBuffer(dataJson))

	if err != nil {
		t.Error(err)
		return
	}

	if r.StatusCode != http.StatusBadRequest {
		t.Errorf("Неверный статус от сервера: %v. Ожидается: %v.", r.StatusCode, http.StatusBadRequest)
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

	data := map[string]string{
		"last_name":             "",
		"first_name":            "",
		"username":              "ededededed",
		"password":              "pass",
		"password_confirmation": "passd",
	}
	dataJson, _ := json.Marshal(data)
	r, err := http.Post(fmt.Sprintf("%s/signup", srv.URL), "application/json", bytes.NewBuffer(dataJson))
	if err != nil {
		t.Error(err)
		return
	}

	if r.StatusCode != http.StatusBadRequest {
		t.Errorf("Неверный статус от сервера: %v. Ожидается: %v.", r.StatusCode, http.StatusBadRequest)
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
	
	_, err = TestGet(client, url, http.StatusOK)
	if err != nil {
		t.Error(err)
		return
	}

	_, err = TestGet(client, url, http.StatusUnauthorized)
	if err != nil {
		t.Error(err)
		return
	}
}