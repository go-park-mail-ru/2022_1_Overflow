package session

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/proto/utils_proto"
	"encoding/gob"
	"net/http"

	"github.com/gorilla/sessions"
	"google.golang.org/protobuf/types/known/wrapperspb"

	log "github.com/sirupsen/logrus"
)

var session_name string = "OverflowMail"

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
		MaxAge:   10*365*24*60*60,
		HttpOnly: false,
		Secure:   false,
	}

	gob.Register(&utils_proto.Session{})
	return
}

func (s *StandardManager) CreateSession(w http.ResponseWriter, r *http.Request, username string) error {
	session, err := s.store.Get(r, session_name)
	if err != nil {
		log.Debug("Ошибка декодирования сессии.")
		err := s.DeleteSession(w, r)
		if err != nil {
			log.Debug("Ошибка удаления сессии.")
			return err
		}
		session, err = s.store.Get(r, session_name)
		if err != nil {
			log.Debug("Повторная ошибка декодирования сессии.")
			return err
		}
	}
	data := &utils_proto.Session{
		Username:      username,
		Authenticated: wrapperspb.Bool(true),
	}
	session.Values["data"] = data
	err = session.Save(r, w)
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
	session, err := s.store.Get(r, session_name)
	if err != nil {
		return nil, err
	}
	sessionData := session.Values["data"].(*utils_proto.Session)
	return sessionData, nil
}
