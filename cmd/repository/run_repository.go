package main

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/proto/repository_proto"
	"OverflowBackend/services/repository"
	"OverflowBackend/services/repository/mock"
	"OverflowBackend/services/repository/postgres"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net"

	"google.golang.org/grpc"
)

func StartRepositoryServer(config *config.Config) {
	log.Info("Запуск сервера")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", config.Server.Services.Database.Port))
	if err != nil {
		log.Fatal(err)
	}
	var dbUrl string
	var repositoryService repository.DatabaseRepository
	log.Info("Подключение к БД.")
	if config.Database.Type == "postgres" {
		dbUrl = fmt.Sprintf("postgres://%v:%v@%v:%v/%v",
			config.Database.User,
			config.Database.Password,
			config.Database.Host,
			config.Database.Port,
			config.Database.Name,
		)

		repositoryService = &postgres.Database{}
	} else {
		dbUrl = "mock"
		repositoryService = &mock.MockDB{}
	}
	err = repositoryService.Create(dbUrl)
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}
	repositoryServer := grpc.NewServer()
	repository_proto.RegisterDatabaseRepositoryServer(repositoryServer, repositoryService)
	repositoryServer.Serve(lis)
}

func main() {
	log.Info("Запуск сервиса Repository")
	config, err := config.NewConfig("./configs/main.yml")
	if err != nil {
		log.Fatal(err)
	}
	StartRepositoryServer(config)
}