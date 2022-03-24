package cmd

import (
	"OverflowBackend/internal/config"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func HandleServer(config *config.Config, router RouterManager) {
	server := &http.Server{
		Addr:         fmt.Sprintf(":%v", config.Server.Port),
		Handler:      router.NewRouter(config.Server.Port),
		ReadTimeout:  config.Server.Timeout.Read,
		WriteTimeout: config.Server.Timeout.Write,
		IdleTimeout:  config.Server.Timeout.Idle,
	}

	log.Info("Запускаю сервер по адресу: ", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		if err != http.ErrServerClosed {
			log.Fatalf("Невозможно запустить сервер: %v", err)
		}
	}
}
