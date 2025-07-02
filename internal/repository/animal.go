package repository

import (
	"pet_shelter_and_store/internal/db"
	"pet_shelter_and_store/internal/models"
	"pet_shelter_and_store/logger"
)

func GetAllAnimals() (animals []models.Animal, err error) {
	if err = db.GetDBConn().Model(&models.Animal{}).Find(&animals).Error; err != nil {
		logger.Error.Printf("[repository.GetAllAnimals] Error while getting all animals: %v", err)

		return nil, TranslateGormError(err)
	}

	return animals, nil
}

func GetAnimalByID(animalID uint) (animal *models.Animal, err error) {
	if err = db.GetDBConn().Model(&models.Animal{}).Where("id = ?", animalID).First(&animal).Error; err != nil {
		logger.Error.Printf("[repository.GetAnimalByID] Error while getting  animal by id: %v", err)

		return nil, TranslateGormError(err)
	}

	return animal, nil
}

func CreateAnimal(animal models.Animal) (err error) {
	if err := db.GetDBConn().Create(&animal).Error; err != nil {
		logger.Error.Printf("[repository.CreateAnimal] Error while creating animal: %v", err)

		return TranslateGormError(err)
	}

	return nil
}

func UpdateAnimal(animal models.Animal) (err error) {
	if err = db.GetDBConn().Model(&models.Animal{}).Where("id = ?", animal.ID).Updates(&animal).Error; err != nil {
		logger.Error.Printf("[repository.UpdateAnimal] Error: %v", err)

		return TranslateGormError(err)
	}

	return nil
}

func DeleteAnimal(animalID uint) (err error) {
	if err = db.GetDBConn().Delete(&models.Animal{}, animalID).Error; err != nil {
		logger.Error.Printf("[repository.DeleteAnimal] Error: %v", err)

		return TranslateGormError(err)
	}

	return nil
}
