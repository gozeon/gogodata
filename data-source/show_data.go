package datasource

import (
	jsonU "encoding/json"
	"gogodata/conf"
	"gogodata/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DoShowData(c *gin.Context) {
	id := c.Param("id")
	var json model.DataSource

	result := conf.DB.Where("id = ?", id).First(&json)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": result.Error.Error()})
		return
	}

	jsonMap := make(map[string]interface{})

	err := jsonU.Unmarshal([]byte(json.Data), &jsonMap)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, jsonMap)
}