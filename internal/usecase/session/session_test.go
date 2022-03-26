package session

import (
	"OverflowBackend/internal/models"

	"encoding/gob"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

func InitTestSession() {
	authKeyOne := securecookie.GenerateRandomKey(64)
	encryptionKeyOne := securecookie.GenerateRandomKey(32)

	store := sessions.NewCookieStore(
		authKeyOne,
		encryptionKeyOne,
	)

	store.Options = &sessions.Options{
		//Path: "/",
		MaxAge:   60 * 15, // 15 мин
		HttpOnly: false,
		Secure:   false,
	}

	gob.Register(models.Session{})
}

func TestGetData(t *testing.T) {
	InitTestSession()

	r := httptest.NewRequest(http.MethodGet, "/", nil)
	//w := httptest.NewRecorder()

	data, err := GetData(r)
	if (err == nil) {
		t.Errorf("Отсутствие ошибки получения данных сессии.")
		return
	}
	if (data != nil) {
		t.Errorf("Данные сессии не являются пустыми.")
		return
	}
}

/*

func TestCreateSession(t *testing.T) {
	InitTestSession()

	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	err := CreateSession(w, r, "test")
	if err != nil {
		t.Error(err)
		return
	}

	r = httptest.NewRequest(http.MethodGet, "/", nil)

	data, err := GetData(r)
	if (err != nil) {
		t.Error(err)
		return
	}
	if (data == nil) {
		t.Errorf("Данные сессии являются пустыми.")
		return
	}
}
*/