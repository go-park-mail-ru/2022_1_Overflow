package cmd

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/internal/delivery"
	"log"
)

type Application struct{}

func (app *Application) Run(configPath string) {
	config, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatalf("Ошибка при чтении конфигурационного файла сервера: %v", err)
	}

	db, err := HandleDatabase(config)
	
	if err != nil {
		log.Fatalf("Ошибка при подключении к БД: %v", err)
	}

	router:= delivery.RouterManager{}
	router.Init(db)

	HandleServer(config, router)
}