package db

import (
	"pet_shelter_and_store/internal/models"
	"pet_shelter_and_store/logger"
)

func Migrate() error {
	err := dbConn.AutoMigrate(
		&models.User{},
		&models.Store{},
		&models.Product{},
		&models.Animal{},
		&models.Order{},
		&models.Request{},
	)
	if err != nil {
		logger.Error.Printf("[db.Migrate] Error while migrating models: %v", err)

		return err
	}

	return nil
}
