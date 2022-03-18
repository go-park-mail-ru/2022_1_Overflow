package main

import "OverflowBackend/cmd"

var configPath string = "../../configs/main.yml"

func main() {
	server := cmd.Application{}
	server.Run(configPath)
}