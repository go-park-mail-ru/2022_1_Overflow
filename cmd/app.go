package cmd

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/internal/middlewares"
	"OverflowBackend/internal/usecase/session"

	log "github.com/sirupsen/logrus"
)

type Application struct{}

// @title OverMail API
// @version 1.0
// @description API почтового сервиса команды Overflow.

// @contact.name Роман Медников
// @contact.url https://vk.com/l____l____l____l____l____l
// @contact.username jellybe@yandex.ru

// @BasePath /
func (app *Application) Run(configPath string) {
	log.Info("Чтение конфигурационного файла сервера.")
	config, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatalf("Ошибка при чтении конфигурационного файла сервера: %v", err)
	}

	log.Info("Подключение к БД.")
	db, err := HandleDatabase(config)

	if err != nil {
		log.Fatalf("Ошибка при подключении к БД: %v", err)
	}

	log.Info("Инициализация менджера сессий.")
	err = session.Init(config)
	if err != nil {
		log.Fatalf("Ошибка при инициализации менеджера сессий: %v", err)
	}
	
	log.Info("Инициализация роутеров.")
	middlewares.Init(config)
	router := RouterManager{}
	router.Init(db, config)

	HandleServer(config, router)
}
