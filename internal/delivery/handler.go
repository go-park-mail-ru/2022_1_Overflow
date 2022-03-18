package delivery

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/internal/repository"
	"OverflowBackend/internal/usecase"
	"OverflowBackend/internal/usecase/auth"
	"OverflowBackend/internal/usecase/mailbox"
	"OverflowBackend/internal/usecase/profile"

	"net/http"

	"github.com/gorilla/mux"
)

type RouterManager struct {
	auth usecase.AuthUseCase
	profile usecase.ProfileUseCase
	mailbox usecase.MailBoxUseCase
	handlers map[string]func(http.ResponseWriter, *http.Request) 
}

func (d *RouterManager) Init(repo repository.DatabaseRepository) {
	d.auth = &auth.Auth{}
	d.profile = &profile.Profile{}
	d.mailbox = &mailbox.MailBox{}
	d.auth.Init(repo)
	d.profile.Init(repo)
	d.mailbox.Init(repo)
	d.handlers = map[string]func(http.ResponseWriter, *http.Request) {
		"/signin": d.auth.SignIn,
		"/logout": d.auth.SignOut,
		"/signup": d.auth.SignUp,
		"/profile": d.profile.GetInfo,
		"/income": d.mailbox.Income,
		"/outcome": d.mailbox.Outcome,
	}
}

func (d *RouterManager) NewRouter() http.Handler {
	router := mux.NewRouter()
	for k,v := range(d.handlers) {
		router.HandleFunc(k, v)
	}
	return config.SetupCORS(router)
}
