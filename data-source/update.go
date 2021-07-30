package datasource

import (
	"gogodata/conf"
	"gogodata/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DoUpdate(c *gin.Context) {
	id := c.Param("id")
	var json model.DataSource

	if err := conf.DB.Where("id = ?", id).First(&json).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	result := conf.DB.Save(&json)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": result.Error.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "no group"})
		return
	}

	c.JSON(http.StatusOK, json)
}
