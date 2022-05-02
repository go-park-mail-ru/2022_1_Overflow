package main

import (
	"OverflowBackend/cmd"
	_"google.golang.org/protobuf/types/known/timestamppb"
	_"google.golang.org/protobuf/types/known/wrapperspb"
)

const configPath string = "configs/main.yml"

func main() {
	server := cmd.Application{}
	server.Run(configPath)
}