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

	"github.com/emersion/go-smtp"
	log "github.com/sirupsen/logrus"
	enmime "github.com/jhillyerd/enmime"
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
	log.Debug("Попытка анонимного входа: ", *state)
	sess := &Session{}
	sess.Init(s.config, s.auth, s.profile, s.mailbox, s.folderManager, &utils_proto.Session{
		Username: "anonymous",
		Authenticated: true,
	})
	return sess, nil
}

func (s *SMTPServer) Login(state *smtp.ConnectionState, username string, password string)  (smtp.Session, error) {
	log.Debug("Попытка аутентификации: ", *state)
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
	sess := &Session{}
	sess.Init(s.config, s.auth, s.profile, s.mailbox, s.folderManager, &utils_proto.Session{
		Username: username,
		Authenticated: true,
	})
	log.Debug("Успешный вход.")
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
	log.Debug("MAIL from ", from)
	s.userSession.Username = from
	/*
	if from != s.userSession.Username {
		return errors.New("имя отправителя отличается от имени пользователя")
	}
	*/
	return nil
}

func (s *Session) Rcpt(to string) error {
	log.Debug("RCPT to ", to)
	s.mailForm.Addressee = to
	return nil
}

func (s *Session) Data(r io.Reader) error {
	log.Debug("DATA")
	msg, err := enmime.ReadEnvelope(r)
	if err != nil {
		return err
	}
	s.mailForm.Theme = msg.GetHeader("Subject")
	s.mailForm.Text = msg.Text
	mailFormBytes, _ := json.Marshal(s.mailForm)
	resp, err := s.mailbox.SendMail(context.Background(), &mailbox_proto.SendMailRequest{
		Data: s.userSession,
		Form: mailFormBytes,
	})
	if err != nil {
		log.Error("Ошибка отправки сообщения: ", err)
		return err
	}
	var response pkg.JsonResponse
	err = json.Unmarshal(resp.Response, &response)
	if err != nil {
		log.Error("Ошибка отправки сообщения: ", err)
		return err
	}
	if response != pkg.NO_ERR {
		err = errors.New(response.Message)
		log.Error("Ошибка отправки сообщения: ", err)
		return err
	}
	return nil
}

func (s *Session) Reset() {
	log.Debug("RESET")
	s.mailForm = models.MailForm{}
}

func (s *Session) Logout() error {
	log.Debug("LOGOUT")
	return nil
}