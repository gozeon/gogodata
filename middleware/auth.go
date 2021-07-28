package middleware

import (
	"gogodata/auth"
	"gogodata/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenString := utils.ExtractToken(c.Request)

		if len(tokenString) <= 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "no auth",
			})
			return
		}

		user, err := auth.ParseToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "unauthorized",
			})
			return
		}

		c.Set("user", user)

		c.Next()
	}
}
