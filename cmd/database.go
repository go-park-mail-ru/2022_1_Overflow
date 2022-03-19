package cmd

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/internal/repository"
	"OverflowBackend/internal/repository/mock"
	"OverflowBackend/internal/repository/postgres"
	"fmt"
)

func HandleDatabase(config *config.Config) (repository.DatabaseRepository, error) {
	var dbUrl string
	var db repository.DatabaseRepository
	if config.Database.Type == "postgres" {
		dbUrl = fmt.Sprintf("postgres://%v:%v@%v:%v/%v",
			config.Database.User,
			config.Database.Password,
			config.Database.Host,
			config.Database.Port,
			config.Database.Name,
		)

		db = &postgres.Database{}
	} else {
		dbUrl = "mock"
		db = &mock.MockDB{}
	}
	err := db.Create(dbUrl)
	return db, err
}