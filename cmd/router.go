package cmd

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/internal/delivery"
	"OverflowBackend/internal/repository"
	"OverflowBackend/internal/usecase"

	"net/http"

	"github.com/gorilla/mux"
)

type RouterManager struct {
	d *delivery.Delivery
	uc usecase.UseCaseInterface
	handlers map[string]func(http.ResponseWriter, *http.Request) 
}

func (rm *RouterManager) Init(repo repository.DatabaseRepository) {
	rm.d = &delivery.Delivery{}
	rm.uc = &usecase.UseCase{}
	rm.uc.Init(repo)
	rm.d.Init(rm.uc)
	rm.handlers = map[string]func(http.ResponseWriter, *http.Request) {
		"/signin": rm.d.SignIn,
		"/logout": rm.d.SignOut,
		"/signup": rm.d.SignUp,
		"/profile": rm.d.GetInfo,
		"/income": rm.d.Income,
		"/outcome": rm.d.Outcome,
	}
}

func (d *RouterManager) NewRouter() http.Handler {
	router := mux.NewRouter()
	for k,v := range(d.handlers) {
		router.HandleFunc(k, v)
	}
	return config.SetupCORS(router)
}
