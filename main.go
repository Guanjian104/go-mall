package main

import (
	"github.com/Guanjian104/go-mall/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.New()
	// TODO: 后面会把应用日志统一收集到文件， 这里根据运行环境判断, 只在dev环境下才使用gin.Logger()输出信息到控制台
	g.Use(gin.Logger(), gin.Recovery())
	g.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	g.GET("/config-read", func(c *gin.Context) {
		app := config.App
		database := config.Database
		c.JSON(http.StatusOK, gin.H{
			"env": 		app.Env,
			"type":     database.Type,
			"max_life": database.MaxLifeTime,
		})
	})
	g.Run(":8080")
}
