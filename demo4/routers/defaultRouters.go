package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func DefaultRouterInit(router *gin.Engine) {
	defaultRouters := router.Group("/")
	{
		defaultRouters.GET("/", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"success": "我是一个首页API",
			})
		})

		defaultRouters.GET("/list", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"success": "HelloWorld",
			})
		})
	}
}
