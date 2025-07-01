package db

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"pet_shelter_and_store/internal/configs"
)

var (
	dbConn *gorm.DB
)

func ConnectToDB() error {
	cfg := configs.AppSettings.PostgresParams

	connStr := fmt.Sprintf(`host=%s
							port=%s 
							user=%s 
							password=%s 
							dbname=%s 
							sslmode=disable`,
		cfg.Host,
		cfg.Port,
		cfg.User,
		os.Getenv("DB_PASSWORD"),
		cfg.Database,
	)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return err
	}

	dbConn = db

	return nil
}

func GetDBConn() *gorm.DB {
	return dbConn
}

func CloseDBConn() error {
	if sqlDB, err := GetDBConn().DB(); err == nil {
		if err = sqlDB.Close(); err != nil {
			log.Fatalf("Error while closing DB: %s", err)
		}
		fmt.Println("Connection closed successfully")
	} else {
		log.Fatalf("Error while getting *sql.DB from GORM: %s", err)
	}

	return nil
}
