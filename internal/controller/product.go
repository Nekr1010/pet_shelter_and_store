package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pet_shelter_and_store/internal/errs"
	"pet_shelter_and_store/internal/service"
	"strconv"
)

func GetProductByID(c *gin.Context) {
	productIDStr := c.Param("id")
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}
	product, err := service.GetProductByID(uint(productID))
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}
	c.JSON(http.StatusOK, product)
}

func GetProductByCategory(c *gin.Context) {
	productByCategory := c.Query("category")
	products, err := service.GetProductsByCategory(productByCategory)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}
	c.JSON(http.StatusOK, products)
}

//func GetProductsByStore(c *gin.Context) {
//	storeIDStr := c.Query("stodeID")
//	productID, err := strconv.Atoi(productIDStr)
//	if err != nil {
//		HandleError(c, errs.ErrInvalidID)
//		return
//	}
//	c.JSON(http.StatusOK, products)
//}
