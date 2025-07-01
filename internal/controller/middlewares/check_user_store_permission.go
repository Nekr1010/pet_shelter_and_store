package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"pet_shelter_and_store/internal/errs"
	"pet_shelter_and_store/internal/service"
	"strconv"
)

const (
	StoreIDCtx = "storeID"
)

func CheckUserStorePermission(c *gin.Context) {
	storeIDStr := c.Param("id")
	storeID, err := strconv.Atoi(storeIDStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": errs.ErrValidationFailed.Error(),
		})
		return
	}

	store, err := service.GetStoreByID(uint(storeID))
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": errs.ErrRecordNotFound,
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": errs.ErrSomethingWentWrong,
		})
		return
	}

	userID := c.GetUint(UserIDCtx)
	userRole := c.GetString(UserRoleCtx)

	if userRole != "admin" || userID != store.OwnerID {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": errs.ErrPermissionDenied,
		})
		return
	}

	c.Set(StoreIDCtx, store.ID)

	c.Next()
}
