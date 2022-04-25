package main

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/pkg"
	"OverflowBackend/proto/mailbox_proto"
	"OverflowBackend/proto/profile_proto"
	"OverflowBackend/proto/repository_proto"
	"OverflowBackend/services/mailbox"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net"

	"google.golang.org/grpc"
)

func StartMailBoxServer(config *config.Config, db repository_proto.DatabaseRepositoryClient, profile profile_proto.ProfileClient) {
	log.Info("Запуск сервера")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", config.Server.Services.MailBox.Port))
	if err != nil {
		log.Fatal(err)
	}
	mailboxServer := grpc.NewServer()
	mailboxService := mailbox.MailBoxService{}
	mailboxService.Init(config, db, profile)
	mailbox_proto.RegisterMailboxServer(mailboxServer, &mailboxService)
	mailboxServer.Serve(lis)
}

func main() {
	log.Info("Запуск сервиса Mailbox")
	config, err := config.NewConfig("./configs/main.yml")
	if err != nil {
		log.Fatal(err)
	}
	dbConn, err := pkg.CreateGRPCDial(fmt.Sprintf("%v:%v", config.Server.Services.Database.Address, config.Server.Services.Database.Port))
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()
	db := repository_proto.NewDatabaseRepositoryClient(dbConn)
	profileConn, err := pkg.CreateGRPCDial(fmt.Sprintf("%v:%v", config.Server.Services.Profile.Address, config.Server.Services.Profile.Port))
	if err != nil {
		log.Fatal(err)
	}
	defer profileConn.Close()
	profile := profile_proto.NewProfileClient(profileConn)
	StartMailBoxServer(config, db, profile)
}