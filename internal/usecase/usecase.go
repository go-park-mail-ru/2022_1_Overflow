package usecase

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/internal/models"
	"OverflowBackend/internal/repository"
	"net/http"
)

type UseCase struct {
	db repository.DatabaseRepository
	config *config.Config
}

func (uc *UseCase) Init(repo repository.DatabaseRepository, config *config.Config) {
	uc.db = repo
	uc.config = config
}

type UseCaseInterface interface {
	AuthUseCase
	ProfileUseCase
	MailBoxUseCase
}

type AuthUseCase interface {
	Init(repository.DatabaseRepository, *config.Config)
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
	Init(repository.DatabaseRepository, *config.Config)
	GetInfo(data *models.Session) (userJson []byte, err error)
	GetAvatar(data *models.Session) (avatarUrl string, err error)
	SetAvatar(data *models.Session, avatar *models.Avatar) error
	SetInfo(data *models.Session, settings *models.SettingsForm) error
}

type MailBoxUseCase interface {
	Init(repository.DatabaseRepository, *config.Config)
	Income(data *models.Session) (parsed []byte, err error)
	Outcome(data *models.Session) (parsed []byte, err error)
}