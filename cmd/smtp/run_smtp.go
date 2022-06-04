package smtp_server

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/internal/smtp"
	"OverflowBackend/pkg"
	"OverflowBackend/proto/auth_proto"
	"OverflowBackend/proto/folder_manager_proto"
	"OverflowBackend/proto/mailbox_proto"
	"OverflowBackend/proto/profile_proto"
	"crypto/tls"
	"fmt"
	"time"

	"github.com/emersion/go-smtp"
	log "github.com/sirupsen/logrus"
	sasl "github.com/emersion/go-sasl"
)

func initServer(configPath string) *smtp_server.SMTPServer {
	log.Info("Чтение конфигурационного файла сервера.")
	config, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatalf("Ошибка при чтении конфигурационного файла сервера: %v", err)
	}

	authDial, err := pkg.CreateGRPCDial(fmt.Sprintf("%v:%v", config.Server.Services.Auth.Address, config.Server.Services.Auth.Port))
	if err != nil {
		log.Fatal("Ошибка подключения к микросервису Auth:", err)
	}
	log.Info("Успешное подключение к микросервису Auth.")
	profileDial, err := pkg.CreateGRPCDial(fmt.Sprintf("%v:%v", config.Server.Services.Profile.Address, config.Server.Services.Profile.Port))
	if err != nil {
		log.Fatal("Ошибка подключения к микросервису Profile:", err)
	}
	log.Info("Успешное подключение к микросервису Profile.")
	mailboxDial, err := pkg.CreateGRPCDial(fmt.Sprintf("%v:%v", config.Server.Services.MailBox.Address, config.Server.Services.MailBox.Port))
	if err != nil {
		log.Fatal("Ошибка подключения к микросервису Mailbox:", err)
	}
	log.Info("Успешное подключение к микросервису Mailbox.")
	folderManagerDial, err := pkg.CreateGRPCDial(fmt.Sprintf("%v:%v", config.Server.Services.FolderManager.Address, config.Server.Services.FolderManager.Port))
	if err != nil {
		log.Fatal("Ошибка подключения к микросервису FolderManager:", err)
	}
	log.Info("Успешное подключение к микросервису FolderManager.")
	server := &smtp_server.SMTPServer{}
	server.Init(
		config,
		auth_proto.NewAuthClient(authDial),
		profile_proto.NewProfileClient(profileDial),
		mailbox_proto.NewMailboxClient(mailboxDial),
		folder_manager_proto.NewFolderManagerClient(folderManagerDial),
	)
	return server
}

func SetupServer() {
	log.SetLevel(log.DebugLevel)
	server := initServer("./configs/main.yml")
	s := smtp.NewServer(server)
	s.Addr = ":25"
	s.Domain = "localhost"
	s.ReadTimeout = 10 * time.Second
	s.WriteTimeout = 10 * time.Second
	s.MaxMessageBytes = 1024 * 1024
	s.MaxRecipients = 50
	s.AllowInsecureAuth = true
	// force TLS for auth
	s.AllowInsecureAuth = false
	// Load the certificate and key
	cer, err := tls.LoadX509KeyPair("/etc/server.crt", "/etc/server.key")
	if err != nil {
		log.Fatal(err)
		return
	}
	// Configure the TLS support
	// Add deprecated LOGIN auth method as some clients haven't learned
	s.EnableAuth(sasl.Login, func(conn *smtp.Conn) sasl.Server {
		return sasl.NewLoginServer(func(username, password string) error {
			state := conn.State()
			session, err := server.Login(&state, username, password)
			if err != nil {
				return err
			}

			conn.SetSession(session)
			return nil
		})
	})
	s.TLSConfig = &tls.Config{Certificates: []tls.Certificate{cer}}
	log.Println("Starting server at", s.Addr)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func main() {

}