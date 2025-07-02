package service

import (
	"pet_shelter_and_store/internal/models"
	"pet_shelter_and_store/internal/repository"
)

func GetProductByID(productID uint) (product models.Product, err error) {
	product, err = repository.GetProductByID(productID)
	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func GetProductsByCategory(productCategory string) (products []models.Product, err error) {
	products, err = repository.GetProductsByCategory(productCategory)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func GetProductsByStore(storeID uint) (products []models.Product, err error) {
	products, err = repository.GetProductsByStore(storeID)
	if err != nil {
		return nil, err
	}

	return products, nil
}
func UpdateProduct(product models.Product) (err error) {
	if err = repository.UpdateProduct(product); err != nil {
		return err
	}
	return nil
}
func DeleteProduct(productID uint) (err error) {
	if err = repository.DeleteProduct(productID); err != nil {
		return err
	}
	return nil
}
