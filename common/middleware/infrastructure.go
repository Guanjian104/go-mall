package middleware

import (
    "github.com/Guanjian104/go-mall/common/util"
    "github.com/gin-gonic/gin"
)

func StartTrace() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        traceId := ctx.Request.Header.Get("traceid")
        pSpanId := ctx.Request.Header.Get("spanid")
        spanId := util.GenerateSpanID(ctx.Request.RemoteAddr)
        if traceId == "" { // 如果traceId 为空，证明是链路的发端，把它设置成此次的spanId，发端的spanId是root spanId
            traceId = spanId // trace 标识整个请求的链路, span则标识链路中的不同服务
        }
        ctx.Set("traceid", traceId)
        ctx.Set("spanid", spanId)
        ctx.Set("pspanid", pSpanId)
        ctx.Next()
    }
}