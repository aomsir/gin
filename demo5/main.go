package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 时间戳转换成日期
func UnixToTime(timestamp int) string {
	fmt.Println(timestamp)
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

// 抽取的中间件
func initMiddleware(context *gin.Context) {
	fmt.Println("我是一个中间件 - 1")

	context.Next() // 调用该请求的剩余处理程序,剩余的程序处理完成才会继续执行这个中间件

	fmt.Println("我是一个中间件 - 2")
}

func main() {
	router := gin.Default()

	router.Use(initMiddleware) // 全局中间件

	router.GET("/", initMiddleware, func(context *gin.Context) {
		fmt.Println("我是首页 - ")
		context.JSON(http.StatusOK, gin.H{
			"success": "HelloWorld",
		})
	})

	router.Run()
}
