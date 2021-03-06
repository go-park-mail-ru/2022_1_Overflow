package session

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/proto/utils_proto"
	"net/http"
)

var Manager SessionManager

type SessionManager interface {
	Init(config *config.Config) (err error)
	CreateSession(w http.ResponseWriter, r *http.Request, username string) error
	DeleteSession(w http.ResponseWriter, r *http.Request) error
	IsLoggedIn(r *http.Request) bool
	GetData(r *http.Request) (data *utils_proto.Session, err error)
	GetDataFull(r *http.Request, storeName string, field string) (interface{}, error)
	SetDataFull(w http.ResponseWriter, r *http.Request, storeName, field string, value interface{}) (err error)
}

func Init(config *config.Config) error {
	if config.Server.SessionType == "postgres" {
		Manager = &PostgresManager{}
	} else {
		Manager = &StandardManager{}
	}
	return Manager.Init(config)
}
