package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DoMe(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"user": c.MustGet("user").(string),
	})
}
