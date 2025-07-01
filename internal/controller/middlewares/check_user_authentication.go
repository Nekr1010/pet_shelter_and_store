package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pet_shelter_and_store/utils"
	"strings"
)

const (
	AuthorizationHeader = "Authorization"
	UserIDCtx           = "userID"
	UserRoleCtx         = "userRole"
)

func CheckUserAuthentication(c *gin.Context) {
	header := c.GetHeader(AuthorizationHeader)

	if header == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "empty auth header",
		})
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "invalid auth header",
		})
		return
	}

	if len(headerParts[1]) == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "token is empty",
		})
		return
	}

	accessToken := headerParts[1]

	claims, err := utils.ParseToken(accessToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.Set(UserIDCtx, claims.UserID)
	c.Set(UserRoleCtx, claims.UserRole)

	c.Next()
}
