package service

import (
	"pet_shelter_and_store/internal/models"
	"pet_shelter_and_store/internal/repository"
	"time"
)

func GetStoreRequests(storeID uint) (requests models.Request, err error) {
	requests, err = repository.GetStoreRequests(storeID, false)
	if err != nil {
		return models.Request{}, err
	}

	return requests, nil
}

func GetStoreAdoptions(storeID uint) (requests models.Request, err error) {
	requests, err = repository.GetStoreRequests(storeID, true)
	if err != nil {
		return models.Request{}, err
	}

	return requests, nil
}

func GetStoreRequestByID(requestID uint) (requests models.Request, err error) {
	requests, err = repository.GetStoreRequestByID(requestID)
	if err != nil {
		return models.Request{}, err
	}

	return requests, nil
}

func CreateStoreRequest(request models.Request) (err error) {
	err = repository.CreateStoreRequest(request)
	if err != nil {
		return err
	}

	return nil
}

// принимает заявки как на усыновления(от юзера создавшего заявку) также на прием в приют(от юзера создавшего заявку)
func AcceptStoreRequest(requestID uint) (err error) {
	request, err := GetStoreRequestByID(requestID)
	if err != nil {
		return err
	}

	if request.IsAdoption {
		request.IsAccepted = true
		animal, err := repository.GetAnimalByID(request.AnimalID)
		if err != nil {
			return err
		}

		animal.TakenAt = time.Now()
		err = repository.UpdateAnimal(*animal)
		if err != nil {
			return err
		}

		err = repository.AcceptStoreRequest(request)
		if err != nil {
			return err
		}

		return nil
	}

	request.IsAccepted = true

	animal, err := repository.GetAnimalByID(request.AnimalID)
	if err != nil {
		return err
	}

	animal.IsActive = true
	err = repository.UpdateAnimal(*animal)
	if err != nil {
		return err
	}

	err = repository.AcceptStoreRequest(request)
	if err != nil {
		return err
	}

	return nil
}

func DeleteStoreRequest(requestID uint) (err error) {
	request, err := GetStoreRequestByID(requestID)
	if err != nil {
		return err
	}

	if request.IsAdoption {
		animal, err := repository.GetAnimalByID(request.AnimalID)
		if err != nil {
			return err
		}

		animal.IsReserved = false
		err = repository.UpdateAnimal(*animal)
		if err != nil {
			return err
		}

		err = repository.DeleteStoreRequest(request)
		if err != nil {
			return err
		}

		return nil
	}

	err = repository.DeleteAnimal(request.AnimalID)
	if err != nil {
		return err
	}

	err = repository.DeleteStoreRequest(request)
	if err != nil {
		return err
	}

	return nil
}
