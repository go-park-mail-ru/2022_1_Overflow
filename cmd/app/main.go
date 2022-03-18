package main

import "OverflowBackend/internal/app"

var configPath string = "../../configs/main.yml"

func main() {
	server := app.Application{}
	server.Run(configPath)
}