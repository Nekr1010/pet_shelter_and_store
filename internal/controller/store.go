package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pet_shelter_and_store/internal/controller/middlewares"
	"pet_shelter_and_store/internal/errs"
	"pet_shelter_and_store/internal/models"
	"pet_shelter_and_store/internal/service"
	"strconv"
)

func GetAllStores(c *gin.Context) {
	stores, err := service.GetAllStores()
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, stores)
}

func GetStoreByID(c *gin.Context) {
	storeIDStr := c.Param("id")
	storeID, err := strconv.Atoi(storeIDStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	store, err := service.GetStoreByID(uint(storeID))
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, store)
}

func CreateStore(c *gin.Context) {
	userRole := c.GetString(middlewares.UserRoleCtx)

	if userRole != "admin" {
		HandleError(c, errs.ErrPermissionDenied)
		return
	}

	var store models.Store
	if err := c.BindJSON(&store); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	err := service.CreateStore(store)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Store was created successfully",
	})
}

func UpdateStore(c *gin.Context) {
	storeID := c.GetUint(middlewares.StoreIDCtx)
	var store models.Store
	if err := c.BindJSON(&store); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}
	store.ID = storeID

	err := service.UpdateStore(store)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Store was updated successfully",
	})
}

func DeleteStore(c *gin.Context) {
	storeID := c.GetUint(middlewares.StoreIDCtx)

	err := service.DeleteStore(storeID)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Store was deleted successfully",
	})
}
