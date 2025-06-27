package main

import (
	"log"
	"pet_shelter_and_store/internal/configs"
	"pet_shelter_and_store/internal/controller"
	"pet_shelter_and_store/internal/db"
	"pet_shelter_and_store/logger"
)

func main() {
	// Reading configs
	if err := configs.ReadSettings(); err != nil {
		log.Fatalf("Ошибка чтения настроек: %s", err)
	}

	// Initializing logger
	if err := logger.Init(); err != nil {
		return
	}
	logger.Info.Println("Loggers initialized successfully!")

	// Connecting to db
	if err := db.ConnectDB(); err != nil {
		return
	}
	logger.Info.Println("Connection to database established successfully!")

	// Initializing db-migrations
	if err := db.InitMigrations(); err != nil {
		return
	}
	logger.Info.Println("Migrations initialized successfully!")

	// Running http-server
	if err := controller.RunServer(); err != nil {
		return
	}

}
