package main

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/pkg"
	"OverflowBackend/proto/folder_manager_proto"
	"OverflowBackend/proto/profile_proto"
	"OverflowBackend/proto/repository_proto"
	"OverflowBackend/services/folder_manager"
	"fmt"
	"net"

	log "github.com/sirupsen/logrus"

	"google.golang.org/grpc"
)

func StartFolderManagerServer(config *config.Config, db repository_proto.DatabaseRepositoryClient, profile profile_proto.ProfileClient) {
	log.Info("Запуск сервера")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", config.Server.Services.FolderManager.Port))
	if err != nil {
		log.Fatal(err)
	}
	folderManagerServer := grpc.NewServer()
	folderManagerService := folder_manager.FolderManagerService{}
	folderManagerService.Init(config, db, profile)
	folder_manager_proto.RegisterFolderManagerServer(folderManagerServer, &folderManagerService)
	folderManagerServer.Serve(lis)
}

func main() {
	log.Info("Запуск сервиса FolderManager")
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
	profileConn, err := pkg.CreateGRPCDial(fmt.Sprintf("%v:%v", config.Server.Services.Profile.Address, config.Server.Services.Profile.Port))
	if err != nil {
		log.Fatal(err)
	}
	defer profileConn.Close()
	profile := profile_proto.NewProfileClient(profileConn)
	StartFolderManagerServer(config, db, profile)
}