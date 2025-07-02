package service

import (
	"pet_shelter_and_store/internal/models"
	"pet_shelter_and_store/internal/repository"
)

func GetAllAnimals() (animals []models.Animal, err error) {
	animals, err = repository.GetAllAnimals()
	if err != nil {
		return nil, err
	}

	return animals, nil
}

func GetAnimalByID(animalID uint) (animal *models.Animal, err error) {
	animal, err = repository.GetAnimalByID(animalID)
	if err != nil {
		return nil, err
	}

	return animal, nil
}

func CreateAnimal(animal models.Animal) (err error) {
	err = repository.CreateAnimal(animal)
	if err != nil {
		return err
	}

	return nil
}

func UpdateAnimal(animal models.Animal) (err error) {
	err = repository.UpdateAnimal(animal)
	if err != nil {
		return err
	}

	return nil
}

func DeleteAnimal(animalID uint) (err error) {
	err = repository.DeleteAnimal(animalID)
	if err != nil {
		return err
	}

	return nil
}
