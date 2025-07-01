package service

import (
	"pet_shelter_and_store/internal/errs"
	"pet_shelter_and_store/internal/models"
	"pet_shelter_and_store/internal/repository"
)

func GetAllStores() (stores []models.Store, err error) {
	stores, err = repository.GetAllStores()
	if err != nil {
		return nil, err
	}

	return stores, nil
}

func GetStoreByID(storeID uint) (store models.Store, err error) {
	store, err = repository.GetStoreByID(storeID)
	if err != nil {
		return models.Store{}, err
	}

	return store, nil
}

func CreateStore(store models.Store) (err error) {
	_, err = GetUserByID(store.OwnerID)
	if err != nil {
		return errs.ErrUserNotFound
	}

	err = repository.CreateStore(store)
	if err != nil {
		return err
	}

	return nil
}

func UpdateStore(store models.Store) (err error) {
	err = repository.UpdateStore(store)
	if err != nil {
		return err
	}

	return nil
}

func DeleteStore(storeID uint) (err error) {
	err = repository.DeleteStore(storeID)
	if err != nil {
		return err
	}
	return nil
}
