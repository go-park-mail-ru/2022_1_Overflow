package usecase

import (
	"OverflowBackend/internal/models"
	"net/http"
)

type GeneralUseCase interface {
	AuthUseCase
	SessionManagerUseCase
	ProfileUseCase
	MailBoxUseCase
}

type AuthUseCase interface {
	SignOut(w http.ResponseWriter, r *http.Request)
	SignIn(w http.ResponseWriter, r *http.Request)
	SignUp(w http.ResponseWriter, r *http.Request)
}

type SessionManagerUseCase interface {
	CreateSession(w http.ResponseWriter, r *http.Request, email string) error
	DeleteSession(w http.ResponseWriter, r *http.Request) error
	IsLoggedIn(r *http.Request) bool
	GetData(r *http.Request) (data *models.Session, err error)
}

type ProfileUseCase interface {
	Info(w http.ResponseWriter, r *http.Request)
}

type MailBoxUseCase interface {
	Income(w http.ResponseWriter, r *http.Request)
	Outcome(w http.ResponseWriter, r *http.Request)
	Send(w http.ResponseWriter, r *http.Request)
}