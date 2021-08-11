package datasource

import (
	"gogodata/conf"
	"gogodata/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DoFindAll(c *gin.Context) {
	var json []model.DataSource
	groupId := c.DefaultQuery("groupId", "")

	result := conf.DB.Where("group_id = ?", groupId).Order("updated_at desc").Find(&json)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, json)
}
