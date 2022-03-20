package cmd

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/internal/delivery"
	"OverflowBackend/internal/middlewares"
	"OverflowBackend/internal/repository"
	"OverflowBackend/internal/usecase"

	"net/http"

	"github.com/gorilla/mux"

	_ "OverflowBackend/docs"

	"github.com/swaggo/http-swagger"
)

type RouterManager struct {
	d *delivery.Delivery
	uc usecase.UseCaseInterface
}

func (rm *RouterManager) Init(repo repository.DatabaseRepository, config *config.Config) {
	rm.d = &delivery.Delivery{}
	rm.uc = &usecase.UseCase{}
	rm.uc.Init(repo, config)
	rm.d.Init(rm.uc, config)
}

func (rm *RouterManager) NewRouter(swaggerPort string) http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/signin", rm.d.SignIn)
	router.HandleFunc("/logout", rm.d.SignOut)
	router.HandleFunc("/signup", rm.d.SignUp)
	router.HandleFunc("/profile", rm.d.GetInfo)
	router.HandleFunc("/set_profile", rm.d.SetInfo)
	router.HandleFunc("/set_profile/avatar", rm.d.SetAvatar)
	router.HandleFunc("/income", rm.d.Income)
	router.HandleFunc("/outcome", rm.d.Outcome)
	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	))
	router.Use(middlewares.Middleware)
	return config.SetupCORS(router)
}
