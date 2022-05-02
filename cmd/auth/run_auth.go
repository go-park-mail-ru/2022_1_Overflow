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

func StartAuthServer(config *config.Config, db repository_proto.DatabaseRepositoryClient) {
	log.Info("Запуск сервера")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", config.Server.Services.Auth.Port))
	if err != nil {
		log.Fatal(err)
	}
	authServer := grpc.NewServer()
	authService := auth.AuthService{}
	authService.Init(config, db)
	auth_proto.RegisterAuthServer(authServer, &authService)
	authServer.Serve(lis)
}

func main() {
	log.Info("Запуск сервиса Auth")
	config, err := config.NewConfig("./configs/main.yml")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := pkg.CreateGRPCDial(fmt.Sprintf("%v:%v", config.Server.Services.Database.Address, config.Server.Services.Database.Port))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	db := repository_proto.NewDatabaseRepositoryClient(conn)
	StartAuthServer(config, db)
}