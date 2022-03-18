package delivery

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/internal/usecase"
	"OverflowBackend/internal/usecase/auth"

	"net/http"

	"github.com/gorilla/mux"
)

type RouterManager struct {
	auth usecase.AuthUseCase
	profile usecase.ProfileUseCase
	mailbox usecase.MailBoxUseCase
	handlers map[string]func(http.ResponseWriter, *http.Request) 
}

func (d *RouterManager) Init() {
	d.auth = auth.Auth{}
	d.handlers = map[string]func(http.ResponseWriter, *http.Request) {
		"/login": d.auth.SignIn,
		"/logout": d.auth.SignOut,
		"/signup": d.auth.SignUp,
		"/profile": d.profile.Info,
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
