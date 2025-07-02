package repository

import (
	"pet_shelter_and_store/internal/db"
	"pet_shelter_and_store/internal/models"
	"pet_shelter_and_store/logger"
)

func GetStoreRequests(storeID uint, isAdoption bool) (requests models.Request, err error) {
	if err = db.GetDBConn().Model(&models.Request{}).Where("store_id = ? AND is_adoption = ?", storeID, isAdoption).Find(&requests).Error; err != nil {
		logger.Error.Printf("Failed to get store requests from database: %v", err)

		return requests, TranslateGormError(err)
	}

	return requests, nil
}

func GetStoreRequestByID(requestID uint) (requests models.Request, err error) {
	if err = db.GetDBConn().Model(&models.Request{}).Where("id = ?", requestID).Find(&requests).Error; err != nil {
		logger.Error.Printf("Failed to get store requests from database: %v", err)

		return requests, TranslateGormError(err)
	}

	return requests, nil
}

func CreateStoreRequest(request models.Request) (err error) {
	if err = db.GetDBConn().Create(&request).Error; err != nil {
		logger.Error.Printf("Failed to create store request: %v", err)

		return TranslateGormError(err)
	}

	return nil
}

func AcceptStoreRequest(request models.Request) (err error) {
	if err = db.GetDBConn().Save(&request).Error; err != nil {
		logger.Error.Printf("Failed to accept store request: %v", err)

		return TranslateGormError(err)
	}

	return nil
}

func DeleteStoreRequest(request models.Request) (err error) {
	if err = db.GetDBConn().Delete(&request).Error; err != nil {
		logger.Error.Printf("Failed to delete store request: %v", err)

		return TranslateGormError(err)
	}

	return nil
}
