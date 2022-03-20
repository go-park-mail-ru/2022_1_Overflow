package main

import "OverflowBackend/cmd"

const configPath string = "configs/main.yml"

func main() {
	server := cmd.Application{}
	server.Run(configPath)
}