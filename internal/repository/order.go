package repository

import (
	"pet_shelter_and_store/internal/db"
	"pet_shelter_and_store/internal/models"
	"pet_shelter_and_store/logger"
)

func GetOrderByID(orderID uint) (order models.Order, err error) {
	if err = db.GetDBConn().Model(&models.Order{}).Where("id = ?", orderID).First(&order).Error; err != nil {
		logger.Error.Printf("[repository.GetOrderByID] Error while getting  order by id: %v", err)
		return models.Order{}, TranslateGormError(err)
	}
	return order, nil
}
func GetUserOrders(userID uint) (orders []models.Order, err error) {
	if err = db.GetDBConn().Model(&models.Order{}).Where("userID = ?", userID).First(&orders).Error; err != nil {
		logger.Error.Printf("[repository.GetOrderByID] Error while getting  order by id: %v", err)
		return nil, TranslateGormError(err)
	}
	return orders, nil
}
func GetStoreOrders(storeID uint) (orders []models.Order, err error) {
	if err = db.GetDBConn().Model(&models.Order{}).Where("storeID = ?", storeID).First(&orders).Error; err != nil {
		logger.Error.Printf("[repository.GetOrderByID] Error while getting  order by id: %v", err)
		return nil, TranslateGormError(err)
	}
	return orders, nil
}
func PatchOrder(orderID uint) (err error) {
	order, err := GetOrderByID(orderID)
	if err != nil {
		return err
	}
	order.IsActive = false
	if err := db.GetDBConn().Save(&order).Error; err != nil {
		logger.Error.Printf("[repository.PatchOrder] Error while updating order: %v", err)
		return TranslateGormError(err)
	}
	return nil
}
func CreateOrder(order models.Order) (err error) {
	if err = db.GetDBConn().Create(&order).Error; err != nil {
		logger.Error.Printf("[repository.CreateOrder] Error while creating order: %v", err)
		return TranslateGormError(err)
	}
	return nil
}
