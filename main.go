package main

import (
	"fmt"
	"net/http"

	"gogodata/auth"
	_ "gogodata/conf"
	"gogodata/middleware"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
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
	}
	r.Run(fmt.Sprintf(":%d", viper.GetInt("port")))
}
