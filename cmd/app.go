package cmd

import (
	"OverflowBackend/internal/config"
	"log"
)

type Application struct{}

// @title OverMail API
// @version 1.0
// @description API почтового сервиса команды Overflow.

// @contact.name Роман Медников
// @contact.url https://vk.com/l____l____l____l____l____l
// @contact.email jellybe@yandex.ru

// @BasePath /
func (app *Application) Run(configPath string) {
	config, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatalf("Ошибка при чтении конфигурационного файла сервера: %v", err)
	}

	db, err := HandleDatabase(config)
	
	if err != nil {
		log.Fatalf("Ошибка при подключении к БД: %v", err)
	}

	router:= RouterManager{}
	router.Init(db, config)

	HandleServer(config, router)
}