package repository

import (
	"pet_shelter_and_store/internal/db"
	"pet_shelter_and_store/internal/models"
	"pet_shelter_and_store/logger"
)

func GetAllStores() (stores []models.Store, err error) {
	if err = db.GetDBConn().Model(&models.Store{}).Find(&stores).Error; err != nil {
		logger.Error.Printf("[repository.GetAllStores] Error while getting all stores: %v", err)

		return nil, TranslateGormError(err)
	}

	return stores, nil
}

func GetStoreByID(storeID uint) (store models.Store, err error) {
	if err = db.GetDBConn().Model(&models.Store{}).Where("id = ?", storeID).First(&store).Error; err != nil {
		logger.Error.Printf("[repository.GetStoreByID] Error: %v", err)

		return store, TranslateGormError(err)
	}

	return store, nil
}

func CreateStore(store models.Store) (err error) {
	if err = db.GetDBConn().Model(&models.Store{}).Create(&store).Error; err != nil {
		logger.Error.Printf("[repository.CreateStore] Error: %v", err)

		return TranslateGormError(err)
	}

	return nil
}

func UpdateStore(store models.Store) (err error) {
	if err = db.GetDBConn().Model(&models.Store{}).Where("id = ?", store.ID).Updates(&store).Error; err != nil {
		logger.Error.Printf("[repository.UpdateStore] Error: %v", err)
		return TranslateGormError(err)
	}

	return nil
}

func DeleteStore(storeID uint) (err error) {
	if err = db.GetDBConn().Delete(&models.Store{}, storeID).Error; err != nil {
		logger.Error.Printf("[repository.DeleteStore] Error: %v", err)

		return TranslateGormError(err)
	}

	return nil
}
