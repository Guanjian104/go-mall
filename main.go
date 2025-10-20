package main

import (
	"net/http"

	"github.com/Guanjian104/go-mall/common/logger"
	"github.com/Guanjian104/go-mall/common/middleware"
	"github.com/Guanjian104/go-mall/config"

	"github.com/gin-gonic/gin"
)

func main() {
    g := gin.New()

    g.Use(gin.Logger(), middleware.StartTrace())
    g.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "pong",
        })
    })

    g.GET("/config-read", func(c *gin.Context) {
        database := config.Database

        c.JSON(http.StatusOK, gin.H{
            "type":     database.Type,
            "max_life": database.MaxLifeTime,
        })
    })

    g.GET("/logger-test", func(c *gin.Context) {
        logger.New(c).Info("logger test", "key", "keyName", "val", 2)
        c.JSON(http.StatusOK, gin.H{
            "status": "ok",
        })
    })

    g.Run(":8080")
}
