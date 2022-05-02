package main

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/proto/repository_proto"
	"OverflowBackend/services/repository"
	"OverflowBackend/services/repository/postgres"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net"

	"google.golang.org/grpc"
)

var SERVICE_PREFIX = "Repository:"

func StartRepositoryServer(config *config.Config) {
	log.Info(SERVICE_PREFIX, "Запуск сервера")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", config.Server.Services.Database.Port))
	if err != nil {
		log.Fatal(err)
	}
	var dbUrl string
	var repositoryService repository.DatabaseRepository
	log.Info(SERVICE_PREFIX, "Подключение к БД.")
	dbUrl = fmt.Sprintf("postgres://%v:%v@%v:%v/%v",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Name,
	)
	repositoryService = &postgres.Database{}
	err = repositoryService.Create(dbUrl)
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}
	repositoryServer := grpc.NewServer()
	repository_proto.RegisterDatabaseRepositoryServer(repositoryServer, repositoryService)
	repositoryServer.Serve(lis)
}

func main() {
	log.Info(SERVICE_PREFIX, "Запуск сервиса Repository")
	config, err := config.NewConfig("./configs/main.yml")
	if err != nil {
		log.Fatal(SERVICE_PREFIX, err)
	}
	StartRepositoryServer(config)
}