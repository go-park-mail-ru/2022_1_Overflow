package app

import (
	"OverflowBackend/internal/config"
	handlers "OverflowBackend/internal/delivery/http"
	"OverflowBackend/internal/repository/postgres"

	"fmt"

	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Application struct{}

func (app *Application) Run(configPath string) {
	var runChan = make(chan os.Signal, 1)

	config, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatalf("Ошибка при чтении конфигурационного файла сервера: %v", err)
	}

	var dbUrl string = fmt.Sprintf("postgres://%v:%v@%v:%v/%v",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Name,
	)

	db := postgres.Database{}
	err = db.Create(dbUrl)
	if err != nil {
		log.Fatalf("Ошибка при подключении к БД: %v", err)
	}

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(config.Server.Timeout.Server),
	)
	defer cancel()

	server := &http.Server {
		Handler: handlers.NewRouter(),
		ReadTimeout: time.Duration(config.Server.Timeout.Read) * time.Second,
		WriteTimeout: time.Duration(config.Server.Timeout.Write) * time.Second,
		IdleTimeout: time.Duration(config.Server.Timeout.Idle) * time.Second,
	}

	signal.Notify(runChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				log.Fatalf("Невозможно запустить сервер: %v", err)
			}
		}
	}()

	interrupt := <-runChan

	log.Printf("Сервер останавливается по причине: %+v\n", interrupt)
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Сервер не смог остановиться по причине: %+v", err)
	}
}