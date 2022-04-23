package session

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/internal/models"
	"fmt"
	"net/http"

	"github.com/antonlindstrom/pgstore"
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
	return
}

func (pm *PostgresManager) CreateSession(w http.ResponseWriter, r *http.Request, username string) error {
	session, err := pm.store.Get(r, session_name)
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

func (pm *PostgresManager) DeleteSession(w http.ResponseWriter, r *http.Request) error {
	session, err := pm.store.Get(r, session_name)
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

func (pm *PostgresManager) IsLoggedIn(r *http.Request) bool {
	session, err := pm.store.Get(r, session_name)
	if err != nil {
		return false
	}
	return !session.IsNew
}

func (pm *PostgresManager) GetData(r *http.Request) (data *models.Session, err error) {
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
	sessionData := session.Values["data"].(models.Session)
	return &sessionData, nil
}
