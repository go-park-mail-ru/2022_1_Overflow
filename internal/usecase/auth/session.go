package auth

import (
	"OverflowBackend/internal/models"
	"encoding/gob"
	"net/http"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

var session_name string = "OveflowMail"

// Данные всех сессий.
var store *sessions.CookieStore

// Запускается при инициализации пакета
func init() {
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

	gob.Register(models.Session{})
}

func (sm SessionManager) CreateSession(w http.ResponseWriter, r *http.Request, email string) error {
	session, err := store.Get(r, session_name)
	if err != nil {
		return err
	}
	data := &models.Session{
		Email:         email,
		Authenticated: true,
	}
	session.Values["data"] = data
	err = session.Save(r, w)
	return err
}

func (sm SessionManager) DeleteSession(w http.ResponseWriter, r *http.Request) error {
	session, err := store.Get(r, session_name)
	if err != nil {
		return err
	}

	session.Values["data"] = models.Session{}
	session.Options.MaxAge = -1

	err = session.Save(r, w)
	if err != nil {
		return err
	}

	return nil
}

func (sm SessionManager) IsLoggedIn(r *http.Request) bool {
	session, err := store.Get(r, session_name)
	if err != nil {
		return false
	}
	return !session.IsNew
}

func (sm SessionManager) GetData(r *http.Request) (data *models.Session, err error) {
	defer func() {
		errRecover := recover()
		if errRecover != nil {
			data, err = nil, errRecover.(error)
		}
	}()
	session, err := store.Get(r, session_name)
	if (err != nil) {
		return nil, err
	}
	sessionData := session.Values["data"].(models.Session)
	return &sessionData, nil
}
