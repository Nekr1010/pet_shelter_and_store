package controller

import (
	"github.com/gin-gonic/gin"
	"pet_shelter_and_store/internal/configs"
	"pet_shelter_and_store/internal/controller/middlewares"
	"pet_shelter_and_store/logger"
)

func RunServer() error {
	router := gin.Default()

	router.GET("/", Ping)

	authG := router.Group("/auth")
	{
		authG.POST("/sign-up", SignUp)
		authG.POST("/sign-in", SignIn)
		authG.POST("/refresh", RefreshToken)
	}

	apiG := router.Group("", middlewares.CheckUserAuthentication)
	router.GET("store", GetAllStores)
	router.GET("store/:id", GetStoreByID)
	storesG := apiG.Group("/store")
	{
		storesG.POST("", CreateStore)
		storesG.PATCH("/:id", middlewares.CheckUserStorePermission, UpdateStore)
		storesG.DELETE("/:id", middlewares.CheckUserStorePermission, DeleteStore)
	}

	requestG := apiG.Group("requests", middlewares.CheckUserStorePermission)
	{
		requestG.GET("/store/:id", GetAllStoreRequests)
		requestG.GET("/:id", GetAllStoreRequestByID)
		requestG.PATCH("/:id", AcceptStoreRequest)
		requestG.DELETE("/:id", DeleteStoreRequest)
	}

	adoptionRequestG := requestG.Group("/adoption")
	{
		adoptionRequestG.GET("store/:id", GetAllStoreAdoptions)
		adoptionRequestG.POST("/:id", CreateStoreAdoption)
	}

	animalsG := router.Group("/animals")
	{
		animalsG.GET("", GetAllAnimals)
		animalsG.GET(":id", GetAnimalByID)
		animalsG.POST("", middlewares.CheckUserAuthentication, CreateAnimal)
		animalsG.PATCH("/:id", UpdateAnimal)
	}

	if err := router.Run(configs.AppSettings.AppParams.PortRun); err != nil {
		logger.Error.Printf("[controller] RunServer():  Error during running HTTP server: %s", err.Error())
		return err
	}

	return nil
}
