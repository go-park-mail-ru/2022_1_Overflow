package main

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/pkg"
	"OverflowBackend/proto/attach_proto"
	"OverflowBackend/proto/repository_proto"
	"OverflowBackend/services/attach"
	"fmt"
	"github.com/minio/minio-go/v7"
	log "github.com/sirupsen/logrus"
	"net"

	"google.golang.org/grpc"
)

var SERVICE_PREFIX = "Attach: "

func StartAttachServer(config *config.Config, db repository_proto.DatabaseRepositoryClient, s3 *minio.Client) {
	address := fmt.Sprintf(":%v", config.Server.Services.Attach.Port)
	log.Info(SERVICE_PREFIX, "Запуск сервера по адресу ", address)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(SERVICE_PREFIX, err)
	}

	attachServer := grpc.NewServer()
	attachService := attach.AttachService{}
	attachService.Init(config, db, s3)
	attach_proto.RegisterAttachServer(attachServer, &attachService)
	attachServer.Serve(lis)
}

func main() {
	log.SetReportCaller(true)
	log.SetLevel(log.DebugLevel)
	log.Info(SERVICE_PREFIX, "Запуск сервиса")
	config, err := config.NewConfig("./configs/main.yml")
	if err != nil {
		log.Fatal(SERVICE_PREFIX, err)
	}

	//Repository microservice connect
	conn, err := pkg.CreateGRPCDial(fmt.Sprintf("%v:%v", config.Server.Services.Database.Address, config.Server.Services.Database.Port))
	if err != nil {
		log.Fatal(SERVICE_PREFIX, err)
	}
	defer conn.Close()
	db := repository_proto.NewDatabaseRepositoryClient(conn)

	//Minio connect
	s3, err := pkg.NewMinioClient(config)
	if err != nil {
		log.Fatal(SERVICE_PREFIX, err)
	}
	log.Info(SERVICE_PREFIX, "Successful connect to minio.")

	StartAttachServer(config, db, s3)
}
