package repository

import (
	"pet_shelter_and_store/internal/db"
	"pet_shelter_and_store/internal/models"
	"pet_shelter_and_store/logger"
)

func GetAllRequests() (requests []models.Request, err error) {
	if err = db.GetDBConn().Model(&models.Request{}).Find(&requests).Error; err != nil {
		logger.Error.Print(err)

		return []models.Request{}, err
	}

	return requests, nil

}
func GetAllAdoptionRequestsByStoreID(storeID uint) (requests []models.Request, err error) {
	if err = db.GetDBConn().Model(&models.Request{}).Where("store_id = ? AND is_adoption = true", storeID).Find(&requests).Error; err != nil {
		logger.Error.Print(err)

		return []models.Request{}, err
	}

	return requests, nil
}

func GetAllSurrenderRequestsByStoreID(storeID uint) (requests []models.Request, err error) {
	if err = db.GetDBConn().Model(&models.Request{}).Where("is_adoption = false").Find(&requests).Error; err != nil {
		logger.Error.Print(err)

		return []models.Request{}, err
	}

	return requests, nil
}
