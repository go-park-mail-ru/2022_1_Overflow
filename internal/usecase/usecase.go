package usecase

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/internal/models"
	"OverflowBackend/internal/repository"
	"OverflowBackend/pkg"
	"net/http"
)

type UseCase struct {
	db     repository.DatabaseRepository
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
	SignIn(data models.SignInForm) pkg.JsonResponse
	SignUp(data models.SignUpForm) pkg.JsonResponse
}

type SessionManagerUseCase interface {
	CreateSession(w http.ResponseWriter, r *http.Request, username string) error
	DeleteSession(w http.ResponseWriter, r *http.Request) error
	IsLoggedIn(r *http.Request) bool
	GetData(r *http.Request) (data *models.Session, err error)
}

type ProfileUseCase interface {
	Init(repository.DatabaseRepository, *config.Config)
	GetInfo(data *models.Session) ([]byte, pkg.JsonResponse)
	GetAvatar(data *models.Session) (string, pkg.JsonResponse)
	SetAvatar(data *models.Session, avatar *models.Avatar) pkg.JsonResponse
	SetInfo(data *models.Session, settings *models.SettingsForm) pkg.JsonResponse
}

type MailBoxUseCase interface {
	Init(repository.DatabaseRepository, *config.Config)
	Income(data *models.Session) ([]byte, pkg.JsonResponse)
	Outcome(data *models.Session) ([]byte, pkg.JsonResponse)
	GetMail(data *models.Session, mail_id int32) ([]byte, pkg.JsonResponse)
	DeleteMail(data *models.Session, id int32) pkg.JsonResponse
	ReadMail(data *models.Session, id int32) pkg.JsonResponse
	SendMail(data *models.Session, form models.MailForm) pkg.JsonResponse
	ForwardMail(data *models.Session, form models.MailForm, mail_id int32) pkg.JsonResponse
	RespondMail(data *models.Session, form models.MailForm, mail_id int32) pkg.JsonResponse
}
