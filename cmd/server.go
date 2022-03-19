package cmd

import (
	"OverflowBackend/internal/config"
	"fmt"
	"log"
	"net/http"
)

func HandleServer(config *config.Config, router RouterManager) {
	server := &http.Server {
		Addr: fmt.Sprintf(":%v", config.Server.Port),
		Handler: router.NewRouter(),
		ReadTimeout: config.Server.Timeout.Read,
		WriteTimeout: config.Server.Timeout.Write,
		IdleTimeout: config.Server.Timeout.Idle,
	}

	log.Printf("Запускаю сервер по адресу: %v\n", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		if err != http.ErrServerClosed {
			log.Fatalf("Невозможно запустить сервер: %v", err)
		}
	}
}