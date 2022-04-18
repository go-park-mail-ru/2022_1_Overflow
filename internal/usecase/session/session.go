package session

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/internal/models"
	"encoding/gob"
	"net/http"

	"github.com/gorilla/sessions"
)

var session_name string = "OveflowMail"

// Данные всех сессий.
var store *sessions.CookieStore

func Init(config *config.Config) {
	authKeyOne := config.Server.Keys.AuthKey
	encryptionKeyOne := config.Server.Keys.EncKey

	store = sessions.NewCookieStore(
		[]byte(authKeyOne),
		[]byte(encryptionKeyOne),
	)

	store.Options = &sessions.Options{
		MaxAge:   0,
		HttpOnly: false,
		Secure:   false,
	}

	gob.Register(models.Session{})
}

func CreateSession(w http.ResponseWriter, r *http.Request, username string) error {
	session, err := store.Get(r, session_name)
	if err != nil {
		return err
	}
	data := &models.Session{
		Username:      username,
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

	session.Values["data"] = models.Session{}
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

func GetData(r *http.Request) (data *models.Session, err error) {
	defer func() {
		errRecover := recover()
		if errRecover != nil {
			data, err = nil, errRecover.(error)
		}
	}()
	session, err := store.Get(r, session_name)
	if err != nil {
		return nil, err
	}
	sessionData := session.Values["data"].(models.Session)
	return &sessionData, nil
}
