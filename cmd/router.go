package cmd

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/internal/delivery"
	"OverflowBackend/internal/middlewares"
	"OverflowBackend/internal/repository"
	"OverflowBackend/internal/usecase"
	"fmt"

	"net/http"

	"github.com/gorilla/mux"

	_ "OverflowBackend/docs"

	"github.com/swaggo/http-swagger"
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

func (rm *RouterManager) NewRouter(swaggerPort string) http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/signin", middlewares.Middleware(rm.d.SignIn, false))
	router.HandleFunc("/logout", middlewares.Middleware(rm.d.SignOut, true))
	router.HandleFunc("/signup", middlewares.Middleware(rm.d.SignUp, false))
	router.HandleFunc("/profile", middlewares.Middleware(rm.d.GetInfo, true))
	router.HandleFunc("/income", middlewares.Middleware(rm.d.Income, true))
	router.HandleFunc("/outcome", middlewares.Middleware(rm.d.Outcome, true))
	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("http://0.0.0.0:%v/swagger/doc.json", swaggerPort)), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	))
	return config.SetupCORS(router)
}
