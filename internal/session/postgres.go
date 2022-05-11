package session

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/proto/utils_proto"
	"encoding/gob"
	"fmt"
	"net/http"

	"github.com/antonlindstrom/pgstore"
	"github.com/gorilla/sessions"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/wrapperspb"
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
		Authenticated: wrapperspb.Bool(true),
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
			data, err = nil, errRecover.(error)
		}
	}()
	session, err := pm.store.Get(r, session_name)
	if err != nil {
		return nil, err
	}
	sessionData := session.Values["data"].(*utils_proto.Session)
	return sessionData, nil
}
