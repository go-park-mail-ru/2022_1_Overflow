package session

import (
	"net/http"
	"encoding/gob"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

var session_name string = "OveflowMail"

// Структура сессии.
type Session struct {
	Email         string
	Authenticated bool
}

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

	gob.Register(Session{})
}

func CreateSession(w http.ResponseWriter, r *http.Request, email string) error {
	session, err := store.Get(r, session_name)
	if err != nil {
		return err
	}
	data := &Session{
		Email:         email,
		Authenticated: true,
	}
	session.Values["data"] = data
	err = session.Save(r, w)
	return err
}

func DeleteSession(w http.ResponseWriter, r *http.Request) error {
	session, err := store.Get(r, session_name)
	if err != nil {
		return err
	}

	session.Values["data"] = Session{}
	session.Options.MaxAge = -1

	err = session.Save(r, w)
	if err != nil {
		return err
	}

	return nil
}

func IsLoggedIn(r *http.Request) bool {
	session, err := store.Get(r, session_name)
	if err != nil {
		return false
	}
	return !session.IsNew
}

func GetData(r *http.Request) (Session, error) {
	session, err := store.Get(r, session_name)
	if (err != nil) {
		return Session{}, err
	}
	return session.Values["data"].(Session), nil
}
