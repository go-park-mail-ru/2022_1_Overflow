package usecase

import (
	"OverflowBackend/internal/models"
	"OverflowBackend/internal/repository"
	"net/http"
)

type GeneralUseCase interface {
	AuthUseCase
	SessionManagerUseCase
	ProfileUseCase
	MailBoxUseCase
}

type AuthUseCase interface {
	Init(repository.DatabaseRepository)
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
	Init(repository.DatabaseRepository)
	GetInfo(w http.ResponseWriter, r *http.Request)
}

type MailBoxUseCase interface {
	Init(repository.DatabaseRepository)
	Income(w http.ResponseWriter, r *http.Request)
	Outcome(w http.ResponseWriter, r *http.Request)
	//Send(w http.ResponseWriter, r *http.Request)
}