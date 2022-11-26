package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("./templates/*")

	router.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "userAdd.html", nil)
	})

	router.POST("/admin/doUpload", func(context *gin.Context) {
		username := context.PostForm("username")
		file, error := context.FormFile("face") // 上传的单个文件
		fileName := file.Filename

		if error == nil {
			dist := path.Join("./static", fileName) // 上传目录
			context.SaveUploadedFile(file, dist)    // 保存文件
			context.JSON(http.StatusOK, gin.H{
				"success": username + "上传成功",
			})
		} else {
			context.JSON(http.StatusOK, gin.H{
				"error": username + "上传失败",
			})
		}

	})
	router.Run()
}
