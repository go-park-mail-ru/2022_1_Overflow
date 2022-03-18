package cmd

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/internal/repository"
	"OverflowBackend/internal/repository/postgres"
	"fmt"
)

func HandleDatabase(config *config.Config) (repository.DatabaseRepository, error) {
	var dbUrl string = fmt.Sprintf("postgres://%v:%v@%v:%v/%v",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Name,
	)

	db := postgres.Database{}
	err := db.Create(dbUrl)
	return &db, err
}