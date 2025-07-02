package repository

import (
	"pet_shelter_and_store/internal/db"
	"pet_shelter_and_store/internal/models"
	"pet_shelter_and_store/logger"
)

func GetProductByID(productID uint) (product models.Product, err error) {
	if err = db.GetDBConn().Model(&models.Product{}).Where("id = ?", productID).First(&product).Error; err != nil {
		logger.Error.Printf("[repository.GetProductByID] Error while getting product by id")
		return models.Product{}, TranslateGormError(err)
	}
	return product, nil
}

func GetProductsByCategory(productCategory string) (products []models.Product, err error) {
	if err = db.GetDBConn().Model(&models.Product{}).Where("category = ?", productCategory).Find(&products).Error; err != nil {
		logger.Error.Printf("[repository.GetProductByID] Error while getting product by id")
		return nil, TranslateGormError(err)
	}
	return products, nil
}

func GetProductsByStore(storeID uint) (products []models.Product, err error) {
	if err = db.GetDBConn().Model(&models.Product{}).Where("storeID = ?", storeID).Find(&products).Error; err != nil {
		logger.Error.Printf("[repository.GetProductByID] Error while getting product by id")
		return nil, TranslateGormError(err)
	}
	return products, nil
}

func UpdateProduct(product models.Product) (err error) {
	if err = db.GetDBConn().Model(&models.Product{}).Where("id = ?", product.ID).Updates(&product).Error; err != nil {
		logger.Error.Printf("[repository.GetProductByID] Error while getting product by id")
		return TranslateGormError(err)
	}
	return nil
}
func DeleteProduct(productID uint) (err error) {
	if err = db.GetDBConn().Delete(&models.Store{}, productID).Error; err != nil {
		logger.Error.Printf("[repository.DeleteProduct] Error: %v", err)
		return TranslateGormError(err)
	}
	return nil
}
