package datasource

import (
	"gogodata/conf"
	"gogodata/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DoDelete(c *gin.Context) {
	id := c.Param("id")

	result := conf.DB.Delete(&model.DataSource{}, id)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": result.Error.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "no data source"})
		return
	}

	c.Status(http.StatusOK)
}
