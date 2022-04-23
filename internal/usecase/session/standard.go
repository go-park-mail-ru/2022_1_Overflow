package session

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/internal/models"
	"encoding/gob"
	"net/http"

	"github.com/gorilla/sessions"
)

var session_name string = "OveflowMail"

type StandardManager struct {
	// Данные всех сессий.
	store *sessions.CookieStore
}

func (s *StandardManager) Init(config *config.Config) (err error) {
	authKeyOne := config.Server.Keys.AuthKey
	encryptionKeyOne := config.Server.Keys.EncKey

	s.store = sessions.NewCookieStore(
		[]byte(authKeyOne),
		[]byte(encryptionKeyOne),
	)

	s.store.Options = &sessions.Options{
		MaxAge:   0,
		HttpOnly: false,
		Secure:   false,
	}

	gob.Register(models.Session{})
	return
}

func (s *StandardManager) CreateSession(w http.ResponseWriter, r *http.Request, username string) error {
	session, err := s.store.Get(r, session_name)
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

func (s *StandardManager) DeleteSession(w http.ResponseWriter, r *http.Request) error {
	session, err := s.store.Get(r, session_name)
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

func (s *StandardManager) IsLoggedIn(r *http.Request) bool {
	session, err := s.store.Get(r, session_name)
	if err != nil {
		return false
	}
	return !session.IsNew
}

func (s *StandardManager) GetData(r *http.Request) (data *models.Session, err error) {
	defer func() {
		errRecover := recover()
		if errRecover != nil {
			data, err = nil, errRecover.(error)
		}
	}()
	session, err := s.store.Get(r, session_name)
	if err != nil {
		return nil, err
	}
	sessionData := session.Values["data"].(models.Session)
	return &sessionData, nil
}
