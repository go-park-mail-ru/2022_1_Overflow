package main

import (
	"OverflowBackend/cmd"

	log "github.com/sirupsen/logrus"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
)

const configPath string = "configs/main.yml"

func main() {
	log.SetReportCaller(true)
	server := cmd.Application{}
	server.Run(configPath)
}