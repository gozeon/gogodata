package group

import (
	"gogodata/conf"
	"gogodata/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DoFindAll(c *gin.Context) {
	var json []model.Group

	result := conf.DB.Find(&json)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, json)
}
