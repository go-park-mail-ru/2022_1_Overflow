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
	config *config.Config
}

func (rm *RouterManager) Init(repo repository.DatabaseRepository, config *config.Config) {
	rm.d = &delivery.Delivery{}
	rm.uc = &usecase.UseCase{}
	rm.uc.Init(repo, config)
	rm.d.Init(rm.uc, config)
	rm.config = config
}

func (rm *RouterManager) NewRouter(swaggerPort string) http.Handler {
	router := mux.NewRouter()
	fs := http.FileServer(http.Dir(rm.config.Server.Static.Dir))
	router.PathPrefix(rm.config.Server.Static.Handle).Handler(
		http.StripPrefix(rm.config.Server.Static.Handle, fs,
	))
	router.HandleFunc("/signin", rm.d.SignIn)
	router.HandleFunc("/logout", rm.d.SignOut)
	router.HandleFunc("/signup", rm.d.SignUp)
	router.HandleFunc("/profile", rm.d.GetInfo)
	router.HandleFunc("/profile/avatar", rm.d.GetAvatar)
	router.HandleFunc("/profile/set", rm.d.SetInfo)
	router.HandleFunc("/profile/avatar/set", rm.d.SetAvatar)
	router.HandleFunc("/mail/income", rm.d.Income)
	router.HandleFunc("/mail/outcome", rm.d.Outcome)
	router.HandleFunc("/mail/delete", rm.d.DeleteMail)
	router.HandleFunc("/mail/read", rm.d.ReadMail)
	router.HandleFunc("/mail/send", rm.d.SendMail)
	router.HandleFunc("/get_token", rm.d.CSRFHandler)
	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	))
	router.Use(middlewares.Middleware)
	return config.SetupCORS(router)
}
