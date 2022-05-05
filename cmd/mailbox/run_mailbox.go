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

var SERVICE_PREFIX = "MailBox: "

func StartMailBoxServer(config *config.Config, db repository_proto.DatabaseRepositoryClient, profile profile_proto.ProfileClient) {
	address := fmt.Sprintf(":%v", config.Server.Services.MailBox.Port)
	log.Info(SERVICE_PREFIX, "Запуск сервера по адресу ", address)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(SERVICE_PREFIX, err)
	}
	mailboxServer := grpc.NewServer()
	mailboxService := mailbox.MailBoxService{}
	mailboxService.Init(config, db, profile)
	mailbox_proto.RegisterMailboxServer(mailboxServer, &mailboxService)
	mailboxServer.Serve(lis)
}

func main() {
	log.SetReportCaller(true)
	log.SetLevel(log.DebugLevel)
	log.Info(SERVICE_PREFIX, "Запуск сервиса")
	config, err := config.NewConfig("./configs/main.yml")
	if err != nil {
		log.Fatal(SERVICE_PREFIX, err)
	}
	dbConn, err := pkg.CreateGRPCDial(fmt.Sprintf("%v:%v", config.Server.Services.Database.Address, config.Server.Services.Database.Port))
	if err != nil {
		log.Fatal(SERVICE_PREFIX, err)
	}
	defer dbConn.Close()
	db := repository_proto.NewDatabaseRepositoryClient(dbConn)
	profileConn, err := pkg.CreateGRPCDial(fmt.Sprintf("%v:%v", config.Server.Services.Profile.Address, config.Server.Services.Profile.Port))
	if err != nil {
		log.Fatal(SERVICE_PREFIX, err)
	}
	defer profileConn.Close()
	profile := profile_proto.NewProfileClient(profileConn)
	StartMailBoxServer(config, db, profile)
}