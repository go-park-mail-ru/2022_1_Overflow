package auth

import (
	"OverflowBackend/internal/repository"
	"OverflowBackend/internal/usecase"

	"github.com/gorilla/mux"
)

type Auth struct {
	db repository.DatabaseRepository
	sm usecase.SessionManagerUseCase
}

type SessionManager struct {}

func (a *Auth) Init(router *mux.Router, repo repository.DatabaseRepository) {
	a.db = repo
	a.sm = SessionManager{}
}