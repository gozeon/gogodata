package main

import (
	"fmt"
	"gogodata/auth"
	"gogodata/conf"
	_ "gogodata/conf"
	dataSource "gogodata/data-source"
	"gogodata/group"
	"gogodata/middleware"
	"gogodata/model"
	"gogodata/stats"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	sqlDB := conf.InitDB()
	defer sqlDB.Close()

	conf.DB.AutoMigrate(&model.User{}, &model.Group{}, &model.DataSource{}, &model.Stats{})

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	v1 := r.Group(fmt.Sprintf("/%s", viper.Get("APP_NAME")))
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
				"appName": viper.GetString("APP_NAME"),
				"port":    viper.GetInt("port"),
			})
		})

		authRouter := v1.Group("/auth")
		{
			authRouter.POST("/login", auth.DoLogin)
			authRouter.GET("/me", middleware.AuthMiddleware(), auth.DoMe)
		}

		groupRouter := v1.Group("/group")
		groupRouter.Use(middleware.AuthMiddleware())
		{
			groupRouter.POST("/", group.DoCreate)
			groupRouter.GET("/", group.DoFindAll)
			groupRouter.GET("/:id", group.DoFindById)
			groupRouter.DELETE("/:id", group.DoDelete)
			groupRouter.PUT("/:id", group.DoUpdate)
		}

		dataSourceRouter := v1.Group("/ds")
		// 渲染data，勿需权限
		dataSourceRouter.GET("/info/:id", middleware.CollectMiddleware(), dataSource.DoShowData)
		dataSourceRouter.Use(middleware.AuthMiddleware())
		{
			dataSourceRouter.POST("/", dataSource.DoCreate)
			dataSourceRouter.GET("/", dataSource.DoFindAll)
			dataSourceRouter.GET("/:id", dataSource.DoFindById)
			dataSourceRouter.DELETE("/:id", dataSource.DoDelete)
			dataSourceRouter.PUT("/:id", dataSource.DoUpdate)
		}

		stateRouter := v1.Group("/stats")
		{
			stateRouter.GET("/", stats.DoFindAll)
		}
	}
	r.Run(fmt.Sprintf(":%d", viper.GetInt("port")))
}
