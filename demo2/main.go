package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/json1", func(context *gin.Context) {
		context.JSON(http.StatusOK, map[string]interface{}{
			"success": "123",
		})
	})

	// 或者
	router.GET("/json2", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"success": "123",
		})
	})

	router.GET("/json3", func(context *gin.Context) {
		article := &Article{
			Title:   "我是标题",
			Desc:    "我是描述",
			Content: "我是内容",
		}

		context.JSON(http.StatusOK, gin.H{
			"article": article,
		})
	})

	router.GET("/news", func(context *gin.Context) {
		context.HTML(http.StatusOK, "news.html", gin.H{
			"title": "来自后台的新闻数据",
			"score": 89,
		})
	})

	router.Run()
}
