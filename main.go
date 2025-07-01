package main

import (
	"github.com/joho/godotenv"
	"log"
	"pet_shelter_and_store/internal/configs"
	"pet_shelter_and_store/internal/controller"
	"pet_shelter_and_store/internal/db"
	"pet_shelter_and_store/logger"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	if err := configs.ReadSettings(); err != nil {
		log.Fatalf("Ошибка чтения настроек: %s", err)
	}

	if err := logger.Init(); err != nil {
		panic(err)
	}

	if err := db.ConnectToDB(); err != nil {
		panic(err)
	}

	if err := db.Migrate(); err != nil {
		panic(err)
	}

	if err := controller.RunServer(); err != nil {
		return
	}
}
