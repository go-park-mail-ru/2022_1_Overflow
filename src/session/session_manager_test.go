package session

import (
	"encoding/gob"
	"net/http"
	"net/http/httptest"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"testing"
)

func Init() {
	authKeyOne := securecookie.GenerateRandomKey(64)
	encryptionKeyOne := securecookie.GenerateRandomKey(32)

	store = sessions.NewCookieStore(
		authKeyOne,
		encryptionKeyOne,
	)

	store.Options = &sessions.Options{
		//Path: "/",
		MaxAge:   60 * 15, // 15 мин
		HttpOnly: false,
		Secure:   false,
	}

	gob.Register(Session{})
}

func TestSessionManager(t *testing.T) {
	Init()

	r := httptest.NewRequest(http.MethodGet, "/", nil)
	//w := httptest.NewRecorder()

	data, err := GetData(r)
	if (err == nil) {
		t.Error(err)
		return
	}
	if (data != Session{}) {
		t.Errorf("Данные сессии не являются пустыми.")
		return
	}
}