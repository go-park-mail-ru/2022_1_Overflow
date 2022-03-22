package test

import (
	"OverflowBackend/internal/models"
	//"OverflowBackend/internal/usecase/session"

	"encoding/gob"
	//"net/http"
	//"net/http/httptest"
	//"testing"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

func Init() {
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

/*
func TestSessionManager(t *testing.T) {
	Init()

	r := httptest.NewRequest(http.MethodGet, "/", nil)
	//w := httptest.NewRecorder()

	data, err := session.GetData(r)
	if (err == nil) {
		t.Error(err)
		return
	}
	if (*data != models.Session{}) {
		t.Errorf("Данные сессии не являются пустыми.")
		return
	}
}
*/