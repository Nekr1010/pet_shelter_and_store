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
	router.GET("/store", GetAllStores)
	router.GET("/:id", GetStoreByID)
	storesG := apiG.Group("/store")
	{
		storesG.POST("", CreateStore)
		storesG.PATCH("/:id", middlewares.CheckUserStorePermission, UpdateStore)
		storesG.DELETE("/:id", middlewares.CheckUserStorePermission, DeleteStore)
	}
	requestG := apiG.Group("requests")
	adoptionRequestG := requestG.Group("/adoption")
	{
		adoptionRequestG.POST("")
		adoptionRequestG.GET("")
		adoptionRequestG.GET("/:id")
		adoptionRequestG.PUT("/:id")
		adoptionRequestG.DELETE("")
	}
	//
	//animalsG := router.Group("/animals")
	//{
	//	animalsG.GET("", GetAllAnimals)
	//	animalsG.GET("/:id", GetAnimalByID)
	//}
	//
	//animalsSurender := apiG.POST("", AnimalsSurendig)
	//
	//requests := apiG.Group("/requests")
	//{
	//	requests.POST("", CreateRequest)
	//	requests.GET("", GetAllRequests)
	//	requests.GET("/:id", GetRequestByID)
	//}

	if err := router.Run(configs.AppSettings.AppParams.PortRun); err != nil {
		logger.Error.Printf("[controller] RunServer():  Error during running HTTP server: %s", err.Error())
		return err
	}

	return nil
}
