package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func main() {
	router := gin.Default()

	// GET请求传值
	router.GET("/getUser", func(context *gin.Context) {
		uid := context.Query("uid")
		page := context.DefaultQuery("page", "1") // 没有就传默认值

		context.JSON(http.StatusOK, gin.H{
			"uid":  uid,
			"page": page,
		})
	})

	// POST请求传值
	router.POST("/addUser", func(context *gin.Context) {
		uid := context.PostForm("uid")
		username := context.PostForm("username")
		context.JSON(http.StatusOK, gin.H{
			"uid":      uid,
			"username": username,
		})
	})

	// 获取数据绑定到结构体
	router.GET("/", func(context *gin.Context) {
		user := &UserInfo{}
		error := context.ShouldBind(user)
		if error != nil {
			context.JSON(http.StatusOK, gin.H{
				"error": error.Error(),
			})
		} else {
			context.JSON(http.StatusOK, user)
		}

	})

	router.Run()
}
