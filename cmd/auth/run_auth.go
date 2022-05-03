package main

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/pkg"
	"OverflowBackend/proto/auth_proto"
	"OverflowBackend/proto/repository_proto"
	"OverflowBackend/services/auth"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net"

	"google.golang.org/grpc"
)

var SERVICE_PREFIX = "Auth: "

func StartAuthServer(config *config.Config, db repository_proto.DatabaseRepositoryClient) {
	address := fmt.Sprintf(":%v", config.Server.Services.Auth.Port)
	log.Info(SERVICE_PREFIX, "Запуск сервера по адресу ", address)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(SERVICE_PREFIX, err)
	}
	authServer := grpc.NewServer()
	authService := auth.AuthService{}
	authService.Init(config, db)
	auth_proto.RegisterAuthServer(authServer, &authService)
	authServer.Serve(lis)
}

func main() {
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
	StartAuthServer(config, db)
}