package imap_server

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/internal/models"
	"OverflowBackend/pkg"
	"OverflowBackend/proto/auth_proto"
	"OverflowBackend/proto/folder_manager_proto"
	"OverflowBackend/proto/mailbox_proto"
	"OverflowBackend/proto/profile_proto"
	"OverflowBackend/proto/utils_proto"
	"context"
	"encoding/json"
	"errors"

	imap "github.com/emersion/go-imap"
	imap_backend "github.com/emersion/go-imap/backend"
	"github.com/emersion/go-imap/server"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type IMAPServer struct {
	auth auth_proto.AuthClient
	profile profile_proto.ProfileClient
	mailbox mailbox_proto.MailboxClient
	folderManager folder_manager_proto.FolderManagerClient
	config *config.Config
}

func (s *IMAPServer) Init(
	config *config.Config,
	auth auth_proto.AuthClient,
	profile profile_proto.ProfileClient,
	mailbox mailbox_proto.MailboxClient,
	folderManager folder_manager_proto.FolderManagerClient,
	) {
	s.config = config
	s.auth = auth
	s.profile = profile
	s.mailbox = mailbox
	s.folderManager = folderManager
}

func (s *IMAPServer) Login(connInfo *imap.ConnInfo, username, password string) (imap_backend.User, error) {
	log.Debug("Попытка аутентификации: ", *connInfo)
	form := models.SignInForm{
		Username: username,
		Password: password,
	}
	formBytes, _ := json.Marshal(form)
	resp, err := s.auth.SignIn(context.Background(), &auth_proto.SignInRequest{
		Form: formBytes,
	})
	if err != nil {
		log.Error("Ошибка входа: ", err)
		return nil, err
	}
	var response pkg.JsonResponse
	err = json.Unmarshal(resp.Response, &response)
	if err != nil {
		log.Error("Ошибка входа: ", err)
		return nil, err
	}
	if response != pkg.NO_ERR {
		err = errors.New(response.Message)
		log.Error("Ошибка входа: ", err)
		return nil, err
	}
	user := &User{}
	user.Init(s.config, s.auth, s.profile, s.mailbox, s.folderManager, &utils_proto.Session{
		Username: username,
		Authenticated: wrapperspb.Bool(true),
	})
	log.Debug("Успешный вход.")
	return user, nil
}

func main() {
	be := &IMAPServer{}
	s := server.New(be)
	s.Addr = ":1143"
	// Since we will use this server for testing only, we can allow plain text
	// authentication over unencrypted connections
	s.AllowInsecureAuth = true

	log.Println("Starting IMAP server at localhost:1143")
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}