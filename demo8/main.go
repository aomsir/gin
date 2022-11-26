package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	// 设置Cookie
	router.GET("/", func(context *gin.Context) {
		context.SetCookie("user", "Aomsir", 3600, "/", "localhost", false, false)
		context.JSON(http.StatusOK, gin.H{
			"success": "HelloWorld",
		})
	})

	// 获取Cookie
	router.GET("/getCookie", func(context *gin.Context) {
		cookieValue, _ := context.Cookie("user")
		context.JSON(http.StatusOK, gin.H{
			"cookie": cookieValue,
		})
	})

	// 删除Cookie
	router.GET("/delCookie", func(context *gin.Context) {
		context.SetCookie("user", "", -1, "/", "localhost", false, false)
		context.JSON(http.StatusOK, gin.H{
			"success": "删除成功",
		})
	})

	router.Run()
}
