package cmd

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/internal/delivery"
	"OverflowBackend/internal/delivery/middlewares"
	"OverflowBackend/internal/repository"
	"OverflowBackend/internal/usecase"

	"net/http"

	"github.com/gorilla/mux"
)

type RouterManager struct {
	d *delivery.Delivery
	uc usecase.UseCaseInterface
}

func (rm *RouterManager) Init(repo repository.DatabaseRepository) {
	rm.d = &delivery.Delivery{}
	rm.uc = &usecase.UseCase{}
	rm.uc.Init(repo)
	rm.d.Init(rm.uc)
}

func (rm *RouterManager) NewRouter() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/signin", middlewares.Middleware(rm.d.SignIn, false))
	router.HandleFunc("/logout", middlewares.Middleware(rm.d.SignOut, true))
	router.HandleFunc("/signup", middlewares.Middleware(rm.d.SignUp, false))
	router.HandleFunc("/profile", middlewares.Middleware(rm.d.GetInfo, true))
	router.HandleFunc("/income", middlewares.Middleware(rm.d.Income, true))
	router.HandleFunc("/outcome", middlewares.Middleware(rm.d.Outcome, true))
	InitSwagger(router)
	return config.SetupCORS(router)
}
