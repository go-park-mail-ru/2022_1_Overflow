package usecase

import (
	"OverflowBackend/internal/models"
	"OverflowBackend/internal/repository"
	"net/http"
)

type UseCase struct {
	db repository.DatabaseRepository
}

func (uc *UseCase) Init(repo repository.DatabaseRepository) {
	uc.db = repo
}

type UseCaseInterface interface {
	AuthUseCase
	ProfileUseCase
	MailBoxUseCase
}

type AuthUseCase interface {
	Init(repository.DatabaseRepository)
	SignIn(data models.SignInForm) error
	SignUp(data models.SignUpForm) error
}

type SessionManagerUseCase interface {
	CreateSession(w http.ResponseWriter, r *http.Request, email string) error
	DeleteSession(w http.ResponseWriter, r *http.Request) error
	IsLoggedIn(r *http.Request) bool
	GetData(r *http.Request) (data *models.Session, err error)
}

type ProfileUseCase interface {
	Init(repository.DatabaseRepository)
	GetInfo(data *models.Session) (userJson []byte, err error)
}

type MailBoxUseCase interface {
	Init(repository.DatabaseRepository)
	Income(data *models.Session) (parsed []byte, err error)
	Outcome(data *models.Session) (parsed []byte, err error)
}