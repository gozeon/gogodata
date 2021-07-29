package group

import (
	"gogodata/conf"
	"gogodata/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DoFindById(c *gin.Context) {
	id := c.Param("id")
	var json model.Group

	result := conf.DB.Where("id = ?", id).First(&json)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, json)
}
