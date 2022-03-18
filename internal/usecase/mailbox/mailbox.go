package mailbox

import (
	"OverflowBackend/internal/repository"
	"OverflowBackend/internal/usecase"
	"OverflowBackend/internal/usecase/auth"
)

type MailBox struct {
	db repository.DatabaseRepository
	sm usecase.SessionManagerUseCase
}

func (mb *MailBox) Init(repo repository.DatabaseRepository) {
	mb.db = repo
	mb.sm = auth.SessionManager{}
}