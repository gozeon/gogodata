package group

import (
	"gogodata/conf"
	"gogodata/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DoDelete(c *gin.Context) {
	id := c.Param("id")

	result := conf.DB.Delete(&model.Group{}, id)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": result.Error.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "no group"})
		return
	}

	c.Status(http.StatusOK)
}
