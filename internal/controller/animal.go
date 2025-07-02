package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"pet_shelter_and_store/internal/controller/middlewares"
	"pet_shelter_and_store/internal/errs"
	"pet_shelter_and_store/internal/models"
	"pet_shelter_and_store/internal/service"
	"strconv"
)

func GetAllAnimals(c *gin.Context) {
	animals, err := service.GetAllAnimals()
	if err != nil {
		HandleError(c, err)
	}
	c.JSON(http.StatusOK, animals)
}

func GetAnimalByID(c *gin.Context) {
	animalIDStr := c.Param("id")
	animalID, err := strconv.Atoi(animalIDStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
	}
	animal, err := service.GetAnimalByID(uint(animalID))
	if err != nil {
		HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, animal)
}

func CreateAnimal(c *gin.Context) {
	userID, userRole := c.GetUint(middlewares.UserIDCtx), c.GetString(middlewares.UserRoleCtx)

	var animal models.Animal
	if err := c.ShouldBindJSON(&animal); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	if userRole == "volunteer" || userRole == "admin" {
		animal.IsActive = true
		err := service.CreateAnimal(animal)
		if err != nil {
			HandleError(c, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Animal created successfully",
		})
		return
	} else {
		animal.IsActive = false
		err := service.CreateAnimal(animal)
		if err != nil {
			HandleError(c, err)
			return
		}
		fmt.Println(userID)

		err = service.CreateStoreRequest(
			models.Request{
				IsAdoption: false,
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

		c.JSON(http.StatusCreated, gin.H{
			"message": "Заявка на добавление животного успешно создана",
		})
		return
	}
}

func UpdateAnimal(c *gin.Context) {
	userRole := c.GetString(middlewares.UserRoleCtx)

	animalIDStr := c.Param("id")
	animalID, err := strconv.Atoi(animalIDStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	_, err = service.GetAnimalByID(uint(animalID))
	if err != nil {
		HandleError(c, err)
		return
	}

	if userRole != "volunteer" && userRole != "admin" {
		HandleError(c, errs.ErrPermissionDenied)
		return
	}

	var animalFromBody models.Animal
	if err := c.ShouldBindJSON(&animalFromBody); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	err = service.UpdateAnimal(animalFromBody)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Animal updated successfully",
	})
}
