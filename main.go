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

    logger.Init()

    g.Use(gin.Logger(), middleware.StartTrace(), middleware.LogAccess(), middleware.GinPanicRecovery())
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
        logger.Info(c, "logger test", "key", "keyName", "val", 3)
        c.JSON(http.StatusOK, gin.H{
            "status": "ok",
        })
    })

    g.POST("/access-log-test", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "status": "ok",
        })
    })

    g.GET("/panic-log-test", func(c *gin.Context) {
        var a map[string]string
        // TODO: panic 测试接口，需要的时候在这里 Panic 一下
        // a["k"] = "v"
        c.JSON(http.StatusOK, gin.H{
            "status": "ok",
            "data": a,
        })
    })

    g.Run(":8080")
}
