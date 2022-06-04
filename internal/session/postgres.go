package session

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/proto/utils_proto"
	"encoding/gob"
	"errors"
	"fmt"
	"net/http"

	"github.com/antonlindstrom/pgstore"
	"github.com/gorilla/sessions"
	log "github.com/sirupsen/logrus"
)

type PostgresManager struct {
	store *pgstore.PGStore
}

func (pm *PostgresManager) Init(config *config.Config) (err error) {
	authKeyOne := config.Server.Keys.AuthKey
	encryptionKeyOne := config.Server.Keys.EncKey

	dbUrl := fmt.Sprintf("postgres://%v:%v@%v:%v/%v",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Name,
	)
	pm.store, err = pgstore.NewPGStore(
		dbUrl,
		[]byte(authKeyOne),
		[]byte(encryptionKeyOne),
	)
	pm.store.Options = &sessions.Options{
		MaxAge:   10*365*24*60*60,
		HttpOnly: false,
		Secure:   false,
	}

	gob.Register(&utils_proto.Session{})
	return
}

func (pm *PostgresManager) CreateSession(w http.ResponseWriter, r *http.Request, username string) error {
	session, _ := pm.store.Get(r, session_name)
	data := &utils_proto.Session{
		Username:      username,
		Authenticated: true,
	}
	session.Values["data"] = data
	err := session.Save(r, w)
	return err
}

func (pm *PostgresManager) DeleteSession(w http.ResponseWriter, r *http.Request) error {
	session, _ := pm.store.Get(r, session_name)
	session.Values["data"] = &utils_proto.Session{}
	session.Options.MaxAge = -1

	err := session.Save(r, w)
	if err != nil {
		return err
	}

	/*
	session, err = pm.store.Get(r, AddStoreName)
	if err != nil {
		return nil
	}
	session.Options.MaxAge = -1
	session.Save(r, w)
	*/
	
	return nil
}

func (pm *PostgresManager) IsLoggedIn(r *http.Request) bool {
	session, err := pm.store.Get(r, session_name)
	if err != nil {
		return false
	}
	return !session.IsNew
}

func (pm *PostgresManager) GetData(r *http.Request) (data *utils_proto.Session, err error) {
	defer func() {
		errRecover := recover()
		if errRecover != nil {
			log.Error(errRecover)
			data, err = nil, errRecover.(error)
		}
	}()
	val, err := pm.GetDataFull(r, session_name, "data")
	if err != nil {
		return nil, err
	}
	sessionData := val.(*utils_proto.Session)
	return sessionData, nil
}

func (pm *PostgresManager) GetDataFull(r *http.Request, storeName string, field string) (interface{}, error) {
	session, err := pm.store.Get(r, storeName)
	if err != nil {
		return nil, err
	}
	if val, ok := session.Values[field]; ok {
		return val, nil
	}
	return nil, errors.New("Поле не существует.")
}

func (pm *PostgresManager) SetDataFull(w http.ResponseWriter, r *http.Request, storeName, field string, value interface{}) (err error) {
	defer func() {
		errRecover := recover()
		if errRecover != nil {
			log.Error(errRecover)
			err = errRecover.(error)
		}
	}()
	session, _ := pm.store.Get(r, storeName)
	session.Values[field] = value
	err = session.Save(r, w)
	return
}
