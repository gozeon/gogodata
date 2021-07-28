package main

import (
	"fmt"
	"gogodata/auth"
	"gogodata/conf"
	_ "gogodata/conf"
	"gogodata/middleware"
	"gorm.io/gorm"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type User struct {
	gorm.Model
	Name string
	Age int
}
func main() {
	sqlDB := conf.InitDB()
	defer sqlDB.Close()

	conf.DB.AutoMigrate(&User{})

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
