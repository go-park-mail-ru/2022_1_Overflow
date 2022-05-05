package main

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/pkg"
	"OverflowBackend/proto/profile_proto"
	"OverflowBackend/proto/repository_proto"
	"OverflowBackend/services/profile"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net"

	"google.golang.org/grpc"
)

var SERVICE_PREFIX = "Profile: "

func StartProfileServer(config *config.Config, db repository_proto.DatabaseRepositoryClient) {
	address := fmt.Sprintf(":%v", config.Server.Services.Profile.Port)
	log.Info(SERVICE_PREFIX, "Запуск сервера по адресу ", address)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(SERVICE_PREFIX, err)
	}
	profileServer := grpc.NewServer()
	profileService := profile.ProfileService{}
	profileService.Init(config, db)
	profile_proto.RegisterProfileServer(profileServer, &profileService)
	profileServer.Serve(lis)
}

func main() {
	log.SetReportCaller(true)
	log.SetLevel(log.DebugLevel)
	log.Info(SERVICE_PREFIX, "Запуск сервиса")
	config, err := config.NewConfig("./configs/main.yml")
	if err != nil {
		log.Fatal(SERVICE_PREFIX, err)
	}
	conn, err := pkg.CreateGRPCDial(fmt.Sprintf("%v:%v", config.Server.Services.Database.Address, config.Server.Services.Database.Port))
	if err != nil {
		log.Fatal(SERVICE_PREFIX, err)
	}
	defer conn.Close()
	db := repository_proto.NewDatabaseRepositoryClient(conn)
	StartProfileServer(config, db)
}
