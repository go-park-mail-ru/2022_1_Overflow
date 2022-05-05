package cmd

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/internal/middlewares"
	"OverflowBackend/internal/session"
	"OverflowBackend/pkg"
	"fmt"

	log "github.com/sirupsen/logrus"

	_ "OverflowBackend/docs"
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

	log.Info("Инициализация менджера сессий.")
	err = session.Init(config)
	if err != nil {
		log.Fatalf("Ошибка при инициализации менеджера сессий: %v", err)
	}

	log.Info("Инициализация роутеров.")
	middlewares.Init(config)
	router := RouterManager{}
	authDial, err := pkg.CreateGRPCDial(fmt.Sprintf("%v:%v", config.Server.Services.Auth.Address, config.Server.Services.Auth.Port))
	if err != nil {
		log.Fatal("Ошибка подключения к микросервису Auth:", err)
	}
	log.Info("Успешное подключение к микросервису Auth.")
	profileDial, err := pkg.CreateGRPCDial(fmt.Sprintf("%v:%v", config.Server.Services.Profile.Address, config.Server.Services.Profile.Port))
	if err != nil {
		log.Fatal("Ошибка подключения к микросервису Profile:", err)
	}
	log.Info("Успешное подключение к микросервису Profile.")
	mailboxDial, err := pkg.CreateGRPCDial(fmt.Sprintf("%v:%v", config.Server.Services.MailBox.Address, config.Server.Services.MailBox.Port))
	if err != nil {
		log.Fatal("Ошибка подключения к микросервису Mailbox:", err)
	}
	log.Info("Успешное подключение к микросервису Mailbox.")
	folderManagerDial, err := pkg.CreateGRPCDial(fmt.Sprintf("%v:%v", config.Server.Services.FolderManager.Address, config.Server.Services.FolderManager.Port))
	if err != nil {
		log.Fatal("Ошибка подключения к микросервису FolderManager:", err)
	}
	log.Info("Успешное подключение к микросервису FolderManager.")
	router.Init(config, authDial, profileDial, mailboxDial, folderManagerDial)
	HandleServer(config, router)
}
