package cmd

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/internal/delivery"
	"OverflowBackend/internal/middlewares"
	"OverflowBackend/proto/attach_proto"
	"OverflowBackend/proto/auth_proto"
	"OverflowBackend/proto/folder_manager_proto"
	"OverflowBackend/proto/mailbox_proto"
	"OverflowBackend/proto/profile_proto"

	"net/http"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"

	"github.com/swaggo/http-swagger"
)

type RouterManager struct {
	d      *delivery.Delivery
	config *config.Config
}

func (rm *RouterManager) Init(
	config *config.Config,
	authDial grpc.ClientConnInterface,
	profileDial grpc.ClientConnInterface,
	mailboxDial grpc.ClientConnInterface,
	folderManagerDial grpc.ClientConnInterface,
	attachDial grpc.ClientConnInterface,
) {
	rm.d = &delivery.Delivery{}
	rm.d.Init(config, auth_proto.NewAuthClient(authDial), profile_proto.NewProfileClient(profileDial), mailbox_proto.NewMailboxClient(mailboxDial), folder_manager_proto.NewFolderManagerClient(folderManagerDial), attach_proto.NewAttachClient(attachDial))
	rm.config = config
}

func (rm *RouterManager) NewRouter(swaggerPort string) http.Handler {
	router := mux.NewRouter()
	fs := http.FileServer(http.Dir(rm.config.Server.Static.Dir))
	router.PathPrefix(rm.config.Server.Static.Handle).Handler(
		http.StripPrefix(rm.config.Server.Static.Handle, fs))

	routerAPI := router.PathPrefix("/api/v1").Subrouter()

	routerAPI.HandleFunc("/signin", rm.d.SignIn)
	routerAPI.HandleFunc("/logout", rm.d.SignOut)
	routerAPI.HandleFunc("/signup", rm.d.SignUp)
	// ======================================================================
	routerAPI.HandleFunc("/profile", rm.d.GetInfo)
	routerAPI.HandleFunc("/profile/avatar", rm.d.GetAvatar)
	routerAPI.HandleFunc("/profile/set", rm.d.SetInfo)
	routerAPI.HandleFunc("/profile/avatar/set", rm.d.SetAvatar)
	routerAPI.HandleFunc("/profile/change_password", rm.d.ChangePassword)
	// ======================================================================
	routerAPI.HandleFunc("/mail/income", rm.d.Income)
	routerAPI.HandleFunc("/mail/outcome", rm.d.Outcome)
	routerAPI.HandleFunc("/mail/get", rm.d.GetMail)
	routerAPI.HandleFunc("/mail/delete", rm.d.DeleteMail)
	routerAPI.HandleFunc("/mail/read", rm.d.ReadMail)
	routerAPI.HandleFunc("/mail/send", rm.d.SendMail)
	routerAPI.HandleFunc("/mail/attach/add", rm.d.UploadAttach)
	routerAPI.HandleFunc("/mail/attach/get", rm.d.GetAttach)
	routerAPI.HandleFunc("/mail/attach/list", rm.d.ListAttach)
	routerAPI.HandleFunc("/mail/countunread", rm.d.GetCountUnread)
	// ======================================================================
	routerAPI.HandleFunc("/folder/add", rm.d.AddFolder)
	routerAPI.HandleFunc("/folder/mail/add", rm.d.AddMailToFolderById)
	routerAPI.HandleFunc("/folder/mail/add_form", rm.d.AddMailToFolderByObject)
	routerAPI.HandleFunc("/folder/mail/move", rm.d.MoveFolderMail)
	routerAPI.HandleFunc("/folder/mail/delete", rm.d.DeleteFolderMail)
	routerAPI.HandleFunc("/folder/rename", rm.d.ChangeFolder)
	routerAPI.HandleFunc("/folder/delete", rm.d.DeleteFolder)
	routerAPI.HandleFunc("/folder/list", rm.d.ListFolders)
	// ======================================================================
	routerAPI.HandleFunc("/ws", rm.d.WSConnect)
	// ======================================================================
	routerAPI.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("/api/v1/swagger/doc.json"), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	))
	router.Use(middlewares.Middleware)
	return config.SetupCORS(router)
}
