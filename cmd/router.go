package cmd

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/internal/delivery"
	"OverflowBackend/internal/middlewares"

	"net/http"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"

	"github.com/swaggo/http-swagger"
)

type RouterManager struct {
	d *delivery.Delivery
	config *config.Config
}

func (rm *RouterManager) Init(
	config *config.Config,
	authDial grpc.ClientConnInterface,
	profileDial grpc.ClientConnInterface,
	mailboxDial grpc.ClientConnInterface,
	folderManagerDial grpc.ClientConnInterface,
	) {
	rm.d = &delivery.Delivery{}
	rm.d.Init(config, authDial, profileDial, mailboxDial, folderManagerDial)
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
	// ======================================================================
	router.HandleFunc("/profile", rm.d.GetInfo)
	router.HandleFunc("/profile/avatar", rm.d.GetAvatar)
	router.HandleFunc("/profile/set", rm.d.SetInfo)
	router.HandleFunc("/profile/avatar/set", rm.d.SetAvatar)
	router.HandleFunc("/profile/change_password", rm.d.ChangePassword)
	// ======================================================================
	router.HandleFunc("/mail/income", rm.d.Income)
	router.HandleFunc("/mail/outcome", rm.d.Outcome)
	router.HandleFunc("/mail/get", rm.d.GetMail)
	router.HandleFunc("/mail/delete", rm.d.DeleteMail)
	router.HandleFunc("/mail/read", rm.d.ReadMail)
	router.HandleFunc("/mail/send", rm.d.SendMail)
	// ======================================================================
	router.HandleFunc("/folder/add", rm.d.AddFolder)
	router.HandleFunc("/folder/mail/add", rm.d.AddMailToFolder)
	router.HandleFunc("/folder/mail/delete", rm.d.DeleteFolderMail)
	router.HandleFunc("/folder/rename", rm.d.ChangeFolder)
	router.HandleFunc("/folder/delete", rm.d.DeleteFolder)
	router.HandleFunc("/folder/list", rm.d.ListFolders)
	// ======================================================================
	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	))
	router.Use(middlewares.Middleware)
	return config.SetupCORS(router)
}
