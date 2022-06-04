package session

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/proto/utils_proto"
	"encoding/gob"
	"errors"
	"net/http"

	"github.com/gorilla/sessions"

	log "github.com/sirupsen/logrus"
)

const session_name string = "OverMail"
const AddStoreName string = "OverMailAdd"

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
		MaxAge:   10 * 365 * 24 * 60 * 60,
		Path:     "/",
		HttpOnly: false,
		Secure:   false,
	}

	gob.Register(&utils_proto.Session{})
	return
}

func (s *StandardManager) CreateSession(w http.ResponseWriter, r *http.Request, username string) error {
	session, _ := s.store.Get(r, session_name)
	data := &utils_proto.Session{
		Username:      username,
		Authenticated: true,
	}
	session.Values["data"] = data
	err := session.Save(r, w)
	return err
}

func (s *StandardManager) DeleteSession(w http.ResponseWriter, r *http.Request) error {
	session, _ := s.store.Get(r, session_name)
	session.Values["data"] = &utils_proto.Session{}
	session.Options.MaxAge = -1

	err := session.Save(r, w)
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

func (s *StandardManager) GetData(r *http.Request) (data *utils_proto.Session, err error) {
	defer func() {
		errRecover := recover()
		if errRecover != nil {
			log.Error(errRecover)
			data, err = nil, errRecover.(error)
		}
	}()
	val, err := s.GetDataFull(r, session_name, "data")
	if err != nil {
		return nil, err
	}
	sessionData := val.(*utils_proto.Session)
	return sessionData, nil
}

func (s *StandardManager) GetDataFull(r *http.Request, storeName string, field string) (interface{}, error) {
	session, err := s.store.Get(r, storeName)
	if err != nil {
		return nil, err
	}
	if val, ok := session.Values[field]; ok {
		return val, nil
	}
	return nil, errors.New("Поле не существует.")
}

func (s *StandardManager) SetDataFull(w http.ResponseWriter, r *http.Request, storeName, field string, value interface{}) (err error) {
	defer func() {
		errRecover := recover()
		if errRecover != nil {
			log.Error(errRecover)
			err = errRecover.(error)
		}
	}()
	session, _ := s.store.Get(r, storeName)
	session.Values[field] = value
	err = session.Save(r, w)
	return
}
