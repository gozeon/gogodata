package middleware

import (
	"gogodata/conf"
	"gogodata/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CollectMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var stats model.Stats
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()
		// 结束时间
		endTime := time.Now()
		// 执行时间
		stats.ResponseTime = endTime.Sub(startTime).String()
		// 状态码
		stats.ResponseCode = c.Writer.Status()

		if stats.ResponseCode == 200 {

			// 请求方式
			stats.RequestMethod = c.Request.Method

			// 请求路由
			stats.RequestURI = c.Request.RequestURI

			// 请求IP
			stats.RequestIP = c.ClientIP()

			// groupID
			stats.GroupID = c.MustGet("show_data_group_id").(uint)

			// dataSourceID
			val, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			stats.DataSourceID = uint(val)

			conf.DB.Create(&stats)

		}

	}
}
