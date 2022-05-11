package smtp_server

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
	"io"
	"io/ioutil"
	"net/mail"

	"github.com/emersion/go-smtp"
	//log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/wrapperspb"
)


type SMTPServer struct{
	auth auth_proto.AuthClient
	profile profile_proto.ProfileClient
	mailbox mailbox_proto.MailboxClient
	folderManager folder_manager_proto.FolderManagerClient
	config *config.Config
}

func (s *SMTPServer) Init(
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

func (s *SMTPServer) AnonymousLogin(state *smtp.ConnectionState) (smtp.Session, error) {
	return nil, smtp.ErrAuthUnsupported
}

func (s *SMTPServer) Login(state *smtp.ConnectionState, username string, password string)  (smtp.Session, error) {
	form := models.SignInForm{
		Username: username,
		Password: password,
	}
	formBytes, _ := json.Marshal(form)
	resp, err := s.auth.SignIn(context.Background(), &auth_proto.SignInRequest{
		Form: formBytes,
	})
	if err != nil {
		return nil, err
	}
	var response pkg.JsonResponse
	err = json.Unmarshal(resp.Response, &response)
	if err != nil {
		return nil, err
	}
	if response != pkg.NO_ERR {
		return nil, errors.New(response.Message)
	}
	sess := &Session{}
	sess.Init(s.config, s.auth, s.profile, s.mailbox, s.folderManager, &utils_proto.Session{
		Username: username,
		Authenticated: wrapperspb.Bool(true),
	})
	return sess, nil
}

// A Session is returned after EHLO.
type Session struct{
	mailForm models.MailForm

	auth auth_proto.AuthClient
	profile profile_proto.ProfileClient
	mailbox mailbox_proto.MailboxClient
	folderManager folder_manager_proto.FolderManagerClient
	config *config.Config

	userSession *utils_proto.Session
}

func (s *Session) Init(
	config *config.Config,
	auth auth_proto.AuthClient,
	profile profile_proto.ProfileClient,
	mailbox mailbox_proto.MailboxClient,
	folderManager folder_manager_proto.FolderManagerClient,
	userSession *utils_proto.Session,
	) {
	s.config = config
	s.auth = auth
	s.profile = profile
	s.mailbox = mailbox
	s.folderManager = folderManager
	s.userSession = userSession
}

func (s *Session) Mail(from string, opts smtp.MailOptions) error {
	if opts.Auth == nil {
		return smtp.ErrAuthRequired
	}
	if from != s.userSession.Username {
		return errors.New("имя отправителя отличается от имени пользователя")
	}
	return nil
}

func (s *Session) Rcpt(to string) error {
	s.mailForm.Addressee = to
	return nil
}

func (s *Session) Data(r io.Reader) error {
	msg, err := mail.ReadMessage(r)
	if err != nil {
		return err
	}
	s.mailForm.Theme = msg.Header.Get("Subject")
	body, err := ioutil.ReadAll(msg.Body)
	if err != nil {
		return err
	}
	// body has \r\n on the end of the message
	if len(body) > 0 {
		body = body[:len(body)-2]
	}
	s.mailForm.Text = string(body)
	mailFormBytes, _ := json.Marshal(s.mailForm)
	s.mailbox.SendMail(context.Background(), &mailbox_proto.SendMailRequest{
		Data: s.userSession,
		Form: mailFormBytes,
	})
	return nil
}

func (s *Session) Reset() {
	s.mailForm = models.MailForm{}
}

func (s *Session) Logout() error {
	return nil
}