package stats

import (
	"gogodata/conf"
	"gogodata/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GourpCount struct {
	GroupID int
	Count   int
}

func DoFindAll(c *gin.Context) {
	var total int64
	var allGroup int64
	var allDataSource int64
	var result *gorm.DB
	ap := make([]GourpCount, 0)

	result = conf.DB.Model(&model.Stats{}).Select("group_id, count(*) as count").Group("group_id").Scan(&ap)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": result.Error.Error()})
		return
	}

	result = conf.DB.Model(&model.Stats{}).Count(&total)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": result.Error.Error()})
		return
	}

	result = conf.DB.Model(&model.Group{}).Count(&allGroup)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": result.Error.Error()})
		return
	}

	result = conf.DB.Model(&model.DataSource{}).Count(&allDataSource)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"detail": ap,
		"total": gin.H{
			"stats":        total,
			"groups":       allGroup,
			"data_sources": allDataSource,
		},
	})
}
