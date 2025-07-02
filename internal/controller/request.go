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

func GetAllStoreRequests(c *gin.Context) {
	storeIDStr := c.Param("id")
	storeID, err := strconv.Atoi(storeIDStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	requests, err := service.GetStoreRequests(uint(storeID))
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, requests)
}

func GetAllStoreAdoptions(c *gin.Context) {
	storeIDStr := c.Param("id")
	storeID, err := strconv.Atoi(storeIDStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	requests, err := service.GetStoreAdoptions(uint(storeID))
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, requests)
}

func GetAllStoreRequestByID(c *gin.Context) {
	requestIDStr := c.Param("id")
	requestID, err := strconv.Atoi(requestIDStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	request, err := service.GetStoreRequestByID(uint(requestID))
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, request)
}

// содание заяки на сдачу животного из бд
func CreateStoreAdoption(c *gin.Context) {
	userID := c.GetUint(middlewares.UserIDCtx)

	animalIDStr := c.Param("id")
	animalID, err := strconv.Atoi(animalIDStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	animal, err := service.GetAnimalByID(uint(animalID))
	if err != nil {
		HandleError(c, err)
		return
	}

	err = service.CreateStoreRequest(
		models.Request{
			IsAdoption: true,
			IsAccepted: false,
			UserID:     userID,
			StoreID:    animal.StoreID,
			AnimalID:   animal.ID,
		},
	)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "store adoption added successfully",
	})
}

func AcceptStoreRequest(c *gin.Context) {
	requestIDStr := c.Param("id")
	requestID, err := strconv.Atoi(requestIDStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	err = service.AcceptStoreRequest(uint(requestID))
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "accept store request successfully",
	})
}

func DeleteStoreRequest(c *gin.Context) {
	requestIDStr := c.Param("id")
	requestID, err := strconv.Atoi(requestIDStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	err = service.DeleteStoreRequest(uint(requestID))
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "deleted store request successfully",
	})
}
