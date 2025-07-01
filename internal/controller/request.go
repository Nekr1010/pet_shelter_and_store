package controller

//
//import (
//	"github.com/gin-gonic/gin"
//	"net/http"
//	"pet_shelter_and_store/internal/controller/middlewares"
//	"pet_shelter_and_store/internal/errs"
//	"pet_shelter_and_store/internal/models"
//	"pet_shelter_and_store/internal/service"
//	"strconv"
//)
//
//func CreateRequest(c *gin.Context) {
//	userID := c.GetUint(middlewares.UserIDCtx)
//
//	var request models.Request
//	if err := c.BindJSON(&request); err != nil {
//		HandleError(c, errs.ErrValidationFailed)
//		return
//	}
//
//	request.UserID = userID
//
//	err := service.CreateRequest(request)
//	if err != nil {
//		HandleError(c, err)
//		return
//	}
//
//	c.JSON(http.StatusCreated, gin.H{
//		"message": "Request was created successfully",
//	})
//}
//
//func GetRequestByID(c *gin.Context) {
//	requestIDStr := c.Param("id")
//	requestID, err := strconv.Atoi(requestIDStr)
//	if err != nil {
//		HandleError(c, errs.ErrInvalidID)
//		return
//	}
//
//	request, err := service.GetRequestByID(uint(requestID))
//	if err != nil {
//		HandleError(c, err)
//		return
//	}
//
//	// Проверка прав доступа
//	userID := c.GetUint(middlewares.UserIDCtx)
//	userRole := c.GetString(middlewares.UserRoleCtx)
//
//	if userRole != "admin" && userRole != "shop_admin" && request.UserID != userID {
//		HandleError(c, errs.ErrPermissionDenied)
//		return
//	}
//
//	c.JSON(http.StatusOK, request)
//}
//
//func GetUserRequests(c *gin.Context) {
//	userID := c.GetUint(middlewares.UserIDCtx)
//
//	requests, err := service.GetRequestsByUserID(userID)
//	if err != nil {
//		HandleError(c, err)
//		return
//	}
//
//	c.JSON(http.StatusOK, requests)
//}
//
//func GetShopRequests(c *gin.Context) {
//	userRole := c.GetString(middlewares.UserRoleCtx)
//	shopIDStr := c.Param("shop_id")
//	shopID, err := strconv.Atoi(shopIDStr)
//	if err != nil {
//		HandleError(c, errs.ErrInvalidID)
//		return
//	}
//
//	// Проверка что пользователь админ этого магазина
//	if userRole != "admin" && userRole != "shop_admin" {
//		HandleError(c, errs.ErrPermissionDenied)
//		return
//	}
//
//	requests, err := service.GetRequestsByShopID(uint(shopID))
//	if err != nil {
//		HandleError(c, err)
//		return
//	}
//
//	c.JSON(http.StatusOK, requests)
//}
//
//func UpdateRequestStatus(c *gin.Context) {
//	userRole := c.GetString(middlewares.UserRoleCtx)
//	if userRole != "admin" && userRole != "shop_admin" {
//		HandleError(c, errs.ErrPermissionDenied)
//		return
//	}
//
//	requestIDStr := c.Param("id")
//	requestID, err := strconv.Atoi(requestIDStr)
//	if err != nil {
//		HandleError(c, errs.ErrInvalidID)
//		return
//	}
//
//	var statusUpdate struct {
//		Status string `json:"status"`
//		Notes  string `json:"notes"`
//	}
//	if err := c.BindJSON(&statusUpdate); err != nil {
//		HandleError(c, errs.ErrValidationFailed)
//		return
//	}
//
//	err = service.UpdateRequestStatus(uint(requestID), statusUpdate.Status, statusUpdate.Notes, c.GetUint(middlewares.UserIDCtx))
//	if err != nil {
//		HandleError(c, err)
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"message": "Request status updated successfully",
//	})
//}
//
//func CancelRequest(c *gin.Context) {
//	userID := c.GetUint(middlewares.UserIDCtx)
//	requestIDStr := c.Param("id")
//	requestID, err := strconv.Atoi(requestIDStr)
//	if err != nil {
//		HandleError(c, errs.ErrInvalidID)
//		return
//	}
//
//	// Проверка что запрос принадлежит пользователю
//	request, err := service.GetRequestByID(uint(requestID))
//	if err != nil {
//		HandleError(c, err)
//		return
//	}
//
//	if request.UserID != userID {
//		HandleError(c, errs.ErrPermissionDenied)
//		return
//	}
//
//	err = service.CancelRequest(uint(requestID))
//	if err != nil {
//		HandleError(c, err)
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"message": "Request cancelled successfully",
//	})
//}
