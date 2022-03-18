package profile

import (
	"OverflowBackend/internal/repository"
	"OverflowBackend/internal/usecase"
	"OverflowBackend/internal/usecase/auth"
)

type Profile struct {
	db repository.DatabaseRepository
	sm usecase.SessionManagerUseCase
}

func (p *Profile) Init(repo repository.DatabaseRepository) {
	p.db = repo
	p.sm = auth.SessionManager{}
}