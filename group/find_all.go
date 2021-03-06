package group

import (
	"gogodata/conf"
	"gogodata/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func DoFindAll(c *gin.Context) {
	var json []model.Group
	search := c.DefaultQuery("search", "")
	var searchLike strings.Builder
	searchLike.WriteString("%")
	searchLike.WriteString(search)
	searchLike.WriteString("%")
	result := conf.DB.Where("create_user = ?", c.MustGet("user").(string)).Where("name LIKE ?", searchLike.String()).Or("description LIKE ?", searchLike.String()).Order("updated_at desc").Find(&json)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, json)
}
