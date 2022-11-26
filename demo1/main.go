package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default() // 创建一个默认路由引擎

	// 配置路由,可以配置多个
	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "值:%v ", "你好gin")
	})

	router.Run(":9092") // 启动一个服务,默认端口为8080
}
