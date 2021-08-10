package auth

import (
	"gogodata/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DoLogin(c *gin.Context) {
	var json model.Login
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if json.User != "admin" && json.Password != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	token, err := CreateToken("admin")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "gen token fail"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "user": json.User})
}
