package main

import (
	"OverflowBackend/cmd"

	log "github.com/sirupsen/logrus"
)

const configPath string = "configs/main.yml"

func main() {
	log.SetReportCaller(true)
	server := cmd.Application{}
	server.Run(configPath)
}