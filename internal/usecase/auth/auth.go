package auth

import (
	"OverflowBackend/internal/repository"
	"OverflowBackend/internal/usecase"
)

type Auth struct {
	db repository.DatabaseRepository
	sm usecase.SessionManagerUseCase
}

type SessionManager struct {}

func (a *Auth) Init(repo repository.DatabaseRepository) {
	a.db = repo
	a.sm = SessionManager{}
}